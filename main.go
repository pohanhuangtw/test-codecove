package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul2(a int, b int) int {
	return a * b
}

func div11(a int, b int) int {
	return a / b
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(sub(1, 2))
	fmt.Println(mul2(1, 2))
	fmt.Println(div11(1, 2))
}
