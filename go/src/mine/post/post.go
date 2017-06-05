package main

import (
	"fmt"
	"io/ioutil"
	"mine/common"
	"mine/state"
	"net/http"
	"net/url"
	"os"
)

var usage = "post.py <username> <state_path>\n\n" +
	"username     username to post to server with\n" +
	"state_path   path to YAML state file\n"

var beatTheHashURI = "http://beatthehash.com/hash"

// postIt posts value to the server
func postIt(serverURL string, username string, value string) {
	data := url.Values{
		"username": {username},
		"value":    {value},
	}

	resp, err := http.PostForm(serverURL, data)
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
	postIt(beatTheHashURI, username, lastState.Value)
}
