package lem

type Room struct {
	Name     string
	NextRoom *Room
	CoordX   string
	CoordY   string
	Visited  bool
	IsEnd    bool
	IsStart  bool
}

type Ant struct {
	ID string
}

type Farm struct {
	PathLen int
	RoomNum int
	Rooms   []*Room
	Ants    []*Ant
	Links   []*Link
	NumAnt   int
	StartRoom *Room
	EndRoom *Room
}

type Link struct {
	RoomName string
	Dash string
	NextRoomId string
}
