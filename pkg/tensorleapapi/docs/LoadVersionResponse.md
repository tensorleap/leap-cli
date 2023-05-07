# LoadVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | [**Project**](Project.md) |  | 
**Version** | [**Version**](Version.md) |  | 

## Methods

### NewLoadVersionResponse

`func NewLoadVersionResponse(project Project, version Version, ) *LoadVersionResponse`

NewLoadVersionResponse instantiates a new LoadVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadVersionResponseWithDefaults

`func NewLoadVersionResponseWithDefaults() *LoadVersionResponse`

NewLoadVersionResponseWithDefaults instantiates a new LoadVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *LoadVersionResponse) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *LoadVersionResponse) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *LoadVersionResponse) SetProject(v Project)`

SetProject sets Project field to given value.


### GetVersion

`func (o *LoadVersionResponse) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LoadVersionResponse) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LoadVersionResponse) SetVersion(v Version)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


