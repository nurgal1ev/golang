package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите команду: ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("ошибка ввода")
			return
		}

		text := scanner.Text()

		fields := strings.Fields(text)

		if len(fields) == 0 {
			fmt.Println("вы ничего не ввели")
			return
		}

		cmd := fields[0]
		if cmd == "выйти" {
			fmt.Println("выход из программы")
			return
		}

		if cmd == "добавить" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]
				if i != len(fields)-1 {
					str += " "
				}
			}

			fmt.Println("вы хотите добавить:", str)
			fmt.Println("")
		} else if cmd == "удалить" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]
				if i != len(fields)-1 {
					str += " "
				}
			}
			fmt.Println("вы хотите удалить:", str)
			fmt.Println("")
		} else if cmd == "help" {
			fmt.Println("команда: добавить")
			fmt.Println("команда: удалить")
			fmt.Println("")
		} else {
			fmt.Println("вы ввели неизвестную команду")
		}
	}
}
