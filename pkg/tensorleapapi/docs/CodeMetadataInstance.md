# CodeMetadataInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | [**CodeMetadataType**](CodeMetadataType.md) |  | 

## Methods

### NewCodeMetadataInstance

`func NewCodeMetadataInstance(name string, type_ CodeMetadataType, ) *CodeMetadataInstance`

NewCodeMetadataInstance instantiates a new CodeMetadataInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeMetadataInstanceWithDefaults

`func NewCodeMetadataInstanceWithDefaults() *CodeMetadataInstance`

NewCodeMetadataInstanceWithDefaults instantiates a new CodeMetadataInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CodeMetadataInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodeMetadataInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodeMetadataInstance) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CodeMetadataInstance) GetType() CodeMetadataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CodeMetadataInstance) GetTypeOk() (*CodeMetadataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CodeMetadataInstance) SetType(v CodeMetadataType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


