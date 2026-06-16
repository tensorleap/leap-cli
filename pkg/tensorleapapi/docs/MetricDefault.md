# MetricDefault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Direction** | [**MetricDirection**](MetricDirection.md) |  | 
**Active** | **bool** |  | 

## Methods

### NewMetricDefault

`func NewMetricDefault(name string, direction MetricDirection, active bool, ) *MetricDefault`

NewMetricDefault instantiates a new MetricDefault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDefaultWithDefaults

`func NewMetricDefaultWithDefaults() *MetricDefault`

NewMetricDefaultWithDefaults instantiates a new MetricDefault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetricDefault) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricDefault) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricDefault) SetName(v string)`

SetName sets Name field to given value.


### GetDirection

`func (o *MetricDefault) GetDirection() MetricDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MetricDefault) GetDirectionOk() (*MetricDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MetricDefault) SetDirection(v MetricDirection)`

SetDirection sets Direction field to given value.


### GetActive

`func (o *MetricDefault) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MetricDefault) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MetricDefault) SetActive(v bool)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


