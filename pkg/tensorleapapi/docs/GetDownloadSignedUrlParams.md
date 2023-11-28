# GetDownloadSignedUrlParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | Pointer to **string** |  | [optional] 
**FileName** | **string** |  | 

## Methods

### NewGetDownloadSignedUrlParams

`func NewGetDownloadSignedUrlParams(fileName string, ) *GetDownloadSignedUrlParams`

NewGetDownloadSignedUrlParams instantiates a new GetDownloadSignedUrlParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDownloadSignedUrlParamsWithDefaults

`func NewGetDownloadSignedUrlParamsWithDefaults() *GetDownloadSignedUrlParams`

NewGetDownloadSignedUrlParamsWithDefaults instantiates a new GetDownloadSignedUrlParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigin

`func (o *GetDownloadSignedUrlParams) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *GetDownloadSignedUrlParams) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *GetDownloadSignedUrlParams) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *GetDownloadSignedUrlParams) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetFileName

`func (o *GetDownloadSignedUrlParams) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *GetDownloadSignedUrlParams) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *GetDownloadSignedUrlParams) SetFileName(v string)`

SetFileName sets FileName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


