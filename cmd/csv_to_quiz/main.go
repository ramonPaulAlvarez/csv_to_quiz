package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-gota/gota/dataframe"
)

func checkAnswer(userResponse string, answer string) bool {
	// Convert response and answer to boolean
	var responseBool bool
	if strings.ToLower(userResponse) == "y" || strings.ToLower(userResponse) == "yes" {
		responseBool = true
	} else {
		responseBool, _ = strconv.ParseBool(userResponse)
	}
	answerBool, _ := strconv.ParseBool(answer)

	// Provide answer comparison result
	return answerBool == responseBool
}

func main() {
	// Load CSV
	csvfile, _ := os.Open("assets/climate_action_quiz.csv")
	quizDF := dataframe.ReadCSV(csvfile)
	correctAnswers := 0
	questionCount := len(quizDF.Records())

	// Iterate over questions, skipping header row
	for i := 0; i < len(quizDF.Records()); i++ {
		record := quizDF.Records()[i]

		if record[0] == "question" {
			questionCount -= 1
			continue
		}

		// Display question to User
		fmt.Printf("Question #%d of #%d: %s ", i, questionCount, record[0])

		// Request User input
		var response string
		fmt.Scan(&response)

		// Check answer and document result
		correct := checkAnswer(response, record[1])
		if correct {
			correctAnswers += 1
		}

		// Display question note(s) when available
		if record[2] != "" {
			fmt.Println(record[2])
		}

	}
	fmt.Printf("Final Grade: %.2f%%\n", float64(correctAnswers)/float64(questionCount)*100)
}
