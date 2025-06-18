package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println("итерация: ", i)
		time.Sleep(500 * time.Millisecond)
	}
}
