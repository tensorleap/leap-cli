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
**Monitor** | Pointer to **bool** |  | [optional] 
**EvaluatedEpoch** | **float64** |  | 
**SessionId** | **string** |  | 
**Type** | [**ExportModelTypeEnum**](ExportModelTypeEnum.md) |  | 
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
**LatentSpaceType** | Pointer to **string** |  | [optional] 
**ForceExecute** | Pointer to **bool** |  | [optional] 
**ElementInstance** | Pointer to **bool** |  | [optional] 
**UseCustomLatentSpace** | Pointer to **bool** |  | [optional] 
**OptionalAnalysis** | Pointer to [**[]OptionalAnalysis**](OptionalAnalysis.md) |  | [optional] 
**ReductionAlgorithm** | [**ReductionAlgorithm**](ReductionAlgorithm.md) |  | 
**ShouldFillRemainingWithUnbalanced** | **bool** |  | 
**BalanceBy** | **[]string** |  | 
**NumOfSamples** | **float64** |  | 
**DomainGapMetadata** | Pointer to **string** |  | [optional] 
**ProjectionMetric** | Pointer to **string** |  | [optional] 
**LabelingAlgorithm** | [**LabelingAlgorithm**](LabelingAlgorithm.md) |  | 
**NumOfSamplesToLabel** | Pointer to **float64** |  | [optional] 
**ExportUrl** | **string** |  | 
**ProjectVersion** | **float64** |  | 
**ExportOptions** | [**ExportOptions**](ExportOptions.md) |  | 
**AlreadyExported** | **bool** |  | 
**ProjectExportMeta** | [**ExportProjectMeta**](ExportProjectMeta.md) |  | 
**CopyToUrl** | Pointer to **string** |  | [optional] 
**ImportUrl** | **string** |  | 
**ProjectMeta** | [**ProjectMeta**](ProjectMeta.md) |  | 
**SecretManagerId** | Pointer to **string** |  | [optional] 
**CodeUrl** | **string** |  | 
**CodeEntryFile** | **string** |  | 
**GenericBaseImageType** | Pointer to **string** |  | [optional] 
**VersionName** | **string** |  | 
**BranchName** | Pointer to **string** |  | [optional] 
**OverwriteVersionId** | Pointer to **string** |  | [optional] 

## Methods

### NewJobParams

`func NewJobParams(epochs float64, batchSize float64, versionId string, projectId string, batchSize float64, dataStates []DataStateType, name string, description string, evaluatedEpoch float64, sessionId string, type_ ExportModelTypeEnum, fromEpoch float64, extId string, title string, epoch float64, digest string, sampleIds []SampleIdentity, limit float64, sessionRunId string, reductionAlgorithm ReductionAlgorithm, shouldFillRemainingWithUnbalanced bool, balanceBy []string, numOfSamples float64, labelingAlgorithm LabelingAlgorithm, exportUrl string, projectVersion float64, exportOptions ExportOptions, alreadyExported bool, projectExportMeta ExportProjectMeta, importUrl string, projectMeta ProjectMeta, codeUrl string, codeEntryFile string, versionName string, ) *JobParams`

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

`func (o *JobParams) GetType() ExportModelTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobParams) GetTypeOk() (*ExportModelTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobParams) SetType(v ExportModelTypeEnum)`

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


### GetLatentSpaceType

`func (o *JobParams) GetLatentSpaceType() string`

GetLatentSpaceType returns the LatentSpaceType field if non-nil, zero value otherwise.

### GetLatentSpaceTypeOk

`func (o *JobParams) GetLatentSpaceTypeOk() (*string, bool)`

GetLatentSpaceTypeOk returns a tuple with the LatentSpaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatentSpaceType

`func (o *JobParams) SetLatentSpaceType(v string)`

SetLatentSpaceType sets LatentSpaceType field to given value.

### HasLatentSpaceType

`func (o *JobParams) HasLatentSpaceType() bool`

HasLatentSpaceType returns a boolean if a field has been set.

### GetForceExecute

`func (o *JobParams) GetForceExecute() bool`

GetForceExecute returns the ForceExecute field if non-nil, zero value otherwise.

### GetForceExecuteOk

`func (o *JobParams) GetForceExecuteOk() (*bool, bool)`

GetForceExecuteOk returns a tuple with the ForceExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceExecute

`func (o *JobParams) SetForceExecute(v bool)`

SetForceExecute sets ForceExecute field to given value.

### HasForceExecute

`func (o *JobParams) HasForceExecute() bool`

HasForceExecute returns a boolean if a field has been set.

### GetElementInstance

`func (o *JobParams) GetElementInstance() bool`

GetElementInstance returns the ElementInstance field if non-nil, zero value otherwise.

### GetElementInstanceOk

`func (o *JobParams) GetElementInstanceOk() (*bool, bool)`

GetElementInstanceOk returns a tuple with the ElementInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementInstance

`func (o *JobParams) SetElementInstance(v bool)`

SetElementInstance sets ElementInstance field to given value.

### HasElementInstance

`func (o *JobParams) HasElementInstance() bool`

HasElementInstance returns a boolean if a field has been set.

### GetUseCustomLatentSpace

`func (o *JobParams) GetUseCustomLatentSpace() bool`

