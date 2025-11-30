# LicenseMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**LicenseType**](LicenseType.md) |  | 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**MaxUsers** | Pointer to **float64** |  | [optional] 

## Methods

### NewLicenseMetadata

`func NewLicenseMetadata(type_ LicenseType, ) *LicenseMetadata`

NewLicenseMetadata instantiates a new LicenseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseMetadataWithDefaults

`func NewLicenseMetadataWithDefaults() *LicenseMetadata`

NewLicenseMetadataWithDefaults instantiates a new LicenseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LicenseMetadata) GetType() LicenseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LicenseMetadata) GetTypeOk() (*LicenseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LicenseMetadata) SetType(v LicenseType)`

SetType sets Type field to given value.


### GetExpirationDate

`func (o *LicenseMetadata) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *LicenseMetadata) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *LicenseMetadata) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *LicenseMetadata) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetMaxUsers

`func (o *LicenseMetadata) GetMaxUsers() float64`

GetMaxUsers returns the MaxUsers field if non-nil, zero value otherwise.

### GetMaxUsersOk

`func (o *LicenseMetadata) GetMaxUsersOk() (*float64, bool)`

GetMaxUsersOk returns a tuple with the MaxUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsers

`func (o *LicenseMetadata) SetMaxUsers(v float64)`

SetMaxUsers sets MaxUsers field to given value.

### HasMaxUsers

`func (o *LicenseMetadata) HasMaxUsers() bool`

HasMaxUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


