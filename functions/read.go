package lem

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func File(file string) {
    f, err := os.Open(file)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer f.Close()

    var farm Farm
    lineNum := 1

    scanner := bufio.NewScanner(f)
    if scanner.Scan() {
        farm.NumAnt, err = strconv.Atoi(scanner.Text())

        if err != nil {
            fmt.Println(err.Error())
            return
        }
        lineNum++
    }

    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)

        switch {
        case len(fields) == 0:
            fmt.Printf("Field is empty on line %d: %s\n", lineNum, line)
            os.Exit(0)
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
}

func handleSingleField(farm *Farm, line string, lineNum int) {
    switch line {
    case "##start":
        if lineNum < len(farm.Rooms) {
            farm.StartRoom = &Farm.Rooms[lineNum-1]
        } else {
            fmt.Printf("Start room not found on line %d: %s\n", lineNum, line)
        }
    case "##end":
        if lineNum < len(farm.Rooms) {
            farm.EndRoom = &Farm.Rooms[lineNum]
        } else {
            fmt.Printf("End room not found on line %d: %s\n", lineNum, line)
        }
    default:
        link := NewLink(line, lineNum)
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

    newRoom := &Room{
        Name:    name,
        CoordX:  coordX,
        CoordY:  coordY,
        Visited: false,
        IsEnd:   false,
        IsStart: false,
    }

    return newRoom
}

func NewLink(line string, lineNum int) *Link {
    linkSplit := strings.Split(line, "-")
    if len(linkSplit) != 2 {
        fmt.Printf("Invalid link definition on line %d: %s\n", lineNum, line)
        return nil
    }

    roomName := linkSplit[0]
    nextRoomId := linkSplit[1]

    newLink := &Link{
        RoomName:   roomName,
        NextRoomId: nextRoomId,
    }

    return newLink
}
