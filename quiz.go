// A quiz program that runs a quiz from questions and answers read from a CSV file
package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// return slice of slices (list of lists)
func ReadCsvFile(file_path string) [][]string {
	// read file_path to memory
	f, _ := os.Open(file_path)

	defer f.Close()

	// use csv package to process file
	csv_file := csv.NewReader(f)
	records, _ := csv_file.ReadAll()

	return records
}

func main() {
	// TODO: input flag for specifying csv filepath
	// TODO: input flag for number of questions
	// TODO: input flag for time limit (default 30s)
	// TODO: input flag to shuffle the quiz order

	// read in the csv file
	csv_file := ReadCsvFile("problems.csv")

	// initialise counters of correct/incorrect answers
	correct := 0
	wrong := 0

	// loop over each row in csv file
	for i, v := range csv_file {
		// print the question to stdout
		fmt.Printf("%v: %v\n", i, v[0])

		// read user input from stdin
		// var answer int
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("-> ")
		answer, _ := reader.ReadString('\n')
		answer = strings.Replace(answer, "\n", "", -1)

		// str to int using "ASCII to integer"
		solution, _ := strconv.Atoi(v[1])
		given_answer, _ := strconv.Atoi(answer)

		if solution == given_answer {
			correct += 1
		} else {
			wrong += 1
		}
	}

	fmt.Println("---------------------")
	fmt.Println("RESULTS:")
	fmt.Println("---------------------")
	fmt.Printf("Correct: %v \t Wrong: %v\n", correct, wrong)
}
