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
	LineNum   int // Added LineNum field
}

type Ant struct {
	ID int
}

type Farm struct {
	Rooms         []*Room
	Ants          []*Ant
	Links         []*Link
	NumAnt        int
	StartRoom     *Room
	EndRoom       *Room
	AntPositions  map[int]*Room //track the ant which ant in which room 
	Paths         []*Path
	StartRoomLine int
	EndRoomLine   int
}

type Link struct {
	Room1 *Room
	Room2 *Room
}

type Path struct {
	Rooms []*Room
	Shortest []Path
}

//paths must be same number as links from the start
//filter the shortest path
