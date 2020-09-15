package main

import (
	"fmt"
	"os"
)

func printName() string {
	return fmt.Sprintf("Hola, %s", os.Getenv("NAME"))
}

func main() {
	fmt.Println(printName())
}
