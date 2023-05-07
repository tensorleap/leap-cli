# VisualizationEntityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**JobId** | **string** |  | 
**JobParams** | Pointer to [**JobParams**](JobParams.md) |  | [optional] 
**Epoch** | **float64** |  | 
**BlobName** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Version** | **string** |  | 
**SessionRunId** | **string** |  | 
**VisualizationUuid** | **string** |  | 
**OrganizationId** | **string** |  | 
**Type** | [**AnalyzeTypeEnum**](AnalyzeTypeEnum.md) |  | 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Layout** | Pointer to [**SizedLayout**](SizedLayout.md) |  | [optional] 

## Methods

### NewVisualizationEntityResponse

`func NewVisualizationEntityResponse(id string, jobId string, epoch float64, blobName string, createdAt time.Time, version string, sessionRunId string, visualizationUuid string, organizationId string, type_ AnalyzeTypeEnum, ) *VisualizationEntityResponse`

NewVisualizationEntityResponse instantiates a new VisualizationEntityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualizationEntityResponseWithDefaults

`func NewVisualizationEntityResponseWithDefaults() *VisualizationEntityResponse`

NewVisualizationEntityResponseWithDefaults instantiates a new VisualizationEntityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VisualizationEntityResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VisualizationEntityResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VisualizationEntityResponse) SetId(v string)`

SetId sets Id field to given value.


### GetJobId

`func (o *VisualizationEntityResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *VisualizationEntityResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *VisualizationEntityResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetJobParams

`func (o *VisualizationEntityResponse) GetJobParams() JobParams`

GetJobParams returns the JobParams field if non-nil, zero value otherwise.

### GetJobParamsOk

`func (o *VisualizationEntityResponse) GetJobParamsOk() (*JobParams, bool)`

GetJobParamsOk returns a tuple with the JobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobParams

`func (o *VisualizationEntityResponse) SetJobParams(v JobParams)`

SetJobParams sets JobParams field to given value.

### HasJobParams

`func (o *VisualizationEntityResponse) HasJobParams() bool`

HasJobParams returns a boolean if a field has been set.

### GetEpoch

`func (o *VisualizationEntityResponse) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *VisualizationEntityResponse) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *VisualizationEntityResponse) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetBlobName

`func (o *VisualizationEntityResponse) GetBlobName() string`

GetBlobName returns the BlobName field if non-nil, zero value otherwise.

### GetBlobNameOk

`func (o *VisualizationEntityResponse) GetBlobNameOk() (*string, bool)`

GetBlobNameOk returns a tuple with the BlobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobName

`func (o *VisualizationEntityResponse) SetBlobName(v string)`

SetBlobName sets BlobName field to given value.


### GetCreatedAt

`func (o *VisualizationEntityResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VisualizationEntityResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VisualizationEntityResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetVersion

`func (o *VisualizationEntityResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VisualizationEntityResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VisualizationEntityResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetSessionRunId

`func (o *VisualizationEntityResponse) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *VisualizationEntityResponse) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *VisualizationEntityResponse) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetVisualizationUuid

`func (o *VisualizationEntityResponse) GetVisualizationUuid() string`

GetVisualizationUuid returns the VisualizationUuid field if non-nil, zero value otherwise.

### GetVisualizationUuidOk

`func (o *VisualizationEntityResponse) GetVisualizationUuidOk() (*string, bool)`

GetVisualizationUuidOk returns a tuple with the VisualizationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizationUuid

`func (o *VisualizationEntityResponse) SetVisualizationUuid(v string)`

SetVisualizationUuid sets VisualizationUuid field to given value.


### GetOrganizationId

`func (o *VisualizationEntityResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *VisualizationEntityResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *VisualizationEntityResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetType

`func (o *VisualizationEntityResponse) GetType() AnalyzeTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VisualizationEntityResponse) GetTypeOk() (*AnalyzeTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VisualizationEntityResponse) SetType(v AnalyzeTypeEnum)`

SetType sets Type field to given value.


### GetCreatedBy

`func (o *VisualizationEntityResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VisualizationEntityResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VisualizationEntityResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *VisualizationEntityResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLayout

`func (o *VisualizationEntityResponse) GetLayout() SizedLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *VisualizationEntityResponse) GetLayoutOk() (*SizedLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *VisualizationEntityResponse) SetLayout(v SizedLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *VisualizationEntityResponse) HasLayout() bool`

HasLayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


