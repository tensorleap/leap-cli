# GetCollectionSamplesMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilteredSamplesCount** | **float64** | Number of indexed docs that match the *current* filter spec across ALL active filters (i.e. the size of the filtered subset). When no filters are active, equals &#x60;esIndexedCount&#x60;. | 
**SamplesCount** | **float64** |  | 
**EsIndexedCount** | **float64** |  | 
**IndexStatus** | [**CollectionIndexStatus**](CollectionIndexStatus.md) |  | 
**Fields** | [**[]MetadataField**](MetadataField.md) |  | 

## Methods

### NewGetCollectionSamplesMetadataResponse

`func NewGetCollectionSamplesMetadataResponse(filteredSamplesCount float64, samplesCount float64, esIndexedCount float64, indexStatus CollectionIndexStatus, fields []MetadataField, ) *GetCollectionSamplesMetadataResponse`

NewGetCollectionSamplesMetadataResponse instantiates a new GetCollectionSamplesMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionSamplesMetadataResponseWithDefaults

`func NewGetCollectionSamplesMetadataResponseWithDefaults() *GetCollectionSamplesMetadataResponse`

NewGetCollectionSamplesMetadataResponseWithDefaults instantiates a new GetCollectionSamplesMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilteredSamplesCount

`func (o *GetCollectionSamplesMetadataResponse) GetFilteredSamplesCount() float64`

GetFilteredSamplesCount returns the FilteredSamplesCount field if non-nil, zero value otherwise.

### GetFilteredSamplesCountOk

`func (o *GetCollectionSamplesMetadataResponse) GetFilteredSamplesCountOk() (*float64, bool)`

GetFilteredSamplesCountOk returns a tuple with the FilteredSamplesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredSamplesCount

`func (o *GetCollectionSamplesMetadataResponse) SetFilteredSamplesCount(v float64)`

SetFilteredSamplesCount sets FilteredSamplesCount field to given value.


### GetSamplesCount

`func (o *GetCollectionSamplesMetadataResponse) GetSamplesCount() float64`

GetSamplesCount returns the SamplesCount field if non-nil, zero value otherwise.

### GetSamplesCountOk

`func (o *GetCollectionSamplesMetadataResponse) GetSamplesCountOk() (*float64, bool)`

GetSamplesCountOk returns a tuple with the SamplesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplesCount

`func (o *GetCollectionSamplesMetadataResponse) SetSamplesCount(v float64)`

SetSamplesCount sets SamplesCount field to given value.


### GetEsIndexedCount

`func (o *GetCollectionSamplesMetadataResponse) GetEsIndexedCount() float64`

GetEsIndexedCount returns the EsIndexedCount field if non-nil, zero value otherwise.

### GetEsIndexedCountOk

`func (o *GetCollectionSamplesMetadataResponse) GetEsIndexedCountOk() (*float64, bool)`

GetEsIndexedCountOk returns a tuple with the EsIndexedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsIndexedCount

`func (o *GetCollectionSamplesMetadataResponse) SetEsIndexedCount(v float64)`

SetEsIndexedCount sets EsIndexedCount field to given value.


### GetIndexStatus

`func (o *GetCollectionSamplesMetadataResponse) GetIndexStatus() CollectionIndexStatus`

GetIndexStatus returns the IndexStatus field if non-nil, zero value otherwise.

### GetIndexStatusOk

`func (o *GetCollectionSamplesMetadataResponse) GetIndexStatusOk() (*CollectionIndexStatus, bool)`

GetIndexStatusOk returns a tuple with the IndexStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexStatus

`func (o *GetCollectionSamplesMetadataResponse) SetIndexStatus(v CollectionIndexStatus)`

SetIndexStatus sets IndexStatus field to given value.


### GetFields

`func (o *GetCollectionSamplesMetadataResponse) GetFields() []MetadataField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GetCollectionSamplesMetadataResponse) GetFieldsOk() (*[]MetadataField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GetCollectionSamplesMetadataResponse) SetFields(v []MetadataField)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


