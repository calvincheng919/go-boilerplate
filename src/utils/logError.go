package utils

import (
	"log"
	"os"
)

func LogError(msg string) {
	file, err := os.OpenFile("errors.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer func() { _ = file.Close() }()

	log.SetOutput(file)
	log.Println(msg)
}
