package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	var personZero, personOne, personTwo string
	personZero, personOne, personTwo = "Max", "Alex", "Tom"

	fmt.Println("Number of secret person 0: Max, 1: Alex, 2: Tom.")

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	secretPersonNumber := r.Intn(3)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter your answer with numer 0 - 2: ")
	scanner.Scan()
	input := scanner.Text()
	answer, err := strconv.Atoi(input)

	if err != nil {
		log.Panic(err)
	}

	if answer == secretPersonNumber {
		if secretPersonNumber == 0 {
			fmt.Printf("Congratulations!! You answer is correct, %s is the secret person.\n", personZero)
		} else if secretPersonNumber == 1 {
			fmt.Printf("Congratulations!! You answer is correct, %s is the secret person.\n", personOne)
		} else if secretPersonNumber == 2 {
			fmt.Printf("Congratulations!! You answer is correct, %s is the secret person.\n", personTwo)
		}
	} else {
		if secretPersonNumber == 0 {
			fmt.Printf("Sorry!! You answer is incorrect, 0: %s is the secret person.\n", personZero)
		} else if secretPersonNumber == 1 {
			fmt.Printf("Sorry!! You answer is incorrect, 1: %s is the secret person.\n", personOne)
		} else if secretPersonNumber == 2 {
			fmt.Printf("Sorry!! You answer is incorrect, 2: %s is the secret person.\n", personTwo)
		}
	}
}

// func main() {
// 	person := [3]string{"Max", "Alex", "Tom"}

// 	source := rand.NewSource(time.Now().UnixNano())
// 	r := rand.New(source)
// 	secretPersonNumber := r.Intn(3)

// 	scanner := bufio.NewScanner(os.Stdin)
// 	fmt.Printf("Enter your answer: ")
// 	scanner.Scan()
// 	input := scanner.Text()
// 	answer, err := strconv.Atoi(input)

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	if answer == secretPersonNumber {
// 		fmt.Printf("Congratulations!! You answer is correct, %s is the secret person.\n", person[answer])
// 	} else {
// 		fmt.Printf("Sorry!! You answer is incorrect, %d: %s is the secret person.\n", answer, person[answer])
// 	}
// }
