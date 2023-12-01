package main

import (
	"fmt"
	"os"
)

func createFileLog() error {
	if err := os.MkdirAll(namePathLog(), os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(nameAllFile())
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func verifyFileLog() (bool, error) {
	_, err := os.Stat(nameAllFile())
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func addLinesFile(msg string) error {
	file, err := os.OpenFile(nameAllFile(), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, msg)
	if err != nil {
		return err
	}
	return nil
}
