package logger

import (
    "testing"
)

// MockHandler для тестирования
type MockHandler struct {
    messages []string
}

func (m *MockHandler) Info(msg string) {
    m.messages = append(m.messages, "INFO: "+msg)
}

func (m *MockHandler) Debug(msg string) {
    m.messages = append(m.messages, "DEBUG: "+msg)
}

func (m *MockHandler) Warning(msg string) {
    m.messages = append(m.messages, "WARNING: "+msg)
}

func (m *MockHandler) Error(msg string) {
    m.messages = append(m.messages, "ERROR: "+msg)
}

func (m *MockHandler) Terminate() error {
    return nil
}

func TestLogger(t *testing.T) {
    mockHandler := &MockHandler{}
    SetHandler(mockHandler)
    
    Info("Test info message")
    Debug("Test debug message")
    Warning("Test warning message")
    Error("Test error message")
    
    expectedLen := 4
    if len(mockHandler.messages) != expectedLen {
        t.Errorf("Expected %d messages, got %d", expectedLen, len(mockHandler.messages))
    }
    
    expectedFirstMessage := "INFO: Test info message"
    if mockHandler.messages[0] != expectedFirstMessage {
        t.Errorf("Expected first message '%s', got '%s'", expectedFirstMessage, mockHandler.messages[0])
    }
}