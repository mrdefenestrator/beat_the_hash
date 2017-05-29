'''Module for containing functions that might have general purpose relevance
'''
import math
import skein


def unicode_to_list(value):
    '''Converts unicode str to list of integers

    Args:
        value (str): unicode string to convert

    Returns:
        list[int]: list of ints
    '''
    return list([ord(_) for _ in value])


def list_to_unicode(value):
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
    n_items = int(math.ceil(math.log(value, base))) if value else 1
    for i in range(0, n_items):
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
    '''Generator for creating value guesses

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


def calc_hamming(truth, guess):
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
