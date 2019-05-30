package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":    "#FF0000",
		"green":  "#GREEN#",
		"yellow": "#YELLOW",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
