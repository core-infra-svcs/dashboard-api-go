/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20079 struct for InlineResponse20079
type InlineResponse20079 struct {
	// Mac address of the server.
	Mac *string `json:"mac,omitempty"`
	// Vlan id of the server.
	Vlan *int32 `json:"vlan,omitempty"`
	// Client id of the server if available.
	ClientId *string `json:"clientId,omitempty"`
	// Whether the server is allowed or blocked. Always true for configured servers.
	IsAllowed *bool `json:"isAllowed,omitempty"`
	// Last time the server was seen.
	LastSeenAt *time.Time `json:"lastSeenAt,omitempty"`
	// Devices that saw the server.
	SeenBy []NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy `json:"seenBy,omitempty"`
	// server type. Can be a 'device', 'stack', or 'discovered' (i.e client).
	Type *string `json:"type,omitempty"`
	Device *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice `json:"device,omitempty"`
	Ipv4 *NetworksNetworkIdSwitchDhcpV4ServersSeenIpv4 `json:"ipv4,omitempty"`
	// Whether the server is configured.
	IsConfigured *bool `json:"isConfigured,omitempty"`
	LastAck *NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck `json:"lastAck,omitempty"`
	LastPacket *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket `json:"lastPacket,omitempty"`
}

// NewInlineResponse20079 instantiates a new InlineResponse20079 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20079() *InlineResponse20079 {
	this := InlineResponse20079{}
	return &this
}

// NewInlineResponse20079WithDefaults instantiates a new InlineResponse20079 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20079WithDefaults() *InlineResponse20079 {
	this := InlineResponse20079{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse20079) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse20079) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse20079) SetMac(v string) {
	o.Mac = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineResponse20079) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineResponse20079) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *InlineResponse20079) SetVlan(v int32) {
	o.Vlan = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *InlineResponse20079) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
    return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *InlineResponse20079) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *InlineResponse20079) SetClientId(v string) {
	o.ClientId = &v
}

// GetIsAllowed returns the IsAllowed field value if set, zero value otherwise.
func (o *InlineResponse20079) GetIsAllowed() bool {
	if o == nil || isNil(o.IsAllowed) {
		var ret bool
		return ret
	}
	return *o.IsAllowed
}

// GetIsAllowedOk returns a tuple with the IsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetIsAllowedOk() (*bool, bool) {
	if o == nil || isNil(o.IsAllowed) {
    return nil, false
	}
	return o.IsAllowed, true
}

// HasIsAllowed returns a boolean if a field has been set.
func (o *InlineResponse20079) HasIsAllowed() bool {
	if o != nil && !isNil(o.IsAllowed) {
		return true
	}

	return false
}

