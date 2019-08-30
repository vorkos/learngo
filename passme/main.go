package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	usage   = "usage - [user] [pass]"
	errMsg  = "Wrong credentials"
	goodMsg = "Access granted for user %s\n"
)

func main() {
	user := "me"
	pass := "no"
	args := os.Args
	if len(args) < 3 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	if args[1] == user && args[2] == pass {
		fmt.Printf(goodMsg, os.Args[1])
	} else {
		fmt.Println(errMsg)
	}

}
