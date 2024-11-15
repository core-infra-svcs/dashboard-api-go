/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20036 struct for InlineResponse20036
type InlineResponse20036 struct {
	// The DHCP mode options for the switch stack interface ('dhcpDisabled', 'dhcpRelay' or 'dhcpServer')
	DhcpMode *string `json:"dhcpMode,omitempty"`
	// The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface
	DhcpRelayServerIps []string `json:"dhcpRelayServerIps,omitempty"`
	// The DHCP lease time config for the dhcp server running on the switch stack interface ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')
	DhcpLeaseTime *string `json:"dhcpLeaseTime,omitempty"`
	// The DHCP name server option for the dhcp server running on the switch stack interface ('googlePublicDns', 'openDns' or 'custom')
	DnsNameserversOption *string `json:"dnsNameserversOption,omitempty"`
	// The DHCP name server IPs when DHCP name server option is 'custom'
	DnsCustomNameservers []string `json:"dnsCustomNameservers,omitempty"`
	// Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface
	BootOptionsEnabled *bool `json:"bootOptionsEnabled,omitempty"`
	// The PXE boot server IP for the DHCP server running on the switch stack interface
	BootNextServer *string `json:"bootNextServer,omitempty"`
	// The PXE boot server file name for the DHCP server running on the switch stack interface
	BootFileName *string `json:"bootFileName,omitempty"`
	// Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface
	DhcpOptions []InlineResponse20036DhcpOptions `json:"dhcpOptions,omitempty"`
	// Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
	ReservedIpRanges []InlineResponse20036ReservedIpRanges `json:"reservedIpRanges,omitempty"`
	// Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
	FixedIpAssignments []InlineResponse20036FixedIpAssignments `json:"fixedIpAssignments,omitempty"`
}

// NewInlineResponse20036 instantiates a new InlineResponse20036 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20036() *InlineResponse20036 {
	this := InlineResponse20036{}
	return &this
}

// NewInlineResponse20036WithDefaults instantiates a new InlineResponse20036 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20036WithDefaults() *InlineResponse20036 {
	this := InlineResponse20036{}
	return &this
}

// GetDhcpMode returns the DhcpMode field value if set, zero value otherwise.
func (o *InlineResponse20036) GetDhcpMode() string {
	if o == nil || isNil(o.DhcpMode) {
		var ret string
		return ret
	}
	return *o.DhcpMode
}

// GetDhcpModeOk returns a tuple with the DhcpMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetDhcpModeOk() (*string, bool) {
	if o == nil || isNil(o.DhcpMode) {
    return nil, false
	}
	return o.DhcpMode, true
}

// HasDhcpMode returns a boolean if a field has been set.
func (o *InlineResponse20036) HasDhcpMode() bool {
	if o != nil && !isNil(o.DhcpMode) {
		return true
	}

	return false
}

// SetDhcpMode gets a reference to the given string and assigns it to the DhcpMode field.
func (o *InlineResponse20036) SetDhcpMode(v string) {
	o.DhcpMode = &v
}

// GetDhcpRelayServerIps returns the DhcpRelayServerIps field value if set, zero value otherwise.
func (o *InlineResponse20036) GetDhcpRelayServerIps() []string {
	if o == nil || isNil(o.DhcpRelayServerIps) {
		var ret []string
		return ret
	}
	return o.DhcpRelayServerIps
}

// GetDhcpRelayServerIpsOk returns a tuple with the DhcpRelayServerIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetDhcpRelayServerIpsOk() ([]string, bool) {
	if o == nil || isNil(o.DhcpRelayServerIps) {
    return nil, false
	}
	return o.DhcpRelayServerIps, true
}

// HasDhcpRelayServerIps returns a boolean if a field has been set.
func (o *InlineResponse20036) HasDhcpRelayServerIps() bool {
	if o != nil && !isNil(o.DhcpRelayServerIps) {
		return true
	}

	return false
}

// SetDhcpRelayServerIps gets a reference to the given []string and assigns it to the DhcpRelayServerIps field.
func (o *InlineResponse20036) SetDhcpRelayServerIps(v []string) {
	o.DhcpRelayServerIps = v
}

// GetDhcpLeaseTime returns the DhcpLeaseTime field value if set, zero value otherwise.
func (o *InlineResponse20036) GetDhcpLeaseTime() string {
	if o == nil || isNil(o.DhcpLeaseTime) {
		var ret string
		return ret
	}
	return *o.DhcpLeaseTime
}

// GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetDhcpLeaseTimeOk() (*string, bool) {
	if o == nil || isNil(o.DhcpLeaseTime) {
    return nil, false
	}
	return o.DhcpLeaseTime, true
}

// HasDhcpLeaseTime returns a boolean if a field has been set.
func (o *InlineResponse20036) HasDhcpLeaseTime() bool {
	if o != nil && !isNil(o.DhcpLeaseTime) {
		return true
	}

	return false
}

// SetDhcpLeaseTime gets a reference to the given string and assigns it to the DhcpLeaseTime field.
func (o *InlineResponse20036) SetDhcpLeaseTime(v string) {
	o.DhcpLeaseTime = &v
}

