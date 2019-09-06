package main

import (
	"fmt"
	"time"
)

func main() {
	switch t := time.Now().Hour(); {
	case t > 5 && t < 11:
		fmt.Println("Good morning")
	case t > 11 && t < 18:
		fmt.Println("Good day")
	case t > 18:
		fmt.Println("Good evening")
	default:
		fmt.Println("Good night")
	}
}
