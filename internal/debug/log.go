package debug

import (
	"fmt"
	"log"
	"os"
)

var logFile *os.File

func UseDebugMode() func() {
	f, err := os.Create("debug.log")
	if err != nil {
		log.Fatalln(err)
	}
	logFile = f

	return func() {
		f.Close()
	}
}

func Printf(format string, arg ...any) {
	logFile.WriteString(fmt.Sprintf(format, arg...) + "\n")
}
