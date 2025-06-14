/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20075Neighbors struct for InlineResponse20075Neighbors
type InlineResponse20075Neighbors struct {
	// The IPv4 address of the neighbor
	Ip *string `json:"ip,omitempty"`
	Ipv6 *InlineResponse20075Ipv6 `json:"ipv6,omitempty"`
	// Remote AS number of the neighbor
	RemoteAsNumber *int32 `json:"remoteAsNumber,omitempty"`
	// The maximum number of routes that the appliance can receive from the neighbor
	ReceiveLimit *int32 `json:"receiveLimit,omitempty"`
	// Whether the appliance will advertise routes learned from other Autonomous Systems
	AllowTransit *bool `json:"allowTransit,omitempty"`
	// The eBGP hold time in seconds for the neighbor
	EbgpHoldTimer *int32 `json:"ebgpHoldTimer,omitempty"`
	// The number of hops the appliance must traverse to establish a peering relationship with the neighbor
	EbgpMultihop *int32 `json:"ebgpMultihop,omitempty"`
	// The output interface the appliance uses to establish a peering relationship with the neighbor
	SourceInterface *string `json:"sourceInterface,omitempty"`
	// The IPv4 address of the neighbor that will establish a TCP session with the appliance
	NextHopIp *string `json:"nextHopIp,omitempty"`
	TtlSecurity *InlineResponse20075TtlSecurity `json:"ttlSecurity,omitempty"`
	Authentication *InlineResponse20075Authentication `json:"authentication,omitempty"`
}

// NewInlineResponse20075Neighbors instantiates a new InlineResponse20075Neighbors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20075Neighbors() *InlineResponse20075Neighbors {
	this := InlineResponse20075Neighbors{}
	return &this
}

// NewInlineResponse20075NeighborsWithDefaults instantiates a new InlineResponse20075Neighbors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20075NeighborsWithDefaults() *InlineResponse20075Neighbors {
	this := InlineResponse20075Neighbors{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineResponse20075Neighbors) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075Neighbors) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineResponse20075Neighbors) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineResponse20075Neighbors) SetIp(v string) {
	o.Ip = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *InlineResponse20075Neighbors) GetIpv6() InlineResponse20075Ipv6 {
	if o == nil || isNil(o.Ipv6) {
		var ret InlineResponse20075Ipv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075Neighbors) GetIpv6Ok() (*InlineResponse20075Ipv6, bool) {
	if o == nil || isNil(o.Ipv6) {
    return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *InlineResponse20075Neighbors) HasIpv6() bool {
	if o != nil && !isNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given InlineResponse20075Ipv6 and assigns it to the Ipv6 field.
func (o *InlineResponse20075Neighbors) SetIpv6(v InlineResponse20075Ipv6) {
	o.Ipv6 = &v
}

// GetRemoteAsNumber returns the RemoteAsNumber field value if set, zero value otherwise.
func (o *InlineResponse20075Neighbors) GetRemoteAsNumber() int32 {
	if o == nil || isNil(o.RemoteAsNumber) {
		var ret int32
		return ret
	}
	return *o.RemoteAsNumber
}

// GetRemoteAsNumberOk returns a tuple with the RemoteAsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075Neighbors) GetRemoteAsNumberOk() (*int32, bool) {
	if o == nil || isNil(o.RemoteAsNumber) {
    return nil, false
	}
	return o.RemoteAsNumber, true
}

// HasRemoteAsNumber returns a boolean if a field has been set.
func (o *InlineResponse20075Neighbors) HasRemoteAsNumber() bool {
	if o != nil && !isNil(o.RemoteAsNumber) {
		return true
	}

	return false
}

// SetRemoteAsNumber gets a reference to the given int32 and assigns it to the RemoteAsNumber field.
func (o *InlineResponse20075Neighbors) SetRemoteAsNumber(v int32) {
	o.RemoteAsNumber = &v
}

// GetReceiveLimit returns the ReceiveLimit field value if set, zero value otherwise.
func (o *InlineResponse20075Neighbors) GetReceiveLimit() int32 {
	if o == nil || isNil(o.ReceiveLimit) {
		var ret int32
		return ret
	}
	return *o.ReceiveLimit
}

// GetReceiveLimitOk returns a tuple with the ReceiveLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075Neighbors) GetReceiveLimitOk() (*int32, bool) {
	if o == nil || isNil(o.ReceiveLimit) {
    return nil, false
	}
	return o.ReceiveLimit, true
}

// HasReceiveLimit returns a boolean if a field has been set.
func (o *InlineResponse20075Neighbors) HasReceiveLimit() bool {
	if o != nil && !isNil(o.ReceiveLimit) {
		return true
	}

	return false
}

// SetReceiveLimit gets a reference to the given int32 and assigns it to the ReceiveLimit field.
func (o *InlineResponse20075Neighbors) SetReceiveLimit(v int32) {
	o.ReceiveLimit = &v
}

// GetAllowTransit returns the AllowTransit field value if set, zero value otherwise.
func (o *InlineResponse20075Neighbors) GetAllowTransit() bool {
	if o == nil || isNil(o.AllowTransit) {
		var ret bool
		return ret
	}
	return *o.AllowTransit
}

// GetAllowTransitOk returns a tuple with the AllowTransit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075Neighbors) GetAllowTransitOk() (*bool, bool) {
	if o == nil || isNil(o.AllowTransit) {
    return nil, false
	}
	return o.AllowTransit, true
}

