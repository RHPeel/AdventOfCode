package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	//"reflect"
)

//This is going to be how we characterize individual games in our dataset.
//Worth noting that these will eventually be the keys in a map.
type game struct {
	gameNumber int
	redCount   int
	greenCount int
	blueCount  int
	power      int
}

//Generic function I use to do the error handling on strconv.Atoi
func string_to_int(a string) int {
	b, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Error:", err)
	}
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

//Convert the file import into an array; each line gets one entry.
func listmaker(a []uint8) []string {
	b := strings.Split(string(a), "\n")
	return b
}

//This is how we convert the input into the "game" struct.
//Our input is one of the lines from the input file.
func process_game(a string) game {
	//First, eliminate the colons, semicolons, and commas. Spaces are enough.
	a = strings.Replace(a, ":", "", -1)
	a = strings.Replace(a, ";", "", -1)
	a = strings.Replace(a, ",", "", -1)
	//We're going to make the line into an array; each word is its own field.
	b := strings.Split(string(a), " ")
	//Now we'll initiate a game and store it as "p".
	p := game{redCount: 0, blueCount: 0, greenCount: 0, power: 0, gameNumber: 0}
	//Set the "numbercheck" variable to zero.
	myNumberCheck := 0
	//FOR loop that iterates through each item in the array.
	for n := 0; n < len(b); n++ {
		//The switch evaluates how we should handle each case.
		switch b[n] {
		//If the value is "Game" we know we can get the gameNumber in the next spot over.
		case "Game":
			p.gameNumber = string_to_int(b[n+1])
		//If the value is "green," we know that the number in the previous spot is the relevant value.
		case "green":
			myNumberCheck = string_to_int(b[n-1])
			//We just want to retain the highest number displayed in the record.
			if p.greenCount < myNumberCheck {
				p.greenCount = myNumberCheck
			}
		case "blue":
			myNumberCheck = string_to_int(b[n-1])
			if p.blueCount < myNumberCheck {
				p.blueCount = myNumberCheck
			}
		case "red":
			myNumberCheck = string_to_int(b[n-1])
			if p.redCount < myNumberCheck {
				p.redCount = myNumberCheck
			}
		}
	}
	//For 2B we calculate "power" as the product of the three block counts.
	p.power = p.blueCount * p.redCount * p.greenCount
	return p
}

//This is our process for actually taking the struct produced in process_game and making it into the map.
func build_games(a []string) map[game]string {
	//Create our map.
	theGames := make(map[game]string)
	//Iterate through every item in the array; each item is a separate line from the input.
	for x := 0; x < len(a); x++ {
		//process_game returns a game struct, so we'll store that in newGame.
		newGame := process_game(a[x])
		//Once we have the game completed we can determine if it is valid.
		theGames[newGame] = determine_if_game_is_valid(newGame)
	}
	//We should return the map here.
	return theGames
}

//This is pretty straightforward: we default to "valid..."
//... and switch to invalid if any of our counts are higher than the max.
func determine_if_game_is_valid(theGame game) string {
	gameDecision := "valid"
	if theGame.redCount > 12 {
		gameDecision = "invalid"
	}
	if theGame.blueCount > 14 {
		gameDecision = "invalid"
	}
	if theGame.greenCount > 13 {
		gameDecision = "invalid"
	}
	return gameDecision
}

//This was the 2A problem.
//It was to just identify the "valid" games and sum up the game numbers.
//Pretty straightforward program for this. We take the map of games as input and output an integer.
func get_sum_of_valid_games(theGame map[game]string) int {
	mySum := 0
	//FOR loop syntax here: when you iterate on a map, you can designate a variable for key and value.
	//Our key is thisGame; our value is gameDecision.
	for thisGame, gameDecision := range theGame {
		if gameDecision == "valid" {
			mySum += thisGame.gameNumber
		}
	}
	return mySum

}

//This was the 2B problem.
//This is even simpler; we computed "power" during process_games.
//So all this does is sum the powers across every object.
func get_sum_of_powers(theGame map[game]string) int {
	myPower := 0
	//FOR loop syntax here: when you iterate on a map, you can designate a variable for key and value.
	//Our key is thisGame; our value is not included at all because we don't need it.
	for thisGame, _ := range theGame {
		myPower += thisGame.power
	}
	return myPower

}

func main() {
	inputtext := fileimport("dcs_day2_input.txt")
	instructions := listmaker(inputtext)
	myGames := build_games(instructions)
	fmt.Println("\n2A: ", get_sum_of_valid_games(myGames))
	fmt.Println("\n2B: ", get_sum_of_powers(myGames))
}
