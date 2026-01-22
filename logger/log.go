package logger

import "fmt"

type Logger struct {
	Prefix string
}

func (l Logger) Log(msg string) {
	fmt.Printf("%s: %s\n", l.Prefix, msg)
}
type TimestampedLogger struct{
	Logger
	Timestamp int
}

func (t TimestampedLogger) Log(msg string) {
	fmt.Printf("%d %s: %s\n", t.Timestamp, t.Prefix, msg)
}
