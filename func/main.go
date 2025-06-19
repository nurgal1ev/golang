package main

import "fmt"

func main() {
	fmt.Println("до вызова функции")
	hello()
	square(5)
	helloUser("ivan")
	sum(5, 7)
	fmt.Println("после вызова функции")
}

func hello() {
	fmt.Println("функция hello")
}

func square(x int) {
	fmt.Println("функция square: x * x =", x*x)
}

func helloUser(name string) {
	fmt.Println(" hello,", name)
}

func sum(a int, b int) int {
	result := a + b
	return result
}
