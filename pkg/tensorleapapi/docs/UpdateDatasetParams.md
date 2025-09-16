# UpdateDatasetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetId** | **string** |  | 
**Name** | **string** |  | 
**Pippin** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateDatasetParams

`func NewUpdateDatasetParams(datasetId string, name string, ) *UpdateDatasetParams`

NewUpdateDatasetParams instantiates a new UpdateDatasetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDatasetParamsWithDefaults

`func NewUpdateDatasetParamsWithDefaults() *UpdateDatasetParams`

NewUpdateDatasetParamsWithDefaults instantiates a new UpdateDatasetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetId

`func (o *UpdateDatasetParams) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *UpdateDatasetParams) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *UpdateDatasetParams) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.


### GetName

`func (o *UpdateDatasetParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDatasetParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDatasetParams) SetName(v string)`

SetName sets Name field to given value.


### GetPippin

`func (o *UpdateDatasetParams) GetPippin() bool`

GetPippin returns the Pippin field if non-nil, zero value otherwise.

### GetPippinOk

`func (o *UpdateDatasetParams) GetPippinOk() (*bool, bool)`

GetPippinOk returns a tuple with the Pippin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPippin

`func (o *UpdateDatasetParams) SetPippin(v bool)`

SetPippin sets Pippin field to given value.

### HasPippin

`func (o *UpdateDatasetParams) HasPippin() bool`

HasPippin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


