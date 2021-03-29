package logs

import (
	"fmt"
	"os"

	"github.com/edgex-go-api/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logLevel zapcore.Level

func InitLogs() {
	logLevel = zapcore.DebugLevel
	if config.AppConfig.LogLevel == "Info" {
		logLevel = zapcore.InfoLevel
	}
	setLogConfig()
}

func setLogConfig() {
	// 设置日志输出格式
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "name",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	})

	// 添加日志切割归档功能
	hook := lumberjack.Logger{
		Filename:   config.AppConfig.FileName,   // 日志文件路径
		MaxSize:    config.AppConfig.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: config.AppConfig.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     config.AppConfig.MaxAge,     // 文件最多保存多少天
		Compress:   config.AppConfig.Compress,   // 是否压缩
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), zapcore.AddSync(&hook)),
		zap.NewAtomicLevelAt(logLevel),
	)

	caller := zap.AddCaller()                    // 开启文件及行号
	development := zap.Development()             // 开启开发模式，堆栈跟踪
	logger := zap.New(core, caller, development) // 构造日志
	zap.ReplaceGlobals(logger)                   // 将自定义的logger替换为全局的logger
}

func ZapInfo(msg string, fields ...zapcore.Field) {
	zap.L().Info(msg, fields...)
}

func ZapWarn(msg string, fields ...zapcore.Field) {
	zap.L().Warn(msg, fields...)
}

func ZapDebug(msg string, fields ...zapcore.Field) {
	zap.L().Debug(msg, fields...)
}

func ZapError(msg string, fields ...zapcore.Field) {
	zap.L().Error(msg, fields...)
}

func ZapDPanic(msg string, fields ...zapcore.Field) {
	zap.L().DPanic(msg, fields...)
}

func ZapPanic(msg string, fields ...zapcore.Field) {
	zap.L().Panic(msg, fields...)
}

func ZapFatal(msg string, fields ...zapcore.Field) {
	zap.L().Fatal(msg, fields...)
}

func Info(format string, v ...interface{}) {
	if zapcore.InfoLevel < logLevel {
		return
	}
	if len(v) == 0 {
		zap.L().Info(format)
		return
	}
	zap.L().Info(fmt.Sprintf(format, v...))
}

func Warn(format string, v ...interface{}) {
	if zapcore.WarnLevel < logLevel {
		return
	}
	if len(v) == 0 {
		zap.L().Warn(format)
		return
	}
	zap.L().Warn(fmt.Sprintf(format, v...))
}

func Debug(format string, v ...interface{}) {
	if zapcore.DebugLevel < logLevel {
		return
	}
	if len(v) == 0 {
		zap.L().Debug(format)
		return
	}
	zap.L().Debug(fmt.Sprintf(format, v...))
}

func Error(format string, v ...interface{}) {
	if zapcore.ErrorLevel < logLevel {
		return
	}
	if len(v) == 0 {
		zap.L().Error(format)
		return
	}
	zap.L().Error(fmt.Sprintf(format, v...))
}

func Fatal(format string, v ...interface{}) {
	if zapcore.FatalLevel < logLevel {
		return
	}
	if len(v) == 0 {
		zap.L().Fatal(format)
		return
	}
	zap.L().Fatal(fmt.Sprintf(format, v...))
}
