package main

import (
	logger "github.com/c0de4un/go-logger-module.git/pkg"
)

func main() {
	handler := &logger.DefaultHandler{LogsDir: "./logs"}
	logger.SetHandler(handler)

	logger.Info("Пример информационного сообщения")
	logger.Debug("Пример отладочного сообщения")
	logger.Warning("Пример предупреждения")
	logger.Error("Пример сообщения об ошибке")

	logger.Terminate()
}
