'''Module for saving and loading state
'''
import os
import yaml


THIS_PATH = os.path.dirname(os.path.abspath(__file__))
STATE_PATH = os.path.join(THIS_PATH, 'state.yml')


class State(object):
    def __init__(self, hamming, value, last_guess):
        '''
        Args:
            hamming (int): the hamming distance
            value (str): the value that produced the distance
            last_guess (str): the last value guessed
        '''
        self.hamming = hamming
        self.value = value
        self.last_guess = last_guess

    def save(self, path):
        '''Save the result to the path

        Args:
            path (str): path to the yaml results file
        '''
        result = {
            'hamming':    self.hamming,
            'value':      self.value,
            'last_guess': self.last_guess,
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
            result['last_guess'],
        )
