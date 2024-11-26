package log

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
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

// Spinner is a custom type that wraps a spinner with additional functionality.
type Spinner struct {
	s          *spinner.Spinner
	message    string
	percentage int
}

// NewSpinner creates a new Spinner instance with the specified message.
func NewSpinner(message string, a ...any) *Spinner {
	message = fmt.Sprintf(message, a...)
	s := spinner.New(spinner.CharSets[33], 500*time.Millisecond)
	s.Suffix = fmt.Sprintf(" %s", message)

	return &Spinner{
		s:       s,
		message: message,
	}
}

// Start begins the spinner animation.
func (sp *Spinner) Start() {
	sp.s.Start()
}

// Stop stops the spinner animation and prints a completion message.
func (sp *Spinner) Stop() {
	sp.s.Stop()
	Infof("%s - Completed\n", sp.message)
}

// UpdateProgress updates the progress percentage displayed with the spinner.
func (sp *Spinner) UpdateProgress(percent int) {
	if percent > 100 {
		percent = 100
	}
	sp.percentage = percent
	sp.s.Suffix = fmt.Sprintf(" %s [%d%%]", sp.message, percent)
}
