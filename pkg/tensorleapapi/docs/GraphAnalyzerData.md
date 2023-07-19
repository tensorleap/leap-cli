# GraphAnalyzerData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 
**Hash** | Pointer to **string** |  | [optional] 
**GeneralError** | Pointer to **string** |  | [optional] 

## Methods

### NewGraphAnalyzerData

`func NewGraphAnalyzerData(nodes map[string]interface{}, ) *GraphAnalyzerData`

NewGraphAnalyzerData instantiates a new GraphAnalyzerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphAnalyzerDataWithDefaults

`func NewGraphAnalyzerDataWithDefaults() *GraphAnalyzerData`

NewGraphAnalyzerDataWithDefaults instantiates a new GraphAnalyzerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *GraphAnalyzerData) GetNodes() map[string]interface{}`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GraphAnalyzerData) GetNodesOk() (*map[string]interface{}, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GraphAnalyzerData) SetNodes(v map[string]interface{})`

SetNodes sets Nodes field to given value.


### GetHash

`func (o *GraphAnalyzerData) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GraphAnalyzerData) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GraphAnalyzerData) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *GraphAnalyzerData) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetGeneralError

`func (o *GraphAnalyzerData) GetGeneralError() string`

GetGeneralError returns the GeneralError field if non-nil, zero value otherwise.

### GetGeneralErrorOk

`func (o *GraphAnalyzerData) GetGeneralErrorOk() (*string, bool)`

GetGeneralErrorOk returns a tuple with the GeneralError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralError

`func (o *GraphAnalyzerData) SetGeneralError(v string)`

SetGeneralError sets GeneralError field to given value.

### HasGeneralError

`func (o *GraphAnalyzerData) HasGeneralError() bool`

HasGeneralError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


