# InlineResponse20040

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpMode** | Pointer to **string** | The DHCP mode options for the switch stack interface (&#39;dhcpDisabled&#39;, &#39;dhcpRelay&#39; or &#39;dhcpServer&#39;) | [optional] 
**DhcpRelayServerIps** | Pointer to **[]string** | The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface | [optional] 
**DhcpLeaseTime** | Pointer to **string** | The DHCP lease time config for the dhcp server running on the switch stack interface (&#39;30 minutes&#39;, &#39;1 hour&#39;, &#39;4 hours&#39;, &#39;12 hours&#39;, &#39;1 day&#39; or &#39;1 week&#39;) | [optional] 
**DnsNameserversOption** | Pointer to **string** | The DHCP name server option for the dhcp server running on the switch stack interface (&#39;googlePublicDns&#39;, &#39;openDns&#39; or &#39;custom&#39;) | [optional] 
**DnsCustomNameservers** | Pointer to **[]string** | The DHCP name server IPs when DHCP name server option is &#39;custom&#39; | [optional] 
**BootOptionsEnabled** | Pointer to **bool** | Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface | [optional] 
**BootNextServer** | Pointer to **string** | The PXE boot server IP for the DHCP server running on the switch stack interface | [optional] 
**BootFileName** | Pointer to **string** | The PXE boot server file name for the DHCP server running on the switch stack interface | [optional] 
**DhcpOptions** | Pointer to [**[]InlineResponse20040DhcpOptions**](InlineResponse20040DhcpOptions.md) | Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface | [optional] 
**ReservedIpRanges** | Pointer to [**[]InlineResponse20040ReservedIpRanges**](InlineResponse20040ReservedIpRanges.md) | Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface | [optional] 
**FixedIpAssignments** | Pointer to [**[]InlineResponse20040FixedIpAssignments**](InlineResponse20040FixedIpAssignments.md) | Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface | [optional] 

## Methods

### NewInlineResponse20040

`func NewInlineResponse20040() *InlineResponse20040`

NewInlineResponse20040 instantiates a new InlineResponse20040 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20040WithDefaults

`func NewInlineResponse20040WithDefaults() *InlineResponse20040`

NewInlineResponse20040WithDefaults instantiates a new InlineResponse20040 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpMode

`func (o *InlineResponse20040) GetDhcpMode() string`

GetDhcpMode returns the DhcpMode field if non-nil, zero value otherwise.

### GetDhcpModeOk

`func (o *InlineResponse20040) GetDhcpModeOk() (*string, bool)`

GetDhcpModeOk returns a tuple with the DhcpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpMode

`func (o *InlineResponse20040) SetDhcpMode(v string)`

SetDhcpMode sets DhcpMode field to given value.

### HasDhcpMode

`func (o *InlineResponse20040) HasDhcpMode() bool`

HasDhcpMode returns a boolean if a field has been set.

### GetDhcpRelayServerIps

`func (o *InlineResponse20040) GetDhcpRelayServerIps() []string`

GetDhcpRelayServerIps returns the DhcpRelayServerIps field if non-nil, zero value otherwise.

### GetDhcpRelayServerIpsOk

`func (o *InlineResponse20040) GetDhcpRelayServerIpsOk() (*[]string, bool)`

GetDhcpRelayServerIpsOk returns a tuple with the DhcpRelayServerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelayServerIps

`func (o *InlineResponse20040) SetDhcpRelayServerIps(v []string)`

SetDhcpRelayServerIps sets DhcpRelayServerIps field to given value.

### HasDhcpRelayServerIps

`func (o *InlineResponse20040) HasDhcpRelayServerIps() bool`

HasDhcpRelayServerIps returns a boolean if a field has been set.

### GetDhcpLeaseTime

`func (o *InlineResponse20040) GetDhcpLeaseTime() string`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *InlineResponse20040) GetDhcpLeaseTimeOk() (*string, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *InlineResponse20040) SetDhcpLeaseTime(v string)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *InlineResponse20040) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetDnsNameserversOption

`func (o *InlineResponse20040) GetDnsNameserversOption() string`

GetDnsNameserversOption returns the DnsNameserversOption field if non-nil, zero value otherwise.

### GetDnsNameserversOptionOk

`func (o *InlineResponse20040) GetDnsNameserversOptionOk() (*string, bool)`

GetDnsNameserversOptionOk returns a tuple with the DnsNameserversOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameserversOption

`func (o *InlineResponse20040) SetDnsNameserversOption(v string)`

SetDnsNameserversOption sets DnsNameserversOption field to given value.

### HasDnsNameserversOption

`func (o *InlineResponse20040) HasDnsNameserversOption() bool`

HasDnsNameserversOption returns a boolean if a field has been set.

### GetDnsCustomNameservers

`func (o *InlineResponse20040) GetDnsCustomNameservers() []string`

GetDnsCustomNameservers returns the DnsCustomNameservers field if non-nil, zero value otherwise.

### GetDnsCustomNameserversOk

`func (o *InlineResponse20040) GetDnsCustomNameserversOk() (*[]string, bool)`

GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCustomNameservers

`func (o *InlineResponse20040) SetDnsCustomNameservers(v []string)`

SetDnsCustomNameservers sets DnsCustomNameservers field to given value.

### HasDnsCustomNameservers

`func (o *InlineResponse20040) HasDnsCustomNameservers() bool`

HasDnsCustomNameservers returns a boolean if a field has been set.

### GetBootOptionsEnabled

`func (o *InlineResponse20040) GetBootOptionsEnabled() bool`

GetBootOptionsEnabled returns the BootOptionsEnabled field if non-nil, zero value otherwise.

### GetBootOptionsEnabledOk

`func (o *InlineResponse20040) GetBootOptionsEnabledOk() (*bool, bool)`

GetBootOptionsEnabledOk returns a tuple with the BootOptionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOptionsEnabled

`func (o *InlineResponse20040) SetBootOptionsEnabled(v bool)`

SetBootOptionsEnabled sets BootOptionsEnabled field to given value.

### HasBootOptionsEnabled

`func (o *InlineResponse20040) HasBootOptionsEnabled() bool`

HasBootOptionsEnabled returns a boolean if a field has been set.

### GetBootNextServer

`func (o *InlineResponse20040) GetBootNextServer() string`

GetBootNextServer returns the BootNextServer field if non-nil, zero value otherwise.

### GetBootNextServerOk

`func (o *InlineResponse20040) GetBootNextServerOk() (*string, bool)`

GetBootNextServerOk returns a tuple with the BootNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootNextServer

`func (o *InlineResponse20040) SetBootNextServer(v string)`

SetBootNextServer sets BootNextServer field to given value.

### HasBootNextServer

`func (o *InlineResponse20040) HasBootNextServer() bool`

HasBootNextServer returns a boolean if a field has been set.

### GetBootFileName

`func (o *InlineResponse20040) GetBootFileName() string`

GetBootFileName returns the BootFileName field if non-nil, zero value otherwise.

### GetBootFileNameOk

`func (o *InlineResponse20040) GetBootFileNameOk() (*string, bool)`

GetBootFileNameOk returns a tuple with the BootFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFileName

`func (o *InlineResponse20040) SetBootFileName(v string)`

SetBootFileName sets BootFileName field to given value.

### HasBootFileName

`func (o *InlineResponse20040) HasBootFileName() bool`

HasBootFileName returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *InlineResponse20040) GetDhcpOptions() []InlineResponse20040DhcpOptions`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *InlineResponse20040) GetDhcpOptionsOk() (*[]InlineResponse20040DhcpOptions, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *InlineResponse20040) SetDhcpOptions(v []InlineResponse20040DhcpOptions)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *InlineResponse20040) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetReservedIpRanges

`func (o *InlineResponse20040) GetReservedIpRanges() []InlineResponse20040ReservedIpRanges`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *InlineResponse20040) GetReservedIpRangesOk() (*[]InlineResponse20040ReservedIpRanges, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *InlineResponse20040) SetReservedIpRanges(v []InlineResponse20040ReservedIpRanges)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *InlineResponse20040) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *InlineResponse20040) GetFixedIpAssignments() []InlineResponse20040FixedIpAssignments`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *InlineResponse20040) GetFixedIpAssignmentsOk() (*[]InlineResponse20040FixedIpAssignments, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *InlineResponse20040) SetFixedIpAssignments(v []InlineResponse20040FixedIpAssignments)`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *InlineResponse20040) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


