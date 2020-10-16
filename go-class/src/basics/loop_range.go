package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5} // define s slice

	// index loop
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] = %d\n", i, s[i])
	}

	// range loop
	for idx, item := range s {
		fmt.Printf("At index %d, there is item %d\n", idx, item)
	}
}
