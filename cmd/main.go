package main

import (
	newlogger "github.com/Dev-Anderson/NewLogger"
)

func main() {
	msg := newlogger.LogInfo("Jorge e mateus", "Nao deixe apagar a fogueira do meu coracao")
	newlogger.GenerateNewLog(msg)

}
