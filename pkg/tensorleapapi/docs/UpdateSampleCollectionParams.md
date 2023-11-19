# UpdateSampleCollectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | Pointer to [**[]SampleIdentity**](SampleIdentity.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SampleCollectionId** | **string** |  | 
**ProjectId** | **string** |  | 

## Methods

### NewUpdateSampleCollectionParams

`func NewUpdateSampleCollectionParams(sampleCollectionId string, projectId string, ) *UpdateSampleCollectionParams`

NewUpdateSampleCollectionParams instantiates a new UpdateSampleCollectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSampleCollectionParamsWithDefaults

`func NewUpdateSampleCollectionParamsWithDefaults() *UpdateSampleCollectionParams`

NewUpdateSampleCollectionParamsWithDefaults instantiates a new UpdateSampleCollectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamples

`func (o *UpdateSampleCollectionParams) GetSamples() []SampleIdentity`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *UpdateSampleCollectionParams) GetSamplesOk() (*[]SampleIdentity, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *UpdateSampleCollectionParams) SetSamples(v []SampleIdentity)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *UpdateSampleCollectionParams) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSampleCollectionParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSampleCollectionParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSampleCollectionParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSampleCollectionParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateSampleCollectionParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSampleCollectionParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSampleCollectionParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSampleCollectionParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSampleCollectionId

`func (o *UpdateSampleCollectionParams) GetSampleCollectionId() string`

GetSampleCollectionId returns the SampleCollectionId field if non-nil, zero value otherwise.

### GetSampleCollectionIdOk

`func (o *UpdateSampleCollectionParams) GetSampleCollectionIdOk() (*string, bool)`

GetSampleCollectionIdOk returns a tuple with the SampleCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleCollectionId

`func (o *UpdateSampleCollectionParams) SetSampleCollectionId(v string)`

SetSampleCollectionId sets SampleCollectionId field to given value.


### GetProjectId

`func (o *UpdateSampleCollectionParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateSampleCollectionParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateSampleCollectionParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


