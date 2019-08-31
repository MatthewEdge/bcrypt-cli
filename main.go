package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: echo -n 'STR_TO_HASH' | %s\n", os.Args[0])
		flag.PrintDefaults()
	}

	var compare = flag.String("compare", "", "Compare string with stdin. stdin must be a bcrypt hashed string")
	flag.Parse()

	in := bufio.NewReader(os.Stdin)

	text, _, err := in.ReadLine()
	if err != nil {
		fmt.Println("Failed to read from stdin: " + err.Error())
		os.Exit(1)
	}

	if *compare != "" {
		err := bcrypt.CompareHashAndPassword(text, []byte(*compare))

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		os.Exit(0)
	}

	hashed, err := bcrypt.GenerateFromPassword(text, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Failed to hash stdin: " + err.Error())
		os.Exit(1)
	}

	fmt.Println(string(hashed))
}
