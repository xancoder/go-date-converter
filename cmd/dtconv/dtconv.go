package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// source: https://www.rfc-editor.org/rfc/rfc2616#section-3.3
	months := map[string]int{
		"Jan": 1,
		"Feb": 2,
		"Mar": 3,
		"Apr": 4,
		"May": 5,
		"Jun": 6,
		"Jul": 7,
		"Aug": 8,
		"Sep": 9,
		"Oct": 10,
		"Nov": 11,
		"Dec": 12,
	}

	//result from "$(wget -qSO- google.com 2>&1 | sed -n 's/^ *Date: *//p' | head -n 1)"
	//tmp := "Mon, 22 Apr 2024 15:10:40 GMT"
	tmp := os.Args[1]
	parts := strings.Split(tmp, " ")
	partsTime := strings.Split(parts[4], ":")

	yyyy := fmt.Sprintf("%04d", convertStringToInt(parts[3]))
	m := fmt.Sprintf("%02d", months[parts[2]])
	dd := fmt.Sprintf("%02d", convertStringToInt(parts[1]))
	hh := fmt.Sprintf("%02d", convertStringToInt(partsTime[0]))
	mm := fmt.Sprintf("%02d", convertStringToInt(partsTime[1]))
	s := fmt.Sprintf("%02d", convertStringToInt(partsTime[2]))

	fmt.Printf("%s%s%s%s%s.%s", yyyy, m, dd, hh, mm, s)
}

func convertStringToInt(t string) int {
	r, err := strconv.Atoi(t)
	if err != nil {
		panic(err)
	}
	return r
}
