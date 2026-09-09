# GetAnalysisSampleAssetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractVersion** | **float64** |  | 
**Samples** | [**[]AnalysisSampleAssets**](AnalysisSampleAssets.md) |  | 

## Methods

### NewGetAnalysisSampleAssetsResponse

`func NewGetAnalysisSampleAssetsResponse(contractVersion float64, samples []AnalysisSampleAssets, ) *GetAnalysisSampleAssetsResponse`

NewGetAnalysisSampleAssetsResponse instantiates a new GetAnalysisSampleAssetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalysisSampleAssetsResponseWithDefaults

`func NewGetAnalysisSampleAssetsResponseWithDefaults() *GetAnalysisSampleAssetsResponse`

NewGetAnalysisSampleAssetsResponseWithDefaults instantiates a new GetAnalysisSampleAssetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractVersion

`func (o *GetAnalysisSampleAssetsResponse) GetContractVersion() float64`

GetContractVersion returns the ContractVersion field if non-nil, zero value otherwise.

### GetContractVersionOk

`func (o *GetAnalysisSampleAssetsResponse) GetContractVersionOk() (*float64, bool)`

GetContractVersionOk returns a tuple with the ContractVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractVersion

`func (o *GetAnalysisSampleAssetsResponse) SetContractVersion(v float64)`

SetContractVersion sets ContractVersion field to given value.


### GetSamples

`func (o *GetAnalysisSampleAssetsResponse) GetSamples() []AnalysisSampleAssets`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *GetAnalysisSampleAssetsResponse) GetSamplesOk() (*[]AnalysisSampleAssets, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *GetAnalysisSampleAssetsResponse) SetSamples(v []AnalysisSampleAssets)`

SetSamples sets Samples field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


