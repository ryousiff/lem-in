package lem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func File(file string) *Farm {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("ERROR: invalid data format, unable to open file")
		os.Exit(1)
	}
	defer f.Close()

	var farm Farm
	farm.AntPositions = make(map[int]*Room)
	farm.RoomNames = make(map[string]bool) // Initialize the map to track room names
	lineNum := 1

	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		farm.NumAnt, err = strconv.Atoi(scanner.Text())
		if err != nil || farm.NumAnt <= 0 {
			fmt.Println("ERROR: invalid data format, invalid number of ants")
			os.Exit(1)
		}
		lineNum++
	}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		switch {
		case len(fields) == 0:
			fmt.Printf("Warning: Empty line on line %d: %s\n", lineNum, line)
		case len(fields) == 1:
			handleSingleField(&farm, line, lineNum)
		case len(fields) == 3:
			room := NewRoom(line, lineNum)
			if room != nil {
				if _, exists := farm.RoomNames[room.Name]; exists {
					fmt.Printf("ERROR: duplicate room name '%s' found on line %d\n", room.Name, lineNum)
					os.Exit(1)
				}
				farm.Rooms = append(farm.Rooms, room)
				farm.RoomNames[room.Name] = true
			}
		default:
			fmt.Printf("ERROR: invalid data format on line %d: %s\n", lineNum, line)
			os.Exit(1)
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR: invalid data format, error reading file")
		os.Exit(1)
	}

	for _, link := range farm.Links {
		link.Room1.Links = append(link.Room1.Links, link)
		link.Room2.Links = append(link.Room2.Links, link)
	}

	for _, room := range farm.Rooms {
		if farm.StartRoomLine > 0 && farm.StartRoomLine == room.LineNum {
			farm.StartRoom = room
			room.IsStart = true
		}
		if farm.EndRoomLine > 0 && farm.EndRoomLine == room.LineNum {
			farm.EndRoom = room
			room.IsEnd = true
		}
	}

	if farm.StartRoom == nil || farm.EndRoom == nil {
		fmt.Println("ERROR: invalid data format, no start or end room found")
		os.Exit(1)
	}

	return &farm
}

func handleSingleField(farm *Farm, line string, lineNum int) {
	switch line {
	case "##start":
		farm.StartRoomLine = lineNum + 1 // Set StartRoomLine to the next line
	case "##end":
		farm.EndRoomLine = lineNum + 1 // Set EndRoomLine to the next line
	default:
		link := NewLink(line, lineNum, farm)
		if link != nil {
			farm.Links = append(farm.Links, link)
		}
	}
}

func NewRoom(line string, lineNum int) *Room {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		fmt.Printf("ERROR: invalid data format on line %d: %s\n", lineNum, line)
		os.Exit(0)
		return nil
	}

	name := fields[0]
	coordX := fields[1]
	coordY := fields[2]

	// Validate coordinates
	_, errX := strconv.Atoi(coordX)
	_, errY := strconv.Atoi(coordY)
	if errX != nil || errY != nil {
		fmt.Printf("ERROR: invalid data format on line %d: %s\n", lineNum, line)
		os.Exit(0)
		return nil
	}

	newRoom := &Room{
		Name:    name,
		CoordX:  coordX,
		CoordY:  coordY,
		Visited: false,
		IsEnd:   false,
		IsStart: false,
		Links:   make([]*Link, 0),
		LineNum: lineNum, // Set the LineNum field
	}

	return newRoom
}

func NewLink(line string, lineNum int, farm *Farm) *Link {
	linkSplit := strings.Split(line, "-")
	if len(linkSplit) != 2 {
		fmt.Printf("ERROR: invalid data format on line %d: %s\n", lineNum, line)
		os.Exit(0)
		return nil
	}

	roomName1 := linkSplit[0]
	roomName2 := linkSplit[1]

	room1 := findRoomByName(roomName1, farm.Rooms)
	room2 := findRoomByName(roomName2, farm.Rooms)

	if room1 == nil || room2 == nil {
		fmt.Printf("ERROR: invalid data format, invalid room name(s) on line %d: %s\n", lineNum, line)
		os.Exit(0)
		return nil
	} else if room1 == room2 {
		fmt.Printf("ERROR: invalid data format, link to itself on line %d: %s\n", lineNum, line)
		os.Exit(0)
		return nil
	}

	newLink := &Link{
		Room1: room1,
		Room2: room2,
	}

	return newLink
}

func findRoomByName(name string, rooms []*Room) *Room {
	for _, room := range rooms {
		if room.Name == name {
			return room
		}
	}
	return nil
}
