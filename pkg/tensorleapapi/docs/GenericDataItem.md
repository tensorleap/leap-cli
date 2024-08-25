# GenericDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 
**Data** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 
**InnerKey** | Pointer to [**NumberOrString**](NumberOrString.md) |  | [optional] 

## Methods

### NewGenericDataItem

`func NewGenericDataItem(data map[string]interface{}, ) *GenericDataItem`

NewGenericDataItem instantiates a new GenericDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericDataItemWithDefaults

`func NewGenericDataItemWithDefaults() *GenericDataItem`

NewGenericDataItemWithDefaults instantiates a new GenericDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GenericDataItem) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GenericDataItem) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GenericDataItem) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GenericDataItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetData

`func (o *GenericDataItem) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GenericDataItem) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GenericDataItem) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetInnerKey

`func (o *GenericDataItem) GetInnerKey() NumberOrString`

GetInnerKey returns the InnerKey field if non-nil, zero value otherwise.

### GetInnerKeyOk

`func (o *GenericDataItem) GetInnerKeyOk() (*NumberOrString, bool)`

GetInnerKeyOk returns a tuple with the InnerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerKey

`func (o *GenericDataItem) SetInnerKey(v NumberOrString)`

SetInnerKey sets InnerKey field to given value.

### HasInnerKey

`func (o *GenericDataItem) HasInnerKey() bool`

HasInnerKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


