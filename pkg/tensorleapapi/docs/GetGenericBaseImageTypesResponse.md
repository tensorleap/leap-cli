# GetGenericBaseImageTypesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseImageTypes** | [**[]GenericBaseImage**](GenericBaseImage.md) |  | 
**DefaultBaseImageType** | **string** |  | 

## Methods

### NewGetGenericBaseImageTypesResponse

`func NewGetGenericBaseImageTypesResponse(baseImageTypes []GenericBaseImage, defaultBaseImageType string, ) *GetGenericBaseImageTypesResponse`

NewGetGenericBaseImageTypesResponse instantiates a new GetGenericBaseImageTypesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGenericBaseImageTypesResponseWithDefaults

`func NewGetGenericBaseImageTypesResponseWithDefaults() *GetGenericBaseImageTypesResponse`

NewGetGenericBaseImageTypesResponseWithDefaults instantiates a new GetGenericBaseImageTypesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseImageTypes

`func (o *GetGenericBaseImageTypesResponse) GetBaseImageTypes() []GenericBaseImage`

GetBaseImageTypes returns the BaseImageTypes field if non-nil, zero value otherwise.

### GetBaseImageTypesOk

`func (o *GetGenericBaseImageTypesResponse) GetBaseImageTypesOk() (*[]GenericBaseImage, bool)`

GetBaseImageTypesOk returns a tuple with the BaseImageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseImageTypes

`func (o *GetGenericBaseImageTypesResponse) SetBaseImageTypes(v []GenericBaseImage)`

SetBaseImageTypes sets BaseImageTypes field to given value.


### GetDefaultBaseImageType

`func (o *GetGenericBaseImageTypesResponse) GetDefaultBaseImageType() string`

GetDefaultBaseImageType returns the DefaultBaseImageType field if non-nil, zero value otherwise.

### GetDefaultBaseImageTypeOk

`func (o *GetGenericBaseImageTypesResponse) GetDefaultBaseImageTypeOk() (*string, bool)`

GetDefaultBaseImageTypeOk returns a tuple with the DefaultBaseImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBaseImageType

`func (o *GetGenericBaseImageTypesResponse) SetDefaultBaseImageType(v string)`

SetDefaultBaseImageType sets DefaultBaseImageType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


