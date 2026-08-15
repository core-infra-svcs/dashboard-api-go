# InlineResponse200488SsidAddressingMdns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, mDNS is enabled for the SSID or cluster | [optional] 
**Rules** | Pointer to [**[]InlineResponse200488SsidAddressingMdnsRules**](InlineResponse200488SsidAddressingMdnsRules.md) | List of mDNS forwarding rules | [optional] 

## Methods

### NewInlineResponse200488SsidAddressingMdns

`func NewInlineResponse200488SsidAddressingMdns() *InlineResponse200488SsidAddressingMdns`

NewInlineResponse200488SsidAddressingMdns instantiates a new InlineResponse200488SsidAddressingMdns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200488SsidAddressingMdnsWithDefaults

`func NewInlineResponse200488SsidAddressingMdnsWithDefaults() *InlineResponse200488SsidAddressingMdns`

NewInlineResponse200488SsidAddressingMdnsWithDefaults instantiates a new InlineResponse200488SsidAddressingMdns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200488SsidAddressingMdns) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200488SsidAddressingMdns) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200488SsidAddressingMdns) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200488SsidAddressingMdns) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse200488SsidAddressingMdns) GetRules() []InlineResponse200488SsidAddressingMdnsRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200488SsidAddressingMdns) GetRulesOk() (*[]InlineResponse200488SsidAddressingMdnsRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200488SsidAddressingMdns) SetRules(v []InlineResponse200488SsidAddressingMdnsRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200488SsidAddressingMdns) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


