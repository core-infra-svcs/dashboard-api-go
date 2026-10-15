# InlineObject94

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The VLAN ID of the new VLAN (must be between 1 and 4094) | 
**Name** | **string** | The name of the new VLAN | 
**Subnet** | Pointer to **string** | The subnet of the VLAN | [optional] 
**ApplianceIp** | Pointer to **string** | The local IP of the appliance on the VLAN | [optional] 
**GroupPolicyId** | Pointer to **NullableString** | The id of the desired group policy to apply to the VLAN | [optional] 
**TemplateVlanType** | Pointer to **string** | Type of subnetting of the VLAN. Applicable only for template network. | [optional] [default to "same"]
**Cidr** | Pointer to **string** | CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN. | [optional] 
**Mask** | Pointer to **int32** | Mask used for the subnet of all bound to the template networks. Applicable only for template network. | [optional] 
**Ipv6** | Pointer to [**NetworksNetworkIdApplianceVlansIpv61**](NetworksNetworkIdApplianceVlansIpv61.md) |  | [optional] 
**DhcpHandling** | Pointer to **string** | The appliance&#39;s handling of DHCP requests on this VLAN. One of: &#39;Run a DHCP server&#39;, &#39;Relay DHCP to another server&#39; or &#39;Do not respond to DHCP requests&#39; | [optional] 
**DhcpRelayServerIps** | Pointer to **[]string** | The IPs (IPv4) of the DHCP servers that DHCP requests should be relayed to. CIDR/subnet notation and hostnames are not supported. | [optional] 
**DhcpLeaseTime** | Pointer to **string** | The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: &#39;30 minutes&#39;, &#39;1 hour&#39;, &#39;4 hours&#39;, &#39;12 hours&#39;, &#39;1 day&#39; or &#39;1 week&#39; | [optional] 
**MandatoryDhcp** | Pointer to [**NetworksNetworkIdApplianceVlansMandatoryDhcp**](NetworksNetworkIdApplianceVlansMandatoryDhcp.md) |  | [optional] 
**DhcpBootOptionsEnabled** | Pointer to **bool** | Use DHCP boot options specified in other properties | [optional] 
**DhcpBootNextServer** | Pointer to **string** | DHCP boot option to direct boot clients to the server to load the boot file from | [optional] 
**DhcpBootFilename** | Pointer to **string** | DHCP boot option for boot filename | [optional] 
**DhcpOptions** | Pointer to [**[]NetworksNetworkIdApplianceVlansDhcpOptions**](NetworksNetworkIdApplianceVlansDhcpOptions.md) | The list of DHCP options that will be included in DHCP responses. Each object in the list should have \&quot;code\&quot;, \&quot;type\&quot;, and \&quot;value\&quot; properties. | [optional] 
**Sgt** | Pointer to [**NetworksNetworkIdApplianceVlansSgt**](NetworksNetworkIdApplianceVlansSgt.md) |  | [optional] 
**Vrf** | Pointer to [**NullableNetworksNetworkIdApplianceVlansVrf1**](NetworksNetworkIdApplianceVlansVrf1.md) |  | [optional] 
**Uplinks** | Pointer to [**[]NetworksNetworkIdApplianceUplinksNatUplinks**](NetworksNetworkIdApplianceUplinksNatUplinks.md) | Per-uplink NAT exception override configuration on the VLAN. Applicable only for networks that support NAT exceptions. | [optional] 

## Methods

### NewInlineObject94

`func NewInlineObject94(id string, name string, ) *InlineObject94`

NewInlineObject94 instantiates a new InlineObject94 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject94WithDefaults

`func NewInlineObject94WithDefaults() *InlineObject94`

NewInlineObject94WithDefaults instantiates a new InlineObject94 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineObject94) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineObject94) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineObject94) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *InlineObject94) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject94) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject94) SetName(v string)`

SetName sets Name field to given value.


### GetSubnet

`func (o *InlineObject94) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject94) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject94) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineObject94) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetApplianceIp

`func (o *InlineObject94) GetApplianceIp() string`

GetApplianceIp returns the ApplianceIp field if non-nil, zero value otherwise.

### GetApplianceIpOk

`func (o *InlineObject94) GetApplianceIpOk() (*string, bool)`

GetApplianceIpOk returns a tuple with the ApplianceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceIp

`func (o *InlineObject94) SetApplianceIp(v string)`

