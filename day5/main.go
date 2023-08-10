package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Towers
var towers [][]string

func moveCrates9000(fromTower int, toTower int) {
	towerLen := len(towers[fromTower])
	if towerLen > 0 {
		moveValue := string(towers[fromTower][towerLen-1])
		towers[fromTower] = towers[fromTower][:towerLen-1]
		towers[toTower] = append(towers[toTower], moveValue)
	}
}
func moveCrates9001(moveQty int, fromTower int, toTower int) {
	towerLen := len(towers[fromTower])
	if towerLen > 0 {
		//moveValue := towers[fromTower][(towerLen - moveQty) : towerLen-1]
		moveValue := towers[fromTower][(towerLen - moveQty):towerLen]
		towers[fromTower] = towers[fromTower][:towerLen-moveQty]
		towers[toTower] = append(towers[toTower], moveValue...)
	}
}

func main() {
	// define up the towers
	tower1 := []string{"J", "H", "P", "M", "S", "F", "N", "V"}
	tower2 := []string{"S", "R", "L", "M", "J", "D", "Q"}
	tower3 := []string{"N", "Q", "D", "H", "C", "S", "W", "B"}
	tower4 := []string{"R", "S", "C", "L"}
	tower5 := []string{"M", "V", "T", "P", "F", "B"}
	tower6 := []string{"T", "R", "Q", "N", "C"}
	tower7 := []string{"G", "V", "R"}
	tower8 := []string{"C", "Z", "S", "P", "D", "L", "R"}
	tower9 := []string{"D", "S", "J", "V", "G", "P", "B", "F"}

	towers = append(towers, tower1, tower2, tower3, tower4, tower5, tower6, tower7, tower8, tower9)

	for i := range towers {
		fmt.Printf("Value of the Tower %d is now %s\n", i, towers[i])
	}
	fmt.Printf("Now to start moving crates\n")

	// Open the file of movements
	var file *os.File
	var err error
	file, err = os.Open("puzzle5_moves.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		moves := strings.Fields(line)
		numMoves, _ := strconv.Atoi(moves[1])
		part2 := true
		if !part2 {
			for i := 1; i <= numMoves; i++ {
				fromtower, _ := strconv.Atoi(moves[3])
				totower, _ := strconv.Atoi(moves[5])
				moveCrates9000(fromtower-1, totower-1)
			}
		} else {
			fromtower, _ := strconv.Atoi(moves[3])
			totower, _ := strconv.Atoi(moves[5])
			moveCrates9001(numMoves, fromtower-1, totower-1)
		}

	}
	for i := range towers {
		fmt.Printf("Value of the Tower %d is now %s\n", i, towers[i])
	}
}
