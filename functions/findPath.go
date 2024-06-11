package lem

import "math"

func FindPath(farm Farm) []Path {
    start := []*Room{farm.StartRoom}
    return findAllPaths(farm, start, make(map[*Room]bool))
}

func findAllPaths(farm Farm, path []*Room, visited map[*Room]bool) []Path {
    var paths []Path
    lastRoom := path[len(path)-1]
    visited[lastRoom] = true

    if roomsEqual(*lastRoom, *farm.EndRoom) {
        pathCopy := make([]*Room, len(path))
        copy(pathCopy, path)
        paths = append(paths, Path{Rooms: pathCopy})
    } else {
        for _, nextRoom := range lastRoom.Neighbors {
            if !visited[nextRoom] && !roomsEqual(*nextRoom, *farm.StartRoom) {
                newPath := make([]*Room, len(path)+1)
                copy(newPath, path)
                newPath[len(path)] = nextRoom
                newVisited := make(map[*Room]bool)
                for room, visited := range visited {
                    newVisited[room] = visited
                }
                paths = append(paths, findAllPaths(farm, newPath, newVisited)...)
            }
        }
    }

    return paths
}



func roomsEqual(r1, r2 Room) bool {
    return r1.Name == r2.Name && r1.CoordX == r2.CoordX && r1.CoordY == r2.CoordY
}

func ShortestPathsFromNeighbors(farm Farm) Path {
    shortestPaths := []Path{}

    for _, link := range farm.StartRoom.Links {
        neighbor := link.Room1
        if link.Room1 == farm.StartRoom {
            neighbor = link.Room2
        }

        allPaths := findAllPaths(farm, []*Room{neighbor}, make(map[*Room]bool))

        var shortestPath Path
        minLength := math.MaxInt32
        for _, path := range allPaths {
            bottlenecks := countBottlenecks(path.Rooms)
            pathLength := len(path.Rooms)
            if bottlenecks == 0 && pathLength < minLength && isUnique(shortestPaths, path) {
                minLength = pathLength
                shortestPath = path
            }
        }

        if shortestPath.Rooms != nil {
            shortestPath.Rooms = append([]*Room{farm.StartRoom}, shortestPath.Rooms...)
            if len(shortestPaths) > 0 && equalPaths(shortestPaths[len(shortestPaths)-1].Rooms, shortestPath.Rooms) {
                continue
            }
            shortestPaths = append(shortestPaths, shortestPath)
        }
    }

    return Path{
        Rooms:    []*Room{farm.StartRoom},
        Shortest: shortestPaths,
    }
}





func countBottlenecks(rooms []*Room) int {
    bottlenecks := 0
    visited := make(map[*Room]bool)

    for _, room := range rooms {
        if room.IsStart || room.IsEnd {
            continue
        }

        if visited[room] {
            bottlenecks++
        } else {
            visited[room] = true
        }
    }

    return bottlenecks
}

func isUnique(paths []Path, path Path) bool {
    for _, p := range paths {
        if equalPaths(p.Rooms, path.Rooms) {
            return false
        }
    }
    return true
}

func equalPaths(path1, path2 []*Room) bool {
    if len(path1) != len(path2) {
        return false
    }
    for i := range path1 {
        if !roomsEqual(*path1[i], *path2[i]) {
            return false
        }
    }
    return true
}
