package main

import "fmt"

type SuperString struct {
	A string
}

func (f *SuperString) Add(s string) {
	f.A = f.A + s
}

func main() {
	val := SuperString{"super"}
	val.Add(" string")
	fmt.Println(val) // => compiler error?

	ptr := &val
	ptr.Add("!")
	fmt.Println(ptr)
}
