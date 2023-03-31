package main

import (
	"github.com/take-2405/defaultcasechecker"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(
		defaultcasechecker.Analyzer,
	)
}
