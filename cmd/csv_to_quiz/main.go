package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/inancgumus/screen"
)

func loadCsv() (dataframe.DataFrame, error) {
	/* Load a CSV file into a DataFrame */

	// Identify CSV
	var csvPath string
	flag.StringVar(&csvPath, "c", "assets/climate_action_quiz.csv", "Specify CSV file path.")
	flag.Parse()

	// Load CSV
	csvfile, err := os.Open(csvPath)
	if err != nil {
		fmt.Printf("Unable to open %s!\n", csvPath)
		return dataframe.New(), err
	}
	return dataframe.ReadCSV(csvfile), err
}

func checkAnswer(userResponse string, answer string) bool {
	/* Compare the User response to the answer */

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
	/* Display a quiz to the User */

	// Load the default or User supplied CSV into a DataFrame
	screen.Clear()
	screen.MoveTopLeft()
	quizDF, err := loadCsv()
	if err != nil {
		fmt.Println("Exiting.")
		return
	}

	// Iterate over questions, skipping header row
	questionCount := len(quizDF.Records())
	correctAnswers := 0
	for i := 0; i < len(quizDF.Records()); i++ {
		record := quizDF.Records()[i]

		// Skip header row when present
		if record[0] == "question" {
			questionCount -= 1
			continue
		}

		// Display question and request input
		fmt.Printf("Question #%d of #%d: %s ", i, questionCount, record[0])
		var response string
		fmt.Scan(&response)

		// Check answer if applicable and document result
		correct := checkAnswer(response, record[1])
		if correct {
			correctAnswers += 1
		}

		// Display note(s) when available
		if record[2] != "" {
			fmt.Printf("%s\n", record[2])
		}
		fmt.Println("")

	}
	fmt.Printf("Final Grade: %.2f%%\n", float64(correctAnswers)/float64(questionCount)*100)
}
