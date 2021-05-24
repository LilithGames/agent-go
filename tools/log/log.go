package log

import (
	"bytes"
	"go.uber.org/zap"
	"runtime"
)

var zapLogger *zap.Logger

func init() {
	zapLogger, _ = zap.NewProduction(zap.AddCaller(), zap.AddCallerSkip(1))
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(msg string, fields ...zap.Field) {
	field := zap.String("goroutine", getGoroutineID())
	fields = append(fields, field)
	zapLogger.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(msg string, fields ...zap.Field) {
	field := zap.String("goroutine", getGoroutineID())
	fields = append(fields, field)
	zapLogger.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(msg string, fields ...zap.Field) {
	field := zap.String("goroutine", getGoroutineID())
	fields = append(fields, field)
	zapLogger.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(msg string, fields ...zap.Field) {
	field := zap.String("goroutine", getGoroutineID())
	fields = append(fields, field)
	zapLogger.Error(msg, fields...)
}

// DPanic logs a message at DPanicLevel. The message includes any fields
// passed at the log site, as well as any fields accumulated on the logger.
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func DPanic(msg string, fields ...zap.Field) {
	field := zap.String("goroutine", getGoroutineID())
	fields = append(fields, field)
	zapLogger.DPanic(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func Panic(msg string, fields ...zap.Field) {
	field := zap.String("goroutine", getGoroutineID())
	fields = append(fields, field)
	zapLogger.Panic(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func Fatal(msg string, fields ...zap.Field) {
	field := zap.String("goroutine", getGoroutineID())
	fields = append(fields, field)
	zapLogger.Fatal(msg, fields...)
}

func getGoroutineID() string {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	return string(b)
}
