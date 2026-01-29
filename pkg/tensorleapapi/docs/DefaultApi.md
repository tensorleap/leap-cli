# \DefaultAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Activate**](DefaultAPI.md#Activate) | **Post** /auth/activate | 
[**AddDashboard**](DefaultAPI.md#AddDashboard) | **Post** /dashboards/addDashboard | 
[**AddExportModelJob**](DefaultAPI.md#AddExportModelJob) | **Post** /exportedsessionruns/addExportModelJob | 
[**AddIssue**](DefaultAPI.md#AddIssue) | **Post** /issues/addIssue | 
[**AddProject**](DefaultAPI.md#AddProject) | **Post** /projects/addProject | 
[**AddSampleCollection**](DefaultAPI.md#AddSampleCollection) | **Post** /sample-collection/addSampleCollection | 
[**AddSecretManager**](DefaultAPI.md#AddSecretManager) | **Post** /secret-manager/addSecretManager | 
[**ApplyInsightsStatus**](DefaultAPI.md#ApplyInsightsStatus) | **Post** /insights/applyInsightsStatus | 
[**CalcPopulationExplorationDigest**](DefaultAPI.md#CalcPopulationExplorationDigest) | **Post** /dashboards/calcPopulationExplorationDigest | 
[**ClearUserJobs**](DefaultAPI.md#ClearUserJobs) | **Post** /users/clearUserJobs | 
[**ClearUserNotifications**](DefaultAPI.md#ClearUserNotifications) | **Post** /users/clearUserNotifications | 
[**ContinueEvaluate**](DefaultAPI.md#ContinueEvaluate) | **Post** /evaluate/continueEvaluate | 
[**CreateSamplesVisualizations**](DefaultAPI.md#CreateSamplesVisualizations) | **Post** /visualizations/createSamplesVisualizations | 
[**CreateSessionTest**](DefaultAPI.md#CreateSessionTest) | **Post** /sessions-tests/createSessionTest | 
[**CreateTeam**](DefaultAPI.md#CreateTeam) | **Post** /teams/createTeam | 
[**DeleteDashboard**](DefaultAPI.md#DeleteDashboard) | **Post** /dashboards/deleteDashboard | 
[**DeleteGeneratedLabel**](DefaultAPI.md#DeleteGeneratedLabel) | **Post** /datasetcuration/deleteGeneratedLabel | 
[**DeleteIssue**](DefaultAPI.md#DeleteIssue) | **Post** /issues/deleteIssue | 
[**DeleteProject**](DefaultAPI.md#DeleteProject) | **Post** /projects/deleteProject | 
[**DeleteSampleAnalysis**](DefaultAPI.md#DeleteSampleAnalysis) | **Post** /visualizations/deleteSampleAnalysis | 
[**DeleteSession**](DefaultAPI.md#DeleteSession) | **Post** /sessions/deleteSession | 
[**DeleteSessionRun**](DefaultAPI.md#DeleteSessionRun) | **Post** /sessions/deleteSessionRun | 
[**DeleteSessionTest**](DefaultAPI.md#DeleteSessionTest) | **Post** /sessions-tests/deleteSessionTest | 
[**DeleteSyntheticData**](DefaultAPI.md#DeleteSyntheticData) | **Post** /datasetcuration/deleteSyntheticData | 
[**DeleteTeam**](DefaultAPI.md#DeleteTeam) | **Post** /teams/deleteTeam | 
[**DeleteUserById**](DefaultAPI.md#DeleteUserById) | **Post** /users/deleteUserById | 
[**DeleteVersion**](DefaultAPI.md#DeleteVersion) | **Post** /versions/deleteVersion | 
[**DeleteVisualizations**](DefaultAPI.md#DeleteVisualizations) | **Post** /visualizations/deleteVisualizations | 
[**DownloadProject**](DefaultAPI.md#DownloadProject) | **Get** /projects/downloadProject/{projectId} | 
[**Evaluate**](DefaultAPI.md#Evaluate) | **Post** /evaluate/evaluate | 
[**ExportProject**](DefaultAPI.md#ExportProject) | **Post** /projects/exportProject | 
[**ExtendTrial**](DefaultAPI.md#ExtendTrial) | **Post** /auth/extendTrial | 
[**FetchSimilar**](DefaultAPI.md#FetchSimilar) | **Post** /visualizations/fetchSimilar | 
[**GenerateInsights**](DefaultAPI.md#GenerateInsights) | **Post** /insights/generateInsights | 
[**GenerateLabels**](DefaultAPI.md#GenerateLabels) | **Post** /datasetcuration/generateLabels | 
[**GenerateSyntheticData**](DefaultAPI.md#GenerateSyntheticData) | **Post** /datasetcuration/generateSyntheticData | 
[**GetAllProjectSessionTests**](DefaultAPI.md#GetAllProjectSessionTests) | **Post** /sessions-tests/getAllProjectSessionTests | 
[**GetAllSlimUserData**](DefaultAPI.md#GetAllSlimUserData) | **Post** /users/getAllSlimUserData | 
[**GetApiKeyByCode**](DefaultAPI.md#GetApiKeyByCode) | **Post** /auth/getApiKeyByCode | 
[**GetAuthProvider**](DefaultAPI.md#GetAuthProvider) | **Post** /auth/getAuthProvider | 
[**GetAuthStatus**](DefaultAPI.md#GetAuthStatus) | **Post** /auth/getAuthStatus | 
[**GetBalancedAccuracy**](DefaultAPI.md#GetBalancedAccuracy) | **Post** /sessionmetrics/getBalancedAccuracy | 
[**GetCodeSnapshot**](DefaultAPI.md#GetCodeSnapshot) | **Post** /versions/getCodeSnapshot | 
[**GetCodeSnapshotUploadUrl**](DefaultAPI.md#GetCodeSnapshotUploadUrl) | **Post** /versions/getCodeSnapshotUploadUrl | 
[**GetConfusionMatrixLabels**](DefaultAPI.md#GetConfusionMatrixLabels) | **Post** /sessionmetrics/getConfusionMatrixLabels | 
[**GetConfusionMatrixResultCombinations**](DefaultAPI.md#GetConfusionMatrixResultCombinations) | **Post** /sessionmetrics/getConfusionMatrixResultCombinations | 
[**GetConfusionMatrixTable**](DefaultAPI.md#GetConfusionMatrixTable) | **Post** /sessionmetrics/getConfusionMatrixTable | 
[**GetConfusionMetricNames**](DefaultAPI.md#GetConfusionMetricNames) | **Post** /sessionmetrics/getConfusionMetricNames | 
[**GetCurrentProjectVersion**](DefaultAPI.md#GetCurrentProjectVersion) | **Post** /projects/getCurrentProjectVersion | 
[**GetDashboard**](DefaultAPI.md#GetDashboard) | **Post** /dashboards/getDashboard | 
[**GetDashletFields**](DefaultAPI.md#GetDashletFields) | **Post** /dashboards/getDashletFields | 
[**GetDownloadSignedUrl**](DefaultAPI.md#GetDownloadSignedUrl) | **Post** /versions/getDownloadSignedUrl | 
[**GetEngineSettings**](DefaultAPI.md#GetEngineSettings) | **Post** /settings/getEngineSettings | 
[**GetEnvironmentInfo**](DefaultAPI.md#GetEnvironmentInfo) | **Post** /metadata/getEnvironmentInfo | 
[**GetExportedSessionJobs**](DefaultAPI.md#GetExportedSessionJobs) | **Post** /exportedsessionruns/getExportedSessionJobs | 
[**GetF1Score**](DefaultAPI.md#GetF1Score) | **Post** /sessionmetrics/getF1Score | 
[**GetFetchSimilarStatus**](DefaultAPI.md#GetFetchSimilarStatus) | **Post** /visualizations/getFetchSimilarStatus | 
[**GetFieldsValues**](DefaultAPI.md#GetFieldsValues) | **Post** /sessionmetrics/getFieldsValues | 
[**GetGeneratedLabels**](DefaultAPI.md#GetGeneratedLabels) | **Post** /datasetcuration/getGeneratedLabels | 
[**GetGenericBaseImageTypes**](DefaultAPI.md#GetGenericBaseImageTypes) | **Post** /versions/getGenericBaseImageTypes | 
[**GetHeatmapChart**](DefaultAPI.md#GetHeatmapChart) | **Post** /sessionmetrics/getHeatmapChart | 
[**GetInsights**](DefaultAPI.md#GetInsights) | **Post** /insights/getInsights | 
[**GetIssueFileUploadSignedUrl**](DefaultAPI.md#GetIssueFileUploadSignedUrl) | **Post** /issues/getIssueFileUploadSignedUrl | 
[**GetJobLogs**](DefaultAPI.md#GetJobLogs) | **Post** /jobs/getJobLogs | 
[**GetLatestExportedProject**](DefaultAPI.md#GetLatestExportedProject) | **Post** /projects/getLatestExportedProject | 
[**GetMachineTypes**](DefaultAPI.md#GetMachineTypes) | **Post** /teams/getMachineTypes | 
[**GetMaxActiveUsers**](DefaultAPI.md#GetMaxActiveUsers) | **Post** /metadata/getMaxActiveUsers | 
[**GetMeanAveragePrecision**](DefaultAPI.md#GetMeanAveragePrecision) | **Post** /sessionmetrics/getMeanAveragePrecision | 
[**GetNotifications**](DefaultAPI.md#GetNotifications) | **Post** /notifications/getNotifications | 
[**GetNotificationsByFilter**](DefaultAPI.md#GetNotificationsByFilter) | **Post** /notifications/getNotificationsByFilter | 
[**GetPopulationExplorationStatus**](DefaultAPI.md#GetPopulationExplorationStatus) | **Post** /visualizations/getPopulationExplorationStatus | 
[**GetPrCurve**](DefaultAPI.md#GetPrCurve) | **Post** /sessionmetrics/getPrCurve | 
[**GetPrecisionScore**](DefaultAPI.md#GetPrecisionScore) | **Post** /sessionmetrics/getPrecisionScore | 
[**GetProjectDashboards**](DefaultAPI.md#GetProjectDashboards) | **Post** /dashboards/getProjectDashboards | 
[**GetProjectIssues**](DefaultAPI.md#GetProjectIssues) | **Post** /issues/getProjectIssues | 
[**GetProjectSlimVersions**](DefaultAPI.md#GetProjectSlimVersions) | **Post** /versions/getProjectSlimVersions | 
[**GetProjects**](DefaultAPI.md#GetProjects) | **Post** /projects/getProjects | 
[**GetRecallScore**](DefaultAPI.md#GetRecallScore) | **Post** /sessionmetrics/getRecallScore | 
[**GetRecentTeamSessions**](DefaultAPI.md#GetRecentTeamSessions) | **Post** /sessions/getRecentTeamSessions | 
[**GetRoc**](DefaultAPI.md#GetRoc) | **Post** /sessionmetrics/getRoc | 
[**GetSampleCollections**](DefaultAPI.md#GetSampleCollections) | **Post** /sample-collection/getSampleCollections | 
[**GetSampleVisualizationsPath**](DefaultAPI.md#GetSampleVisualizationsPath) | **Post** /visualizations/getSampleVisualizationsPath | 
[**GetScatterSampleVisualizations**](DefaultAPI.md#GetScatterSampleVisualizations) | **Post** /visualizations/getScatterSampleVisualizations | 
[**GetSecretManagerList**](DefaultAPI.md#GetSecretManagerList) | **Post** /secret-manager/getSecretManagerList | 
[**GetSessionRunsEvaluate**](DefaultAPI.md#GetSessionRunsEvaluate) | **Post** /sessions/getSessionRunsEvaluate | 
[**GetSessionRunsVisualizations**](DefaultAPI.md#GetSessionRunsVisualizations) | **Post** /visualizations/getSessionRunsVisualizations | 
[**GetSessionTestResult**](DefaultAPI.md#GetSessionTestResult) | **Post** /sessions-tests/getSessionTestResult | 
[**GetSessionsByHash**](DefaultAPI.md#GetSessionsByHash) | **Post** /sessions/getSessionsByHash | 
[**GetSessionsByVersionId**](DefaultAPI.md#GetSessionsByVersionId) | **Post** /sessions/getSessionsByVersionId | 
[**GetSessionsEpochs**](DefaultAPI.md#GetSessionsEpochs) | **Post** /versions/getSessionsEpochs | 
[**GetSignedUrl**](DefaultAPI.md#GetSignedUrl) | **Post** /versions/getSignedUrl | 
[**GetSingleIssue**](DefaultAPI.md#GetSingleIssue) | **Post** /issues/getSingleIssue | 
[**GetSingleSessionTest**](DefaultAPI.md#GetSingleSessionTest) | **Post** /sessions-tests/getSingleSessionTest | 
[**GetSlimJobs**](DefaultAPI.md#GetSlimJobs) | **Post** /jobs/getSlimJobs | 
[**GetSlimVisualization**](DefaultAPI.md#GetSlimVisualization) | **Post** /visualizations/getSlimVisualization | 
[**GetState**](DefaultAPI.md#GetState) | **Post** /projectstate/getState | 
[**GetStatistics**](DefaultAPI.md#GetStatistics) | **Post** /metadata/getStatistics | 
[**GetSyntheticData**](DefaultAPI.md#GetSyntheticData) | **Post** /datasetcuration/getSyntheticData | 
[**GetTableChart**](DefaultAPI.md#GetTableChart) | **Post** /sessionmetrics/getTableChart | 
[**GetTeamJobs**](DefaultAPI.md#GetTeamJobs) | **Post** /jobs/getTeamJobs | 
[**GetTeamSlimUserData**](DefaultAPI.md#GetTeamSlimUserData) | **Post** /users/getTeamSlimUserData | 
[**GetTeams**](DefaultAPI.md#GetTeams) | **Post** /teams/getTeams | 
[**GetUploadModelSignedUrl**](DefaultAPI.md#GetUploadModelSignedUrl) | **Post** /versions/getUploadModelSignedUrl | 
[**GetUploadSignedUrl**](DefaultAPI.md#GetUploadSignedUrl) | **Post** /versions/getUploadSignedUrl | 
[**GetVisualization**](DefaultAPI.md#GetVisualization) | **Post** /visualizations/getVisualization | 
[**GetXYChart**](DefaultAPI.md#GetXYChart) | **Post** /sessionmetrics/getXYChart | 
[**HealthCheck**](DefaultAPI.md#HealthCheck) | **Get** /monitor/healthCheck | 
[**ImportExternalModel**](DefaultAPI.md#ImportExternalModel) | **Post** /versions/importExternalModel | 
[**ImportModel**](DefaultAPI.md#ImportModel) | **Post** /versions/importModel | 
[**ImportProject**](DefaultAPI.md#ImportProject) | **Post** /projects/importProject | 
[**InitExperiment**](DefaultAPI.md#InitExperiment) | **Post** /versions/initExperiment | 
[**KeyGen**](DefaultAPI.md#KeyGen) | **Post** /auth/keygen | 
[**LoadModel**](DefaultAPI.md#LoadModel) | **Post** /projects/loadModel | 
[**LoadVersion**](DefaultAPI.md#LoadVersion) | **Post** /versions/loadVersion | 
[**LocalAuth**](DefaultAPI.md#LocalAuth) | **Post** /auth/localAuth | 
[**LogExternalEpochData**](DefaultAPI.md#LogExternalEpochData) | **Post** /externalepochdata/logExternalEpochData | 
[**Login**](DefaultAPI.md#Login) | **Post** /auth/login | 
[**Logout**](DefaultAPI.md#Logout) | **Post** /auth/logout | 
[**OverwriteModel**](DefaultAPI.md#OverwriteModel) | **Post** /versions/overwriteModel | 
[**PopulationExploration**](DefaultAPI.md#PopulationExploration) | **Post** /visualizations/populationExploration | 
[**PushCodeSnapshot**](DefaultAPI.md#PushCodeSnapshot) | **Post** /versions/pushCodeSnapshot | 
[**RefreshLocalAuth**](DefaultAPI.md#RefreshLocalAuth) | **Post** /auth/refreshLocalAuth | 
[**RemoveSamplesCollection**](DefaultAPI.md#RemoveSamplesCollection) | **Post** /sample-collection/removeSampleCollections | 
[**ResolveConcurrentUsersConflict**](DefaultAPI.md#ResolveConcurrentUsersConflict) | **Post** /auth/resolveConcurrentUsersConflict | 
[**SampleAnalysis**](DefaultAPI.md#SampleAnalysis) | **Post** /visualizations/sampleAnalysis | 
[**SamplesVisualizationsRefresh**](DefaultAPI.md#SamplesVisualizationsRefresh) | **Post** /visualizations/samplesVisualizationsRefresh | 
[**SaveAnalyzerLayout**](DefaultAPI.md#SaveAnalyzerLayout) | **Post** /visualizations/saveAnalyzerLayout | 
[**SendUserMessage**](DefaultAPI.md#SendUserMessage) | **Post** /users/sendUserMessage | 
[**SetCodeChallenge**](DefaultAPI.md#SetCodeChallenge) | **Post** /auth/setCodeChallenge | 
[**SetDefaultTeam**](DefaultAPI.md#SetDefaultTeam) | **Post** /teams/setDefaultTeam | 
[**SetExperimentProperties**](DefaultAPI.md#SetExperimentProperties) | **Post** /versions/setExperimentProperties | 
[**SetMachineType**](DefaultAPI.md#SetMachineType) | **Post** /teams/setMachineType | 
[**SetUserNotificationsAsRead**](DefaultAPI.md#SetUserNotificationsAsRead) | **Post** /notifications/setUserNotificationsAsRead | 
[**StartTrial**](DefaultAPI.md#StartTrial) | **Post** /auth/startTrial | 
[**StopJob**](DefaultAPI.md#StopJob) | **Post** /jobs/stopJob | 
[**TagModel**](DefaultAPI.md#TagModel) | **Post** /versions/tagModel | 
[**TerminateAllJobs**](DefaultAPI.md#TerminateAllJobs) | **Post** /jobs/terminateAllJobs | 
[**TerminateJob**](DefaultAPI.md#TerminateJob) | **Post** /jobs/terminateJob | 
[**TrashSecretManager**](DefaultAPI.md#TrashSecretManager) | **Post** /secret-manager/trashSecretManager | 
[**UpdateDashboard**](DefaultAPI.md#UpdateDashboard) | **Post** /dashboards/updateDashboard | 
[**UpdateEngineSettings**](DefaultAPI.md#UpdateEngineSettings) | **Post** /settings/updateSetting | 
[**UpdateIssue**](DefaultAPI.md#UpdateIssue) | **Post** /issues/updateIssue | 
[**UpdateProjectMeta**](DefaultAPI.md#UpdateProjectMeta) | **Post** /projects/updateProjectMeta | 
[**UpdateSampleCollection**](DefaultAPI.md#UpdateSampleCollection) | **Post** /sample-collection/updateSampleCollection | 
[**UpdateSecretManager**](DefaultAPI.md#UpdateSecretManager) | **Post** /secret-manager/updateSecretManager | 
[**UpdateSessionName**](DefaultAPI.md#UpdateSessionName) | **Post** /sessions/updateSessionName | 
[**UpdateSessionRun**](DefaultAPI.md#UpdateSessionRun) | **Post** /sessionsruns/updateSessionRun | 
[**UpdateSessionTest**](DefaultAPI.md#UpdateSessionTest) | **Post** /sessions-tests/updateSessionTest | 
[**UpdateTeamPublicName**](DefaultAPI.md#UpdateTeamPublicName) | **Post** /teams/updateTeamPublicName | 
[**UpdateUserName**](DefaultAPI.md#UpdateUserName) | **Post** /users/updateUserName | 
[**UpdateUserRole**](DefaultAPI.md#UpdateUserRole) | **Post** /users/updateUserRole | 
[**UpdateUserStatus**](DefaultAPI.md#UpdateUserStatus) | **Post** /users/updateUserStatus | 
[**UpdateUserTeam**](DefaultAPI.md#UpdateUserTeam) | **Post** /users/updateUserTeam | 
[**UpdateVersion**](DefaultAPI.md#UpdateVersion) | **Post** /versions/updateVersion | 
[**UpdateVersionName**](DefaultAPI.md#UpdateVersionName) | **Post** /versions/updateVersionName | 
[**UploadProject**](DefaultAPI.md#UploadProject) | **Put** /projects/uploadProject/{projectName} | 
[**UpsertState**](DefaultAPI.md#UpsertState) | **Post** /projectstate/upsertState | 
[**Warmup**](DefaultAPI.md#Warmup) | **Post** /jobs/warmup | 
[**WhoAmI**](DefaultAPI.md#WhoAmI) | **Post** /auth/whoAmI | 



## Activate

> UserData Activate(ctx).ActivationParams(activationParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	activationParams := *openapiclient.NewActivationParams("Token_example") // ActivationParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Activate(context.Background()).ActivationParams(activationParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Activate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Activate`: UserData
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Activate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activationParams** | [**ActivationParams**](ActivationParams.md) |  | 

### Return type

[**UserData**](UserData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddDashboard

> AddDashboardResponse AddDashboard(ctx).AddDashboardParams(addDashboardParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	addDashboardParams := *openapiclient.NewAddDashboardParams("ProjectId_example", "Name_example", []openapiclient.Dashlet{*openapiclient.NewDashlet("Cid_example", *openapiclient.NewSizedLayout(*openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123))), "Type_example", *openapiclient.NewPopulationExplorationDashletData(map[string]interface{}(123), "Name_example", "Type_example"), "Name_example", []string{"CollectionIds_example"})}) // AddDashboardParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AddDashboard(context.Background()).AddDashboardParams(addDashboardParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDashboard`: AddDashboardResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addDashboardParams** | [**AddDashboardParams**](AddDashboardParams.md) |  | 

### Return type

[**AddDashboardResponse**](AddDashboardResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddExportModelJob

> Job AddExportModelJob(ctx).AddExportModelJobParams(addExportModelJobParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	addExportModelJobParams := *openapiclient.NewAddExportModelJobParams("ProjectId_example", "SessionWeightId_example", openapiclient.ExportModelTypeEnum("JSON_TF2"), "Title_example", false) // AddExportModelJobParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AddExportModelJob(context.Background()).AddExportModelJobParams(addExportModelJobParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddExportModelJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddExportModelJob`: Job
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddExportModelJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddExportModelJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addExportModelJobParams** | [**AddExportModelJobParams**](AddExportModelJobParams.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddIssue

> Issue AddIssue(ctx).AddIssueParams(addIssueParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	addIssueParams := *openapiclient.NewAddIssueParams("ProjectId_example", "Title_example", openapiclient.IssueStatus("Open")) // AddIssueParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AddIssue(context.Background()).AddIssueParams(addIssueParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddIssue`: Issue
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddIssue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addIssueParams** | [**AddIssueParams**](AddIssueParams.md) |  | 

### Return type

[**Issue**](Issue.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddProject

> AddProjectResponse AddProject(ctx).ProjectMeta(projectMeta).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	projectMeta := *openapiclient.NewProjectMeta("Name_example", "Description_example", []string{"Tags_example"}, openapiclient.HubPublishPolicy("public")) // ProjectMeta | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AddProject(context.Background()).ProjectMeta(projectMeta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddProject`: AddProjectResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectMeta** | [**ProjectMeta**](ProjectMeta.md) |  | 

### Return type

[**AddProjectResponse**](AddProjectResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSampleCollection

> AddSampleCollectionResponse AddSampleCollection(ctx).AddSampleCollectionParams(addSampleCollectionParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	addSampleCollectionParams := *openapiclient.NewAddSampleCollectionParams([]openapiclient.SampleIdentity{*openapiclient.NewSampleIdentity(openapiclient.DataStateType("training"), *openapiclient.NewSampleIdType())}, "ProjectId_example", "Name_example") // AddSampleCollectionParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AddSampleCollection(context.Background()).AddSampleCollectionParams(addSampleCollectionParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddSampleCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSampleCollection`: AddSampleCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddSampleCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSampleCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addSampleCollectionParams** | [**AddSampleCollectionParams**](AddSampleCollectionParams.md) |  | 

### Return type

[**AddSampleCollectionResponse**](AddSampleCollectionResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSecretManager

> AddSecretManagerResponse AddSecretManager(ctx).AddSecretManagerParams(addSecretManagerParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	addSecretManagerParams := *openapiclient.NewAddSecretManagerParams("Name_example", "AuthData_example") // AddSecretManagerParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AddSecretManager(context.Background()).AddSecretManagerParams(addSecretManagerParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddSecretManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSecretManager`: AddSecretManagerResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddSecretManager`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSecretManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addSecretManagerParams** | [**AddSecretManagerParams**](AddSecretManagerParams.md) |  | 

### Return type

[**AddSecretManagerResponse**](AddSecretManagerResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyInsightsStatus

> ApplyInsightsStatus(ctx).ApplyInsightsStatusParams(applyInsightsStatusParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	applyInsightsStatusParams := *openapiclient.NewApplyInsightsStatusParams("InsightId_example", openapiclient.InsightStatus("InReview"), "ProjectId_example") // ApplyInsightsStatusParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ApplyInsightsStatus(context.Background()).ApplyInsightsStatusParams(applyInsightsStatusParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApplyInsightsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyInsightsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applyInsightsStatusParams** | [**ApplyInsightsStatusParams**](ApplyInsightsStatusParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CalcPopulationExplorationDigest

> CalcPopulationExplorationDigestResponse CalcPopulationExplorationDigest(ctx).CalcPopulationExplorationDigestParams(calcPopulationExplorationDigestParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	calcPopulationExplorationDigestParams := *openapiclient.NewCalcPopulationExplorationDigestParams(*openapiclient.NewPickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail("SessionRunId_example", "ProjectId_example", float64(123), float64(123), float64(123), []string{"BalanceBy_example"}, false, openapiclient.ReductionAlgorithm("TSNE"))) // CalcPopulationExplorationDigestParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CalcPopulationExplorationDigest(context.Background()).CalcPopulationExplorationDigestParams(calcPopulationExplorationDigestParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CalcPopulationExplorationDigest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalcPopulationExplorationDigest`: CalcPopulationExplorationDigestResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CalcPopulationExplorationDigest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCalcPopulationExplorationDigestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calcPopulationExplorationDigestParams** | [**CalcPopulationExplorationDigestParams**](CalcPopulationExplorationDigestParams.md) |  | 

### Return type

[**CalcPopulationExplorationDigestResponse**](CalcPopulationExplorationDigestResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearUserJobs

> ClearUserJobs(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ClearUserJobs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ClearUserJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClearUserJobsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearUserNotifications

> ClearUserNotifications(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ClearUserNotifications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ClearUserNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClearUserNotificationsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContinueEvaluate

> Job ContinueEvaluate(ctx).ContinueEvaluateParams(continueEvaluateParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	continueEvaluateParams := *openapiclient.NewContinueEvaluateParams("SessionRunId_example", "ProjectId_example") // ContinueEvaluateParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ContinueEvaluate(context.Background()).ContinueEvaluateParams(continueEvaluateParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContinueEvaluate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContinueEvaluate`: Job
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ContinueEvaluate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContinueEvaluateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **continueEvaluateParams** | [**ContinueEvaluateParams**](ContinueEvaluateParams.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSamplesVisualizations

> Job CreateSamplesVisualizations(ctx).CreateSampleVisualizationsParams(createSampleVisualizationsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	createSampleVisualizationsParams := *openapiclient.NewCreateSampleVisualizationsParams("ProjectId_example", "SessionRunId_example", float64(123), "Digest_example") // CreateSampleVisualizationsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateSamplesVisualizations(context.Background()).CreateSampleVisualizationsParams(createSampleVisualizationsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateSamplesVisualizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSamplesVisualizations`: Job
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateSamplesVisualizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSamplesVisualizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSampleVisualizationsParams** | [**CreateSampleVisualizationsParams**](CreateSampleVisualizationsParams.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSessionTest

> string CreateSessionTest(ctx).CreateSessionTestRequest(createSessionTestRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	createSessionTestRequest := *openapiclient.NewCreateSessionTestRequest("ProjectId_example", "Name_example", *openapiclient.NewClientFilterParams("Field_example", openapiclient.FilterOperatorType("between"), float64(123))) // CreateSessionTestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateSessionTest(context.Background()).CreateSessionTestRequest(createSessionTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateSessionTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSessionTest`: string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateSessionTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSessionTestRequest** | [**CreateSessionTestRequest**](CreateSessionTestRequest.md) |  | 

### Return type

**string**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTeam

> CreateTeamResponse CreateTeam(ctx).CreateTeamRequest(createTeamRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	createTeamRequest := *openapiclient.NewCreateTeamRequest("Name_example", "PublicName_example") // CreateTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateTeam(context.Background()).CreateTeamRequest(createTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTeam`: CreateTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTeamRequest** | [**CreateTeamRequest**](CreateTeamRequest.md) |  | 

### Return type

[**CreateTeamResponse**](CreateTeamResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboard

> DeleteDashboard(ctx).DeleteDashboardParams(deleteDashboardParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	deleteDashboardParams := *openapiclient.NewDeleteDashboardParams("DashboardId_example", "ProjectId_example") // DeleteDashboardParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteDashboard(context.Background()).DeleteDashboardParams(deleteDashboardParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDashboardParams** | [**DeleteDashboardParams**](DeleteDashboardParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGeneratedLabel

> DeleteGeneratedLabel(ctx).DeleteGeneratedLabelParams(deleteGeneratedLabelParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	deleteGeneratedLabelParams := *openapiclient.NewDeleteGeneratedLabelParams("ProjectId_example", "LabelId_example") // DeleteGeneratedLabelParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteGeneratedLabel(context.Background()).DeleteGeneratedLabelParams(deleteGeneratedLabelParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteGeneratedLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGeneratedLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteGeneratedLabelParams** | [**DeleteGeneratedLabelParams**](DeleteGeneratedLabelParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIssue

> DeleteIssue(ctx).DeleteIssueParams(deleteIssueParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	deleteIssueParams := *openapiclient.NewDeleteIssueParams("Cid_example", "ProjectId_example") // DeleteIssueParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteIssue(context.Background()).DeleteIssueParams(deleteIssueParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteIssueParams** | [**DeleteIssueParams**](DeleteIssueParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProject

> DeleteProject(ctx).DeleteProjectParams(deleteProjectParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	deleteProjectParams := *openapiclient.NewDeleteProjectParams("ProjectId_example") // DeleteProjectParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteProject(context.Background()).DeleteProjectParams(deleteProjectParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteProjectParams** | [**DeleteProjectParams**](DeleteProjectParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSampleAnalysis

> DeleteSampleAnalysis(ctx).DeleteSamplesAnalysisParams(deleteSamplesAnalysisParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	deleteSamplesAnalysisParams := *openapiclient.NewDeleteSamplesAnalysisParams("ProjectId_example", []string{"SessionRunIds_example"}, []openapiclient.SampleIdentity{*openapiclient.NewSampleIdentity(openapiclient.DataStateType("training"), *openapiclient.NewSampleIdType())}) // DeleteSamplesAnalysisParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteSampleAnalysis(context.Background()).DeleteSamplesAnalysisParams(deleteSamplesAnalysisParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteSampleAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSampleAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSamplesAnalysisParams** | [**DeleteSamplesAnalysisParams**](DeleteSamplesAnalysisParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSession

> DeleteSession(ctx).DeleteSessionParams(deleteSessionParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	deleteSessionParams := *openapiclient.NewDeleteSessionParams("SessionId_example", "ProjectId_example") // DeleteSessionParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteSession(context.Background()).DeleteSessionParams(deleteSessionParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSessionParams** | [**DeleteSessionParams**](DeleteSessionParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSessionRun

> DeleteSessionRun(ctx).DeleteSessionRunParams(deleteSessionRunParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	deleteSessionRunParams := *openapiclient.NewDeleteSessionRunParams("SessionRunId_example", "ProjectId_example") // DeleteSessionRunParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteSessionRun(context.Background()).DeleteSessionRunParams(deleteSessionRunParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteSessionRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSessionRunParams** | [**DeleteSessionRunParams**](DeleteSessionRunParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSessionTest

> DeleteSessionTest(ctx).DeleteSessionTestRequest(deleteSessionTestRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	deleteSessionTestRequest := *openapiclient.NewDeleteSessionTestRequest("Cid_example", "ProjectId_example") // DeleteSessionTestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteSessionTest(context.Background()).DeleteSessionTestRequest(deleteSessionTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteSessionTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSessionTestRequest** | [**DeleteSessionTestRequest**](DeleteSessionTestRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSyntheticData

> DeleteSyntheticData(ctx).DeleteSyntheticDataParams(deleteSyntheticDataParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	deleteSyntheticDataParams := *openapiclient.NewDeleteSyntheticDataParams("ProjectId_example", "SyntheticDataId_example") // DeleteSyntheticDataParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteSyntheticData(context.Background()).DeleteSyntheticDataParams(deleteSyntheticDataParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteSyntheticData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSyntheticDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSyntheticDataParams** | [**DeleteSyntheticDataParams**](DeleteSyntheticDataParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeam

> DeleteTeam(ctx).DeleteTeamRequest(deleteTeamRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	deleteTeamRequest := *openapiclient.NewDeleteTeamRequest("Cid_example") // DeleteTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteTeam(context.Background()).DeleteTeamRequest(deleteTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteTeamRequest** | [**DeleteTeamRequest**](DeleteTeamRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserById

> DeleteUserById(ctx).DeleteUserByIdRequest(deleteUserByIdRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	deleteUserByIdRequest := *openapiclient.NewDeleteUserByIdRequest("UserId_example") // DeleteUserByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteUserById(context.Background()).DeleteUserByIdRequest(deleteUserByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteUserById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteUserByIdRequest** | [**DeleteUserByIdRequest**](DeleteUserByIdRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVersion

> DeleteVersion(ctx).DeleteVersionParams(deleteVersionParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	deleteVersionParams := *openapiclient.NewDeleteVersionParams("VersionId_example", "ProjectId_example") // DeleteVersionParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteVersion(context.Background()).DeleteVersionParams(deleteVersionParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVersionParams** | [**DeleteVersionParams**](DeleteVersionParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVisualizations

> DeleteVisualizations(ctx).DeleteVisualizationsParams(deleteVisualizationsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	deleteVisualizationsParams := *openapiclient.NewDeleteVisualizationsParams([]string{"VisualizationIdsToDelete_example"}, "ProjectId_example") // DeleteVisualizationsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteVisualizations(context.Background()).DeleteVisualizationsParams(deleteVisualizationsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteVisualizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVisualizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVisualizationsParams** | [**DeleteVisualizationsParams**](DeleteVisualizationsParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadProject

> DownloadProject(ctx, projectId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	projectId := "projectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DownloadProject(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Evaluate

> Job Evaluate(ctx).EvaluateParams(evaluateParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	evaluateParams := *openapiclient.NewEvaluateParams("VersionId_example", "ProjectId_example", float64(123), "Name_example", "Description_example", float64(123), "SessionId_example") // EvaluateParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Evaluate(context.Background()).EvaluateParams(evaluateParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Evaluate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Evaluate`: Job
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Evaluate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evaluateParams** | [**EvaluateParams**](EvaluateParams.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportProject

> ExportProjectResponse ExportProject(ctx).ExportProjectRequest(exportProjectRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	exportProjectRequest := *openapiclient.NewExportProjectRequest("ProjectId_example") // ExportProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ExportProject(context.Background()).ExportProjectRequest(exportProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ExportProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportProject`: ExportProjectResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ExportProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportProjectRequest** | [**ExportProjectRequest**](ExportProjectRequest.md) |  | 

### Return type

[**ExportProjectResponse**](ExportProjectResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtendTrial

> ExtendTrialResponse ExtendTrial(ctx).ExtendTrialParams(extendTrialParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	extendTrialParams := *openapiclient.NewExtendTrialParams("Token_example") // ExtendTrialParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ExtendTrial(context.Background()).ExtendTrialParams(extendTrialParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ExtendTrial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExtendTrial`: ExtendTrialResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ExtendTrial`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExtendTrialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extendTrialParams** | [**ExtendTrialParams**](ExtendTrialParams.md) |  | 

### Return type

[**ExtendTrialResponse**](ExtendTrialResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSimilar

> FetchSimilarResponse FetchSimilar(ctx).FetchSimilarRequestParams(fetchSimilarRequestParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	fetchSimilarRequestParams := *openapiclient.NewFetchSimilarRequestParams("SessionRunId_example", "ProjectId_example", float64(123), []openapiclient.SampleIdentity{*openapiclient.NewSampleIdentity(openapiclient.DataStateType("training"), *openapiclient.NewSampleIdType())}, float64(123), "Digest_example") // FetchSimilarRequestParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FetchSimilar(context.Background()).FetchSimilarRequestParams(fetchSimilarRequestParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FetchSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchSimilar`: FetchSimilarResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FetchSimilar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchSimilarRequestParams** | [**FetchSimilarRequestParams**](FetchSimilarRequestParams.md) |  | 

### Return type

[**FetchSimilarResponse**](FetchSimilarResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateInsights

> GenerateInsights(ctx).GenerateInsightsParams(generateInsightsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	generateInsightsParams := *openapiclient.NewGenerateInsightsParams("ProjectId_example", "SessionRunId_example") // GenerateInsightsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GenerateInsights(context.Background()).GenerateInsightsParams(generateInsightsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GenerateInsights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateInsightsParams** | [**GenerateInsightsParams**](GenerateInsightsParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateLabels

> GenerateLabels(ctx).GenerateLabelParams(generateLabelParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	generateLabelParams := *openapiclient.NewGenerateLabelParams("ProjectId_example", "SessionRunId_example", float64(123)) // GenerateLabelParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GenerateLabels(context.Background()).GenerateLabelParams(generateLabelParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GenerateLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateLabelParams** | [**GenerateLabelParams**](GenerateLabelParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateSyntheticData

> GenerateSyntheticData(ctx).GenerateSyntheticDataParams(generateSyntheticDataParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	generateSyntheticDataParams := *openapiclient.NewGenerateSyntheticDataParams("ProjectId_example", "SessionRunId_example", float64(123), []openapiclient.GenerateSyntheticDataParamsSourcesInner{*openapiclient.NewGenerateSyntheticDataParamsSourcesInner([]string{"MetadataKeys_example"}, []openapiclient.ESFilter{*openapiclient.NewESFilter(openapiclient.FilterOperatorType("between"), "Field_example", *openapiclient.NewESFilterValue())})}, []openapiclient.ESFilter{*openapiclient.NewESFilter(openapiclient.FilterOperatorType("between"), "Field_example", *openapiclient.NewESFilterValue())}) // GenerateSyntheticDataParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GenerateSyntheticData(context.Background()).GenerateSyntheticDataParams(generateSyntheticDataParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GenerateSyntheticData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSyntheticDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateSyntheticDataParams** | [**GenerateSyntheticDataParams**](GenerateSyntheticDataParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllProjectSessionTests

> []SessionTest GetAllProjectSessionTests(ctx).GetAllProjectSessionTestsRequest(getAllProjectSessionTestsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getAllProjectSessionTestsRequest := *openapiclient.NewGetAllProjectSessionTestsRequest("ProjectId_example") // GetAllProjectSessionTestsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetAllProjectSessionTests(context.Background()).GetAllProjectSessionTestsRequest(getAllProjectSessionTestsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAllProjectSessionTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllProjectSessionTests`: []SessionTest
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAllProjectSessionTests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProjectSessionTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getAllProjectSessionTestsRequest** | [**GetAllProjectSessionTestsRequest**](GetAllProjectSessionTestsRequest.md) |  | 

### Return type

[**[]SessionTest**](SessionTest.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSlimUserData

> GetTeamUsersResponse GetAllSlimUserData(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetAllSlimUserData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAllSlimUserData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSlimUserData`: GetTeamUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAllSlimUserData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSlimUserDataRequest struct via the builder pattern


### Return type

[**GetTeamUsersResponse**](GetTeamUsersResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKeyByCode

> RotateApiKeyResponse GetApiKeyByCode(ctx).GetApiKeyByCodeRequest(getApiKeyByCodeRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getApiKeyByCodeRequest := *openapiclient.NewGetApiKeyByCodeRequest("CodeVerifier_example") // GetApiKeyByCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetApiKeyByCode(context.Background()).GetApiKeyByCodeRequest(getApiKeyByCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetApiKeyByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiKeyByCode`: RotateApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetApiKeyByCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getApiKeyByCodeRequest** | [**GetApiKeyByCodeRequest**](GetApiKeyByCodeRequest.md) |  | 

### Return type

[**RotateApiKeyResponse**](RotateApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthProvider

> GetAuthProviderResponse GetAuthProvider(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetAuthProvider(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAuthProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthProvider`: GetAuthProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAuthProvider`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthProviderRequest struct via the builder pattern


### Return type

[**GetAuthProviderResponse**](GetAuthProviderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthStatus

> GetAuthStatusResponse GetAuthStatus(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetAuthStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAuthStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthStatus`: GetAuthStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAuthStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthStatusRequest struct via the builder pattern


### Return type

[**GetAuthStatusResponse**](GetAuthStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalancedAccuracy

> MultiChartsResponse GetBalancedAccuracy(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	body := ConfusionMatrixParams(987) // ConfusionMatrixParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetBalancedAccuracy(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetBalancedAccuracy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalancedAccuracy`: MultiChartsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetBalancedAccuracy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBalancedAccuracyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **ConfusionMatrixParams** |  | 

### Return type

[**MultiChartsResponse**](MultiChartsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCodeSnapshot

> GetCodeSnapshotResponse GetCodeSnapshot(ctx).GetCodeSnapshotParams(getCodeSnapshotParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getCodeSnapshotParams := *openapiclient.NewGetCodeSnapshotParams("ProjectId_example", "CodeSnapshotId_example") // GetCodeSnapshotParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetCodeSnapshot(context.Background()).GetCodeSnapshotParams(getCodeSnapshotParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCodeSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCodeSnapshot`: GetCodeSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCodeSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCodeSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCodeSnapshotParams** | [**GetCodeSnapshotParams**](GetCodeSnapshotParams.md) |  | 

### Return type

[**GetCodeSnapshotResponse**](GetCodeSnapshotResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCodeSnapshotUploadUrl

> CodeSnapshotUploadUrlResponse GetCodeSnapshotUploadUrl(ctx).GetCodeSnapshotUploadUrlParams(getCodeSnapshotUploadUrlParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getCodeSnapshotUploadUrlParams := *openapiclient.NewGetCodeSnapshotUploadUrlParams("ProjectId_example") // GetCodeSnapshotUploadUrlParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetCodeSnapshotUploadUrl(context.Background()).GetCodeSnapshotUploadUrlParams(getCodeSnapshotUploadUrlParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCodeSnapshotUploadUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCodeSnapshotUploadUrl`: CodeSnapshotUploadUrlResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCodeSnapshotUploadUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCodeSnapshotUploadUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCodeSnapshotUploadUrlParams** | [**GetCodeSnapshotUploadUrlParams**](GetCodeSnapshotUploadUrlParams.md) |  | 

### Return type

[**CodeSnapshotUploadUrlResponse**](CodeSnapshotUploadUrlResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfusionMatrixLabels

> ConfusionMatrixLabelsResponse GetConfusionMatrixLabels(ctx).GetConfusionMatrixLabels(getConfusionMatrixLabels).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getConfusionMatrixLabels := *openapiclient.NewGetConfusionMatrixLabels("ProjectId_example", []string{"SessionRunIds_example"}) // GetConfusionMatrixLabels | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetConfusionMatrixLabels(context.Background()).GetConfusionMatrixLabels(getConfusionMatrixLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetConfusionMatrixLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfusionMatrixLabels`: ConfusionMatrixLabelsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetConfusionMatrixLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfusionMatrixLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getConfusionMatrixLabels** | [**GetConfusionMatrixLabels**](GetConfusionMatrixLabels.md) |  | 

### Return type

[**ConfusionMatrixLabelsResponse**](ConfusionMatrixLabelsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfusionMatrixResultCombinations

> MultiChartsResponse GetConfusionMatrixResultCombinations(ctx).GetConfusionMatrixResultCombinationsParams(getConfusionMatrixResultCombinationsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getConfusionMatrixResultCombinationsParams := *openapiclient.NewGetConfusionMatrixResultCombinationsParams("ProjectId_example", []openapiclient.SessionRunToEpoch{*openapiclient.NewSessionRunToEpoch("SessionRunId_example", float64(123))}, *openapiclient.NewSplitAgg(NullableFloat64(123), openapiclient.OrderType("desc"), "Field_example", float64(123), "Distribution_example"), "CustomMetricName_example") // GetConfusionMatrixResultCombinationsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetConfusionMatrixResultCombinations(context.Background()).GetConfusionMatrixResultCombinationsParams(getConfusionMatrixResultCombinationsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetConfusionMatrixResultCombinations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfusionMatrixResultCombinations`: MultiChartsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetConfusionMatrixResultCombinations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfusionMatrixResultCombinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getConfusionMatrixResultCombinationsParams** | [**GetConfusionMatrixResultCombinationsParams**](GetConfusionMatrixResultCombinationsParams.md) |  | 

### Return type

[**MultiChartsResponse**](MultiChartsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfusionMatrixTable

> MultiChartsResponse GetConfusionMatrixTable(ctx).ConfusionMatrixTableParams(confusionMatrixTableParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	confusionMatrixTableParams := *openapiclient.NewConfusionMatrixTableParams([]openapiclient.SessionRunToEpoch{*openapiclient.NewSessionRunToEpoch("SessionRunId_example", float64(123))}, "ProjectId_example", "CustomMetricName_example", false) // ConfusionMatrixTableParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetConfusionMatrixTable(context.Background()).ConfusionMatrixTableParams(confusionMatrixTableParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetConfusionMatrixTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfusionMatrixTable`: MultiChartsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetConfusionMatrixTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfusionMatrixTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **confusionMatrixTableParams** | [**ConfusionMatrixTableParams**](ConfusionMatrixTableParams.md) |  | 

### Return type

[**MultiChartsResponse**](MultiChartsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfusionMetricNames

> ConfusionMetricNamesResponse GetConfusionMetricNames(ctx).ConfusionMetricNamesParams(confusionMetricNamesParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	confusionMetricNamesParams := *openapiclient.NewConfusionMetricNamesParams([]string{"SessionRunIds_example"}, "ProjectId_example") // ConfusionMetricNamesParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetConfusionMetricNames(context.Background()).ConfusionMetricNamesParams(confusionMetricNamesParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetConfusionMetricNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfusionMetricNames`: ConfusionMetricNamesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetConfusionMetricNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfusionMetricNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **confusionMetricNamesParams** | [**ConfusionMetricNamesParams**](ConfusionMetricNamesParams.md) |  | 

### Return type

[**ConfusionMetricNamesResponse**](ConfusionMetricNamesResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentProjectVersion

> GetCurrentProjectVersionResponse GetCurrentProjectVersion(ctx).GetCurrentProjectVersionParams(getCurrentProjectVersionParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getCurrentProjectVersionParams := *openapiclient.NewGetCurrentProjectVersionParams("ProjectId_example") // GetCurrentProjectVersionParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetCurrentProjectVersion(context.Background()).GetCurrentProjectVersionParams(getCurrentProjectVersionParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCurrentProjectVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentProjectVersion`: GetCurrentProjectVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCurrentProjectVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentProjectVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCurrentProjectVersionParams** | [**GetCurrentProjectVersionParams**](GetCurrentProjectVersionParams.md) |  | 

### Return type

[**GetCurrentProjectVersionResponse**](GetCurrentProjectVersionResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboard

> GetDashboardResponse GetDashboard(ctx).GetDashboardParams(getDashboardParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getDashboardParams := *openapiclient.NewGetDashboardParams("DashboardId_example", "ProjectId_example") // GetDashboardParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetDashboard(context.Background()).GetDashboardParams(getDashboardParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboard`: GetDashboardResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getDashboardParams** | [**GetDashboardParams**](GetDashboardParams.md) |  | 

### Return type

[**GetDashboardResponse**](GetDashboardResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashletFields

> GetDashletFieldsResponse GetDashletFields(ctx).GetDashletFieldsParams(getDashletFieldsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getDashletFieldsParams := *openapiclient.NewGetDashletFieldsParams("ProjectId_example", []string{"SessionRunIds_example"}) // GetDashletFieldsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetDashletFields(context.Background()).GetDashletFieldsParams(getDashletFieldsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetDashletFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashletFields`: GetDashletFieldsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetDashletFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDashletFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getDashletFieldsParams** | [**GetDashletFieldsParams**](GetDashletFieldsParams.md) |  | 

### Return type

[**GetDashletFieldsResponse**](GetDashletFieldsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDownloadSignedUrl

> GetDownloadSignedUrlResponse GetDownloadSignedUrl(ctx).GetDownloadSignedUrlParams(getDownloadSignedUrlParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getDownloadSignedUrlParams := *openapiclient.NewGetDownloadSignedUrlParams("FileName_example") // GetDownloadSignedUrlParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetDownloadSignedUrl(context.Background()).GetDownloadSignedUrlParams(getDownloadSignedUrlParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetDownloadSignedUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDownloadSignedUrl`: GetDownloadSignedUrlResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetDownloadSignedUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadSignedUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getDownloadSignedUrlParams** | [**GetDownloadSignedUrlParams**](GetDownloadSignedUrlParams.md) |  | 

### Return type

[**GetDownloadSignedUrlResponse**](GetDownloadSignedUrlResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEngineSettings

> SettingsAndValuesWrapper GetEngineSettings(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetEngineSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEngineSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEngineSettings`: SettingsAndValuesWrapper
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEngineSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEngineSettingsRequest struct via the builder pattern


### Return type

[**SettingsAndValuesWrapper**](SettingsAndValuesWrapper.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentInfo

> GetEnvironmentInfoResponse GetEnvironmentInfo(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetEnvironmentInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEnvironmentInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentInfo`: GetEnvironmentInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEnvironmentInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentInfoRequest struct via the builder pattern


### Return type

[**GetEnvironmentInfoResponse**](GetEnvironmentInfoResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExportedSessionJobs

> GetExportedSessionRunJobsResponse GetExportedSessionJobs(ctx).GetExportedSessionRunJobsParams(getExportedSessionRunJobsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getExportedSessionRunJobsParams := *openapiclient.NewGetExportedSessionRunJobsParams("SessionId_example", "ProjectId_example") // GetExportedSessionRunJobsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetExportedSessionJobs(context.Background()).GetExportedSessionRunJobsParams(getExportedSessionRunJobsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetExportedSessionJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExportedSessionJobs`: GetExportedSessionRunJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetExportedSessionJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExportedSessionJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getExportedSessionRunJobsParams** | [**GetExportedSessionRunJobsParams**](GetExportedSessionRunJobsParams.md) |  | 

### Return type

[**GetExportedSessionRunJobsResponse**](GetExportedSessionRunJobsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetF1Score

> MultiChartsResponse GetF1Score(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	body := ConfusionMatrixParams(987) // ConfusionMatrixParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetF1Score(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetF1Score``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetF1Score`: MultiChartsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetF1Score`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetF1ScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **ConfusionMatrixParams** |  | 

### Return type

[**MultiChartsResponse**](MultiChartsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFetchSimilarStatus

> FetchSimilarResponse GetFetchSimilarStatus(ctx).FetchSimilarRequestParams(fetchSimilarRequestParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	fetchSimilarRequestParams := *openapiclient.NewFetchSimilarRequestParams("SessionRunId_example", "ProjectId_example", float64(123), []openapiclient.SampleIdentity{*openapiclient.NewSampleIdentity(openapiclient.DataStateType("training"), *openapiclient.NewSampleIdType())}, float64(123), "Digest_example") // FetchSimilarRequestParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetFetchSimilarStatus(context.Background()).FetchSimilarRequestParams(fetchSimilarRequestParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetFetchSimilarStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFetchSimilarStatus`: FetchSimilarResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetFetchSimilarStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFetchSimilarStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchSimilarRequestParams** | [**FetchSimilarRequestParams**](FetchSimilarRequestParams.md) |  | 

### Return type

[**FetchSimilarResponse**](FetchSimilarResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFieldsValues

> GetFieldsValuesResponse GetFieldsValues(ctx).GetFieldsValuesRequest(getFieldsValuesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getFieldsValuesRequest := *openapiclient.NewGetFieldsValuesRequest([]openapiclient.ESFilter{*openapiclient.NewESFilter(openapiclient.FilterOperatorType("between"), "Field_example", *openapiclient.NewESFilterValue())}, []string{"SessionRunIds_example"}, []openapiclient.QueryFieldValues{*openapiclient.NewQueryFieldValues("Type_example", "Field_example")}, "ProjectId_example") // GetFieldsValuesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetFieldsValues(context.Background()).GetFieldsValuesRequest(getFieldsValuesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetFieldsValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFieldsValues`: GetFieldsValuesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetFieldsValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldsValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getFieldsValuesRequest** | [**GetFieldsValuesRequest**](GetFieldsValuesRequest.md) |  | 

### Return type

[**GetFieldsValuesResponse**](GetFieldsValuesResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGeneratedLabels

> GeneratedLabelsResponse GetGeneratedLabels(ctx).GetGeneratedLabelsParams(getGeneratedLabelsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getGeneratedLabelsParams := *openapiclient.NewGetGeneratedLabelsParams("ProjectId_example") // GetGeneratedLabelsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGeneratedLabels(context.Background()).GetGeneratedLabelsParams(getGeneratedLabelsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGeneratedLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGeneratedLabels`: GeneratedLabelsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGeneratedLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGeneratedLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getGeneratedLabelsParams** | [**GetGeneratedLabelsParams**](GetGeneratedLabelsParams.md) |  | 

### Return type

[**GeneratedLabelsResponse**](GeneratedLabelsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGenericBaseImageTypes

> GetGenericBaseImageTypesResponse GetGenericBaseImageTypes(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGenericBaseImageTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGenericBaseImageTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenericBaseImageTypes`: GetGenericBaseImageTypesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGenericBaseImageTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenericBaseImageTypesRequest struct via the builder pattern


### Return type

[**GetGenericBaseImageTypesResponse**](GetGenericBaseImageTypesResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHeatmapChart

> MultiChartsResponse GetHeatmapChart(ctx).HeatmapChartsParams(heatmapChartsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	heatmapChartsParams := *openapiclient.NewHeatmapChartsParams("ProjectId_example", *openapiclient.NewSplitAgg(NullableFloat64(123), openapiclient.OrderType("desc"), "Field_example", float64(123), "Distribution_example"), *openapiclient.NewSplitAgg(NullableFloat64(123), openapiclient.OrderType("desc"), "Field_example", float64(123), "Distribution_example"), *openapiclient.NewAggregations("Field_example", openapiclient.AggregationMethod("Average")), []openapiclient.SessionRunToEpoch{*openapiclient.NewSessionRunToEpoch("SessionRunId_example", float64(123))}, false) // HeatmapChartsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetHeatmapChart(context.Background()).HeatmapChartsParams(heatmapChartsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetHeatmapChart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHeatmapChart`: MultiChartsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetHeatmapChart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHeatmapChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **heatmapChartsParams** | [**HeatmapChartsParams**](HeatmapChartsParams.md) |  | 

### Return type

[**MultiChartsResponse**](MultiChartsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsights

> GetInsightsResponse GetInsights(ctx).GetInsightsParams(getInsightsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getInsightsParams := *openapiclient.NewGetInsightsParams("SessionRunId_example", "ProjectId_example") // GetInsightsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetInsights(context.Background()).GetInsightsParams(getInsightsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetInsights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInsights`: GetInsightsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getInsightsParams** | [**GetInsightsParams**](GetInsightsParams.md) |  | 

### Return type

[**GetInsightsResponse**](GetInsightsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueFileUploadSignedUrl

> IssueFileUploadSignedUrl GetIssueFileUploadSignedUrl(ctx).GetIssueFileUploadSignedUrl(getIssueFileUploadSignedUrl).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getIssueFileUploadSignedUrl := *openapiclient.NewGetIssueFileUploadSignedUrl("FileName_example", "ProjectId_example", "IssueId_example") // GetIssueFileUploadSignedUrl | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetIssueFileUploadSignedUrl(context.Background()).GetIssueFileUploadSignedUrl(getIssueFileUploadSignedUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetIssueFileUploadSignedUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueFileUploadSignedUrl`: IssueFileUploadSignedUrl
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetIssueFileUploadSignedUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueFileUploadSignedUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getIssueFileUploadSignedUrl** | [**GetIssueFileUploadSignedUrl**](GetIssueFileUploadSignedUrl.md) |  | 

### Return type

[**IssueFileUploadSignedUrl**](IssueFileUploadSignedUrl.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobLogs

> GetJobLogsResponse GetJobLogs(ctx).GetJobLogsParams(getJobLogsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getJobLogsParams := *openapiclient.NewGetJobLogsParams("JobId_example") // GetJobLogsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetJobLogs(context.Background()).GetJobLogsParams(getJobLogsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetJobLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobLogs`: GetJobLogsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetJobLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getJobLogsParams** | [**GetJobLogsParams**](GetJobLogsParams.md) |  | 

### Return type

[**GetJobLogsResponse**](GetJobLogsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestExportedProject

> LatestExportedProject GetLatestExportedProject(ctx).GetLatestExportedProjectParams(getLatestExportedProjectParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getLatestExportedProjectParams := *openapiclient.NewGetLatestExportedProjectParams("ProjectId_example") // GetLatestExportedProjectParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetLatestExportedProject(context.Background()).GetLatestExportedProjectParams(getLatestExportedProjectParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetLatestExportedProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestExportedProject`: LatestExportedProject
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetLatestExportedProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestExportedProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getLatestExportedProjectParams** | [**GetLatestExportedProjectParams**](GetLatestExportedProjectParams.md) |  | 

### Return type

[**LatestExportedProject**](LatestExportedProject.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMachineTypes

> GetMachineTypesResponse GetMachineTypes(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetMachineTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMachineTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMachineTypes`: GetMachineTypesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMachineTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMachineTypesRequest struct via the builder pattern


### Return type

[**GetMachineTypesResponse**](GetMachineTypesResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaxActiveUsers

> GetMaxActiveUsersResponse GetMaxActiveUsers(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetMaxActiveUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMaxActiveUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMaxActiveUsers`: GetMaxActiveUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMaxActiveUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxActiveUsersRequest struct via the builder pattern


### Return type

[**GetMaxActiveUsersResponse**](GetMaxActiveUsersResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeanAveragePrecision

> MultiChartsResponse GetMeanAveragePrecision(ctx).ConfusionMatrixParams(confusionMatrixParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	confusionMatrixParams := *openapiclient.NewConfusionMatrixParams("ProjectId_example", []openapiclient.SessionRunToEpoch{*openapiclient.NewSessionRunToEpoch("SessionRunId_example", float64(123))}, *openapiclient.NewSplitAgg(NullableFloat64(123), openapiclient.OrderType("desc"), "Field_example", float64(123), "Distribution_example"), "CustomMetricName_example") // ConfusionMatrixParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetMeanAveragePrecision(context.Background()).ConfusionMatrixParams(confusionMatrixParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMeanAveragePrecision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeanAveragePrecision`: MultiChartsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMeanAveragePrecision`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMeanAveragePrecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **confusionMatrixParams** | [**ConfusionMatrixParams**](ConfusionMatrixParams.md) |  | 

### Return type

[**MultiChartsResponse**](MultiChartsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotifications

> GetNotificationsResponse GetNotifications(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetNotifications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotifications`: GetNotificationsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetNotifications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsRequest struct via the builder pattern


### Return type

[**GetNotificationsResponse**](GetNotificationsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsByFilter

> GetNotificationsByFilterResponse GetNotificationsByFilter(ctx).GetNotificationsByFilterParams(getNotificationsByFilterParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getNotificationsByFilterParams := *openapiclient.NewGetNotificationsByFilterParams() // GetNotificationsByFilterParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetNotificationsByFilter(context.Background()).GetNotificationsByFilterParams(getNotificationsByFilterParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetNotificationsByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationsByFilter`: GetNotificationsByFilterResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetNotificationsByFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getNotificationsByFilterParams** | [**GetNotificationsByFilterParams**](GetNotificationsByFilterParams.md) |  | 

### Return type

[**GetNotificationsByFilterResponse**](GetNotificationsByFilterResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPopulationExplorationStatus

> PopulationExplorationResponse GetPopulationExplorationStatus(ctx).PopulationExplorationParams(populationExplorationParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	populationExplorationParams := *openapiclient.NewPopulationExplorationParams("SessionRunId_example", "ProjectId_example", float64(123), float64(123), "Digest_example", float64(123), []string{"BalanceBy_example"}, false, openapiclient.ReductionAlgorithm("TSNE")) // PopulationExplorationParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPopulationExplorationStatus(context.Background()).PopulationExplorationParams(populationExplorationParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPopulationExplorationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPopulationExplorationStatus`: PopulationExplorationResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPopulationExplorationStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPopulationExplorationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **populationExplorationParams** | [**PopulationExplorationParams**](PopulationExplorationParams.md) |  | 

### Return type

[**PopulationExplorationResponse**](PopulationExplorationResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrCurve

> MultiChartsResponse GetPrCurve(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	body := MultiThresholdConfusionMatrixParams(987) // MultiThresholdConfusionMatrixParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPrCurve(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPrCurve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrCurve`: MultiChartsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPrCurve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrCurveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **MultiThresholdConfusionMatrixParams** |  | 

### Return type

[**MultiChartsResponse**](MultiChartsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrecisionScore

> MultiChartsResponse GetPrecisionScore(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	body := ConfusionMatrixParams(987) // ConfusionMatrixParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPrecisionScore(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPrecisionScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrecisionScore`: MultiChartsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPrecisionScore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrecisionScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **ConfusionMatrixParams** |  | 

### Return type

[**MultiChartsResponse**](MultiChartsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectDashboards

> GetProjectDashboardsResponse GetProjectDashboards(ctx).GetProjectDashboardsParams(getProjectDashboardsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getProjectDashboardsParams := *openapiclient.NewGetProjectDashboardsParams("ProjectId_example") // GetProjectDashboardsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetProjectDashboards(context.Background()).GetProjectDashboardsParams(getProjectDashboardsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetProjectDashboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectDashboards`: GetProjectDashboardsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetProjectDashboards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectDashboardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getProjectDashboardsParams** | [**GetProjectDashboardsParams**](GetProjectDashboardsParams.md) |  | 

### Return type

[**GetProjectDashboardsResponse**](GetProjectDashboardsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectIssues

> GetProjectIssuesResponse GetProjectIssues(ctx).GetProjectIssuesParams(getProjectIssuesParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getProjectIssuesParams := *openapiclient.NewGetProjectIssuesParams("ProjectId_example") // GetProjectIssuesParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetProjectIssues(context.Background()).GetProjectIssuesParams(getProjectIssuesParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetProjectIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectIssues`: GetProjectIssuesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetProjectIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getProjectIssuesParams** | [**GetProjectIssuesParams**](GetProjectIssuesParams.md) |  | 

### Return type

[**GetProjectIssuesResponse**](GetProjectIssuesResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectSlimVersions

> GetProjectSlimVersionsResponse GetProjectSlimVersions(ctx).GetProjectVersionsParams(getProjectVersionsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getProjectVersionsParams := *openapiclient.NewGetProjectVersionsParams("ProjectId_example") // GetProjectVersionsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetProjectSlimVersions(context.Background()).GetProjectVersionsParams(getProjectVersionsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetProjectSlimVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectSlimVersions`: GetProjectSlimVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetProjectSlimVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectSlimVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getProjectVersionsParams** | [**GetProjectVersionsParams**](GetProjectVersionsParams.md) |  | 

### Return type

[**GetProjectSlimVersionsResponse**](GetProjectSlimVersionsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjects

> GetProjectsResponse GetProjects(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjects`: GetProjectsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsRequest struct via the builder pattern


### Return type

[**GetProjectsResponse**](GetProjectsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecallScore

> MultiChartsResponse GetRecallScore(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	body := ConfusionMatrixParams(987) // ConfusionMatrixParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRecallScore(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRecallScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecallScore`: MultiChartsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRecallScore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecallScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **ConfusionMatrixParams** |  | 

### Return type

[**MultiChartsResponse**](MultiChartsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecentTeamSessions

> RecentSessionsResponse GetRecentTeamSessions(ctx).RecentTeamSessionsRequestParams(recentTeamSessionsRequestParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	recentTeamSessionsRequestParams := *openapiclient.NewRecentTeamSessionsRequestParams(float64(123), "ProjectId_example") // RecentTeamSessionsRequestParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRecentTeamSessions(context.Background()).RecentTeamSessionsRequestParams(recentTeamSessionsRequestParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRecentTeamSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecentTeamSessions`: RecentSessionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRecentTeamSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecentTeamSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recentTeamSessionsRequestParams** | [**RecentTeamSessionsRequestParams**](RecentTeamSessionsRequestParams.md) |  | 

### Return type

[**RecentSessionsResponse**](RecentSessionsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoc

> MultiChartsResponse GetRoc(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	body := RocConfusionMatrixParams(987) // RocConfusionMatrixParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRoc(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRoc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoc`: MultiChartsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRoc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRocRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **RocConfusionMatrixParams** |  | 

### Return type

[**MultiChartsResponse**](MultiChartsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSampleCollections

> GetSampleCollectionsResponse GetSampleCollections(ctx).GetSampleCollectionsParams(getSampleCollectionsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getSampleCollectionsParams := *openapiclient.NewGetSampleCollectionsParams("ProjectId_example") // GetSampleCollectionsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSampleCollections(context.Background()).GetSampleCollectionsParams(getSampleCollectionsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSampleCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSampleCollections`: GetSampleCollectionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSampleCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSampleCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSampleCollectionsParams** | [**GetSampleCollectionsParams**](GetSampleCollectionsParams.md) |  | 

### Return type

[**GetSampleCollectionsResponse**](GetSampleCollectionsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSampleVisualizationsPath

> GetSampleVisualizationsPathsResponse GetSampleVisualizationsPath(ctx).GetSampleVisualizationsPathsParams(getSampleVisualizationsPathsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getSampleVisualizationsPathsParams := *openapiclient.NewGetSampleVisualizationsPathsParams("ScatterSampleVisualizationsPrefix_example", "FileNameMatch_example", "SampleId_example") // GetSampleVisualizationsPathsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSampleVisualizationsPath(context.Background()).GetSampleVisualizationsPathsParams(getSampleVisualizationsPathsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSampleVisualizationsPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSampleVisualizationsPath`: GetSampleVisualizationsPathsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSampleVisualizationsPath`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSampleVisualizationsPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSampleVisualizationsPathsParams** | [**GetSampleVisualizationsPathsParams**](GetSampleVisualizationsPathsParams.md) |  | 

### Return type

[**GetSampleVisualizationsPathsResponse**](GetSampleVisualizationsPathsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScatterSampleVisualizations

> GetScatterSampleVisualizationsResponse GetScatterSampleVisualizations(ctx).GetScatterSampleVisualizationsParams(getScatterSampleVisualizationsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getScatterSampleVisualizationsParams := *openapiclient.NewGetScatterSampleVisualizationsParams("SessionRunId_example", "ProjectId_example", float64(123)) // GetScatterSampleVisualizationsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetScatterSampleVisualizations(context.Background()).GetScatterSampleVisualizationsParams(getScatterSampleVisualizationsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetScatterSampleVisualizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScatterSampleVisualizations`: GetScatterSampleVisualizationsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetScatterSampleVisualizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetScatterSampleVisualizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getScatterSampleVisualizationsParams** | [**GetScatterSampleVisualizationsParams**](GetScatterSampleVisualizationsParams.md) |  | 

### Return type

[**GetScatterSampleVisualizationsResponse**](GetScatterSampleVisualizationsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretManagerList

> GetSecretManagerListResponse GetSecretManagerList(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSecretManagerList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSecretManagerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecretManagerList`: GetSecretManagerListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSecretManagerList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretManagerListRequest struct via the builder pattern


### Return type

[**GetSecretManagerListResponse**](GetSecretManagerListResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionRunsEvaluate

> GetSessionRunsEvaluateResponse GetSessionRunsEvaluate(ctx).GetSessionRunsEvaluateParams(getSessionRunsEvaluateParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getSessionRunsEvaluateParams := *openapiclient.NewGetSessionRunsEvaluateParams("ProjectId_example") // GetSessionRunsEvaluateParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSessionRunsEvaluate(context.Background()).GetSessionRunsEvaluateParams(getSessionRunsEvaluateParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSessionRunsEvaluate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionRunsEvaluate`: GetSessionRunsEvaluateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSessionRunsEvaluate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionRunsEvaluateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSessionRunsEvaluateParams** | [**GetSessionRunsEvaluateParams**](GetSessionRunsEvaluateParams.md) |  | 

### Return type

[**GetSessionRunsEvaluateResponse**](GetSessionRunsEvaluateResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionRunsVisualizations

> GetSessionRunsVisualizationsResponse GetSessionRunsVisualizations(ctx).GetSessionRunsVisualizationsParams(getSessionRunsVisualizationsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getSessionRunsVisualizationsParams := *openapiclient.NewGetSessionRunsVisualizationsParams([]string{"SessionRunIds_example"}, "ProjectId_example") // GetSessionRunsVisualizationsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSessionRunsVisualizations(context.Background()).GetSessionRunsVisualizationsParams(getSessionRunsVisualizationsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSessionRunsVisualizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionRunsVisualizations`: GetSessionRunsVisualizationsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSessionRunsVisualizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionRunsVisualizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSessionRunsVisualizationsParams** | [**GetSessionRunsVisualizationsParams**](GetSessionRunsVisualizationsParams.md) |  | 

### Return type

[**GetSessionRunsVisualizationsResponse**](GetSessionRunsVisualizationsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionTestResult

> []AllSessionsTestResults GetSessionTestResult(ctx).GetSessionTestResultsRequest(getSessionTestResultsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getSessionTestResultsRequest := *openapiclient.NewGetSessionTestResultsRequest("ProjectId_example", []openapiclient.SessionTestData{*openapiclient.NewSessionTestData("SessionRunId_example", "ProjectId_example", float64(123))}) // GetSessionTestResultsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSessionTestResult(context.Background()).GetSessionTestResultsRequest(getSessionTestResultsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSessionTestResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionTestResult`: []AllSessionsTestResults
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSessionTestResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionTestResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSessionTestResultsRequest** | [**GetSessionTestResultsRequest**](GetSessionTestResultsRequest.md) |  | 

### Return type

[**[]AllSessionsTestResults**](AllSessionsTestResults.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionsByHash

> SessionsResponse GetSessionsByHash(ctx).SessionHashRequestParams(sessionHashRequestParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	sessionHashRequestParams := *openapiclient.NewSessionHashRequestParams("Hash_example", "ProjectId_example") // SessionHashRequestParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSessionsByHash(context.Background()).SessionHashRequestParams(sessionHashRequestParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSessionsByHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionsByHash`: SessionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSessionsByHash`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionsByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionHashRequestParams** | [**SessionHashRequestParams**](SessionHashRequestParams.md) |  | 

### Return type

[**SessionsResponse**](SessionsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionsByVersionId

> SessionsResponse GetSessionsByVersionId(ctx).SessionVersionIdRequestParams(sessionVersionIdRequestParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	sessionVersionIdRequestParams := *openapiclient.NewSessionVersionIdRequestParams("VersionId_example", "ProjectId_example") // SessionVersionIdRequestParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSessionsByVersionId(context.Background()).SessionVersionIdRequestParams(sessionVersionIdRequestParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSessionsByVersionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionsByVersionId`: SessionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSessionsByVersionId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionsByVersionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionVersionIdRequestParams** | [**SessionVersionIdRequestParams**](SessionVersionIdRequestParams.md) |  | 

### Return type

[**SessionsResponse**](SessionsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionsEpochs

> GetSessionEpochsResponse GetSessionsEpochs(ctx).GetSessionsEpochsRequest(getSessionsEpochsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getSessionsEpochsRequest := *openapiclient.NewGetSessionsEpochsRequest("ProjectId_example", []string{"SessionIds_example"}) // GetSessionsEpochsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSessionsEpochs(context.Background()).GetSessionsEpochsRequest(getSessionsEpochsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSessionsEpochs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionsEpochs`: GetSessionEpochsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSessionsEpochs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionsEpochsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSessionsEpochsRequest** | [**GetSessionsEpochsRequest**](GetSessionsEpochsRequest.md) |  | 

### Return type

[**GetSessionEpochsResponse**](GetSessionEpochsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignedUrl

> ExternalImportModelStorage GetSignedUrl(ctx).GetSignedUrlParams(getSignedUrlParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getSignedUrlParams := *openapiclient.NewGetSignedUrlParams("FileName_example", float64(123), openapiclient.HttpMethods("GET")) // GetSignedUrlParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSignedUrl(context.Background()).GetSignedUrlParams(getSignedUrlParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSignedUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSignedUrl`: ExternalImportModelStorage
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSignedUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignedUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSignedUrlParams** | [**GetSignedUrlParams**](GetSignedUrlParams.md) |  | 

### Return type

[**ExternalImportModelStorage**](ExternalImportModelStorage.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleIssue

> Issue GetSingleIssue(ctx).GetSingleIssueParams(getSingleIssueParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getSingleIssueParams := *openapiclient.NewGetSingleIssueParams("Cid_example", "ProjectId_example") // GetSingleIssueParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSingleIssue(context.Background()).GetSingleIssueParams(getSingleIssueParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSingleIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSingleIssue`: Issue
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSingleIssue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSingleIssueParams** | [**GetSingleIssueParams**](GetSingleIssueParams.md) |  | 

### Return type

[**Issue**](Issue.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleSessionTest

> SessionTest GetSingleSessionTest(ctx).GetSingleSessionTestRequest(getSingleSessionTestRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getSingleSessionTestRequest := *openapiclient.NewGetSingleSessionTestRequest("Cid_example", "ProjectId_example") // GetSingleSessionTestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSingleSessionTest(context.Background()).GetSingleSessionTestRequest(getSingleSessionTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSingleSessionTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSingleSessionTest`: SessionTest
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSingleSessionTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleSessionTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSingleSessionTestRequest** | [**GetSingleSessionTestRequest**](GetSingleSessionTestRequest.md) |  | 

### Return type

[**SessionTest**](SessionTest.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlimJobs

> GetSlimJobsResponse GetSlimJobs(ctx).GetJobsFilterParams(getJobsFilterParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getJobsFilterParams := *openapiclient.NewGetJobsFilterParams() // GetJobsFilterParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSlimJobs(context.Background()).GetJobsFilterParams(getJobsFilterParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSlimJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSlimJobs`: GetSlimJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSlimJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSlimJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getJobsFilterParams** | [**GetJobsFilterParams**](GetJobsFilterParams.md) |  | 

### Return type

[**GetSlimJobsResponse**](GetSlimJobsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlimVisualization

> GetSlimVisualizationResponse GetSlimVisualization(ctx).GetSlimVisualizationParams(getSlimVisualizationParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getSlimVisualizationParams := *openapiclient.NewGetSlimVisualizationParams("VisualizationId_example", "ProjectId_example") // GetSlimVisualizationParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSlimVisualization(context.Background()).GetSlimVisualizationParams(getSlimVisualizationParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSlimVisualization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSlimVisualization`: GetSlimVisualizationResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSlimVisualization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSlimVisualizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSlimVisualizationParams** | [**GetSlimVisualizationParams**](GetSlimVisualizationParams.md) |  | 

### Return type

[**GetSlimVisualizationResponse**](GetSlimVisualizationResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetState

> GetStateResponse GetState(ctx).GetStateParams(getStateParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getStateParams := *openapiclient.NewGetStateParams("ProjectId_example", "Digest_example") // GetStateParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetState(context.Background()).GetStateParams(getStateParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetState`: GetStateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getStateParams** | [**GetStateParams**](GetStateParams.md) |  | 

### Return type

[**GetStateResponse**](GetStateResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatistics

> GetStatisticsResponse GetStatistics(ctx).GetStatisticsParams(getStatisticsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getStatisticsParams := *openapiclient.NewGetStatisticsParams("ProjectId_example") // GetStatisticsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetStatistics(context.Background()).GetStatisticsParams(getStatisticsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatistics`: GetStatisticsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetStatistics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getStatisticsParams** | [**GetStatisticsParams**](GetStatisticsParams.md) |  | 

### Return type

[**GetStatisticsResponse**](GetStatisticsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyntheticData

> SyntheticDataResponse GetSyntheticData(ctx).GetSyntheticDataParams(getSyntheticDataParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getSyntheticDataParams := *openapiclient.NewGetSyntheticDataParams("ProjectId_example") // GetSyntheticDataParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSyntheticData(context.Background()).GetSyntheticDataParams(getSyntheticDataParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSyntheticData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyntheticData`: SyntheticDataResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSyntheticData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSyntheticDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSyntheticDataParams** | [**GetSyntheticDataParams**](GetSyntheticDataParams.md) |  | 

### Return type

[**SyntheticDataResponse**](SyntheticDataResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTableChart

> MultiChartsResponse GetTableChart(ctx).GenericDataQueryParams(genericDataQueryParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	genericDataQueryParams := *openapiclient.NewGenericDataQueryParams("ProjectId_example", []openapiclient.SessionRunToEpoch{*openapiclient.NewSessionRunToEpoch("SessionRunId_example", float64(123))}, false, []openapiclient.Aggregations{*openapiclient.NewAggregations("Field_example", openapiclient.AggregationMethod("Average"))}, []openapiclient.SplitAgg{*openapiclient.NewSplitAgg(NullableFloat64(123), openapiclient.OrderType("desc"), "Field_example", float64(123), "Distribution_example")}) // GenericDataQueryParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTableChart(context.Background()).GenericDataQueryParams(genericDataQueryParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTableChart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTableChart`: MultiChartsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTableChart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTableChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **genericDataQueryParams** | [**GenericDataQueryParams**](GenericDataQueryParams.md) |  | 

### Return type

[**MultiChartsResponse**](MultiChartsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamJobs

> GetJobsResponse GetTeamJobs(ctx).GetJobsFilterParams(getJobsFilterParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getJobsFilterParams := *openapiclient.NewGetJobsFilterParams() // GetJobsFilterParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTeamJobs(context.Background()).GetJobsFilterParams(getJobsFilterParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTeamJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamJobs`: GetJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTeamJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getJobsFilterParams** | [**GetJobsFilterParams**](GetJobsFilterParams.md) |  | 

### Return type

[**GetJobsResponse**](GetJobsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamSlimUserData

> GetTeamUsersResponse GetTeamSlimUserData(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTeamSlimUserData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTeamSlimUserData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamSlimUserData`: GetTeamUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTeamSlimUserData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamSlimUserDataRequest struct via the builder pattern


### Return type

[**GetTeamUsersResponse**](GetTeamUsersResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeams

> GetTeamsResponse GetTeams(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTeams(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeams`: GetTeamsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTeams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsRequest struct via the builder pattern


### Return type

[**GetTeamsResponse**](GetTeamsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUploadModelSignedUrl

> ExternalImportModelStorage GetUploadModelSignedUrl(ctx).GetUploadModelSignedUrlRequest(getUploadModelSignedUrlRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getUploadModelSignedUrlRequest := *openapiclient.NewGetUploadModelSignedUrlRequest(float64(123), "ExperimentId_example", openapiclient.UploadModelFileType("onnx"), "VersionId_example", "ProjectId_example") // GetUploadModelSignedUrlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetUploadModelSignedUrl(context.Background()).GetUploadModelSignedUrlRequest(getUploadModelSignedUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetUploadModelSignedUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUploadModelSignedUrl`: ExternalImportModelStorage
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetUploadModelSignedUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUploadModelSignedUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getUploadModelSignedUrlRequest** | [**GetUploadModelSignedUrlRequest**](GetUploadModelSignedUrlRequest.md) |  | 

### Return type

[**ExternalImportModelStorage**](ExternalImportModelStorage.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUploadSignedUrl

> ExternalImportModelStorage GetUploadSignedUrl(ctx).GetUploadSignedUrlParams(getUploadSignedUrlParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getUploadSignedUrlParams := *openapiclient.NewGetUploadSignedUrlParams("FileName_example") // GetUploadSignedUrlParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetUploadSignedUrl(context.Background()).GetUploadSignedUrlParams(getUploadSignedUrlParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetUploadSignedUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUploadSignedUrl`: ExternalImportModelStorage
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetUploadSignedUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUploadSignedUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getUploadSignedUrlParams** | [**GetUploadSignedUrlParams**](GetUploadSignedUrlParams.md) |  | 

### Return type

[**ExternalImportModelStorage**](ExternalImportModelStorage.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVisualization

> Visualization GetVisualization(ctx).GetSlimVisualizationParams(getSlimVisualizationParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	getSlimVisualizationParams := *openapiclient.NewGetSlimVisualizationParams("VisualizationId_example", "ProjectId_example") // GetSlimVisualizationParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetVisualization(context.Background()).GetSlimVisualizationParams(getSlimVisualizationParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetVisualization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVisualization`: Visualization
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetVisualization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVisualizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSlimVisualizationParams** | [**GetSlimVisualizationParams**](GetSlimVisualizationParams.md) |  | 

### Return type

[**Visualization**](Visualization.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetXYChart

> MultiChartsResponse GetXYChart(ctx).MultiChartsParams(multiChartsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	multiChartsParams := *openapiclient.NewMultiChartsParams("ProjectId_example", *openapiclient.NewSplitAgg(NullableFloat64(123), openapiclient.OrderType("desc"), "Field_example", float64(123), "Distribution_example"), *openapiclient.NewAggregations("Field_example", openapiclient.AggregationMethod("Average")), []openapiclient.SessionRunToEpoch{*openapiclient.NewSessionRunToEpoch("SessionRunId_example", float64(123))}, false) // MultiChartsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetXYChart(context.Background()).MultiChartsParams(multiChartsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetXYChart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetXYChart`: MultiChartsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetXYChart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetXYChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multiChartsParams** | [**MultiChartsParams**](MultiChartsParams.md) |  | 

### Return type

[**MultiChartsResponse**](MultiChartsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthCheck

> HealthCheckResponse HealthCheck(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.HealthCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthCheck`: HealthCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HealthCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthCheckRequest struct via the builder pattern


### Return type

[**HealthCheckResponse**](HealthCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportExternalModel

> RunImportModelResponse ImportExternalModel(ctx).ImportExternalModelParams(importExternalModelParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	importExternalModelParams := *openapiclient.NewImportExternalModelParams("ProjectId_example", "SessionId_example", float64(123)) // ImportExternalModelParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ImportExternalModel(context.Background()).ImportExternalModelParams(importExternalModelParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ImportExternalModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportExternalModel`: RunImportModelResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ImportExternalModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportExternalModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importExternalModelParams** | [**ImportExternalModelParams**](ImportExternalModelParams.md) |  | 

### Return type

[**RunImportModelResponse**](RunImportModelResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportModel

> RunImportModelResponse ImportModel(ctx).ImportNewModelParams(importNewModelParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	importNewModelParams := *openapiclient.NewImportNewModelParams("ProjectId_example", "VersionId_example", *openapiclient.NewImportModelInfo("FileName_example", openapiclient.ImportModelType("JSON_TF2"))) // ImportNewModelParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ImportModel(context.Background()).ImportNewModelParams(importNewModelParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ImportModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportModel`: RunImportModelResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ImportModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importNewModelParams** | [**ImportNewModelParams**](ImportNewModelParams.md) |  | 

### Return type

[**RunImportModelResponse**](RunImportModelResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportProject

> ImportProjectResponse ImportProject(ctx).ImportProjectRequest(importProjectRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	importProjectRequest := *openapiclient.NewImportProjectRequest("Name_example", "Description_example", []string{"Tags_example"}, openapiclient.HubPublishPolicy("public"), "ImportUrl_example") // ImportProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ImportProject(context.Background()).ImportProjectRequest(importProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ImportProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportProject`: ImportProjectResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ImportProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importProjectRequest** | [**ImportProjectRequest**](ImportProjectRequest.md) |  | 

### Return type

[**ImportProjectResponse**](ImportProjectResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitExperiment

> InitExperimentResponse InitExperiment(ctx).InitExperimentRequest(initExperimentRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	initExperimentRequest := *openapiclient.NewInitExperimentRequest("ExperimentName_example", "CodeSnapshotId_example") // InitExperimentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.InitExperiment(context.Background()).InitExperimentRequest(initExperimentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InitExperiment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitExperiment`: InitExperimentResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.InitExperiment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitExperimentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **initExperimentRequest** | [**InitExperimentRequest**](InitExperimentRequest.md) |  | 

### Return type

[**InitExperimentResponse**](InitExperimentResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyGen

> RotateApiKeyResponse KeyGen(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.KeyGen(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.KeyGen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyGen`: RotateApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.KeyGen`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKeyGenRequest struct via the builder pattern


### Return type

[**RotateApiKeyResponse**](RotateApiKeyResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadModel

> LoadSessionResponse LoadModel(ctx).LoadSessionParams(loadSessionParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	loadSessionParams := *openapiclient.NewLoadSessionParams("SessionId_example", "ProjectId_example") // LoadSessionParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.LoadModel(context.Background()).LoadSessionParams(loadSessionParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LoadModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadModel`: LoadSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.LoadModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoadModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loadSessionParams** | [**LoadSessionParams**](LoadSessionParams.md) |  | 

### Return type

[**LoadSessionResponse**](LoadSessionResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadVersion

> LoadVersionResponse LoadVersion(ctx).LoadVersionParams(loadVersionParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	loadVersionParams := *openapiclient.NewLoadVersionParams("VersionId_example", "ProjectId_example") // LoadVersionParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.LoadVersion(context.Background()).LoadVersionParams(loadVersionParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LoadVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadVersion`: LoadVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.LoadVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoadVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loadVersionParams** | [**LoadVersionParams**](LoadVersionParams.md) |  | 

### Return type

[**LoadVersionResponse**](LoadVersionResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocalAuth

> LocalLoginResponse LocalAuth(ctx).LocalLoginParams(localLoginParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	localLoginParams := *openapiclient.NewLocalLoginParams("Email_example", "Password_example") // LocalLoginParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.LocalAuth(context.Background()).LocalLoginParams(localLoginParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LocalAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocalAuth`: LocalLoginResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.LocalAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocalAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **localLoginParams** | [**LocalLoginParams**](LocalLoginParams.md) |  | 

### Return type

[**LocalLoginResponse**](LocalLoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogExternalEpochData

> LogExternalEpochData(ctx).LogExternalEpochDataRequest(logExternalEpochDataRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	logExternalEpochDataRequest := *openapiclient.NewLogExternalEpochDataRequest("ProjectId_example", "ExperimentId_example", float64(123), map[string]EpochMetricsValue{"key": *openapiclient.NewEpochMetricsValue("Value_example", "Type_example")}) // LogExternalEpochDataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.LogExternalEpochData(context.Background()).LogExternalEpochDataRequest(logExternalEpochDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LogExternalEpochData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogExternalEpochDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logExternalEpochDataRequest** | [**LogExternalEpochDataRequest**](LogExternalEpochDataRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> UserData Login(ctx).LoginParams(loginParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	loginParams := *openapiclient.NewLoginParams("Email_example", "UserName_example", "FirstName_example", "LastName_example") // LoginParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Login(context.Background()).LoginParams(loginParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Login``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Login`: UserData
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginParams** | [**LoginParams**](LoginParams.md) |  | 

### Return type

[**UserData**](UserData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> Logout(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.Logout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Logout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverwriteModel

> RunImportModelResponse OverwriteModel(ctx).OverwriteModelParams(overwriteModelParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	overwriteModelParams := *openapiclient.NewOverwriteModelParams("ProjectId_example", "VersionId_example") // OverwriteModelParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.OverwriteModel(context.Background()).OverwriteModelParams(overwriteModelParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OverwriteModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OverwriteModel`: RunImportModelResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OverwriteModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOverwriteModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **overwriteModelParams** | [**OverwriteModelParams**](OverwriteModelParams.md) |  | 

### Return type

[**RunImportModelResponse**](RunImportModelResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PopulationExploration

> PopulationExplorationResponse PopulationExploration(ctx).PopulationExplorationParams(populationExplorationParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	populationExplorationParams := *openapiclient.NewPopulationExplorationParams("SessionRunId_example", "ProjectId_example", float64(123), float64(123), "Digest_example", float64(123), []string{"BalanceBy_example"}, false, openapiclient.ReductionAlgorithm("TSNE")) // PopulationExplorationParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PopulationExploration(context.Background()).PopulationExplorationParams(populationExplorationParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PopulationExploration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PopulationExploration`: PopulationExplorationResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PopulationExploration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPopulationExplorationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **populationExplorationParams** | [**PopulationExplorationParams**](PopulationExplorationParams.md) |  | 

### Return type

[**PopulationExplorationResponse**](PopulationExplorationResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PushCodeSnapshot

> PushCodeSnapshotResponse PushCodeSnapshot(ctx).PushCodeSnapshotParams(pushCodeSnapshotParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	pushCodeSnapshotParams := *openapiclient.NewPushCodeSnapshotParams("ProjectId_example", "CodeUrl_example", "CodeEntryFile_example", "VersionName_example") // PushCodeSnapshotParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PushCodeSnapshot(context.Background()).PushCodeSnapshotParams(pushCodeSnapshotParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PushCodeSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PushCodeSnapshot`: PushCodeSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PushCodeSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPushCodeSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushCodeSnapshotParams** | [**PushCodeSnapshotParams**](PushCodeSnapshotParams.md) |  | 

### Return type

[**PushCodeSnapshotResponse**](PushCodeSnapshotResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshLocalAuth

> LocalLoginResponse RefreshLocalAuth(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RefreshLocalAuth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RefreshLocalAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshLocalAuth`: LocalLoginResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RefreshLocalAuth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshLocalAuthRequest struct via the builder pattern


### Return type

[**LocalLoginResponse**](LocalLoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSamplesCollection

> RemoveSamplesCollection(ctx).RemoveSampleCollectionParams(removeSampleCollectionParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	removeSampleCollectionParams := *openapiclient.NewRemoveSampleCollectionParams([]string{"SampleCollectionIds_example"}, "ProjectId_example") // RemoveSampleCollectionParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RemoveSamplesCollection(context.Background()).RemoveSampleCollectionParams(removeSampleCollectionParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RemoveSamplesCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSamplesCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeSampleCollectionParams** | [**RemoveSampleCollectionParams**](RemoveSampleCollectionParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveConcurrentUsersConflict

> ResolveConcurrentUsersConflict(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ResolveConcurrentUsersConflict(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ResolveConcurrentUsersConflict``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResolveConcurrentUsersConflictRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SampleAnalysis

> Job SampleAnalysis(ctx).SampleAnalysisParams(sampleAnalysisParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	sampleAnalysisParams := *openapiclient.NewSampleAnalysisParams("SessionRunId_example", "ProjectId_example", *openapiclient.NewSampleIdentity(openapiclient.DataStateType("training"), *openapiclient.NewSampleIdType()), float64(123), openapiclient.SampleAnalysisAlgo("focus_layer_cam")) // SampleAnalysisParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SampleAnalysis(context.Background()).SampleAnalysisParams(sampleAnalysisParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SampleAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SampleAnalysis`: Job
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SampleAnalysis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSampleAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sampleAnalysisParams** | [**SampleAnalysisParams**](SampleAnalysisParams.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SamplesVisualizationsRefresh

> SamplesVisualizationsRefresh(ctx).SamplesVisualizationsRefreshParams(samplesVisualizationsRefreshParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	samplesVisualizationsRefreshParams := *openapiclient.NewSamplesVisualizationsRefreshParams("ProjectId_example", "SessionRunId_example") // SamplesVisualizationsRefreshParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.SamplesVisualizationsRefresh(context.Background()).SamplesVisualizationsRefreshParams(samplesVisualizationsRefreshParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SamplesVisualizationsRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSamplesVisualizationsRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **samplesVisualizationsRefreshParams** | [**SamplesVisualizationsRefreshParams**](SamplesVisualizationsRefreshParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveAnalyzerLayout

> SaveAnalyzerLayout(ctx).SaveAnalyzerLayoutParams(saveAnalyzerLayoutParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	saveAnalyzerLayoutParams := *openapiclient.NewSaveAnalyzerLayoutParams([]openapiclient.PanelLayout{*openapiclient.NewPanelLayout("VisualizationId_example", *openapiclient.NewSizedLayout(*openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123))))}, "ProjectId_example") // SaveAnalyzerLayoutParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.SaveAnalyzerLayout(context.Background()).SaveAnalyzerLayoutParams(saveAnalyzerLayoutParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SaveAnalyzerLayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveAnalyzerLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveAnalyzerLayoutParams** | [**SaveAnalyzerLayoutParams**](SaveAnalyzerLayoutParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendUserMessage

> SendUserMessageResponse SendUserMessage(ctx).SendUserMessageParams(sendUserMessageParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	sendUserMessageParams := *openapiclient.NewSendUserMessageParams("FirstName_example", "LastName_example", "Email_example", map[string]interface{}(123)) // SendUserMessageParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SendUserMessage(context.Background()).SendUserMessageParams(sendUserMessageParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SendUserMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendUserMessage`: SendUserMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SendUserMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendUserMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendUserMessageParams** | [**SendUserMessageParams**](SendUserMessageParams.md) |  | 

### Return type

[**SendUserMessageResponse**](SendUserMessageResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCodeChallenge

> SetCodeChallenge(ctx).SetCodeChallengeRequest(setCodeChallengeRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	setCodeChallengeRequest := *openapiclient.NewSetCodeChallengeRequest("CodeChallenge_example") // SetCodeChallengeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.SetCodeChallenge(context.Background()).SetCodeChallengeRequest(setCodeChallengeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SetCodeChallenge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetCodeChallengeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setCodeChallengeRequest** | [**SetCodeChallengeRequest**](SetCodeChallengeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultTeam

> SetDefaultTeam(ctx).SetDefaultTeamRequest(setDefaultTeamRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	setDefaultTeamRequest := *openapiclient.NewSetDefaultTeamRequest("Cid_example") // SetDefaultTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.SetDefaultTeam(context.Background()).SetDefaultTeamRequest(setDefaultTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SetDefaultTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setDefaultTeamRequest** | [**SetDefaultTeamRequest**](SetDefaultTeamRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetExperimentProperties

> SetExperimentProperties(ctx).SetExperimentPropertiesRequest(setExperimentPropertiesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	setExperimentPropertiesRequest := *openapiclient.NewSetExperimentPropertiesRequest("ProjectId_example", "ExperimentId_example", map[string]interface{}(123)) // SetExperimentPropertiesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.SetExperimentProperties(context.Background()).SetExperimentPropertiesRequest(setExperimentPropertiesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SetExperimentProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetExperimentPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setExperimentPropertiesRequest** | [**SetExperimentPropertiesRequest**](SetExperimentPropertiesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMachineType

> SetMachineType(ctx).SetTeamMachineTypeParams(setTeamMachineTypeParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	setTeamMachineTypeParams := *openapiclient.NewSetTeamMachineTypeParams("MachineTypeId_example") // SetTeamMachineTypeParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.SetMachineType(context.Background()).SetTeamMachineTypeParams(setTeamMachineTypeParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SetMachineType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetMachineTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setTeamMachineTypeParams** | [**SetTeamMachineTypeParams**](SetTeamMachineTypeParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetUserNotificationsAsRead

> SetUserNotificationsAsRead200Response SetUserNotificationsAsRead(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SetUserNotificationsAsRead(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SetUserNotificationsAsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetUserNotificationsAsRead`: SetUserNotificationsAsRead200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SetUserNotificationsAsRead`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSetUserNotificationsAsReadRequest struct via the builder pattern


### Return type

[**SetUserNotificationsAsRead200Response**](SetUserNotificationsAsRead200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartTrial

> UserData StartTrial(ctx).StartTrialParams(startTrialParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	startTrialParams := *openapiclient.NewStartTrialParams("FirstName_example", "LastName_example", "Email_example") // StartTrialParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.StartTrial(context.Background()).StartTrialParams(startTrialParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.StartTrial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartTrial`: UserData
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.StartTrial`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartTrialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTrialParams** | [**StartTrialParams**](StartTrialParams.md) |  | 

### Return type

[**UserData**](UserData.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopJob

> StopJobResponse StopJob(ctx).StopJobParams(stopJobParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	stopJobParams := *openapiclient.NewStopJobParams("JobId_example") // StopJobParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.StopJob(context.Background()).StopJobParams(stopJobParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.StopJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopJob`: StopJobResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.StopJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStopJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stopJobParams** | [**StopJobParams**](StopJobParams.md) |  | 

### Return type

[**StopJobResponse**](StopJobResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagModel

> TagModel(ctx).TagModelRequest(tagModelRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	tagModelRequest := *openapiclient.NewTagModelRequest("ProjectId_example", "ExperimentId_example", []string{"Tags_example"}, float64(123)) // TagModelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.TagModel(context.Background()).TagModelRequest(tagModelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TagModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTagModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagModelRequest** | [**TagModelRequest**](TagModelRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateAllJobs

> TerminateAllJobsResponse TerminateAllJobs(ctx).TerminateAllJobsParams(terminateAllJobsParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	terminateAllJobsParams := *openapiclient.NewTerminateAllJobsParams() // TerminateAllJobsParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TerminateAllJobs(context.Background()).TerminateAllJobsParams(terminateAllJobsParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TerminateAllJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TerminateAllJobs`: TerminateAllJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TerminateAllJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTerminateAllJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **terminateAllJobsParams** | [**TerminateAllJobsParams**](TerminateAllJobsParams.md) |  | 

### Return type

[**TerminateAllJobsResponse**](TerminateAllJobsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateJob

> TerminateJobResponse TerminateJob(ctx).TerminateJobParams(terminateJobParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	terminateJobParams := *openapiclient.NewTerminateJobParams("JobId_example") // TerminateJobParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TerminateJob(context.Background()).TerminateJobParams(terminateJobParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TerminateJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TerminateJob`: TerminateJobResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TerminateJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTerminateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **terminateJobParams** | [**TerminateJobParams**](TerminateJobParams.md) |  | 

### Return type

[**TerminateJobResponse**](TerminateJobResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrashSecretManager

> TrashSecretManagerResponse TrashSecretManager(ctx).TrashSecretManagerParams(trashSecretManagerParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	trashSecretManagerParams := *openapiclient.NewTrashSecretManagerParams("SecretManagerId_example") // TrashSecretManagerParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TrashSecretManager(context.Background()).TrashSecretManagerParams(trashSecretManagerParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TrashSecretManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrashSecretManager`: TrashSecretManagerResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TrashSecretManager`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrashSecretManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trashSecretManagerParams** | [**TrashSecretManagerParams**](TrashSecretManagerParams.md) |  | 

### Return type

[**TrashSecretManagerResponse**](TrashSecretManagerResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboard

> UpdateDashboard(ctx).UpdateDashboardParams(updateDashboardParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	updateDashboardParams := *openapiclient.NewUpdateDashboardParams("DashboardId_example", "Name_example", []openapiclient.Dashlet{*openapiclient.NewDashlet("Cid_example", *openapiclient.NewSizedLayout(*openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123))), "Type_example", *openapiclient.NewPopulationExplorationDashletData(map[string]interface{}(123), "Name_example", "Type_example"), "Name_example", []string{"CollectionIds_example"})}, "ProjectId_example") // UpdateDashboardParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateDashboard(context.Background()).UpdateDashboardParams(updateDashboardParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDashboardParams** | [**UpdateDashboardParams**](UpdateDashboardParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEngineSettings

> UpdateEngineSettings(ctx).SetSettingValueWrapper(setSettingValueWrapper).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	setSettingValueWrapper := *openapiclient.NewSetSettingValueWrapper() // SetSettingValueWrapper | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateEngineSettings(context.Background()).SetSettingValueWrapper(setSettingValueWrapper).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateEngineSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEngineSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setSettingValueWrapper** | [**SetSettingValueWrapper**](SetSettingValueWrapper.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIssue

> Issue UpdateIssue(ctx).UpdateIssueParams(updateIssueParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	updateIssueParams := *openapiclient.NewUpdateIssueParams("Cid_example", "ProjectId_example") // UpdateIssueParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateIssue(context.Background()).UpdateIssueParams(updateIssueParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIssue`: Issue
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateIssue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateIssueParams** | [**UpdateIssueParams**](UpdateIssueParams.md) |  | 

### Return type

[**Issue**](Issue.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectMeta

> UpdateProjectMeta(ctx).UpdateProjectMetaRequest(updateProjectMetaRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	updateProjectMetaRequest := *openapiclient.NewUpdateProjectMetaRequest("Name_example", "Description_example", []string{"Tags_example"}, openapiclient.HubPublishPolicy("public"), "ProjectId_example") // UpdateProjectMetaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateProjectMeta(context.Background()).UpdateProjectMetaRequest(updateProjectMetaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateProjectMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateProjectMetaRequest** | [**UpdateProjectMetaRequest**](UpdateProjectMetaRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSampleCollection

> UpdateSampleCollection(ctx).UpdateSampleCollectionParams(updateSampleCollectionParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	updateSampleCollectionParams := *openapiclient.NewUpdateSampleCollectionParams("SampleCollectionId_example", "ProjectId_example") // UpdateSampleCollectionParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateSampleCollection(context.Background()).UpdateSampleCollectionParams(updateSampleCollectionParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateSampleCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSampleCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSampleCollectionParams** | [**UpdateSampleCollectionParams**](UpdateSampleCollectionParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecretManager

> UpdateSecretManagerResponse UpdateSecretManager(ctx).UpdateSecretManagerParams(updateSecretManagerParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	updateSecretManagerParams := *openapiclient.NewUpdateSecretManagerParams("Cid_example", "Name_example") // UpdateSecretManagerParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateSecretManager(context.Background()).UpdateSecretManagerParams(updateSecretManagerParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateSecretManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSecretManager`: UpdateSecretManagerResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateSecretManager`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecretManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSecretManagerParams** | [**UpdateSecretManagerParams**](UpdateSecretManagerParams.md) |  | 

### Return type

[**UpdateSecretManagerResponse**](UpdateSecretManagerResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSessionName

> UpdateSessionName(ctx).UpdateSessionNameParams(updateSessionNameParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	updateSessionNameParams := *openapiclient.NewUpdateSessionNameParams("Cid_example", "ProjectId_example", "Name_example") // UpdateSessionNameParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateSessionName(context.Background()).UpdateSessionNameParams(updateSessionNameParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateSessionName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSessionNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSessionNameParams** | [**UpdateSessionNameParams**](UpdateSessionNameParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSessionRun

> UpdateSessionRun(ctx).UpdateSessionRunNameParams(updateSessionRunNameParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	updateSessionRunNameParams := *openapiclient.NewUpdateSessionRunNameParams("Cid_example", "Name_example", "Description_example", "ProjectId_example") // UpdateSessionRunNameParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateSessionRun(context.Background()).UpdateSessionRunNameParams(updateSessionRunNameParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateSessionRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSessionRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSessionRunNameParams** | [**UpdateSessionRunNameParams**](UpdateSessionRunNameParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSessionTest

> UpdateSessionTest(ctx).UpdateSessionTestRequest(updateSessionTestRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	updateSessionTestRequest := *openapiclient.NewUpdateSessionTestRequest("Cid_example", "ProjectId_example") // UpdateSessionTestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateSessionTest(context.Background()).UpdateSessionTestRequest(updateSessionTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateSessionTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSessionTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSessionTestRequest** | [**UpdateSessionTestRequest**](UpdateSessionTestRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeamPublicName

> UpdateTeamPublicName(ctx).UpdateTeamPublicNameRequest(updateTeamPublicNameRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	updateTeamPublicNameRequest := *openapiclient.NewUpdateTeamPublicNameRequest("Cid_example", "PublicName_example") // UpdateTeamPublicNameRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateTeamPublicName(context.Background()).UpdateTeamPublicNameRequest(updateTeamPublicNameRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateTeamPublicName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamPublicNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTeamPublicNameRequest** | [**UpdateTeamPublicNameRequest**](UpdateTeamPublicNameRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserName

> UserData UpdateUserName(ctx).UpdateUserNameRequest(updateUserNameRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	updateUserNameRequest := *openapiclient.NewUpdateUserNameRequest("UserName_example") // UpdateUserNameRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateUserName(context.Background()).UpdateUserNameRequest(updateUserNameRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateUserName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserName`: UserData
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateUserName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserNameRequest** | [**UpdateUserNameRequest**](UpdateUserNameRequest.md) |  | 

### Return type

[**UserData**](UserData.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserRole

> SlimUserData UpdateUserRole(ctx).UpdateUserRoleRequest(updateUserRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	updateUserRoleRequest := *openapiclient.NewUpdateUserRoleRequest("UserId_example", "Role_example") // UpdateUserRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateUserRole(context.Background()).UpdateUserRoleRequest(updateUserRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateUserRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserRole`: SlimUserData
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateUserRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserRoleRequest** | [**UpdateUserRoleRequest**](UpdateUserRoleRequest.md) |  | 

### Return type

[**SlimUserData**](SlimUserData.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserStatus

> SlimUserData UpdateUserStatus(ctx).UpdateUserStatusRequest(updateUserStatusRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	updateUserStatusRequest := *openapiclient.NewUpdateUserStatusRequest("UserId_example", false) // UpdateUserStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateUserStatus(context.Background()).UpdateUserStatusRequest(updateUserStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateUserStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserStatus`: SlimUserData
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateUserStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserStatusRequest** | [**UpdateUserStatusRequest**](UpdateUserStatusRequest.md) |  | 

### Return type

[**SlimUserData**](SlimUserData.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserTeam

> SlimUserData UpdateUserTeam(ctx).UpdateUserTeamRequest(updateUserTeamRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	updateUserTeamRequest := *openapiclient.NewUpdateUserTeamRequest("UserId_example", "TeamId_example") // UpdateUserTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateUserTeam(context.Background()).UpdateUserTeamRequest(updateUserTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateUserTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserTeam`: SlimUserData
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateUserTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserTeamRequest** | [**UpdateUserTeamRequest**](UpdateUserTeamRequest.md) |  | 

### Return type

[**SlimUserData**](SlimUserData.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVersion

> UpdateVersionResponse UpdateVersion(ctx).UpdateVersionParams(updateVersionParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	updateVersionParams := *openapiclient.NewUpdateVersionParams("VersionId_example", "ProjectId_example", *openapiclient.NewModelGraph("Id_example", map[string]interface{}(123))) // UpdateVersionParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateVersion(context.Background()).UpdateVersionParams(updateVersionParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVersion`: UpdateVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateVersionParams** | [**UpdateVersionParams**](UpdateVersionParams.md) |  | 

### Return type

[**UpdateVersionResponse**](UpdateVersionResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVersionName

> UpdateVersionName(ctx).UpdateVersionNameParams(updateVersionNameParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	updateVersionNameParams := *openapiclient.NewUpdateVersionNameParams("Cid_example", "ProjectId_example", "Name_example") // UpdateVersionNameParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateVersionName(context.Background()).UpdateVersionNameParams(updateVersionNameParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateVersionName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVersionNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateVersionNameParams** | [**UpdateVersionNameParams**](UpdateVersionNameParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadProject

> UploadProject(ctx, projectName).Description(description).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	projectName := "projectName_example" // string | 
	description := "description_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UploadProject(context.Background(), projectName).Description(description).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UploadProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **string** |  | [default to &quot;&quot;]

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertState

> UpsertStateResponse UpsertState(ctx).UpsertStateParams(upsertStateParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {
	upsertStateParams := *openapiclient.NewUpsertStateParams("ProjectId_example", "State_example") // UpsertStateParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpsertState(context.Background()).UpsertStateParams(upsertStateParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpsertState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertState`: UpsertStateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpsertState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertStateParams** | [**UpsertStateParams**](UpsertStateParams.md) |  | 

### Return type

[**UpsertStateResponse**](UpsertStateResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Warmup

> Warmup(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.Warmup(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Warmup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWarmupRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WhoAmI

> UserData WhoAmI(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi/tensorleapapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.WhoAmI(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WhoAmI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WhoAmI`: UserData
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.WhoAmI`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWhoAmIRequest struct via the builder pattern


### Return type

[**UserData**](UserData.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

