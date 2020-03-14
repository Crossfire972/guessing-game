package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

var (
	difficulty  string
	answerRange int
)

func generateValidRange(start, end int) []int {
	s := make([]int, end-start+1)
	for i := range s {
		s[i] = start
		start++
	}
	return s
}

func validateAnswer(valid []int, answer int) bool {
	for _, i := range valid {
		if answer == i {
			return true
		}
	}
	return false
}

func calculateChances(diff int) int {
	var i float64
	var found bool
	for !found {
		if int(math.Pow(2, i)) > diff {
			break
		}
		i++
	}
	return int(i)
}

func doRetry() bool {
	val := getInput()
	switch val.(string) {
	case "y":
		return true
	case "Y":
		return true
	case "Yes":
		return true
	case "yes":
		return true
	}
	return false
}

func getInput() interface{} {
	var answer string
	for answer == "" {
		_, err := fmt.Scanln(&answer)
		if err != nil {
			fmt.Printf("Error: invalid input\n")
		}
	}
	return answer
}

func getAnswer(start int, end int) interface{} {
	var answer int = 0
	for answer == 0 {
		fmt.Printf("Your answer: ")
		a := getInput().(string)
		answer, _ = strconv.Atoi(a)
		valid := generateValidRange(start, end)
		validated := validateAnswer(valid, answer)
		if !validated {
			fmt.Printf("Invalid value provided, must one of %v\n", valid)
			answer = 0
		}
	}
	fmt.Println("")
	return answer
}

func setDifficulty(val int) {
	switch val {

	case 1:
		difficulty = "Easy"
		answerRange = 15
	case 2:
		difficulty = "Medium"
		answerRange = 300
	case 3:
		difficulty = "Hard"
		answerRange = 600
	default:
		panic("Value out of range provided")
	}
}

func gameRun() {
	var retry bool = true
	for retry {
		var down = 0
		var up = answerRange
		// a := generateValidRange(0, answerRange)
		chances := calculateChances(answerRange)
		fmt.Printf("The computer have chosen a number between 0 and %d.\n", answerRange)
		fmt.Printf("You have a maximum of %d guess to find out which one it is.\n", chances)
		cpuChoice := rand.Intn(answerRange)
		for chances > 0 {
			fmt.Printf("Enter your guess between [%d and %d]: ", down, up)
			var answer int = 0
			a := getInput().(string)
			answer, _ = strconv.Atoi(a)
			if answer > cpuChoice {
				chances--
				fmt.Printf("No ! my choice is lower than %d [%d choices left]\n", answer, chances)
				up = answer
			}
			if answer < cpuChoice {
				chances--
				fmt.Printf("No ! my choice is greater than %d [%d choices left]\n", answer, chances)
				down = answer
			}
			if answer == cpuChoice {
				// Display Win message
				fmt.Printf("Congratulation, my choice was %d.\n", cpuChoice)
				fmt.Printf("Youve found it in %d attempt\n", calculateChances(answerRange)-chances+1)
				break
			}
			if chances == 0 {
				// Display Lose message
				fmt.Printf("Too Bad you lose, my choice was %d.\n", cpuChoice)
			}
		}
		// Prompt Retry
		fmt.Println("")
		fmt.Printf("Do you want to try again ? (enter yes/no): ")
		retry = doRetry()
	}
}

func main() {
	// Welcome message
	fmt.Println("Guessing game")
	fmt.Println("In this game you will have to guess what number the computer have chosen in a maximum of guess.")
	fmt.Println("While providing a guess the computer will tell you if it's choice is lower or greater than your guess.")
	// Select Game difficulty
	fmt.Println("Choose your Level")
	fmt.Println("1 - Easy")
	fmt.Println("2 - Medium")
	fmt.Println("3 - Hard")
	answer := getAnswer(1, 3)
	setDifficulty(answer.(int))
	// Enter Game
	gameRun()
	// Exit game
	os.Exit(0)
}
