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

var usage string = "post.py <username> <state_path>\n\n" +
		"username     username to post to server with\n" +
		"state_path   path to YAML state file\n"

var beatTheHashUri string = "http://beatthehash.com/hash"

// Post value to the server
func postIt(serverUrl string, username string, value string) {
	data := url.Values{
		"username": {username},
		"value":    {value},
	}

	resp, err := http.PostForm(serverUrl, data)
	common.Check(err)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	common.Check(err)

	fmt.Println(data)
	fmt.Println(resp.StatusCode)
	fmt.Println(string(bodyBytes))
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		os.Exit(-1)
	}

	username := os.Args[1]
	statePath := os.Args[2]

	lastState := state.State{}
	lastState.Load(statePath)
	postIt(beatTheHashUri, username, lastState.Value)
}
