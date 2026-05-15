# InlineObject150

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | The updated mac address of the trusted server | [optional] 
**Vlan** | Pointer to **int32** | The updated VLAN of the trusted server. It must be between 1 and 4094 | [optional] 
**Ipv4** | Pointer to [**NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4**](NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4.md) |  | [optional] 

## Methods

### NewInlineObject150

`func NewInlineObject150() *InlineObject150`

NewInlineObject150 instantiates a new InlineObject150 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject150WithDefaults

`func NewInlineObject150WithDefaults() *InlineObject150`

NewInlineObject150WithDefaults instantiates a new InlineObject150 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *InlineObject150) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineObject150) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineObject150) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineObject150) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetVlan

`func (o *InlineObject150) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineObject150) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineObject150) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineObject150) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetIpv4

`func (o *InlineObject150) GetIpv4() NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *InlineObject150) GetIpv4Ok() (*NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *InlineObject150) SetIpv4(v NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *InlineObject150) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


