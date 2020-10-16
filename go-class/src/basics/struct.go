package main

import "fmt"

type T struct {
	a, b int
}

func main() {
	var t0 T = T{1, 2}
	t0.b = 3

	t1 := T{ // type inference
		a: 3,
		b: 4,
	}
	t1.a = 2
	fmt.Printf("%+v, %+v", t0, t1)
}
