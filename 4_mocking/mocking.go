package main

import (
	interfaces "avalith/testing/mocking/interfaces"
	"fmt"
)

// Struct def

// Item struct for item
type Item struct {
	ID   string
	Name string
}

// NewItem item generator
func NewItem(IDManager interfaces.IDManager, name string) Item {
	return Item{
		ID:   IDManager.GenID(),
		Name: name,
	}
}

func main() {
	IDManager := interfaces.NewIDManagerLocal()
	fmt.Println(NewItem(IDManager, "Item 1"))
}
