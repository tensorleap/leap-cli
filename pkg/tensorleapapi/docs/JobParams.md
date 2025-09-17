# JobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Epochs** | **float64** |  | 
**BatchSize** | **float64** |  | 
**EarlyStopParams** | Pointer to [**EarlyStopParams**](EarlyStopParams.md) |  | [optional] 
**VersionId** | **string** |  | 
**ProjectId** | **string** |  | 
**BatchSize** | **float64** |  | 
**DataStates** | [**[]DataStateType**](DataStateType.md) |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**SkipMetricsEstimation** | **bool** |  | 
**Monitor** | Pointer to **bool** |  | [optional] 
**EvaluatedEpoch** | **float64** |  | 
**SessionId** | **string** |  | 
**Type** | **string** |  | 
**FromEpoch** | **float64** |  | 
**SampleIdentity** | Pointer to [**SampleIdentity**](SampleIdentity.md) |  | [optional] 
**FromDatasetSlice** | Pointer to [**DataStateType**](DataStateType.md) |  | [optional] 
**ExtId** | **string** |  | 
**Title** | **string** |  | 
**Epoch** | **float64** |  | 
**Digest** | **string** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**SampleIds** | [**[]SampleIdentity**](SampleIdentity.md) |  | 
**Limit** | **float64** |  | 
**SessionRunId** | **string** |  | 
**NumOfSamplesToLabel** | Pointer to **float64** |  | [optional] 
**RecommendLabeling** | Pointer to **bool** |  | [optional] 
**TriggerCreateSampleVisualiztions** | Pointer to **bool** |  | [optional] 
**DisplayParams** | [**PopulationExplorationDisplayParams**](PopulationExplorationDisplayParams.md) |  | 
**DomainGapMetadata** | Pointer to **string** |  | [optional] 
**ProjectionMetric** | Pointer to **string** |  | [optional] 
**ExportUrl** | **string** |  | 
**ProjectVersion** | **float64** |  | 
**ExportOptions** | [**ExportOptions**](ExportOptions.md) |  | 
**AlreadyExported** | **bool** |  | 
**ProjectExportMeta** | [**ExportProjectMeta**](ExportProjectMeta.md) |  | 
**CopyToUrl** | Pointer to **string** |  | [optional] 
**ImportUrl** | **string** |  | 
**ProjectMeta** | [**ProjectMeta**](ProjectMeta.md) |  | 
**DatasetId** | **string** |  | 
**Setup** | Pointer to [**DatasetSetup**](DatasetSetup.md) |  | [optional] 
**SecretManagerId** | Pointer to **string** |  | [optional] 
**CodeUrl** | **string** |  | 
**CodeEntryFile** | **string** |  | 
**Branch** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**GenericBaseImageType** | Pointer to **string** |  | [optional] 

## Methods

### NewJobParams

`func NewJobParams(epochs float64, batchSize float64, versionId string, projectId string, batchSize float64, dataStates []DataStateType, name string, description string, skipMetricsEstimation bool, evaluatedEpoch float64, sessionId string, type_ string, fromEpoch float64, extId string, title string, epoch float64, digest string, sampleIds []SampleIdentity, limit float64, sessionRunId string, displayParams PopulationExplorationDisplayParams, exportUrl string, projectVersion float64, exportOptions ExportOptions, alreadyExported bool, projectExportMeta ExportProjectMeta, importUrl string, projectMeta ProjectMeta, datasetId string, codeUrl string, codeEntryFile string, ) *JobParams`

NewJobParams instantiates a new JobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobParamsWithDefaults

`func NewJobParamsWithDefaults() *JobParams`

NewJobParamsWithDefaults instantiates a new JobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpochs

`func (o *JobParams) GetEpochs() float64`

GetEpochs returns the Epochs field if non-nil, zero value otherwise.

### GetEpochsOk

`func (o *JobParams) GetEpochsOk() (*float64, bool)`

GetEpochsOk returns a tuple with the Epochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochs

`func (o *JobParams) SetEpochs(v float64)`

SetEpochs sets Epochs field to given value.


### GetBatchSize