// GetDnsNameserversOption returns the DnsNameserversOption field value if set, zero value otherwise.
func (o *InlineResponse20036) GetDnsNameserversOption() string {
	if o == nil || isNil(o.DnsNameserversOption) {
		var ret string
		return ret
	}
	return *o.DnsNameserversOption
}

// GetDnsNameserversOptionOk returns a tuple with the DnsNameserversOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetDnsNameserversOptionOk() (*string, bool) {
	if o == nil || isNil(o.DnsNameserversOption) {
    return nil, false
	}
	return o.DnsNameserversOption, true
}

// HasDnsNameserversOption returns a boolean if a field has been set.
func (o *InlineResponse20036) HasDnsNameserversOption() bool {
	if o != nil && !isNil(o.DnsNameserversOption) {
		return true
	}

	return false
}

// SetDnsNameserversOption gets a reference to the given string and assigns it to the DnsNameserversOption field.
func (o *InlineResponse20036) SetDnsNameserversOption(v string) {
	o.DnsNameserversOption = &v
}

// GetDnsCustomNameservers returns the DnsCustomNameservers field value if set, zero value otherwise.
func (o *InlineResponse20036) GetDnsCustomNameservers() []string {
	if o == nil || isNil(o.DnsCustomNameservers) {
		var ret []string
		return ret
	}
	return o.DnsCustomNameservers
}

// GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetDnsCustomNameserversOk() ([]string, bool) {
	if o == nil || isNil(o.DnsCustomNameservers) {
    return nil, false
	}
	return o.DnsCustomNameservers, true
}

// HasDnsCustomNameservers returns a boolean if a field has been set.
func (o *InlineResponse20036) HasDnsCustomNameservers() bool {
	if o != nil && !isNil(o.DnsCustomNameservers) {
		return true
	}

	return false
}

// SetDnsCustomNameservers gets a reference to the given []string and assigns it to the DnsCustomNameservers field.
func (o *InlineResponse20036) SetDnsCustomNameservers(v []string) {
	o.DnsCustomNameservers = v
}

// GetBootOptionsEnabled returns the BootOptionsEnabled field value if set, zero value otherwise.
func (o *InlineResponse20036) GetBootOptionsEnabled() bool {
	if o == nil || isNil(o.BootOptionsEnabled) {
		var ret bool
		return ret
	}
	return *o.BootOptionsEnabled
}

// GetBootOptionsEnabledOk returns a tuple with the BootOptionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetBootOptionsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.BootOptionsEnabled) {
    return nil, false
	}
	return o.BootOptionsEnabled, true
}

// HasBootOptionsEnabled returns a boolean if a field has been set.
func (o *InlineResponse20036) HasBootOptionsEnabled() bool {
	if o != nil && !isNil(o.BootOptionsEnabled) {
		return true
	}

	return false
}

// SetBootOptionsEnabled gets a reference to the given bool and assigns it to the BootOptionsEnabled field.
func (o *InlineResponse20036) SetBootOptionsEnabled(v bool) {
	o.BootOptionsEnabled = &v
}

// GetBootNextServer returns the BootNextServer field value if set, zero value otherwise.
func (o *InlineResponse20036) GetBootNextServer() string {
	if o == nil || isNil(o.BootNextServer) {
		var ret string
		return ret
	}
	return *o.BootNextServer
}

// GetBootNextServerOk returns a tuple with the BootNextServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetBootNextServerOk() (*string, bool) {
	if o == nil || isNil(o.BootNextServer) {
    return nil, false
	}
	return o.BootNextServer, true
}

// HasBootNextServer returns a boolean if a field has been set.
func (o *InlineResponse20036) HasBootNextServer() bool {
	if o != nil && !isNil(o.BootNextServer) {
		return true
	}

	return false
}

// SetBootNextServer gets a reference to the given string and assigns it to the BootNextServer field.
func (o *InlineResponse20036) SetBootNextServer(v string) {
	o.BootNextServer = &v
}

// GetBootFileName returns the BootFileName field value if set, zero value otherwise.
func (o *InlineResponse20036) GetBootFileName() string {
	if o == nil || isNil(o.BootFileName) {
		var ret string
		return ret
	}
	return *o.BootFileName
}

// GetBootFileNameOk returns a tuple with the BootFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetBootFileNameOk() (*string, bool) {
	if o == nil || isNil(o.BootFileName) {
    return nil, false
	}
	return o.BootFileName, true
}

// HasBootFileName returns a boolean if a field has been set.
func (o *InlineResponse20036) HasBootFileName() bool {
	if o != nil && !isNil(o.BootFileName) {
		return true
	}

	return false
}

// SetBootFileName gets a reference to the given string and assigns it to the BootFileName field.
func (o *InlineResponse20036) SetBootFileName(v string) {
	o.BootFileName = &v
}

// GetDhcpOptions returns the DhcpOptions field value if set, zero value otherwise.
func (o *InlineResponse20036) GetDhcpOptions() []InlineResponse20036DhcpOptions {
	if o == nil || isNil(o.DhcpOptions) {
		var ret []InlineResponse20036DhcpOptions
		return ret
	}
	return o.DhcpOptions
}

