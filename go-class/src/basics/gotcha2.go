package main

import "fmt"

type SuperString struct {
	A string
}

func (f *SuperString) Add(s string) {
	f.A = f.A + s
}

func main() {
	fo := SuperString{"super"}
	fo.Add(" string")
	fmt.Println(fo) // => compiler error?

	foo := &fo
	foo.Add("!")
	fmt.Println(foo)
}
