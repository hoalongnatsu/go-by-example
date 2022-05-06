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

	var answer int
	var err error
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Please enter your answer with numer 0 - 2: ")
		scanner.Scan()
		input := scanner.Text()
		answer, err = strconv.Atoi(input)

		if err != nil {
			log.Panic(err)
		}

		if answer >= 0 && answer <= 2 {
			break
		}
	}

	var secretPerson string
	switch secretPersonNumber {
	case 0:
		secretPerson = personZero
	case 1:
		secretPerson = personOne
	default:
		secretPerson = personTwo
	}

	if answer == secretPersonNumber {
		fmt.Printf("Congratulations!! You answer is correct, %s is the secret person.\n", secretPerson)
	} else {
		fmt.Printf("Sorry!! You answer is incorrect, %d: %s is the secret person.\n", secretPersonNumber, secretPerson)
	}
}
