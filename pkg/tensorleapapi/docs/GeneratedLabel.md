# GeneratedLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**JobId** | **string** |  | 
**SessionRunId** | **string** |  | 
**SessionRunName** | **string** |  | 
**Epoch** | **float64** |  | 
**NumOfSamples** | Pointer to **float64** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**FilePath** | **string** |  | 
**FileUrl** | **string** |  | 
**Status** | [**JobStatus**](JobStatus.md) |  | 
**IsDeleted** | **bool** |  | 

## Methods

### NewGeneratedLabel

`func NewGeneratedLabel(id string, jobId string, sessionRunId string, sessionRunName string, epoch float64, createdAt time.Time, createdBy string, filePath string, fileUrl string, status JobStatus, isDeleted bool, ) *GeneratedLabel`

NewGeneratedLabel instantiates a new GeneratedLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneratedLabelWithDefaults

`func NewGeneratedLabelWithDefaults() *GeneratedLabel`

NewGeneratedLabelWithDefaults instantiates a new GeneratedLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GeneratedLabel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GeneratedLabel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GeneratedLabel) SetId(v string)`

SetId sets Id field to given value.


### GetJobId

`func (o *GeneratedLabel) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *GeneratedLabel) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *GeneratedLabel) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetSessionRunId

`func (o *GeneratedLabel) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *GeneratedLabel) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *GeneratedLabel) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetSessionRunName

`func (o *GeneratedLabel) GetSessionRunName() string`

GetSessionRunName returns the SessionRunName field if non-nil, zero value otherwise.

### GetSessionRunNameOk

`func (o *GeneratedLabel) GetSessionRunNameOk() (*string, bool)`

GetSessionRunNameOk returns a tuple with the SessionRunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunName

`func (o *GeneratedLabel) SetSessionRunName(v string)`

SetSessionRunName sets SessionRunName field to given value.


### GetEpoch

`func (o *GeneratedLabel) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *GeneratedLabel) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *GeneratedLabel) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetNumOfSamples

`func (o *GeneratedLabel) GetNumOfSamples() float64`

GetNumOfSamples returns the NumOfSamples field if non-nil, zero value otherwise.

### GetNumOfSamplesOk

`func (o *GeneratedLabel) GetNumOfSamplesOk() (*float64, bool)`

GetNumOfSamplesOk returns a tuple with the NumOfSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfSamples

`func (o *GeneratedLabel) SetNumOfSamples(v float64)`

SetNumOfSamples sets NumOfSamples field to given value.

### HasNumOfSamples

`func (o *GeneratedLabel) HasNumOfSamples() bool`

HasNumOfSamples returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GeneratedLabel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GeneratedLabel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GeneratedLabel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *GeneratedLabel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GeneratedLabel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GeneratedLabel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetFilePath

`func (o *GeneratedLabel) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *GeneratedLabel) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *GeneratedLabel) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.


### GetFileUrl

`func (o *GeneratedLabel) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *GeneratedLabel) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *GeneratedLabel) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetStatus

`func (o *GeneratedLabel) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GeneratedLabel) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GeneratedLabel) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.


### GetIsDeleted

`func (o *GeneratedLabel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *GeneratedLabel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *GeneratedLabel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


