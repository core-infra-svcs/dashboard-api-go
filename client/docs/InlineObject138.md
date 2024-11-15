# InlineObject138

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | The updated mac address of the trusted server | [optional] 
**Vlan** | Pointer to **int32** | The updated VLAN of the trusted server. It must be between 1 and 4094 | [optional] 
**Ipv4** | Pointer to [**NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4**](NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4.md) |  | [optional] 

## Methods

### NewInlineObject138

`func NewInlineObject138() *InlineObject138`

NewInlineObject138 instantiates a new InlineObject138 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject138WithDefaults

`func NewInlineObject138WithDefaults() *InlineObject138`

NewInlineObject138WithDefaults instantiates a new InlineObject138 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *InlineObject138) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineObject138) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineObject138) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineObject138) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetVlan

`func (o *InlineObject138) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineObject138) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineObject138) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineObject138) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetIpv4

`func (o *InlineObject138) GetIpv4() NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *InlineObject138) GetIpv4Ok() (*NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *InlineObject138) SetIpv4(v NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *InlineObject138) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


