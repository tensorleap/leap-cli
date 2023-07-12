package log

import (
	"fmt"
	"os"
)

func ConnectFileToVerboseLogOutput(filePath string) (close func(), err error) {

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		err = fmt.Errorf("Error opening log file: %s", err)
		return
	}

	VerboseLoggerOutputs.Add(file)

	close = func() {
		VerboseLoggerOutputs.Remove(file)
		file.Close()
	}
	return
}
