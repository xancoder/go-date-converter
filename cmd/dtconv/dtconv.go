package main

import (
	"fmt"
	"os"

	"go-date-converter/internal/pkg"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("one argument is required")
		fmt.Println("usage: dtconv <datetime>")
		os.Exit(1)
	}

	s := os.Args[1]
	dateTime, err := pkg.ConvertDateTime(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(dateTime)
}