`func (o *JobParams) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *JobParams) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *JobParams) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.


### GetEarlyStopParams

`func (o *JobParams) GetEarlyStopParams() EarlyStopParams`

GetEarlyStopParams returns the EarlyStopParams field if non-nil, zero value otherwise.

### GetEarlyStopParamsOk

`func (o *JobParams) GetEarlyStopParamsOk() (*EarlyStopParams, bool)`

GetEarlyStopParamsOk returns a tuple with the EarlyStopParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyStopParams

`func (o *JobParams) SetEarlyStopParams(v EarlyStopParams)`

SetEarlyStopParams sets EarlyStopParams field to given value.

### HasEarlyStopParams

`func (o *JobParams) HasEarlyStopParams() bool`

HasEarlyStopParams returns a boolean if a field has been set.

### GetVersionId

`func (o *JobParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *JobParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *JobParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetProjectId

`func (o *JobParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *JobParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *JobParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetBatchSize

`func (o *JobParams) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *JobParams) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *JobParams) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.


### GetDataStates

`func (o *JobParams) GetDataStates() []DataStateType`

GetDataStates returns the DataStates field if non-nil, zero value otherwise.

### GetDataStatesOk

`func (o *JobParams) GetDataStatesOk() (*[]DataStateType, bool)`

GetDataStatesOk returns a tuple with the DataStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStates

`func (o *JobParams) SetDataStates(v []DataStateType)`

SetDataStates sets DataStates field to given value.


### GetName

`func (o *JobParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobParams) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *JobParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobParams) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSkipMetricsEstimation

`func (o *JobParams) GetSkipMetricsEstimation() bool`

GetSkipMetricsEstimation returns the SkipMetricsEstimation field if non-nil, zero value otherwise.

### GetSkipMetricsEstimationOk

`func (o *JobParams) GetSkipMetricsEstimationOk() (*bool, bool)`

GetSkipMetricsEstimationOk returns a tuple with the SkipMetricsEstimation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipMetricsEstimation

`func (o *JobParams) SetSkipMetricsEstimation(v bool)`

SetSkipMetricsEstimation sets SkipMetricsEstimation field to given value.


### GetMonitor

`func (o *JobParams) GetMonitor() bool`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *JobParams) GetMonitorOk() (*bool, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *JobParams) SetMonitor(v bool)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *JobParams) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetEvaluatedEpoch

`func (o *JobParams) GetEvaluatedEpoch() float64`

GetEvaluatedEpoch returns the EvaluatedEpoch field if non-nil, zero value otherwise.

### GetEvaluatedEpochOk

`func (o *JobParams) GetEvaluatedEpochOk() (*float64, bool)`

GetEvaluatedEpochOk returns a tuple with the EvaluatedEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedEpoch

`func (o *JobParams) SetEvaluatedEpoch(v float64)`

SetEvaluatedEpoch sets EvaluatedEpoch field to given value.


### GetSessionId

`func (o *JobParams) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *JobParams) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *JobParams) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetType

`func (o *JobParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobParams) SetType(v string)`

SetType sets Type field to given value.


### GetFromEpoch

`func (o *JobParams) GetFromEpoch() float64`

GetFromEpoch returns the FromEpoch field if non-nil, zero value otherwise.

### GetFromEpochOk

`func (o *JobParams) GetFromEpochOk() (*float64, bool)`

GetFromEpochOk returns a tuple with the FromEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEpoch

`func (o *JobParams) SetFromEpoch(v float64)`

SetFromEpoch sets FromEpoch field to given value.


### GetSampleIdentity

`func (o *JobParams) GetSampleIdentity() SampleIdentity`

GetSampleIdentity returns the SampleIdentity field if non-nil, zero value otherwise.

### GetSampleIdentityOk

`func (o *JobParams) GetSampleIdentityOk() (*SampleIdentity, bool)`

GetSampleIdentityOk returns a tuple with the SampleIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleIdentity

`func (o *JobParams) SetSampleIdentity(v SampleIdentity)`

SetSampleIdentity sets SampleIdentity field to given value.

### HasSampleIdentity

`func (o *JobParams) HasSampleIdentity() bool`

HasSampleIdentity returns a boolean if a field has been set.

### GetFromDatasetSlice

`func (o *JobParams) GetFromDatasetSlice() DataStateType`

GetFromDatasetSlice returns the FromDatasetSlice field if non-nil, zero value otherwise.

### GetFromDatasetSliceOk

`func (o *JobParams) GetFromDatasetSliceOk() (*DataStateType, bool)`

