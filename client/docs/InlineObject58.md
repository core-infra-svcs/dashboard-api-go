# InlineObject58

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]InlineResponse20077Rules**](InlineResponse20077Rules.md) | An ordered array of the MX L7 firewall rules. Each rule is an object with &#39;policy&#39;, &#39;type&#39;, and &#39;value&#39;. The &#39;value&#39; shape depends on &#39;type&#39;: object for application/applicationCategory, string for host/port/ipRange, and an array of 2-letter ISO 3166-1 alpha-2 country codes for allowedCountries/blockedCountries. For backward compatibility, request types also accept whitelistedCountries/blacklistedCountries. | [optional] 

## Methods

### NewInlineObject58

`func NewInlineObject58() *InlineObject58`

NewInlineObject58 instantiates a new InlineObject58 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject58WithDefaults

`func NewInlineObject58WithDefaults() *InlineObject58`

NewInlineObject58WithDefaults instantiates a new InlineObject58 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineObject58) GetRules() []InlineResponse20077Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject58) GetRulesOk() (*[]InlineResponse20077Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject58) SetRules(v []InlineResponse20077Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject58) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


