package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printPrompt()

		if scanner.Scan() {
			input := scanner.Text()

			if input == ".exit" {
				fmt.Println("goodbye")
				break
			} else {
				fmt.Println("Unrecognized Command")
			}

		}
	}
}

func printPrompt() {
	fmt.Print("db > ")
}
