# ValidateGraphResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** |  | 
**Data** | Pointer to [**GraphValidatorData**](GraphValidatorData.md) |  | [optional] 
**Status** | [**JobStatus**](JobStatus.md) |  | 

## Methods

### NewValidateGraphResponse

`func NewValidateGraphResponse(digest string, status JobStatus, ) *ValidateGraphResponse`

NewValidateGraphResponse instantiates a new ValidateGraphResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateGraphResponseWithDefaults

`func NewValidateGraphResponseWithDefaults() *ValidateGraphResponse`

NewValidateGraphResponseWithDefaults instantiates a new ValidateGraphResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *ValidateGraphResponse) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ValidateGraphResponse) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ValidateGraphResponse) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetData

`func (o *ValidateGraphResponse) GetData() GraphValidatorData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ValidateGraphResponse) GetDataOk() (*GraphValidatorData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ValidateGraphResponse) SetData(v GraphValidatorData)`

SetData sets Data field to given value.

### HasData

`func (o *ValidateGraphResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *ValidateGraphResponse) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ValidateGraphResponse) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ValidateGraphResponse) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


