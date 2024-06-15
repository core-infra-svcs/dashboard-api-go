# InlineResponse200189

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]InlineResponse200189Rules**](InlineResponse200189Rules.md) | An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule). | [optional] 

## Methods

### NewInlineResponse200189

`func NewInlineResponse200189() *InlineResponse200189`

NewInlineResponse200189 instantiates a new InlineResponse200189 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200189WithDefaults

`func NewInlineResponse200189WithDefaults() *InlineResponse200189`

NewInlineResponse200189WithDefaults instantiates a new InlineResponse200189 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineResponse200189) GetRules() []InlineResponse200189Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200189) GetRulesOk() (*[]InlineResponse200189Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200189) SetRules(v []InlineResponse200189Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200189) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


