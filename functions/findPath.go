package lem

import "math"

func Edmonds(farm *Farm) []*Path {
    start := []*Room{farm.StartRoom}
    end := farm.EndRoom
    queue := []*Path{{Rooms: start}}
    var paths []*Path

    for len(queue) > 0 {
        path := queue[0]
        queue = queue[1:]
        currentRoom := path.Rooms[len(path.Rooms)-1]

        if currentRoom == end {
            newPath := &Path{Rooms: make([]*Room, len(path.Rooms))}
            copy(newPath.Rooms, path.Rooms)
            paths = append(paths, newPath)
            continue
        }

        for _, link := range currentRoom.Links {
            nextRoom := link.Room2
            if nextRoom == currentRoom {
                nextRoom = link.Room1
            }
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

func ChooseOptimalPaths(paths []*Path, startRoom *Room) []*Path {
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

func contains(paths []*Path, path *Path) bool {
    for _, p := range paths {
        if p == path {
            return true
        }
    }
    return false
}

func containsRoom(rooms []*Room, room *Room) bool {
    for _, r := range rooms {
        if r == room {
            return true
        }
    }
    return false
}
