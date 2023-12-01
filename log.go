package main

import (
	"fmt"
	"log"
)

func formatMsg(cod, msg string) string {
	log := fmt.Sprintf(" %s %s", cod, msg)
	message := timeNow() + log
	return message
}

func LogInfo(cod, msg string) error {
	err := addLinesFile(formatMsg(cod, msg))
	if err != nil {
		log.Panic("Erro ao adicionar log", err)
	}
	return nil
}
