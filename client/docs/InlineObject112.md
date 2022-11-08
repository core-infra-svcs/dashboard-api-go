# InlineObject112

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | **string** | The mac address of the trusted server being added | 
**Vlan** | **int32** | The VLAN of the trusted server being added. It must be between 1 and 4094 | 
**Ipv4** | [**NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41**](NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41.md) |  | 

## Methods

### NewInlineObject112

`func NewInlineObject112(mac string, vlan int32, ipv4 NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41, ) *InlineObject112`

NewInlineObject112 instantiates a new InlineObject112 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject112WithDefaults

`func NewInlineObject112WithDefaults() *InlineObject112`

NewInlineObject112WithDefaults instantiates a new InlineObject112 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *InlineObject112) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineObject112) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineObject112) SetMac(v string)`

SetMac sets Mac field to given value.


### GetVlan

`func (o *InlineObject112) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineObject112) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineObject112) SetVlan(v int32)`

SetVlan sets Vlan field to given value.


### GetIpv4

`func (o *InlineObject112) GetIpv4() NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *InlineObject112) GetIpv4Ok() (*NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *InlineObject112) SetIpv4(v NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41)`

SetIpv4 sets Ipv4 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


