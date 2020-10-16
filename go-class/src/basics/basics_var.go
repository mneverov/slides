package main

import "fmt"

const c int = 1
const (
	one     = 1
	two int = 2
)

var x int // package global
var initializedVar = 42

func main() {
	var y, z int // function local
	x = 1 + y

	const greeting string = "Hello, World!"
	fmt.Println(x, y, z, greeting)
}
