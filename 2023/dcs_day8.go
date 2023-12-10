package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

type node struct {
	designator      string
	leftDirection   string
	rightDirection  string
	repeatTurnCount int
	endsInZ         []int
}

//Generic function I use to do the error handling on strconv.Atoi
func string_to_int(a string) int {
	b, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return b
}

func int_to_string(c int) string {
	d := strconv.Itoa(c)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//}
	return d
}

func isDigit(char rune) bool {
	return unicode.IsDigit(char)
}

func containsItem(itemToFind int, slice []int) bool {
	for _, v := range slice {
		if v == itemToFind {
			return true
		}
	}
	return false
}

func power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func listmaker(a []uint8) []string {
	b := strings.Split(string(a), "\n")
	return b
}

//My standard file import.
func fileimport(a string) []uint8 {

	b, err := ioutil.ReadFile(a)

	if err != nil {
		fmt.Print(err)
	}
	return b
}

func build_node_map(a []string) map[string]node {
	myNodes := make(map[string]node)
	for i := 2; i < len(a); i++ {
		singleNode := node{designator: a[i][0:3], leftDirection: a[i][7:10], rightDirection: a[i][12:15]}
		myNodes[a[i][0:3]] = singleNode
	}
	return myNodes
}

func execute_turns(myTurns string, myNodeMap map[string]node) int {
	currentLocation := "AAA"
	targetLocation := "ZZZ"
	turnCount := 0
	for i := 0; i < len(myTurns); i++ {
		switch string(myTurns[i]) {
		case "L":
			currentLocation = myNodeMap[currentLocation].leftDirection
		case "R":
			currentLocation = myNodeMap[currentLocation].rightDirection
		}
		turnCount++
		if currentLocation == targetLocation {
			break
		}
		if i+1 == len(myTurns) {
			i = -1
		}
	}
	return turnCount
}

func build_target_set(myNodeMap map[string]node, myTarget string) []string {
	currentLocations := make([]string, 0)
	for designator, _ := range myNodeMap {
		if string(designator[2]) == myTarget {
			currentLocations = append(currentLocations, string(designator))
		}
	}
	return currentLocations
}

func get_loop_details(myTurns2 string, currentNode node, myNodeMap map[string]node) node {
	loopNumberKnown := 0
	haveZs := 0
	turnCount := 0
	currentLocation := currentNode.designator
	locationHistory := make([]string, 0)
	for i := 0; i < len(myTurns2); i++ {
		switch string(myTurns2[i]) {
		case "L":
			currentLocation = myNodeMap[currentLocation].leftDirection
		case "R":
			currentLocation = myNodeMap[currentLocation].rightDirection
		}
		locationHistory = append(locationHistory, currentLocation)
		if turnCount > len(myTurns2) {
			if locationHistory[turnCount%(len(myTurns2))] == currentLocation {
				loopNumberKnown = 1
				currentNode.repeatTurnCount = turnCount
				break
			}
		}
		turnCount++
		if string(currentLocation[2]) == "Z" {
			haveZs = 1
			currentNode.endsInZ = append(currentNode.endsInZ, turnCount)
		}
		if i+1 == len(myTurns2) {
			i = -1
		}
		if loopNumberKnown == 1 && haveZs == 1 {
			fmt.Println(locationHistory)
			break
		}
	}
	return currentNode
}

func execute_turns_8b(myTurns string, myNodeMap map[string]node) map[string]node {
	currentLocations := build_target_set(myNodeMap, "A")
	myNodeMapZ := make(map[string]node)
	for _, item := range currentLocations {
		updatedNode := get_loop_details(myTurns, myNodeMap[item], myNodeMap)
		myNodeMapZ[item] = updatedNode
	}
	return myNodeMapZ
}

func get_mins(a map[string]node) []int {
	min_list := make([]int, 0)
	for _, item := range a {
		myList := item.endsInZ
		min_list = append(min_list, findMin(myList))
	}
	return min_list
}

func findMin(slice []int) (min int) {
	if len(slice) == 0 {
		// Handle empty slice case
		return 0
	}

	// Initialize min with the first element of the slice
	min = slice[0]

	// Iterate through the slice starting from the second element
	for _, value := range slice[1:] {
		if value < min {
			min = value
		}
	}

	return min
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to calculate the least common multiple (LCM) of two numbers
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// Function to calculate the least common multiple (LCM) of multiple numbers
func calculateLCM(numbers []int) int {
	if len(numbers) < 2 {
		// Handle case with less than two numbers
		return 0
	}

	// Initialize lcm with the LCM of the first two numbers
	result := lcm(numbers[0], numbers[1])

	// Iterate through the remaining numbers
	for i := 2; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	return result
}

func main() {
	inputtext := fileimport("dcs_day8_input.txt")
	instructions := listmaker(inputtext)
	turns := instructions[0]
	//fmt.Println("turns: ", turns)
	myNodes := build_node_map(instructions)
	fmt.Println("8A: ", execute_turns(turns, myNodes))
	nodes8b := execute_turns_8b(turns, myNodes)
	minZs := get_mins(nodes8b)
	fmt.Println("8B: ", calculateLCM(minZs))
}
