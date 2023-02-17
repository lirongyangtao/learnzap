package core

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel

	_minLevel = DebugLevel
	_maxLevel = FatalLevel

	InvalidLevel = _maxLevel + 1
)
