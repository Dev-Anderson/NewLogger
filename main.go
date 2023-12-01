package main

import (
	"fmt"
	"log"
)

func Initialization() {
	exist, err := verifyFileLog()
	if err != nil {
		log.Panic("Erro ao verificar arquivo de log", err)
	}
	if !exist {
		createFileLog()
	}
}

func main() {
	fmt.Println("Iniciando o projeto")
}
