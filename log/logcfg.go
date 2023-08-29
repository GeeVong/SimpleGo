package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"sync"
	"time"
)

const (
	LogConsole        = 0
	LogConsoleColor   = 1
	LogFile           = 2
	LogConsoleAndFile = 3
)

var LogFlag int
var LogLevel = zerolog.DebugLevel // 日志级别选项
var Logger zerolog.Logger

var once sync.Once

// 初始化 zerolog
func InitZeroLogCfg() {
	once.Do(func() {
		LogFlag = LogConsole
		switch LogFlag {
		case LogConsole:
			Logger = zeroLogConsole()
		case LogConsoleColor:
			Logger = zeroWithColor()
		case LogFile:
			Logger = zeroLogWriteFile()
		case LogConsoleAndFile:
			Logger = zeroLogConsoleAndFile()
		}
	})
}

func zeroLogConsoleAndFile() zerolog.Logger {
	// 创建log目录
	logDir := "./run_log/"
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		fmt.Println("Mkdir failed, err:", err)
		return zerolog.Logger{} // todo
	}

	// 文件目录格式
	fileName := logDir + time.Now().Format("2006-01-02") + ".log"
	logFile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	lw := zerolog.MultiLevelWriter(logFile)
	// 日志文件，数据时间格式
	zerolog.TimeFieldFormat = time.RFC3339

	// 定义标签
	zerolog.TimestampFieldName = "[TS]"
	zerolog.LevelFieldName = "[Level]"
	zerolog.CallerFieldName = "[Caller]"
	zerolog.MessageFieldName = "[Msg]"
	zerolog.SetGlobalLevel(LogLevel)

	cw := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}

	multi := zerolog.MultiLevelWriter(cw, lw)
	return zerolog.New(multi).
		With().
		Timestamp().
		Logger().
		Level(LogLevel)
}

func zeroLogWriteFile() zerolog.Logger {
	// 创建log目录
	logDir := "./run_log/"
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		fmt.Println("Mkdir failed, err:", err)
		return zerolog.Logger{} // todo
	}

	// 文件目录格式
	fileName := logDir + time.Now().Format("2006-01-02") + ".log"
	logFile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	lw := zerolog.MultiLevelWriter(logFile)
	// 日志文件，数据时间格式
	zerolog.TimeFieldFormat = time.RFC3339

	// 定义标签
	zerolog.TimestampFieldName = "[TS]"
	zerolog.LevelFieldName = "[Level]"
	zerolog.MessageFieldName = "[Msg]"
	zerolog.CallerFieldName = "[Caller]"
	zerolog.SetGlobalLevel(LogLevel)

	return zerolog.New(lw).
		With().
		Logger().
		Level(LogLevel)

}

func zeroWithColor() zerolog.Logger {
	cw := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	return log.Output(cw).
		With().
		Logger().
		Level(LogLevel)

}

func zeroLogConsole() zerolog.Logger {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	zerolog.SetGlobalLevel(LogLevel)
	return zerolog.New(consoleWriter).
		With().
		Timestamp().
		Logger()

}

// todo hook
