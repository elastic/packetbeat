package logp

import (
	"go.uber.org/zap"
)

// LogOption configures a Logger.
type LogOption = zap.Option

// Logger logs messages to the configured output.
type Logger struct {
	sugar *zap.SugaredLogger
}

// NewLogger returns a new Logger labeled with the name of the selector. This
// should never be used from any global contexts (instead create "per instance"
// loggers).
func NewLogger(selector string, options ...LogOption) *Logger {
	log := loadLogger().rootLogger.
		WithOptions(zap.AddCallerSkip(1)).
		WithOptions(options...).
		Named(selector)
	return &Logger{log.Sugar()}
}

// With creates a child logger and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
func (l *Logger) With(args ...interface{}) *Logger {
	return &Logger{l.sugar.With(args...)}
}

// Sprint

// Debug uses fmt.Sprint to construct and log a message.
func (l *Logger) Debug(args ...interface{}) {
	l.sugar.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func (l *Logger) Info(args ...interface{}) {
	l.sugar.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *Logger) Warn(args ...interface{}) {
	l.sugar.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *Logger) Error(args ...interface{}) {
	l.sugar.Error(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit(1).
func (l *Logger) Fatal(args ...interface{}) {
	l.sugar.Fatal(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *Logger) Panic(args ...interface{}) {
	l.sugar.Panic(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics.
func (l *Logger) DPanic(args ...interface{}) {
	l.sugar.DPanic(args...)
}

// Sprintf

// Debugf uses fmt.Sprintf to construct and log a message.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.sugar.Debugf(format, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.sugar.Infof(format, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.sugar.Warnf(format, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.sugar.Errorf(format, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit(1).
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.sugar.Fatalf(format, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.sugar.Panicf(format, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics.
func (l *Logger) DPanicf(format string, args ...interface{}) {
	l.sugar.DPanicf(format, args...)
}

// With context (reflection based)

// Debugw logs a message with some additional context. The additional context
// is added in the form of key-value pairs. The optimal way to write the value
// to the log message will be inferred by the value's type. To explicitly
// specify a type you can pass a Field such as logp.Stringer.
func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.sugar.Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The additional context
// is added in the form of key-value pairs. The optimal way to write the value
// to the log message will be inferred by the value's type. To explicitly
// specify a type you can pass a Field such as logp.Stringer.
func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.sugar.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The additional context
// is added in the form of key-value pairs. The optimal way to write the value
// to the log message will be inferred by the value's type. To explicitly
// specify a type you can pass a Field such as logp.Stringer.
func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.sugar.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The additional context
// is added in the form of key-value pairs. The optimal way to write the value
// to the log message will be inferred by the value's type. To explicitly
// specify a type you can pass a Field such as logp.Stringer.
func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.sugar.Errorw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit(1).
// The additional context is added in the form of key-value pairs. The optimal
// way to write the value to the log message will be inferred by the value's
// type. To explicitly specify a type you can pass a Field such as
// logp.Stringer.
func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.sugar.Fatalw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// additional context is added in the form of key-value pairs. The optimal way
// to write the value to the log message will be inferred by the value's type.
// To explicitly specify a type you can pass a Field such as logp.Stringer.
func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.sugar.Panicw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. The logger panics only
// in Development mode.  The additional context is added in the form of
// key-value pairs. The optimal way to write the value to the log message will
// be inferred by the value's type. To explicitly specify a type you can pass a
// Field such as logp.Stringer.
func (l *Logger) DPanicw(msg string, keysAndValues ...interface{}) {
	l.sugar.DPanicw(msg, keysAndValues...)
}
