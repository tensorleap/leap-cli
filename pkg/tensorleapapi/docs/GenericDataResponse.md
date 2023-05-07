# GenericDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GenericDataItem**](GenericDataItem.md) |  | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


