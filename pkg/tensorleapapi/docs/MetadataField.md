# MetadataField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | [**[]MetadataFieldValue**](MetadataFieldValue.md) |  | 
**Type** | **string** |  | 
**Key** | **string** |  | 

## Methods

### NewMetadataField

`func NewMetadataField(values []MetadataFieldValue, type_ string, key string, ) *MetadataField`

NewMetadataField instantiates a new MetadataField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataFieldWithDefaults

`func NewMetadataFieldWithDefaults() *MetadataField`

NewMetadataFieldWithDefaults instantiates a new MetadataField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *MetadataField) GetValues() []MetadataFieldValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MetadataField) GetValuesOk() (*[]MetadataFieldValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MetadataField) SetValues(v []MetadataFieldValue)`

SetValues sets Values field to given value.


### GetType

`func (o *MetadataField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetadataField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetadataField) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *MetadataField) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetadataField) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetadataField) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


