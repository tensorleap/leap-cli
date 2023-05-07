# VisualizationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | [**[]VizType**](VizType.md) |  | 
**Title** | **string** |  | 
**Epoch** | **float64** |  | 

## Methods

### NewVisualizationData

`func NewVisualizationData(payload []VizType, title string, epoch float64, ) *VisualizationData`

NewVisualizationData instantiates a new VisualizationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualizationDataWithDefaults

`func NewVisualizationDataWithDefaults() *VisualizationData`

NewVisualizationDataWithDefaults instantiates a new VisualizationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *VisualizationData) GetPayload() []VizType`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *VisualizationData) GetPayloadOk() (*[]VizType, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *VisualizationData) SetPayload(v []VizType)`

SetPayload sets Payload field to given value.


### GetTitle

`func (o *VisualizationData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VisualizationData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VisualizationData) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetEpoch

`func (o *VisualizationData) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *VisualizationData) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *VisualizationData) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


