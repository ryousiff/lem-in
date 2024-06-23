package main

import (
    "fmt"
    lem "lem/functions"
    "os"
    "sort"
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

    // Find paths from the start room to the end room using Edmonds algorithm
    paths := lem.Edmonds(farm)

    // Sort paths by length to find the quickest paths
    sort.Slice(paths, func(i, j int) bool {
        return len(paths[i].Rooms) < len(paths[j].Rooms)
    })

    // Choose the optimal paths using the ChooseOptimalPaths function
    optimalPaths := lem.ChooseOptimalPaths(paths, farm.StartRoom)

    // Print the optimal paths found
    fmt.Println("Optimal Paths:")
    for _, path := range optimalPaths {
        fmt.Print("path: ")
        for j, room := range path.Rooms {
            if j > 0 {
                fmt.Print(" -> ")
            }
            fmt.Printf("%s (%s, %s)", room.Name, room.CoordX, room.CoordY)
        }
        fmt.Println()
    }
}
