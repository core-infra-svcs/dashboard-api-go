# InlineResponse200493SsidAddressingMdns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, mDNS is enabled for the SSID or cluster | [optional] 
**Rules** | Pointer to [**[]InlineResponse200493SsidAddressingMdnsRules**](InlineResponse200493SsidAddressingMdnsRules.md) | List of mDNS forwarding rules | [optional] 

## Methods

### NewInlineResponse200493SsidAddressingMdns

`func NewInlineResponse200493SsidAddressingMdns() *InlineResponse200493SsidAddressingMdns`

NewInlineResponse200493SsidAddressingMdns instantiates a new InlineResponse200493SsidAddressingMdns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200493SsidAddressingMdnsWithDefaults

`func NewInlineResponse200493SsidAddressingMdnsWithDefaults() *InlineResponse200493SsidAddressingMdns`

NewInlineResponse200493SsidAddressingMdnsWithDefaults instantiates a new InlineResponse200493SsidAddressingMdns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200493SsidAddressingMdns) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200493SsidAddressingMdns) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200493SsidAddressingMdns) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200493SsidAddressingMdns) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse200493SsidAddressingMdns) GetRules() []InlineResponse200493SsidAddressingMdnsRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200493SsidAddressingMdns) GetRulesOk() (*[]InlineResponse200493SsidAddressingMdnsRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200493SsidAddressingMdns) SetRules(v []InlineResponse200493SsidAddressingMdnsRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200493SsidAddressingMdns) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


