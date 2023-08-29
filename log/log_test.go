package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"os"
	"runtime"
	"strconv"
	"testing"
	"time"
)

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
