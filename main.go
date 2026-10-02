package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello from Go running inside a container!")
	fmt.Println("Application started at:", time.Now().Format(time.RFC3339))
}
