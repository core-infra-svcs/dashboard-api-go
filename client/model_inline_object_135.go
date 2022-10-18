/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject135 struct for InlineObject135
type InlineObject135 struct {
	// The DHCP mode options for the switch stack interface ('dhcpDisabled', 'dhcpRelay' or 'dhcpServer')
	DhcpMode *string `json:"dhcpMode,omitempty"`
	// The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface
	DhcpRelayServerIps *[]string `json:"dhcpRelayServerIps,omitempty"`
	// The DHCP lease time config for the dhcp server running on switch stack interface ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')
	DhcpLeaseTime *string `json:"dhcpLeaseTime,omitempty"`
	// The DHCP name server option for the dhcp server running on the switch stack interface ('googlePublicDns', 'openDns' or 'custom')
	DnsNameserversOption *string `json:"dnsNameserversOption,omitempty"`
	// The DHCP name server IPs when DHCP name server option is 'custom'
	DnsCustomNameservers *[]string `json:"dnsCustomNameservers,omitempty"`
	// Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface
	BootOptionsEnabled *bool `json:"bootOptionsEnabled,omitempty"`
	// The PXE boot server IP for the DHCP server running on the switch stack interface
	BootNextServer *string `json:"bootNextServer,omitempty"`
	// The PXE boot server file name for the DHCP server running on the switch stack interface
	BootFileName *string `json:"bootFileName,omitempty"`
	// Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface
	DhcpOptions *[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions `json:"dhcpOptions,omitempty"`
	// Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
	ReservedIpRanges *[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges `json:"reservedIpRanges,omitempty"`
	// Array of DHCP fixed IP assignments for the DHCP server running on the switch stack interface
	FixedIpAssignments *[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments `json:"fixedIpAssignments,omitempty"`
}

// NewInlineObject135 instantiates a new InlineObject135 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject135() *InlineObject135 {
	this := InlineObject135{}
	return &this
}

// NewInlineObject135WithDefaults instantiates a new InlineObject135 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject135WithDefaults() *InlineObject135 {
	this := InlineObject135{}
	return &this
}

// GetDhcpMode returns the DhcpMode field value if set, zero value otherwise.
func (o *InlineObject135) GetDhcpMode() string {
	if o == nil || o.DhcpMode == nil {
		var ret string
		return ret
	}
	return *o.DhcpMode
}

// GetDhcpModeOk returns a tuple with the DhcpMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject135) GetDhcpModeOk() (*string, bool) {
	if o == nil || o.DhcpMode == nil {
		return nil, false
	}
	return o.DhcpMode, true
}

// HasDhcpMode returns a boolean if a field has been set.
func (o *InlineObject135) HasDhcpMode() bool {
	if o != nil && o.DhcpMode != nil {
		return true
	}

	return false
}

// SetDhcpMode gets a reference to the given string and assigns it to the DhcpMode field.
func (o *InlineObject135) SetDhcpMode(v string) {
	o.DhcpMode = &v
}

// GetDhcpRelayServerIps returns the DhcpRelayServerIps field value if set, zero value otherwise.
func (o *InlineObject135) GetDhcpRelayServerIps() []string {
	if o == nil || o.DhcpRelayServerIps == nil {
		var ret []string
		return ret
	}
	return *o.DhcpRelayServerIps
}

// GetDhcpRelayServerIpsOk returns a tuple with the DhcpRelayServerIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject135) GetDhcpRelayServerIpsOk() (*[]string, bool) {
	if o == nil || o.DhcpRelayServerIps == nil {
		return nil, false
	}
	return o.DhcpRelayServerIps, true
}

// HasDhcpRelayServerIps returns a boolean if a field has been set.
func (o *InlineObject135) HasDhcpRelayServerIps() bool {
	if o != nil && o.DhcpRelayServerIps != nil {
		return true
	}

	return false
}

// SetDhcpRelayServerIps gets a reference to the given []string and assigns it to the DhcpRelayServerIps field.
func (o *InlineObject135) SetDhcpRelayServerIps(v []string) {
	o.DhcpRelayServerIps = &v
}

// GetDhcpLeaseTime returns the DhcpLeaseTime field value if set, zero value otherwise.
func (o *InlineObject135) GetDhcpLeaseTime() string {
	if o == nil || o.DhcpLeaseTime == nil {
		var ret string
		return ret
	}
	return *o.DhcpLeaseTime
}

// GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject135) GetDhcpLeaseTimeOk() (*string, bool) {
	if o == nil || o.DhcpLeaseTime == nil {
		return nil, false
	}
	return o.DhcpLeaseTime, true
}

// HasDhcpLeaseTime returns a boolean if a field has been set.
func (o *InlineObject135) HasDhcpLeaseTime() bool {
	if o != nil && o.DhcpLeaseTime != nil {
		return true
	}

	return false
}

