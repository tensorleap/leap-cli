package run

import (
	"archive/tar"
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/tensorleap/leap-cli/pkg/api"
	tlApi "github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func GetRunLogs(ctx context.Context, jobId string) ([]tlApi.PodLogs, error) {

	job, _, err := api.ApiClient.GetJobLogs(ctx).GetJobLogsParams(tlApi.GetJobLogsParams{JobId: jobId}).Execute()
	if err != nil {
		return nil, err
	}

	return job.PodsLogs, nil
}

func GetTopLogs(logs []tlApi.PodLogs, pat *regexp.Regexp, count int) []string {
	var result []string
	seen := make(map[string]bool)

	for _, podLog := range logs {
		lines := strings.Split(podLog.Logs, "\n")

		var matchingLines []string
		for _, line := range lines {
			if strings.TrimSpace(line) == "" {
				continue
			}
			if pat.MatchString(line) {
				matchingLines = append(matchingLines, line)
			}
		}

		// Take the last `count` lines from this pod (end of file = most recent)
		startIdx := len(matchingLines) - count
		if startIdx < 0 {
			startIdx = 0
		}
		topFromPod := matchingLines[startIdx:]

		// Add unique lines to result
		for _, line := range topFromPod {
			if !seen[line] {
				seen[line] = true
				result = append(result, line)
			}
		}
	}

	return result
}

func WrapLogsInTarFile(logs []tlApi.PodLogs, output string, defaultFileName string) (string, error) {
	isOutputDir := strings.HasSuffix(output, "/")
	if !strings.HasSuffix(output, ".tar") {
		if isOutputDir {
			output = fmt.Sprintf("%slogs-%s.tar", output, defaultFileName)
		} else {
			output = fmt.Sprintf("%s.tar", output)
		}
	}
	outputFile, err := os.Create(output)
	if err != nil {
		return "", err
	}
	defer func() { _ = outputFile.Close() }()
	tarWriter := tar.NewWriter(outputFile)
	for index, log := range logs {
		fileName := log.Name
		if fileName == "" {
			fileName = fmt.Sprintf("log-%d.txt", index)
		}
		err := tarWriter.WriteHeader(&tar.Header{
			Name: fileName,
			Size: int64(len(log.Logs)),
		})
		if err != nil {
			return "", err
		}
		_, err = tarWriter.Write([]byte(log.Logs))
		if err != nil {
			return "", err
		}
	}
	_ = tarWriter.Close()
	return output, nil
}

func PrintLogs(logs []tlApi.PodLogs) {
	for _, log := range logs {
		fmt.Printf("Logs for %s:\n", log.Name)
		fmt.Println("--------------------------------")
		fmt.Println("")
		fmt.Println(log.Logs)
		fmt.Println("--------------------------------")
		fmt.Println("")
	}
}
