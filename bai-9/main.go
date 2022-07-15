package main

import "fmt"

func main() {
	persons := []string{
		"Ronaldo", "Messi", "Pogba", "Hazard", "Hazard",
		"Messi", "Ronaldo", "Pogba", "Ronaldo", "Messi",
		"Hazard", "Ronaldo", "Ronaldo", "Messi",
	}
	countPerson := map[string]int{}

	for _, p := range persons {
		_, exist := countPerson[p]

		if exist {
			countPerson[p]++
		} else {
			countPerson[p] = 1
		}
	}

	fmt.Printf("%+v\n", countPerson)

	var max int
	var person string

	for key, value := range countPerson {
		if value > max {
			max = value
			person = key
		}
	}

	fmt.Println(person)
}
