package main

import (
	"fmt"
	"net/url"
	"net/http"
	"os"
	"mine/common"
	"mine/state"
	"io/ioutil"
)

var beatthehash_uri string = "http://beatthehash.com/hash"
var usage string = "post.py <username> <state_path>\n\n" +
		"username     username to post to server with\n" +
		"state_path   path to YAML state file\n"

// Post value to the server
func post_it(server_url string, username string, value string) {
	data := url.Values{
		"username": {username},
		"value":    {value},
	}

	resp, err := http.PostForm(server_url, data)
	common.Check(err)

	body_bytes, err := ioutil.ReadAll(resp.Body)
	common.Check(err)

	fmt.Println(data)
	fmt.Println(resp.StatusCode)
	fmt.Println(string(body_bytes))
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		os.Exit(-1)
	}

	username := os.Args[1]
	state_path := os.Args[2]

	last_state := state.State{}
	last_state.Load(state_path)
	post_it(beatthehash_uri, username, last_state.Value)
}
