package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// DefaultHandler реализует интерфейс Handler с записью в файл и консоль
type DefaultHandler struct {
	LogsDir string
}

// Info записывает информационное сообщение
func (dl *DefaultHandler) Info(msg string) {
	dl.printMsg("INFO", msg)
}

// Debug записывает отладочное сообщение
func (dl *DefaultHandler) Debug(msg string) {
	dl.printMsg("DEBUG", msg)
}

// Warning записывает предупреждение
func (dl *DefaultHandler) Warning(msg string) {
	dl.printMsg("WARNING", msg)
}

// Error записывает сообщение об ошибке
func (dl *DefaultHandler) Error(msg string) {
	dl.printMsg("ERROR", msg)
}

// Terminate завершает работу обработчика
func (dl *DefaultHandler) Terminate() error {
	// В этой реализации нет ресурсов, требующих освобождения
	return nil
}

// printMsg форматирует и записывает сообщение в файл и консоль
func (dl *DefaultHandler) printMsg(tag string, msg string) {
	currentTime := time.Now()
	strToWrite := fmt.Sprintf("[%s] %s: %s\n",
		currentTime.Format("02-01-2006-15:04:05"), tag, msg)

	fmt.Print(strToWrite)

	// Безопасное формирование пути к файлу
	logFileName := fmt.Sprintf("%s_log.txt", currentTime.Format("02_01_2006"))
	filePath := filepath.Join(dl.safePath(dl.LogsDir), logFileName)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			fmt.Printf("Error closing log file: %v\n", closeErr)
		}
	}()

	_, err = file.WriteString(strToWrite)
	if err != nil {
		fmt.Printf("Error writing to log file: %v\n", err)
		return
	}
}

// safePath предотвращает path traversal уязвимости
func (dl *DefaultHandler) safePath(dir string) string {
	if dir == "" {
		dir = "." // Используем текущую директорию по умолчанию
	}

	// Очищаем путь от потенциально опасных элементов
	cleanPath := filepath.Clean(dir)
	if strings.HasPrefix(cleanPath, "../") ||
		strings.Contains(cleanPath, "/../") ||
		cleanPath == ".." {
		return "."
	}

	return cleanPath
}
