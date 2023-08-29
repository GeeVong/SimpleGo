package log

import (
	"fmt"
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
	msg := fmt.Sprintf(str, param...)
	Logger.Info().Caller(1).Str("[tag]", tag).Msg(msg)

}

func LogError(tag, str string, param ...any) {
	msg := fmt.Sprintf(str, param...)
	Logger.Error().Caller(1).Str("[tag]", tag).Msg(msg)
}

func LogWarn(tag, str string, param ...any) {
	msg := fmt.Sprintf(str, param...)
	Logger.Warn().Caller(1).Str("[tag]", tag).Msg(msg)
}

func LogDebug(tag, str string, param ...any) {
	msg := fmt.Sprintf(str, param...)
	Logger.Debug().Caller(1).Str("[tag]", tag).Msg(msg)
}
