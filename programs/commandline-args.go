package programs

import (
	"fmt"
	"os"
)

// printCommandLineArgs prints the commandline args
func PrintCommandLineArgs() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("%+v\n", os.Args[i])
	}
}
