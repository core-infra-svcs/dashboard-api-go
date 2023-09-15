# InlineObject140

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpMode** | Pointer to **string** | The DHCP mode options for the switch stack interface         (&#39;dhcpDisabled&#39;, &#39;dhcpRelay&#39; or &#39;dhcpServer&#39;) | [optional] 
**DhcpRelayServerIps** | Pointer to **[]string** | The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface | [optional] 
**DhcpLeaseTime** | Pointer to **string** | The DHCP lease time config for the dhcp server running on switch stack interface         (&#39;30 minutes&#39;, &#39;1 hour&#39;, &#39;4 hours&#39;, &#39;12 hours&#39;, &#39;1 day&#39; or &#39;1 week&#39;) | [optional] 
**DnsNameserversOption** | Pointer to **string** | The DHCP name server option for the dhcp server running on the switch stack interface         (&#39;googlePublicDns&#39;, &#39;openDns&#39; or &#39;custom&#39;) | [optional] 
**DnsCustomNameservers** | Pointer to **[]string** | The DHCP name server IPs when DHCP name server option is &#39;         custom&#39; | [optional] 
**BootOptionsEnabled** | Pointer to **bool** | Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch         stack interface | [optional] 
**BootNextServer** | Pointer to **string** | The PXE boot server IP for the DHCP server running on the switch stack interface | [optional] 
**BootFileName** | Pointer to **string** | The PXE boot server file name for the DHCP server running on the switch stack interface | [optional] 
**DhcpOptions** | Pointer to [**[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions**](DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions.md) | Array of DHCP options consisting of code, type and value for the DHCP server running on the         switch stack interface | [optional] 
**ReservedIpRanges** | Pointer to [**[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges**](DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges.md) | Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface | [optional] 
**FixedIpAssignments** | Pointer to [**[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments**](DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments.md) | Array of DHCP fixed IP assignments for the DHCP server running on the switch stack interface | [optional] 

## Methods

### NewInlineObject140

`func NewInlineObject140() *InlineObject140`

NewInlineObject140 instantiates a new InlineObject140 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject140WithDefaults

`func NewInlineObject140WithDefaults() *InlineObject140`

NewInlineObject140WithDefaults instantiates a new InlineObject140 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpMode

`func (o *InlineObject140) GetDhcpMode() string`

GetDhcpMode returns the DhcpMode field if non-nil, zero value otherwise.

### GetDhcpModeOk

`func (o *InlineObject140) GetDhcpModeOk() (*string, bool)`

GetDhcpModeOk returns a tuple with the DhcpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpMode

`func (o *InlineObject140) SetDhcpMode(v string)`

SetDhcpMode sets DhcpMode field to given value.

### HasDhcpMode

`func (o *InlineObject140) HasDhcpMode() bool`

HasDhcpMode returns a boolean if a field has been set.

### GetDhcpRelayServerIps

`func (o *InlineObject140) GetDhcpRelayServerIps() []string`

GetDhcpRelayServerIps returns the DhcpRelayServerIps field if non-nil, zero value otherwise.

### GetDhcpRelayServerIpsOk

`func (o *InlineObject140) GetDhcpRelayServerIpsOk() (*[]string, bool)`

GetDhcpRelayServerIpsOk returns a tuple with the DhcpRelayServerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelayServerIps

`func (o *InlineObject140) SetDhcpRelayServerIps(v []string)`

SetDhcpRelayServerIps sets DhcpRelayServerIps field to given value.

### HasDhcpRelayServerIps

`func (o *InlineObject140) HasDhcpRelayServerIps() bool`

HasDhcpRelayServerIps returns a boolean if a field has been set.

### GetDhcpLeaseTime

`func (o *InlineObject140) GetDhcpLeaseTime() string`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *InlineObject140) GetDhcpLeaseTimeOk() (*string, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *InlineObject140) SetDhcpLeaseTime(v string)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *InlineObject140) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetDnsNameserversOption

`func (o *InlineObject140) GetDnsNameserversOption() string`

GetDnsNameserversOption returns the DnsNameserversOption field if non-nil, zero value otherwise.

### GetDnsNameserversOptionOk

`func (o *InlineObject140) GetDnsNameserversOptionOk() (*string, bool)`

GetDnsNameserversOptionOk returns a tuple with the DnsNameserversOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameserversOption

`func (o *InlineObject140) SetDnsNameserversOption(v string)`

SetDnsNameserversOption sets DnsNameserversOption field to given value.

### HasDnsNameserversOption

`func (o *InlineObject140) HasDnsNameserversOption() bool`

HasDnsNameserversOption returns a boolean if a field has been set.

### GetDnsCustomNameservers

`func (o *InlineObject140) GetDnsCustomNameservers() []string`

GetDnsCustomNameservers returns the DnsCustomNameservers field if non-nil, zero value otherwise.

### GetDnsCustomNameserversOk

`func (o *InlineObject140) GetDnsCustomNameserversOk() (*[]string, bool)`

GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCustomNameservers

`func (o *InlineObject140) SetDnsCustomNameservers(v []string)`

SetDnsCustomNameservers sets DnsCustomNameservers field to given value.

### HasDnsCustomNameservers

`func (o *InlineObject140) HasDnsCustomNameservers() bool`

HasDnsCustomNameservers returns a boolean if a field has been set.

### GetBootOptionsEnabled

`func (o *InlineObject140) GetBootOptionsEnabled() bool`

GetBootOptionsEnabled returns the BootOptionsEnabled field if non-nil, zero value otherwise.

### GetBootOptionsEnabledOk

`func (o *InlineObject140) GetBootOptionsEnabledOk() (*bool, bool)`

GetBootOptionsEnabledOk returns a tuple with the BootOptionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOptionsEnabled

`func (o *InlineObject140) SetBootOptionsEnabled(v bool)`

SetBootOptionsEnabled sets BootOptionsEnabled field to given value.

### HasBootOptionsEnabled

`func (o *InlineObject140) HasBootOptionsEnabled() bool`

HasBootOptionsEnabled returns a boolean if a field has been set.

### GetBootNextServer

`func (o *InlineObject140) GetBootNextServer() string`

GetBootNextServer returns the BootNextServer field if non-nil, zero value otherwise.

### GetBootNextServerOk

`func (o *InlineObject140) GetBootNextServerOk() (*string, bool)`

GetBootNextServerOk returns a tuple with the BootNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootNextServer

`func (o *InlineObject140) SetBootNextServer(v string)`

SetBootNextServer sets BootNextServer field to given value.

### HasBootNextServer

`func (o *InlineObject140) HasBootNextServer() bool`

HasBootNextServer returns a boolean if a field has been set.

### GetBootFileName

`func (o *InlineObject140) GetBootFileName() string`

GetBootFileName returns the BootFileName field if non-nil, zero value otherwise.

### GetBootFileNameOk

`func (o *InlineObject140) GetBootFileNameOk() (*string, bool)`

GetBootFileNameOk returns a tuple with the BootFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFileName

`func (o *InlineObject140) SetBootFileName(v string)`

SetBootFileName sets BootFileName field to given value.

### HasBootFileName

`func (o *InlineObject140) HasBootFileName() bool`

HasBootFileName returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *InlineObject140) GetDhcpOptions() []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *InlineObject140) GetDhcpOptionsOk() (*[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *InlineObject140) SetDhcpOptions(v []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *InlineObject140) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetReservedIpRanges

`func (o *InlineObject140) GetReservedIpRanges() []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *InlineObject140) GetReservedIpRangesOk() (*[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *InlineObject140) SetReservedIpRanges(v []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *InlineObject140) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *InlineObject140) GetFixedIpAssignments() []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *InlineObject140) GetFixedIpAssignmentsOk() (*[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *InlineObject140) SetFixedIpAssignments(v []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments)`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *InlineObject140) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