// HasAllowTransit returns a boolean if a field has been set.
func (o *InlineResponse20075Neighbors) HasAllowTransit() bool {
	if o != nil && !isNil(o.AllowTransit) {
		return true
	}

	return false
}

// SetAllowTransit gets a reference to the given bool and assigns it to the AllowTransit field.
func (o *InlineResponse20075Neighbors) SetAllowTransit(v bool) {
	o.AllowTransit = &v
}

// GetEbgpHoldTimer returns the EbgpHoldTimer field value if set, zero value otherwise.
func (o *InlineResponse20075Neighbors) GetEbgpHoldTimer() int32 {
	if o == nil || isNil(o.EbgpHoldTimer) {
		var ret int32
		return ret
	}
	return *o.EbgpHoldTimer
}

// GetEbgpHoldTimerOk returns a tuple with the EbgpHoldTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075Neighbors) GetEbgpHoldTimerOk() (*int32, bool) {
	if o == nil || isNil(o.EbgpHoldTimer) {
    return nil, false
	}
	return o.EbgpHoldTimer, true
}

// HasEbgpHoldTimer returns a boolean if a field has been set.
func (o *InlineResponse20075Neighbors) HasEbgpHoldTimer() bool {
	if o != nil && !isNil(o.EbgpHoldTimer) {
		return true
	}

	return false
}

// SetEbgpHoldTimer gets a reference to the given int32 and assigns it to the EbgpHoldTimer field.
func (o *InlineResponse20075Neighbors) SetEbgpHoldTimer(v int32) {
	o.EbgpHoldTimer = &v
}

// GetEbgpMultihop returns the EbgpMultihop field value if set, zero value otherwise.
func (o *InlineResponse20075Neighbors) GetEbgpMultihop() int32 {
	if o == nil || isNil(o.EbgpMultihop) {
		var ret int32
		return ret
	}
	return *o.EbgpMultihop
}

// GetEbgpMultihopOk returns a tuple with the EbgpMultihop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075Neighbors) GetEbgpMultihopOk() (*int32, bool) {
	if o == nil || isNil(o.EbgpMultihop) {
    return nil, false
	}
	return o.EbgpMultihop, true
}

// HasEbgpMultihop returns a boolean if a field has been set.
func (o *InlineResponse20075Neighbors) HasEbgpMultihop() bool {
	if o != nil && !isNil(o.EbgpMultihop) {
		return true
	}

	return false
}

// SetEbgpMultihop gets a reference to the given int32 and assigns it to the EbgpMultihop field.
func (o *InlineResponse20075Neighbors) SetEbgpMultihop(v int32) {
	o.EbgpMultihop = &v
}

// GetSourceInterface returns the SourceInterface field value if set, zero value otherwise.
func (o *InlineResponse20075Neighbors) GetSourceInterface() string {
	if o == nil || isNil(o.SourceInterface) {
		var ret string
		return ret
	}
	return *o.SourceInterface
}

// GetSourceInterfaceOk returns a tuple with the SourceInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075Neighbors) GetSourceInterfaceOk() (*string, bool) {
	if o == nil || isNil(o.SourceInterface) {
    return nil, false
	}
	return o.SourceInterface, true
}

