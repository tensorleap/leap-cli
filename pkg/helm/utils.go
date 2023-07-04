package helm

type Record = map[string]interface{}

const (
	RELEASE_NAME="tensorleap"
	CHART_NAME="tensorleap"
	REPO_URL="https://helm.tensorleap.ai"
)

func CreateTensorleapChartValues(useGpu bool, dataDir string) Record {
	return Record{
		"tensorleap-engine": Record{
			"gpu":                useGpu,
			"localDataDirectory": dataDir,
		},
	}
}