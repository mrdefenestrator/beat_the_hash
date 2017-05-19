'''Unit tests for verifying operation
'''
import unittest

from mine import *


class HashTests(unittest.TestCase):
    def test_ham(self):
        '''Test the operation of the hamming distance calculation
        '''
        vectors = [
            (
                bytes([0x10]),
                bytes([0x01]),
                2
            ),
            (
                bytes([0x01]),
                bytes([0x01]),
                0
            ),
            (
                bytes([0x00, 0x10]),
                bytes([0x00, 0x01]),
                2
            ),
            (
                bytes([0x10, 0x10]),
                bytes([0x01, 0x01]),
                4
            ),
            (
                bytes([0x00] * 128),
                bytes([0xff] * 128),
                1024
            ),
            (
                bytes([0xff] * 128),
                bytes([0xff] * 128),
                0
            ),
        ]

        for vector in vectors:
            result = hamming_it(vector[0], vector[1])
            self.assertEqual(vector[2], result)

    def test_values(self):
        '''Check values against scores from beatthehash.com to determine proper
        binary encoding
        '''
        tests = [
            (bytes('AA==', 'utf-8'), 525),
            (bytes('AQ==', 'utf-8'), 511),
            (bytes('asdf', 'utf-8'), 500),
            (bytes('%00', 'utf-8'), 510),
            (bytes('%20', 'utf-8'), 513),
            (bytes([0x00]), 539),
            (bytes([0x01]), 525),
            # TODO: use percent encoding when submitting to server
        ]

        for test in tests:
            hamming = hamming_it(BIN_HASH, hash_it(test[0]))
            # self.assertEqual(test[1], hamming)
            print(test[0], test[1], hamming)

        self.assertEqual(HEX_HASH, BIN_HASH.hex())
