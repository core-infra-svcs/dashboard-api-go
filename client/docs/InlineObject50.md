# InlineObject50

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | Pointer to **string** | The subnet of the single LAN configuration | [optional] 
**ApplianceIp** | Pointer to **string** | The appliance IP address of the single LAN | [optional] 
**Ipv6** | Pointer to [**NetworksNetworkIdApplianceSingleLanIpv6**](NetworksNetworkIdApplianceSingleLanIpv6.md) |  | [optional] 
**MandatoryDhcp** | Pointer to [**NetworksNetworkIdApplianceSingleLanMandatoryDhcp**](NetworksNetworkIdApplianceSingleLanMandatoryDhcp.md) |  | [optional] 

## Methods

### NewInlineObject50

`func NewInlineObject50() *InlineObject50`

NewInlineObject50 instantiates a new InlineObject50 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject50WithDefaults

`func NewInlineObject50WithDefaults() *InlineObject50`

NewInlineObject50WithDefaults instantiates a new InlineObject50 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *InlineObject50) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject50) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject50) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineObject50) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetApplianceIp

`func (o *InlineObject50) GetApplianceIp() string`

GetApplianceIp returns the ApplianceIp field if non-nil, zero value otherwise.

### GetApplianceIpOk

`func (o *InlineObject50) GetApplianceIpOk() (*string, bool)`

GetApplianceIpOk returns a tuple with the ApplianceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceIp

`func (o *InlineObject50) SetApplianceIp(v string)`

SetApplianceIp sets ApplianceIp field to given value.

### HasApplianceIp

`func (o *InlineObject50) HasApplianceIp() bool`

HasApplianceIp returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineObject50) GetIpv6() NetworksNetworkIdApplianceSingleLanIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineObject50) GetIpv6Ok() (*NetworksNetworkIdApplianceSingleLanIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineObject50) SetIpv6(v NetworksNetworkIdApplianceSingleLanIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineObject50) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetMandatoryDhcp

`func (o *InlineObject50) GetMandatoryDhcp() NetworksNetworkIdApplianceSingleLanMandatoryDhcp`

GetMandatoryDhcp returns the MandatoryDhcp field if non-nil, zero value otherwise.

### GetMandatoryDhcpOk

`func (o *InlineObject50) GetMandatoryDhcpOk() (*NetworksNetworkIdApplianceSingleLanMandatoryDhcp, bool)`

GetMandatoryDhcpOk returns a tuple with the MandatoryDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryDhcp

`func (o *InlineObject50) SetMandatoryDhcp(v NetworksNetworkIdApplianceSingleLanMandatoryDhcp)`

SetMandatoryDhcp sets MandatoryDhcp field to given value.

### HasMandatoryDhcp

`func (o *InlineObject50) HasMandatoryDhcp() bool`

HasMandatoryDhcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