// SetIsAllowed gets a reference to the given bool and assigns it to the IsAllowed field.
func (o *InlineResponse20079) SetIsAllowed(v bool) {
	o.IsAllowed = &v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise.
func (o *InlineResponse20079) GetLastSeenAt() time.Time {
	if o == nil || isNil(o.LastSeenAt) {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastSeenAt) {
    return nil, false
	}
	return o.LastSeenAt, true
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *InlineResponse20079) HasLastSeenAt() bool {
	if o != nil && !isNil(o.LastSeenAt) {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given time.Time and assigns it to the LastSeenAt field.
func (o *InlineResponse20079) SetLastSeenAt(v time.Time) {
	o.LastSeenAt = &v
}

// GetSeenBy returns the SeenBy field value if set, zero value otherwise.
func (o *InlineResponse20079) GetSeenBy() []NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy {
	if o == nil || isNil(o.SeenBy) {
		var ret []NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy
		return ret
	}
	return o.SeenBy
}

// GetSeenByOk returns a tuple with the SeenBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetSeenByOk() ([]NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy, bool) {
	if o == nil || isNil(o.SeenBy) {
    return nil, false
	}
	return o.SeenBy, true
}

// HasSeenBy returns a boolean if a field has been set.
func (o *InlineResponse20079) HasSeenBy() bool {
	if o != nil && !isNil(o.SeenBy) {
		return true
	}

	return false
}

// SetSeenBy gets a reference to the given []NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy and assigns it to the SeenBy field.
func (o *InlineResponse20079) SetSeenBy(v []NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) {
	o.SeenBy = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20079) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20079) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20079) SetType(v string) {
	o.Type = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InlineResponse20079) GetDevice() NetworksNetworkIdSwitchDhcpV4ServersSeenDevice {
	if o == nil || isNil(o.Device) {
		var ret NetworksNetworkIdSwitchDhcpV4ServersSeenDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetDeviceOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenDevice, bool) {
	if o == nil || isNil(o.Device) {
    return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InlineResponse20079) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NetworksNetworkIdSwitchDhcpV4ServersSeenDevice and assigns it to the Device field.
func (o *InlineResponse20079) SetDevice(v NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) {
	o.Device = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *InlineResponse20079) GetIpv4() NetworksNetworkIdSwitchDhcpV4ServersSeenIpv4 {
	if o == nil || isNil(o.Ipv4) {
		var ret NetworksNetworkIdSwitchDhcpV4ServersSeenIpv4
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetIpv4Ok() (*NetworksNetworkIdSwitchDhcpV4ServersSeenIpv4, bool) {
	if o == nil || isNil(o.Ipv4) {
    return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *InlineResponse20079) HasIpv4() bool {
	if o != nil && !isNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given NetworksNetworkIdSwitchDhcpV4ServersSeenIpv4 and assigns it to the Ipv4 field.
func (o *InlineResponse20079) SetIpv4(v NetworksNetworkIdSwitchDhcpV4ServersSeenIpv4) {
	o.Ipv4 = &v
}

// GetIsConfigured returns the IsConfigured field value if set, zero value otherwise.
func (o *InlineResponse20079) GetIsConfigured() bool {
	if o == nil || isNil(o.IsConfigured) {
		var ret bool
		return ret
	}
	return *o.IsConfigured
}

// GetIsConfiguredOk returns a tuple with the IsConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetIsConfiguredOk() (*bool, bool) {
	if o == nil || isNil(o.IsConfigured) {
    return nil, false
	}
	return o.IsConfigured, true
}

// HasIsConfigured returns a boolean if a field has been set.
func (o *InlineResponse20079) HasIsConfigured() bool {
	if o != nil && !isNil(o.IsConfigured) {
		return true
	}

	return false
}

// SetIsConfigured gets a reference to the given bool and assigns it to the IsConfigured field.
func (o *InlineResponse20079) SetIsConfigured(v bool) {
	o.IsConfigured = &v
}

// GetLastAck returns the LastAck field value if set, zero value otherwise.
func (o *InlineResponse20079) GetLastAck() NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck {
	if o == nil || isNil(o.LastAck) {
		var ret NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck
		return ret
	}
	return *o.LastAck
}

// GetLastAckOk returns a tuple with the LastAck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetLastAckOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck, bool) {
	if o == nil || isNil(o.LastAck) {
    return nil, false
	}
	return o.LastAck, true
}

// HasLastAck returns a boolean if a field has been set.
func (o *InlineResponse20079) HasLastAck() bool {
	if o != nil && !isNil(o.LastAck) {
		return true
	}

	return false
}

// SetLastAck gets a reference to the given NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck and assigns it to the LastAck field.
func (o *InlineResponse20079) SetLastAck(v NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) {
	o.LastAck = &v
}

// GetLastPacket returns the LastPacket field value if set, zero value otherwise.
func (o *InlineResponse20079) GetLastPacket() NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket {
	if o == nil || isNil(o.LastPacket) {
		var ret NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket
		return ret
	}
	return *o.LastPacket
}

// GetLastPacketOk returns a tuple with the LastPacket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetLastPacketOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket, bool) {
	if o == nil || isNil(o.LastPacket) {
    return nil, false
	}
	return o.LastPacket, true
}

// HasLastPacket returns a boolean if a field has been set.
func (o *InlineResponse20079) HasLastPacket() bool {
	if o != nil && !isNil(o.LastPacket) {
		return true
	}

	return false
}

// SetLastPacket gets a reference to the given NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket and assigns it to the LastPacket field.
func (o *InlineResponse20079) SetLastPacket(v NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) {
	o.LastPacket = &v
}

func (o InlineResponse20079) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !isNil(o.IsAllowed) {
		toSerialize["isAllowed"] = o.IsAllowed
	}
	if !isNil(o.LastSeenAt) {
		toSerialize["lastSeenAt"] = o.LastSeenAt
	}
	if !isNil(o.SeenBy) {
		toSerialize["seenBy"] = o.SeenBy
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !isNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	if !isNil(o.IsConfigured) {
		toSerialize["isConfigured"] = o.IsConfigured
	}
	if !isNil(o.LastAck) {
		toSerialize["lastAck"] = o.LastAck
	}
	if !isNil(o.LastPacket) {
		toSerialize["lastPacket"] = o.LastPacket
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20079 struct {
	value *InlineResponse20079
	isSet bool
}

func (v NullableInlineResponse20079) Get() *InlineResponse20079 {
	return v.value
}

func (v *NullableInlineResponse20079) Set(val *InlineResponse20079) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20079) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20079) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20079(val *InlineResponse20079) *NullableInlineResponse20079 {
	return &NullableInlineResponse20079{value: val, isSet: true}
}

func (v NullableInlineResponse20079) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20079) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


