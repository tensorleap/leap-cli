# AnalyzeGraphResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** |  | 
**Response** | Pointer to [**GraphAnalyzerData**](GraphAnalyzerData.md) |  | [optional] 
**Status** | [**JobStatus**](JobStatus.md) |  | 

## Methods

### NewAnalyzeGraphResponse

`func NewAnalyzeGraphResponse(digest string, status JobStatus, ) *AnalyzeGraphResponse`

NewAnalyzeGraphResponse instantiates a new AnalyzeGraphResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyzeGraphResponseWithDefaults

`func NewAnalyzeGraphResponseWithDefaults() *AnalyzeGraphResponse`

NewAnalyzeGraphResponseWithDefaults instantiates a new AnalyzeGraphResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *AnalyzeGraphResponse) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *AnalyzeGraphResponse) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *AnalyzeGraphResponse) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetResponse

`func (o *AnalyzeGraphResponse) GetResponse() GraphAnalyzerData`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *AnalyzeGraphResponse) GetResponseOk() (*GraphAnalyzerData, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *AnalyzeGraphResponse) SetResponse(v GraphAnalyzerData)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *AnalyzeGraphResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetStatus

`func (o *AnalyzeGraphResponse) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnalyzeGraphResponse) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnalyzeGraphResponse) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


