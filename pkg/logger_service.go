package logger

import (
    "sync"
)

type service struct {
    handler Handler
    mutex   sync.RWMutex // Добавлен мьютекс для потокобезопасности
}

var (
    serviceInstance *service
    serviceOnce     sync.Once
)

// getService возвращает экземпляр сервиса, инициализируя его при необходимости
func getService() *service {
    serviceOnce.Do(func() {
        serviceInstance = &service{
            handler: nil, // Изначально без обработчика
        }
    })
    return serviceInstance
}

// SetHandler устанавливает обработчик логгера
func SetHandler(h Handler) {
    s := getService()
    s.mutex.Lock()
    defer s.mutex.Unlock()
    s.handler = h
}

// getHandler возвращает текущий обработчик
func getHandler() Handler {
    s := getService()
    s.mutex.RLock()
    defer s.mutex.RUnlock()
    return s.handler
}

// Terminate корректно завершает работу логгера
func Terminate() error {
    handler := getHandler()
    if handler != nil {
        return handler.Terminate()
    }
    return nil
}

// Info отправляет информационное сообщение
func Info(msg string) {
    if handler := getHandler(); handler != nil {
        handler.Info(msg)
    }
}

// Debug отправляет отладочное сообщение
func Debug(msg string) {
    if handler := getHandler(); handler != nil {
        handler.Debug(msg)
    }
}

// Warning отправляет предупреждение
func Warning(msg string) {
    if handler := getHandler(); handler != nil {
        handler.Warning(msg)
    }
}

// Error отправляет сообщение об ошибке
func Error(msg string) {
    if handler := getHandler(); handler != nil {
        handler.Error(msg)
    }
}
