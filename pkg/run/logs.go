package run

import (
	"archive/tar"
	"context"
	"fmt"
	"os"
	"regexp"
	"slices"
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
	// Track last occurrence index of each unique line
	var allMatchingLines []string

	// Collect all matching lines from all pods, tracking last occurrence
	for _, log := range logs {
		lines := strings.Split(log.Logs, "\n")
		for _, line := range lines {
			trimmedLine := strings.TrimSpace(line)
			if trimmedLine == "" {
				continue
			}
			if pat.MatchString(line) {
				allMatchingLines = append(allMatchingLines, line)
			}
		}
	}

	// Build result with only last occurrence of each unique line, preserving order
	top := []string{}
	seen := make(map[string]bool)
	slices.Reverse(allMatchingLines)
	for _, line := range allMatchingLines {
		if !seen[line] {
			seen[line] = true
			top = append(top, line)
		}
		if len(top) >= count {
			break
		}
	}
	return top
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
