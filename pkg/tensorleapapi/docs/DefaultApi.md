# \DefaultApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Activate**](DefaultApi.md#Activate) | **Post** /auth/activate | 
[**AddDashboard**](DefaultApi.md#AddDashboard) | **Post** /dashboards/addDashboard | 
[**AddDataset**](DefaultApi.md#AddDataset) | **Post** /datasets/addDataset | 
[**AddExportModelJob**](DefaultApi.md#AddExportModelJob) | **Post** /jobs/addExportModelJob | 
[**AddIssue**](DefaultApi.md#AddIssue) | **Post** /issues/addIssue | 
[**AddProject**](DefaultApi.md#AddProject) | **Post** /projects/addProject | 
[**AddSecretManager**](DefaultApi.md#AddSecretManager) | **Post** /secret-manager/addSecretManager | 
[**AddVersion**](DefaultApi.md#AddVersion) | **Post** /versions/addVersion | 
[**ContinueTrain**](DefaultApi.md#ContinueTrain) | **Post** /jobs/continueTrain | 
[**CreateOrganization**](DefaultApi.md#CreateOrganization) | **Post** /organizations/createOrganization | 
[**CreateSessionTest**](DefaultApi.md#CreateSessionTest) | **Post** /sessions-tests/createSessionTest | 
[**DeleteDashboard**](DefaultApi.md#DeleteDashboard) | **Post** /dashboards/deleteDashboard | 
[**DeleteIssue**](DefaultApi.md#DeleteIssue) | **Post** /issues/deleteIssue | 
[**DeleteOrganization**](DefaultApi.md#DeleteOrganization) | **Post** /organizations/deleteOrganization | 
[**DeleteProject**](DefaultApi.md#DeleteProject) | **Post** /projects/deleteProject | 
[**DeleteSession**](DefaultApi.md#DeleteSession) | **Post** /sessions/deleteSession | 
[**DeleteSessionRun**](DefaultApi.md#DeleteSessionRun) | **Post** /sessions/deleteSessionRun | 
[**DeleteSessionTest**](DefaultApi.md#DeleteSessionTest) | **Post** /sessions-tests/deleteSessionTest | 
[**DeleteUserById**](DefaultApi.md#DeleteUserById) | **Post** /users/deleteUserById | 
[**DeleteVersion**](DefaultApi.md#DeleteVersion) | **Post** /versions/deleteVersion | 
[**DeleteVisualizations**](DefaultApi.md#DeleteVisualizations) | **Post** /visualizations/deleteVisualizations | 
[**Evaluate**](DefaultApi.md#Evaluate) | **Post** /jobs/evaluate | 
[**ExtendTrial**](DefaultApi.md#ExtendTrial) | **Post** /auth/extendTrial | 
[**GetAllProjectSessionTests**](DefaultApi.md#GetAllProjectSessionTests) | **Post** /sessions-tests/getAllProjectSessionTests | 
[**GetAllSlimUserData**](DefaultApi.md#GetAllSlimUserData) | **Post** /users/getAllSlimUserData | 
[**GetBalancedAccuracy**](DefaultApi.md#GetBalancedAccuracy) | **Post** /sessionmetrics/getBalancedAccuracy | 
[**GetConfusionMetricNames**](DefaultApi.md#GetConfusionMetricNames) | **Post** /sessionmetrics/getConfusionMetricNames | 
[**GetCurrentProjectVersion**](DefaultApi.md#GetCurrentProjectVersion) | **Post** /projects/getCurrentProjectVersion | 
[**GetDashboard**](DefaultApi.md#GetDashboard) | **Post** /dashboards/getDashboard | 
[**GetDashletFields**](DefaultApi.md#GetDashletFields) | **Post** /dashboards/getDashletFields | 
[**GetDatasetVersion**](DefaultApi.md#GetDatasetVersion) | **Post** /datasets/getDatasetVersion | 
[**GetDatasetVersionUploadUrl**](DefaultApi.md#GetDatasetVersionUploadUrl) | **Post** /datasets/getDatasetVersionUploadUrl | 
[**GetDatasetVersions**](DefaultApi.md#GetDatasetVersions) | **Post** /datasets/getDatasetVersions | 
[**GetDatasets**](DefaultApi.md#GetDatasets) | **Post** /datasets/getDatasets | 
[**GetExportedSessionJobs**](DefaultApi.md#GetExportedSessionJobs) | **Post** /exportedsessionruns/getExportedSessionRunJobs | 
[**GetF1Score**](DefaultApi.md#GetF1Score) | **Post** /sessionmetrics/getF1Score | 
[**GetHeatmapChart**](DefaultApi.md#GetHeatmapChart) | **Post** /sessionmetrics/getHeatmapChart | 
[**GetJobs**](DefaultApi.md#GetJobs) | **Post** /jobs/getJobs | 
[**GetLatestDatasetVersion**](DefaultApi.md#GetLatestDatasetVersion) | **Post** /datasets/getLatestDatasetVersion | 
[**GetMachineTypes**](DefaultApi.md#GetMachineTypes) | **Post** /jobs/getMachineTypes | 
[**GetMaxActiveUsers**](DefaultApi.md#GetMaxActiveUsers) | **Post** /metadata/getMaxActiveUsers | 
[**GetNotifications**](DefaultApi.md#GetNotifications) | **Post** /notifications/getNotifications | 
[**GetOrganizationJobs**](DefaultApi.md#GetOrganizationJobs) | **Post** /jobs/getOrganizationJobs | 
[**GetOrganizationSlimUserData**](DefaultApi.md#GetOrganizationSlimUserData) | **Post** /users/getOrganizationSlimUserData | 
[**GetOrganizations**](DefaultApi.md#GetOrganizations) | **Post** /organizations/getOrganizations | 
[**GetPrCurve**](DefaultApi.md#GetPrCurve) | **Post** /sessionmetrics/getPrCurve | 
[**GetProjectDashboards**](DefaultApi.md#GetProjectDashboards) | **Post** /dashboards/getProjectDashboards | 
[**GetProjectIssues**](DefaultApi.md#GetProjectIssues) | **Post** /issues/getProjectIssues | 
[**GetProjectSlimVersions**](DefaultApi.md#GetProjectSlimVersions) | **Post** /versions/getProjectSlimVersions | 
[**GetProjects**](DefaultApi.md#GetProjects) | **Post** /projects/getProjects | 
[**GetRecentOrganizationSessions**](DefaultApi.md#GetRecentOrganizationSessions) | **Post** /sessions/getRecentOrganizationSessions | 
[**GetRoc**](DefaultApi.md#GetRoc) | **Post** /sessionmetrics/getRoc | 
[**GetSecretManagerList**](DefaultApi.md#GetSecretManagerList) | **Post** /secret-manager/getSecretManagerList | 
[**GetSessionInsights**](DefaultApi.md#GetSessionInsights) | **Post** /insights/getSessionInsights | 
[**GetSessionTestResult**](DefaultApi.md#GetSessionTestResult) | **Post** /sessions-tests/getSessionTestResult | 
[**GetSessionVisualizations**](DefaultApi.md#GetSessionVisualizations) | **Post** /visualizations/getSessionVisualizations | 
[**GetSessionsByHash**](DefaultApi.md#GetSessionsByHash) | **Post** /sessions/getSessionsByHash | 
[**GetSessionsByVersionId**](DefaultApi.md#GetSessionsByVersionId) | **Post** /sessions/getSessionsByVersionId | 
[**GetSignedStoredResourceUrl**](DefaultApi.md#GetSignedStoredResourceUrl) | **Post** /visualizations/getSignedStoredResourceUrl | 
[**GetSignedStoredUrlByVisualizationId**](DefaultApi.md#GetSignedStoredUrlByVisualizationId) | **Post** /visualizations/getSignedStoredUrlByVisualizationId | 
[**GetSingleIssue**](DefaultApi.md#GetSingleIssue) | **Post** /issues/getSingleIssue | 
[**GetSingleSessionTest**](DefaultApi.md#GetSingleSessionTest) | **Post** /sessions-tests/getSingleSessionTest | 
[**GetStatistics**](DefaultApi.md#GetStatistics) | **Post** /metadata/getStatistics | 
[**GetStoredExportedSessionRunResourceUrl**](DefaultApi.md#GetStoredExportedSessionRunResourceUrl) | **Post** /exportedsessionruns/getStoredExportedSessionRunResourceUrl | 
[**GetTableChart**](DefaultApi.md#GetTableChart) | **Post** /sessionmetrics/getTableChart | 
[**GetTrainingJobs**](DefaultApi.md#GetTrainingJobs) | **Post** /jobs/getTrainingJobs | 
[**GetUploadSignedUrl**](DefaultApi.md#GetUploadSignedUrl) | **Post** /versions/getUploadSignedUrl | 
[**GetVisualization**](DefaultApi.md#GetVisualization) | **Post** /visualizations/getVisualization | 
[**GetXYChart**](DefaultApi.md#GetXYChart) | **Post** /sessionmetrics/getXYChart | 
[**HealthCheck**](DefaultApi.md#HealthCheck) | **Get** /monitor/healthCheck | 
[**ImportModel**](DefaultApi.md#ImportModel) | **Post** /versions/importModel | 
[**IsTrainingJobRunning**](DefaultApi.md#IsTrainingJobRunning) | **Post** /jobs/isTrainingJobRunning | 
[**KeyGen**](DefaultApi.md#KeyGen) | **Post** /auth/keygen | 
[**KeyLogin**](DefaultApi.md#KeyLogin) | **Post** /auth/keylogin | 
[**LoadModel**](DefaultApi.md#LoadModel) | **Post** /projects/loadModel | 
[**LoadVersion**](DefaultApi.md#LoadVersion) | **Post** /versions/loadVersion | 
[**Login**](DefaultApi.md#Login) | **Post** /auth/login | 
[**Logout**](DefaultApi.md#Logout) | **Post** /auth/logout | 
[**ModifyDatasetVersionNote**](DefaultApi.md#ModifyDatasetVersionNote) | **Post** /datasets/modifyDatasetVersionNote | 
[**PopulationExploration**](DefaultApi.md#PopulationExploration) | **Post** /jobs/populationExploration | 
[**SampleAnalysis**](DefaultApi.md#SampleAnalysis) | **Post** /jobs/sampleAnalysis | 
[**SampleSelection**](DefaultApi.md#SampleSelection) | **Post** /jobs/sampleSelection | 
[**SaveAnalyzerLayout**](DefaultApi.md#SaveAnalyzerLayout) | **Post** /visualizations/saveAnalyzerLayout | 
[**SaveDatasetVersion**](DefaultApi.md#SaveDatasetVersion) | **Post** /datasets/saveDatasetVersion | 
[**SendUserMessage**](DefaultApi.md#SendUserMessage) | **Post** /users/sendUserMessage | 
[**SetDefaultOrganization**](DefaultApi.md#SetDefaultOrganization) | **Post** /organizations/setDefaultOrganization | 
[**SetOrganizationMachineType**](DefaultApi.md#SetOrganizationMachineType) | **Post** /jobs/setMachineType | 
[**SetUserNotificationsAsRead**](DefaultApi.md#SetUserNotificationsAsRead) | **Post** /notifications/setUserNotificationsAsRead | 
[**Signup**](DefaultApi.md#Signup) | **Post** /auth/signup | 
[**StartTrial**](DefaultApi.md#StartTrial) | **Post** /auth/startTrial | 
[**StopJob**](DefaultApi.md#StopJob) | **Post** /jobs/stopJob | 
[**TerminateJob**](DefaultApi.md#TerminateJob) | **Post** /jobs/terminateJob | 
[**TrainFromInitialWeights**](DefaultApi.md#TrainFromInitialWeights) | **Post** /jobs/trainFromInitialWeights | 
[**TrainFromScratch**](DefaultApi.md#TrainFromScratch) | **Post** /jobs/trainFromScratch | 
[**TrashDataset**](DefaultApi.md#TrashDataset) | **Post** /datasets/trashDataset | 
[**TrashSecretManager**](DefaultApi.md#TrashSecretManager) | **Post** /secret-manager/trashSecretManager | 
[**UpdateDashboard**](DefaultApi.md#UpdateDashboard) | **Post** /dashboards/updateDashboard | 
[**UpdateIssue**](DefaultApi.md#UpdateIssue) | **Post** /issues/updateIssue | 
[**UpdateOrganizationPublicName**](DefaultApi.md#UpdateOrganizationPublicName) | **Post** /organizations/updateOrganizationPublicName | 
[**UpdateSecretManager**](DefaultApi.md#UpdateSecretManager) | **Post** /secret-manager/updateSecretManager | 
[**UpdateSessionTest**](DefaultApi.md#UpdateSessionTest) | **Post** /sessions-tests/updateSessionTest | 
[**UpdateUserName**](DefaultApi.md#UpdateUserName) | **Post** /users/updateUserName | 
[**UpdateUserOrganization**](DefaultApi.md#UpdateUserOrganization) | **Post** /users/updateUserOrganization | 
[**UpdateUserRole**](DefaultApi.md#UpdateUserRole) | **Post** /users/updateUserRole | 
[**UpdateUserStatus**](DefaultApi.md#UpdateUserStatus) | **Post** /users/updateUserStatus | 
[**UpdateVersion**](DefaultApi.md#UpdateVersion) | **Post** /versions/updateVersion | 
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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    addDashboardParams := *openapiclient.NewAddDashboardParams("ProjectId_example", "Name_example", []openapiclient.DashboardItem{*openapiclient.NewDashboardItem("Id_example", *openapiclient.NewSizedLayout(*openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123))), "Type_example", *openapiclient.NewCustomVisualizationData(map[string]interface{}(123), "Name_example", openapiclient.CustomVisualizationType("Bar")))}) // AddDashboardParams | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    addExportModelJobParams := *openapiclient.NewAddExportModelJobParams("TargetSessionRunId_example", openapiclient.ExportModelTypeEnum("JSON_TF2"), "Title_example", false, float64(123)) // AddExportModelJobParams | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddProject

> AddProjectResponse AddProject(ctx).AddProjectParams(addProjectParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    addProjectParams := *openapiclient.NewAddProjectParams("Name_example", "Description_example") // AddProjectParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddProject(context.Background()).AddProjectParams(addProjectParams).Execute()
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
 **addProjectParams** | [**AddProjectParams**](AddProjectParams.md) |  | 

### Return type

[**AddProjectResponse**](AddProjectResponse.md)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    addVersionParams := *openapiclient.NewAddVersionParams("ProjectId_example", *openapiclient.NewModelGraph("Id_example", map[string]interface{}(123)), "BranchName_example", "Description_example", *openapiclient.NewDatasetSetup(*openapiclient.NewDatasetPreprocess(float64(123), float64(123)), []openapiclient.DatasetInputInstance{*openapiclient.NewDatasetInputInstance("Name_example", []float64{float64(123)})}, []openapiclient.DatasetMetadataInstance{*openapiclient.NewDatasetMetadataInstance("Name_example", openapiclient.DatasetMetadataType("float"))}, []openapiclient.DatasetOutputInstance{*openapiclient.NewDatasetOutputInstance("Name_example", []float64{float64(123)})}, []openapiclient.VisualizerInstance{*openapiclient.NewVisualizerInstance("Name_example", openapiclient.LeapDataType("Image"), []string{"ArgNames_example"})}, []openapiclient.PredictionTypeInstance{*openapiclient.NewPredictionTypeInstance("Name_example", []string{"Labels_example"})}, []openapiclient.CustomLossInstance{*openapiclient.NewCustomLossInstance("Name_example", []string{"ArgNames_example"})}, []openapiclient.MetricInstance{*openapiclient.NewMetricInstance("Name_example", []string{"ArgNames_example"})})) // AddVersionParams | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContinueTrain

> Job ContinueTrain(ctx).ContinueTrainParams(continueTrainParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    continueTrainParams := *openapiclient.NewContinueTrainParams("VersionId_example", "SessionRunId_example", float64(123), *openapiclient.NewTrainingParams(float64(123), float64(123)), false) // ContinueTrainParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ContinueTrain(context.Background()).ContinueTrainParams(continueTrainParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContinueTrain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContinueTrain`: Job
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContinueTrain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContinueTrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **continueTrainParams** | [**ContinueTrainParams**](ContinueTrainParams.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganization

> CreateOrganizationResponse CreateOrganization(ctx).CreateOrganizationRequest(createOrganizationRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    createOrganizationRequest := *openapiclient.NewCreateOrganizationRequest("Name_example", "PublicName_example") // CreateOrganizationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateOrganization(context.Background()).CreateOrganizationRequest(createOrganizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganization`: CreateOrganizationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrganizationRequest** | [**CreateOrganizationRequest**](CreateOrganizationRequest.md) |  | 

### Return type

[**CreateOrganizationResponse**](CreateOrganizationResponse.md)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    deleteDashboardParams := *openapiclient.NewDeleteDashboardParams("DashboardId_example") // DeleteDashboardParams | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    deleteIssueParams := *openapiclient.NewDeleteIssueParams("IssueId_example") // DeleteIssueParams | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganization

> DeleteOrganization(ctx).DeleteOrganizationRequest(deleteOrganizationRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    deleteOrganizationRequest := *openapiclient.NewDeleteOrganizationRequest("Id_example") // DeleteOrganizationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteOrganization(context.Background()).DeleteOrganizationRequest(deleteOrganizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteOrganizationRequest** | [**DeleteOrganizationRequest**](DeleteOrganizationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    deleteSessionParams := *openapiclient.NewDeleteSessionParams("SessionId_example") // DeleteSessionParams | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    deleteSessionRunParams := *openapiclient.NewDeleteSessionRunParams("SessionRunId_example") // DeleteSessionRunParams | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    deleteSessionTestRequest := *openapiclient.NewDeleteSessionTestRequest("Id_example") // DeleteSessionTestRequest | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    deleteVersionParams := *openapiclient.NewDeleteVersionParams("VersionId_example") // DeleteVersionParams | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    deleteVisualizationsParams := *openapiclient.NewDeleteVisualizationsParams([]string{"VisualizationIdsToDelete_example"}) // DeleteVisualizationsParams | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    evaluateParams := *openapiclient.NewEvaluateParams("VersionId_example", "SessionRunId_example", float64(123), []openapiclient.DataStateForEval{openapiclient.DataStateForEval("training")}, false, float64(123)) // EvaluateParams | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSlimUserData

> GetOrganizationUsersResponse GetAllSlimUserData(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAllSlimUserData(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllSlimUserData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSlimUserData`: GetOrganizationUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllSlimUserData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSlimUserDataRequest struct via the builder pattern


### Return type

[**GetOrganizationUsersResponse**](GetOrganizationUsersResponse.md)

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfusionMetricNames

> []GetConfusionMetricNamesObject GetConfusionMetricNames(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetConfusionMetricNames(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetConfusionMetricNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfusionMetricNames`: []GetConfusionMetricNamesObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetConfusionMetricNames`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfusionMetricNamesRequest struct via the builder pattern


### Return type

[**[]GetConfusionMetricNamesObject**](GetConfusionMetricNamesObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    getDashboardParams := *openapiclient.NewGetDashboardParams("DashboardId_example") // GetDashboardParams | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashletFields

> GetDashletFieldsResponse GetDashletFields(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetDashletFields(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDashletFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashletFields`: GetDashletFieldsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDashletFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashletFieldsRequest struct via the builder pattern


### Return type

[**GetDashletFieldsResponse**](GetDashletFieldsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasetVersionUploadUrl

> DatasetVersionUploadUrlResponse GetDatasetVersionUploadUrl(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetDatasetVersionUploadUrl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDatasetVersionUploadUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatasetVersionUploadUrl`: DatasetVersionUploadUrlResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDatasetVersionUploadUrl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetVersionUploadUrlRequest struct via the builder pattern


### Return type

[**DatasetVersionUploadUrlResponse**](DatasetVersionUploadUrlResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    getExportedSessionRunJobsParams := *openapiclient.NewGetExportedSessionRunJobsParams("SessionRunId_example") // GetExportedSessionRunJobsParams | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    heatmapChartsParams := *openapiclient.NewHeatmapChartsParams("XField_example", float64(123), openapiclient.DataDistributionType("distinct"), "YField_example", float64(123), openapiclient.DataDistributionType("distinct"), "ColorField_example", openapiclient.AggregationMethod("Average"), []string{"SessionRunIds_example"}) // HeatmapChartsParams | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobs

> GetJobsResponse GetJobs(ctx).GetJobsParams(getJobsParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    getJobsParams := *openapiclient.NewGetJobsParams() // GetJobsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetJobs(context.Background()).GetJobsParams(getJobsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobs`: GetJobsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getJobsParams** | [**GetJobsParams**](GetJobsParams.md) |  | 

### Return type

[**GetJobsResponse**](GetJobsResponse.md)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationJobs

> GetJobsResponse GetOrganizationJobs(ctx).GetOrganizationJobsParams(getOrganizationJobsParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    getOrganizationJobsParams := *openapiclient.NewGetOrganizationJobsParams() // GetOrganizationJobsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetOrganizationJobs(context.Background()).GetOrganizationJobsParams(getOrganizationJobsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOrganizationJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationJobs`: GetJobsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetOrganizationJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getOrganizationJobsParams** | [**GetOrganizationJobsParams**](GetOrganizationJobsParams.md) |  | 

### Return type

[**GetJobsResponse**](GetJobsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSlimUserData

> GetOrganizationUsersResponse GetOrganizationSlimUserData(ctx).GetOrganizationUsersRequest(getOrganizationUsersRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    getOrganizationUsersRequest := *openapiclient.NewGetOrganizationUsersRequest("OrganizationId_example") // GetOrganizationUsersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetOrganizationSlimUserData(context.Background()).GetOrganizationUsersRequest(getOrganizationUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOrganizationSlimUserData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSlimUserData`: GetOrganizationUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetOrganizationSlimUserData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSlimUserDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getOrganizationUsersRequest** | [**GetOrganizationUsersRequest**](GetOrganizationUsersRequest.md) |  | 

### Return type

[**GetOrganizationUsersResponse**](GetOrganizationUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizations

> GetOrganizationsResponse GetOrganizations(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetOrganizations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizations`: GetOrganizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetOrganizations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsRequest struct via the builder pattern


### Return type

[**GetOrganizationsResponse**](GetOrganizationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecentOrganizationSessions

> RecentSessionsResponse GetRecentOrganizationSessions(ctx).RecentOrganizationSessionsRequestParams(recentOrganizationSessionsRequestParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    recentOrganizationSessionsRequestParams := *openapiclient.NewRecentOrganizationSessionsRequestParams(float64(123)) // RecentOrganizationSessionsRequestParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRecentOrganizationSessions(context.Background()).RecentOrganizationSessionsRequestParams(recentOrganizationSessionsRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRecentOrganizationSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecentOrganizationSessions`: RecentSessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRecentOrganizationSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecentOrganizationSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recentOrganizationSessionsRequestParams** | [**RecentOrganizationSessionsRequestParams**](RecentOrganizationSessionsRequestParams.md) |  | 

### Return type

[**RecentSessionsResponse**](RecentSessionsResponse.md)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    body := MultiThresholdConfusionMatrixParams(987) // MultiThresholdConfusionMatrixParams | 

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
 **body** | **MultiThresholdConfusionMatrixParams** |  | 

### Return type

[**MultiChartsResponse**](MultiChartsResponse.md)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionInsights

> GetSessionInsightsResponse GetSessionInsights(ctx).GetSessionInsightsParams(getSessionInsightsParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    getSessionInsightsParams := *openapiclient.NewGetSessionInsightsParams("SessionRunId_example") // GetSessionInsightsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSessionInsights(context.Background()).GetSessionInsightsParams(getSessionInsightsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSessionInsights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionInsights`: GetSessionInsightsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSessionInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSessionInsightsParams** | [**GetSessionInsightsParams**](GetSessionInsightsParams.md) |  | 

### Return type

[**GetSessionInsightsResponse**](GetSessionInsightsResponse.md)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    getSessionTestResultsRequest := *openapiclient.NewGetSessionTestResultsRequest("ProjectId_example", []openapiclient.SessionTestData{*openapiclient.NewSessionTestData("SessionRunId_example", float64(123))}) // GetSessionTestResultsRequest | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionVisualizations

> GetSessionVisualizationsResponse GetSessionVisualizations(ctx).GetSessionVisualizationsParams(getSessionVisualizationsParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    getSessionVisualizationsParams := *openapiclient.NewGetSessionVisualizationsParams("SessionRunId_example") // GetSessionVisualizationsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSessionVisualizations(context.Background()).GetSessionVisualizationsParams(getSessionVisualizationsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSessionVisualizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionVisualizations`: GetSessionVisualizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSessionVisualizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionVisualizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSessionVisualizationsParams** | [**GetSessionVisualizationsParams**](GetSessionVisualizationsParams.md) |  | 

### Return type

[**GetSessionVisualizationsResponse**](GetSessionVisualizationsResponse.md)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    sessionHashRequestParams := *openapiclient.NewSessionHashRequestParams("Hash_example") // SessionHashRequestParams | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    sessionVersionIdRequestParams := *openapiclient.NewSessionVersionIdRequestParams("VersionId_example") // SessionVersionIdRequestParams | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignedStoredResourceUrl

> SignedStoredResourseUrl GetSignedStoredResourceUrl(ctx).GetStoredResourceUrlParams(getStoredResourceUrlParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    getStoredResourceUrlParams := *openapiclient.NewGetStoredResourceUrlParams("FileName_example") // GetStoredResourceUrlParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSignedStoredResourceUrl(context.Background()).GetStoredResourceUrlParams(getStoredResourceUrlParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSignedStoredResourceUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignedStoredResourceUrl`: SignedStoredResourseUrl
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSignedStoredResourceUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignedStoredResourceUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getStoredResourceUrlParams** | [**GetStoredResourceUrlParams**](GetStoredResourceUrlParams.md) |  | 

### Return type

[**SignedStoredResourseUrl**](SignedStoredResourseUrl.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignedStoredUrlByVisualizationId

> SignedStoredVisualization GetSignedStoredUrlByVisualizationId(ctx).GetVisualizationParams(getVisualizationParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    getVisualizationParams := *openapiclient.NewGetVisualizationParams("VisualizationId_example") // GetVisualizationParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSignedStoredUrlByVisualizationId(context.Background()).GetVisualizationParams(getVisualizationParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSignedStoredUrlByVisualizationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignedStoredUrlByVisualizationId`: SignedStoredVisualization
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSignedStoredUrlByVisualizationId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignedStoredUrlByVisualizationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getVisualizationParams** | [**GetVisualizationParams**](GetVisualizationParams.md) |  | 

### Return type

[**SignedStoredVisualization**](SignedStoredVisualization.md)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    getSingleIssueParams := *openapiclient.NewGetSingleIssueParams("Id_example") // GetSingleIssueParams | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    getSingleSessionTestRequest := *openapiclient.NewGetSingleSessionTestRequest("Id_example") // GetSingleSessionTestRequest | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatistics

> GetStatisticsResponse GetStatistics(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetStatistics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatistics`: GetStatisticsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStatistics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatisticsRequest struct via the builder pattern


### Return type

[**GetStatisticsResponse**](GetStatisticsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStoredExportedSessionRunResourceUrl

> GetStoredExportedSessionRunResourceUrlResponse GetStoredExportedSessionRunResourceUrl(ctx).GetStoredExportedSessionRunResourceUrlParams(getStoredExportedSessionRunResourceUrlParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    getStoredExportedSessionRunResourceUrlParams := *openapiclient.NewGetStoredExportedSessionRunResourceUrlParams("ExportedSessionRunId_example") // GetStoredExportedSessionRunResourceUrlParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetStoredExportedSessionRunResourceUrl(context.Background()).GetStoredExportedSessionRunResourceUrlParams(getStoredExportedSessionRunResourceUrlParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStoredExportedSessionRunResourceUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStoredExportedSessionRunResourceUrl`: GetStoredExportedSessionRunResourceUrlResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStoredExportedSessionRunResourceUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStoredExportedSessionRunResourceUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getStoredExportedSessionRunResourceUrlParams** | [**GetStoredExportedSessionRunResourceUrlParams**](GetStoredExportedSessionRunResourceUrlParams.md) |  | 

### Return type

[**GetStoredExportedSessionRunResourceUrlResponse**](GetStoredExportedSessionRunResourceUrlResponse.md)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    genericDataQueryParams := *openapiclient.NewGenericDataQueryParams([]string{"SessionRunIds_example"}, []openapiclient.Aggregations{*openapiclient.NewAggregations("Field_example", openapiclient.AggregationMethod("Average"))}, []openapiclient.BucketAggregation{*openapiclient.NewBucketAggregation("XField_example", openapiclient.DataDistributionType("distinct"), float64(123), "OrderParams_example")}) // GenericDataQueryParams | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrainingJobs

> []Job GetTrainingJobs(ctx).GetTrainingJobParams(getTrainingJobParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    getTrainingJobParams := *openapiclient.NewGetTrainingJobParams("VersionId_example") // GetTrainingJobParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetTrainingJobs(context.Background()).GetTrainingJobParams(getTrainingJobParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTrainingJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrainingJobs`: []Job
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTrainingJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrainingJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getTrainingJobParams** | [**GetTrainingJobParams**](GetTrainingJobParams.md) |  | 

### Return type

[**[]Job**](Job.md)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVisualization

> Visualization GetVisualization(ctx).GetVisualizationParams(getVisualizationParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    getVisualizationParams := *openapiclient.NewGetVisualizationParams("VisualizationId_example") // GetVisualizationParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVisualization(context.Background()).GetVisualizationParams(getVisualizationParams).Execute()
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
 **getVisualizationParams** | [**GetVisualizationParams**](GetVisualizationParams.md) |  | 

### Return type

[**Visualization**](Visualization.md)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    multiChartsParams := *openapiclient.NewMultiChartsParams("XField_example", "YField_example", openapiclient.AggregationMethod("Average"), []string{"SessionRunIds_example"}, openapiclient.DataDistributionType("distinct"), float64(123)) // MultiChartsParams | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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


## ImportModel

> ExternalImportModelStorageResponse ImportModel(ctx).ImportNewModelParams(importNewModelParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsTrainingJobRunning

> IsTrainingJobRunningResponse IsTrainingJobRunning(ctx).IsTrainingJobRunningParams(isTrainingJobRunningParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    isTrainingJobRunningParams := *openapiclient.NewIsTrainingJobRunningParams("VersionId_example", "SessionRunId_example") // IsTrainingJobRunningParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.IsTrainingJobRunning(context.Background()).IsTrainingJobRunningParams(isTrainingJobRunningParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.IsTrainingJobRunning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IsTrainingJobRunning`: IsTrainingJobRunningResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.IsTrainingJobRunning`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIsTrainingJobRunningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isTrainingJobRunningParams** | [**IsTrainingJobRunningParams**](IsTrainingJobRunningParams.md) |  | 

### Return type

[**IsTrainingJobRunningResponse**](IsTrainingJobRunningResponse.md)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyLogin

> UserData KeyLogin(ctx).KeyLoginParams(keyLoginParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    keyLoginParams := *openapiclient.NewKeyLoginParams("ApiId_example", "ApiKey_example") // KeyLoginParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.KeyLogin(context.Background()).KeyLoginParams(keyLoginParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.KeyLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyLogin`: UserData
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.KeyLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyLoginParams** | [**KeyLoginParams**](KeyLoginParams.md) |  | 

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


## LoadModel

> LoadSessionResponse LoadModel(ctx).LoadSessionParams(loadSessionParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    loadSessionParams := *openapiclient.NewLoadSessionParams("SessionId_example") // LoadSessionParams | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    loadVersionParams := *openapiclient.NewLoadVersionParams("VersionId_example") // LoadVersionParams | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    loginParams := *openapiclient.NewLoginParams("Email_example", "Password_example") // LoginParams | 

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PopulationExploration

> Job PopulationExploration(ctx).PopulationExplorationParams(populationExplorationParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    populationExplorationParams := *openapiclient.NewPopulationExplorationParams("SessionRunId_example", float64(123), float64(123), float64(123)) // PopulationExplorationParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PopulationExploration(context.Background()).PopulationExplorationParams(populationExplorationParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PopulationExploration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PopulationExploration`: Job
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

[**Job**](Job.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    sampleAnalysisParams := *openapiclient.NewSampleAnalysisParams("SessionRunId_example", *openapiclient.NewSampleIdentity(openapiclient.DataStateType("training"), float64(123)), float64(123)) // SampleAnalysisParams | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SampleSelection

> Job SampleSelection(ctx).SampleSelectionParams(sampleSelectionParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    sampleSelectionParams := *openapiclient.NewSampleSelectionParams("SessionRunId_example", *openapiclient.NewSampleIdentity(openapiclient.DataStateType("training"), float64(123)), openapiclient.DataStateType("training"), float64(123), float64(123), float64(123)) // SampleSelectionParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SampleSelection(context.Background()).SampleSelectionParams(sampleSelectionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SampleSelection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SampleSelection`: Job
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SampleSelection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSampleSelectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sampleSelectionParams** | [**SampleSelectionParams**](SampleSelectionParams.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    saveAnalyzerLayoutParams := *openapiclient.NewSaveAnalyzerLayoutParams([]openapiclient.PanelLayout{*openapiclient.NewPanelLayout("VisualizationId_example", *openapiclient.NewSizedLayout(*openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123))))}) // SaveAnalyzerLayoutParams | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultOrganization

> SetDefaultOrganization(ctx).SetDefaultOrganizationRequest(setDefaultOrganizationRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    setDefaultOrganizationRequest := *openapiclient.NewSetDefaultOrganizationRequest("Id_example") // SetDefaultOrganizationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.SetDefaultOrganization(context.Background()).SetDefaultOrganizationRequest(setDefaultOrganizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetDefaultOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setDefaultOrganizationRequest** | [**SetDefaultOrganizationRequest**](SetDefaultOrganizationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOrganizationMachineType

> SetOrganizationMachineType(ctx).SetOrganizationMachineTypeParams(setOrganizationMachineTypeParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    setOrganizationMachineTypeParams := *openapiclient.NewSetOrganizationMachineTypeParams("OrganizationId_example", "MachineTypeId_example") // SetOrganizationMachineTypeParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.SetOrganizationMachineType(context.Background()).SetOrganizationMachineTypeParams(setOrganizationMachineTypeParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetOrganizationMachineType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetOrganizationMachineTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setOrganizationMachineTypeParams** | [**SetOrganizationMachineTypeParams**](SetOrganizationMachineTypeParams.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Signup

> SignupResponse Signup(ctx).SignupParams(signupParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    signupParams := *openapiclient.NewSignupParams("Name_example", "Email_example", "Password_example") // SignupParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Signup(context.Background()).SignupParams(signupParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Signup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Signup`: SignupResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Signup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signupParams** | [**SignupParams**](SignupParams.md) |  | 

### Return type

[**SignupResponse**](SignupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrainFromInitialWeights

> Job TrainFromInitialWeights(ctx).TrainFromInitialWeightsParams(trainFromInitialWeightsParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    trainFromInitialWeightsParams := *openapiclient.NewTrainFromInitialWeightsParams("VersionId_example", "FromSessionRunId_example", float64(123), "ModelName_example", *openapiclient.NewTrainingParams(float64(123), float64(123)), false) // TrainFromInitialWeightsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TrainFromInitialWeights(context.Background()).TrainFromInitialWeightsParams(trainFromInitialWeightsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TrainFromInitialWeights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TrainFromInitialWeights`: Job
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TrainFromInitialWeights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrainFromInitialWeightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trainFromInitialWeightsParams** | [**TrainFromInitialWeightsParams**](TrainFromInitialWeightsParams.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrainFromScratch

> Job TrainFromScratch(ctx).TrainFromScratchParams(trainFromScratchParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    trainFromScratchParams := *openapiclient.NewTrainFromScratchParams("VersionId_example", "ModelName_example", *openapiclient.NewTrainingParams(float64(123), float64(123)), false) // TrainFromScratchParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TrainFromScratch(context.Background()).TrainFromScratchParams(trainFromScratchParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TrainFromScratch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TrainFromScratch`: Job
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TrainFromScratch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrainFromScratchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trainFromScratchParams** | [**TrainFromScratchParams**](TrainFromScratchParams.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    updateDashboardParams := *openapiclient.NewUpdateDashboardParams("DashboardId_example", "Name_example", []openapiclient.DashboardItem{*openapiclient.NewDashboardItem("Id_example", *openapiclient.NewSizedLayout(*openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewLayout(float64(123), float64(123), float64(123), float64(123))), "Type_example", *openapiclient.NewCustomVisualizationData(map[string]interface{}(123), "Name_example", openapiclient.CustomVisualizationType("Bar")))}) // UpdateDashboardParams | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    updateIssueParams := *openapiclient.NewUpdateIssueParams("Id_example") // UpdateIssueParams | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationPublicName

> UpdateOrganizationPublicName(ctx).UpdateOrganizationPublicNameRequest(updateOrganizationPublicNameRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    updateOrganizationPublicNameRequest := *openapiclient.NewUpdateOrganizationPublicNameRequest("Id_example", "PublicName_example") // UpdateOrganizationPublicNameRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.UpdateOrganizationPublicName(context.Background()).UpdateOrganizationPublicNameRequest(updateOrganizationPublicNameRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateOrganizationPublicName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationPublicNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateOrganizationPublicNameRequest** | [**UpdateOrganizationPublicNameRequest**](UpdateOrganizationPublicNameRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    updateSecretManagerParams := *openapiclient.NewUpdateSecretManagerParams("Id_example", "Name_example") // UpdateSecretManagerParams | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    updateSessionTestRequest := *openapiclient.NewUpdateSessionTestRequest("Id_example") // UpdateSessionTestRequest | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserOrganization

> SlimUserData UpdateUserOrganization(ctx).UpdateUserOrganizationRequest(updateUserOrganizationRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    updateUserOrganizationRequest := *openapiclient.NewUpdateUserOrganizationRequest("UserId_example", "OrganizationId_example") // UpdateUserOrganizationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateUserOrganization(context.Background()).UpdateUserOrganizationRequest(updateUserOrganizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUserOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserOrganization`: SlimUserData
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateUserOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserOrganizationRequest** | [**UpdateUserOrganizationRequest**](UpdateUserOrganizationRequest.md) |  | 

### Return type

[**SlimUserData**](SlimUserData.md)

### Authorization

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func main() {
    updateVersionParams := *openapiclient.NewUpdateVersionParams("VersionId_example", *openapiclient.NewModelGraph("Id_example", map[string]interface{}(123)), *openapiclient.NewDatasetSetup(*openapiclient.NewDatasetPreprocess(float64(123), float64(123)), []openapiclient.DatasetInputInstance{*openapiclient.NewDatasetInputInstance("Name_example", []float64{float64(123)})}, []openapiclient.DatasetMetadataInstance{*openapiclient.NewDatasetMetadataInstance("Name_example", openapiclient.DatasetMetadataType("float"))}, []openapiclient.DatasetOutputInstance{*openapiclient.NewDatasetOutputInstance("Name_example", []float64{float64(123)})}, []openapiclient.VisualizerInstance{*openapiclient.NewVisualizerInstance("Name_example", openapiclient.LeapDataType("Image"), []string{"ArgNames_example"})}, []openapiclient.PredictionTypeInstance{*openapiclient.NewPredictionTypeInstance("Name_example", []string{"Labels_example"})}, []openapiclient.CustomLossInstance{*openapiclient.NewCustomLossInstance("Name_example", []string{"ArgNames_example"})}, []openapiclient.MetricInstance{*openapiclient.NewMetricInstance("Name_example", []string{"ArgNames_example"})}), false) // UpdateVersionParams | 

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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

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
    openapiclient "github.com/tensorleap/cli-go/pkg/tensorleapapi"
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

