package main

import "fmt"

func main() {
	multilineString := `first
second\n
third`

	fmt.Println(multilineString)

	const s string = "🤔🤨"

	for index, runeValue := range s {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}
