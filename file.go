package newlogger

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Path() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}

func PathToFile() string {
	toFile := fmt.Sprintf("%s/%s", Path(), "logs")
	return toFile
}

func NameFile() string {
	return fmt.Sprintf("%s.log", CurrentDate())
}

func NameFileWithPath() string {
	return fmt.Sprintf(PathToFile() + "/" + NameFile())
}

func DeleteFileInForDay(day int64) {
	t := time.Now()
	folder, err := os.Open(PathToFile())
	if err != nil {
		log.Println("Error when searching for file", err)
	}
	defer folder.Close()
	files, err := folder.Readdir(-1)
	if err != nil {
		log.Println("Error reading the directory", err)
	}

	for _, file := range files {
		if t.Sub(file.ModTime()) > (time.Hour * 24 * time.Duration(day)) {
			os.Remove(PathToFile() + "/" + file.Name())
		}
	}
}

func GenerateNewLog(logMessage string) {
	os.Mkdir(fmt.Sprintf("%s/%s", Path(), "logs"), 0755)
	_, err := os.Stat(NameFileWithPath())
	if err != nil {
		file, err := os.Create(NameFileWithPath())
		if err != nil {
			log.Println("Erro in create file", err)
		}
		defer file.Close()
	}
	logFile, err := os.OpenFile(NameFileWithPath(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Erro open file", err)
	}

	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println(logMessage)
	defer logFile.Close()
}
