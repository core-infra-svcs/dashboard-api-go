/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceVpnBgpNeighbors struct for NetworksNetworkIdApplianceVpnBgpNeighbors
type NetworksNetworkIdApplianceVpnBgpNeighbors struct {
	// The IPv4 address of the neighbor
	Ip *string `json:"ip,omitempty"`
	Ipv6 *NetworksNetworkIdApplianceVpnBgpIpv6 `json:"ipv6,omitempty"`
	// Remote ASN of the neighbor. The remote ASN must be an integer between 1 and 4294967295.
	RemoteAsNumber int32 `json:"remoteAsNumber"`
	// The receive limit is the maximum number of routes that can be received from any BGP peer. The receive limit must be an integer between 0 and 4294967295. When absent, it defaults to 0.
	ReceiveLimit *int32 `json:"receiveLimit,omitempty"`
	// When this feature is on, the Meraki device will advertise routes learned from other Autonomous Systems, thereby allowing traffic between Autonomous Systems to transit this AS. When absent, it defaults to false.
	AllowTransit *bool `json:"allowTransit,omitempty"`
	// The EBGP hold timer in seconds for each neighbor. The EBGP hold timer must be an integer between 12 and 240.
	EbgpHoldTimer int32 `json:"ebgpHoldTimer"`
	// Configure this if the neighbor is not adjacent. The EBGP multi-hop must be an integer between 1 and 255.
	EbgpMultihop int32 `json:"ebgpMultihop"`
	// The output interface for peering with the remote BGP peer. Valid values are: 'wired0', 'wired1' or 'vlan{VLAN ID}'(e.g. 'vlan123').
	SourceInterface *string `json:"sourceInterface,omitempty"`
	// The IPv4 address of the remote BGP peer that will establish a TCP session with the local MX.
	NextHopIp *string `json:"nextHopIp,omitempty"`
	TtlSecurity *NetworksNetworkIdApplianceVpnBgpTtlSecurity `json:"ttlSecurity,omitempty"`
	Authentication *NetworksNetworkIdApplianceVpnBgpAuthentication `json:"authentication,omitempty"`
}

// NewNetworksNetworkIdApplianceVpnBgpNeighbors instantiates a new NetworksNetworkIdApplianceVpnBgpNeighbors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceVpnBgpNeighbors(remoteAsNumber int32, ebgpHoldTimer int32, ebgpMultihop int32) *NetworksNetworkIdApplianceVpnBgpNeighbors {
	this := NetworksNetworkIdApplianceVpnBgpNeighbors{}
	this.RemoteAsNumber = remoteAsNumber
	this.EbgpHoldTimer = ebgpHoldTimer
	this.EbgpMultihop = ebgpMultihop
	return &this
}

// NewNetworksNetworkIdApplianceVpnBgpNeighborsWithDefaults instantiates a new NetworksNetworkIdApplianceVpnBgpNeighbors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceVpnBgpNeighborsWithDefaults() *NetworksNetworkIdApplianceVpnBgpNeighbors {
	this := NetworksNetworkIdApplianceVpnBgpNeighbors{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetIp(v string) {
	o.Ip = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetIpv6() NetworksNetworkIdApplianceVpnBgpIpv6 {
	if o == nil || isNil(o.Ipv6) {
		var ret NetworksNetworkIdApplianceVpnBgpIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetIpv6Ok() (*NetworksNetworkIdApplianceVpnBgpIpv6, bool) {
	if o == nil || isNil(o.Ipv6) {
    return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) HasIpv6() bool {
	if o != nil && !isNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given NetworksNetworkIdApplianceVpnBgpIpv6 and assigns it to the Ipv6 field.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetIpv6(v NetworksNetworkIdApplianceVpnBgpIpv6) {
	o.Ipv6 = &v
}

// GetRemoteAsNumber returns the RemoteAsNumber field value
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetRemoteAsNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RemoteAsNumber
}

// GetRemoteAsNumberOk returns a tuple with the RemoteAsNumber field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetRemoteAsNumberOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RemoteAsNumber, true
}

// SetRemoteAsNumber sets field value
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetRemoteAsNumber(v int32) {
	o.RemoteAsNumber = v
}

// GetReceiveLimit returns the ReceiveLimit field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetReceiveLimit() int32 {
	if o == nil || isNil(o.ReceiveLimit) {
		var ret int32
		return ret
	}
	return *o.ReceiveLimit
}

// GetReceiveLimitOk returns a tuple with the ReceiveLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetReceiveLimitOk() (*int32, bool) {
	if o == nil || isNil(o.ReceiveLimit) {
    return nil, false
	}
	return o.ReceiveLimit, true
}

// HasReceiveLimit returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) HasReceiveLimit() bool {
	if o != nil && !isNil(o.ReceiveLimit) {
		return true
	}

	return false
}

// SetReceiveLimit gets a reference to the given int32 and assigns it to the ReceiveLimit field.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetReceiveLimit(v int32) {
	o.ReceiveLimit = &v
}

// GetAllowTransit returns the AllowTransit field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetAllowTransit() bool {
	if o == nil || isNil(o.AllowTransit) {
		var ret bool
		return ret
	}
	return *o.AllowTransit
}

// GetAllowTransitOk returns a tuple with the AllowTransit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetAllowTransitOk() (*bool, bool) {
	if o == nil || isNil(o.AllowTransit) {
    return nil, false
	}
	return o.AllowTransit, true
}

// HasAllowTransit returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) HasAllowTransit() bool {
	if o != nil && !isNil(o.AllowTransit) {
		return true
	}

	return false
}

// SetAllowTransit gets a reference to the given bool and assigns it to the AllowTransit field.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetAllowTransit(v bool) {
	o.AllowTransit = &v
}

