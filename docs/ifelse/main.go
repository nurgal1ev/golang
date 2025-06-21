package main

import "fmt"

func main() {
	score := 18
	/*
		if score > 15 {
			if score > 15 {
				fmt.Println("great")
			} else {
				fmt.Println("good")
			}
		} else {
			fmt.Println("bad")
		}*/

	if score > 15 {
		fmt.Println("great")
	} else if score > 10 {
		fmt.Println("good")
	} else {
		fmt.Println("bad")
	}
}
