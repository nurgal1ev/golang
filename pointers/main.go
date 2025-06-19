package main

import "fmt"

func main() {
	number := 5
	num(number)
	//fmt.Println(number)
	pointer := &number
	point(pointer)
	fmt.Println(number)
	nilPointer()
}

func point(n *int) {
	fmt.Println("указатель на тип данных:", n)
	fmt.Println("указатель на значение переменной:", *n)
	*n = 10
}

func num(n int) {
	n = 10
}

func nilPointer() {
	number := 2
	fmt.Println(number)

	var ptr *int
	fmt.Println(ptr)
}
