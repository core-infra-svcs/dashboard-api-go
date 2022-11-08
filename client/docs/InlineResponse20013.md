# InlineResponse20013

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The VLAN ID of the VLAN | [optional] 
**InterfaceId** | Pointer to **string** | The interface ID of the VLAN | [optional] 
**Name** | Pointer to **string** | The name of the VLAN | [optional] 
**Subnet** | Pointer to **string** | The subnet of the VLAN | [optional] 
**ApplianceIp** | Pointer to **string** | The local IP of the appliance on the VLAN | [optional] 
**GroupPolicyId** | Pointer to **string** | The id of the desired group policy to apply to the VLAN | [optional] 
**TemplateVlanType** | Pointer to **string** | Type of subnetting of the VLAN. Applicable only for template network. | [optional] [default to "same"]
**Cidr** | Pointer to **string** | CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN. | [optional] 
**Mask** | Pointer to **int32** | Mask used for the subnet of all bound to the template networks. Applicable only for template network. | [optional] 
**DhcpRelayServerIps** | Pointer to **[]string** | The IPs of the DHCP servers that DHCP requests should be relayed to | [optional] 
**DhcpHandling** | Pointer to **string** | The appliance&#39;s handling of DHCP requests on this VLAN. One of: &#39;Run a DHCP server&#39;, &#39;Relay DHCP to another server&#39; or &#39;Do not respond to DHCP requests&#39; | [optional] 
**DhcpLeaseTime** | Pointer to **string** | The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: &#39;30 minutes&#39;, &#39;1 hour&#39;, &#39;4 hours&#39;, &#39;12 hours&#39;, &#39;1 day&#39; or &#39;1 week&#39; | [optional] 
**DhcpBootOptionsEnabled** | Pointer to **bool** | Use DHCP boot options specified in other properties | [optional] 
**DhcpBootNextServer** | Pointer to **string** | DHCP boot option to direct boot clients to the server to load the boot file from | [optional] 
**DhcpBootFilename** | Pointer to **string** | DHCP boot option for boot filename | [optional] 
**FixedIpAssignments** | Pointer to **map[string]interface{}** | The DHCP fixed IP assignments on the VLAN. This should be an object that contains mappings from MAC addresses to objects that themselves each contain \&quot;ip\&quot; and \&quot;name\&quot; string fields. See the sample request/response for more details. | [optional] 
**ReservedIpRanges** | Pointer to [**[]NetworksNetworkIdApplianceVlansReservedIpRanges**](NetworksNetworkIdApplianceVlansReservedIpRanges.md) | The DHCP reserved IP ranges on the VLAN | [optional] 
**DnsNameservers** | Pointer to **string** | The DNS nameservers used for DHCP responses, either \&quot;upstream_dns\&quot;, \&quot;google_dns\&quot;, \&quot;opendns\&quot;, or a newline seperated string of IP addresses or domain names | [optional] 
**DhcpOptions** | Pointer to [**[]NetworksNetworkIdApplianceVlansDhcpOptions**](NetworksNetworkIdApplianceVlansDhcpOptions.md) | The list of DHCP options that will be included in DHCP responses. Each object in the list should have \&quot;code\&quot;, \&quot;type\&quot;, and \&quot;value\&quot; properties. | [optional] 
**VpnNatSubnet** | Pointer to **string** | The translated VPN subnet if VPN and VPN subnet translation are enabled on the VLAN | [optional] 
**Ipv6** | Pointer to [**NetworksNetworkIdApplianceVlansIpv6**](NetworksNetworkIdApplianceVlansIpv6.md) |  | [optional] 

## Methods

### NewInlineResponse20013

`func NewInlineResponse20013() *InlineResponse20013`

NewInlineResponse20013 instantiates a new InlineResponse20013 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20013WithDefaults

`func NewInlineResponse20013WithDefaults() *InlineResponse20013`

NewInlineResponse20013WithDefaults instantiates a new InlineResponse20013 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20013) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20013) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20013) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20013) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceId

