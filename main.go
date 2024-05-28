package main

import (
	"fmt"
	"os"
	lem "lem/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: lem <input-file>")
		return
	}

	// Parse the farm from the input file
	farm := lem.File(os.Args[1])
	if farm == nil {
		fmt.Println("Failed to parse farm from input file.")
		return
	}

	// Print the number of ants
	fmt.Printf("Number of ants: %d\n", farm.NumAnt)

	// Print the rooms and their links
	fmt.Println("Rooms:")
	for _, room := range farm.Rooms {
		fmt.Printf("  %s (%s, %s) - Start: %t, End: %t\n", room.Name, room.CoordX, room.CoordY, room.IsStart, room.IsEnd)
		fmt.Print("    Links: ")
		for i, link := range room.Links {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%s-%s", link.Room1.Name, link.Room2.Name)
		}
		fmt.Println()
	}

	// Build the neighbors list for each room based on links
	for _, room := range farm.Rooms {
		for _, link := range room.Links {
			if link.Room1 == room {
				room.Neighbors = append(room.Neighbors, link.Room2)
			} else {
				room.Neighbors = append(room.Neighbors, link.Room1)
			}
		}
	}

	// Find paths from the start room to the end room
	paths := lem.FindPath(*farm)

	// Print the paths found
	fmt.Println("Paths found:")
	for i, path := range paths {
		fmt.Printf("Path %d: ", i+1)
		for j, room := range path.Rooms {
			if j > 0 {
				fmt.Print(" -> ")
			}
			fmt.Print(room.Name)
		}
		fmt.Println()
	}
}
