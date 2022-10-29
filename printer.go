package main

import (
	"fmt"
	"strings"
)

type Color string

const (
	reset  Color = "\033[0m"
	black  Color = "\033[30m"
	red    Color = "\033[31m"
	green  Color = "\033[32m"
	yellow Color = "\033[33m"
	blue   Color = "\033[34m"
	purple Color = "\033[35m"
	cyan   Color = "\033[36m"
	white  Color = "\033[37m"
)

const (
	boxTopLeft     = "╭"
	boxTopRight    = "╮"
	boxBottomLeft  = "╰"
	boxBottomRight = "╯"
	boxSide        = "│"
	boxBoundary    = "─"
)

func setColor(color Color) {
	fmt.Printf("%v", color)
}

func resetColor() {
	fmt.Printf("%v", reset)
}

func boxPrint(color Color, s []string) {
	setColor(color)
	length := maxLen(s)
	boundary := strings.Repeat(boxBoundary, length+3)
	fmt.Printf("%v%v%v\n", boxTopLeft, boundary, boxTopRight)
	for _, v := range s {
		gap := strings.Repeat(" ", length-len(v))
		fmt.Printf("%v %v %v %v\n", boxSide, v, gap, boxSide)
	}
	fmt.Printf("%v%v%v\n", boxBottomLeft, boundary, boxBottomRight)
	resetColor()
}

func colorPrintf(color Color, s string, args ...any) {
	setColor(color)
	fmt.Printf(s, args...)
	resetColor()
}
