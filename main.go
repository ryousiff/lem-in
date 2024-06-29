package main

import (
	"fmt"
	lem "lem/functions"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}
	fileName := os.Args[1]
	basename := filepath.Base(fileName)
	if strings.HasPrefix(basename, "#") || strings.HasPrefix(basename, "L") {
		fmt.Println("ERROR: filename starts with invalid character (# or L)")
		os.Exit(1)

	}

	// Parse the farm from the input file
	farm := lem.File(fileName)

	if farm == nil {
		fmt.Println("Failed to parse farm from input file.")
		return
	}

	// Print the farm configuration (this echoes the input)
	inputText := lem.PrintFarmConfiguration(farm)
	fmt.Print(inputText)
	fmt.Println()

	// Find all possible routes from start to end
	paths := lem.Edmonds(farm)
	if len(paths) == 0 {
		fmt.Printf("ERROR: invalid data format\nNo path from start room to end room\n")
		return
	}

	// Choose optimal paths
	farm.Paths = lem.ChooseOptimalPaths(paths, farm.NumAnt)

	// Distribute ants to paths
	lem.DistributeAnts(farm)

	// Simulate ant movements
	movements := lem.SimulateAntsMovement(farm)

	// Print the movements
	for _, move := range movements {
		if move != "" {
			fmt.Println(move)
		}
	}
	fmt.Println("$")
}
