package log

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
}

func NewLogger() Logger {
	return &logger{}
}

type logger struct {
}

func (l *logger) Debug(args ...interface{}) {}
func (l *logger) Info(args ...interface{})  {}
func (l *logger) Error(args ...interface{}) {}
func (l *logger) Warn(args ...interface{})  {}
