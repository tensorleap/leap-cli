# SampleCsvColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeaderName** | **string** |  | 
**Field** | **string** | The row key, as in &#x60;CollectionSampleRow&#x60; (&#x60;state&#x60;, &#x60;metadata.x&#x60;, …). | 

## Methods

### NewSampleCsvColumn

`func NewSampleCsvColumn(headerName string, field string, ) *SampleCsvColumn`

NewSampleCsvColumn instantiates a new SampleCsvColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleCsvColumnWithDefaults

`func NewSampleCsvColumnWithDefaults() *SampleCsvColumn`

NewSampleCsvColumnWithDefaults instantiates a new SampleCsvColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaderName

`func (o *SampleCsvColumn) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *SampleCsvColumn) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *SampleCsvColumn) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.


### GetField

`func (o *SampleCsvColumn) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *SampleCsvColumn) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *SampleCsvColumn) SetField(v string)`

SetField sets Field field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


