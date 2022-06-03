# InlineObject96

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | **string** | The mac address of the trusted server being added | 
**Vlan** | **int32** | The VLAN of the trusted server being added. It must be between 1 and 4094 | 
**Ipv4** | [**NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4**](NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4.md) |  | 

## Methods

### NewInlineObject96

`func NewInlineObject96(mac string, vlan int32, ipv4 NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4, ) *InlineObject96`

NewInlineObject96 instantiates a new InlineObject96 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject96WithDefaults

`func NewInlineObject96WithDefaults() *InlineObject96`

NewInlineObject96WithDefaults instantiates a new InlineObject96 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *InlineObject96) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineObject96) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineObject96) SetMac(v string)`

SetMac sets Mac field to given value.


### GetVlan

`func (o *InlineObject96) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineObject96) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineObject96) SetVlan(v int32)`

SetVlan sets Vlan field to given value.


### GetIpv4

`func (o *InlineObject96) GetIpv4() NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *InlineObject96) GetIpv4Ok() (*NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *InlineObject96) SetIpv4(v NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4)`

SetIpv4 sets Ipv4 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


