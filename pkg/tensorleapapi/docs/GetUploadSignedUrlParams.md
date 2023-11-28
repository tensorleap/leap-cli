# GetUploadSignedUrlParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | Pointer to **string** |  | [optional] 
**FileName** | **string** |  | 

## Methods

### NewGetUploadSignedUrlParams

`func NewGetUploadSignedUrlParams(fileName string, ) *GetUploadSignedUrlParams`

NewGetUploadSignedUrlParams instantiates a new GetUploadSignedUrlParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUploadSignedUrlParamsWithDefaults

`func NewGetUploadSignedUrlParamsWithDefaults() *GetUploadSignedUrlParams`

NewGetUploadSignedUrlParamsWithDefaults instantiates a new GetUploadSignedUrlParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigin

`func (o *GetUploadSignedUrlParams) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *GetUploadSignedUrlParams) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *GetUploadSignedUrlParams) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *GetUploadSignedUrlParams) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetFileName

`func (o *GetUploadSignedUrlParams) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *GetUploadSignedUrlParams) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *GetUploadSignedUrlParams) SetFileName(v string)`

SetFileName sets FileName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


