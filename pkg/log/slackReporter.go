package log

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/sys/unix"
)

const (
	reportURL = "https://us-central1-tensorleap-ops3.cloudfunctions.net/demo-contact-bot"
)

var cloudMessageUuid string
var commandName string
var disableResporting bool

type MessageLevel string

const (
	InfoLog    MessageLevel = "info"
	WarningLog MessageLevel = "warning"
	ErrorLog   MessageLevel = "error"
)

type MessageState string

const (
	Starting MessageState = "Starting"
	Running  MessageState = "Running"
	Failed   MessageState = "Failed"
	Success  MessageState = "Success"
)

func SendCloudReport(messageLevel MessageLevel, message string, messageState MessageState, payload *map[string]interface{}) {
	switch messageState {
	case Starting, Running, Success:
		VerboseLogger.Infof("%s, %v", message, payload)
	case Failed:
		VerboseLogger.Errorf("%s, %v", message, payload)
	}
	if disableResporting {
		return
	}

	unameData := getUnameData()
	fullMessageObject := map[string]interface{}{"message": message, "state": messageState, "osName": runtime.GOOS, "commandName": commandName, "level": messageLevel,
		"installId": cloudMessageUuid, "payload": payload, "unameData": unameData}

	jsonData, err := json.Marshal(fullMessageObject)
	if err != nil {
		VerboseLogger.Error("Error:", err)
		return
	}

	req, _ := http.NewRequest("POST", reportURL, bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		VerboseLogger.Errorf("POST request failed: %v\n", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		VerboseLogger.Errorf("POST request failed with status code: %d\n", res.StatusCode)
		body, _ := io.ReadAll(res.Body)
		VerboseLogger.Errorf("res body: %s\n", string(body))
		return
	}

	defer res.Body.Close()
	return
}

func getUnameData() map[string]interface{} {
	var unameInfo unix.Utsname

	err := unix.Uname(&unameInfo)
	if err != nil {
		return map[string]interface{}{}
	}

	sysname := strings.ReplaceAll(string(unameInfo.Sysname[:]), "\u0000", "")
	nodename := strings.ReplaceAll(string(unameInfo.Nodename[:]), "\u0000", "")
	release := strings.ReplaceAll(string(unameInfo.Release[:]), "\u0000", "")
	version := strings.ReplaceAll(string(unameInfo.Version[:]), "\u0000", "")
	machine := strings.ReplaceAll(string(unameInfo.Machine[:]), "\u0000", "")

	return map[string]interface{}{"sysname": sysname, "nodename": nodename, "release": release, "version": version, "machine": machine}
}

func SetCommandName(name string) {
	commandName = name
}

func init() {
	cloudMessageUuid = uuid.New().String()
	disableResporting = os.Getenv("DISABLE_REPORTING") == "true"
}