GetFromDatasetSliceOk returns a tuple with the FromDatasetSlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDatasetSlice

`func (o *JobParams) SetFromDatasetSlice(v DataStateType)`

SetFromDatasetSlice sets FromDatasetSlice field to given value.

### HasFromDatasetSlice

`func (o *JobParams) HasFromDatasetSlice() bool`

HasFromDatasetSlice returns a boolean if a field has been set.

### GetExtId

`func (o *JobParams) GetExtId() string`

GetExtId returns the ExtId field if non-nil, zero value otherwise.

### GetExtIdOk

`func (o *JobParams) GetExtIdOk() (*string, bool)`

GetExtIdOk returns a tuple with the ExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtId

`func (o *JobParams) SetExtId(v string)`

SetExtId sets ExtId field to given value.


### GetTitle

`func (o *JobParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *JobParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *JobParams) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetEpoch

`func (o *JobParams) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *JobParams) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *JobParams) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetDigest

`func (o *JobParams) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *JobParams) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *JobParams) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetFilters

`func (o *JobParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *JobParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *JobParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *JobParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSampleIds

`func (o *JobParams) GetSampleIds() []SampleIdentity`

GetSampleIds returns the SampleIds field if non-nil, zero value otherwise.

### GetSampleIdsOk

`func (o *JobParams) GetSampleIdsOk() (*[]SampleIdentity, bool)`

GetSampleIdsOk returns a tuple with the SampleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleIds

`func (o *JobParams) SetSampleIds(v []SampleIdentity)`

SetSampleIds sets SampleIds field to given value.


### GetLimit

`func (o *JobParams) GetLimit() float64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *JobParams) GetLimitOk() (*float64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *JobParams) SetLimit(v float64)`

SetLimit sets Limit field to given value.


### GetSessionRunId

`func (o *JobParams) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *JobParams) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *JobParams) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetNumOfSamplesToLabel

`func (o *JobParams) GetNumOfSamplesToLabel() float64`

GetNumOfSamplesToLabel returns the NumOfSamplesToLabel field if non-nil, zero value otherwise.

### GetNumOfSamplesToLabelOk

`func (o *JobParams) GetNumOfSamplesToLabelOk() (*float64, bool)`

GetNumOfSamplesToLabelOk returns a tuple with the NumOfSamplesToLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfSamplesToLabel

`func (o *JobParams) SetNumOfSamplesToLabel(v float64)`

SetNumOfSamplesToLabel sets NumOfSamplesToLabel field to given value.

### HasNumOfSamplesToLabel

`func (o *JobParams) HasNumOfSamplesToLabel() bool`

HasNumOfSamplesToLabel returns a boolean if a field has been set.

### GetRecommendLabeling

`func (o *JobParams) GetRecommendLabeling() bool`

GetRecommendLabeling returns the RecommendLabeling field if non-nil, zero value otherwise.

### GetRecommendLabelingOk

`func (o *JobParams) GetRecommendLabelingOk() (*bool, bool)`

GetRecommendLabelingOk returns a tuple with the RecommendLabeling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendLabeling

`func (o *JobParams) SetRecommendLabeling(v bool)`

SetRecommendLabeling sets RecommendLabeling field to given value.

### HasRecommendLabeling

`func (o *JobParams) HasRecommendLabeling() bool`

HasRecommendLabeling returns a boolean if a field has been set.

### GetTriggerCreateSampleVisualiztions

`func (o *JobParams) GetTriggerCreateSampleVisualiztions() bool`

GetTriggerCreateSampleVisualiztions returns the TriggerCreateSampleVisualiztions field if non-nil, zero value otherwise.

### GetTriggerCreateSampleVisualiztionsOk

`func (o *JobParams) GetTriggerCreateSampleVisualiztionsOk() (*bool, bool)`

GetTriggerCreateSampleVisualiztionsOk returns a tuple with the TriggerCreateSampleVisualiztions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCreateSampleVisualiztions

`func (o *JobParams) SetTriggerCreateSampleVisualiztions(v bool)`

SetTriggerCreateSampleVisualiztions sets TriggerCreateSampleVisualiztions field to given value.

### HasTriggerCreateSampleVisualiztions

`func (o *JobParams) HasTriggerCreateSampleVisualiztions() bool`

HasTriggerCreateSampleVisualiztions returns a boolean if a field has been set.

### GetDisplayParams

`func (o *JobParams) GetDisplayParams() PopulationExplorationDisplayParams`

GetDisplayParams returns the DisplayParams field if non-nil, zero value otherwise.

### GetDisplayParamsOk

`func (o *JobParams) GetDisplayParamsOk() (*PopulationExplorationDisplayParams, bool)`

GetDisplayParamsOk returns a tuple with the DisplayParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayParams

`func (o *JobParams) SetDisplayParams(v PopulationExplorationDisplayParams)`

SetDisplayParams sets DisplayParams field to given value.


### GetDomainGapMetadata

`func (o *JobParams) GetDomainGapMetadata() string`

GetDomainGapMetadata returns the DomainGapMetadata field if non-nil, zero value otherwise.

### GetDomainGapMetadataOk

`func (o *JobParams) GetDomainGapMetadataOk() (*string, bool)`

GetDomainGapMetadataOk returns a tuple with the DomainGapMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGapMetadata

`func (o *JobParams) SetDomainGapMetadata(v string)`

SetDomainGapMetadata sets DomainGapMetadata field to given value.

### HasDomainGapMetadata

`func (o *JobParams) HasDomainGapMetadata() bool`

HasDomainGapMetadata returns a boolean if a field has been set.

### GetProjectionMetric

`func (o *JobParams) GetProjectionMetric() string`

GetProjectionMetric returns the ProjectionMetric field if non-nil, zero value otherwise.

### GetProjectionMetricOk

`func (o *JobParams) GetProjectionMetricOk() (*string, bool)`

GetProjectionMetricOk returns a tuple with the ProjectionMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectionMetric

`func (o *JobParams) SetProjectionMetric(v string)`

SetProjectionMetric sets ProjectionMetric field to given value.

### HasProjectionMetric

`func (o *JobParams) HasProjectionMetric() bool`

HasProjectionMetric returns a boolean if a field has been set.

### GetExportUrl

`func (o *JobParams) GetExportUrl() string`

GetExportUrl returns the ExportUrl field if non-nil, zero value otherwise.

### GetExportUrlOk

`func (o *JobParams) GetExportUrlOk() (*string, bool)`

GetExportUrlOk returns a tuple with the ExportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUrl

`func (o *JobParams) SetExportUrl(v string)`

SetExportUrl sets ExportUrl field to given value.


### GetProjectVersion

`func (o *JobParams) GetProjectVersion() float64`

GetProjectVersion returns the ProjectVersion field if non-nil, zero value otherwise.

### GetProjectVersionOk

`func (o *JobParams) GetProjectVersionOk() (*float64, bool)`

GetProjectVersionOk returns a tuple with the ProjectVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectVersion

`func (o *JobParams) SetProjectVersion(v float64)`

SetProjectVersion sets ProjectVersion field to given value.


### GetExportOptions

`func (o *JobParams) GetExportOptions() ExportOptions`

GetExportOptions returns the ExportOptions field if non-nil, zero value otherwise.

### GetExportOptionsOk

`func (o *JobParams) GetExportOptionsOk() (*ExportOptions, bool)`

GetExportOptionsOk returns a tuple with the ExportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportOptions

`func (o *JobParams) SetExportOptions(v ExportOptions)`

SetExportOptions sets ExportOptions field to given value.


### GetAlreadyExported

`func (o *JobParams) GetAlreadyExported() bool`

GetAlreadyExported returns the AlreadyExported field if non-nil, zero value otherwise.

### GetAlreadyExportedOk

`func (o *JobParams) GetAlreadyExportedOk() (*bool, bool)`

GetAlreadyExportedOk returns a tuple with the AlreadyExported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyExported

`func (o *JobParams) SetAlreadyExported(v bool)`

SetAlreadyExported sets AlreadyExported field to given value.


### GetProjectExportMeta

`func (o *JobParams) GetProjectExportMeta() ExportProjectMeta`

GetProjectExportMeta returns the ProjectExportMeta field if non-nil, zero value otherwise.

### GetProjectExportMetaOk

`func (o *JobParams) GetProjectExportMetaOk() (*ExportProjectMeta, bool)`

GetProjectExportMetaOk returns a tuple with the ProjectExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectExportMeta

`func (o *JobParams) SetProjectExportMeta(v ExportProjectMeta)`

SetProjectExportMeta sets ProjectExportMeta field to given value.


### GetCopyToUrl

`func (o *JobParams) GetCopyToUrl() string`

GetCopyToUrl returns the CopyToUrl field if non-nil, zero value otherwise.

### GetCopyToUrlOk

`func (o *JobParams) GetCopyToUrlOk() (*string, bool)`

GetCopyToUrlOk returns a tuple with the CopyToUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyToUrl

`func (o *JobParams) SetCopyToUrl(v string)`

SetCopyToUrl sets CopyToUrl field to given value.

### HasCopyToUrl

`func (o *JobParams) HasCopyToUrl() bool`

HasCopyToUrl returns a boolean if a field has been set.

### GetImportUrl

`func (o *JobParams) GetImportUrl() string`

GetImportUrl returns the ImportUrl field if non-nil, zero value otherwise.

### GetImportUrlOk

`func (o *JobParams) GetImportUrlOk() (*string, bool)`

GetImportUrlOk returns a tuple with the ImportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportUrl

`func (o *JobParams) SetImportUrl(v string)`

SetImportUrl sets ImportUrl field to given value.


### GetProjectMeta

`func (o *JobParams) GetProjectMeta() ProjectMeta`

GetProjectMeta returns the ProjectMeta field if non-nil, zero value otherwise.

### GetProjectMetaOk

`func (o *JobParams) GetProjectMetaOk() (*ProjectMeta, bool)`

GetProjectMetaOk returns a tuple with the ProjectMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectMeta

`func (o *JobParams) SetProjectMeta(v ProjectMeta)`

SetProjectMeta sets ProjectMeta field to given value.


### GetDatasetId

`func (o *JobParams) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *JobParams) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *JobParams) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.


