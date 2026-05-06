# EnsureCollectionIndexResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**SamplesCount** | **float64** |  | 
**EsIndexedCount** | **float64** |  | 
**Status** | [**CollectionIndexStatus**](CollectionIndexStatus.md) |  | 

## Methods

### NewEnsureCollectionIndexResponse

`func NewEnsureCollectionIndexResponse(samplesCount float64, esIndexedCount float64, status CollectionIndexStatus, ) *EnsureCollectionIndexResponse`

NewEnsureCollectionIndexResponse instantiates a new EnsureCollectionIndexResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnsureCollectionIndexResponseWithDefaults

`func NewEnsureCollectionIndexResponseWithDefaults() *EnsureCollectionIndexResponse`

NewEnsureCollectionIndexResponseWithDefaults instantiates a new EnsureCollectionIndexResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *EnsureCollectionIndexResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EnsureCollectionIndexResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EnsureCollectionIndexResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EnsureCollectionIndexResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSamplesCount

`func (o *EnsureCollectionIndexResponse) GetSamplesCount() float64`

GetSamplesCount returns the SamplesCount field if non-nil, zero value otherwise.

### GetSamplesCountOk

`func (o *EnsureCollectionIndexResponse) GetSamplesCountOk() (*float64, bool)`

GetSamplesCountOk returns a tuple with the SamplesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplesCount

`func (o *EnsureCollectionIndexResponse) SetSamplesCount(v float64)`

SetSamplesCount sets SamplesCount field to given value.


### GetEsIndexedCount

`func (o *EnsureCollectionIndexResponse) GetEsIndexedCount() float64`

GetEsIndexedCount returns the EsIndexedCount field if non-nil, zero value otherwise.

### GetEsIndexedCountOk

`func (o *EnsureCollectionIndexResponse) GetEsIndexedCountOk() (*float64, bool)`

GetEsIndexedCountOk returns a tuple with the EsIndexedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsIndexedCount

`func (o *EnsureCollectionIndexResponse) SetEsIndexedCount(v float64)`

SetEsIndexedCount sets EsIndexedCount field to given value.


### GetStatus

`func (o *EnsureCollectionIndexResponse) GetStatus() CollectionIndexStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnsureCollectionIndexResponse) GetStatusOk() (*CollectionIndexStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnsureCollectionIndexResponse) SetStatus(v CollectionIndexStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


