# InlineResponse200207

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]InlineResponse200207Rules**](InlineResponse200207Rules.md) | An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule). | [optional] 

## Methods

### NewInlineResponse200207

`func NewInlineResponse200207() *InlineResponse200207`

NewInlineResponse200207 instantiates a new InlineResponse200207 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200207WithDefaults

`func NewInlineResponse200207WithDefaults() *InlineResponse200207`

NewInlineResponse200207WithDefaults instantiates a new InlineResponse200207 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineResponse200207) GetRules() []InlineResponse200207Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200207) GetRulesOk() (*[]InlineResponse200207Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200207) SetRules(v []InlineResponse200207Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200207) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


