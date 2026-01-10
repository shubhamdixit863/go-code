package log

import "fmt"

type Logger struct {
	Prefix string
}

func (l Logger) Log(message string) {
	fmt.Printf("%s:%s", l.Prefix, message)
}

type TimestampedLogger struct {
	Logger    // if you are mebedding another struct its fields are accesible directly
	Timestamp int
}

func (l TimestampedLogger) Log(message string) {
	fmt.Printf("%d:%s:%s", l.Timestamp, l.Prefix, message)
}

type Service struct {
	Logger Logger // composition
	Name   string
}

func (s Service) Start(message string) {
	s.Logger.Log(message)
}
