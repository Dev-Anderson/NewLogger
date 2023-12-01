package log

import (
	"fmt"
	"log"

	"github.com/Dev-Anderson/NewLogger/internal/file"
	"github.com/Dev-Anderson/NewLogger/internal/util"
)

func formatMsg(cod, msg string) string {
	log := fmt.Sprintf(" %s %s", cod, msg)
	message := util.TimeNow() + log
	return message
}

func LogInfo(cod, msg string) error {
	err := file.AddLinesFile(formatMsg(cod, msg))
	if err != nil {
		log.Panic("Erro ao adicionar log", err)
	}
	return nil
}
