package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	directed := flag.Bool("d", false, "directed graph")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	edges := []string{}

	for scanner.Scan() {
		var u, v string
		fmt.Sscanf(scanner.Text(), "%s %s", &u, &v)
		if *directed {
			edges = append(edges, fmt.Sprintf("    %s -> %s;", u, v))
		} else {
			edges = append(edges, fmt.Sprintf("    %s -- %s;", u, v))
		}
	}

	if *directed {
		fmt.Println("digraph G {")
	} else {
		fmt.Println("graph G {")
	}

	for _, edge := range edges {
		fmt.Println(edge)
	}
	fmt.Println("}")
}
