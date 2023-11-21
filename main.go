package main

import (
	"fmt"
	"os"

	"github.com/Xhofe/wopan-sdk-go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please input access token")
		return
	}

	w := wopan.DefaultWithAccessToken(os.Args[1])
	err := w.EmptyRecycleData()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println("success")
}
