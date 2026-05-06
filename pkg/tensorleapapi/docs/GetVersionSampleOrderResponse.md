# GetVersionSampleOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **float64** |  | [optional] 
**Rows** | [**[]CollectionSampleRow**](CollectionSampleRow.md) |  | 
**Samples** | [**[]SampleIdentity**](SampleIdentity.md) |  | 

## Methods

### NewGetVersionSampleOrderResponse

`func NewGetVersionSampleOrderResponse(rows []CollectionSampleRow, samples []SampleIdentity, ) *GetVersionSampleOrderResponse`

NewGetVersionSampleOrderResponse instantiates a new GetVersionSampleOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVersionSampleOrderResponseWithDefaults

`func NewGetVersionSampleOrderResponseWithDefaults() *GetVersionSampleOrderResponse`

NewGetVersionSampleOrderResponseWithDefaults instantiates a new GetVersionSampleOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *GetVersionSampleOrderResponse) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetVersionSampleOrderResponse) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetVersionSampleOrderResponse) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GetVersionSampleOrderResponse) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetTotal

`func (o *GetVersionSampleOrderResponse) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetVersionSampleOrderResponse) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetVersionSampleOrderResponse) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetVersionSampleOrderResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetRows

`func (o *GetVersionSampleOrderResponse) GetRows() []CollectionSampleRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetVersionSampleOrderResponse) GetRowsOk() (*[]CollectionSampleRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetVersionSampleOrderResponse) SetRows(v []CollectionSampleRow)`

SetRows sets Rows field to given value.


### GetSamples

`func (o *GetVersionSampleOrderResponse) GetSamples() []SampleIdentity`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *GetVersionSampleOrderResponse) GetSamplesOk() (*[]SampleIdentity, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *GetVersionSampleOrderResponse) SetSamples(v []SampleIdentity)`

SetSamples sets Samples field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