`func (o *InlineResponse20013) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *InlineResponse20013) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *InlineResponse20013) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *InlineResponse20013) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20013) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20013) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20013) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20013) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse20013) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse20013) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse20013) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse20013) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetApplianceIp

`func (o *InlineResponse20013) GetApplianceIp() string`

GetApplianceIp returns the ApplianceIp field if non-nil, zero value otherwise.

### GetApplianceIpOk

`func (o *InlineResponse20013) GetApplianceIpOk() (*string, bool)`

GetApplianceIpOk returns a tuple with the ApplianceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceIp

`func (o *InlineResponse20013) SetApplianceIp(v string)`

SetApplianceIp sets ApplianceIp field to given value.

### HasApplianceIp

`func (o *InlineResponse20013) HasApplianceIp() bool`

HasApplianceIp returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *InlineResponse20013) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineResponse20013) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineResponse20013) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *InlineResponse20013) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.

### GetTemplateVlanType

`func (o *InlineResponse20013) GetTemplateVlanType() string`

GetTemplateVlanType returns the TemplateVlanType field if non-nil, zero value otherwise.

### GetTemplateVlanTypeOk

`func (o *InlineResponse20013) GetTemplateVlanTypeOk() (*string, bool)`

GetTemplateVlanTypeOk returns a tuple with the TemplateVlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVlanType

`func (o *InlineResponse20013) SetTemplateVlanType(v string)`

SetTemplateVlanType sets TemplateVlanType field to given value.

### HasTemplateVlanType

`func (o *InlineResponse20013) HasTemplateVlanType() bool`

HasTemplateVlanType returns a boolean if a field has been set.

### GetCidr

`func (o *InlineResponse20013) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InlineResponse20013) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InlineResponse20013) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *InlineResponse20013) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetMask

`func (o *InlineResponse20013) GetMask() int32`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *InlineResponse20013) GetMaskOk() (*int32, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *InlineResponse20013) SetMask(v int32)`

SetMask sets Mask field to given value.

### HasMask

`func (o *InlineResponse20013) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetDhcpRelayServerIps

`func (o *InlineResponse20013) GetDhcpRelayServerIps() []string`

GetDhcpRelayServerIps returns the DhcpRelayServerIps field if non-nil, zero value otherwise.

### GetDhcpRelayServerIpsOk

`func (o *InlineResponse20013) GetDhcpRelayServerIpsOk() (*[]string, bool)`

GetDhcpRelayServerIpsOk returns a tuple with the DhcpRelayServerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelayServerIps

`func (o *InlineResponse20013) SetDhcpRelayServerIps(v []string)`

SetDhcpRelayServerIps sets DhcpRelayServerIps field to given value.

### HasDhcpRelayServerIps

`func (o *InlineResponse20013) HasDhcpRelayServerIps() bool`

HasDhcpRelayServerIps returns a boolean if a field has been set.

### GetDhcpHandling

`func (o *InlineResponse20013) GetDhcpHandling() string`

GetDhcpHandling returns the DhcpHandling field if non-nil, zero value otherwise.

### GetDhcpHandlingOk

`func (o *InlineResponse20013) GetDhcpHandlingOk() (*string, bool)`

GetDhcpHandlingOk returns a tuple with the DhcpHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpHandling

`func (o *InlineResponse20013) SetDhcpHandling(v string)`

SetDhcpHandling sets DhcpHandling field to given value.

### HasDhcpHandling

`func (o *InlineResponse20013) HasDhcpHandling() bool`

HasDhcpHandling returns a boolean if a field has been set.

### GetDhcpLeaseTime

`func (o *InlineResponse20013) GetDhcpLeaseTime() string`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *InlineResponse20013) GetDhcpLeaseTimeOk() (*string, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *InlineResponse20013) SetDhcpLeaseTime(v string)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *InlineResponse20013) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetDhcpBootOptionsEnabled

`func (o *InlineResponse20013) GetDhcpBootOptionsEnabled() bool`

GetDhcpBootOptionsEnabled returns the DhcpBootOptionsEnabled field if non-nil, zero value otherwise.

### GetDhcpBootOptionsEnabledOk

`func (o *InlineResponse20013) GetDhcpBootOptionsEnabledOk() (*bool, bool)`

GetDhcpBootOptionsEnabledOk returns a tuple with the DhcpBootOptionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpBootOptionsEnabled

`func (o *InlineResponse20013) SetDhcpBootOptionsEnabled(v bool)`

SetDhcpBootOptionsEnabled sets DhcpBootOptionsEnabled field to given value.

### HasDhcpBootOptionsEnabled

`func (o *InlineResponse20013) HasDhcpBootOptionsEnabled() bool`

