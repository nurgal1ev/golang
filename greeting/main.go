package main

import (
	"fmt"
	"golanggit/greeting"
)

func main() {
	fmt.Println("hello")
	greeting.SayHello()
	greeting.SayBad()
	i := greeting.GiveMeInt()
	fmt.Println(i)
}
