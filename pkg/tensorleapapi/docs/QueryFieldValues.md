# QueryFieldValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **float64** |  | [optional] 
**Size** | Pointer to **float64** |  | [optional] 
**Type** | **string** |  | 
**Query** | Pointer to [**NumberOrString**](NumberOrString.md) |  | [optional] 
**Field** | **string** |  | 

## Methods

### NewQueryFieldValues

`func NewQueryFieldValues(type_ string, field string, ) *QueryFieldValues`

NewQueryFieldValues instantiates a new QueryFieldValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryFieldValuesWithDefaults

`func NewQueryFieldValuesWithDefaults() *QueryFieldValues`

NewQueryFieldValuesWithDefaults instantiates a new QueryFieldValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *QueryFieldValues) GetPage() float64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *QueryFieldValues) GetPageOk() (*float64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *QueryFieldValues) SetPage(v float64)`

SetPage sets Page field to given value.

### HasPage

`func (o *QueryFieldValues) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *QueryFieldValues) GetSize() float64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *QueryFieldValues) GetSizeOk() (*float64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *QueryFieldValues) SetSize(v float64)`

SetSize sets Size field to given value.

### HasSize

`func (o *QueryFieldValues) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *QueryFieldValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryFieldValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryFieldValues) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *QueryFieldValues) GetQuery() NumberOrString`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryFieldValues) GetQueryOk() (*NumberOrString, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryFieldValues) SetQuery(v NumberOrString)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *QueryFieldValues) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetField

`func (o *QueryFieldValues) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *QueryFieldValues) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *QueryFieldValues) SetField(v string)`

SetField sets Field field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


