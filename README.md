# go-logger-module

Простой пакет логгера с поддержкой различных обработчиков. Совместим с другими логгерами.

## Установка

```bash
go get github.com/c0de4un/go-logger-module
```

## Использование

```go
package main

import (
    "github.com/c0de4un/go-logger-module.git/pkg"
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
```

## Особенности

- Потокобезопасный синглтон сервис
- Поддержка различных обработчиков логов
- Безопасная запись в файлы с предотвращением path traversal
- Поддержка завершения работы с освобождением ресурсов
- Модульные тесты

## Интерфейс Handler

Для создания собственного обработчика реализуйте интерфейс `Handler`:

```go
type Handler interface {
    Info(msg string)
    Debug(msg string)
    Warning(msg string)
    Error(msg string)
    Terminate() error
}
```
