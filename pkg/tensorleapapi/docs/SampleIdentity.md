# SampleIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**DataStateType**](DataStateType.md) |  | 
**Index** | [**SampleIdType**](SampleIdType.md) |  | 
**HashedIndex** | Pointer to **string** |  | [optional] 

## Methods

### NewSampleIdentity

`func NewSampleIdentity(state DataStateType, index SampleIdType, ) *SampleIdentity`

NewSampleIdentity instantiates a new SampleIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleIdentityWithDefaults

`func NewSampleIdentityWithDefaults() *SampleIdentity`

NewSampleIdentityWithDefaults instantiates a new SampleIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *SampleIdentity) GetState() DataStateType`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SampleIdentity) GetStateOk() (*DataStateType, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SampleIdentity) SetState(v DataStateType)`

SetState sets State field to given value.


### GetIndex

`func (o *SampleIdentity) GetIndex() SampleIdType`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SampleIdentity) GetIndexOk() (*SampleIdType, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SampleIdentity) SetIndex(v SampleIdType)`

SetIndex sets Index field to given value.


### GetHashedIndex

`func (o *SampleIdentity) GetHashedIndex() string`

GetHashedIndex returns the HashedIndex field if non-nil, zero value otherwise.

### GetHashedIndexOk

`func (o *SampleIdentity) GetHashedIndexOk() (*string, bool)`

GetHashedIndexOk returns a tuple with the HashedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedIndex

`func (o *SampleIdentity) SetHashedIndex(v string)`

SetHashedIndex sets HashedIndex field to given value.

### HasHashedIndex

`func (o *SampleIdentity) HasHashedIndex() bool`

HasHashedIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


