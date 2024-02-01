# ESFilterValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gt** | Pointer to **float64** |  | [optional] 
**Lt** | Pointer to **float64** |  | [optional] 
**Eq** | Pointer to [**NumberOrString**](NumberOrString.md) |  | [optional] 
**Lst** | Pointer to [**[]NumberOrString**](NumberOrString.md) |  | [optional] 
**BlobPath** | Pointer to **string** |  | [optional] 

## Methods

### NewESFilterValue

`func NewESFilterValue() *ESFilterValue`

NewESFilterValue instantiates a new ESFilterValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewESFilterValueWithDefaults

`func NewESFilterValueWithDefaults() *ESFilterValue`

NewESFilterValueWithDefaults instantiates a new ESFilterValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGt

`func (o *ESFilterValue) GetGt() float64`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *ESFilterValue) GetGtOk() (*float64, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *ESFilterValue) SetGt(v float64)`

SetGt sets Gt field to given value.

### HasGt

`func (o *ESFilterValue) HasGt() bool`

HasGt returns a boolean if a field has been set.

### GetLt

`func (o *ESFilterValue) GetLt() float64`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *ESFilterValue) GetLtOk() (*float64, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *ESFilterValue) SetLt(v float64)`

SetLt sets Lt field to given value.

### HasLt

`func (o *ESFilterValue) HasLt() bool`

HasLt returns a boolean if a field has been set.

### GetEq

`func (o *ESFilterValue) GetEq() NumberOrString`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *ESFilterValue) GetEqOk() (*NumberOrString, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *ESFilterValue) SetEq(v NumberOrString)`

SetEq sets Eq field to given value.

### HasEq

`func (o *ESFilterValue) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetLst

`func (o *ESFilterValue) GetLst() []NumberOrString`

GetLst returns the Lst field if non-nil, zero value otherwise.

### GetLstOk

`func (o *ESFilterValue) GetLstOk() (*[]NumberOrString, bool)`

GetLstOk returns a tuple with the Lst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLst

`func (o *ESFilterValue) SetLst(v []NumberOrString)`

SetLst sets Lst field to given value.

### HasLst

`func (o *ESFilterValue) HasLst() bool`

HasLst returns a boolean if a field has been set.

### GetBlobPath

`func (o *ESFilterValue) GetBlobPath() string`

GetBlobPath returns the BlobPath field if non-nil, zero value otherwise.

### GetBlobPathOk

`func (o *ESFilterValue) GetBlobPathOk() (*string, bool)`

GetBlobPathOk returns a tuple with the BlobPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobPath

`func (o *ESFilterValue) SetBlobPath(v string)`

SetBlobPath sets BlobPath field to given value.

### HasBlobPath

`func (o *ESFilterValue) HasBlobPath() bool`

HasBlobPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


