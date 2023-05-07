# ModelGraph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Nodes** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 

## Methods

### NewModelGraph

`func NewModelGraph(id string, nodes map[string]interface{}, ) *ModelGraph`

NewModelGraph instantiates a new ModelGraph object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelGraphWithDefaults

`func NewModelGraphWithDefaults() *ModelGraph`

NewModelGraphWithDefaults instantiates a new ModelGraph object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelGraph) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelGraph) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelGraph) SetId(v string)`

SetId sets Id field to given value.


### GetNodes

`func (o *ModelGraph) GetNodes() map[string]interface{}`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ModelGraph) GetNodesOk() (*map[string]interface{}, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ModelGraph) SetNodes(v map[string]interface{})`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


