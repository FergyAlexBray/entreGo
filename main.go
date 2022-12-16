package main

import (
	entrego "github.com/FergyAlexBray/entreGo/src"
)

func main() {
	core := entrego.Core{}

	entrego.Parser(&core)

	core.Run()

	entrego.EndMessage(core)
}
