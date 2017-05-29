'''Module for saving and loading state
'''
import os
import yaml


THIS_PATH = os.path.dirname(os.path.abspath(__file__))


class State(object):
    def __init__(self, hamming=None, value=None, last_value=None):
        '''
        Args:
            hamming (int): the best hamming distance
            value (str): the value that produced the best distance
            last_value (str): the last value checked
        '''
        self.hamming = hamming
        self.value = value
        self.last_value = last_value

    def save(self, path):
        '''Save the result to the path

        Args:
            path (str): path to the yaml results file
        '''
        result = {
            'hamming':    self.hamming,
            'value':      self.value,
            'last_value': self.last_value,
        }

        with open(path, 'w') as fp:
            fp.write(yaml.dump(result))

    @classmethod
    def load(cls, path):
        '''Load the result from the path

        Args:
            path (str): path to the yaml results file

        Returns:
            State: the loaded result
        '''
        with open(path, 'r') as fp:
            result = yaml.load(fp.read())

        return cls(
            result['hamming'],
            result['value'],
            result['last_value'],
        )
