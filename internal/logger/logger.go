package logger

import (
	"fmt"
	"time"
)

type Color string

const (
	Red     Color = "\033[31m"
	Green   Color = "\033[32m"
	Yellow  Color = "\033[33m"
	Blue    Color = "\033[34m"
	Magenta Color = "\033[35m"
	Pink    Color = "\033[95m"
	Cyan    Color = "\033[36m"
	White   Color = "\033[37m"

	Bold Color = "\033[1m"

	BgGreen Color = "\033[42m"
	BgRed   Color = "\033[41m"
	BgCyan  Color = "\033[46m"
	BgBlue  Color = "\033[44m"
	BgReset Color = "\033[49m"
	BgGrey  Color = "\033[100m"

	Reset Color = "\033[0m"
)

func PrintStatusLineKV(color1 Color, process string, color2 Color, tag string, color3 Color, value string) {
	// ANSI Grey
	fmt.Printf("\033[90m%s\033[0m ", time.Now().Format("15:04:05"))
	fmt.Printf("%s%s%s ", color1, process, Reset)
	fmt.Printf("%s%s%s ", color2, tag, Reset)
	fmt.Printf("%s%s%s\n", color3, value, Reset)
}
