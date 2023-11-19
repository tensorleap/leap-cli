# AddSampleCollectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | [**[]SampleIdentity**](SampleIdentity.md) |  | 
**ProjectId** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewAddSampleCollectionParams

`func NewAddSampleCollectionParams(samples []SampleIdentity, projectId string, name string, ) *AddSampleCollectionParams`

NewAddSampleCollectionParams instantiates a new AddSampleCollectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSampleCollectionParamsWithDefaults

`func NewAddSampleCollectionParamsWithDefaults() *AddSampleCollectionParams`

NewAddSampleCollectionParamsWithDefaults instantiates a new AddSampleCollectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamples

`func (o *AddSampleCollectionParams) GetSamples() []SampleIdentity`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *AddSampleCollectionParams) GetSamplesOk() (*[]SampleIdentity, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *AddSampleCollectionParams) SetSamples(v []SampleIdentity)`

SetSamples sets Samples field to given value.


### GetProjectId

`func (o *AddSampleCollectionParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AddSampleCollectionParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AddSampleCollectionParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetDescription

`func (o *AddSampleCollectionParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSampleCollectionParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSampleCollectionParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSampleCollectionParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AddSampleCollectionParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddSampleCollectionParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddSampleCollectionParams) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


