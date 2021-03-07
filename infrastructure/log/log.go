package log

type Logger interface {
	Log(keyvals ...interface{}) error // same as `Info`
	Debug(keyvals ...interface{}) error
	Info(keyvals ...interface{}) error
	Warn(keyvals ...interface{}) error
	Error(keyvals ...interface{}) error
	Fatal(keyvals ...interface{}) error
}
