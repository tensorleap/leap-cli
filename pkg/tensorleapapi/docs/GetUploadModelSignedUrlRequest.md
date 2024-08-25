# GetUploadModelSignedUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Epoch** | **float64** |  | 
**Origin** | Pointer to **string** |  | [optional] 
**ExperimentId** | **string** |  | 
**FileType** | [**UploadModelFileType**](UploadModelFileType.md) |  | 
**VersionId** | **string** |  | 
**ProjectId** | **string** |  | 

## Methods

### NewGetUploadModelSignedUrlRequest

`func NewGetUploadModelSignedUrlRequest(epoch float64, experimentId string, fileType UploadModelFileType, versionId string, projectId string, ) *GetUploadModelSignedUrlRequest`

NewGetUploadModelSignedUrlRequest instantiates a new GetUploadModelSignedUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUploadModelSignedUrlRequestWithDefaults

`func NewGetUploadModelSignedUrlRequestWithDefaults() *GetUploadModelSignedUrlRequest`

NewGetUploadModelSignedUrlRequestWithDefaults instantiates a new GetUploadModelSignedUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpoch

`func (o *GetUploadModelSignedUrlRequest) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *GetUploadModelSignedUrlRequest) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *GetUploadModelSignedUrlRequest) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetOrigin

`func (o *GetUploadModelSignedUrlRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *GetUploadModelSignedUrlRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *GetUploadModelSignedUrlRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *GetUploadModelSignedUrlRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetExperimentId

`func (o *GetUploadModelSignedUrlRequest) GetExperimentId() string`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *GetUploadModelSignedUrlRequest) GetExperimentIdOk() (*string, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *GetUploadModelSignedUrlRequest) SetExperimentId(v string)`

SetExperimentId sets ExperimentId field to given value.


### GetFileType

`func (o *GetUploadModelSignedUrlRequest) GetFileType() UploadModelFileType`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *GetUploadModelSignedUrlRequest) GetFileTypeOk() (*UploadModelFileType, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *GetUploadModelSignedUrlRequest) SetFileType(v UploadModelFileType)`

SetFileType sets FileType field to given value.


### GetVersionId

`func (o *GetUploadModelSignedUrlRequest) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *GetUploadModelSignedUrlRequest) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *GetUploadModelSignedUrlRequest) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetProjectId

`func (o *GetUploadModelSignedUrlRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetUploadModelSignedUrlRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetUploadModelSignedUrlRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


