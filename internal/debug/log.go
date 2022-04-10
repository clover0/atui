package debug

import (
	"fmt"
	"log"
	"os"
)

var logFile *os.File

func init() {
	f, err := os.Create("debug.log")
	if err != nil {
		log.Fatalln(err)
	}
	logFile = f
}

func Printf(format string, arg ...any) {
	logFile.WriteString(fmt.Sprintf(format, arg...) + "\n")
}