### GetSetup

`func (o *JobParams) GetSetup() DatasetSetup`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *JobParams) GetSetupOk() (*DatasetSetup, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *JobParams) SetSetup(v DatasetSetup)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *JobParams) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetSecretManagerId

`func (o *JobParams) GetSecretManagerId() string`

GetSecretManagerId returns the SecretManagerId field if non-nil, zero value otherwise.

### GetSecretManagerIdOk

`func (o *JobParams) GetSecretManagerIdOk() (*string, bool)`

GetSecretManagerIdOk returns a tuple with the SecretManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretManagerId

`func (o *JobParams) SetSecretManagerId(v string)`

SetSecretManagerId sets SecretManagerId field to given value.

### HasSecretManagerId

`func (o *JobParams) HasSecretManagerId() bool`

HasSecretManagerId returns a boolean if a field has been set.

### GetCodeUrl

`func (o *JobParams) GetCodeUrl() string`

GetCodeUrl returns the CodeUrl field if non-nil, zero value otherwise.

### GetCodeUrlOk

`func (o *JobParams) GetCodeUrlOk() (*string, bool)`

GetCodeUrlOk returns a tuple with the CodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeUrl

`func (o *JobParams) SetCodeUrl(v string)`

SetCodeUrl sets CodeUrl field to given value.


