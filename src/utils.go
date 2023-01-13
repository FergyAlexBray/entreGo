package entrego

import (
	"log"
	"strconv"
)

func convertStringToNumber(s string) int {
	var data int
	var err error

	data, err = strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}
	return data
}