SetApplianceIp sets ApplianceIp field to given value.

### HasApplianceIp

`func (o *InlineObject94) HasApplianceIp() bool`

HasApplianceIp returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *InlineObject94) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineObject94) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineObject94) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *InlineObject94) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.

### SetGroupPolicyIdNil

`func (o *InlineObject94) SetGroupPolicyIdNil(b bool)`

 SetGroupPolicyIdNil sets the value for GroupPolicyId to be an explicit nil

### UnsetGroupPolicyId
`func (o *InlineObject94) UnsetGroupPolicyId()`

UnsetGroupPolicyId ensures that no value is present for GroupPolicyId, not even an explicit nil
### GetTemplateVlanType

`func (o *InlineObject94) GetTemplateVlanType() string`

GetTemplateVlanType returns the TemplateVlanType field if non-nil, zero value otherwise.

### GetTemplateVlanTypeOk

`func (o *InlineObject94) GetTemplateVlanTypeOk() (*string, bool)`

GetTemplateVlanTypeOk returns a tuple with the TemplateVlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVlanType

`func (o *InlineObject94) SetTemplateVlanType(v string)`

SetTemplateVlanType sets TemplateVlanType field to given value.

### HasTemplateVlanType

`func (o *InlineObject94) HasTemplateVlanType() bool`

HasTemplateVlanType returns a boolean if a field has been set.

### GetCidr

`func (o *InlineObject94) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InlineObject94) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InlineObject94) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *InlineObject94) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetMask

`func (o *InlineObject94) GetMask() int32`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *InlineObject94) GetMaskOk() (*int32, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *InlineObject94) SetMask(v int32)`

SetMask sets Mask field to given value.

### HasMask

`func (o *InlineObject94) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineObject94) GetIpv6() NetworksNetworkIdApplianceVlansIpv61`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineObject94) GetIpv6Ok() (*NetworksNetworkIdApplianceVlansIpv61, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineObject94) SetIpv6(v NetworksNetworkIdApplianceVlansIpv61)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineObject94) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetDhcpHandling

`func (o *InlineObject94) GetDhcpHandling() string`

GetDhcpHandling returns the DhcpHandling field if non-nil, zero value otherwise.

### GetDhcpHandlingOk

`func (o *InlineObject94) GetDhcpHandlingOk() (*string, bool)`

GetDhcpHandlingOk returns a tuple with the DhcpHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpHandling

`func (o *InlineObject94) SetDhcpHandling(v string)`

SetDhcpHandling sets DhcpHandling field to given value.

### HasDhcpHandling

`func (o *InlineObject94) HasDhcpHandling() bool`

HasDhcpHandling returns a boolean if a field has been set.

### GetDhcpRelayServerIps

`func (o *InlineObject94) GetDhcpRelayServerIps() []string`

GetDhcpRelayServerIps returns the DhcpRelayServerIps field if non-nil, zero value otherwise.

### GetDhcpRelayServerIpsOk

`func (o *InlineObject94) GetDhcpRelayServerIpsOk() (*[]string, bool)`

GetDhcpRelayServerIpsOk returns a tuple with the DhcpRelayServerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelayServerIps

`func (o *InlineObject94) SetDhcpRelayServerIps(v []string)`

SetDhcpRelayServerIps sets DhcpRelayServerIps field to given value.

### HasDhcpRelayServerIps

`func (o *InlineObject94) HasDhcpRelayServerIps() bool`

HasDhcpRelayServerIps returns a boolean if a field has been set.

### GetDhcpLeaseTime

`func (o *InlineObject94) GetDhcpLeaseTime() string`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *InlineObject94) GetDhcpLeaseTimeOk() (*string, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *InlineObject94) SetDhcpLeaseTime(v string)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *InlineObject94) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetMandatoryDhcp

`func (o *InlineObject94) GetMandatoryDhcp() NetworksNetworkIdApplianceVlansMandatoryDhcp`

GetMandatoryDhcp returns the MandatoryDhcp field if non-nil, zero value otherwise.

### GetMandatoryDhcpOk

`func (o *InlineObject94) GetMandatoryDhcpOk() (*NetworksNetworkIdApplianceVlansMandatoryDhcp, bool)`

GetMandatoryDhcpOk returns a tuple with the MandatoryDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryDhcp

`func (o *InlineObject94) SetMandatoryDhcp(v NetworksNetworkIdApplianceVlansMandatoryDhcp)`

