# GetSampleEnrichmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | [**[]GetSampleEnrichmentResponseSamplesInner**](GetSampleEnrichmentResponseSamplesInner.md) | One entry per matched ES document, keyed by sample_id (\&quot;{state}_{index}\&quot;). | 

## Methods

### NewGetSampleEnrichmentResponse

`func NewGetSampleEnrichmentResponse(samples []GetSampleEnrichmentResponseSamplesInner, ) *GetSampleEnrichmentResponse`

NewGetSampleEnrichmentResponse instantiates a new GetSampleEnrichmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSampleEnrichmentResponseWithDefaults

`func NewGetSampleEnrichmentResponseWithDefaults() *GetSampleEnrichmentResponse`

NewGetSampleEnrichmentResponseWithDefaults instantiates a new GetSampleEnrichmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamples

`func (o *GetSampleEnrichmentResponse) GetSamples() []GetSampleEnrichmentResponseSamplesInner`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *GetSampleEnrichmentResponse) GetSamplesOk() (*[]GetSampleEnrichmentResponseSamplesInner, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *GetSampleEnrichmentResponse) SetSamples(v []GetSampleEnrichmentResponseSamplesInner)`

SetSamples sets Samples field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


