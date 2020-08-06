package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#ff0001",
		"black": "#000000",
		"white": "#ffffff",
	}
	fmt.Println(colors)

	auxColors := make(map[string]string)
	var auxColors2 map[string]string
	delete(auxColors2, "orange")
	auxColors["white"] = "#ffffff"

	delete(auxColors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