// SetDhcpLeaseTime gets a reference to the given string and assigns it to the DhcpLeaseTime field.
func (o *InlineObject135) SetDhcpLeaseTime(v string) {
	o.DhcpLeaseTime = &v
}

// GetDnsNameserversOption returns the DnsNameserversOption field value if set, zero value otherwise.
func (o *InlineObject135) GetDnsNameserversOption() string {
	if o == nil || o.DnsNameserversOption == nil {
		var ret string
		return ret
	}
	return *o.DnsNameserversOption
}

// GetDnsNameserversOptionOk returns a tuple with the DnsNameserversOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject135) GetDnsNameserversOptionOk() (*string, bool) {
	if o == nil || o.DnsNameserversOption == nil {
		return nil, false
	}
	return o.DnsNameserversOption, true
}

// HasDnsNameserversOption returns a boolean if a field has been set.
func (o *InlineObject135) HasDnsNameserversOption() bool {
	if o != nil && o.DnsNameserversOption != nil {
		return true
	}

	return false
}

// SetDnsNameserversOption gets a reference to the given string and assigns it to the DnsNameserversOption field.
func (o *InlineObject135) SetDnsNameserversOption(v string) {
	o.DnsNameserversOption = &v
}

// GetDnsCustomNameservers returns the DnsCustomNameservers field value if set, zero value otherwise.
func (o *InlineObject135) GetDnsCustomNameservers() []string {
	if o == nil || o.DnsCustomNameservers == nil {
		var ret []string
		return ret
	}
	return *o.DnsCustomNameservers
}

// GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject135) GetDnsCustomNameserversOk() (*[]string, bool) {
	if o == nil || o.DnsCustomNameservers == nil {
		return nil, false
	}
	return o.DnsCustomNameservers, true
}

// HasDnsCustomNameservers returns a boolean if a field has been set.
func (o *InlineObject135) HasDnsCustomNameservers() bool {
	if o != nil && o.DnsCustomNameservers != nil {
		return true
	}

	return false
}

// SetDnsCustomNameservers gets a reference to the given []string and assigns it to the DnsCustomNameservers field.
func (o *InlineObject135) SetDnsCustomNameservers(v []string) {
	o.DnsCustomNameservers = &v
}

// GetBootOptionsEnabled returns the BootOptionsEnabled field value if set, zero value otherwise.
func (o *InlineObject135) GetBootOptionsEnabled() bool {
	if o == nil || o.BootOptionsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BootOptionsEnabled
}

// GetBootOptionsEnabledOk returns a tuple with the BootOptionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject135) GetBootOptionsEnabledOk() (*bool, bool) {
	if o == nil || o.BootOptionsEnabled == nil {
		return nil, false
	}
	return o.BootOptionsEnabled, true
}

// HasBootOptionsEnabled returns a boolean if a field has been set.
func (o *InlineObject135) HasBootOptionsEnabled() bool {
	if o != nil && o.BootOptionsEnabled != nil {
		return true
	}

	return false
}

// SetBootOptionsEnabled gets a reference to the given bool and assigns it to the BootOptionsEnabled field.
func (o *InlineObject135) SetBootOptionsEnabled(v bool) {
	o.BootOptionsEnabled = &v
}

// GetBootNextServer returns the BootNextServer field value if set, zero value otherwise.
func (o *InlineObject135) GetBootNextServer() string {
	if o == nil || o.BootNextServer == nil {
		var ret string
		return ret
	}
	return *o.BootNextServer
}

// GetBootNextServerOk returns a tuple with the BootNextServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject135) GetBootNextServerOk() (*string, bool) {
	if o == nil || o.BootNextServer == nil {
		return nil, false
	}
	return o.BootNextServer, true
}

// HasBootNextServer returns a boolean if a field has been set.
func (o *InlineObject135) HasBootNextServer() bool {
	if o != nil && o.BootNextServer != nil {
		return true
	}

	return false
}

// SetBootNextServer gets a reference to the given string and assigns it to the BootNextServer field.
func (o *InlineObject135) SetBootNextServer(v string) {
	o.BootNextServer = &v
}

// GetBootFileName returns the BootFileName field value if set, zero value otherwise.
func (o *InlineObject135) GetBootFileName() string {
	if o == nil || o.BootFileName == nil {
		var ret string
		return ret
	}
	return *o.BootFileName
}

// GetBootFileNameOk returns a tuple with the BootFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject135) GetBootFileNameOk() (*string, bool) {
	if o == nil || o.BootFileName == nil {
		return nil, false
	}
	return o.BootFileName, true
}

// HasBootFileName returns a boolean if a field has been set.
func (o *InlineObject135) HasBootFileName() bool {
	if o != nil && o.BootFileName != nil {
		return true
	}

	return false
}

