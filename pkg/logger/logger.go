package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	s *zap.SugaredLogger
}

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func NewLogger(logLevel string, logFile string) (*Logger, error) {
	level, err := getLogLever(logLevel)
	if err != nil {
		return nil, err
	}

	writeSyncer := getLogWriter(logFile)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zap.NewAtomicLevelAt(level))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
	return &Logger{logger.Sugar()}, nil
}

func getLogLever(logLevel string) (zapcore.Level, error) {
	if level, ok := levelMap[logLevel]; ok {
		return level, nil
	}

	return 0, fmt.Errorf("unknown log level %s", logLevel)
}

func getLogWriter(logFile string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile, // 日志文件的位置
		MaxSize:    1024,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 30,      // 保留旧文件的最大个数
		MaxAge:     90,      // 保留旧文件的最大天数
		Compress:   true,    // 是否压缩/归档旧文件
	}

	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func (l *Logger) Debug(args ...interface{}) {
	l.s.Debug(args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.s.Debugf(format, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.s.Info(args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.s.Infof(format, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.s.Warn(args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.s.Warnf(format, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.s.Error(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.s.Errorf(format, args...)
}

func (l *Logger) DPanic(args ...interface{}) {
	l.s.DPanic(args...)
}

func (l *Logger) DPanicf(format string, args ...interface{}) {
	l.s.DPanicf(format, args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.s.Panic(args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.s.Panicf(format, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.s.Fatal(args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.s.Fatalf(format, args...)
}