// GetEbgpHoldTimer returns the EbgpHoldTimer field value
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetEbgpHoldTimer() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EbgpHoldTimer
}

// GetEbgpHoldTimerOk returns a tuple with the EbgpHoldTimer field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetEbgpHoldTimerOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EbgpHoldTimer, true
}

// SetEbgpHoldTimer sets field value
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetEbgpHoldTimer(v int32) {
	o.EbgpHoldTimer = v
}

// GetEbgpMultihop returns the EbgpMultihop field value
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetEbgpMultihop() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EbgpMultihop
}

// GetEbgpMultihopOk returns a tuple with the EbgpMultihop field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetEbgpMultihopOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EbgpMultihop, true
}

// SetEbgpMultihop sets field value
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetEbgpMultihop(v int32) {
	o.EbgpMultihop = v
}

// GetSourceInterface returns the SourceInterface field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetSourceInterface() string {
	if o == nil || isNil(o.SourceInterface) {
		var ret string
		return ret
	}
	return *o.SourceInterface
}

// GetSourceInterfaceOk returns a tuple with the SourceInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetSourceInterfaceOk() (*string, bool) {
	if o == nil || isNil(o.SourceInterface) {
    return nil, false
	}
	return o.SourceInterface, true
}

// HasSourceInterface returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) HasSourceInterface() bool {
	if o != nil && !isNil(o.SourceInterface) {
		return true
	}

	return false
}

// SetSourceInterface gets a reference to the given string and assigns it to the SourceInterface field.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetSourceInterface(v string) {
	o.SourceInterface = &v
}

// GetNextHopIp returns the NextHopIp field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetNextHopIp() string {
	if o == nil || isNil(o.NextHopIp) {
		var ret string
		return ret
	}
	return *o.NextHopIp
}

// GetNextHopIpOk returns a tuple with the NextHopIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetNextHopIpOk() (*string, bool) {
	if o == nil || isNil(o.NextHopIp) {
    return nil, false
	}
	return o.NextHopIp, true
}

// HasNextHopIp returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) HasNextHopIp() bool {
	if o != nil && !isNil(o.NextHopIp) {
		return true
	}

	return false
}

// SetNextHopIp gets a reference to the given string and assigns it to the NextHopIp field.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetNextHopIp(v string) {
	o.NextHopIp = &v
}

// GetTtlSecurity returns the TtlSecurity field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetTtlSecurity() NetworksNetworkIdApplianceVpnBgpTtlSecurity {
	if o == nil || isNil(o.TtlSecurity) {
		var ret NetworksNetworkIdApplianceVpnBgpTtlSecurity
		return ret
	}
	return *o.TtlSecurity
}

// GetTtlSecurityOk returns a tuple with the TtlSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetTtlSecurityOk() (*NetworksNetworkIdApplianceVpnBgpTtlSecurity, bool) {
	if o == nil || isNil(o.TtlSecurity) {
    return nil, false
	}
	return o.TtlSecurity, true
}

// HasTtlSecurity returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) HasTtlSecurity() bool {
	if o != nil && !isNil(o.TtlSecurity) {
		return true
	}

	return false
}

// SetTtlSecurity gets a reference to the given NetworksNetworkIdApplianceVpnBgpTtlSecurity and assigns it to the TtlSecurity field.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetTtlSecurity(v NetworksNetworkIdApplianceVpnBgpTtlSecurity) {
	o.TtlSecurity = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetAuthentication() NetworksNetworkIdApplianceVpnBgpAuthentication {
	if o == nil || isNil(o.Authentication) {
		var ret NetworksNetworkIdApplianceVpnBgpAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetAuthenticationOk() (*NetworksNetworkIdApplianceVpnBgpAuthentication, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given NetworksNetworkIdApplianceVpnBgpAuthentication and assigns it to the Authentication field.
func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetAuthentication(v NetworksNetworkIdApplianceVpnBgpAuthentication) {
	o.Authentication = &v
}

func (o NetworksNetworkIdApplianceVpnBgpNeighbors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	if true {
		toSerialize["remoteAsNumber"] = o.RemoteAsNumber
	}
	if !isNil(o.ReceiveLimit) {
		toSerialize["receiveLimit"] = o.ReceiveLimit
	}
	if !isNil(o.AllowTransit) {
		toSerialize["allowTransit"] = o.AllowTransit
	}
	if true {
		toSerialize["ebgpHoldTimer"] = o.EbgpHoldTimer
	}
	if true {
		toSerialize["ebgpMultihop"] = o.EbgpMultihop
	}
	if !isNil(o.SourceInterface) {
		toSerialize["sourceInterface"] = o.SourceInterface
	}
	if !isNil(o.NextHopIp) {
		toSerialize["nextHopIp"] = o.NextHopIp
	}
	if !isNil(o.TtlSecurity) {
		toSerialize["ttlSecurity"] = o.TtlSecurity
	}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceVpnBgpNeighbors struct {
	value *NetworksNetworkIdApplianceVpnBgpNeighbors
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceVpnBgpNeighbors) Get() *NetworksNetworkIdApplianceVpnBgpNeighbors {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceVpnBgpNeighbors) Set(val *NetworksNetworkIdApplianceVpnBgpNeighbors) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceVpnBgpNeighbors) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceVpnBgpNeighbors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceVpnBgpNeighbors(val *NetworksNetworkIdApplianceVpnBgpNeighbors) *NullableNetworksNetworkIdApplianceVpnBgpNeighbors {
	return &NullableNetworksNetworkIdApplianceVpnBgpNeighbors{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceVpnBgpNeighbors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceVpnBgpNeighbors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


