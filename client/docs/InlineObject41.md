# InlineObject41

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpMode** | Pointer to **string** | The DHCP mode options for the switch interface        (&#39;dhcpDisabled&#39;, &#39;dhcpRelay&#39; or &#39;dhcpServer&#39;) | [optional] 
**DhcpRelayServerIps** | Pointer to **[]string** | The DHCP relay server IPs to which DHCP packets would get relayed for the switch interface | [optional] 
**DhcpLeaseTime** | Pointer to **string** | The DHCP lease time config for the dhcp server running on switch interface         (&#39;30 minutes&#39;, &#39;1 hour&#39;, &#39;4 hours&#39;, &#39;12 hours&#39;, &#39;1 day&#39; or &#39;1 week&#39;) | [optional] 
**DnsNameserversOption** | Pointer to **string** | The DHCP name server option for the dhcp server running on the switch interface         (&#39;googlePublicDns&#39;, &#39;openDns&#39; or &#39;custom&#39;) | [optional] 
**DnsCustomNameservers** | Pointer to **[]string** | The DHCP name server IPs when DHCP name server option is         &#39;custom&#39; | [optional] 
**BootOptionsEnabled** | Pointer to **bool** | Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch         interface | [optional] 
**BootNextServer** | Pointer to **string** | The PXE boot server IP for the DHCP server running on the switch interface | [optional] 
**BootFileName** | Pointer to **string** | The PXE boot server filename for the DHCP server running on the switch interface | [optional] 
**DhcpOptions** | Pointer to [**[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions**](DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions.md) | Array of DHCP options consisting of code, type and value for the DHCP server running on the switch interface | [optional] 
**ReservedIpRanges** | Pointer to [**[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges**](DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges.md) | Array of DHCP reserved IP assignments for the DHCP server running on the switch interface | [optional] 
**FixedIpAssignments** | Pointer to [**[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments**](DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments.md) | Array of DHCP fixed IP assignments for the DHCP server running on the switch interface | [optional] 
**DhcpDefaultRouterIps** | Pointer to **string** | The DHCP default router IPs for the DHCP server running on the switch interface | [optional] 
**DhcpDomainName** | Pointer to **NullableString** | The DHCP domain name for the DHCP server running on the switch interface | [optional] 

## Methods

### NewInlineObject41

`func NewInlineObject41() *InlineObject41`

NewInlineObject41 instantiates a new InlineObject41 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject41WithDefaults

`func NewInlineObject41WithDefaults() *InlineObject41`

NewInlineObject41WithDefaults instantiates a new InlineObject41 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpMode

`func (o *InlineObject41) GetDhcpMode() string`

GetDhcpMode returns the DhcpMode field if non-nil, zero value otherwise.

### GetDhcpModeOk

`func (o *InlineObject41) GetDhcpModeOk() (*string, bool)`

GetDhcpModeOk returns a tuple with the DhcpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpMode

`func (o *InlineObject41) SetDhcpMode(v string)`

SetDhcpMode sets DhcpMode field to given value.

### HasDhcpMode

`func (o *InlineObject41) HasDhcpMode() bool`

HasDhcpMode returns a boolean if a field has been set.

### GetDhcpRelayServerIps

`func (o *InlineObject41) GetDhcpRelayServerIps() []string`

GetDhcpRelayServerIps returns the DhcpRelayServerIps field if non-nil, zero value otherwise.

### GetDhcpRelayServerIpsOk

`func (o *InlineObject41) GetDhcpRelayServerIpsOk() (*[]string, bool)`

GetDhcpRelayServerIpsOk returns a tuple with the DhcpRelayServerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelayServerIps

`func (o *InlineObject41) SetDhcpRelayServerIps(v []string)`

SetDhcpRelayServerIps sets DhcpRelayServerIps field to given value.

### HasDhcpRelayServerIps

`func (o *InlineObject41) HasDhcpRelayServerIps() bool`

HasDhcpRelayServerIps returns a boolean if a field has been set.

### GetDhcpLeaseTime

`func (o *InlineObject41) GetDhcpLeaseTime() string`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *InlineObject41) GetDhcpLeaseTimeOk() (*string, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *InlineObject41) SetDhcpLeaseTime(v string)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *InlineObject41) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetDnsNameserversOption

`func (o *InlineObject41) GetDnsNameserversOption() string`

GetDnsNameserversOption returns the DnsNameserversOption field if non-nil, zero value otherwise.

### GetDnsNameserversOptionOk

`func (o *InlineObject41) GetDnsNameserversOptionOk() (*string, bool)`

GetDnsNameserversOptionOk returns a tuple with the DnsNameserversOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameserversOption

`func (o *InlineObject41) SetDnsNameserversOption(v string)`

SetDnsNameserversOption sets DnsNameserversOption field to given value.

### HasDnsNameserversOption

`func (o *InlineObject41) HasDnsNameserversOption() bool`

HasDnsNameserversOption returns a boolean if a field has been set.

### GetDnsCustomNameservers

`func (o *InlineObject41) GetDnsCustomNameservers() []string`

GetDnsCustomNameservers returns the DnsCustomNameservers field if non-nil, zero value otherwise.

### GetDnsCustomNameserversOk

`func (o *InlineObject41) GetDnsCustomNameserversOk() (*[]string, bool)`

GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCustomNameservers

`func (o *InlineObject41) SetDnsCustomNameservers(v []string)`

SetDnsCustomNameservers sets DnsCustomNameservers field to given value.

### HasDnsCustomNameservers

`func (o *InlineObject41) HasDnsCustomNameservers() bool`

HasDnsCustomNameservers returns a boolean if a field has been set.

### GetBootOptionsEnabled

`func (o *InlineObject41) GetBootOptionsEnabled() bool`

GetBootOptionsEnabled returns the BootOptionsEnabled field if non-nil, zero value otherwise.

### GetBootOptionsEnabledOk

`func (o *InlineObject41) GetBootOptionsEnabledOk() (*bool, bool)`

GetBootOptionsEnabledOk returns a tuple with the BootOptionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOptionsEnabled

`func (o *InlineObject41) SetBootOptionsEnabled(v bool)`

SetBootOptionsEnabled sets BootOptionsEnabled field to given value.

### HasBootOptionsEnabled

`func (o *InlineObject41) HasBootOptionsEnabled() bool`

HasBootOptionsEnabled returns a boolean if a field has been set.

### GetBootNextServer

`func (o *InlineObject41) GetBootNextServer() string`

GetBootNextServer returns the BootNextServer field if non-nil, zero value otherwise.

### GetBootNextServerOk

`func (o *InlineObject41) GetBootNextServerOk() (*string, bool)`

GetBootNextServerOk returns a tuple with the BootNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootNextServer

`func (o *InlineObject41) SetBootNextServer(v string)`

SetBootNextServer sets BootNextServer field to given value.

### HasBootNextServer

`func (o *InlineObject41) HasBootNextServer() bool`

HasBootNextServer returns a boolean if a field has been set.

### GetBootFileName

`func (o *InlineObject41) GetBootFileName() string`

GetBootFileName returns the BootFileName field if non-nil, zero value otherwise.

### GetBootFileNameOk

`func (o *InlineObject41) GetBootFileNameOk() (*string, bool)`

GetBootFileNameOk returns a tuple with the BootFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFileName

`func (o *InlineObject41) SetBootFileName(v string)`

SetBootFileName sets BootFileName field to given value.

### HasBootFileName

`func (o *InlineObject41) HasBootFileName() bool`

HasBootFileName returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *InlineObject41) GetDhcpOptions() []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *InlineObject41) GetDhcpOptionsOk() (*[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *InlineObject41) SetDhcpOptions(v []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *InlineObject41) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetReservedIpRanges

`func (o *InlineObject41) GetReservedIpRanges() []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *InlineObject41) GetReservedIpRangesOk() (*[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *InlineObject41) SetReservedIpRanges(v []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *InlineObject41) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *InlineObject41) GetFixedIpAssignments() []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *InlineObject41) GetFixedIpAssignmentsOk() (*[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *InlineObject41) SetFixedIpAssignments(v []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments)`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *InlineObject41) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.

### GetDhcpDefaultRouterIps

`func (o *InlineObject41) GetDhcpDefaultRouterIps() string`

GetDhcpDefaultRouterIps returns the DhcpDefaultRouterIps field if non-nil, zero value otherwise.

### GetDhcpDefaultRouterIpsOk

`func (o *InlineObject41) GetDhcpDefaultRouterIpsOk() (*string, bool)`

GetDhcpDefaultRouterIpsOk returns a tuple with the DhcpDefaultRouterIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpDefaultRouterIps

`func (o *InlineObject41) SetDhcpDefaultRouterIps(v string)`

SetDhcpDefaultRouterIps sets DhcpDefaultRouterIps field to given value.

### HasDhcpDefaultRouterIps

`func (o *InlineObject41) HasDhcpDefaultRouterIps() bool`

HasDhcpDefaultRouterIps returns a boolean if a field has been set.

### GetDhcpDomainName

`func (o *InlineObject41) GetDhcpDomainName() string`

GetDhcpDomainName returns the DhcpDomainName field if non-nil, zero value otherwise.

### GetDhcpDomainNameOk

`func (o *InlineObject41) GetDhcpDomainNameOk() (*string, bool)`

GetDhcpDomainNameOk returns a tuple with the DhcpDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpDomainName

`func (o *InlineObject41) SetDhcpDomainName(v string)`

SetDhcpDomainName sets DhcpDomainName field to given value.

### HasDhcpDomainName

`func (o *InlineObject41) HasDhcpDomainName() bool`

HasDhcpDomainName returns a boolean if a field has been set.

### SetDhcpDomainNameNil

`func (o *InlineObject41) SetDhcpDomainNameNil(b bool)`

 SetDhcpDomainNameNil sets the value for DhcpDomainName to be an explicit nil

### UnsetDhcpDomainName
`func (o *InlineObject41) UnsetDhcpDomainName()`

UnsetDhcpDomainName ensures that no value is present for DhcpDomainName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


