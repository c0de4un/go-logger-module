package logger

// Handler интерфейс для различных реализаций логгера
type Handler interface {
    Info(msg string)
    Debug(msg string)
    Warning(msg string)
    Error(msg string)
    Terminate() error // Добавлен метод для корректного завершения
}
