package main

import (
	"PotionomicsCalculator/src"
	"bufio"
	"os"
)

func main() {
	src.In = bufio.NewReader(os.Stdin)
	src.Out = bufio.NewWriter(os.Stdout)
	src.PrintWithBufio("Starting program... ")
	src.Initialize()
	src.PrintWithBufio("Initialization complete.\n----------------------\n")
	src.MainLoop()
}
