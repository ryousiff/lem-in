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
		fmt.Println(err.Error())
		return nil
	}
	defer f.Close()

	var farm Farm
	farm.AntPositions = make(map[int]*Room)
	lineNum := 1

	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		farm.NumAnt, err = strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println(err.Error())
			return nil
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
				farm.Rooms = append(farm.Rooms, room)
			}
		default:
			fmt.Printf("Invalid line format on line %d: %s\n", lineNum, line)
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
	}

	for _, link := range farm.Links {
		link.Room1.Links = append(link.Room1.Links, link)
		link.Room2.Links = append(link.Room2.Links, link)
	}
	return &farm
}

func handleSingleField(farm *Farm, line string, lineNum int) {
	switch line {
	case "##start":
		if lineNum < len(farm.Rooms) {
			farm.StartRoom = farm.Rooms[lineNum-1]
			farm.StartRoom.IsStart = true
		} else {
			fmt.Printf("Start room not found on line %d: %s\n", lineNum, line)
		}
	case "##end":
		if lineNum < len(farm.Rooms) {
			farm.EndRoom = farm.Rooms[lineNum]
			farm.EndRoom.IsEnd = true
		} else {
			fmt.Printf("End room not found on line %d: %s\n", lineNum, line)
		}
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
		fmt.Printf("Invalid room definition on line %d: %s\n", lineNum, line)
		return nil
	}

	name := fields[0]
	coordX := fields[1]
	coordY := fields[2]

	// Validate room name
	if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") || strings.Contains(name, " ") {
		fmt.Printf("Invalid room name on line %d: %s\n", lineNum, line)
		return nil
	}

	// Validate coordinates
	_, errX := strconv.Atoi(coordX)
	_, errY := strconv.Atoi(coordY)
	if errX != nil || errY != nil {
		fmt.Printf("Invalid room coordinates on line %d: %s\n", lineNum, line)
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
	}

	return newRoom
}

func NewLink(line string, lineNum int, farm *Farm) *Link {
	linkSplit := strings.Split(line, "-")
	if len(linkSplit) != 2 {
		fmt.Printf("Invalid link definition on line %d: %s\n", lineNum, line)
		return nil
	}

	roomName1 := linkSplit[0]
	roomName2 := linkSplit[1]

	// Validate room names
	// if strings.HasPrefix(roomName1, "L") || strings.HasPrefix(roomName1, "#") || strings.Contains(roomName1, " ") {
	//     fmt.Printf("Invalid room name on line %d: %s\n", lineNum, line)
	//     return nil
	// }
	// if strings.HasPrefix(roomName2, "L") || strings.HasPrefix(roomName2, "#") || strings.Contains(roomName2, " ") {
	//     fmt.Printf("Invalid room name on line %d: %s\n", lineNum, line)
	//     return nil
	// }

	room1 := findRoomByName(roomName1, farm.Rooms)
	room2 := findRoomByName(roomName2, farm.Rooms)

	if room1 == nil || room2 == nil {
		fmt.Printf("Invalid room name(s) on line %d: %s\n", lineNum, line)
		return nil
	} else if room1 == room2 {
		fmt.Printf("link to itself %d: %s\n", lineNum, line)
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
