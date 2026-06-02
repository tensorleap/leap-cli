# CodeSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preprocess** | [**CodePreprocess**](CodePreprocess.md) |  | 
**Inputs** | [**[]CodeInputInstance**](CodeInputInstance.md) |  | 
**Metadata** | [**[]CodeMetadataInstance**](CodeMetadataInstance.md) |  | 
**Outputs** | [**[]CodeOutputInstance**](CodeOutputInstance.md) |  | 
**Visualizers** | [**[]VisualizerInstance**](VisualizerInstance.md) |  | 
**PredictionTypes** | [**[]PredictionTypeInstance**](PredictionTypeInstance.md) |  | 
**CustomLosses** | [**[]CustomLossInstance**](CustomLossInstance.md) |  | 
**Metrics** | [**[]MetricInstance**](MetricInstance.md) |  | 

## Methods

### NewCodeSetup

`func NewCodeSetup(preprocess CodePreprocess, inputs []CodeInputInstance, metadata []CodeMetadataInstance, outputs []CodeOutputInstance, visualizers []VisualizerInstance, predictionTypes []PredictionTypeInstance, customLosses []CustomLossInstance, metrics []MetricInstance, ) *CodeSetup`

NewCodeSetup instantiates a new CodeSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeSetupWithDefaults

`func NewCodeSetupWithDefaults() *CodeSetup`

NewCodeSetupWithDefaults instantiates a new CodeSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreprocess

`func (o *CodeSetup) GetPreprocess() CodePreprocess`

GetPreprocess returns the Preprocess field if non-nil, zero value otherwise.

### GetPreprocessOk

`func (o *CodeSetup) GetPreprocessOk() (*CodePreprocess, bool)`

GetPreprocessOk returns a tuple with the Preprocess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprocess

`func (o *CodeSetup) SetPreprocess(v CodePreprocess)`

SetPreprocess sets Preprocess field to given value.


### GetInputs

`func (o *CodeSetup) GetInputs() []CodeInputInstance`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *CodeSetup) GetInputsOk() (*[]CodeInputInstance, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *CodeSetup) SetInputs(v []CodeInputInstance)`

SetInputs sets Inputs field to given value.


### GetMetadata

`func (o *CodeSetup) GetMetadata() []CodeMetadataInstance`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CodeSetup) GetMetadataOk() (*[]CodeMetadataInstance, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CodeSetup) SetMetadata(v []CodeMetadataInstance)`

SetMetadata sets Metadata field to given value.


### GetOutputs

`func (o *CodeSetup) GetOutputs() []CodeOutputInstance`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *CodeSetup) GetOutputsOk() (*[]CodeOutputInstance, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *CodeSetup) SetOutputs(v []CodeOutputInstance)`

SetOutputs sets Outputs field to given value.


### GetVisualizers

`func (o *CodeSetup) GetVisualizers() []VisualizerInstance`

GetVisualizers returns the Visualizers field if non-nil, zero value otherwise.

### GetVisualizersOk

`func (o *CodeSetup) GetVisualizersOk() (*[]VisualizerInstance, bool)`

GetVisualizersOk returns a tuple with the Visualizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizers

`func (o *CodeSetup) SetVisualizers(v []VisualizerInstance)`

SetVisualizers sets Visualizers field to given value.


### GetPredictionTypes

`func (o *CodeSetup) GetPredictionTypes() []PredictionTypeInstance`

GetPredictionTypes returns the PredictionTypes field if non-nil, zero value otherwise.

### GetPredictionTypesOk

`func (o *CodeSetup) GetPredictionTypesOk() (*[]PredictionTypeInstance, bool)`

GetPredictionTypesOk returns a tuple with the PredictionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionTypes

`func (o *CodeSetup) SetPredictionTypes(v []PredictionTypeInstance)`

SetPredictionTypes sets PredictionTypes field to given value.


### GetCustomLosses

`func (o *CodeSetup) GetCustomLosses() []CustomLossInstance`

GetCustomLosses returns the CustomLosses field if non-nil, zero value otherwise.

### GetCustomLossesOk

`func (o *CodeSetup) GetCustomLossesOk() (*[]CustomLossInstance, bool)`

GetCustomLossesOk returns a tuple with the CustomLosses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLosses

`func (o *CodeSetup) SetCustomLosses(v []CustomLossInstance)`

SetCustomLosses sets CustomLosses field to given value.


### GetMetrics

`func (o *CodeSetup) GetMetrics() []MetricInstance`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *CodeSetup) GetMetricsOk() (*[]MetricInstance, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *CodeSetup) SetMetrics(v []MetricInstance)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