// GetDhcpOptionsOk returns a tuple with the DhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetDhcpOptionsOk() ([]InlineResponse20036DhcpOptions, bool) {
	if o == nil || isNil(o.DhcpOptions) {
    return nil, false
	}
	return o.DhcpOptions, true
}

// HasDhcpOptions returns a boolean if a field has been set.
func (o *InlineResponse20036) HasDhcpOptions() bool {
	if o != nil && !isNil(o.DhcpOptions) {
		return true
	}

	return false
}

// SetDhcpOptions gets a reference to the given []InlineResponse20036DhcpOptions and assigns it to the DhcpOptions field.
func (o *InlineResponse20036) SetDhcpOptions(v []InlineResponse20036DhcpOptions) {
	o.DhcpOptions = v
}

// GetReservedIpRanges returns the ReservedIpRanges field value if set, zero value otherwise.
func (o *InlineResponse20036) GetReservedIpRanges() []InlineResponse20036ReservedIpRanges {
	if o == nil || isNil(o.ReservedIpRanges) {
		var ret []InlineResponse20036ReservedIpRanges
		return ret
	}
	return o.ReservedIpRanges
}

// GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetReservedIpRangesOk() ([]InlineResponse20036ReservedIpRanges, bool) {
	if o == nil || isNil(o.ReservedIpRanges) {
    return nil, false
	}
	return o.ReservedIpRanges, true
}

// HasReservedIpRanges returns a boolean if a field has been set.
func (o *InlineResponse20036) HasReservedIpRanges() bool {
	if o != nil && !isNil(o.ReservedIpRanges) {
		return true
	}

	return false
}

// SetReservedIpRanges gets a reference to the given []InlineResponse20036ReservedIpRanges and assigns it to the ReservedIpRanges field.
func (o *InlineResponse20036) SetReservedIpRanges(v []InlineResponse20036ReservedIpRanges) {
	o.ReservedIpRanges = v
}

// GetFixedIpAssignments returns the FixedIpAssignments field value if set, zero value otherwise.
func (o *InlineResponse20036) GetFixedIpAssignments() []InlineResponse20036FixedIpAssignments {
	if o == nil || isNil(o.FixedIpAssignments) {
		var ret []InlineResponse20036FixedIpAssignments
		return ret
	}
	return o.FixedIpAssignments
}

// GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetFixedIpAssignmentsOk() ([]InlineResponse20036FixedIpAssignments, bool) {
	if o == nil || isNil(o.FixedIpAssignments) {
    return nil, false
	}
	return o.FixedIpAssignments, true
}

// HasFixedIpAssignments returns a boolean if a field has been set.
func (o *InlineResponse20036) HasFixedIpAssignments() bool {
	if o != nil && !isNil(o.FixedIpAssignments) {
		return true
	}

	return false
}

// SetFixedIpAssignments gets a reference to the given []InlineResponse20036FixedIpAssignments and assigns it to the FixedIpAssignments field.
func (o *InlineResponse20036) SetFixedIpAssignments(v []InlineResponse20036FixedIpAssignments) {
	o.FixedIpAssignments = v
}

func (o InlineResponse20036) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DhcpMode) {
		toSerialize["dhcpMode"] = o.DhcpMode
	}
	if !isNil(o.DhcpRelayServerIps) {
		toSerialize["dhcpRelayServerIps"] = o.DhcpRelayServerIps
	}
	if !isNil(o.DhcpLeaseTime) {
		toSerialize["dhcpLeaseTime"] = o.DhcpLeaseTime
	}
	if !isNil(o.DnsNameserversOption) {
		toSerialize["dnsNameserversOption"] = o.DnsNameserversOption
	}
	if !isNil(o.DnsCustomNameservers) {
		toSerialize["dnsCustomNameservers"] = o.DnsCustomNameservers
	}
	if !isNil(o.BootOptionsEnabled) {
		toSerialize["bootOptionsEnabled"] = o.BootOptionsEnabled
	}
	if !isNil(o.BootNextServer) {
		toSerialize["bootNextServer"] = o.BootNextServer
	}
	if !isNil(o.BootFileName) {
		toSerialize["bootFileName"] = o.BootFileName
	}
	if !isNil(o.DhcpOptions) {
		toSerialize["dhcpOptions"] = o.DhcpOptions
	}
	if !isNil(o.ReservedIpRanges) {
		toSerialize["reservedIpRanges"] = o.ReservedIpRanges
	}
	if !isNil(o.FixedIpAssignments) {
		toSerialize["fixedIpAssignments"] = o.FixedIpAssignments
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20036 struct {
	value *InlineResponse20036
	isSet bool
}

func (v NullableInlineResponse20036) Get() *InlineResponse20036 {
	return v.value
}

func (v *NullableInlineResponse20036) Set(val *InlineResponse20036) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20036) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20036) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20036(val *InlineResponse20036) *NullableInlineResponse20036 {
	return &NullableInlineResponse20036{value: val, isSet: true}
}

func (v NullableInlineResponse20036) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20036) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


