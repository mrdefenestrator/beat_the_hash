#!/usr/bin/env python3
'''Script for posting best result from state to site
'''
import sys
import requests

import state


BEATTHEHASH_URI = 'http://beatthehash.com/hash'

USAGE = '''post.py <username> <state_path>

    username     username to post to server with
    state_path   path to YAML state file
'''


def post_it(uri, username, state_path):
    '''Posts the value to the server

    Args:
        uri (str): uri of the server
        username (str): username to post with
        state_path (str): path to yaml state file
    '''
    last_state = state.State.load(state_path)

    data = {
        'username': username,
        'value': last_state.value
    }
    print(data)
    result = requests.post(uri, data=data)
    print(result.status_code)
    print(result.text)


def main():
    '''Perform post of best value from state to site
    '''
    if len(sys.argv) != 3:
        print(USAGE)
        exit(-1)

    username = sys.argv[1]
    state_path = sys.argv[2]

    post_it(BEATTHEHASH_URI, username, state_path)


if __name__ == '__main__':
    main()
