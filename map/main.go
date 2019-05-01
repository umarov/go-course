package main

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hex := range colors {
		println(color, hex)
	}
}
