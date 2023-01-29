package logger

import "fmt"

// Logger ...
type Logger struct {
	logLevel string
}

// Log ...
func (l *Logger) Log(s string) {
	fmt.Println(s)
}
