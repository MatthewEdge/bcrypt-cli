package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	text, _, err := in.ReadLine()
	if err != nil {
		panic("Failed to read from stdin: " + err.Error())
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to hash stdin: " + err.Error())
	}

	fmt.Println(string(hashed))
}
