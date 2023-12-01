package main

import (
	"fmt"
	"log"

	l "github.com/Dev-Anderson/NewLogger/pkg/log"
)

func Initialization() {
	exist, err := l.VerifyFileLog()
	if err != nil {
		log.Panic("Erro ao verificar arquivo de log", err)
	}
	if !exist {
		l.CreateFileLog()
	}
}

func main() {
	fmt.Println("Iniciando o projeto")
}
