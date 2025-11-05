# GetEnvironmentInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientStoragePrefixUrl** | **string** |  | 
**SchemaVersion** | **float64** |  | 
**HubDefaultNamespace** | **string** |  | 
**HubUrl** | **string** |  | 
**DisableDatadogMetrics** | **bool** |  | 
**IsCloud** | **bool** |  | 
**GatewayUrl** | **string** |  | 
**ProxyUrl** | **string** |  | 
**LicenseMetadata** | Pointer to [**LicenseMetadata**](LicenseMetadata.md) |  | [optional] 

## Methods

### NewGetEnvironmentInfoResponse

`func NewGetEnvironmentInfoResponse(clientStoragePrefixUrl string, schemaVersion float64, hubDefaultNamespace string, hubUrl string, disableDatadogMetrics bool, isCloud bool, gatewayUrl string, proxyUrl string, ) *GetEnvironmentInfoResponse`

NewGetEnvironmentInfoResponse instantiates a new GetEnvironmentInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEnvironmentInfoResponseWithDefaults

`func NewGetEnvironmentInfoResponseWithDefaults() *GetEnvironmentInfoResponse`

NewGetEnvironmentInfoResponseWithDefaults instantiates a new GetEnvironmentInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientStoragePrefixUrl

`func (o *GetEnvironmentInfoResponse) GetClientStoragePrefixUrl() string`

GetClientStoragePrefixUrl returns the ClientStoragePrefixUrl field if non-nil, zero value otherwise.

### GetClientStoragePrefixUrlOk

`func (o *GetEnvironmentInfoResponse) GetClientStoragePrefixUrlOk() (*string, bool)`

GetClientStoragePrefixUrlOk returns a tuple with the ClientStoragePrefixUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientStoragePrefixUrl

`func (o *GetEnvironmentInfoResponse) SetClientStoragePrefixUrl(v string)`

SetClientStoragePrefixUrl sets ClientStoragePrefixUrl field to given value.


### GetSchemaVersion

`func (o *GetEnvironmentInfoResponse) GetSchemaVersion() float64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetEnvironmentInfoResponse) GetSchemaVersionOk() (*float64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetEnvironmentInfoResponse) SetSchemaVersion(v float64)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetHubDefaultNamespace

`func (o *GetEnvironmentInfoResponse) GetHubDefaultNamespace() string`

GetHubDefaultNamespace returns the HubDefaultNamespace field if non-nil, zero value otherwise.

### GetHubDefaultNamespaceOk

`func (o *GetEnvironmentInfoResponse) GetHubDefaultNamespaceOk() (*string, bool)`

GetHubDefaultNamespaceOk returns a tuple with the HubDefaultNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubDefaultNamespace

`func (o *GetEnvironmentInfoResponse) SetHubDefaultNamespace(v string)`

SetHubDefaultNamespace sets HubDefaultNamespace field to given value.


### GetHubUrl

`func (o *GetEnvironmentInfoResponse) GetHubUrl() string`

GetHubUrl returns the HubUrl field if non-nil, zero value otherwise.

### GetHubUrlOk

`func (o *GetEnvironmentInfoResponse) GetHubUrlOk() (*string, bool)`

GetHubUrlOk returns a tuple with the HubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubUrl

`func (o *GetEnvironmentInfoResponse) SetHubUrl(v string)`

SetHubUrl sets HubUrl field to given value.


### GetDisableDatadogMetrics

`func (o *GetEnvironmentInfoResponse) GetDisableDatadogMetrics() bool`

GetDisableDatadogMetrics returns the DisableDatadogMetrics field if non-nil, zero value otherwise.

### GetDisableDatadogMetricsOk

`func (o *GetEnvironmentInfoResponse) GetDisableDatadogMetricsOk() (*bool, bool)`

GetDisableDatadogMetricsOk returns a tuple with the DisableDatadogMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDatadogMetrics

`func (o *GetEnvironmentInfoResponse) SetDisableDatadogMetrics(v bool)`

SetDisableDatadogMetrics sets DisableDatadogMetrics field to given value.


### GetIsCloud

`func (o *GetEnvironmentInfoResponse) GetIsCloud() bool`

GetIsCloud returns the IsCloud field if non-nil, zero value otherwise.

### GetIsCloudOk

`func (o *GetEnvironmentInfoResponse) GetIsCloudOk() (*bool, bool)`

GetIsCloudOk returns a tuple with the IsCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloud

`func (o *GetEnvironmentInfoResponse) SetIsCloud(v bool)`

SetIsCloud sets IsCloud field to given value.


### GetGatewayUrl

`func (o *GetEnvironmentInfoResponse) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *GetEnvironmentInfoResponse) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *GetEnvironmentInfoResponse) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.


### GetProxyUrl

`func (o *GetEnvironmentInfoResponse) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *GetEnvironmentInfoResponse) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *GetEnvironmentInfoResponse) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.


### GetLicenseMetadata

`func (o *GetEnvironmentInfoResponse) GetLicenseMetadata() LicenseMetadata`

GetLicenseMetadata returns the LicenseMetadata field if non-nil, zero value otherwise.

### GetLicenseMetadataOk

`func (o *GetEnvironmentInfoResponse) GetLicenseMetadataOk() (*LicenseMetadata, bool)`

GetLicenseMetadataOk returns a tuple with the LicenseMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseMetadata

`func (o *GetEnvironmentInfoResponse) SetLicenseMetadata(v LicenseMetadata)`

SetLicenseMetadata sets LicenseMetadata field to given value.

### HasLicenseMetadata

`func (o *GetEnvironmentInfoResponse) HasLicenseMetadata() bool`

HasLicenseMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


