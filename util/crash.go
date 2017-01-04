package util

import "github.com/Sirupsen/logrus"

func Crash(message string, err error) {
	if err != nil {
		logrus.Fatal(message, "\n\n", err)
	}
}
