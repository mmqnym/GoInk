package logger

import (
	"os"

	"goink/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger
var Sugar *zap.SugaredLogger

func init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  "./logs/" + config.Log.FileName,
		MaxSize:   config.Log.MaxSize,
		MaxAge:    config.Log.MaxAge,
		LocalTime: config.Log.LocalTime,
		Compress:  config.Log.Compress,
	})

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	level := judgeLevel()

	fileCore := zapcore.NewCore(
		appEncoder{encoder},
		w,
		level,
	)

	consoleCore := zapcore.NewCore(
		appEncoder{encoder},
		zapcore.AddSync(os.Stdout),
		level,
	)

	core := zapcore.NewTee(fileCore, consoleCore)

	Logger = zap.New(core)
	Sugar = Logger.Sugar()
}

func judgeLevel() zapcore.Level {
	switch config.Log.Level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}
