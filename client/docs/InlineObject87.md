# InlineObject87

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | **[]string** | Domain names to exclude from Umbrella DNS routing (e.g., &#39;example.com&#39;, &#39;corp.example.org&#39;). Standard FQDNs only — wildcards are not supported. Values are lowercased before saving. Each call replaces the full exclusion list. | 

## Methods

### NewInlineObject87

`func NewInlineObject87(domains []string, ) *InlineObject87`

NewInlineObject87 instantiates a new InlineObject87 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject87WithDefaults

`func NewInlineObject87WithDefaults() *InlineObject87`

NewInlineObject87WithDefaults instantiates a new InlineObject87 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *InlineObject87) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *InlineObject87) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *InlineObject87) SetDomains(v []string)`

SetDomains sets Domains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