GetUseCustomLatentSpace returns the UseCustomLatentSpace field if non-nil, zero value otherwise.

### GetUseCustomLatentSpaceOk

`func (o *JobParams) GetUseCustomLatentSpaceOk() (*bool, bool)`

GetUseCustomLatentSpaceOk returns a tuple with the UseCustomLatentSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomLatentSpace

`func (o *JobParams) SetUseCustomLatentSpace(v bool)`

SetUseCustomLatentSpace sets UseCustomLatentSpace field to given value.

### HasUseCustomLatentSpace

`func (o *JobParams) HasUseCustomLatentSpace() bool`

HasUseCustomLatentSpace returns a boolean if a field has been set.

### GetOptionalAnalysis

`func (o *JobParams) GetOptionalAnalysis() []OptionalAnalysis`

GetOptionalAnalysis returns the OptionalAnalysis field if non-nil, zero value otherwise.

### GetOptionalAnalysisOk

`func (o *JobParams) GetOptionalAnalysisOk() (*[]OptionalAnalysis, bool)`

GetOptionalAnalysisOk returns a tuple with the OptionalAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalAnalysis

`func (o *JobParams) SetOptionalAnalysis(v []OptionalAnalysis)`

SetOptionalAnalysis sets OptionalAnalysis field to given value.

### HasOptionalAnalysis

`func (o *JobParams) HasOptionalAnalysis() bool`

HasOptionalAnalysis returns a boolean if a field has been set.

### GetReductionAlgorithm

`func (o *JobParams) GetReductionAlgorithm() ReductionAlgorithm`

GetReductionAlgorithm returns the ReductionAlgorithm field if non-nil, zero value otherwise.

### GetReductionAlgorithmOk

`func (o *JobParams) GetReductionAlgorithmOk() (*ReductionAlgorithm, bool)`

GetReductionAlgorithmOk returns a tuple with the ReductionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReductionAlgorithm

`func (o *JobParams) SetReductionAlgorithm(v ReductionAlgorithm)`

SetReductionAlgorithm sets ReductionAlgorithm field to given value.


### GetShouldFillRemainingWithUnbalanced

`func (o *JobParams) GetShouldFillRemainingWithUnbalanced() bool`

GetShouldFillRemainingWithUnbalanced returns the ShouldFillRemainingWithUnbalanced field if non-nil, zero value otherwise.

### GetShouldFillRemainingWithUnbalancedOk

`func (o *JobParams) GetShouldFillRemainingWithUnbalancedOk() (*bool, bool)`

GetShouldFillRemainingWithUnbalancedOk returns a tuple with the ShouldFillRemainingWithUnbalanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldFillRemainingWithUnbalanced

`func (o *JobParams) SetShouldFillRemainingWithUnbalanced(v bool)`

SetShouldFillRemainingWithUnbalanced sets ShouldFillRemainingWithUnbalanced field to given value.


### GetBalanceBy

`func (o *JobParams) GetBalanceBy() []string`

GetBalanceBy returns the BalanceBy field if non-nil, zero value otherwise.

### GetBalanceByOk

`func (o *JobParams) GetBalanceByOk() (*[]string, bool)`

GetBalanceByOk returns a tuple with the BalanceBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceBy

`func (o *JobParams) SetBalanceBy(v []string)`

SetBalanceBy sets BalanceBy field to given value.


### GetNumOfSamples

`func (o *JobParams) GetNumOfSamples() float64`

GetNumOfSamples returns the NumOfSamples field if non-nil, zero value otherwise.

### GetNumOfSamplesOk

`func (o *JobParams) GetNumOfSamplesOk() (*float64, bool)`

GetNumOfSamplesOk returns a tuple with the NumOfSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfSamples

`func (o *JobParams) SetNumOfSamples(v float64)`

SetNumOfSamples sets NumOfSamples field to given value.


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

### GetLabelingAlgorithm

`func (o *JobParams) GetLabelingAlgorithm() LabelingAlgorithm`

GetLabelingAlgorithm returns the LabelingAlgorithm field if non-nil, zero value otherwise.

### GetLabelingAlgorithmOk

`func (o *JobParams) GetLabelingAlgorithmOk() (*LabelingAlgorithm, bool)`

GetLabelingAlgorithmOk returns a tuple with the LabelingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelingAlgorithm

`func (o *JobParams) SetLabelingAlgorithm(v LabelingAlgorithm)`

SetLabelingAlgorithm sets LabelingAlgorithm field to given value.


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

### GetVersionName

`func (o *JobParams) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *JobParams) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *JobParams) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.


### GetBranchName

`func (o *JobParams) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *JobParams) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *JobParams) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *JobParams) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetOverwriteVersionId

`func (o *JobParams) GetOverwriteVersionId() string`

GetOverwriteVersionId returns the OverwriteVersionId field if non-nil, zero value otherwise.

### GetOverwriteVersionIdOk

`func (o *JobParams) GetOverwriteVersionIdOk() (*string, bool)`

GetOverwriteVersionIdOk returns a tuple with the OverwriteVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteVersionId

`func (o *JobParams) SetOverwriteVersionId(v string)`

SetOverwriteVersionId sets OverwriteVersionId field to given value.

### HasOverwriteVersionId

`func (o *JobParams) HasOverwriteVersionId() bool`

HasOverwriteVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


