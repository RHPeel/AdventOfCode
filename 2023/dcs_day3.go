//This doesn't quite work yet but I understand the bug.

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

type schematicSymbol struct {
	symbplType string
	rowMin     int
	rowMax     int
	colMin     int
	colMax     int
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

func doesNotcontainItem(itemToFind int, slice []int) bool {
	for _, v := range slice {
		if v == itemToFind {
			return false
		}
	}
	return true
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
/* func listmaker(a []uint8) [][]string {
	mySchematic := make([][]string, 0)
	b := strings.Split(string(a), "\n")
	fmt.Println("\n", b)
	for item := range b {
		d := strings.Split(b[item], "")
		if len(d) > 0 {
			mySchematic = append(mySchematic, d)
		}
		fmt.Println("\n", d)
	}
	return mySchematic
}
*/

func listmaker(a []uint8) []string {
	b := strings.Split(string(a), "\n")
	return b
}

func build_number_set(a []string) map[[2]int]int {
	inProgressNumber := ""
	myNumberSet := make(map[[2]int]int)
	for j := 0; j < len(a); j++ {
		for i := 0; i < len(a[j]); i++ {
			fmt.Println(i)
			myChar := rune(a[j][i])
			//fmt.Println(myChar)
			switch {
			case unicode.IsDigit(myChar):
				inProgressNumber = inProgressNumber + string(myChar)
			default:
				if inProgressNumber != "" {
					//fmt.Println(i)
					//fmt.Println(inProgressNumber)func build_number_set(a []string) map[[2]int]int {
					inProgressNumber := ""
					myNumberSet := make(map[[2]int]int)
					for j := 0; j < len(a); j++ {
						for i := 0; i < len(a[j]); i++ {
							myChar := rune(a[j][i])
							//fmt.Println(myChar)
							switch {
							case unicode.IsDigit(myChar):
								inProgressNumber = inProgressNumber + string(myChar)
								if i == len(a[j])-1 {
									myCoordinate := [2]int{j, i - len(inProgressNumber)}
									myNumberSet[myCoordinate] = string_to_int(inProgressNumber)
									inProgressNumber = ""
								}
							default:
								if inProgressNumber != "" {
									//fmt.Println(i)
									//fmt.Println(inProgressNumber)
									myCoordinate := [2]int{j, i - len(inProgressNumber)}
									//fmt.Println(i - len(inProgressNumber))
									myNumberSet[myCoordinate] = string_to_int(inProgressNumber)
									inProgressNumber = ""
								}
							}

						}
					}
					return myNumberSet

					myCoordinate := [2]int{j, i - len(inProgressNumber)}
					//fmt.Println(i - len(inProgressNumber))
					myNumberSet[myCoordinate] = string_to_int(inProgressNumber)
					inProgressNumber = ""
				}
			}

		}
	}
	return myNumberSet
}

func iterate_through(a []string, b map[[2]int]int) int {
	myAdjacentSum := 0
	overallMaxRow := len(a) - 2
	//fmt.Println("Overall Max Row: ", overallMaxRow)
	overallMaxCol := len(a[0]) - 1
	//fmt.Println("Overall Max Col: ", overallMaxCol)
	//fmt.Println("hello")
	var minR, maxR, minC, maxC int
	//myDecision := 0
	for myPosition, actualValue := range b {
		myDecision := 0
		//fmt.Println(myPosition)
		//fmt.Println(actualValue)
		if myPosition[0] == 0 {
			minR = 0
		} else {
			minR = myPosition[0] - 1
		}
		if myPosition[0] == overallMaxRow {
			maxR = overallMaxRow
		} else {
			maxR = myPosition[0] + 1
		}
		if myPosition[1] == 0 {
			minC = 0
		} else {
			minC = myPosition[1] - 1
		}
		if myPosition[1]+len(int_to_string(actualValue)) == overallMaxCol {
			maxC = overallMaxCol
		} else {
			maxC = myPosition[1] + len(int_to_string(actualValue))
		}
		//fmt.Println("minC: ", minC, "\nmaxC: ", maxC, "\nminR: ", minR, "\nmaxR: ", maxR)
		for j := minR; j <= maxR; j++ {
			for i := minC; i <= maxC; i++ {
				myChar := rune(a[j][i])
				//fmt.Println(string(myChar))
				switch {
				case unicode.IsDigit(myChar):
					continue
				case string(myChar) == ".":
					continue
				default:
					myDecision = 1
				}
			}
		}
		if myDecision == 1 {
			myAdjacentSum += actualValue
			//fmt.Println("add me")
		} //else {
		//fmt.Println("don't add me")
		//}
	}
	return myAdjacentSum
}

func make_list_of_possible_adjacents(myRow int, myCol int, sourceNumber int, self int) []string {
	adjacentSpaces := make([]string, 0)
	r := myRow
	c := myCol
	my_len := len(int_to_string(sourceNumber))
	switch self {
	case 1:
		for x := 0; x < my_len; x++ {
			adjacentSpaces = append(adjacentSpaces, "r"+int_to_string(r)+"c"+(int_to_string(c+x)))
		}
	default:
		adjacentSpaces = append(adjacentSpaces, "r"+(int_to_string(r+1))+"c"+int_to_string(c))
		adjacentSpaces = append(adjacentSpaces, "r"+(int_to_string(r-1))+"c"+int_to_string(c))
		adjacentSpaces = append(adjacentSpaces, "r"+int_to_string(r)+"c"+(int_to_string(c-1)))
		adjacentSpaces = append(adjacentSpaces, "r"+int_to_string(r)+"c"+(int_to_string(c+1)))
		adjacentSpaces = append(adjacentSpaces, "r"+(int_to_string(r+1))+"c"+(int_to_string(c-1)))
		adjacentSpaces = append(adjacentSpaces, "r"+(int_to_string(r-1))+"c"+(int_to_string(c-1)))
		adjacentSpaces = append(adjacentSpaces, "r"+(int_to_string(r+1))+"c"+(int_to_string(c+1)))
		adjacentSpaces = append(adjacentSpaces, "r"+(int_to_string(r-1))+"c"+(int_to_string(c+1)))
	}
	return adjacentSpaces
}

func make_adjacent_spaces_map(a []string, b map[[2]int]int) map[int][]string {
	allPossibleSpaces := make([]string, 0)
	for j := 0; j < len(a); j++ {
		for i := 0; i < len(a[j]); i++ {
			myStringName := "r" + int_to_string(j) + "c" + int_to_string(i)
			allPossibleSpaces = append(allPossibleSpaces, myStringName)
		}
	}
	theActuals := make(map[int][]string)
	for myPosition, actualValue := range b {
		myPossibleSet := make_list_of_possible_adjacents(myPosition[0], myPosition[1], actualValue, 1)
		for _, item1 := range myPossibleSet {
			for _, item2 := range allPossibleSpaces {
				if item1 == item2 {
					theActuals[actualValue] = append(theActuals[actualValue], item1)
				}
			}

		}
	}
	return theActuals
}

func iterate_3b(a []string, b map[int][]string) int {
	myReturnSum := 0
	for j := 0; j < len(a); j++ {
		for i := 0; i < len(a[j]); i++ {
			myChar := string(a[j][i])
			if myChar == "*" {
				myMatchList := make([]int, 0)
				spacesToCheck := make_list_of_possible_adjacents(j, i, 0, 0)
				fmt.Println(spacesToCheck)
				//theDoor := 0
				for myVal, item1 := range b {
					theDoor := 0
					for _, item2 := range item1 {
						for _, item3 := range spacesToCheck {
							if item2 == item3 && theDoor == 0 && doesNotcontainItem(myVal, myMatchList) {
								myMatchList = append(myMatchList, myVal)

							}
						}
					}
				}
				fmt.Println(myMatchList)
				if len(myMatchList) == 2 {
					myReturnSum = myReturnSum + (myMatchList[0] * myMatchList[1])
					fmt.Println(myMatchList[0] * myMatchList[1])
				}
			}
		}
	}
	return myReturnSum
}

func main() {
	inputtext := fileimport("dcs_day3_input.txt")
	instructions := listmaker(inputtext)
	//fmt.Println(instructions)
	//fmt.Println(instructions[0][2])
	dansNumberSet := build_number_set(instructions)
	fmt.Println(dansNumberSet)
	fmt.Println(iterate_through(instructions, dansNumberSet))
	dansNumberSet2 := make_adjacent_spaces_map(instructions, dansNumberSet)
	fmt.Println(dansNumberSet2)
	fmt.Println(iterate_3b(instructions, dansNumberSet2))

}

//B solution
//Iterate through the whole map again.
//On "asterisk," iterate throught the entire myNumberSet.
//If the number matches one of the 8 spaces, count it; increment matchCount.
//If matchCount == 2, add it to your sum.
//Reset matchCount after each asterisk is processed.
