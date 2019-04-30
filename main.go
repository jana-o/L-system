package main

import (
	"code/lsystem-v2/ls"
	"code/lsystem-v2/turtlego"
)

func main() {
	l := ls.NewLsystem("F", "F", "F+F-")
	l.Generate(2)

	turtlego.ToPNG(l, "l-system.png")
}
