package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Welcome the User
	fmt.Println("Welcome to the Adventure Game!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // Trim any whitespace and newline characters

	fmt.Printf("Hello, %s! Let the adventure begin!\n", name)

	// Start the Game loop
	for {
		fmt.Print(">")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Process the player's command
		if input == "quit" || input == "exit" {
			fmt.Println("Thank you for playing!")
			break
		} else {
			processCommand(input)
		}
	}
}

func processCommand(command string) {
	switch command {
	case "look":
		fmt.Println("You are standing in an open field west of a white house.")
	case "help":
		fmt.Println("Available Commands: look, help, quit")
	default:
		fmt.Printf("I don't understand '%s'. Type 'help' for available commands.\n", command)
	}
}
