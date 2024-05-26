package lem

type Room struct {
    Name     string
    CoordX   string
    CoordY   string
    Visited  bool
    IsEnd    bool
    IsStart  bool
    Links    []*Link
}

type Ant struct {
    ID int
}

type Farm struct {
    Rooms        []*Room
    Ants         []*Ant
    Links        []*Link
    NumAnt       int
    StartRoom    *Room
    EndRoom      *Room
    AntPositions map[int]*Room
}

type Link struct {
    Room1 *Room
    Room2 *Room
}

//next room have to have ##start each next room
//return the error in read file