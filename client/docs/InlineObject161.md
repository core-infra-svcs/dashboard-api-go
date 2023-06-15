# InlineObject161

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules**](NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules.md) | An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule) | [optional] 
**AllowLanAccess** | Pointer to **bool** | Allow wireless client access to local LAN (boolean value - true allows access and false denies access) (optional) | [optional] 

## Methods

### NewInlineObject161

`func NewInlineObject161() *InlineObject161`

NewInlineObject161 instantiates a new InlineObject161 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject161WithDefaults

`func NewInlineObject161WithDefaults() *InlineObject161`

NewInlineObject161WithDefaults instantiates a new InlineObject161 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineObject161) GetRules() []NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject161) GetRulesOk() (*[]NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject161) SetRules(v []NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject161) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetAllowLanAccess

`func (o *InlineObject161) GetAllowLanAccess() bool`

GetAllowLanAccess returns the AllowLanAccess field if non-nil, zero value otherwise.

### GetAllowLanAccessOk

`func (o *InlineObject161) GetAllowLanAccessOk() (*bool, bool)`

GetAllowLanAccessOk returns a tuple with the AllowLanAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLanAccess

`func (o *InlineObject161) SetAllowLanAccess(v bool)`

SetAllowLanAccess sets AllowLanAccess field to given value.

### HasAllowLanAccess

`func (o *InlineObject161) HasAllowLanAccess() bool`

HasAllowLanAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


