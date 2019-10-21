package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const usage = `
	Give me the size of the table`

func main() {
	var tableSize int64

	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	tableSize, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("Wrong size")
		return
	}
	fmt.Printf("%5s ", "X")

	for i := 1; i <= int(tableSize); i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Printf("\n \n")

	for i := 1; i <= int(tableSize); i++ {
		fmt.Printf("%5d ", i)
		for j := 1; j <= int(tableSize); j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
