package aocutils

import (
	"log"
	"strconv"
)

func Stoi(s string) int {
	n, err := strconv.Atoi(s)

	if err != nil {
		log.Fatalf("problem converting %s to int", s)
	}

	return n
}
