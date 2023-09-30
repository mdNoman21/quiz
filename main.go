package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	problemsM := make(map[string]string)
	problems := make([]string, 0, len(problemsM))
	rand.Seed(time.Now().UnixNano())

	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		problem, err := reader.Read()
		if err != nil {
			break
		}
		problemsM[problem[0]] = problem[1]
		problems = append(problems, problem[0])
	}

	correctAnswers := 0
	totalQuestions := 5
	timeLimit := 5

	scanner := bufio.NewScanner(os.Stdin)
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	fmt.Println("Loading problems...")

	for totalQuestions > 0 {
		randomIndex := rand.Intn(len(problems))
		randomProblem := problems[randomIndex]
		fmt.Printf("%s:", randomProblem)
		answerCh := make(chan string)
		scanner.Scan()

		go func() {
			answerCh <- scanner.Text()
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime's up!")
			fmt.Printf("%v Correct answers out of %v\n", correctAnswers, totalQuestions)
			return
		case answer := <-answerCh:
			answer = strings.TrimSpace(answer)
			if problemsM[randomProblem] == answer {
				correctAnswers++
			}
		}
		totalQuestions--

	}

	fmt.Printf("%v Correct answers out of %v", correctAnswers, totalQuestions)
}
