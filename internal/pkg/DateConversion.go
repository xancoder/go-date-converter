package pkg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ConvertDateTime(d string) (string, error) {
	if d == "" {
		return "", errors.New("empty datetime")
	}

	// source: https://www.rfc-editor.org/rfc/rfc2822.html#section-3.3
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

	split := strings.Split(d, " ")
	yyyy := fmt.Sprintf("%04d", convertStringToInt(split[3]))
	m := fmt.Sprintf("%02d", months[split[2]])
	dd := fmt.Sprintf("%02d", convertStringToInt(split[1]))

	time := strings.Split(split[4], ":")
	hh := fmt.Sprintf("%02d", convertStringToInt(time[0]))
	mm := fmt.Sprintf("%02d", convertStringToInt(time[1]))
	s := fmt.Sprintf("%02d", convertStringToInt(time[2]))

	r := fmt.Sprintf("%s%s%s%s%s.%s", yyyy, m, dd, hh, mm, s)

	return r, nil
}

func convertStringToInt(t string) int {
	r, err := strconv.Atoi(t)
	if err != nil {
		panic(err)
	}
	return r
}
