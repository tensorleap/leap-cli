# ManualInsightJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** |  | 
**Filters** | [**[]ESFilter**](ESFilter.md) |  | 
**VersionId** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewManualInsightJobParams

`func NewManualInsightJobParams(digest string, filters []ESFilter, versionId string, type_ string, ) *ManualInsightJobParams`

NewManualInsightJobParams instantiates a new ManualInsightJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualInsightJobParamsWithDefaults

`func NewManualInsightJobParamsWithDefaults() *ManualInsightJobParams`

NewManualInsightJobParamsWithDefaults instantiates a new ManualInsightJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *ManualInsightJobParams) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ManualInsightJobParams) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ManualInsightJobParams) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetFilters

`func (o *ManualInsightJobParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ManualInsightJobParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ManualInsightJobParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.


### GetVersionId

`func (o *ManualInsightJobParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ManualInsightJobParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ManualInsightJobParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetType

`func (o *ManualInsightJobParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManualInsightJobParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManualInsightJobParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


