package main

import "fmt"

func main() {
	/*
		type User struct {
			Name    string
			Rating  float64
			Premium bool
		}

		userArray := []User{
			User{
				Name:    "Jack",
				Rating:  5.5,
				Premium: true,
			},

			User{
				Name:    "Mary",
				Rating:  3.5,
				Premium: false,
			},

			User{
				Name:    "Aboba",
				Rating:  7.5,
				Premium: true,
			},
		}

		fmt.Println("до")
		fmt.Println("len:", len(userArray))
		fmt.Println("cap:", cap(userArray))

		userArray = append(
			userArray, User{
				Name:    "Alex",
				Rating:  6.5,
				Premium: true,
			},
		)
		fmt.Println("после")
		fmt.Println("len:", len(userArray))
		fmt.Println("cap:", cap(userArray))

		for _, user := range userArray {
			pp.Println(user)
		}

		for index, user := range userArray {
			if user.Premium {
				userArray[index].Rating += 1.0
			}
		}

		pp.Println(userArray)*/

	intSlice := make([]int, 0)
	fmt.Println(intSlice)
	intSlice = append(intSlice, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(intSlice)
}
