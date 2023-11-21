package main

import (
	"fmt"
	"os"

	"github.com/Xhofe/wopan-sdk-go"
)

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func main() {
	actions := []string{"EmptyRecycleData", "AppQueryUser"}
	if len(os.Args) < 2 || !IsContain(actions, os.Args[1]) {
		fmt.Printf("please action method: %s", actions)
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("please input access token")
		return
	}

	action := os.Args[1]
	access_token := os.Args[2]

	w := wopan.DefaultWithAccessToken(access_token)
	if action == "EmptyRecycleData" {
		err := w.EmptyRecycleData()
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		fmt.Println("success")
	} else if action == "AppQueryUser" {
		res, err := w.AppQueryUser()
		if err != nil {
			fmt.Printf("AppQueryUser() error = %v", err)
		} else {
			fmt.Printf("AppQueryUser() = %+v", res)
		}
	} else {
		fmt.Println("[ERROR] not suppport action")
	}
}
