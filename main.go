package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	http "net/http"
	"os"
)

func main() {
	n := getUserGithubName()

	bodyByte := getUserInfo(n)

	reposArr := parseBody(bodyByte)

	createTable(reposArr)
}

func getUserGithubName() string {
	var n string

	fmt.Println("Enter your Github name: ")
	fmt.Scan(&n)

	return n
}

func getUserInfo(n string) []byte {
	resp, err := http.Get("https://api.github.com/users/" + n + "/repos")

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 404 {
		fmt.Printf("No user was found with given name: %v\n", n)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

func parseBody(b []byte) []Repo {
	var repos []Repo

	err := json.Unmarshal(b, &repos)

	if err != nil {
		fmt.Println("error:", err)
	}

	return repos
}
