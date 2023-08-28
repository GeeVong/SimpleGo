package log

import (
	"github.com/anqiansong/ketty/console"
)

const (
	LogPackage_type         = 0
	LogPackage_kettyconsole = 1
	LogPackage_zerolog      = 2
)

var LogPackage = LogPackage_zerolog

func LogInfo(str string) {
	switch LogPackage {
	case LogPackage_kettyconsole:
		console.Info(str)
	case LogPackage_zerolog:
		Logger.Info().Str("key", "value").Msg(str)
	}
}

/*
 	1. 多个日志str 输出问题
	2. 日志颜色显示问题
	3. 日志格式设置问题
*/
