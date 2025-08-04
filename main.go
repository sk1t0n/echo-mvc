package main

import (
	"github.com/fatih/color"
	"github.com/sk1t0n/echo-mvc/internal/cli"
)

const version = "v0.1.0"

func main() {
	color.Blue("CLI echo-mvc " + version)
	cli.Run()
}
