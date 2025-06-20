package main

import (
	"fmt"
	"golanggit/greeting"
	"golanggit/user"
)

func main() {
	fmt.Println("hello")
	greeting.SayHello()
	greeting.SayBad()
	i := greeting.GiveMeInt()
	fmt.Println(i)

	u := user.User{}
}
