package logger

import (
	"log"
	"os"
)

type Logger struct {
	logger *log.Logger
}

var PublicStdoutLogger = NewLogger(log.New(os.Stdout, "", log.LstdFlags))

func NewLogger(logger *log.Logger) *Logger {
	return &Logger{logger: log.New(os.Stdout, "", log.LstdFlags)}
}

func (l *Logger) Log(message string) {
	l.logger.Println(message)
}

func (l *Logger) Info(message string) {
	l.logger.Println("INFO:", message)
}

func (l *Logger) Success(message string) {
	l.logger.Println("SUCCESS:", message)
}

func (l *Logger) Warn(message string) {
	l.logger.Println("WARN:", message)
}

func (l *Logger) Error(message string) {
	l.logger.Println("ERROR:", message)
}

func (l *Logger) Assert(condition bool, message string) {
	if !condition {
		l.logger.Println("ASSERT:", message)
	}
}

func (l *Logger) Fatal(message string) {
	l.logger.Fatalln("FATAL:", message)
}

func (l *Logger) Panic(message string) {
	l.logger.Panicln("PANIC:", message)
}

func (l *Logger) PanicIfError(err error) {
	if err != nil {
		l.logger.Panicln("PANIC:", err)
	}
}

func (l *Logger) FatalIfError(err error) {
	if err != nil {
		l.logger.Fatalln("FATAL:", err)
	}
}

func (l *Logger) LogIfError(err error) {
	if err != nil {
		l.logger.Println("ERROR:", err)
	}
}
