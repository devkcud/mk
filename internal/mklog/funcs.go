package mklog

import (
	"github.com/fatih/color"
	"time"
)

func date() string {
	now := time.Now().Format("15:04:05")
	return color.HiBlackString(now)
}
