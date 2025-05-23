# \DefaultApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Activate**](DefaultApi.md#Activate) | **Post** /auth/activate | 
[**AddDashboard**](DefaultApi.md#AddDashboard) | **Post** /dashboards/addDashboard | 
[**AddDataset**](DefaultApi.md#AddDataset) | **Post** /datasets/addDataset | 
[**AddExportModelJob**](DefaultApi.md#AddExportModelJob) | **Post** /exportedsessionruns/addExportModelJob | 
[**AddIssue**](DefaultApi.md#AddIssue) | **Post** /issues/addIssue | 
[**AddProject**](DefaultApi.md#AddProject) | **Post** /projects/addProject | 
[**AddSampleCollection**](DefaultApi.md#AddSampleCollection) | **Post** /sample-collection/addSampleCollection | 
[**AddSecretManager**](DefaultApi.md#AddSecretManager) | **Post** /secret-manager/addSecretManager | 
[**AddVersion**](DefaultApi.md#AddVersion) | **Post** /versions/addVersion | 
[**AnalyzeGraph**](DefaultApi.md#AnalyzeGraph) | **Post** /graph/analyzeGraph | 
[**ApplyDatasetMapping**](DefaultApi.md#ApplyDatasetMapping) | **Post** /graph/applyDatasetMapping | 
[**ArchiveInsight**](DefaultApi.md#ArchiveInsight) | **Post** /insights/archiveInsight | 
[**ClearUserJobs**](DefaultApi.md#ClearUserJobs) | **Post** /users/clearUserJobs | 
[**ClearUserNotifications**](DefaultApi.md#ClearUserNotifications) | **Post** /users/clearUserNotifications | 
[**ContinueEvaluate**](DefaultApi.md#ContinueEvaluate) | **Post** /evaluate/continueEvaluate | 
[**CreateSamplesVisualizations**](DefaultApi.md#CreateSamplesVisualizations) | **Post** /visualizations/createSamplesVisualizations | 
[**CreateSessionTest**](DefaultApi.md#CreateSessionTest) | **Post** /sessions-tests/createSessionTest | 
[**CreateTeam**](DefaultApi.md#CreateTeam) | **Post** /teams/createTeam | 
[**DeleteDashboard**](DefaultApi.md#DeleteDashboard) | **Post** /dashboards/deleteDashboard | 
[**DeleteIssue**](DefaultApi.md#DeleteIssue) | **Post** /issues/deleteIssue | 
[**DeleteProject**](DefaultApi.md#DeleteProject) | **Post** /projects/deleteProject | 
[**DeleteSampleAnalysis**](DefaultApi.md#DeleteSampleAnalysis) | **Post** /visualizations/deleteSampleAnalysis | 
[**DeleteSession**](DefaultApi.md#DeleteSession) | **Post** /sessions/deleteSession | 
[**DeleteSessionRun**](DefaultApi.md#DeleteSessionRun) | **Post** /sessions/deleteSessionRun | 
[**DeleteSessionTest**](DefaultApi.md#DeleteSessionTest) | **Post** /sessions-tests/deleteSessionTest | 
[**DeleteTeam**](DefaultApi.md#DeleteTeam) | **Post** /teams/deleteTeam | 
[**DeleteUserById**](DefaultApi.md#DeleteUserById) | **Post** /users/deleteUserById | 
[**DeleteVersion**](DefaultApi.md#DeleteVersion) | **Post** /versions/deleteVersion | 
[**DeleteVisualizations**](DefaultApi.md#DeleteVisualizations) | **Post** /visualizations/deleteVisualizations | 
[**DownloadProject**](DefaultApi.md#DownloadProject) | **Get** /projects/downloadProject/{projectId} | 
[**Evaluate**](DefaultApi.md#Evaluate) | **Post** /evaluate/evaluate | 
[**ExportProject**](DefaultApi.md#ExportProject) | **Post** /projects/exportProject | 
[**ExtendTrial**](DefaultApi.md#ExtendTrial) | **Post** /auth/extendTrial | 
[**FetchSimilar**](DefaultApi.md#FetchSimilar) | **Post** /visualizations/fetchSimilar | 
[**GetAllProjectSessionTests**](DefaultApi.md#GetAllProjectSessionTests) | **Post** /sessions-tests/getAllProjectSessionTests | 
[**GetAllSlimUserData**](DefaultApi.md#GetAllSlimUserData) | **Post** /users/getAllSlimUserData | 
[**GetApiKeyByCode**](DefaultApi.md#GetApiKeyByCode) | **Post** /auth/getApiKeyByCode | 
[**GetArchiveInsights**](DefaultApi.md#GetArchiveInsights) | **Post** /insights/getArchiveInsights | 
[**GetAuthProvider**](DefaultApi.md#GetAuthProvider) | **Post** /auth/getAuthProvider | 
[**GetBalancedAccuracy**](DefaultApi.md#GetBalancedAccuracy) | **Post** /sessionmetrics/getBalancedAccuracy | 
[**GetCodeIntegrationMappingErrorsByVersionId**](DefaultApi.md#GetCodeIntegrationMappingErrorsByVersionId) | **Post** /graph/getCodeIntegrationMappingErrorsByVersionId | 
[**GetCodeIntegrationVersionFromBinder**](DefaultApi.md#GetCodeIntegrationVersionFromBinder) | **Post** /versions/getCodeIntegrationVersionFromBinder | 
[**GetConfusionMatrixLabels**](DefaultApi.md#GetConfusionMatrixLabels) | **Post** /sessionmetrics/getConfusionMatrixLabels | 
[**GetConfusionMatrixResultCombinations**](DefaultApi.md#GetConfusionMatrixResultCombinations) | **Post** /sessionmetrics/getConfusionMatrixResultCombinations | 
[**GetConfusionMatrixTable**](DefaultApi.md#GetConfusionMatrixTable) | **Post** /sessionmetrics/getConfusionMatrixTable | 
[**GetConfusionMetricNames**](DefaultApi.md#GetConfusionMetricNames) | **Post** /sessionmetrics/getConfusionMetricNames | 
[**GetCurrentProjectVersion**](DefaultApi.md#GetCurrentProjectVersion) | **Post** /projects/getCurrentProjectVersion | 
[**GetDashboard**](DefaultApi.md#GetDashboard) | **Post** /dashboards/getDashboard | 
[**GetDashletFields**](DefaultApi.md#GetDashletFields) | **Post** /dashboards/getDashletFields | 
[**GetDatasetVersion**](DefaultApi.md#GetDatasetVersion) | **Post** /datasetVersions/getDatasetVersion | 
[**GetDatasetVersionUploadUrl**](DefaultApi.md#GetDatasetVersionUploadUrl) | **Post** /datasetVersions/getDatasetVersionUploadUrl | 
[**GetDatasetVersions**](DefaultApi.md#GetDatasetVersions) | **Post** /datasetVersions/getDatasetVersions | 
[**GetDatasets**](DefaultApi.md#GetDatasets) | **Post** /datasets/getDatasets | 
[**GetDownloadSignedUrl**](DefaultApi.md#GetDownloadSignedUrl) | **Post** /versions/getDownloadSignedUrl | 
[**GetEngineSettings**](DefaultApi.md#GetEngineSettings) | **Post** /settings/getEngineSettings | 
[**GetEnvironmentInfo**](DefaultApi.md#GetEnvironmentInfo) | **Post** /metadata/getEnvironmentInfo | 
[**GetExportedSessionJobs**](DefaultApi.md#GetExportedSessionJobs) | **Post** /exportedsessionruns/getExportedSessionJobs | 
[**GetF1Score**](DefaultApi.md#GetF1Score) | **Post** /sessionmetrics/getF1Score | 
[**GetFetchSimilarStatus**](DefaultApi.md#GetFetchSimilarStatus) | **Post** /visualizations/getFetchSimilarStatus | 
[**GetFieldsValues**](DefaultApi.md#GetFieldsValues) | **Post** /sessionmetrics/getFieldsValues | 
[**GetGeneratedDashlets**](DefaultApi.md#GetGeneratedDashlets) | **Post** /dashletsGenerator/getGeneratedDashlets | 
[**GetGenericBaseImageTypes**](DefaultApi.md#GetGenericBaseImageTypes) | **Post** /datasetVersions/getGenericBaseImageTypes | 
[**GetHeatmapChart**](DefaultApi.md#GetHeatmapChart) | **Post** /sessionmetrics/getHeatmapChart | 
[**GetIssueFileUploadSignedUrl**](DefaultApi.md#GetIssueFileUploadSignedUrl) | **Post** /issues/getIssueFileUploadSignedUrl | 
[**GetJobLogs**](DefaultApi.md#GetJobLogs) | **Post** /jobs/getJobLogs | 
[**GetLatestDatasetVersion**](DefaultApi.md#GetLatestDatasetVersion) | **Post** /datasetVersions/getLatestDatasetVersion | 
[**GetMachineTypes**](DefaultApi.md#GetMachineTypes) | **Post** /teams/getMachineTypes | 
[**GetMaxActiveUsers**](DefaultApi.md#GetMaxActiveUsers) | **Post** /metadata/getMaxActiveUsers | 
[**GetMeanAveragePrecision**](DefaultApi.md#GetMeanAveragePrecision) | **Post** /sessionmetrics/getMeanAveragePrecision | 
[**GetNotifications**](DefaultApi.md#GetNotifications) | **Post** /notifications/getNotifications | 
[**GetPopulationExplorationStatus**](DefaultApi.md#GetPopulationExplorationStatus) | **Post** /visualizations/getPopulationExplorationStatus | 
[**GetPrCurve**](DefaultApi.md#GetPrCurve) | **Post** /sessionmetrics/getPrCurve | 
[**GetPrecisionScore**](DefaultApi.md#GetPrecisionScore) | **Post** /sessionmetrics/getPrecisionScore | 
[**GetProjectDashboards**](DefaultApi.md#GetProjectDashboards) | **Post** /dashboards/getProjectDashboards | 
[**GetProjectIssues**](DefaultApi.md#GetProjectIssues) | **Post** /issues/getProjectIssues | 
[**GetProjectSlimVersions**](DefaultApi.md#GetProjectSlimVersions) | **Post** /versions/getProjectSlimVersions | 
[**GetProjects**](DefaultApi.md#GetProjects) | **Post** /projects/getProjects | 
[**GetRecallScore**](DefaultApi.md#GetRecallScore) | **Post** /sessionmetrics/getRecallScore | 
[**GetRecentTeamSessions**](DefaultApi.md#GetRecentTeamSessions) | **Post** /sessions/getRecentTeamSessions | 
[**GetRoc**](DefaultApi.md#GetRoc) | **Post** /sessionmetrics/getRoc | 
[**GetSampleCollections**](DefaultApi.md#GetSampleCollections) | **Post** /sample-collection/getSampleCollections | 
[**GetSampleVisualizationsPath**](DefaultApi.md#GetSampleVisualizationsPath) | **Post** /visualizations/getSampleVisualizationsPath | 
[**GetScatterSampleVisualizations**](DefaultApi.md#GetScatterSampleVisualizations) | **Post** /visualizations/getScatterSampleVisualizations | 
[**GetSecretManagerList**](DefaultApi.md#GetSecretManagerList) | **Post** /secret-manager/getSecretManagerList | 
[**GetSessionRunsVisualizations**](DefaultApi.md#GetSessionRunsVisualizations) | **Post** /visualizations/getSessionRunsVisualizations | 
[**GetSessionTestResult**](DefaultApi.md#GetSessionTestResult) | **Post** /sessions-tests/getSessionTestResult | 
[**GetSessionsByHash**](DefaultApi.md#GetSessionsByHash) | **Post** /sessions/getSessionsByHash | 
[**GetSessionsByVersionId**](DefaultApi.md#GetSessionsByVersionId) | **Post** /sessions/getSessionsByVersionId | 
[**GetSessionsEpochs**](DefaultApi.md#GetSessionsEpochs) | **Post** /versions/getSessionsEpochs | 
[**GetSignedUrl**](DefaultApi.md#GetSignedUrl) | **Post** /versions/getSignedUrl | 
[**GetSingleIssue**](DefaultApi.md#GetSingleIssue) | **Post** /issues/getSingleIssue | 
[**GetSingleSessionTest**](DefaultApi.md#GetSingleSessionTest) | **Post** /sessions-tests/getSingleSessionTest | 
[**GetSlimJobs**](DefaultApi.md#GetSlimJobs) | **Post** /jobs/getSlimJobs | 
[**GetSlimVisualization**](DefaultApi.md#GetSlimVisualization) | **Post** /visualizations/getSlimVisualization | 
[**GetState**](DefaultApi.md#GetState) | **Post** /projectstate/getState | 
[**GetStatistics**](DefaultApi.md#GetStatistics) | **Post** /metadata/getStatistics | 
[**GetTableChart**](DefaultApi.md#GetTableChart) | **Post** /sessionmetrics/getTableChart | 
[**GetTeamJobs**](DefaultApi.md#GetTeamJobs) | **Post** /jobs/getTeamJobs | 
[**GetTeamSlimUserData**](DefaultApi.md#GetTeamSlimUserData) | **Post** /users/getTeamSlimUserData | 
[**GetTeams**](DefaultApi.md#GetTeams) | **Post** /teams/getTeams | 
[**GetUploadModelSignedUrl**](DefaultApi.md#GetUploadModelSignedUrl) | **Post** /versions/getUploadModelSignedUrl | 
[**GetUploadSignedUrl**](DefaultApi.md#GetUploadSignedUrl) | **Post** /versions/getUploadSignedUrl | 
[**GetValidateGraphProcessState**](DefaultApi.md#GetValidateGraphProcessState) | **Post** /graph/getValidateGraphProcessState | 
[**GetVisualization**](DefaultApi.md#GetVisualization) | **Post** /visualizations/getVisualization | 
[**GetXYChart**](DefaultApi.md#GetXYChart) | **Post** /sessionmetrics/getXYChart | 
[**HealthCheck**](DefaultApi.md#HealthCheck) | **Get** /monitor/healthCheck | 
[**ImportExternalModel**](DefaultApi.md#ImportExternalModel) | **Post** /versions/importExternalModel | 
[**ImportModel**](DefaultApi.md#ImportModel) | **Post** /versions/importModel | 
[**ImportProject**](DefaultApi.md#ImportProject) | **Post** /projects/importProject | 
[**InitExperiment**](DefaultApi.md#InitExperiment) | **Post** /versions/initExperiment | 
[**KeyGen**](DefaultApi.md#KeyGen) | **Post** /auth/keygen | 
[**LoadModel**](DefaultApi.md#LoadModel) | **Post** /projects/loadModel | 
[**LoadVersion**](DefaultApi.md#LoadVersion) | **Post** /versions/loadVersion | 
[**LocalAuth**](DefaultApi.md#LocalAuth) | **Post** /auth/localAuth | 
[**LogExternalEpochData**](DefaultApi.md#LogExternalEpochData) | **Post** /externalepochdata/logExternalEpochData | 
[**Login**](DefaultApi.md#Login) | **Post** /auth/login | 
[**Logout**](DefaultApi.md#Logout) | **Post** /auth/logout | 
[**ModifyDatasetVersionNote**](DefaultApi.md#ModifyDatasetVersionNote) | **Post** /datasetVersions/modifyDatasetVersionNote | 
[**PopulationExploration**](DefaultApi.md#PopulationExploration) | **Post** /visualizations/populationExploration | 
[**RefreshLocalAuth**](DefaultApi.md#RefreshLocalAuth) | **Post** /auth/refreshLocalAuth | 
[**RemoveSamplesCollection**](DefaultApi.md#RemoveSamplesCollection) | **Post** /sample-collection/removeSampleCollections | 
[**ResolveConcurrentUsersConflict**](DefaultApi.md#ResolveConcurrentUsersConflict) | **Post** /auth/resolveConcurrentUsersConflict | 
[**SampleAnalysis**](DefaultApi.md#SampleAnalysis) | **Post** /visualizations/sampleAnalysis | 
[**SamplesVisualizationsRefresh**](DefaultApi.md#SamplesVisualizationsRefresh) | **Post** /visualizations/samplesVisualizationsRefresh | 
[**SaveAnalyzerLayout**](DefaultApi.md#SaveAnalyzerLayout) | **Post** /visualizations/saveAnalyzerLayout | 
[**SaveDatasetVersion**](DefaultApi.md#SaveDatasetVersion) | **Post** /datasetVersions/saveDatasetVersion | 
[**SaveDatasetVersionNoParse**](DefaultApi.md#SaveDatasetVersionNoParse) | **Post** /datasetVersions/saveDatasetVersionNoParse | 
[**SendUserMessage**](DefaultApi.md#SendUserMessage) | **Post** /users/sendUserMessage | 
[**SetCodeChallenge**](DefaultApi.md#SetCodeChallenge) | **Post** /auth/setCodeChallenge | 
[**SetDefaultTeam**](DefaultApi.md#SetDefaultTeam) | **Post** /teams/setDefaultTeam | 
[**SetExperimentProperties**](DefaultApi.md#SetExperimentProperties) | **Post** /versions/setExperimentProperties | 
[**SetMachineType**](DefaultApi.md#SetMachineType) | **Post** /teams/setMachineType | 
[**SetUserNotificationsAsRead**](DefaultApi.md#SetUserNotificationsAsRead) | **Post** /notifications/setUserNotificationsAsRead | 
[**StartTrial**](DefaultApi.md#StartTrial) | **Post** /auth/startTrial | 
[**StopJob**](DefaultApi.md#StopJob) | **Post** /jobs/stopJob | 
[**TagModel**](DefaultApi.md#TagModel) | **Post** /versions/tagModel | 
[**TerminateAllJobs**](DefaultApi.md#TerminateAllJobs) | **Post** /jobs/terminateAllJobs | 
[**TerminateJob**](DefaultApi.md#TerminateJob) | **Post** /jobs/terminateJob | 
[**TrashDataset**](DefaultApi.md#TrashDataset) | **Post** /datasets/trashDataset | 
[**TrashSecretManager**](DefaultApi.md#TrashSecretManager) | **Post** /secret-manager/trashSecretManager | 
[**UnarchiveInsight**](DefaultApi.md#UnarchiveInsight) | **Post** /insights/unarchiveInsight | 
[**UpdateDashboard**](DefaultApi.md#UpdateDashboard) | **Post** /dashboards/updateDashboard | 
[**UpdateEngineSettings**](DefaultApi.md#UpdateEngineSettings) | **Post** /settings/updateSetting | 
[**UpdateIssue**](DefaultApi.md#UpdateIssue) | **Post** /issues/updateIssue | 
[**UpdateProjectMeta**](DefaultApi.md#UpdateProjectMeta) | **Post** /projects/updateProjectMeta | 
[**UpdateSampleCollection**](DefaultApi.md#UpdateSampleCollection) | **Post** /sample-collection/updateSampleCollection | 
[**UpdateSecretManager**](DefaultApi.md#UpdateSecretManager) | **Post** /secret-manager/updateSecretManager | 
[**UpdateSessionName**](DefaultApi.md#UpdateSessionName) | **Post** /sessions/updateSessionName | 
[**UpdateSessionRun**](DefaultApi.md#UpdateSessionRun) | **Post** /sessionsruns/updateSessionRun | 
[**UpdateSessionTest**](DefaultApi.md#UpdateSessionTest) | **Post** /sessions-tests/updateSessionTest | 
[**UpdateTeamPublicName**](DefaultApi.md#UpdateTeamPublicName) | **Post** /teams/updateTeamPublicName | 
[**UpdateUserName**](DefaultApi.md#UpdateUserName) | **Post** /users/updateUserName | 
[**UpdateUserRole**](DefaultApi.md#UpdateUserRole) | **Post** /users/updateUserRole | 
[**UpdateUserStatus**](DefaultApi.md#UpdateUserStatus) | **Post** /users/updateUserStatus | 
[**UpdateUserTeam**](DefaultApi.md#UpdateUserTeam) | **Post** /users/updateUserTeam | 
[**UpdateVersion**](DefaultApi.md#UpdateVersion) | **Post** /versions/updateVersion | 
[**UpdateVersionName**](DefaultApi.md#UpdateVersionName) | **Post** /versions/updateVersionName | 
[**UploadProject**](DefaultApi.md#UploadProject) | **Put** /projects/uploadProject/{projectName} | 
[**UpsertState**](DefaultApi.md#UpsertState) | **Post** /projectstate/upsertState | 
[**ValidateGraph**](DefaultApi.md#ValidateGraph) | **Post** /graph/validateGraph | 
[**Warmup**](DefaultApi.md#Warmup) | **Post** /jobs/warmup | 
[**WhoAmI**](DefaultApi.md#WhoAmI) | **Post** /auth/whoAmI | 



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
    resp, r, err := apiClient.DefaultApi.Activate(context.Background()).ActivationParams(activationParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Activate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Activate`: UserData
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Activate`: %v\n", resp)
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
    addDashboardParams := *openapiclient.NewAddDashboardParams("ProjectId_example", "Name_example", []openapiclient.Dashlet{*openapiclient.NewDashlet("Cid_example", *openapiclient.NewSizedLayout(*openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123))), "Type_example", *openapiclient.NewPopulationExplorationDashletData(map[string]interface{}(123), "Type_example", "Name_example"), "Name_example", []string{"CollectionIds_example"})}) // AddDashboardParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddDashboard(context.Background()).AddDashboardParams(addDashboardParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDashboard`: AddDashboardResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddDashboard`: %v\n", resp)
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


## AddDataset

> AddDatasetResponse AddDataset(ctx).NewDatasetParams(newDatasetParams).Execute()



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
    newDatasetParams := *openapiclient.NewNewDatasetParams("Name_example") // NewDatasetParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddDataset(context.Background()).NewDatasetParams(newDatasetParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddDataset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDataset`: AddDatasetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddDataset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDatasetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newDatasetParams** | [**NewDatasetParams**](NewDatasetParams.md) |  | 

### Return type

[**AddDatasetResponse**](AddDatasetResponse.md)

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
    resp, r, err := apiClient.DefaultApi.AddExportModelJob(context.Background()).AddExportModelJobParams(addExportModelJobParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddExportModelJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddExportModelJob`: Job
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddExportModelJob`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.AddIssue(context.Background()).AddIssueParams(addIssueParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddIssue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIssue`: Issue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddIssue`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.AddProject(context.Background()).ProjectMeta(projectMeta).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddProject`: AddProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddProject`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.AddSampleCollection(context.Background()).AddSampleCollectionParams(addSampleCollectionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddSampleCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSampleCollection`: AddSampleCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddSampleCollection`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.AddSecretManager(context.Background()).AddSecretManagerParams(addSecretManagerParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddSecretManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSecretManager`: AddSecretManagerResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddSecretManager`: %v\n", resp)
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


## AddVersion

> AddVersionResponse AddVersion(ctx).AddVersionParams(addVersionParams).Execute()



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
    addVersionParams := *openapiclient.NewAddVersionParams("ProjectId_example", *openapiclient.NewModelGraph("Id_example", map[string]interface{}(123)), "BranchName_example", "Description_example") // AddVersionParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddVersion(context.Background()).AddVersionParams(addVersionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVersion`: AddVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addVersionParams** | [**AddVersionParams**](AddVersionParams.md) |  | 

### Return type

[**AddVersionResponse**](AddVersionResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyzeGraph

> AnalyzeGraphResponse AnalyzeGraph(ctx).AnalyzeGraphParams(analyzeGraphParams).Execute()



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
    analyzeGraphParams := *openapiclient.NewAnalyzeGraphParams(*openapiclient.NewModelGraph("Id_example", map[string]interface{}(123)), "VersionId_example", "ProjectId_example", "Digest_example") // AnalyzeGraphParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AnalyzeGraph(context.Background()).AnalyzeGraphParams(analyzeGraphParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AnalyzeGraph``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AnalyzeGraph`: AnalyzeGraphResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AnalyzeGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyzeGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyzeGraphParams** | [**AnalyzeGraphParams**](AnalyzeGraphParams.md) |  | 

### Return type

[**AnalyzeGraphResponse**](AnalyzeGraphResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyDatasetMapping

> ApplyMappingResponse ApplyDatasetMapping(ctx).ApplyDatasetMappingsParams(applyDatasetMappingsParams).Execute()



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
    applyDatasetMappingsParams := *openapiclient.NewApplyDatasetMappingsParams(*openapiclient.NewModelGraph("Id_example", map[string]interface{}(123)), "Yaml_example") // ApplyDatasetMappingsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApplyDatasetMapping(context.Background()).ApplyDatasetMappingsParams(applyDatasetMappingsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApplyDatasetMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyDatasetMapping`: ApplyMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApplyDatasetMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyDatasetMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applyDatasetMappingsParams** | [**ApplyDatasetMappingsParams**](ApplyDatasetMappingsParams.md) |  | 

### Return type

[**ApplyMappingResponse**](ApplyMappingResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchiveInsight

> ArchiveInsight(ctx).ArchiveInsightsParams(archiveInsightsParams).Execute()



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
    archiveInsightsParams := *openapiclient.NewArchiveInsightsParams("PopExpDigest_example", []float64{float64(123)}, "ProjectId_example") // ArchiveInsightsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.ArchiveInsight(context.Background()).ArchiveInsightsParams(archiveInsightsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ArchiveInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archiveInsightsParams** | [**ArchiveInsightsParams**](ArchiveInsightsParams.md) |  | 

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
    r, err := apiClient.DefaultApi.ClearUserJobs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ClearUserJobs``: %v\n", err)
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
    r, err := apiClient.DefaultApi.ClearUserNotifications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ClearUserNotifications``: %v\n", err)
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
    continueEvaluateParams := *openapiclient.NewContinueEvaluateParams("JobId_example", "ProjectId_example") // ContinueEvaluateParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ContinueEvaluate(context.Background()).ContinueEvaluateParams(continueEvaluateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContinueEvaluate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContinueEvaluate`: Job
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContinueEvaluate`: %v\n", resp)
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
    createSampleVisualizationsParams := *openapiclient.NewCreateSampleVisualizationsParams("ProjectId_example", "SessionRunId_example", float64(123), []openapiclient.SampleIdentity{*openapiclient.NewSampleIdentity(openapiclient.DataStateType("training"), *openapiclient.NewSampleIdType())}, "Digest_example") // CreateSampleVisualizationsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSamplesVisualizations(context.Background()).CreateSampleVisualizationsParams(createSampleVisualizationsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSamplesVisualizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSamplesVisualizations`: Job
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSamplesVisualizations`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.CreateSessionTest(context.Background()).CreateSessionTestRequest(createSessionTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSessionTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionTest`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSessionTest`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.CreateTeam(context.Background()).CreateTeamRequest(createTeamRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTeam`: CreateTeamResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTeam`: %v\n", resp)
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
    r, err := apiClient.DefaultApi.DeleteDashboard(context.Background()).DeleteDashboardParams(deleteDashboardParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteDashboard``: %v\n", err)
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
    r, err := apiClient.DefaultApi.DeleteIssue(context.Background()).DeleteIssueParams(deleteIssueParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIssue``: %v\n", err)
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
    r, err := apiClient.DefaultApi.DeleteProject(context.Background()).DeleteProjectParams(deleteProjectParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteProject``: %v\n", err)
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
    r, err := apiClient.DefaultApi.DeleteSampleAnalysis(context.Background()).DeleteSamplesAnalysisParams(deleteSamplesAnalysisParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSampleAnalysis``: %v\n", err)
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
    r, err := apiClient.DefaultApi.DeleteSession(context.Background()).DeleteSessionParams(deleteSessionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSession``: %v\n", err)
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
    r, err := apiClient.DefaultApi.DeleteSessionRun(context.Background()).DeleteSessionRunParams(deleteSessionRunParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSessionRun``: %v\n", err)
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
    r, err := apiClient.DefaultApi.DeleteSessionTest(context.Background()).DeleteSessionTestRequest(deleteSessionTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSessionTest``: %v\n", err)
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
    r, err := apiClient.DefaultApi.DeleteTeam(context.Background()).DeleteTeamRequest(deleteTeamRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTeam``: %v\n", err)
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
    r, err := apiClient.DefaultApi.DeleteUserById(context.Background()).DeleteUserByIdRequest(deleteUserByIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUserById``: %v\n", err)
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
    r, err := apiClient.DefaultApi.DeleteVersion(context.Background()).DeleteVersionParams(deleteVersionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteVersion``: %v\n", err)
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
    r, err := apiClient.DefaultApi.DeleteVisualizations(context.Background()).DeleteVisualizationsParams(deleteVisualizationsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteVisualizations``: %v\n", err)
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
    r, err := apiClient.DefaultApi.DownloadProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DownloadProject``: %v\n", err)
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
    evaluateParams := *openapiclient.NewEvaluateParams("VersionId_example", "ProjectId_example", float64(123), []openapiclient.DataStateType{openapiclient.DataStateType("training")}, "Name_example", "Description_example", false, float64(123), "SessionId_example") // EvaluateParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Evaluate(context.Background()).EvaluateParams(evaluateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Evaluate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Evaluate`: Job
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Evaluate`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.ExportProject(context.Background()).ExportProjectRequest(exportProjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ExportProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportProject`: ExportProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ExportProject`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.ExtendTrial(context.Background()).ExtendTrialParams(extendTrialParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ExtendTrial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtendTrial`: ExtendTrialResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ExtendTrial`: %v\n", resp)
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

[jwt](../README.md#jwt)

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
    resp, r, err := apiClient.DefaultApi.FetchSimilar(context.Background()).FetchSimilarRequestParams(fetchSimilarRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSimilar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSimilar`: FetchSimilarResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSimilar`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetAllProjectSessionTests(context.Background()).GetAllProjectSessionTestsRequest(getAllProjectSessionTestsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllProjectSessionTests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllProjectSessionTests`: []SessionTest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllProjectSessionTests`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetAllSlimUserData(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllSlimUserData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSlimUserData`: GetTeamUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllSlimUserData`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetApiKeyByCode(context.Background()).GetApiKeyByCodeRequest(getApiKeyByCodeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApiKeyByCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiKeyByCode`: RotateApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApiKeyByCode`: %v\n", resp)
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


## GetArchiveInsights

> map[string][]float64 GetArchiveInsights(ctx).GetArchiveInsightsParams(getArchiveInsightsParams).Execute()



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
    getArchiveInsightsParams := *openapiclient.NewGetArchiveInsightsParams([]string{"PopExpDigests_example"}, "ProjectId_example") // GetArchiveInsightsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetArchiveInsights(context.Background()).GetArchiveInsightsParams(getArchiveInsightsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetArchiveInsights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArchiveInsights`: map[string][]float64
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetArchiveInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetArchiveInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getArchiveInsightsParams** | [**GetArchiveInsightsParams**](GetArchiveInsightsParams.md) |  | 

### Return type

[**map[string][]float64**](array.md)

### Authorization

[jwt](../README.md#jwt)

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
    resp, r, err := apiClient.DefaultApi.GetAuthProvider(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAuthProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthProvider`: GetAuthProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAuthProvider`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetBalancedAccuracy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetBalancedAccuracy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBalancedAccuracy`: MultiChartsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetBalancedAccuracy`: %v\n", resp)
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


## GetCodeIntegrationMappingErrorsByVersionId

> CodeIntegrationMappingErrorsResponse GetCodeIntegrationMappingErrorsByVersionId(ctx).CodeIntegrationMappingErrorsParams(codeIntegrationMappingErrorsParams).Execute()



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
    codeIntegrationMappingErrorsParams := *openapiclient.NewCodeIntegrationMappingErrorsParams("ProjectId_example", "VersionId_example", "Mapping_example") // CodeIntegrationMappingErrorsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetCodeIntegrationMappingErrorsByVersionId(context.Background()).CodeIntegrationMappingErrorsParams(codeIntegrationMappingErrorsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCodeIntegrationMappingErrorsByVersionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCodeIntegrationMappingErrorsByVersionId`: CodeIntegrationMappingErrorsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCodeIntegrationMappingErrorsByVersionId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCodeIntegrationMappingErrorsByVersionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codeIntegrationMappingErrorsParams** | [**CodeIntegrationMappingErrorsParams**](CodeIntegrationMappingErrorsParams.md) |  | 

### Return type

[**CodeIntegrationMappingErrorsResponse**](CodeIntegrationMappingErrorsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCodeIntegrationVersionFromBinder

> DatasetVersion GetCodeIntegrationVersionFromBinder(ctx).CodeIntegrationBinder(codeIntegrationBinder).Execute()



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
    codeIntegrationBinder := *openapiclient.NewCodeIntegrationBinder("CodeIntegrationId_example", "Branch_example", "Type_example", "CodeIntegrationVersionId_example") // CodeIntegrationBinder | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetCodeIntegrationVersionFromBinder(context.Background()).CodeIntegrationBinder(codeIntegrationBinder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCodeIntegrationVersionFromBinder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCodeIntegrationVersionFromBinder`: DatasetVersion
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCodeIntegrationVersionFromBinder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCodeIntegrationVersionFromBinderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codeIntegrationBinder** | [**CodeIntegrationBinder**](CodeIntegrationBinder.md) |  | 

### Return type

[**DatasetVersion**](DatasetVersion.md)

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
    resp, r, err := apiClient.DefaultApi.GetConfusionMatrixLabels(context.Background()).GetConfusionMatrixLabels(getConfusionMatrixLabels).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetConfusionMatrixLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfusionMatrixLabels`: ConfusionMatrixLabelsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetConfusionMatrixLabels`: %v\n", resp)
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
    getConfusionMatrixResultCombinationsParams := *openapiclient.NewGetConfusionMatrixResultCombinationsParams("ProjectId_example", []openapiclient.SessionRunToEpoch{*openapiclient.NewSessionRunToEpoch("SessionRunId_example", float64(123))}, *openapiclient.NewSplitAgg("Field_example", openapiclient.OrderType("asc"), "OrderField_example", float64(123), openapiclient.DistributionType("continuous")), "CustomMetricName_example") // GetConfusionMatrixResultCombinationsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetConfusionMatrixResultCombinations(context.Background()).GetConfusionMatrixResultCombinationsParams(getConfusionMatrixResultCombinationsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetConfusionMatrixResultCombinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfusionMatrixResultCombinations`: MultiChartsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetConfusionMatrixResultCombinations`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetConfusionMatrixTable(context.Background()).ConfusionMatrixTableParams(confusionMatrixTableParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetConfusionMatrixTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfusionMatrixTable`: MultiChartsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetConfusionMatrixTable`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetConfusionMetricNames(context.Background()).ConfusionMetricNamesParams(confusionMetricNamesParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetConfusionMetricNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfusionMetricNames`: ConfusionMetricNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetConfusionMetricNames`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetCurrentProjectVersion(context.Background()).GetCurrentProjectVersionParams(getCurrentProjectVersionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCurrentProjectVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentProjectVersion`: GetCurrentProjectVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCurrentProjectVersion`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetDashboard(context.Background()).GetDashboardParams(getDashboardParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboard`: GetDashboardResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDashboard`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetDashletFields(context.Background()).GetDashletFieldsParams(getDashletFieldsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDashletFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashletFields`: GetDashletFieldsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDashletFields`: %v\n", resp)
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


## GetDatasetVersion

> GetDatasetVersionResponse GetDatasetVersion(ctx).GetDatasetVersionParams(getDatasetVersionParams).Execute()



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
    getDatasetVersionParams := *openapiclient.NewGetDatasetVersionParams("DatasetVersionId_example") // GetDatasetVersionParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetDatasetVersion(context.Background()).GetDatasetVersionParams(getDatasetVersionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDatasetVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatasetVersion`: GetDatasetVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDatasetVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getDatasetVersionParams** | [**GetDatasetVersionParams**](GetDatasetVersionParams.md) |  | 

### Return type

[**GetDatasetVersionResponse**](GetDatasetVersionResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasetVersionUploadUrl

> DatasetVersionUploadUrlResponse GetDatasetVersionUploadUrl(ctx).GetDatasetVersionUploadUrlParams(getDatasetVersionUploadUrlParams).Execute()



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
    getDatasetVersionUploadUrlParams := *openapiclient.NewGetDatasetVersionUploadUrlParams("DatasetId_example") // GetDatasetVersionUploadUrlParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetDatasetVersionUploadUrl(context.Background()).GetDatasetVersionUploadUrlParams(getDatasetVersionUploadUrlParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDatasetVersionUploadUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatasetVersionUploadUrl`: DatasetVersionUploadUrlResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDatasetVersionUploadUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetVersionUploadUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getDatasetVersionUploadUrlParams** | [**GetDatasetVersionUploadUrlParams**](GetDatasetVersionUploadUrlParams.md) |  | 

### Return type

[**DatasetVersionUploadUrlResponse**](DatasetVersionUploadUrlResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasetVersions

> GetDatasetVersionsResponse GetDatasetVersions(ctx).GetDatasetVersionsParams(getDatasetVersionsParams).Execute()



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
    getDatasetVersionsParams := *openapiclient.NewGetDatasetVersionsParams("DatasetId_example") // GetDatasetVersionsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetDatasetVersions(context.Background()).GetDatasetVersionsParams(getDatasetVersionsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDatasetVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatasetVersions`: GetDatasetVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDatasetVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getDatasetVersionsParams** | [**GetDatasetVersionsParams**](GetDatasetVersionsParams.md) |  | 

### Return type

[**GetDatasetVersionsResponse**](GetDatasetVersionsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasets

> GetDatasetsResponse GetDatasets(ctx).Execute()



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
    resp, r, err := apiClient.DefaultApi.GetDatasets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDatasets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatasets`: GetDatasetsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDatasets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetsRequest struct via the builder pattern


### Return type

[**GetDatasetsResponse**](GetDatasetsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
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
    resp, r, err := apiClient.DefaultApi.GetDownloadSignedUrl(context.Background()).GetDownloadSignedUrlParams(getDownloadSignedUrlParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDownloadSignedUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDownloadSignedUrl`: GetDownloadSignedUrlResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDownloadSignedUrl`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetEngineSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEngineSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEngineSettings`: SettingsAndValuesWrapper
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEngineSettings`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetEnvironmentInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEnvironmentInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentInfo`: GetEnvironmentInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEnvironmentInfo`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetExportedSessionJobs(context.Background()).GetExportedSessionRunJobsParams(getExportedSessionRunJobsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetExportedSessionJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExportedSessionJobs`: GetExportedSessionRunJobsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetExportedSessionJobs`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetF1Score(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetF1Score``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetF1Score`: MultiChartsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetF1Score`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetFetchSimilarStatus(context.Background()).FetchSimilarRequestParams(fetchSimilarRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetFetchSimilarStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFetchSimilarStatus`: FetchSimilarResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetFetchSimilarStatus`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetFieldsValues(context.Background()).GetFieldsValuesRequest(getFieldsValuesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetFieldsValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFieldsValues`: GetFieldsValuesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetFieldsValues`: %v\n", resp)
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


## GetGeneratedDashlets

> GetGeneratedDashletsResponse GetGeneratedDashlets(ctx).GetGeneratedDashletsParams(getGeneratedDashletsParams).Execute()



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
    getGeneratedDashletsParams := *openapiclient.NewGetGeneratedDashletsParams("ProjectId_example", []string{"SessionRunIds_example"}) // GetGeneratedDashletsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetGeneratedDashlets(context.Background()).GetGeneratedDashletsParams(getGeneratedDashletsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetGeneratedDashlets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGeneratedDashlets`: GetGeneratedDashletsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetGeneratedDashlets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGeneratedDashletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getGeneratedDashletsParams** | [**GetGeneratedDashletsParams**](GetGeneratedDashletsParams.md) |  | 

### Return type

[**GetGeneratedDashletsResponse**](GetGeneratedDashletsResponse.md)

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
    resp, r, err := apiClient.DefaultApi.GetGenericBaseImageTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetGenericBaseImageTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGenericBaseImageTypes`: GetGenericBaseImageTypesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetGenericBaseImageTypes`: %v\n", resp)
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
    heatmapChartsParams := *openapiclient.NewHeatmapChartsParams("ProjectId_example", *openapiclient.NewSplitAgg("Field_example", openapiclient.OrderType("asc"), "OrderField_example", float64(123), openapiclient.DistributionType("continuous")), *openapiclient.NewSplitAgg("Field_example", openapiclient.OrderType("asc"), "OrderField_example", float64(123), openapiclient.DistributionType("continuous")), *openapiclient.NewAggregations("Field_example", openapiclient.AggregationMethod("Average")), []openapiclient.SessionRunToEpoch{*openapiclient.NewSessionRunToEpoch("SessionRunId_example", float64(123))}, false) // HeatmapChartsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetHeatmapChart(context.Background()).HeatmapChartsParams(heatmapChartsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetHeatmapChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHeatmapChart`: MultiChartsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetHeatmapChart`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetIssueFileUploadSignedUrl(context.Background()).GetIssueFileUploadSignedUrl(getIssueFileUploadSignedUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetIssueFileUploadSignedUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIssueFileUploadSignedUrl`: IssueFileUploadSignedUrl
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetIssueFileUploadSignedUrl`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetJobLogs(context.Background()).GetJobLogsParams(getJobLogsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetJobLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobLogs`: GetJobLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetJobLogs`: %v\n", resp)
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


## GetLatestDatasetVersion

> GetLatestDatasetVersionResponse GetLatestDatasetVersion(ctx).GetLatestDatasetVersionParams(getLatestDatasetVersionParams).Execute()



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
    getLatestDatasetVersionParams := *openapiclient.NewGetLatestDatasetVersionParams("DatasetId_example") // GetLatestDatasetVersionParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetLatestDatasetVersion(context.Background()).GetLatestDatasetVersionParams(getLatestDatasetVersionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetLatestDatasetVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestDatasetVersion`: GetLatestDatasetVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetLatestDatasetVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestDatasetVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getLatestDatasetVersionParams** | [**GetLatestDatasetVersionParams**](GetLatestDatasetVersionParams.md) |  | 

### Return type

[**GetLatestDatasetVersionResponse**](GetLatestDatasetVersionResponse.md)

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
    resp, r, err := apiClient.DefaultApi.GetMachineTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMachineTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMachineTypes`: GetMachineTypesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMachineTypes`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetMaxActiveUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMaxActiveUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaxActiveUsers`: GetMaxActiveUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMaxActiveUsers`: %v\n", resp)
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
    confusionMatrixParams := *openapiclient.NewConfusionMatrixParams("ProjectId_example", []openapiclient.SessionRunToEpoch{*openapiclient.NewSessionRunToEpoch("SessionRunId_example", float64(123))}, *openapiclient.NewSplitAgg("Field_example", openapiclient.OrderType("asc"), "OrderField_example", float64(123), openapiclient.DistributionType("continuous")), "CustomMetricName_example") // ConfusionMatrixParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetMeanAveragePrecision(context.Background()).ConfusionMatrixParams(confusionMatrixParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMeanAveragePrecision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMeanAveragePrecision`: MultiChartsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMeanAveragePrecision`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetNotifications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotifications`: GetNotificationsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetNotifications`: %v\n", resp)
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
    populationExplorationParams := *openapiclient.NewPopulationExplorationParams("SessionRunId_example", "ProjectId_example", float64(123), float64(123), "Digest_example", *openapiclient.NewPopulationExplorationDisplayParams(float64(123), []string{"BalanceBy_example"}, false, openapiclient.ReductionAlgorithm("TSNE"))) // PopulationExplorationParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetPopulationExplorationStatus(context.Background()).PopulationExplorationParams(populationExplorationParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPopulationExplorationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPopulationExplorationStatus`: PopulationExplorationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPopulationExplorationStatus`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetPrCurve(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPrCurve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrCurve`: MultiChartsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPrCurve`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetPrecisionScore(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPrecisionScore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrecisionScore`: MultiChartsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPrecisionScore`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetProjectDashboards(context.Background()).GetProjectDashboardsParams(getProjectDashboardsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProjectDashboards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectDashboards`: GetProjectDashboardsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetProjectDashboards`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetProjectIssues(context.Background()).GetProjectIssuesParams(getProjectIssuesParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProjectIssues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectIssues`: GetProjectIssuesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetProjectIssues`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetProjectSlimVersions(context.Background()).GetProjectVersionsParams(getProjectVersionsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProjectSlimVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectSlimVersions`: GetProjectSlimVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetProjectSlimVersions`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetProjects(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjects`: GetProjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetProjects`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetRecallScore(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRecallScore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecallScore`: MultiChartsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRecallScore`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetRecentTeamSessions(context.Background()).RecentTeamSessionsRequestParams(recentTeamSessionsRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRecentTeamSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecentTeamSessions`: RecentSessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRecentTeamSessions`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetRoc(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRoc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoc`: MultiChartsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRoc`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetSampleCollections(context.Background()).GetSampleCollectionsParams(getSampleCollectionsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSampleCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSampleCollections`: GetSampleCollectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSampleCollections`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetSampleVisualizationsPath(context.Background()).GetSampleVisualizationsPathsParams(getSampleVisualizationsPathsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSampleVisualizationsPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSampleVisualizationsPath`: GetSampleVisualizationsPathsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSampleVisualizationsPath`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetScatterSampleVisualizations(context.Background()).GetScatterSampleVisualizationsParams(getScatterSampleVisualizationsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetScatterSampleVisualizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScatterSampleVisualizations`: GetScatterSampleVisualizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetScatterSampleVisualizations`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetSecretManagerList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSecretManagerList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretManagerList`: GetSecretManagerListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSecretManagerList`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetSessionRunsVisualizations(context.Background()).GetSessionRunsVisualizationsParams(getSessionRunsVisualizationsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSessionRunsVisualizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionRunsVisualizations`: GetSessionRunsVisualizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSessionRunsVisualizations`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetSessionTestResult(context.Background()).GetSessionTestResultsRequest(getSessionTestResultsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSessionTestResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionTestResult`: []AllSessionsTestResults
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSessionTestResult`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetSessionsByHash(context.Background()).SessionHashRequestParams(sessionHashRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSessionsByHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionsByHash`: SessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSessionsByHash`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetSessionsByVersionId(context.Background()).SessionVersionIdRequestParams(sessionVersionIdRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSessionsByVersionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionsByVersionId`: SessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSessionsByVersionId`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetSessionsEpochs(context.Background()).GetSessionsEpochsRequest(getSessionsEpochsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSessionsEpochs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionsEpochs`: GetSessionEpochsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSessionsEpochs`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetSignedUrl(context.Background()).GetSignedUrlParams(getSignedUrlParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSignedUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignedUrl`: ExternalImportModelStorage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSignedUrl`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetSingleIssue(context.Background()).GetSingleIssueParams(getSingleIssueParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSingleIssue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleIssue`: Issue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSingleIssue`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetSingleSessionTest(context.Background()).GetSingleSessionTestRequest(getSingleSessionTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSingleSessionTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleSessionTest`: SessionTest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSingleSessionTest`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetSlimJobs(context.Background()).GetJobsFilterParams(getJobsFilterParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSlimJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlimJobs`: GetSlimJobsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSlimJobs`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetSlimVisualization(context.Background()).GetSlimVisualizationParams(getSlimVisualizationParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSlimVisualization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlimVisualization`: GetSlimVisualizationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSlimVisualization`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetState(context.Background()).GetStateParams(getStateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetState`: GetStateResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetState`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetStatistics(context.Background()).GetStatisticsParams(getStatisticsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatistics`: GetStatisticsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStatistics`: %v\n", resp)
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
    genericDataQueryParams := *openapiclient.NewGenericDataQueryParams("ProjectId_example", []openapiclient.SessionRunToEpoch{*openapiclient.NewSessionRunToEpoch("SessionRunId_example", float64(123))}, false, []openapiclient.Aggregations{*openapiclient.NewAggregations("Field_example", openapiclient.AggregationMethod("Average"))}, []openapiclient.SplitAgg{*openapiclient.NewSplitAgg("Field_example", openapiclient.OrderType("asc"), "OrderField_example", float64(123), openapiclient.DistributionType("continuous"))}) // GenericDataQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetTableChart(context.Background()).GenericDataQueryParams(genericDataQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTableChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTableChart`: MultiChartsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTableChart`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetTeamJobs(context.Background()).GetJobsFilterParams(getJobsFilterParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTeamJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamJobs`: GetJobsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTeamJobs`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetTeamSlimUserData(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTeamSlimUserData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamSlimUserData`: GetTeamUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTeamSlimUserData`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetTeams(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeams`: GetTeamsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTeams`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetUploadModelSignedUrl(context.Background()).GetUploadModelSignedUrlRequest(getUploadModelSignedUrlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUploadModelSignedUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUploadModelSignedUrl`: ExternalImportModelStorage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUploadModelSignedUrl`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.GetUploadSignedUrl(context.Background()).GetUploadSignedUrlParams(getUploadSignedUrlParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUploadSignedUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUploadSignedUrl`: ExternalImportModelStorage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUploadSignedUrl`: %v\n", resp)
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


## GetValidateGraphProcessState

> ValidateGraphResponse GetValidateGraphProcessState(ctx).GetValidateGraphProcessStateParams(getValidateGraphProcessStateParams).Execute()



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
    getValidateGraphProcessStateParams := *openapiclient.NewGetValidateGraphProcessStateParams("ProjectId_example", "Digest_example") // GetValidateGraphProcessStateParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetValidateGraphProcessState(context.Background()).GetValidateGraphProcessStateParams(getValidateGraphProcessStateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetValidateGraphProcessState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetValidateGraphProcessState`: ValidateGraphResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetValidateGraphProcessState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetValidateGraphProcessStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getValidateGraphProcessStateParams** | [**GetValidateGraphProcessStateParams**](GetValidateGraphProcessStateParams.md) |  | 

### Return type

[**ValidateGraphResponse**](ValidateGraphResponse.md)

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
    resp, r, err := apiClient.DefaultApi.GetVisualization(context.Background()).GetSlimVisualizationParams(getSlimVisualizationParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVisualization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVisualization`: Visualization
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVisualization`: %v\n", resp)
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
    multiChartsParams := *openapiclient.NewMultiChartsParams("ProjectId_example", *openapiclient.NewSplitAgg("Field_example", openapiclient.OrderType("asc"), "OrderField_example", float64(123), openapiclient.DistributionType("continuous")), *openapiclient.NewAggregations("Field_example", openapiclient.AggregationMethod("Average")), []openapiclient.SessionRunToEpoch{*openapiclient.NewSessionRunToEpoch("SessionRunId_example", float64(123))}, false) // MultiChartsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetXYChart(context.Background()).MultiChartsParams(multiChartsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetXYChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetXYChart`: MultiChartsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetXYChart`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.HealthCheck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.HealthCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HealthCheck`: HealthCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.HealthCheck`: %v\n", resp)
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

> ExternalImportModelStorageResponse ImportExternalModel(ctx).ImportExternalModelParams(importExternalModelParams).Execute()



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
    resp, r, err := apiClient.DefaultApi.ImportExternalModel(context.Background()).ImportExternalModelParams(importExternalModelParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ImportExternalModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportExternalModel`: ExternalImportModelStorageResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ImportExternalModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportExternalModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importExternalModelParams** | [**ImportExternalModelParams**](ImportExternalModelParams.md) |  | 

### Return type

[**ExternalImportModelStorageResponse**](ExternalImportModelStorageResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportModel

> ExternalImportModelStorageResponse ImportModel(ctx).ImportNewModelParams(importNewModelParams).Execute()



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
    importNewModelParams := *openapiclient.NewImportNewModelParams("ProjectId_example", "FileName_example", "ModelName_example", "VersionName_example", openapiclient.ImportModelType("JSON_TF2")) // ImportNewModelParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ImportModel(context.Background()).ImportNewModelParams(importNewModelParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ImportModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportModel`: ExternalImportModelStorageResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ImportModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importNewModelParams** | [**ImportNewModelParams**](ImportNewModelParams.md) |  | 

### Return type

[**ExternalImportModelStorageResponse**](ExternalImportModelStorageResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportProject

> ImportProject(ctx).ImportProjectRequest(importProjectRequest).Execute()



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
    r, err := apiClient.DefaultApi.ImportProject(context.Background()).ImportProjectRequest(importProjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ImportProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importProjectRequest** | [**ImportProjectRequest**](ImportProjectRequest.md) |  | 

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
    initExperimentRequest := *openapiclient.NewInitExperimentRequest("ExperimentName_example") // InitExperimentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.InitExperiment(context.Background()).InitExperimentRequest(initExperimentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InitExperiment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitExperiment`: InitExperimentResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.InitExperiment`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.KeyGen(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.KeyGen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyGen`: RotateApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.KeyGen`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.LoadModel(context.Background()).LoadSessionParams(loadSessionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.LoadModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadModel`: LoadSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.LoadModel`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.LoadVersion(context.Background()).LoadVersionParams(loadVersionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.LoadVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadVersion`: LoadVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.LoadVersion`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.LocalAuth(context.Background()).LocalLoginParams(localLoginParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.LocalAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LocalAuth`: LocalLoginResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.LocalAuth`: %v\n", resp)
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
    r, err := apiClient.DefaultApi.LogExternalEpochData(context.Background()).LogExternalEpochDataRequest(logExternalEpochDataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.LogExternalEpochData``: %v\n", err)
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
    resp, r, err := apiClient.DefaultApi.Login(context.Background()).LoginParams(loginParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Login``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Login`: UserData
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Login`: %v\n", resp)
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
    r, err := apiClient.DefaultApi.Logout(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Logout``: %v\n", err)
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


## ModifyDatasetVersionNote

> ModifyDatasetVersionNoteResponse ModifyDatasetVersionNote(ctx).ModifyDatasetVersionNoteParams(modifyDatasetVersionNoteParams).Execute()



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
    modifyDatasetVersionNoteParams := *openapiclient.NewModifyDatasetVersionNoteParams("DatasetVersionId_example", "Note_example") // ModifyDatasetVersionNoteParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ModifyDatasetVersionNote(context.Background()).ModifyDatasetVersionNoteParams(modifyDatasetVersionNoteParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyDatasetVersionNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyDatasetVersionNote`: ModifyDatasetVersionNoteResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyDatasetVersionNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyDatasetVersionNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modifyDatasetVersionNoteParams** | [**ModifyDatasetVersionNoteParams**](ModifyDatasetVersionNoteParams.md) |  | 

### Return type

[**ModifyDatasetVersionNoteResponse**](ModifyDatasetVersionNoteResponse.md)

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
    populationExplorationParams := *openapiclient.NewPopulationExplorationParams("SessionRunId_example", "ProjectId_example", float64(123), float64(123), "Digest_example", *openapiclient.NewPopulationExplorationDisplayParams(float64(123), []string{"BalanceBy_example"}, false, openapiclient.ReductionAlgorithm("TSNE"))) // PopulationExplorationParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PopulationExploration(context.Background()).PopulationExplorationParams(populationExplorationParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PopulationExploration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PopulationExploration`: PopulationExplorationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PopulationExploration`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.RefreshLocalAuth(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RefreshLocalAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshLocalAuth`: LocalLoginResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RefreshLocalAuth`: %v\n", resp)
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
    r, err := apiClient.DefaultApi.RemoveSamplesCollection(context.Background()).RemoveSampleCollectionParams(removeSampleCollectionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RemoveSamplesCollection``: %v\n", err)
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
    r, err := apiClient.DefaultApi.ResolveConcurrentUsersConflict(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ResolveConcurrentUsersConflict``: %v\n", err)
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
    resp, r, err := apiClient.DefaultApi.SampleAnalysis(context.Background()).SampleAnalysisParams(sampleAnalysisParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SampleAnalysis``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SampleAnalysis`: Job
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SampleAnalysis`: %v\n", resp)
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
    r, err := apiClient.DefaultApi.SamplesVisualizationsRefresh(context.Background()).SamplesVisualizationsRefreshParams(samplesVisualizationsRefreshParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SamplesVisualizationsRefresh``: %v\n", err)
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
    r, err := apiClient.DefaultApi.SaveAnalyzerLayout(context.Background()).SaveAnalyzerLayoutParams(saveAnalyzerLayoutParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SaveAnalyzerLayout``: %v\n", err)
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


## SaveDatasetVersion

> SaveDatasetSetupResponse SaveDatasetVersion(ctx).SaveDatasetVersionParams(saveDatasetVersionParams).Execute()



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
    saveDatasetVersionParams := *openapiclient.NewSaveDatasetVersionParams("DatasetId_example", "CodeUrl_example", "CodeEntryFile_example") // SaveDatasetVersionParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SaveDatasetVersion(context.Background()).SaveDatasetVersionParams(saveDatasetVersionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SaveDatasetVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveDatasetVersion`: SaveDatasetSetupResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SaveDatasetVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveDatasetVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveDatasetVersionParams** | [**SaveDatasetVersionParams**](SaveDatasetVersionParams.md) |  | 

### Return type

[**SaveDatasetSetupResponse**](SaveDatasetSetupResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveDatasetVersionNoParse

> SaveDatasetSetupResponse SaveDatasetVersionNoParse(ctx).SaveDatasetVersionNoParseParams(saveDatasetVersionNoParseParams).Execute()



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
    saveDatasetVersionNoParseParams := *openapiclient.NewSaveDatasetVersionNoParseParams("FromDatasetVersionId_example", "CodeUrl_example", "CodeEntryFile_example") // SaveDatasetVersionNoParseParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SaveDatasetVersionNoParse(context.Background()).SaveDatasetVersionNoParseParams(saveDatasetVersionNoParseParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SaveDatasetVersionNoParse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveDatasetVersionNoParse`: SaveDatasetSetupResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SaveDatasetVersionNoParse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveDatasetVersionNoParseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveDatasetVersionNoParseParams** | [**SaveDatasetVersionNoParseParams**](SaveDatasetVersionNoParseParams.md) |  | 

### Return type

[**SaveDatasetSetupResponse**](SaveDatasetSetupResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
    resp, r, err := apiClient.DefaultApi.SendUserMessage(context.Background()).SendUserMessageParams(sendUserMessageParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SendUserMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendUserMessage`: SendUserMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SendUserMessage`: %v\n", resp)
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
    r, err := apiClient.DefaultApi.SetCodeChallenge(context.Background()).SetCodeChallengeRequest(setCodeChallengeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetCodeChallenge``: %v\n", err)
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
    r, err := apiClient.DefaultApi.SetDefaultTeam(context.Background()).SetDefaultTeamRequest(setDefaultTeamRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetDefaultTeam``: %v\n", err)
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
    r, err := apiClient.DefaultApi.SetExperimentProperties(context.Background()).SetExperimentPropertiesRequest(setExperimentPropertiesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetExperimentProperties``: %v\n", err)
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
    r, err := apiClient.DefaultApi.SetMachineType(context.Background()).SetTeamMachineTypeParams(setTeamMachineTypeParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetMachineType``: %v\n", err)
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
    resp, r, err := apiClient.DefaultApi.SetUserNotificationsAsRead(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetUserNotificationsAsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetUserNotificationsAsRead`: SetUserNotificationsAsRead200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SetUserNotificationsAsRead`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.StartTrial(context.Background()).StartTrialParams(startTrialParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.StartTrial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartTrial`: UserData
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.StartTrial`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.StopJob(context.Background()).StopJobParams(stopJobParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.StopJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopJob`: StopJobResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.StopJob`: %v\n", resp)
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
    r, err := apiClient.DefaultApi.TagModel(context.Background()).TagModelRequest(tagModelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TagModel``: %v\n", err)
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
    resp, r, err := apiClient.DefaultApi.TerminateAllJobs(context.Background()).TerminateAllJobsParams(terminateAllJobsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TerminateAllJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TerminateAllJobs`: TerminateAllJobsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TerminateAllJobs`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.TerminateJob(context.Background()).TerminateJobParams(terminateJobParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TerminateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TerminateJob`: TerminateJobResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TerminateJob`: %v\n", resp)
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


## TrashDataset

> GetDatasetsResponse TrashDataset(ctx).TrashDatasetParams(trashDatasetParams).Execute()



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
    trashDatasetParams := *openapiclient.NewTrashDatasetParams("DatasetId_example") // TrashDatasetParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TrashDataset(context.Background()).TrashDatasetParams(trashDatasetParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TrashDataset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TrashDataset`: GetDatasetsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TrashDataset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrashDatasetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trashDatasetParams** | [**TrashDatasetParams**](TrashDatasetParams.md) |  | 

### Return type

[**GetDatasetsResponse**](GetDatasetsResponse.md)

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
    resp, r, err := apiClient.DefaultApi.TrashSecretManager(context.Background()).TrashSecretManagerParams(trashSecretManagerParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TrashSecretManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TrashSecretManager`: TrashSecretManagerResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TrashSecretManager`: %v\n", resp)
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


## UnarchiveInsight

> UnarchiveInsight(ctx).Body(body).Execute()



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
    body := ArchiveInsightsParams(987) // ArchiveInsightsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.UnarchiveInsight(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UnarchiveInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **ArchiveInsightsParams** |  | 

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
    updateDashboardParams := *openapiclient.NewUpdateDashboardParams("DashboardId_example", "Name_example", []openapiclient.Dashlet{*openapiclient.NewDashlet("Cid_example", *openapiclient.NewSizedLayout(*openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123))), "Type_example", *openapiclient.NewPopulationExplorationDashletData(map[string]interface{}(123), "Type_example", "Name_example"), "Name_example", []string{"CollectionIds_example"})}, "ProjectId_example") // UpdateDashboardParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.UpdateDashboard(context.Background()).UpdateDashboardParams(updateDashboardParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateDashboard``: %v\n", err)
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
    r, err := apiClient.DefaultApi.UpdateEngineSettings(context.Background()).SetSettingValueWrapper(setSettingValueWrapper).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateEngineSettings``: %v\n", err)
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
    resp, r, err := apiClient.DefaultApi.UpdateIssue(context.Background()).UpdateIssueParams(updateIssueParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateIssue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIssue`: Issue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateIssue`: %v\n", resp)
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
    r, err := apiClient.DefaultApi.UpdateProjectMeta(context.Background()).UpdateProjectMetaRequest(updateProjectMetaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateProjectMeta``: %v\n", err)
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
    r, err := apiClient.DefaultApi.UpdateSampleCollection(context.Background()).UpdateSampleCollectionParams(updateSampleCollectionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSampleCollection``: %v\n", err)
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
    resp, r, err := apiClient.DefaultApi.UpdateSecretManager(context.Background()).UpdateSecretManagerParams(updateSecretManagerParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSecretManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecretManager`: UpdateSecretManagerResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSecretManager`: %v\n", resp)
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
    r, err := apiClient.DefaultApi.UpdateSessionName(context.Background()).UpdateSessionNameParams(updateSessionNameParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSessionName``: %v\n", err)
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
    r, err := apiClient.DefaultApi.UpdateSessionRun(context.Background()).UpdateSessionRunNameParams(updateSessionRunNameParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSessionRun``: %v\n", err)
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
    r, err := apiClient.DefaultApi.UpdateSessionTest(context.Background()).UpdateSessionTestRequest(updateSessionTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSessionTest``: %v\n", err)
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
    r, err := apiClient.DefaultApi.UpdateTeamPublicName(context.Background()).UpdateTeamPublicNameRequest(updateTeamPublicNameRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTeamPublicName``: %v\n", err)
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
    resp, r, err := apiClient.DefaultApi.UpdateUserName(context.Background()).UpdateUserNameRequest(updateUserNameRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUserName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserName`: UserData
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateUserName`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.UpdateUserRole(context.Background()).UpdateUserRoleRequest(updateUserRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUserRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserRole`: SlimUserData
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateUserRole`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.UpdateUserStatus(context.Background()).UpdateUserStatusRequest(updateUserStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUserStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserStatus`: SlimUserData
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateUserStatus`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.UpdateUserTeam(context.Background()).UpdateUserTeamRequest(updateUserTeamRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUserTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserTeam`: SlimUserData
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateUserTeam`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultApi.UpdateVersion(context.Background()).UpdateVersionParams(updateVersionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVersion`: UpdateVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateVersion`: %v\n", resp)
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
    r, err := apiClient.DefaultApi.UpdateVersionName(context.Background()).UpdateVersionNameParams(updateVersionNameParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateVersionName``: %v\n", err)
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
    r, err := apiClient.DefaultApi.UploadProject(context.Background(), projectName).Description(description).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UploadProject``: %v\n", err)
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
    resp, r, err := apiClient.DefaultApi.UpsertState(context.Background()).UpsertStateParams(upsertStateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpsertState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpsertState`: UpsertStateResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpsertState`: %v\n", resp)
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


## ValidateGraph

> ValidateGraphResponse ValidateGraph(ctx).ValidateGraphParams(validateGraphParams).Execute()



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
    validateGraphParams := *openapiclient.NewValidateGraphParams(*openapiclient.NewModelGraph("Id_example", map[string]interface{}(123)), "DatasetVersionId_example", "VersionId_example", "ProjectId_example", "Digest_example") // ValidateGraphParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ValidateGraph(context.Background()).ValidateGraphParams(validateGraphParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ValidateGraph``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateGraph`: ValidateGraphResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ValidateGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateGraphParams** | [**ValidateGraphParams**](ValidateGraphParams.md) |  | 

### Return type

[**ValidateGraphResponse**](ValidateGraphResponse.md)

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
    r, err := apiClient.DefaultApi.Warmup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Warmup``: %v\n", err)
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
    resp, r, err := apiClient.DefaultApi.WhoAmI(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.WhoAmI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WhoAmI`: UserData
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.WhoAmI`: %v\n", resp)
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

