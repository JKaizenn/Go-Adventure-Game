package main

import (
	"fmt"
	"strings"
)

type Player struct {
	Name            string
	CurrentLocation string
	Inventory       []Item
}

func NewPlayer(name, startLocation string) Player {
	return Player{
		Name:            name,
		CurrentLocation: startLocation,
		Inventory:       []Item{},
	}
}

func (p *Player) ShowInventory() {
	if len(p.Inventory) == 0 {
		fmt.Println("Your inventory is empty.")
		return
	}
	fmt.Println("You are carrying:")
	for _, item := range p.Inventory {
		fmt.Printf("- %s: %s\n", item.Name, item.Description)
	}
}

func (p *Player) Move(direction string, locations map[string]Location) {
	current, exists := locations[p.CurrentLocation]
	if !exists {
		fmt.Printf("Error: Current location '%s' does not exist.\n", p.CurrentLocation)
		return
	}
	next, ok := current.Exits[direction]
	if !ok {
		fmt.Printf("You can't go '%s' from here.\n", direction)
		return
	}
	if _, exists := locations[next]; !exists {
		fmt.Printf("The location '%s' does not exist.\n", next)
		return
	}
	p.CurrentLocation = next
	fmt.Printf("You move %s.\n", direction)
	fmt.Println(locations[next].Description)
}

func (p *Player) TakeItem(itemName string, locations map[string]Location) {
	current, exists := locations[p.CurrentLocation]
	if !exists {
		fmt.Printf("Error: Current location '%s' does not exist.\n", p.CurrentLocation)
		return
	}
	for i, item := range current.Items {
		if strings.EqualFold(item.Name, itemName) {
			// Add item to inventory
			p.Inventory = append(p.Inventory, item)
			// Remove item from location
			current.Items = append(current.Items[:i], current.Items[i+1:]...)
			locations[p.CurrentLocation] = current // Update location in the map
			fmt.Printf("You have taken the %s.\n", item.Name)
			return
		}
	}
	fmt.Printf("There is no '%s' here to take.\n", itemName)
}
