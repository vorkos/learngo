package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = `
usage - [user] [pass]
`

func main() {
	user := "me"
	pass := "no"
	args := os.Args
	if len(args) < 3 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	if args[1] == user && args[2] == pass {
		fmt.Printf("Access granted for user %s\n", os.Args[1])
	} else {
		fmt.Println("Wrong credentials")
	}

}
