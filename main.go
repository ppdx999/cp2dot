package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	edges := []string{}

	for scanner.Scan() {
		var u, v string
		fmt.Sscanf(scanner.Text(), "%s %s", &u, &v)
		edges = append(edges, fmt.Sprintf("    %s -> %s;", u, v))
	}

	fmt.Println("digraph G {")
	for _, edge := range edges {
		fmt.Println(edge)
	}
	fmt.Println("}")
}
