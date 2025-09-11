package run

import (
	"archive/tar"
	"context"
	"fmt"
	"math"
	"os"
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

func GetTopLogs(logs []tlApi.PodLogs, prefix string, count int) string {
	logsStr := ""
	for _, log := range logs {
		if strings.HasPrefix(log.Name, prefix) {
			s := strings.Split(log.Logs, "\n")
			fromIndex := int(math.Max(0, float64(len(s)-count)))
			logsStr = strings.Join(s[fromIndex:], "\n")
			break
		}
	}
	return logsStr
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
	defer outputFile.Close()
	tarWriter := tar.NewWriter(outputFile)
	for index, log := range logs {
		fileName := log.Name
		if fileName == "" {
			fileName = fmt.Sprintf("log-%d.txt", index)
		}
		err := tarWriter.WriteHeader(&tar.Header{
			Name: fileName,
		})
		if err != nil {
			return "", err
		}
		_, err = tarWriter.Write([]byte(log.Logs))
		if err != nil {
			return "", err
		}
	}
	tarWriter.Close()
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
