package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var file *os.File
var err error

func andPull(opponentTurn, myTurn string) int {
	// decide win, lose, draw and return points
	var points int
	if (opponentTurn == "A") && (myTurn == "X") {
		points = 4
	}
	if (opponentTurn == "A") && (myTurn == "Y") {
		points = 8
	}
	if (opponentTurn == "A") && (myTurn == "Z") {
		points = 3
	}
	if (opponentTurn == "B") && (myTurn == "X") {
		points = 1
	}
	if (opponentTurn == "B") && (myTurn == "Y") {
		points = 5
	}
	if (opponentTurn == "B") && (myTurn == "Z") {
		points = 9
	}
	if (opponentTurn == "C") && (myTurn == "X") {
		points = 7
	}
	if (opponentTurn == "C") && (myTurn == "Y") {
		points = 2
	}
	if (opponentTurn == "C") && (myTurn == "Z") {
		points = 6
	}
	return points
}

func whatsMyGo(opponentTurn, myTurn string) string {
	// X = lose, Y = Draw, Z = Win
	var myGo string
	if (opponentTurn == "A") && (myTurn == "X") {
		myGo = "Z"
	}
	if (opponentTurn == "A") && (myTurn == "Y") {
		myGo = "X"
	}
	if (opponentTurn == "A") && (myTurn == "Z") {
		myGo = "Y"
	}
	if (opponentTurn == "B") && (myTurn == "X") {
		myGo = "X"
	}
	if (opponentTurn == "B") && (myTurn == "Y") {
		myGo = "Y"
	}
	if (opponentTurn == "B") && (myTurn == "Z") {
		myGo = "Z"
	}
	if (opponentTurn == "C") && (myTurn == "X") {
		myGo = "Y"
	}
	if (opponentTurn == "C") && (myTurn == "Y") {
		myGo = "Z"
	}
	if (opponentTurn == "C") && (myTurn == "Z") {
		myGo = "X"
	}
	return myGo
}

func checkRucksacks(rucksacks [3]string) int {
	var ruckSack1 string
	var ruckSack2 string
	var ruckSack3 string

	// find the shortest string
	r1 := len(rucksacks[0])
	r2 := len(rucksacks[1])
	r3 := len(rucksacks[2])

	if (r1 <= r2) && (r1 <= r3) {
		ruckSack1 = rucksacks[0]
		ruckSack2 = rucksacks[1]
		ruckSack3 = rucksacks[2]
	} else if (r2 <= r1) && (r2 <= r3) {
		ruckSack1 = rucksacks[1]
		ruckSack2 = rucksacks[0]
		ruckSack3 = rucksacks[2]
	} else {
		ruckSack1 = rucksacks[2]
		ruckSack2 = rucksacks[0]
		ruckSack3 = rucksacks[1]
	}

	tableMap := make(map[string]int)
	aValue := 1
	for i := 97; i < 123; i++ {
		tableMap[string(i)] = aValue
		aValue += 1
	}
	for i := 65; i < 91; i++ {
		tableMap[string(i)] = aValue
		aValue += 1
	}
	var value int = 0
	for i, char := range ruckSack1 {
		//fmt.Printf("Value of char is %d which is a %s\n", char, string(char))
		if strings.Count(ruckSack2, string(char)) > 0 { //found a character in the second rucksack
			if strings.Count(ruckSack3, string(char)) > 0 { //found a character in all three rucksacks
				value = tableMap[string(ruckSack1[i])]
				//fmt.Print(value)
				break
			}
		}
	}
	return value
}

