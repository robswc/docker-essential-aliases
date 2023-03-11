package util

import (
	"fmt"
	"strings"
)

func HandleVerbose(verbose bool, command []string) {
	if verbose {
		fmt.Printf("Docker Command: %s\n", strings.Join(command, " "))
	}
}
