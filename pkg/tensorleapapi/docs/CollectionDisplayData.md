# CollectionDisplayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SampleIdentitiesPreview** | [**[]SampleIdentity**](SampleIdentity.md) |  | 
**SamplesCount** | **float64** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewCollectionDisplayData

`func NewCollectionDisplayData(sampleIdentitiesPreview []SampleIdentity, samplesCount float64, name string, ) *CollectionDisplayData`

NewCollectionDisplayData instantiates a new CollectionDisplayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionDisplayDataWithDefaults

`func NewCollectionDisplayDataWithDefaults() *CollectionDisplayData`

NewCollectionDisplayDataWithDefaults instantiates a new CollectionDisplayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSampleIdentitiesPreview

`func (o *CollectionDisplayData) GetSampleIdentitiesPreview() []SampleIdentity`

GetSampleIdentitiesPreview returns the SampleIdentitiesPreview field if non-nil, zero value otherwise.

### GetSampleIdentitiesPreviewOk

`func (o *CollectionDisplayData) GetSampleIdentitiesPreviewOk() (*[]SampleIdentity, bool)`

GetSampleIdentitiesPreviewOk returns a tuple with the SampleIdentitiesPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleIdentitiesPreview

`func (o *CollectionDisplayData) SetSampleIdentitiesPreview(v []SampleIdentity)`

SetSampleIdentitiesPreview sets SampleIdentitiesPreview field to given value.


### GetSamplesCount

`func (o *CollectionDisplayData) GetSamplesCount() float64`

GetSamplesCount returns the SamplesCount field if non-nil, zero value otherwise.

### GetSamplesCountOk

`func (o *CollectionDisplayData) GetSamplesCountOk() (*float64, bool)`

GetSamplesCountOk returns a tuple with the SamplesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplesCount

`func (o *CollectionDisplayData) SetSamplesCount(v float64)`

SetSamplesCount sets SamplesCount field to given value.


### GetDescription

`func (o *CollectionDisplayData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionDisplayData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionDisplayData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CollectionDisplayData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CollectionDisplayData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionDisplayData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionDisplayData) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


