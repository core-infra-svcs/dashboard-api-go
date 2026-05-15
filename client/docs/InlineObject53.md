# InlineObject53

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]InlineResponse20072Rules**](InlineResponse20072Rules.md) | An ordered array of the MX L7 firewall rules. Each rule is an object with &#39;policy&#39;, &#39;type&#39;, and &#39;value&#39;. The &#39;value&#39; shape depends on &#39;type&#39;: object for application/applicationCategory, string for host/port/ipRange, and an array of 2-letter ISO 3166-1 alpha-2 country codes for allowedCountries/blockedCountries. For backward compatibility, request types also accept whitelistedCountries/blacklistedCountries. | [optional] 

## Methods

### NewInlineObject53

`func NewInlineObject53() *InlineObject53`

NewInlineObject53 instantiates a new InlineObject53 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject53WithDefaults

`func NewInlineObject53WithDefaults() *InlineObject53`

NewInlineObject53WithDefaults instantiates a new InlineObject53 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineObject53) GetRules() []InlineResponse20072Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject53) GetRulesOk() (*[]InlineResponse20072Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject53) SetRules(v []InlineResponse20072Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject53) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


