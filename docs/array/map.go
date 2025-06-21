package main

import "fmt"

func main() {
	weather := map[int]int{
		11: +3,
		12: -2,
		13: +5,
		14: -1,
	}

	weather[15] = +6

	for k, _ := range weather {
		if weather[k] > 0 {
			weather[k] += 1
		}
	}
	fmt.Println(weather)
	//fmt.Println(weather)

	users := map[string]bool{
		"tom":   true,
		"jack":  false,
		"jills": true,
		"mary":  true,
		"alex":  false,
	}
	c, ok := users["jackil"]

	if !ok {
		fmt.Println("нет в базе")
	} else {
		fmt.Println("есть в базе")
	}

	fmt.Println(c, ok)
}
