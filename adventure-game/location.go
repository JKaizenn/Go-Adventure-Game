package main

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
