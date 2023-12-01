package main

import (
	"fmt"
	"log"
	"os"
)

func namePathLog() string {
	path, err := os.Getwd()
	if err != nil {
		log.Panic("Falha ao pegar o nome da pasta raiz", err)
	}
	return path + "/logs"
}

func nameFileLog() string {
	return fmt.Sprintf("%s.log", currentDate())
}

func nameAllFile() string {
	return namePathLog() + "/" + nameFileLog()
}
