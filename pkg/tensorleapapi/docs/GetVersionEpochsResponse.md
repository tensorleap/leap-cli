# GetVersionEpochsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Epochs** | [**[]EpochData**](EpochData.md) |  | 

## Methods

### NewGetVersionEpochsResponse

`func NewGetVersionEpochsResponse(epochs []EpochData, ) *GetVersionEpochsResponse`

NewGetVersionEpochsResponse instantiates a new GetVersionEpochsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVersionEpochsResponseWithDefaults

`func NewGetVersionEpochsResponseWithDefaults() *GetVersionEpochsResponse`

NewGetVersionEpochsResponseWithDefaults instantiates a new GetVersionEpochsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpochs

`func (o *GetVersionEpochsResponse) GetEpochs() []EpochData`

GetEpochs returns the Epochs field if non-nil, zero value otherwise.

### GetEpochsOk

`func (o *GetVersionEpochsResponse) GetEpochsOk() (*[]EpochData, bool)`

GetEpochsOk returns a tuple with the Epochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochs

`func (o *GetVersionEpochsResponse) SetEpochs(v []EpochData)`

SetEpochs sets Epochs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


