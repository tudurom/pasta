package util

import "log"

func Crash(message string, err error) {
	if err != nil {
		log.Fatal(message, "\n\n", err)
	}
}
