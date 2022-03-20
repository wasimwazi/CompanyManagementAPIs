package cmd

import (
	"XM/common/utils"
	"errors"
	"os"
)

func getPort() (string, error) {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return utils.EmptyString, errors.New("PORT environment variable missing")
	}
	return port, nil
}

func checkEnv() error {
	_, ok := os.LookupEnv("DBConString")
	if !ok {
		return errors.New("DBConString environment variable missing")
	}
	_, ok = os.LookupEnv("PORT")
	if !ok {
		return errors.New("PORT environment variable missing")
	}
	return nil
}
