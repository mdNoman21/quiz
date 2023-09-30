package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	problemsM := make(map[string]string)
	problems := make([]string, 0, len(problemsM))

	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	fmt.Println("Loading problems...")
	for {
		problem, err := reader.Read()
		if err != nil {
			break
		}
		problemsM[problem[0]] = problem[1]
	}
	for key := range problemsM {
		problems = append(problems, key)
	}

	correctAnswers := 0
	scanner := bufio.NewScanner(os.Stdin)
	n := 4
	x := n
	for x > 0 {
		randomIndex := rand.Intn(len(problems))
		randomProblem := problems[randomIndex]
		fmt.Printf("%s:", randomProblem)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		answer := scanner.Text()
		if problemsM[randomProblem] == answer {
			correctAnswers += 1
		}
		x--
	}
	fmt.Printf("%v Correct answers out of %v", correctAnswers, n)
}
