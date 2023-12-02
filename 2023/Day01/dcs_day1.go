package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
	//"reflect"
)

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

func check_if_string_includes_number(a string) int {
	//fmt.Println("\n", a)
	myMatch := 0
	myNumberArray := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9}
	for key, value := range myNumberArray {
		if strings.Contains(a, key) {
			myMatch = value
			break
		}
	}
	return myMatch
}

func find_first_and_last_numbers(c string) string {
	b := ""
	myString := ""
	myStringCheck := 0
	firstDigit := -1
	lastDigit := -1
	for n := 0; n < len(c); n++ {
		myChar := rune(c[n])
		if unicode.IsDigit(myChar) {
			firstDigit = string_to_int(string(myChar))
			break
		} else {
			myString = myString + string(myChar)
			myStringCheck = check_if_string_includes_number(myString)
			if myStringCheck > 0 {
				firstDigit = myStringCheck
				break
			}
		}
	}
	myString = ""
	for n := len(c) - 1; n >= 0; n-- {
		myChar := rune(c[n])
		if unicode.IsDigit(myChar) {
			lastDigit = string_to_int(string(myChar))
			break
		} else {
			myString = string(myChar) + myString
			myStringCheck = check_if_string_includes_number(myString)
			if myStringCheck > 0 {
				lastDigit = myStringCheck
				break
			}
		}
	}
	b = strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit)
	fmt.Println("\n", b)
	return b
}

func extract_digits(c string) string {
	b := ""
	firstDigit := -1
	lastDigit := -1
	//fmt.Println("\na: ", c)
	for n := 0; n < len(c); n++ {
		myChar := rune(c[n])
		if unicode.IsDigit(myChar) {
			lastDigit = string_to_int(string(myChar))
			if firstDigit == -1 {
				firstDigit = string_to_int(string(myChar))
			}
		}
	}
	//fmt.Println("\nFirstDigit: ", firstDigit)
	//fmt.Println("\nLastDigit: ", lastDigit)
	b = strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit)
	//fmt.Println("/n", b)
	return b
}

func run_through_input_1a(aList []string) int {
	digitsSum := 0
	for n := 0; n < len(aList); n++ {
		digitsSum += string_to_int(extract_digits(aList[n]))
	}
	return digitsSum
}

func run_through_input_1b(aList []string) int {
	digitsSum := 0
	for n := 0; n < len(aList); n++ {
		digitsSum += string_to_int(find_first_and_last_numbers(aList[n]))
	}
	return digitsSum
}

func main() {
	inputtext := fileimport("dcs_day1_input.txt")
	instructions := listmaker(inputtext)
	mySum := run_through_input_1a(instructions)
	fmt.Println("\n1A: ", mySum)
	mySum = run_through_input_1b(instructions)
	fmt.Println("\n1B: ", mySum)

}
