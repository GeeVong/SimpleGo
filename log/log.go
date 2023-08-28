package log

import (
	"fmt"
	"github.com/anqiansong/ketty/console"
)

/*
 	理论上日志输出只需要一下几个信息：
		1. time
		2. 日志打印位置
		3. 日志级别
		4. 日志msg

问题：
		1. 多个日志str 输出问题 done
		2. 日志颜色显示问题	done
		3. 日志格式设置问题	done
*/

func LogInfo(tag, str string, param ...any) {
	switch LogPackage {

	case LogPackage_zerolog:
		msg := fmt.Sprintf(str, param...)
		Logger.Info().Caller(1).Str("[tag]", tag).Msg(msg)

	case LogPackage_kettyconsole:
		console.Info(str)
	}
}

func LogError(tag, str string, param ...any) {
	switch LogPackage {
	case LogPackage_zerolog:
		msg := fmt.Sprintf(str, param...)
		Logger.Error().Caller(1).Str("[tag]", tag).Msg(msg)

	case LogPackage_kettyconsole:
		console.Info(str)

	}
}

func LogWarn(tag, str string, param ...any) {
	switch LogPackage {
	case LogPackage_kettyconsole:
		console.Info(str)
	case LogPackage_zerolog:

		msg := fmt.Sprintf(str, param...)
		Logger.Warn().Caller(1).Str("[tag]", tag).Msg(msg)
	}
}

func LogDebug(tag, str string, param ...any) {
	switch LogPackage {
	case LogPackage_kettyconsole:
		console.Info(str)
	case LogPackage_zerolog:

		msg := fmt.Sprintf(str, param...)
		Logger.Debug().Caller(1).Str("[tag]", tag).Msg(msg)
	}
}
