package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	LoggerInLine = GetLogger("LOGGER...")
	LoggerInFile = GetLoggerInFile()
)

// func UnifyLogs(p string) (*Logger, *Logger) {
// 	logger := GetLogger(p)
// 	loggerInFile := GetLoggerInFile()
// 	return logger, loggerInFile
// }

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger := NewLogger(p)
	return logger
}

func GetLoggerInFile() *Logger {
	arquivoLog, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Erro ao abrir o arquivo de log:", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Errorf("Erro ao fechar arquivo de logs: %v", err)
		}
	}(arquivoLog)

	// Cria uma nova instância do logger com arquivo de saída
	logger, err := NewLoggerWithFile("LOG: ", "logs.txt")
	if err != nil {
		log.Fatal("Erro ao iniciar logger:", arquivoLog)
	}

	logger.Info("Iniciando Arquivo...")

	return logger
}
