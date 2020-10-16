package main

import (
	"fmt"
	"strings"
)

type Foo struct {
	A string
}

// implements stringer Interface
func (f Foo) String() string {
	return strings.ToUpper(f.A)
}

func main() {
	fo := Foo{"another Foo"}
	fmt.Println(fo)

	foo := &fo
	fmt.Println(foo)
}
