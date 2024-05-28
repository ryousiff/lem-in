package lem

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
            if !visited[nextRoom] {
                newPath := append(path, nextRoom)
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
