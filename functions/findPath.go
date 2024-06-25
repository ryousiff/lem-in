package lem

import (
    "math"
    "fmt"
)

// Edmonds function finds all possible paths from the start room to the end room in the farm.
func Edmonds(farm *Farm) []*Path {
    start := []*Room{farm.StartRoom} // Initialize the start room
    end := farm.EndRoom // Initialize the end room
    queue := []*Path{{Rooms: start}} // Initialize the queue with the start room
    var paths []*Path // Slice to store all found paths

    // Breadth-First Search (BFS) to find all paths
    for len(queue) > 0 {
        path := queue[0] // Get the first path in the queue
        queue = queue[1:] // Remove the first path from the queue
        currentRoom := path.Rooms[len(path.Rooms)-1] // Get the last room in the current path

        // If the current room is the end room, add the path to the paths slice
        if currentRoom == end {
            newPath := &Path{Rooms: make([]*Room, len(path.Rooms))}
            copy(newPath.Rooms, path.Rooms)
            paths = append(paths, newPath)
            continue
        }

        // Explore all links from the current room
        for _, link := range currentRoom.Links {
            nextRoom := link.Room2
            if nextRoom == currentRoom {
                nextRoom = link.Room1
            }
            // If the next room is not already in the current path, add it to the new path
            if !containsRoom(path.Rooms, nextRoom) {
                newPath := &Path{Rooms: make([]*Room, len(path.Rooms), len(path.Rooms)+1)}
                copy(newPath.Rooms, path.Rooms)
                newPath.Rooms = append(newPath.Rooms, nextRoom)
                queue = append(queue, newPath)
            }
        }
    }

    return paths
}

// ChooseOptimalPaths function selects the optimal paths from the list of all found paths.
func ChooseOptimalPaths(paths []*Path, startRoom *Room) []*Path {
    // Remove redundant paths
    paths = RemoveParents(paths)

    // Group paths by their second room
    groups := make(map[*Room][]*Path)
    for _, path := range paths {
        if len(path.Rooms) > 1 {
            secondRoom := path.Rooms[1]
            groups[secondRoom] = append(groups[secondRoom], path)
        }
    }

    // Find the optimal path for each group
    var optimalPaths []*Path
    for _, pathsInGroup := range groups {
        minSharedRooms := math.MaxFloat64
        minPathLength := math.MaxFloat64
        var optimalPath *Path

        for _, path := range pathsInGroup {
            sharedRooms := 0.0
            for _, otherPaths := range groups {
                if &otherPaths != &pathsInGroup {
                    for _, otherPath := range otherPaths {
                        if hasSharedRooms(path, otherPath) {
                            sharedRooms++
                            break
                        }
                    }
                }
            }

            pathLength := float64(len(path.Rooms))

            if sharedRooms < minSharedRooms || (sharedRooms == minSharedRooms && pathLength < minPathLength) {
                minSharedRooms = sharedRooms
                minPathLength = pathLength
                optimalPath = path
            }
        }

        optimalPaths = append(optimalPaths, optimalPath)
    }

    // Filter optimalPaths to choose one path that has the shortest length between paths that have shared rooms with other paths in other groups
    var filteredPaths []*Path
    for _, path := range optimalPaths {
        hasSharedRoomsWithOthers := false
        for _, otherPath := range optimalPaths {
            if path != otherPath && hasSharedRooms(path, otherPath) {
                hasSharedRoomsWithOthers = true
                break
            }
        }

        if !hasSharedRoomsWithOthers {
            filteredPaths = append(filteredPaths, path)
        }
    }

    if len(filteredPaths) < len(optimalPaths) {
        minPathLength := math.Inf(1)
        var selectedPath *Path

        for _, path := range optimalPaths {
            if !contains(filteredPaths, path) {
                pathLength := float64(len(path.Rooms))
                if pathLength < minPathLength {
                    minPathLength = pathLength
                    selectedPath = path
                }
            }
        }

        filteredPaths = append(filteredPaths, selectedPath)
    }

    return filteredPaths
}

// RemoveParents function removes redundant paths that are subsets of other paths.
func RemoveParents(paths []*Path) []*Path {
    for i := 0; i < len(paths); i++ {
        for j := i + 1; j < len(paths); j++ {
            if len(paths[i].Rooms) != 2 && numOfSameRooms(paths[i].Rooms, paths[j].Rooms) == len(paths[i].Rooms) {
                paths = append(paths[:j], paths[j+1:]...)
                j--
            }
        }
    }
    return paths
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

// hasSharedRooms function checks if two paths share any rooms other than the start and end rooms.
func hasSharedRooms(path1, path2 *Path) bool {
    rooms1 := make(map[*Room]bool)
    for _, room := range path1.Rooms {
        if room != path1.Rooms[0] && room != path1.Rooms[len(path1.Rooms)-1] {
            rooms1[room] = true
        }
    }

    for _, room := range path2.Rooms {
        if room != path2.Rooms[0] && room != path2.Rooms[len(path2.Rooms)-1] {
            if rooms1[room] {
                return true
            }
        }
    }

    return false
}

// contains function checks if a path is in a slice of paths.
func contains(paths []*Path, path *Path) bool {
    for _, p := range paths {
        if p == path {
            return true
        }
    }
    return false
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
