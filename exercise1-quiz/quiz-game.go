package main

/*
TODO: Flag for file name
TODO: Flag for timer
TODO: Timer
*/
import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// The flags for the function
	var fFlag = flag.String("f", "problems.csv", "the file to pull quiz from")
	var tFlag = flag.Int("t", 30, "time limit for the quiz")
	flag.Parse()

	// Reading in the quiz and initializing the metrics
	file, _ := os.Open(*fFlag)
	r := csv.NewReader(file)
	data, _ := r.ReadAll()
	correct, total := 0, len(data)

	// Prompting the user to start the quiz
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Press Enter to begin the quiz.")
	scanner.Scan()

	// Setting up the timer
	dur := time.Duration(*tFlag) * time.Second
	f := func() { report(correct, total); os.Exit(0) }
	time.AfterFunc(dur, f)

	// Loop through
	for i := 0; i < total; i++ {
		// Read question and answer
		var question, user_answer, real_answer string
		record := data[i]
		question, real_answer = record[0], record[1]

		fmt.Printf("%s?: ", question)
		fmt.Scan(&user_answer)

		if user_answer == real_answer {
			correct++
		}
	}
	report(correct, total)
}

func report(correct int, total int) {
	fmt.Printf("Correct: %d, total: %d\n", correct, total-correct)
}
