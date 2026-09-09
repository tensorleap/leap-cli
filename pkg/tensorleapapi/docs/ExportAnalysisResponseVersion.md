# ExportAnalysisResponseVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**SerialNumber** | Pointer to **float64** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewExportAnalysisResponseVersion

`func NewExportAnalysisResponseVersion(createdAt time.Time, name string, ) *ExportAnalysisResponseVersion`

NewExportAnalysisResponseVersion instantiates a new ExportAnalysisResponseVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportAnalysisResponseVersionWithDefaults

`func NewExportAnalysisResponseVersionWithDefaults() *ExportAnalysisResponseVersion`

NewExportAnalysisResponseVersionWithDefaults instantiates a new ExportAnalysisResponseVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ExportAnalysisResponseVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExportAnalysisResponseVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExportAnalysisResponseVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSerialNumber

`func (o *ExportAnalysisResponseVersion) GetSerialNumber() float64`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ExportAnalysisResponseVersion) GetSerialNumberOk() (*float64, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ExportAnalysisResponseVersion) SetSerialNumber(v float64)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ExportAnalysisResponseVersion) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetName

`func (o *ExportAnalysisResponseVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExportAnalysisResponseVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExportAnalysisResponseVersion) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


