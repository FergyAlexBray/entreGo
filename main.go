package main

import (
	entrego "github.com/FergyAlexBray/entreGo/src"
	"os"
)

func main() {
	core := entrego.Core{}

	entrego.Parser(&core, os.Args)

	core.Run()

	entrego.EndMessage()
}
