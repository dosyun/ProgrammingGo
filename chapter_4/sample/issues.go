package main

import (
	"fmt"
	"log"
	"os"
	"work/go/ProgrammingGo/ch4/github/ch4/github"
)

func main() {
	result, err := github.SeachIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
}
