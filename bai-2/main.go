package main

import (
	"flag"
	"fmt"
)

var koreanGreeting string

func main() {
	lang := flag.String("lang", "en", "greeting language")
	flag.Parse()

	koreanGreeting = "안녕, 친구"
	var enGreeting = "Hello"
	chinaGreeting := "你好朋友"

	if *lang == "en" {
		fmt.Println(enGreeting)
	} else if *lang == "korean" {
		fmt.Println(koreanGreeting)
	} else if *lang == "china" {
		fmt.Println(chinaGreeting)
	} else {
		fmt.Println("Language is not support.")
	}
}