HasDhcpBootOptionsEnabled returns a boolean if a field has been set.

### GetDhcpBootNextServer

`func (o *InlineResponse20013) GetDhcpBootNextServer() string`

GetDhcpBootNextServer returns the DhcpBootNextServer field if non-nil, zero value otherwise.

### GetDhcpBootNextServerOk

`func (o *InlineResponse20013) GetDhcpBootNextServerOk() (*string, bool)`

GetDhcpBootNextServerOk returns a tuple with the DhcpBootNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpBootNextServer

`func (o *InlineResponse20013) SetDhcpBootNextServer(v string)`

SetDhcpBootNextServer sets DhcpBootNextServer field to given value.

### HasDhcpBootNextServer

`func (o *InlineResponse20013) HasDhcpBootNextServer() bool`

HasDhcpBootNextServer returns a boolean if a field has been set.

### GetDhcpBootFilename

`func (o *InlineResponse20013) GetDhcpBootFilename() string`

GetDhcpBootFilename returns the DhcpBootFilename field if non-nil, zero value otherwise.

### GetDhcpBootFilenameOk

`func (o *InlineResponse20013) GetDhcpBootFilenameOk() (*string, bool)`

GetDhcpBootFilenameOk returns a tuple with the DhcpBootFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpBootFilename

`func (o *InlineResponse20013) SetDhcpBootFilename(v string)`

SetDhcpBootFilename sets DhcpBootFilename field to given value.

### HasDhcpBootFilename

`func (o *InlineResponse20013) HasDhcpBootFilename() bool`

HasDhcpBootFilename returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *InlineResponse20013) GetFixedIpAssignments() map[string]interface{}`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *InlineResponse20013) GetFixedIpAssignmentsOk() (*map[string]interface{}, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *InlineResponse20013) SetFixedIpAssignments(v map[string]interface{})`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *InlineResponse20013) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.

### GetReservedIpRanges

`func (o *InlineResponse20013) GetReservedIpRanges() []NetworksNetworkIdApplianceVlansReservedIpRanges`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *InlineResponse20013) GetReservedIpRangesOk() (*[]NetworksNetworkIdApplianceVlansReservedIpRanges, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *InlineResponse20013) SetReservedIpRanges(v []NetworksNetworkIdApplianceVlansReservedIpRanges)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *InlineResponse20013) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.

### GetDnsNameservers

`func (o *InlineResponse20013) GetDnsNameservers() string`

GetDnsNameservers returns the DnsNameservers field if non-nil, zero value otherwise.

### GetDnsNameserversOk

`func (o *InlineResponse20013) GetDnsNameserversOk() (*string, bool)`

GetDnsNameserversOk returns a tuple with the DnsNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameservers

`func (o *InlineResponse20013) SetDnsNameservers(v string)`

SetDnsNameservers sets DnsNameservers field to given value.

### HasDnsNameservers

`func (o *InlineResponse20013) HasDnsNameservers() bool`

HasDnsNameservers returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *InlineResponse20013) GetDhcpOptions() []NetworksNetworkIdApplianceVlansDhcpOptions`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *InlineResponse20013) GetDhcpOptionsOk() (*[]NetworksNetworkIdApplianceVlansDhcpOptions, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *InlineResponse20013) SetDhcpOptions(v []NetworksNetworkIdApplianceVlansDhcpOptions)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *InlineResponse20013) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetVpnNatSubnet

`func (o *InlineResponse20013) GetVpnNatSubnet() string`

GetVpnNatSubnet returns the VpnNatSubnet field if non-nil, zero value otherwise.

### GetVpnNatSubnetOk

`func (o *InlineResponse20013) GetVpnNatSubnetOk() (*string, bool)`

GetVpnNatSubnetOk returns a tuple with the VpnNatSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnNatSubnet

`func (o *InlineResponse20013) SetVpnNatSubnet(v string)`

SetVpnNatSubnet sets VpnNatSubnet field to given value.

### HasVpnNatSubnet

`func (o *InlineResponse20013) HasVpnNatSubnet() bool`

HasVpnNatSubnet returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineResponse20013) GetIpv6() NetworksNetworkIdApplianceVlansIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineResponse20013) GetIpv6Ok() (*NetworksNetworkIdApplianceVlansIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineResponse20013) SetIpv6(v NetworksNetworkIdApplianceVlansIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineResponse20013) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


