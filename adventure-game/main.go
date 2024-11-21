package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Welcome the user
	fmt.Println("Welcome to the Adventure Game!")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}
	name = strings.TrimSpace(name)
	fmt.Printf("Hello, %s! Let the adventure begin!\n", name)

	// Load locations
	locations, err := LoadLocations("locations.json")
	if err != nil {
		fmt.Println("Error loading locations:", err)
		return
	}

	// Initialize player
	player := NewPlayer(name, "field")

	// Game loop
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input)

		if input == "quit" || input == "exit" {
			fmt.Println("Thank you for playing!")
			break
		}

		// Process command
		ProcessCommand(input, &player, locations)
	}
}
