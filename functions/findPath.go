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

func ShortestPathsFromNeighbors(farm Farm) Path {
	shortestPaths := []Path{}

	for _, link := range farm.StartRoom.Links {
		neighbor := link.Room1
		if link.Room1 == farm.StartRoom {
			neighbor = link.Room2
		}

		allPaths := findAllPaths(farm, []*Room{neighbor}, make(map[*Room]bool))

		var shortestPath Path
		minSteps := 0
		for _, path := range allPaths {
			if minSteps == 0 || len(path.Rooms) < minSteps {
				minSteps = len(path.Rooms)
				shortestPath = path
			}
		}

		shortestPath.Rooms = append([]*Room{farm.StartRoom}, shortestPath.Rooms...)

		shortestPaths = append(shortestPaths, shortestPath)
	}

	return Path{
		Rooms:    []*Room{farm.StartRoom},
		Shortest: shortestPaths,
	}
}
