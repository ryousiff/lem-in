package lem

func DistributeAnts(farm *Farm) {
    for i := 1; i <= farm.NumAnt; i++ {
        ant := &Ant{ID: i, CurrentRoom: farm.StartRoom}
        farm.Ants = append(farm.Ants, ant)

        // Find the path with the least total "cost"
        minCost := len(farm.Paths[0].Rooms) + len(farm.Paths[0].Ants) //num of rooms in a path + num of ants in a path
        minPath := farm.Paths[0]

        for _, path := range farm.Paths[1:] {
            cost := len(path.Rooms) + len(path.Ants)
            if cost < minCost {
                minCost = cost
                minPath = path
            }
        }

        // Assign the ant to the path with the least cost
        minPath.Ants = append(minPath.Ants, ant)
    }
}

//cost is used to choose the best path 
//cost equation len(path) + len(ant)
//we choose the min cost path