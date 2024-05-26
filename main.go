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

    farm := lem.File(os.Args[1])

    // Print the number of ants
    fmt.Printf("Number of ants: %d\n", farm.NumAnt)

    // Print the rooms
    fmt.Println("Rooms:")
    for _, room := range farm.Rooms {
        fmt.Printf("  %s (%s, %s) - Start: %t, End: %t\n", room.Name, room.CoordX, room.CoordY, room.IsStart, room.IsEnd)
        fmt.Print("    Links: ")
        for i, link := range room.Links {
            if i > 0 {
                fmt.Print(", ")
            }
            fmt.Print(link.Room1.Name, "-", link.Room2.Name)
        }
        fmt.Println()
    }
}

