# MetricInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ArgNames** | **[]string** |  | 

## Methods

### NewMetricInstance

`func NewMetricInstance(name string, argNames []string, ) *MetricInstance`

NewMetricInstance instantiates a new MetricInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricInstanceWithDefaults

`func NewMetricInstanceWithDefaults() *MetricInstance`

NewMetricInstanceWithDefaults instantiates a new MetricInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetricInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricInstance) SetName(v string)`

SetName sets Name field to given value.


### GetArgNames

`func (o *MetricInstance) GetArgNames() []string`

GetArgNames returns the ArgNames field if non-nil, zero value otherwise.

### GetArgNamesOk

`func (o *MetricInstance) GetArgNamesOk() (*[]string, bool)`

GetArgNamesOk returns a tuple with the ArgNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgNames

`func (o *MetricInstance) SetArgNames(v []string)`

SetArgNames sets ArgNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


