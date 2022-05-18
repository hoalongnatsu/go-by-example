package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var UserOneName, UserTwoName string
	fmt.Print("Enter name for user one: ")
	fmt.Scan(&UserOneName)
	fmt.Print("Enter name for user two: ")
	fmt.Scan(&UserTwoName)

	UserOneTotalStep, UserTwoTotalStep := 0, 0

	for round := 1; round <= 10; round++ {
		fmt.Println("========")
		fmt.Printf("Ready for round %d ...\n", round)

		UserOneTotalStep += UserGoToStep(UserOneName)
		UserTwoTotalStep += UserGoToStep(UserTwoName)

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("========")
	if UserOneTotalStep > UserTwoTotalStep {
		fmt.Printf("User %s win with %d step\n", UserOneName, UserOneTotalStep)
	} else {
		fmt.Printf("User %s win with %d step\n", UserTwoName, UserTwoTotalStep)
	}
}

func UserGoToStep(user string) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	step := r.Intn(10)
	fmt.Printf("User %s go %d step\n", user, step)

	return step
}