// HasSourceInterface returns a boolean if a field has been set.
func (o *InlineResponse20075Neighbors) HasSourceInterface() bool {
	if o != nil && !isNil(o.SourceInterface) {
		return true
	}

	return false
}

// SetSourceInterface gets a reference to the given string and assigns it to the SourceInterface field.
func (o *InlineResponse20075Neighbors) SetSourceInterface(v string) {
	o.SourceInterface = &v
}

// GetNextHopIp returns the NextHopIp field value if set, zero value otherwise.
func (o *InlineResponse20075Neighbors) GetNextHopIp() string {
	if o == nil || isNil(o.NextHopIp) {
		var ret string
		return ret
	}
	return *o.NextHopIp
}

// GetNextHopIpOk returns a tuple with the NextHopIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075Neighbors) GetNextHopIpOk() (*string, bool) {
	if o == nil || isNil(o.NextHopIp) {
    return nil, false
	}
	return o.NextHopIp, true
}

// HasNextHopIp returns a boolean if a field has been set.
func (o *InlineResponse20075Neighbors) HasNextHopIp() bool {
	if o != nil && !isNil(o.NextHopIp) {
		return true
	}

	return false
}

// SetNextHopIp gets a reference to the given string and assigns it to the NextHopIp field.
func (o *InlineResponse20075Neighbors) SetNextHopIp(v string) {
	o.NextHopIp = &v
}

// GetTtlSecurity returns the TtlSecurity field value if set, zero value otherwise.
func (o *InlineResponse20075Neighbors) GetTtlSecurity() InlineResponse20075TtlSecurity {
	if o == nil || isNil(o.TtlSecurity) {
		var ret InlineResponse20075TtlSecurity
		return ret
	}
	return *o.TtlSecurity
}

// GetTtlSecurityOk returns a tuple with the TtlSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075Neighbors) GetTtlSecurityOk() (*InlineResponse20075TtlSecurity, bool) {
	if o == nil || isNil(o.TtlSecurity) {
    return nil, false
	}
	return o.TtlSecurity, true
}

// HasTtlSecurity returns a boolean if a field has been set.
func (o *InlineResponse20075Neighbors) HasTtlSecurity() bool {
	if o != nil && !isNil(o.TtlSecurity) {
		return true
	}

	return false
}

// SetTtlSecurity gets a reference to the given InlineResponse20075TtlSecurity and assigns it to the TtlSecurity field.
func (o *InlineResponse20075Neighbors) SetTtlSecurity(v InlineResponse20075TtlSecurity) {
	o.TtlSecurity = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *InlineResponse20075Neighbors) GetAuthentication() InlineResponse20075Authentication {
	if o == nil || isNil(o.Authentication) {
		var ret InlineResponse20075Authentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075Neighbors) GetAuthenticationOk() (*InlineResponse20075Authentication, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InlineResponse20075Neighbors) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given InlineResponse20075Authentication and assigns it to the Authentication field.
func (o *InlineResponse20075Neighbors) SetAuthentication(v InlineResponse20075Authentication) {
	o.Authentication = &v
}

func (o InlineResponse20075Neighbors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	if !isNil(o.RemoteAsNumber) {
		toSerialize["remoteAsNumber"] = o.RemoteAsNumber
	}
	if !isNil(o.ReceiveLimit) {
		toSerialize["receiveLimit"] = o.ReceiveLimit
	}
	if !isNil(o.AllowTransit) {
		toSerialize["allowTransit"] = o.AllowTransit
	}
	if !isNil(o.EbgpHoldTimer) {
		toSerialize["ebgpHoldTimer"] = o.EbgpHoldTimer
	}
	if !isNil(o.EbgpMultihop) {
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

type NullableInlineResponse20075Neighbors struct {
	value *InlineResponse20075Neighbors
	isSet bool
}

func (v NullableInlineResponse20075Neighbors) Get() *InlineResponse20075Neighbors {
	return v.value
}

func (v *NullableInlineResponse20075Neighbors) Set(val *InlineResponse20075Neighbors) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20075Neighbors) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20075Neighbors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20075Neighbors(val *InlineResponse20075Neighbors) *NullableInlineResponse20075Neighbors {
	return &NullableInlineResponse20075Neighbors{value: val, isSet: true}
}

func (v NullableInlineResponse20075Neighbors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20075Neighbors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


