package main

import (
	"fmt"
	task1 "infosec/tasks/1"
	"strings"
)

func enterArgs(amount int) []string {
	args := make([]string, 0)
	switch {
	case amount > 0:
		fmt.Printf("Enter %d args: \n", amount)
		var arg string
		for i := 0; i < amount; i++ {
			fmt.Scan(&arg)
			args = append(args, arg)
		}
	// 0 means that there's undefined amount of arguments
	case amount == 0:
		fmt.Printf("Enter args (to stop enter 'done'): \n")
		var arg string
		for i := 0; true; i++ {
			fmt.Scan(&arg)
			if strings.ToLower(arg) == "done" {
				break
			}
			args = append(args, arg)
		}

	default:
		fmt.Printf("ERROR! Amount of arguments is less than 0 in main.go/enterArgs.\n")
	}
	fmt.Println()
	return args
}

func hr() {
	fmt.Print("\n====================\n")
}

func main() {
	str := "kekus"
	newStr := task1.ShuffleIPtest([]byte(str), true, 1)
	fmt.Println("Shuffled: \"", string(newStr), "\"")
	oldStr := task1.ShuffleIPRevtest(newStr, true, 1)
	fmt.Println("Unshuffled: \"", string(oldStr), "\"")
	return

	done := map[string]bool{}

	fmt.Printf("To see the help page, type 'help', to see the list of exercises, type 'list'.\n")
	fmt.Printf("To exit the program, type 'exit'\n")
	hr()

	for {
		print("What exercise do you wanna check: ")
		var line string
		_, err := fmt.Scan(&line)
		hr()

		if err != nil {
			panic(err)
		} else {
			switch {

			case strings.ToLower(line) == "help" || strings.ToLower(line) == "h":
				fmt.Println("You can pop up this HELP page again by typing 'help' or 'h' when you're choosing exercises.")
				fmt.Println("You can EXIT the program by typing 'exit', 'e' or '-1' when you're choosing exercises.")
				fmt.Println("You can see the list of avaliable exercises by  typing 'list' or 'l' when you're choosing exercises.")
				fmt.Println("The input isn't case-sensetive. And wrong inputs won't crash the program.")
			case strings.ToLower(line) == "exit" || strings.ToLower(line) == "e" || strings.ToLower(line) == "-1":
				fmt.Println("Thank you for using the program. In hope for a good mark :)")
				return
			case strings.ToLower(line) == "list" || strings.ToLower(line) == "l":

			case done[line]:
				switch line {

				}

			default:
				fmt.Println("Sorry, there's no such exercise or command...")
			}
		}

		hr()
	}
}
