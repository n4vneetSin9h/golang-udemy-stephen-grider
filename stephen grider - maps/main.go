package main

import "fmt"

func main() {

	colors := map[string]string{
		"purple": "#800080",
		"blue": "#0000ff",
		"red": "#ff0000",
		"green": "#008000",
		"yellow": "#FFCE30",
		"black": "#000",
		"white": "#fff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {

	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["purple"] = "#800080"

	// delete(colors, "blue")

	// fmt.Println(colors)

	for color, hex := range c {
		fmt.Println("Hex color for", color, "is", hex)
	}
}