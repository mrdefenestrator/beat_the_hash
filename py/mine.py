#!/usr/bin/env python3
'''Script for mining for hash values
'''
import os
import sys
import math
import skein
import multiprocessing

import state


HEX_HASH = '''\
5b4da95f5fa08280fc9879df44f418c8f9f12ba424b7757de02bbdfbae0d4c4f\
df9317c80cc5fe04c6429073466cf29706b8c25999ddd2f6540d4475cc977b87\
f4757be023f19b8f4035d7722886b78869826de916a79cf9c94cc79cd4347d24\
b567aa3e2390a573a373a48a5e676640c79cc70197e1c5e7f902fb53ca1858b6\
'''
BIN_HASH = bytes(bytearray.fromhex(HEX_HASH))

USAGE = '''mine.py <n_values> <n_processes>

    n_values     number of values to mine
    n_processes  number of processes to use for mining
'''


def to_list(value):
    '''Converts unicode str to list of integers
    
    Args:
        value (str): unicode string to convert
        
    Returns:
        list[int]: list of ints
    '''
    return list([ord(_) for _ in value])


def to_unicode(value):
    '''Converts list of integers to unicode str
    
    Args:
        value (list[int]): list of ints to convert
        
    Returns:
        str: unicode string
    '''
    return ''.join([chr(_) for _ in value])


def to_base(value, base):
    '''Converts integer to a list of ints of base x
    
    Args:
        value (int): number to convert
        base (int): base to convert number to
        
    Returns:
        list[int]: list of integers of specified base
    '''
    result = []
    for i in range(0, int(math.ceil(math.log(value, base)))):
        result.append((value // base ** i) % base)

    return result[::-1]


def from_base(value, base):
    '''Converts to integer from a list of ints of base x

    Args:
        value (list[int]): list if integers to convert
        base (int): base to convert number from

    Returns:
        int: converted number
    '''
    result = 0
    for i, num in enumerate(value[::-1]):
        result += num * base ** i

    return result


def gen_guess(n, start=None):
    '''Generator for creating value guesses - should be improved to iterate 
    through n bytes of output then each possible value for that n bytes, this
    currently skips guesses with any amount of forward zero padding.

    Args:
        n (int): number of values to generate
        start (int): starting value to increment from

    Yields:
        int: next value
    '''
    if start is None:
        start = 0
    num = start
    current_n = 0

    while current_n < n:
        yield num
        num += 1
        current_n += 1


def int_to_bytes(num):
    '''Generator to convert an integer to bytes

    Args:
        num (int): number to convert

    Yields:
        bytes: bytes for the number
    '''
    while num > 0:
        yield num & 0xff
        num = num >> 8


def hash_it(value):
    '''Get the skein1024 hash of the value

    Args:
        value (bytes): value to hash

    Returns:
        bytes: the hash digest
    '''
    digest = skein.skein1024()
    digest.update(value)
    return digest.digest()


def hamming_it(truth, guess):
    '''Return the bitwise hamming distance

    Args:
        truth (bytes): truth to compare against
        guess (bytes): guess to compare with truth

    Returns:
        int: the hamming distance
    '''
    dist = 0
    for ub, db in zip(truth, guess):
        diff = ub ^ db
        for x in range(8):
            if diff & 0b1:
                dist += 1
            diff = diff >> 1

    return dist


def worker(start, n, best_hamming, best_value, proc_num, results):
    '''Worker process for generating, hashing & checking values

    Args:
        start (int): value to start with
        n (int): number of values to test
        best_hamming (int): previous best hamming distance
        best_value (bytes): previous best value
        proc_num (int): process number
        results (dict): shared variable for providing results
    '''
    last_value = None
    for guess in gen_guess(n, start):
        # Iterate guesses
        value = bytes(int_to_bytes(guess))

        # Hash the guess and check resulting distance
        hamming = hamming_it(BIN_HASH, hash_it(value))

        if hamming < best_hamming:
            # Maximize
            print('New Best', hamming, value)
            best_hamming = hamming
            best_value = value

        last_value = value

    results[proc_num] = (best_hamming, best_value, last_value)


def main(state_path, n, n_processes):
    '''Do some mining based on persistent state
    '''
    if not os.path.isfile(state_path):
        # Generate initial state
        last_state = state.State(1024, bytes(), bytes(1))
    else:
        # Load persisted state
        last_state = state.State.load(state_path)

    best_hamming = last_state.hamming
    best_value = last_state.value
    last_value = last_state.last_guess

    # Create shared variable
    manager = multiprocessing.Manager()
    results = manager.dict()

    # Divide work & start multiple processes
    start_value = int.from_bytes(last_value, byteorder='big')
    jobs = []
    for i in range(n_processes):
        start = (n // n_processes) * i + start_value
        process = multiprocessing.Process(
            target=worker, args=(
                start, n // n_processes, best_hamming, best_value, i, results))
        jobs.append(process)
        process.start()

    # Wait for all jobs to complete
    for job in jobs:
        job.join()

    # Determine best result
    for result in results.values():
        hamming, value, _ = result
        if hamming < best_hamming:
            best_hamming = hamming
            best_value = value

    last_value = results[n_processes - 1][2]

    print('Best this run', best_hamming, best_value)

    new_state = state.State(best_hamming, best_value, last_value)
    new_state.save(state_path)


if __name__ == '__main__':
    if len(sys.argv) < 3:
        print(USAGE)
        exit(-1)

    main(state.STATE_PATH, int(sys.argv[1]), int(sys.argv[2]))
