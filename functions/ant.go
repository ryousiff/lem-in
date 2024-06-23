package lem

func AssignAnts(farm *Farm) {
    antIndex := 1
    currentPath := 0

    for antIndex <= farm.NumAnt {
        if currentPath >= len(farm.Paths) {
            currentPath = 0
        }

        lowestCostPath := farm.Paths[currentPath]
        lowestCost := cost(lowestCostPath)

        for i, path := range farm.Paths {
            if cost(path) < lowestCost {
                lowestCostPath = path
                lowestCost = cost(path)
                currentPath = i
            }
        }

        ant := &Ant{
            ID:          antIndex,
            Path:        &lowestCostPath,
            CurrentRoom: lowestCostPath.Rooms[0],
        }
        farm.Ants = append(farm.Ants, ant)
        lowestCostPath.Queue = append(lowestCostPath.Queue, ant)
        lowestCostPath.NumNamlaty++
        antIndex++
        currentPath++
    }
}

func cost(path Path) int {
    return len(path.Rooms) + path.NumNamlaty
}
