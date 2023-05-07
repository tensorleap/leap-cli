# ExportModelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ExportModelTypeEnum**](ExportModelTypeEnum.md) |  | 
**Title** | **string** |  | 
**Epoch** | **float64** |  | 

## Methods

### NewExportModelParams

`func NewExportModelParams(type_ ExportModelTypeEnum, title string, epoch float64, ) *ExportModelParams`

NewExportModelParams instantiates a new ExportModelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportModelParamsWithDefaults

`func NewExportModelParamsWithDefaults() *ExportModelParams`

NewExportModelParamsWithDefaults instantiates a new ExportModelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExportModelParams) GetType() ExportModelTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportModelParams) GetTypeOk() (*ExportModelTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportModelParams) SetType(v ExportModelTypeEnum)`

SetType sets Type field to given value.


### GetTitle

`func (o *ExportModelParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExportModelParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExportModelParams) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetEpoch

`func (o *ExportModelParams) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *ExportModelParams) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *ExportModelParams) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


