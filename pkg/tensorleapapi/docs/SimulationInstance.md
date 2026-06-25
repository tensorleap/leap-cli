# SimulationInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ParamNames** | **[]string** |  | 
**SimConfig** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 

## Methods

### NewSimulationInstance

`func NewSimulationInstance(name string, paramNames []string, simConfig map[string]interface{}, ) *SimulationInstance`

NewSimulationInstance instantiates a new SimulationInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulationInstanceWithDefaults

`func NewSimulationInstanceWithDefaults() *SimulationInstance`

NewSimulationInstanceWithDefaults instantiates a new SimulationInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SimulationInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimulationInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimulationInstance) SetName(v string)`

SetName sets Name field to given value.


### GetParamNames

`func (o *SimulationInstance) GetParamNames() []string`

GetParamNames returns the ParamNames field if non-nil, zero value otherwise.

### GetParamNamesOk

`func (o *SimulationInstance) GetParamNamesOk() (*[]string, bool)`

GetParamNamesOk returns a tuple with the ParamNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamNames

`func (o *SimulationInstance) SetParamNames(v []string)`

SetParamNames sets ParamNames field to given value.


### GetSimConfig

`func (o *SimulationInstance) GetSimConfig() map[string]interface{}`

GetSimConfig returns the SimConfig field if non-nil, zero value otherwise.

### GetSimConfigOk

`func (o *SimulationInstance) GetSimConfigOk() (*map[string]interface{}, bool)`

GetSimConfigOk returns a tuple with the SimConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimConfig

`func (o *SimulationInstance) SetSimConfig(v map[string]interface{})`

SetSimConfig sets SimConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


