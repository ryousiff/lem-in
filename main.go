package main

import (
    lem "lem/functions"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run main.go <filename>")
        return
    }

    // Parse the farm from the input file
    farm := lem.File(os.Args[1])
    if farm == nil {
        // fmt.Println("Failed to parse farm from input file.")
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

    // Remove redundant paths (longer routes)
    paths = lem.RemoveParents(paths)

    // Choose optimal paths
    farm.Paths = lem.ChooseOptimalPaths(paths, farm.StartRoom)

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
