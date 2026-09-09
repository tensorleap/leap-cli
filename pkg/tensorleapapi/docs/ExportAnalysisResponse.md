# ExportAnalysisResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractVersion** | **float64** |  | 
**ProjectId** | **string** |  | 
**VersionId** | **string** |  | 
**Version** | [**ExportAnalysisResponseVersion**](ExportAnalysisResponseVersion.md) |  | 
**Insights** | [**[]AnalysisInsight**](AnalysisInsight.md) |  | 
**PopulationCsvUrl** | Pointer to **string** |  | [optional] 
**IntegrationCodeUrl** | Pointer to **string** |  | [optional] 
**IntegrationEntryFile** | Pointer to **string** |  | [optional] 
**PredictionLabels** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 
**Visualizers** | [**[]AnalysisVisualizerInfo**](AnalysisVisualizerInfo.md) |  | 
**DeepLinkPath** | **string** |  | 

## Methods

### NewExportAnalysisResponse

`func NewExportAnalysisResponse(contractVersion float64, projectId string, versionId string, version ExportAnalysisResponseVersion, insights []AnalysisInsight, predictionLabels map[string]interface{}, visualizers []AnalysisVisualizerInfo, deepLinkPath string, ) *ExportAnalysisResponse`

NewExportAnalysisResponse instantiates a new ExportAnalysisResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportAnalysisResponseWithDefaults

`func NewExportAnalysisResponseWithDefaults() *ExportAnalysisResponse`

NewExportAnalysisResponseWithDefaults instantiates a new ExportAnalysisResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractVersion

`func (o *ExportAnalysisResponse) GetContractVersion() float64`

GetContractVersion returns the ContractVersion field if non-nil, zero value otherwise.

### GetContractVersionOk

`func (o *ExportAnalysisResponse) GetContractVersionOk() (*float64, bool)`

GetContractVersionOk returns a tuple with the ContractVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractVersion

`func (o *ExportAnalysisResponse) SetContractVersion(v float64)`

SetContractVersion sets ContractVersion field to given value.


### GetProjectId

`func (o *ExportAnalysisResponse) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ExportAnalysisResponse) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ExportAnalysisResponse) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersionId

`func (o *ExportAnalysisResponse) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ExportAnalysisResponse) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ExportAnalysisResponse) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetVersion

`func (o *ExportAnalysisResponse) GetVersion() ExportAnalysisResponseVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExportAnalysisResponse) GetVersionOk() (*ExportAnalysisResponseVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExportAnalysisResponse) SetVersion(v ExportAnalysisResponseVersion)`

SetVersion sets Version field to given value.


### GetInsights

`func (o *ExportAnalysisResponse) GetInsights() []AnalysisInsight`

GetInsights returns the Insights field if non-nil, zero value otherwise.

### GetInsightsOk

`func (o *ExportAnalysisResponse) GetInsightsOk() (*[]AnalysisInsight, bool)`

GetInsightsOk returns a tuple with the Insights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsights

`func (o *ExportAnalysisResponse) SetInsights(v []AnalysisInsight)`

SetInsights sets Insights field to given value.


### GetPopulationCsvUrl

`func (o *ExportAnalysisResponse) GetPopulationCsvUrl() string`

GetPopulationCsvUrl returns the PopulationCsvUrl field if non-nil, zero value otherwise.

### GetPopulationCsvUrlOk

`func (o *ExportAnalysisResponse) GetPopulationCsvUrlOk() (*string, bool)`

GetPopulationCsvUrlOk returns a tuple with the PopulationCsvUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationCsvUrl

`func (o *ExportAnalysisResponse) SetPopulationCsvUrl(v string)`

SetPopulationCsvUrl sets PopulationCsvUrl field to given value.

### HasPopulationCsvUrl

`func (o *ExportAnalysisResponse) HasPopulationCsvUrl() bool`

HasPopulationCsvUrl returns a boolean if a field has been set.

### GetIntegrationCodeUrl

`func (o *ExportAnalysisResponse) GetIntegrationCodeUrl() string`

GetIntegrationCodeUrl returns the IntegrationCodeUrl field if non-nil, zero value otherwise.

### GetIntegrationCodeUrlOk

`func (o *ExportAnalysisResponse) GetIntegrationCodeUrlOk() (*string, bool)`

GetIntegrationCodeUrlOk returns a tuple with the IntegrationCodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationCodeUrl

`func (o *ExportAnalysisResponse) SetIntegrationCodeUrl(v string)`

SetIntegrationCodeUrl sets IntegrationCodeUrl field to given value.

### HasIntegrationCodeUrl

`func (o *ExportAnalysisResponse) HasIntegrationCodeUrl() bool`

HasIntegrationCodeUrl returns a boolean if a field has been set.

### GetIntegrationEntryFile

`func (o *ExportAnalysisResponse) GetIntegrationEntryFile() string`

GetIntegrationEntryFile returns the IntegrationEntryFile field if non-nil, zero value otherwise.

### GetIntegrationEntryFileOk

`func (o *ExportAnalysisResponse) GetIntegrationEntryFileOk() (*string, bool)`

GetIntegrationEntryFileOk returns a tuple with the IntegrationEntryFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationEntryFile

`func (o *ExportAnalysisResponse) SetIntegrationEntryFile(v string)`

SetIntegrationEntryFile sets IntegrationEntryFile field to given value.

### HasIntegrationEntryFile

`func (o *ExportAnalysisResponse) HasIntegrationEntryFile() bool`

HasIntegrationEntryFile returns a boolean if a field has been set.

### GetPredictionLabels

`func (o *ExportAnalysisResponse) GetPredictionLabels() map[string]interface{}`

GetPredictionLabels returns the PredictionLabels field if non-nil, zero value otherwise.

### GetPredictionLabelsOk

`func (o *ExportAnalysisResponse) GetPredictionLabelsOk() (*map[string]interface{}, bool)`

GetPredictionLabelsOk returns a tuple with the PredictionLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionLabels

`func (o *ExportAnalysisResponse) SetPredictionLabels(v map[string]interface{})`

SetPredictionLabels sets PredictionLabels field to given value.


### GetVisualizers

`func (o *ExportAnalysisResponse) GetVisualizers() []AnalysisVisualizerInfo`

GetVisualizers returns the Visualizers field if non-nil, zero value otherwise.

### GetVisualizersOk

`func (o *ExportAnalysisResponse) GetVisualizersOk() (*[]AnalysisVisualizerInfo, bool)`

GetVisualizersOk returns a tuple with the Visualizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizers

`func (o *ExportAnalysisResponse) SetVisualizers(v []AnalysisVisualizerInfo)`

SetVisualizers sets Visualizers field to given value.


### GetDeepLinkPath

`func (o *ExportAnalysisResponse) GetDeepLinkPath() string`

GetDeepLinkPath returns the DeepLinkPath field if non-nil, zero value otherwise.

### GetDeepLinkPathOk

`func (o *ExportAnalysisResponse) GetDeepLinkPathOk() (*string, bool)`

GetDeepLinkPathOk returns a tuple with the DeepLinkPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepLinkPath

`func (o *ExportAnalysisResponse) SetDeepLinkPath(v string)`

SetDeepLinkPath sets DeepLinkPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


