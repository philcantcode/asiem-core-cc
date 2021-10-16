package main

import (
	"github.com/philcantcode/asiem-core-cc/discovery"
	"github.com/philcantcode/asiem-core-utils/logutils"
)

func main() {
	logutils.Print("a")

	discovery.PingScan()
}
