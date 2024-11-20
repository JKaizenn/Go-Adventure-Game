package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//Welcome the User
	fmt.Println("Welcome to the Adventure Game!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')

	fmt.Printf("Hello, %sLet the adventure begin!\n", name)

	//Start the Game loop
	for {
		fmt.Print(">")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		//Process the player's command
		if input == "quit" || input == "exit" {
			fmt.Println("Thank you for playing!")
			break
		} else {
			fmt.Printf("You entered: %s\n", input)

		}
	}

}
