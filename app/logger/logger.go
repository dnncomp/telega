package logger

import (
	"fmt"
	"github.com/hashicorp/logutils"
	"github.com/natefinch/lumberjack"
	"io"
	"log"
	"os"
	"path/filepath"
)

func New(logPath, logLevel string, logMaxSize, logMaxBackups, logMaxAge int) *log.Logger {

	//каталог и файл логов
	logDir := filepath.Dir(logPath + "/")
	//logFileName := filepath.Base(logPath)
	logFileName := "telega.log"
	logFilePath := filepath.Join(logDir, logFileName)

	// создадим и проверим путь к лог файлу
	if err := os.MkdirAll(logDir, os.ModeDir); err != nil {
		log.Fatalln("Ошибка создания каталога для логов", logDir, err)
	}
	_, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Ошибка открытия файла логов", logFilePath, err)
	}

	// карусель логов
	var lumberjackLogger *lumberjack.Logger
	lumberjackLogger = &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    logMaxSize, // megabytes
		MaxBackups: logMaxBackups,
		MaxAge:     logMaxAge, //days
		LocalTime:  true,
	}
	multi := io.MultiWriter(lumberjackLogger, os.Stderr)

	flags := log.LstdFlags
	if logLevel != "INFO" {
		flags = log.LstdFlags | log.Lshortfile
	}
	logger := log.New(os.Stdout, "", flags)

	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(logLevel),
		Writer:   multi,
	}
	logger.SetOutput(filter)

	if logLevel == "DEBUG" {
		fmt.Println("Каталог логирования:", logFilePath)
	}

	return logger
}