func dayOne() {
	file, err = os.Open("puzzle1_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var elves []int
	scanner := bufio.NewScanner(file)
	runningTotal := 0
	linesRead := 0

	for scanner.Scan() {
		line := scanner.Text()
		linesRead += 1
		if line == "" {
			elves = append(elves, int(runningTotal))
			runningTotal = 0
		} else {
			number, _ := strconv.Atoi(line)
			runningTotal += number
		}
	}
	file.Close()

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	fmt.Printf("Total value of the Elf with most calories is: %v\n", elves[0])
	fmt.Printf("Total value of the 3 Elves with most calories is: %v\n", elves[0]+elves[1]+elves[2])
}

func dayTwo() {
	fmt.Print("\n\nMoving onto the Rock, paper and scissors game\n")

	// Read in the strategy guide line by line
	file, err = os.Open("stratguide.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner2 := bufio.NewScanner(file)
	linesRead := 0
	myScore := 0
	var wld int

	for scanner2.Scan() {
		line := scanner2.Text()
		linesRead += 1
		opponentTurn := string(line[0])
		myTurn := string(line[2])
		wld = andPull(opponentTurn, myTurn)
		myScore += wld
	}
	file.Close()
	fmt.Printf("My Score after applying secret strategy is :- %v\n", myScore)

	// Read in the strategy guide line by line
	file, err = os.Open("stratguide.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner3 := bufio.NewScanner(file)
	linesRead = 0
	myScore = 0
	wld = 0
	var myNewGo string

	for scanner3.Scan() {
		line := scanner3.Text()
		linesRead += 1
		opponentTurn := string(line[0])
		myTurn := string(line[2])
		myNewGo = whatsMyGo(opponentTurn, myTurn)
		wld = andPull(opponentTurn, myNewGo)
		myScore += wld
	}
	file.Close()
	fmt.Printf("My Score after applying the new secret strategy is :- %v\n", myScore)
}

func dayThree() {
	/*
		Read file in and for each line, split into two equal bits.
		Compare each bit for the single duplicate letter
		Assign a value to the letter based on a matrix/map
		Total up values and display

	*/

	tableMap := make(map[string]int)
	aValue := 1
	for i := 97; i < 123; i++ {
		tableMap[string(i)] = aValue
		aValue += 1
	}
	for i := 65; i < 91; i++ {
		tableMap[string(i)] = aValue
		aValue += 1
	}

	file, err = os.Open("puzzle3.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	runningTotal := 0
	var partA, partB string

	for scanner.Scan() {
		line := scanner.Text()
		partA = line[0:(int(len(line) / 2))]
		partB = line[(int(len(line) / 2)):]
		// Need top loop through partA and see if the character is in partB
		for i, _ := range partA {
			if strings.Contains(partB, string(partA[i])) {
				value := tableMap[string(partA[i])]
				runningTotal += value
				break
			}
		}
	}
	file.Close()
	fmt.Printf("The sum of all the priorities is %d\n", runningTotal)

	// Part two. Need to read 3 lines in, find the duplicate value and reset.
	file, err = os.Open("puzzle3.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner = bufio.NewScanner(file)
	counter := 0
	var rucksacks [3]string
	runningTotal = 0

	for scanner.Scan() {
		line := scanner.Text()
		rucksacks[counter] = line
		if counter == 2 {
			aValue = checkRucksacks(rucksacks)
			runningTotal += aValue
			counter = 0
		} else {
			counter++
		}
	}
	file.Close()
	fmt.Printf("The sum of all the badges priorities is %d\n", runningTotal)
}

func dayFour() {
	// Seems each day we read a file :-)
	file, err = os.Open("puzzle4a.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	//contained := 0
	overlapping := 0

	for scanner.Scan() {
		line := scanner.Text()
		// each line has two parts seperated by a comma - split on that
		p1 := strings.Split(line, "-")
		p2 := strings.Split(p1[1], ",")
		r1l, _ := strconv.Atoi(p1[0])
		r1h, _ := strconv.Atoi(p2[0])
		r2l, _ := strconv.Atoi(p2[1])
		r2h, _ := strconv.Atoi(p1[2])
		/*  Commented out so not to get in the way of part 2 of the puzzle day
		if (r1l >= r2l) && (r1h <= r2h) {
			contained++ // range 1 is inside range 2
			continue    // required to skip around equivelent ranges and prevent doube counting
		}
		if (r2l >= r1l) && (r2h <= r1h) {
			contained++ // range 2 is inside range 1
		}
		*/
		if (r1h < r2l) || (r2h < r1l) {
			continue
		} else {
			overlapping++
		}
	}
	// fmt.Printf("There are %d contained overlapping ranges\n", contained)
	fmt.Printf("There are %d patial overlapping ranges\n", overlapping)
}

func main() {
	dayOne()
	dayTwo()
	dayThree()
	dayFour()
}
