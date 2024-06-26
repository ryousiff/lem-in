package lem

type Room struct {
    Name      string
    CoordX    string
    CoordY    string
    Visited   bool
    IsEnd     bool
    IsStart   bool
    Links     []*Link
    Neighbors []*Room
    LineNum   int
    Ants      []*Ant
}

type Ant struct {
    ID          int
    Path        *Path
    CurrentRoom *Room
}

// type Farm struct {
//     Rooms         []*Room
//     Ants          []*Ant
//     Links         []*Link
//     NumAnt        int
//     StartRoom     *Room
//     EndRoom       *Room
//     AntPositions  map[int]*Room
//     Paths         []Path
//     StartRoomLine int
//     EndRoomLine   int
// }

type Link struct {
    Room1 *Room
    Room2 *Room
}

type Path struct {
    Rooms    []*Room
    Shortest []Path
    Queue    []*Ant
    NumNamlaty int
    // Skip       bool
    // Steps      int  // Add this field
    Ants       []*Ant  // Add this field
}


type Farm struct {
    Rooms         []*Room
    Ants          []*Ant
    Links         []*Link
    NumAnt        int
    StartRoom     *Room
    EndRoom       *Room
    AntPositions  map[int]*Room
    Paths         []*Path
    StartRoomLine int
    EndRoomLine   int
}



//no # no L and have no spaces
//tunnel for a room that is not here