package main

import (
	"github.com/philcantcode/asiem-core-utils/discovery"
	"github.com/philcantcode/asiem-core-utils/logutils"
)

func main() {
	logutils.Print("a")

	discovery.PingScan()
}
