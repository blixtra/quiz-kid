package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hi! Let's start quizzing.")
	fmt.Printf("What's your name?: ")

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hi %s! Ready to learn your multiplications?", name[:len(name)-1])

	fmt.Printf("\nWhich number would you like to start with?: ")
	var num int
	fmt.Scanf("%d", &num)

	fmt.Printf("Great, let's multiply %d.\n", num)

	multiplyQuiz(num)
}

func multiplyQuiz(number int) {
	for i := 1; i <= 10; i++ {
		correctAnswer := i * number
		isCorrect := false

		var givenAnswer int
		for !isCorrect {
			fmt.Printf("What is %d * %d?: ", number, i)
			fmt.Scanf("%d", &givenAnswer)
			isCorrect = givenAnswer == correctAnswer
			if !isCorrect {
				fmt.Println("Hmm, let's try that again.")
			}
		}
		fmt.Printf("That's right %d * %d = %d\n", number, i, correctAnswer)
	}

	fmt.Printf("Congratulations, you've finished multiplying %d by numbers 1-10.\n", number)
}
