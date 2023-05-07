# HealthCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllModules** | [**[]ModuleStatus**](ModuleStatus.md) |  | 

## Methods

### NewHealthCheckResponse

`func NewHealthCheckResponse(allModules []ModuleStatus, ) *HealthCheckResponse`

NewHealthCheckResponse instantiates a new HealthCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckResponseWithDefaults

`func NewHealthCheckResponseWithDefaults() *HealthCheckResponse`

NewHealthCheckResponseWithDefaults instantiates a new HealthCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllModules

`func (o *HealthCheckResponse) GetAllModules() []ModuleStatus`

GetAllModules returns the AllModules field if non-nil, zero value otherwise.

### GetAllModulesOk

`func (o *HealthCheckResponse) GetAllModulesOk() (*[]ModuleStatus, bool)`

GetAllModulesOk returns a tuple with the AllModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllModules

`func (o *HealthCheckResponse) SetAllModules(v []ModuleStatus)`

SetAllModules sets AllModules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


