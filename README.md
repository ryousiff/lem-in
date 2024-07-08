# Lem-in: Digital Ant Farm Simulator
Project Overview
Lem-in is a Go-based program designed to simulate the traversal of ants through a digital ant farm. The objective is to find the quickest path for a given number of ants to travel from the start room to the end room, navigating through a network of rooms connected by tunnels.

# How It Works
The program reads the colony layout from a file, which includes the number of ants, the room definitions, and the tunnel connections. The goal is to efficiently move the ants from the start room (##start) to the end room (##end) while avoiding obstacles such as self-linking rooms and invalid paths. The program must handle various edge cases, including missing start or end rooms, duplicate rooms, and invalid input formats.

# Input Format
The input file is structured as follows:
Number of Ants: The first line contains a single integer representing the number of ants.
Rooms: Each room is defined by a line in the format name coord_x coord_y. The ##start and ##end rooms are marked accordingly.
Links: Each link is defined by a line in the format name1-name2.
Comments: Lines starting with # are comments and should be ignored unless they specify ##start or ##end.

# Example Input
3
##start
1 23 3
2 16 7
#comment
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
#another comment
4-2
2-1
7-6
7-2
7-4
6-5

# Usage
To run the program:
$ go run . <input_file>
