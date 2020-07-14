package programs

import (
	"fmt"
	"os"
)

// printCommandLineArgs prints the commandline args
func printCommandLineArgs() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("%+v\n", os.Args[i])
	}
}
