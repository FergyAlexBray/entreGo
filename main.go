package main

import (
	"os"

	entrego "github.com/FergyAlexBray/entreGo/src"
)

func main() {
	core := entrego.Core{}

	entrego.Parser(&core, os.Args)

	core.Run()
}
