package goversion

import (
	"fmt"
	"runtime"
)

// Print func prints go version
func Print() {
	fmt.Println(runtime.Version())
}
