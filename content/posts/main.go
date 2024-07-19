package main

import "fmt"

func main() {
	x := map[string]int{}
	fmt.Printf("x: %v\n", len(x))
	x["chat"] = 2
	fmt.Printf("x: %v\n", x)
}
