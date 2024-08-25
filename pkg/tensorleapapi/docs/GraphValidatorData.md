# GraphValidatorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inputs** | [**[]ValidatedNode**](ValidatedNode.md) |  | 
**Metadata** | [**[]ValidatedNode**](ValidatedNode.md) |  | 
**GroundTruths** | [**[]ValidatedNode**](ValidatedNode.md) |  | 
**PredictionTypes** | [**[]ValidatedNode**](ValidatedNode.md) |  | 
**Visualizers** | [**[]ValidatedNode**](ValidatedNode.md) |  | 
**Losses** | [**[]ValidatedLossNode**](ValidatedLossNode.md) |  | 
**Metrics** | [**[]ValidatedNode**](ValidatedNode.md) |  | 
**CustomLayers** | [**[]ValidatedNode**](ValidatedNode.md) |  | 
**GeneralError** | Pointer to **string** |  | [optional] 

## Methods

### NewGraphValidatorData

`func NewGraphValidatorData(inputs []ValidatedNode, metadata []ValidatedNode, groundTruths []ValidatedNode, predictionTypes []ValidatedNode, visualizers []ValidatedNode, losses []ValidatedLossNode, metrics []ValidatedNode, customLayers []ValidatedNode, ) *GraphValidatorData`

NewGraphValidatorData instantiates a new GraphValidatorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphValidatorDataWithDefaults

`func NewGraphValidatorDataWithDefaults() *GraphValidatorData`

NewGraphValidatorDataWithDefaults instantiates a new GraphValidatorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputs

`func (o *GraphValidatorData) GetInputs() []ValidatedNode`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *GraphValidatorData) GetInputsOk() (*[]ValidatedNode, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *GraphValidatorData) SetInputs(v []ValidatedNode)`

SetInputs sets Inputs field to given value.


### GetMetadata

`func (o *GraphValidatorData) GetMetadata() []ValidatedNode`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GraphValidatorData) GetMetadataOk() (*[]ValidatedNode, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GraphValidatorData) SetMetadata(v []ValidatedNode)`

SetMetadata sets Metadata field to given value.


### GetGroundTruths

`func (o *GraphValidatorData) GetGroundTruths() []ValidatedNode`

GetGroundTruths returns the GroundTruths field if non-nil, zero value otherwise.

### GetGroundTruthsOk

`func (o *GraphValidatorData) GetGroundTruthsOk() (*[]ValidatedNode, bool)`

GetGroundTruthsOk returns a tuple with the GroundTruths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroundTruths

`func (o *GraphValidatorData) SetGroundTruths(v []ValidatedNode)`

SetGroundTruths sets GroundTruths field to given value.


### GetPredictionTypes

`func (o *GraphValidatorData) GetPredictionTypes() []ValidatedNode`

GetPredictionTypes returns the PredictionTypes field if non-nil, zero value otherwise.

### GetPredictionTypesOk

`func (o *GraphValidatorData) GetPredictionTypesOk() (*[]ValidatedNode, bool)`

GetPredictionTypesOk returns a tuple with the PredictionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionTypes

`func (o *GraphValidatorData) SetPredictionTypes(v []ValidatedNode)`

SetPredictionTypes sets PredictionTypes field to given value.


### GetVisualizers

`func (o *GraphValidatorData) GetVisualizers() []ValidatedNode`

GetVisualizers returns the Visualizers field if non-nil, zero value otherwise.

### GetVisualizersOk

`func (o *GraphValidatorData) GetVisualizersOk() (*[]ValidatedNode, bool)`

GetVisualizersOk returns a tuple with the Visualizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizers

`func (o *GraphValidatorData) SetVisualizers(v []ValidatedNode)`

SetVisualizers sets Visualizers field to given value.


### GetLosses

`func (o *GraphValidatorData) GetLosses() []ValidatedLossNode`

GetLosses returns the Losses field if non-nil, zero value otherwise.

### GetLossesOk

`func (o *GraphValidatorData) GetLossesOk() (*[]ValidatedLossNode, bool)`

GetLossesOk returns a tuple with the Losses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLosses

`func (o *GraphValidatorData) SetLosses(v []ValidatedLossNode)`

SetLosses sets Losses field to given value.


### GetMetrics

`func (o *GraphValidatorData) GetMetrics() []ValidatedNode`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GraphValidatorData) GetMetricsOk() (*[]ValidatedNode, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GraphValidatorData) SetMetrics(v []ValidatedNode)`

SetMetrics sets Metrics field to given value.


### GetCustomLayers

`func (o *GraphValidatorData) GetCustomLayers() []ValidatedNode`

GetCustomLayers returns the CustomLayers field if non-nil, zero value otherwise.

### GetCustomLayersOk

`func (o *GraphValidatorData) GetCustomLayersOk() (*[]ValidatedNode, bool)`

GetCustomLayersOk returns a tuple with the CustomLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLayers

`func (o *GraphValidatorData) SetCustomLayers(v []ValidatedNode)`

SetCustomLayers sets CustomLayers field to given value.


### GetGeneralError

`func (o *GraphValidatorData) GetGeneralError() string`

GetGeneralError returns the GeneralError field if non-nil, zero value otherwise.

### GetGeneralErrorOk

`func (o *GraphValidatorData) GetGeneralErrorOk() (*string, bool)`

GetGeneralErrorOk returns a tuple with the GeneralError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralError

`func (o *GraphValidatorData) SetGeneralError(v string)`

SetGeneralError sets GeneralError field to given value.

### HasGeneralError

`func (o *GraphValidatorData) HasGeneralError() bool`

HasGeneralError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


