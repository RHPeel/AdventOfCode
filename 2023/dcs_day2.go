package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	//"reflect"
)

type game struct {
	gameNumber int
	redCount   int
	greenCount int
	blueCount  int
	power      int
}

func string_to_int(a string) int {
	b, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return b
}

func fileimport(a string) []uint8 {

	b, err := ioutil.ReadFile(a)

	if err != nil {
		fmt.Print(err)
	}
	return b
}

func listmaker(a []uint8) []string {
	b := strings.Split(string(a), "\n")
	return b
}

func process_game(a string) game {
	a = strings.Replace(a, ":", "", -1)
	a = strings.Replace(a, ";", "", -1)
	a = strings.Replace(a, ",", "", -1)
	b := strings.Split(string(a), " ")
	p := game{redCount: 0, blueCount: 0, greenCount: 0, power: 0}
	myNumberCheck := 0
	for n := 0; n < len(b); n++ {
		//fmt.Println(string(b[n]))
		switch b[n] {
		case "Game":
			p.gameNumber = string_to_int(b[n+1])
		case "green":
			myNumberCheck = string_to_int(b[n-1])
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
	p.power = p.blueCount * p.redCount * p.greenCount
	return p
}

func build_games(a []string) map[game]string {
	theGames := make(map[game]string)
	for x := 0; x < len(a); x++ {
		newGame := process_game(a[x])
		theGames[newGame] = "unchecked"
	}
	return theGames
}

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

func get_sum_of_valid_games(theGame map[game]string) int {
	mySum := 0
	for thisGame, gameDecision := range theGame {
		if gameDecision == "valid" {
			mySum += thisGame.gameNumber
		}
	}
	return mySum

}

func get_sum_of_powers(theGame map[game]string) int {
	myPower := 0
	for thisGame, _ := range theGame {
		myPower += thisGame.power
	}
	return myPower

}

func main() {
	inputtext := fileimport("dcs_day2_input.txt")
	instructions := listmaker(inputtext)
	myGames := build_games(instructions)
	for gameNumber, _ := range myGames {
		myGames[gameNumber] = determine_if_game_is_valid(gameNumber)
	}
	fmt.Println("\n2A: ", get_sum_of_valid_games(myGames))
	fmt.Println("\n2B: ", get_sum_of_powers(myGames))
}
