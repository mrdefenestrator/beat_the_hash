#!/usr/bin/env python3
'''Script for mining for hash values
'''
import os
import sys
import multiprocessing

import state
import common


HEX_HASH = '''\
5b4da95f5fa08280fc9879df44f418c8f9f12ba424b7757de02bbdfbae0d4c4fdf9317c80cc5fe\
04c6429073466cf29706b8c25999ddd2f6540d4475cc977b87f4757be023f19b8f4035d7722886\
b78869826de916a79cf9c94cc79cd4347d24b567aa3e2390a573a373a48a5e676640c79cc70197\
e1c5e7f902fb53ca1858b6
'''
BIN_HASH = bytes(bytearray.fromhex(HEX_HASH))

USAGE = '''mine.py <n_values> <n_processes>

    n_values     number of values to mine
    n_processes  number of processes to use for mining
'''


def worker(start, n_values, best_hamming, best_value, proc_num, results):
    '''Worker process for generating, hashing & checking values

    Args:
        start (int): value to start with
        n_values (int): number of values to test
        best_hamming (int): previous best hamming distance
        best_value (bytes): previous best value
        proc_num (int): process number
        results (dict): shared variable for providing results
    '''
    last_value = None

    for guess in common.gen_guess(n_values, start):
        # Iterate guesses
        value = common.list_to_unicode(common.to_base(guess, sys.maxunicode))

        # Hash the guess and check resulting distance
        try:
            value_bytes = value.encode()
        except ValueError:
            # Unable to encode string, skip it
            continue
        hamming = common.calc_hamming(BIN_HASH, common.hash_it(value_bytes))

        if hamming < best_hamming:
            # Maximize
            print('New Best', hamming, value)
            best_hamming = hamming
            best_value = value

        last_value = value

    results[proc_num] = (best_hamming, best_value, last_value)


def str_to_int(value):
    '''Converts unicode string to integer
    
    Args:
        value (str): value to convert 

    Returns:
        int: converted value
    '''
    return common.from_base(common.unicode_to_list(value), sys.maxunicode)


def int_to_str(value):
    '''Converts integer to unicode string
    Args:
        value (int): value to convert

    Returns:
        str: converted value
    '''
    return common.list_to_unicode(common.to_base(value, sys.maxunicode))


def main(state_path, n_values, n_processes):
    '''Do some mining based on persistent state
    '''
    if not os.path.isfile(state_path):
        # Generate initial state
        my_state = state.State(1024, '', '')
    else:
        # Load persisted state
        my_state = state.State.load(state_path)

    # Create shared variable
    manager = multiprocessing.Manager()
    results = manager.dict()

    # Divide work & start multiple workers
    start_value = str_to_int(my_state.value)
    jobs = []
    for i in range(n_processes):
        start = (n_values // n_processes) * i + start_value
        process = multiprocessing.Process(target=worker, args=(
            start,
            n_values // n_processes,
            my_state.hamming,
            my_state.value,
            i,
            results
        ))
        jobs.append(process)
        process.start()

    # Wait for all workers to complete
    for job in jobs:
        job.join()

    # Determine best result
    for result in results.values():
        hamming, value, _ = result
        if hamming < my_state.hamming:
            my_state.hamming = hamming
            my_state.value = value

    my_state.last_value = int_to_str(start_value + n_values)

    print('Best this run', my_state.hamming, my_state.value)

    # Persist state
    my_state.save(state_path)


if __name__ == '__main__':
    if len(sys.argv) < 3:
        print(USAGE)
        exit(-1)

    main(state.STATE_PATH, int(sys.argv[1]), int(sys.argv[2]))
