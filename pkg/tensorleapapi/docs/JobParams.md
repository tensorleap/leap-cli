# JobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Epochs** | **float64** |  | 
**BatchSize** | **float64** |  | 
**EarlyStopParams** | Pointer to [**EarlyStopParams**](EarlyStopParams.md) |  | [optional] 
**Type** | **string** |  | 
**FromEpoch** | **float64** |  | 
**BatchSize** | **float64** |  | 
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
**ProjectionMetric** | Pointer to **string** |  | [optional] 
**NumOfSamples** | **float64** |  | 
**ExportUrl** | **string** |  | 
**ProjectVersion** | **float64** |  | 
**CopyToUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewJobParams

`func NewJobParams(epochs float64, batchSize float64, type_ string, fromEpoch float64, batchSize float64, extId string, title string, epoch float64, digest string, sampleIds []SampleIdentity, limit float64, sessionRunId string, numOfSamples float64, exportUrl string, projectVersion float64, ) *JobParams`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


