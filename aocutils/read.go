package aocutils

import (
	"log"
	"os"
)

func GetRawInput() string {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal("problem reading input")
	}

	return string(data)
}
