# GenericDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GenericDataItem**](GenericDataItem.md) |  | 
**ReferenceStats** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewGenericDataResponse

`func NewGenericDataResponse(data []GenericDataItem, ) *GenericDataResponse`

NewGenericDataResponse instantiates a new GenericDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericDataResponseWithDefaults

`func NewGenericDataResponseWithDefaults() *GenericDataResponse`

NewGenericDataResponseWithDefaults instantiates a new GenericDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GenericDataResponse) GetData() []GenericDataItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GenericDataResponse) GetDataOk() (*[]GenericDataItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GenericDataResponse) SetData(v []GenericDataItem)`

SetData sets Data field to given value.


### GetReferenceStats

`func (o *GenericDataResponse) GetReferenceStats() map[string]interface{}`

GetReferenceStats returns the ReferenceStats field if non-nil, zero value otherwise.

### GetReferenceStatsOk

`func (o *GenericDataResponse) GetReferenceStatsOk() (*map[string]interface{}, bool)`

GetReferenceStatsOk returns a tuple with the ReferenceStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceStats

`func (o *GenericDataResponse) SetReferenceStats(v map[string]interface{})`

SetReferenceStats sets ReferenceStats field to given value.

### HasReferenceStats

`func (o *GenericDataResponse) HasReferenceStats() bool`

HasReferenceStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


