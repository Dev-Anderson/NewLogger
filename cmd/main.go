package main

import (
	"fmt"
	"log"

	"github.com/Dev-Anderson/NewLogger/internal/file"
)

func Initialization() {
	exist, err := file.VerifyFileLog()
	if err != nil {
		log.Panic("Erro ao verificar arquivo de log", err)
	}
	if !exist {
		file.CreateFileLog()
	}
}

func main() {
	fmt.Println("Iniciando o projeto")
}
