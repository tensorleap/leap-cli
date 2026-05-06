# GetCollectionSampleOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **float64** |  | [optional] 
**Rows** | [**[]CollectionSampleRow**](CollectionSampleRow.md) | Same rows as &#x60;samples&#x60;, but with metadata/metrics flattened in. The table-pagination hook uses these directly as DataGrid rows. | 
**Samples** | [**[]SampleIdentity**](SampleIdentity.md) | Identities-only — kept for callers that just need state/index. | 
**Status** | [**CollectionIndexStatus**](CollectionIndexStatus.md) |  | 

## Methods

### NewGetCollectionSampleOrderResponse

`func NewGetCollectionSampleOrderResponse(rows []CollectionSampleRow, samples []SampleIdentity, status CollectionIndexStatus, ) *GetCollectionSampleOrderResponse`

NewGetCollectionSampleOrderResponse instantiates a new GetCollectionSampleOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionSampleOrderResponseWithDefaults

`func NewGetCollectionSampleOrderResponseWithDefaults() *GetCollectionSampleOrderResponse`

NewGetCollectionSampleOrderResponseWithDefaults instantiates a new GetCollectionSampleOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *GetCollectionSampleOrderResponse) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetCollectionSampleOrderResponse) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetCollectionSampleOrderResponse) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GetCollectionSampleOrderResponse) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetTotal

`func (o *GetCollectionSampleOrderResponse) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetCollectionSampleOrderResponse) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetCollectionSampleOrderResponse) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetCollectionSampleOrderResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetRows

`func (o *GetCollectionSampleOrderResponse) GetRows() []CollectionSampleRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetCollectionSampleOrderResponse) GetRowsOk() (*[]CollectionSampleRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetCollectionSampleOrderResponse) SetRows(v []CollectionSampleRow)`

SetRows sets Rows field to given value.


### GetSamples

`func (o *GetCollectionSampleOrderResponse) GetSamples() []SampleIdentity`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *GetCollectionSampleOrderResponse) GetSamplesOk() (*[]SampleIdentity, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *GetCollectionSampleOrderResponse) SetSamples(v []SampleIdentity)`

SetSamples sets Samples field to given value.


### GetStatus

`func (o *GetCollectionSampleOrderResponse) GetStatus() CollectionIndexStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCollectionSampleOrderResponse) GetStatusOk() (*CollectionIndexStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCollectionSampleOrderResponse) SetStatus(v CollectionIndexStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


