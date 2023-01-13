package main

import (
	"os"

	entrego "github.com/FergyAlexBray/entreGo/src"
)

func main() {
	core := entrego.Core{}

	ok := entrego.Parser(&core, os.Args)

	if ok {
		core.Run()
	}
}
