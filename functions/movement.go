package lem

import (
    "fmt"
    "strings"
)

func SimulateAntsMovement(farm *Farm) []string {
    var movements []string
    for !allAntsArrived(farm) { //while there are ants not arrived, if all arrived it will stop
        var currentMove []string
        for _, path := range farm.Paths {
            currentMove = append(currentMove, moveAntsOnPath(path, farm.EndRoom)...)
        }
        if len(currentMove) > 0 {
            movements = append(movements, strings.Join(currentMove, " "))
        }
    }
    return movements
}

func allAntsArrived(farm *Farm) bool {
    for _, path := range farm.Paths {
        if len(path.Ants) > 0 {
            return false
        }
    }
    return true //if no ants, all of them arrived 
}

func moveAntsOnPath(path *Path, endRoom *Room) []string {
    var moves []string
    for i := len(path.Rooms) - 1; i > 0; i-- {
        currentRoom := path.Rooms[i]
        previousRoom := path.Rooms[i-1]
        for j := 0; j < len(path.Ants); j++ {
            ant := path.Ants[j]
            if ant.CurrentRoom == previousRoom { //if the ant's current room is the prev room
                ant.CurrentRoom = currentRoom // we move it to current room
                moves = append(moves, fmt.Sprintf("L%d-%s", ant.ID, currentRoom.Name))
                if currentRoom == endRoom {
                    path.Ants = append(path.Ants[:j], path.Ants[j+1:]...)
                    j--
                }
                break
            }
        }
    }
    return moves
}


//movement in 1, 2 and 5 are  wrong
//4 the paths are correct but wrong order