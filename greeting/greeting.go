package greeting

import "fmt"

func SayHello() {
	fmt.Println("hello, golang packages!")
	j := GiveMeInt()
	fmt.Println(j)
}
