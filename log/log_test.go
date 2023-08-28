package log

import (
	"errors"
	"fmt"
	"github.com/anqiansong/ketty/console"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"os"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestLogDebug(t *testing.T) {
	console.Info(`
    {
        "name":"Hello Ketty",
        "description":"a color logger",
        "author":"anqiansong",
        "category":"console",
        "github":"https://github.com/anqiansong/ketty",
        "useage":[
            "info",
            "debug"
        ]
    }`)
	console.Debug("Hello Ketty")
	console.Warn("Hello Ketty")
	console.Error(errors.New("error test"))
}

func TestLogColor(t *testing.T) {
	fmt.Printf("\033[1;31;40m%s\033[0m\n", "Red.")
	fmt.Printf("\033[1;37;41m%s\033[0m\n", "Red.")

}

// 控制台输出带颜色
func TestZeroLogConsoleWriter(t *testing.T) {
	beautifulLogger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	_, file, line, _ := runtime.Caller(0)
	beautifulLogger.Info().Str("file", file).Str("line", strconv.Itoa(line)).Msg("str")
	time.Sleep(1)
	beautifulLogger.Error().Str("file", file).Str("line", strconv.Itoa(line)).Msg("str")
}
