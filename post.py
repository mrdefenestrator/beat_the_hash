#!/usr/bin/env python3
'''Script for posting best result from state to site
'''
import requests

import state


BEATTHEHASH_URI = 'http://beatthehash.com/hash'
USERNAME = 'TEST64'


def post_it(uri, username, value):
    '''Posts the value to the server

    Args:
        uri (str): uri of the server
        username (str): username to post with
        value (bytes): value to post
    '''
    data = {
        'username': username,
        'value': value
    }
    print(data)
    result = requests.post(uri, data=data)
    print(result.status_code)
    print(result.text)


def main():
    '''Perform post of best value from state to site
    '''
    last_state = state.State.load(state.STATE_PATH)
    post_it(BEATTHEHASH_URI, USERNAME, last_state.value)


if __name__ == '__main__':
    main()
