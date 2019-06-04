package logger

import (
	"github.com/op/go-logging"
)

var Logger = logging.MustGetLogger("account")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func init() {
	logging.SetFormatter(format)
	logging.SetLevel(logging.INFO, "")
}
