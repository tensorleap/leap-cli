# CollectionFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to [**CollectionFilterSpecRange**](CollectionFilterSpecRange.md) |  | [optional] 
**HiddenValues** | Pointer to [**[]CollectionFilterSpecHiddenValuesInner**](CollectionFilterSpecHiddenValuesInner.md) |  | [optional] 
**Field** | **string** |  | 

## Methods

### NewCollectionFilterSpec

`func NewCollectionFilterSpec(field string, ) *CollectionFilterSpec`

NewCollectionFilterSpec instantiates a new CollectionFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionFilterSpecWithDefaults

`func NewCollectionFilterSpecWithDefaults() *CollectionFilterSpec`

NewCollectionFilterSpecWithDefaults instantiates a new CollectionFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *CollectionFilterSpec) GetRange() CollectionFilterSpecRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CollectionFilterSpec) GetRangeOk() (*CollectionFilterSpecRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CollectionFilterSpec) SetRange(v CollectionFilterSpecRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *CollectionFilterSpec) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetHiddenValues

`func (o *CollectionFilterSpec) GetHiddenValues() []CollectionFilterSpecHiddenValuesInner`

GetHiddenValues returns the HiddenValues field if non-nil, zero value otherwise.

### GetHiddenValuesOk

`func (o *CollectionFilterSpec) GetHiddenValuesOk() (*[]CollectionFilterSpecHiddenValuesInner, bool)`

GetHiddenValuesOk returns a tuple with the HiddenValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenValues

`func (o *CollectionFilterSpec) SetHiddenValues(v []CollectionFilterSpecHiddenValuesInner)`

SetHiddenValues sets HiddenValues field to given value.

### HasHiddenValues

`func (o *CollectionFilterSpec) HasHiddenValues() bool`

HasHiddenValues returns a boolean if a field has been set.

### GetField

`func (o *CollectionFilterSpec) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *CollectionFilterSpec) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *CollectionFilterSpec) SetField(v string)`

SetField sets Field field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


