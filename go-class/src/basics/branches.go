package main

import (
	"fmt"
	"runtime"
)

func main() {
	const s string = "🤔🤨"
	if len(s) > 0 {
		fmt.Printf("str: %q, len(str)=%d\n", s, len(s))
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%q.", os)
	}
}
