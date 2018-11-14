package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

// UsernameRegex username regex
const UsernameRegex string = `^\w{1,5}$`

func main() {
	var gopherName string
	flag.StringVar(&gopherName, "gophername", "mike", "the name of the Gopher")
	flag.Parse()
	if checkUsernameSyntax(gopherName) {
		fmt.Println("Hi, " + gopherName)
	} else {
		fmt.Println("Wrong gophername!")
	}

}

func checkUsernameSyntax(username string) bool {
	valid := false
	r, err := regexp.Compile(UsernameRegex)
	if err != nil {
		log.Fatal(err)
	}
	valid = r.MatchString(username)
	return valid

}
