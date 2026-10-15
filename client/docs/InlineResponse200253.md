# InlineResponse200253

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]InlineResponse200253Rules**](InlineResponse200253Rules.md) | An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule). | [optional] 

## Methods

### NewInlineResponse200253

`func NewInlineResponse200253() *InlineResponse200253`

NewInlineResponse200253 instantiates a new InlineResponse200253 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200253WithDefaults

`func NewInlineResponse200253WithDefaults() *InlineResponse200253`

NewInlineResponse200253WithDefaults instantiates a new InlineResponse200253 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineResponse200253) GetRules() []InlineResponse200253Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200253) GetRulesOk() (*[]InlineResponse200253Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200253) SetRules(v []InlineResponse200253Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200253) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


