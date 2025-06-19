package main

import "fmt"

func main() {
	fmt.Println("main start")
	defer func() {
		fmt.Println("main end")
	}()
	hello()
}

func hello() {
	fmt.Println("hello start")
	defer func() {
		fmt.Println("hello end")
	}()

	fmt.Println("hello func 1")
	fmt.Println("hello func 2")
	fmt.Println("hello func 3")
}
