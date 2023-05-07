# TrainFromInitialWeightsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**FromSessionRunId** | **string** |  | 
**FromEpoch** | **float64** |  | 
**ModelName** | **string** |  | 
**TrainingParams** | [**TrainingParams**](TrainingParams.md) |  | 
**ShouldRunPopulationExploration** | **bool** |  | 

## Methods

### NewTrainFromInitialWeightsParams

`func NewTrainFromInitialWeightsParams(versionId string, fromSessionRunId string, fromEpoch float64, modelName string, trainingParams TrainingParams, shouldRunPopulationExploration bool, ) *TrainFromInitialWeightsParams`

NewTrainFromInitialWeightsParams instantiates a new TrainFromInitialWeightsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainFromInitialWeightsParamsWithDefaults

`func NewTrainFromInitialWeightsParamsWithDefaults() *TrainFromInitialWeightsParams`

NewTrainFromInitialWeightsParamsWithDefaults instantiates a new TrainFromInitialWeightsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *TrainFromInitialWeightsParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *TrainFromInitialWeightsParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *TrainFromInitialWeightsParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetFromSessionRunId

`func (o *TrainFromInitialWeightsParams) GetFromSessionRunId() string`

GetFromSessionRunId returns the FromSessionRunId field if non-nil, zero value otherwise.

### GetFromSessionRunIdOk

`func (o *TrainFromInitialWeightsParams) GetFromSessionRunIdOk() (*string, bool)`

GetFromSessionRunIdOk returns a tuple with the FromSessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromSessionRunId

`func (o *TrainFromInitialWeightsParams) SetFromSessionRunId(v string)`

SetFromSessionRunId sets FromSessionRunId field to given value.


### GetFromEpoch

`func (o *TrainFromInitialWeightsParams) GetFromEpoch() float64`

GetFromEpoch returns the FromEpoch field if non-nil, zero value otherwise.

### GetFromEpochOk

`func (o *TrainFromInitialWeightsParams) GetFromEpochOk() (*float64, bool)`

GetFromEpochOk returns a tuple with the FromEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEpoch

`func (o *TrainFromInitialWeightsParams) SetFromEpoch(v float64)`

SetFromEpoch sets FromEpoch field to given value.


### GetModelName

`func (o *TrainFromInitialWeightsParams) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *TrainFromInitialWeightsParams) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *TrainFromInitialWeightsParams) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetTrainingParams

`func (o *TrainFromInitialWeightsParams) GetTrainingParams() TrainingParams`

GetTrainingParams returns the TrainingParams field if non-nil, zero value otherwise.

### GetTrainingParamsOk

`func (o *TrainFromInitialWeightsParams) GetTrainingParamsOk() (*TrainingParams, bool)`

GetTrainingParamsOk returns a tuple with the TrainingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingParams

`func (o *TrainFromInitialWeightsParams) SetTrainingParams(v TrainingParams)`

SetTrainingParams sets TrainingParams field to given value.


### GetShouldRunPopulationExploration

`func (o *TrainFromInitialWeightsParams) GetShouldRunPopulationExploration() bool`

GetShouldRunPopulationExploration returns the ShouldRunPopulationExploration field if non-nil, zero value otherwise.

### GetShouldRunPopulationExplorationOk

`func (o *TrainFromInitialWeightsParams) GetShouldRunPopulationExplorationOk() (*bool, bool)`

GetShouldRunPopulationExplorationOk returns a tuple with the ShouldRunPopulationExploration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldRunPopulationExploration

`func (o *TrainFromInitialWeightsParams) SetShouldRunPopulationExploration(v bool)`

SetShouldRunPopulationExploration sets ShouldRunPopulationExploration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


