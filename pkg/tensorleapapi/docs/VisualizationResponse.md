# VisualizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**VisualizationData**](VisualizationData.md) |  | 
**Info** | [**VizInfoType**](VizInfoType.md) |  | 
**Source** | **string** |  | 

## Methods

### NewVisualizationResponse

`func NewVisualizationResponse(data VisualizationData, info VizInfoType, source string, ) *VisualizationResponse`

NewVisualizationResponse instantiates a new VisualizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualizationResponseWithDefaults

`func NewVisualizationResponseWithDefaults() *VisualizationResponse`

NewVisualizationResponseWithDefaults instantiates a new VisualizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VisualizationResponse) GetData() VisualizationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VisualizationResponse) GetDataOk() (*VisualizationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VisualizationResponse) SetData(v VisualizationData)`

SetData sets Data field to given value.


### GetInfo

`func (o *VisualizationResponse) GetInfo() VizInfoType`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *VisualizationResponse) GetInfoOk() (*VizInfoType, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *VisualizationResponse) SetInfo(v VizInfoType)`

SetInfo sets Info field to given value.


### GetSource

`func (o *VisualizationResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VisualizationResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VisualizationResponse) SetSource(v string)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


