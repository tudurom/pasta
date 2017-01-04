package util

import (
	"errors"
	"os"

	"github.com/Sirupsen/logrus"
)

type Storage struct {
	Path string
}

var GlobalStorage Storage

func (s *Storage) NewStorage() {
	logrus.Print("Creating storage")
	info, err := os.Stat(GlobalConfig.StoragePath)
	if os.IsNotExist(err) {
		logrus.Print("Storage directory does not exist. Creating one")
		err = os.Mkdir(GlobalConfig.StoragePath, 0755)
		if err != nil {
			Crash("Couldn't create storage directory", err)
		}
	} else if err != nil {
		logrus.Print("Storage directory available")
		if !info.IsDir() {
			logrus.Fatal(errors.New("The specified storage path is not a directory"))
		}
	} else {
		Crash("Error while creating storage", err)
	}

	// All good
	s.Path = GlobalConfig.StoragePath
}
