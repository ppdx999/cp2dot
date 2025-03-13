package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Println("cp2dot v0.1.0")
		return
	}

	fmt.Println("Welcome to cp2dot!")
}