SetMandatoryDhcp sets MandatoryDhcp field to given value.

### HasMandatoryDhcp

`func (o *InlineObject94) HasMandatoryDhcp() bool`

HasMandatoryDhcp returns a boolean if a field has been set.

### GetDhcpBootOptionsEnabled

`func (o *InlineObject94) GetDhcpBootOptionsEnabled() bool`

GetDhcpBootOptionsEnabled returns the DhcpBootOptionsEnabled field if non-nil, zero value otherwise.

### GetDhcpBootOptionsEnabledOk

`func (o *InlineObject94) GetDhcpBootOptionsEnabledOk() (*bool, bool)`

GetDhcpBootOptionsEnabledOk returns a tuple with the DhcpBootOptionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpBootOptionsEnabled

`func (o *InlineObject94) SetDhcpBootOptionsEnabled(v bool)`

SetDhcpBootOptionsEnabled sets DhcpBootOptionsEnabled field to given value.

### HasDhcpBootOptionsEnabled

`func (o *InlineObject94) HasDhcpBootOptionsEnabled() bool`

HasDhcpBootOptionsEnabled returns a boolean if a field has been set.

### GetDhcpBootNextServer

`func (o *InlineObject94) GetDhcpBootNextServer() string`

GetDhcpBootNextServer returns the DhcpBootNextServer field if non-nil, zero value otherwise.

### GetDhcpBootNextServerOk

`func (o *InlineObject94) GetDhcpBootNextServerOk() (*string, bool)`

GetDhcpBootNextServerOk returns a tuple with the DhcpBootNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpBootNextServer

`func (o *InlineObject94) SetDhcpBootNextServer(v string)`

SetDhcpBootNextServer sets DhcpBootNextServer field to given value.

### HasDhcpBootNextServer

`func (o *InlineObject94) HasDhcpBootNextServer() bool`

HasDhcpBootNextServer returns a boolean if a field has been set.

### GetDhcpBootFilename

`func (o *InlineObject94) GetDhcpBootFilename() string`

GetDhcpBootFilename returns the DhcpBootFilename field if non-nil, zero value otherwise.

### GetDhcpBootFilenameOk

`func (o *InlineObject94) GetDhcpBootFilenameOk() (*string, bool)`

GetDhcpBootFilenameOk returns a tuple with the DhcpBootFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpBootFilename

`func (o *InlineObject94) SetDhcpBootFilename(v string)`

SetDhcpBootFilename sets DhcpBootFilename field to given value.

### HasDhcpBootFilename

`func (o *InlineObject94) HasDhcpBootFilename() bool`

HasDhcpBootFilename returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *InlineObject94) GetDhcpOptions() []NetworksNetworkIdApplianceVlansDhcpOptions`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *InlineObject94) GetDhcpOptionsOk() (*[]NetworksNetworkIdApplianceVlansDhcpOptions, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *InlineObject94) SetDhcpOptions(v []NetworksNetworkIdApplianceVlansDhcpOptions)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *InlineObject94) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetSgt

`func (o *InlineObject94) GetSgt() NetworksNetworkIdApplianceVlansSgt`

GetSgt returns the Sgt field if non-nil, zero value otherwise.

### GetSgtOk

`func (o *InlineObject94) GetSgtOk() (*NetworksNetworkIdApplianceVlansSgt, bool)`

GetSgtOk returns a tuple with the Sgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgt

`func (o *InlineObject94) SetSgt(v NetworksNetworkIdApplianceVlansSgt)`

SetSgt sets Sgt field to given value.

### HasSgt

`func (o *InlineObject94) HasSgt() bool`

HasSgt returns a boolean if a field has been set.

### GetVrf

`func (o *InlineObject94) GetVrf() NetworksNetworkIdApplianceVlansVrf1`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineObject94) GetVrfOk() (*NetworksNetworkIdApplianceVlansVrf1, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineObject94) SetVrf(v NetworksNetworkIdApplianceVlansVrf1)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineObject94) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *InlineObject94) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *InlineObject94) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetUplinks

`func (o *InlineObject94) GetUplinks() []NetworksNetworkIdApplianceUplinksNatUplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineObject94) GetUplinksOk() (*[]NetworksNetworkIdApplianceUplinksNatUplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineObject94) SetUplinks(v []NetworksNetworkIdApplianceUplinksNatUplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineObject94) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


