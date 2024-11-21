package main

import "fmt"

func ProcessCommand(command string, player *Player, locations map[string]Location) {
	switch {
	case command == "look":
		location := locations[player.CurrentLocation]
		fmt.Println(location.Description)
		if len(location.Items) > 0 {
			fmt.Println("You see the following items:")
			for _, item := range location.Items {
				fmt.Printf("- %s: %s\n", item.Name, item.Description)
			}
		}
	case command == "help":
		fmt.Println("Available Commands: look, go [direction], take [item], inventory, help, quit")
	case command == "inventory":
		player.ShowInventory()
	case command[:3] == "go ":
		direction := command[3:]
		player.Move(direction, locations)
	case command[:5] == "take ":
		item := command[5:]
		player.TakeItem(item, locations)
	default:
		fmt.Printf("I don't understand '%s'. Type 'help' for available commands.\n", command)
	}
}
