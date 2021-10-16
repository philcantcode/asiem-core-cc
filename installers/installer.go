package installers

import (
	"fmt"

	"github.com/philcantcode/asiem-core-utils/osutils"
)

var programs [1]string

func installer() {
	programs[0] = "nmap-7.92-setup.exe"
}

func installPrerequisites() {
	for _, program := range programs {
		fmt.Println(program)
	}
}

func nmapInstaller() {
	if osutils.IdentifyOS() == "windows" {
		//osutils.Run("ping", "8.8.8.8")
	}
}
