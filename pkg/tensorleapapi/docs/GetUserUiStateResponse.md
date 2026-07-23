# GetUserUiStateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **map[string]interface{}** |  | 
**BlobPath** | **string** |  | 

## Methods

### NewGetUserUiStateResponse

`func NewGetUserUiStateResponse(data map[string]interface{}, blobPath string, ) *GetUserUiStateResponse`

NewGetUserUiStateResponse instantiates a new GetUserUiStateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserUiStateResponseWithDefaults

`func NewGetUserUiStateResponseWithDefaults() *GetUserUiStateResponse`

NewGetUserUiStateResponseWithDefaults instantiates a new GetUserUiStateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetUserUiStateResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUserUiStateResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUserUiStateResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *GetUserUiStateResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GetUserUiStateResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetBlobPath

`func (o *GetUserUiStateResponse) GetBlobPath() string`

GetBlobPath returns the BlobPath field if non-nil, zero value otherwise.

### GetBlobPathOk

`func (o *GetUserUiStateResponse) GetBlobPathOk() (*string, bool)`

GetBlobPathOk returns a tuple with the BlobPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobPath

`func (o *GetUserUiStateResponse) SetBlobPath(v string)`

SetBlobPath sets BlobPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


