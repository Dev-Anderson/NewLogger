package file

import (
	"fmt"
	"log"
	"os"

	"github.com/Dev-Anderson/NewLogger/internal/util"
)

func namePath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Panic("Falha ao pegar o nome da pasta raiz", err)
	}
	return path + "/logs"
}

func nameFile() string {
	return fmt.Sprintf("%s.log", util.CurrentDate)
}

func nameAllFile() string {
	return namePath() + "/" + nameFile()
}

func CreateFileLog() error {
	if err := os.MkdirAll(namePath(), os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(nameAllFile())
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func VerifyFileLog() (bool, error) {
	_, err := os.Stat(nameAllFile())
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func AddLinesFile(msg string) error {
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
