package lem

import (
    "fmt"
    "math"
    "sort"
)

// Edmonds function finds all possible paths from the start room to the end room in the farm.
func Edmonds(farm *Farm) []*Path {
    start := []*Room{farm.StartRoom}
    end := farm.EndRoom
    queue := []*Path{{Rooms: start}} //queue a slice of path poniters with a path containing the start room 
    var paths []*Path

    for len(queue) > 0 { //loop till the queue is empty
        path := queue[0] //gets first path
        queue = queue[1:] //updates the queue with removing the first one
        currentRoom := path.Rooms[len(path.Rooms)-1] // gets last room in the path

        if currentRoom == end { // checks if the room is the last room
            newPath := &Path{Rooms: append([]*Room(nil), path.Rooms...)} //save the path it to the struct
            paths = append(paths, newPath)
            continue //skips to next itteration since a complete path has been found
        }

        for _, link := range currentRoom.Links { //finds all links from current room
            nextRoom := link.Room2 //gets the next room in the link
            if nextRoom == currentRoom { //make sure next room isnt the current room
                nextRoom = link.Room1
            }
            if !containsRoom(path.Rooms, nextRoom) { //ensure that the next room the linked one is not already in a path
                newPath := &Path{Rooms: append(append([]*Room(nil), path.Rooms...), nextRoom)}
                queue = append(queue, newPath)
            }
        }
    }
    return paths
}

// ChooseOptimalPaths function selects the optimal paths from the list of all found paths.
func ChooseOptimalPaths(paths []*Path, numAnts int) []*Path {
    applyFindMaxFlow(paths)

    var filteredPaths []*Path
    for _, path := range paths {
        if !path.Skip {
            filteredPaths = append(filteredPaths, path)
        }
    }

    sort.Slice(filteredPaths, func(i, j int) bool {
        return len(filteredPaths[i].Rooms) < len(filteredPaths[j].Rooms)
    })

    var optimalPaths []*Path
    minSteps := math.MaxInt32

    for i := 1; i <= len(filteredPaths); i++ {
        selectedPaths := filteredPaths[:i]
        steps := calculateSteps(selectedPaths, numAnts)
        if steps < minSteps {
            minSteps = steps
            optimalPaths = selectedPaths
        }
    }
    return optimalPaths
}

// applyFindMaxFlow applies the FindMaxFlow logic to skip paths with shared rooms.
func applyFindMaxFlow(paths []*Path) {
    linkedTo := make([][]int, len(paths))
    for i := range linkedTo {
        linkedTo[i] = make([]int, len(paths))
    }

    for i := range paths {
        if paths[i].Skip {
            continue
        }
        for j := i + 1; j < len(paths); j++ {
            if paths[j].Skip {
                continue
            }
            if numOfSameRooms(paths[i].Rooms, paths[j].Rooms) > 2 {
                linkedTo[i][j] = 1
                linkedTo[j][i] = 1
            }
        }
    }

    maxSimilarity, maxPath := 0, -1
    for i := len(linkedTo) - 1; i >= 0; i-- {
        sumConnections := 0
        for _, conn := range linkedTo[i] {
            sumConnections += conn
        }
        if sumConnections > maxSimilarity {
            maxSimilarity = sumConnections
            maxPath = i
        }
    }

    if maxSimilarity != 0 {
        paths[maxPath].Skip = true
        applyFindMaxFlow(paths)
    }
}

// calculateSteps calculates the total number of steps required to move all ants using the given paths.
func calculateSteps(paths []*Path, numAnts int) int {
    maxPathLength := 0
    for _, path := range paths {
        if len(path.Rooms) > maxPathLength {
            maxPathLength = len(path.Rooms)
        }
    }

    totalSteps, remainingAnts := 0, numAnts
    for remainingAnts > 0 {
        totalSteps++
        for range paths {
            if remainingAnts > 0 {
                remainingAnts--
            }
        }
    }
    return totalSteps + maxPathLength - 1
}

// numOfSameRooms function counts the number of rooms that are the same in two routes.
func numOfSameRooms(route1, route2 []*Room) int {
    count := 0
    for _, room1 := range route1 {
        for _, room2 := range route2 {
            if room1 == room2 {
                count++
            }
        }
    }
    return count
}

// containsRoom function checks if a room is in a slice of rooms.
func containsRoom(rooms []*Room, room *Room) bool {
    for _, r := range rooms {
        if r == room {
            return true
        }
    }
    return false
}

// PrintFarmConfiguration prints the farm configuration.
func PrintFarmConfiguration(farm *Farm) string {
    var result string
    result += fmt.Sprintf("%d\n", farm.NumAnt)
    for _, room := range farm.Rooms {
        if room.IsStart {
            result += "##start\n"
        }
        if room.IsEnd {
            result += "##end\n"
        }
        result += fmt.Sprintf("%s %s %s\n", room.Name, room.CoordX, room.CoordY)
    }
    for _, link := range farm.Links {
        result += fmt.Sprintf("%s-%s\n", link.Room1.Name, link.Room2.Name)
    }
    return result
}