// SetBootFileName gets a reference to the given string and assigns it to the BootFileName field.
func (o *InlineObject135) SetBootFileName(v string) {
	o.BootFileName = &v
}

// GetDhcpOptions returns the DhcpOptions field value if set, zero value otherwise.
func (o *InlineObject135) GetDhcpOptions() []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions {
	if o == nil || o.DhcpOptions == nil {
		var ret []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions
		return ret
	}
	return *o.DhcpOptions
}

// GetDhcpOptionsOk returns a tuple with the DhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject135) GetDhcpOptionsOk() (*[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions, bool) {
	if o == nil || o.DhcpOptions == nil {
		return nil, false
	}
	return o.DhcpOptions, true
}

// HasDhcpOptions returns a boolean if a field has been set.
func (o *InlineObject135) HasDhcpOptions() bool {
	if o != nil && o.DhcpOptions != nil {
		return true
	}

	return false
}

// SetDhcpOptions gets a reference to the given []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions and assigns it to the DhcpOptions field.
func (o *InlineObject135) SetDhcpOptions(v []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) {
	o.DhcpOptions = &v
}

// GetReservedIpRanges returns the ReservedIpRanges field value if set, zero value otherwise.
func (o *InlineObject135) GetReservedIpRanges() []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges {
	if o == nil || o.ReservedIpRanges == nil {
		var ret []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges
		return ret
	}
	return *o.ReservedIpRanges
}

// GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject135) GetReservedIpRangesOk() (*[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges, bool) {
	if o == nil || o.ReservedIpRanges == nil {
		return nil, false
	}
	return o.ReservedIpRanges, true
}

// HasReservedIpRanges returns a boolean if a field has been set.
func (o *InlineObject135) HasReservedIpRanges() bool {
	if o != nil && o.ReservedIpRanges != nil {
		return true
	}

	return false
}

// SetReservedIpRanges gets a reference to the given []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges and assigns it to the ReservedIpRanges field.
func (o *InlineObject135) SetReservedIpRanges(v []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) {
	o.ReservedIpRanges = &v
}

// GetFixedIpAssignments returns the FixedIpAssignments field value if set, zero value otherwise.
func (o *InlineObject135) GetFixedIpAssignments() []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments {
	if o == nil || o.FixedIpAssignments == nil {
		var ret []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments
		return ret
	}
	return *o.FixedIpAssignments
}

// GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject135) GetFixedIpAssignmentsOk() (*[]DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments, bool) {
	if o == nil || o.FixedIpAssignments == nil {
		return nil, false
	}
	return o.FixedIpAssignments, true
}

// HasFixedIpAssignments returns a boolean if a field has been set.
func (o *InlineObject135) HasFixedIpAssignments() bool {
	if o != nil && o.FixedIpAssignments != nil {
		return true
	}

	return false
}

// SetFixedIpAssignments gets a reference to the given []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments and assigns it to the FixedIpAssignments field.
func (o *InlineObject135) SetFixedIpAssignments(v []DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) {
	o.FixedIpAssignments = &v
}

func (o InlineObject135) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DhcpMode != nil {
		toSerialize["dhcpMode"] = o.DhcpMode
	}
	if o.DhcpRelayServerIps != nil {
		toSerialize["dhcpRelayServerIps"] = o.DhcpRelayServerIps
	}
	if o.DhcpLeaseTime != nil {
		toSerialize["dhcpLeaseTime"] = o.DhcpLeaseTime
	}
	if o.DnsNameserversOption != nil {
		toSerialize["dnsNameserversOption"] = o.DnsNameserversOption
	}
	if o.DnsCustomNameservers != nil {
		toSerialize["dnsCustomNameservers"] = o.DnsCustomNameservers
	}
	if o.BootOptionsEnabled != nil {
		toSerialize["bootOptionsEnabled"] = o.BootOptionsEnabled
	}
	if o.BootNextServer != nil {
		toSerialize["bootNextServer"] = o.BootNextServer
	}
	if o.BootFileName != nil {
		toSerialize["bootFileName"] = o.BootFileName
	}
	if o.DhcpOptions != nil {
		toSerialize["dhcpOptions"] = o.DhcpOptions
	}
	if o.ReservedIpRanges != nil {
		toSerialize["reservedIpRanges"] = o.ReservedIpRanges
	}
	if o.FixedIpAssignments != nil {
		toSerialize["fixedIpAssignments"] = o.FixedIpAssignments
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject135 struct {
	value *InlineObject135
	isSet bool
}

func (v NullableInlineObject135) Get() *InlineObject135 {
	return v.value
}

func (v *NullableInlineObject135) Set(val *InlineObject135) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject135) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject135) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject135(val *InlineObject135) *NullableInlineObject135 {
	return &NullableInlineObject135{value: val, isSet: true}
}

func (v NullableInlineObject135) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject135) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


