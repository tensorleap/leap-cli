# DatasetSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preprocess** | [**DatasetPreprocess**](DatasetPreprocess.md) |  | 
**Inputs** | [**[]DatasetInputInstance**](DatasetInputInstance.md) |  | 
**Metadata** | [**[]DatasetMetadataInstance**](DatasetMetadataInstance.md) |  | 
**Outputs** | [**[]DatasetOutputInstance**](DatasetOutputInstance.md) |  | 
**Visualizers** | [**[]VisualizerInstance**](VisualizerInstance.md) |  | 
**PredictionTypes** | [**[]PredictionTypeInstance**](PredictionTypeInstance.md) |  | 
**CustomLosses** | [**[]CustomLossInstance**](CustomLossInstance.md) |  | 
**Metrics** | [**[]MetricInstance**](MetricInstance.md) |  | 

## Methods

### NewDatasetSetup

`func NewDatasetSetup(preprocess DatasetPreprocess, inputs []DatasetInputInstance, metadata []DatasetMetadataInstance, outputs []DatasetOutputInstance, visualizers []VisualizerInstance, predictionTypes []PredictionTypeInstance, customLosses []CustomLossInstance, metrics []MetricInstance, ) *DatasetSetup`

NewDatasetSetup instantiates a new DatasetSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetSetupWithDefaults

`func NewDatasetSetupWithDefaults() *DatasetSetup`

NewDatasetSetupWithDefaults instantiates a new DatasetSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreprocess

`func (o *DatasetSetup) GetPreprocess() DatasetPreprocess`

GetPreprocess returns the Preprocess field if non-nil, zero value otherwise.

### GetPreprocessOk

`func (o *DatasetSetup) GetPreprocessOk() (*DatasetPreprocess, bool)`

GetPreprocessOk returns a tuple with the Preprocess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprocess

`func (o *DatasetSetup) SetPreprocess(v DatasetPreprocess)`

SetPreprocess sets Preprocess field to given value.


### GetInputs

`func (o *DatasetSetup) GetInputs() []DatasetInputInstance`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *DatasetSetup) GetInputsOk() (*[]DatasetInputInstance, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *DatasetSetup) SetInputs(v []DatasetInputInstance)`

SetInputs sets Inputs field to given value.


### GetMetadata

`func (o *DatasetSetup) GetMetadata() []DatasetMetadataInstance`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DatasetSetup) GetMetadataOk() (*[]DatasetMetadataInstance, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DatasetSetup) SetMetadata(v []DatasetMetadataInstance)`

SetMetadata sets Metadata field to given value.


### GetOutputs

`func (o *DatasetSetup) GetOutputs() []DatasetOutputInstance`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *DatasetSetup) GetOutputsOk() (*[]DatasetOutputInstance, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *DatasetSetup) SetOutputs(v []DatasetOutputInstance)`

SetOutputs sets Outputs field to given value.


### GetVisualizers

`func (o *DatasetSetup) GetVisualizers() []VisualizerInstance`

GetVisualizers returns the Visualizers field if non-nil, zero value otherwise.

### GetVisualizersOk

`func (o *DatasetSetup) GetVisualizersOk() (*[]VisualizerInstance, bool)`

GetVisualizersOk returns a tuple with the Visualizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizers

`func (o *DatasetSetup) SetVisualizers(v []VisualizerInstance)`

SetVisualizers sets Visualizers field to given value.


### GetPredictionTypes

`func (o *DatasetSetup) GetPredictionTypes() []PredictionTypeInstance`

GetPredictionTypes returns the PredictionTypes field if non-nil, zero value otherwise.

### GetPredictionTypesOk

`func (o *DatasetSetup) GetPredictionTypesOk() (*[]PredictionTypeInstance, bool)`

GetPredictionTypesOk returns a tuple with the PredictionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionTypes

`func (o *DatasetSetup) SetPredictionTypes(v []PredictionTypeInstance)`

SetPredictionTypes sets PredictionTypes field to given value.


### GetCustomLosses

`func (o *DatasetSetup) GetCustomLosses() []CustomLossInstance`

GetCustomLosses returns the CustomLosses field if non-nil, zero value otherwise.

### GetCustomLossesOk

`func (o *DatasetSetup) GetCustomLossesOk() (*[]CustomLossInstance, bool)`

GetCustomLossesOk returns a tuple with the CustomLosses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLosses

`func (o *DatasetSetup) SetCustomLosses(v []CustomLossInstance)`

SetCustomLosses sets CustomLosses field to given value.


### GetMetrics

`func (o *DatasetSetup) GetMetrics() []MetricInstance`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *DatasetSetup) GetMetricsOk() (*[]MetricInstance, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *DatasetSetup) SetMetrics(v []MetricInstance)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


