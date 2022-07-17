package util

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SugarLogger *zap.SugaredLogger

func InitLogger() {
	// 进行日志分级
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //error  以下分均为lowPriority
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //error 及以上分为highPriority
		return lev >= zap.ErrorLevel
	})
	LowLogWriter := getLowLogWriter()
	HighLogWriter := getHighLogWriter()
	encoder := GetEncoder()
	LowCore := zapcore.NewCore(encoder, LowLogWriter, lowPriority)
	HighCore := zapcore.NewCore(encoder, HighLogWriter, highPriority)
	logger := zap.New(zapcore.NewTee(HighCore, LowCore), zap.AddCaller())
	SugarLogger = logger.Sugar()

}

func GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

//使用Lumberjack进行日志切割归档
func getLowLogWriter() zapcore.WriteSyncer {
	LowJackLogger := &lumberjack.Logger{
		Filename:   "./log/info.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}

	return zapcore.AddSync(LowJackLogger)
}

func getHighLogWriter() zapcore.WriteSyncer {
	HighJackLogger := &lumberjack.Logger{
		Filename:   "./log/error.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(HighJackLogger)
}