### GetCodeEntryFile

`func (o *JobParams) GetCodeEntryFile() string`

GetCodeEntryFile returns the CodeEntryFile field if non-nil, zero value otherwise.

### GetCodeEntryFileOk

`func (o *JobParams) GetCodeEntryFileOk() (*string, bool)`

GetCodeEntryFileOk returns a tuple with the CodeEntryFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeEntryFile

`func (o *JobParams) SetCodeEntryFile(v string)`

SetCodeEntryFile sets CodeEntryFile field to given value.


### GetBranch

`func (o *JobParams) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *JobParams) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *JobParams) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *JobParams) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetNote

`func (o *JobParams) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *JobParams) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *JobParams) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *JobParams) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetGenericBaseImageType

`func (o *JobParams) GetGenericBaseImageType() string`

GetGenericBaseImageType returns the GenericBaseImageType field if non-nil, zero value otherwise.

### GetGenericBaseImageTypeOk

`func (o *JobParams) GetGenericBaseImageTypeOk() (*string, bool)`

GetGenericBaseImageTypeOk returns a tuple with the GenericBaseImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericBaseImageType

`func (o *JobParams) SetGenericBaseImageType(v string)`

SetGenericBaseImageType sets GenericBaseImageType field to given value.

### HasGenericBaseImageType

`func (o *JobParams) HasGenericBaseImageType() bool`

HasGenericBaseImageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


