package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Location struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Exits       map[string]string `json:"exits"`
	Items       []Item            `json:"items"`
}

func main() {

	// Welcome the User
	fmt.Println("Welcome to the Adventure Game!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Hello, %s! Let the adventure begin!\n", name)

	//Open the JSON File
	file, err := os.Open("locations.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	//Decode the JSON file
	decoder := json.NewDecoder(file)
	var locationsData []Location

	err = decoder.Decode(&locationsData)
	if err != nil {
		fmt.Println("Error decoding JSON data:", err)
		return
	}

	//Convert the slice to map
	locations := make(map[string]Location)
	for _, loc := range locationsData {
		locations[loc.Name] = loc
	}

	//Set the player's starting location
	currentLocation := "field"

	// Start the Game loop
	for {
		fmt.Print(">")
		fmt.Println("Your current options are: look and help")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Process the player's command
		if input == "quit" || input == "exit" {
			fmt.Println("Thank you for playing!")
			break
		} else {
			processCommand(input, locations, &currentLocation)
		}
	}
}

func processCommand(command string, locations map[string]Location, currentLocation *string) {
	switch {
	case command == "look":
		loc := locations[*currentLocation]
		fmt.Println(loc.Description)
		if len(loc.Items) > 0 {
			fmt.Println("You see the following items:")
			for _, item := range loc.Items {
				fmt.Printf("-%s: %s\n", item.Name, item.Description)
			}
		}
	case command == "help":
		fmt.Println("Available Commands: look, help, quit")
	default:
		fmt.Printf("I don't understand '%s'. Type 'help' for available commands.\n", command)
	}
}
