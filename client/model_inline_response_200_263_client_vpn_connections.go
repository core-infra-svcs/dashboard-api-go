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

// InlineResponse200263ClientVpnConnections struct for InlineResponse200263ClientVpnConnections
type InlineResponse200263ClientVpnConnections struct {
	// The IP address of the VPN the client last connected to
	RemoteIp *string `json:"remoteIp,omitempty"`
	// The time the client last connected to the VPN
	ConnectedAt *int32 `json:"connectedAt,omitempty"`
	// The time the client last disconnected from the VPN
	DisconnectedAt *int32 `json:"disconnectedAt,omitempty"`
}

// NewInlineResponse200263ClientVpnConnections instantiates a new InlineResponse200263ClientVpnConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200263ClientVpnConnections() *InlineResponse200263ClientVpnConnections {
	this := InlineResponse200263ClientVpnConnections{}
	return &this
}

// NewInlineResponse200263ClientVpnConnectionsWithDefaults instantiates a new InlineResponse200263ClientVpnConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200263ClientVpnConnectionsWithDefaults() *InlineResponse200263ClientVpnConnections {
	this := InlineResponse200263ClientVpnConnections{}
	return &this
}

// GetRemoteIp returns the RemoteIp field value if set, zero value otherwise.
func (o *InlineResponse200263ClientVpnConnections) GetRemoteIp() string {
	if o == nil || isNil(o.RemoteIp) {
		var ret string
		return ret
	}
	return *o.RemoteIp
}

// GetRemoteIpOk returns a tuple with the RemoteIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200263ClientVpnConnections) GetRemoteIpOk() (*string, bool) {
	if o == nil || isNil(o.RemoteIp) {
    return nil, false
	}
	return o.RemoteIp, true
}

// HasRemoteIp returns a boolean if a field has been set.
func (o *InlineResponse200263ClientVpnConnections) HasRemoteIp() bool {
	if o != nil && !isNil(o.RemoteIp) {
		return true
	}

	return false
}

// SetRemoteIp gets a reference to the given string and assigns it to the RemoteIp field.
func (o *InlineResponse200263ClientVpnConnections) SetRemoteIp(v string) {
	o.RemoteIp = &v
}

// GetConnectedAt returns the ConnectedAt field value if set, zero value otherwise.
func (o *InlineResponse200263ClientVpnConnections) GetConnectedAt() int32 {
	if o == nil || isNil(o.ConnectedAt) {
		var ret int32
		return ret
	}
	return *o.ConnectedAt
}

// GetConnectedAtOk returns a tuple with the ConnectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200263ClientVpnConnections) GetConnectedAtOk() (*int32, bool) {
	if o == nil || isNil(o.ConnectedAt) {
    return nil, false
	}
	return o.ConnectedAt, true
}

// HasConnectedAt returns a boolean if a field has been set.
func (o *InlineResponse200263ClientVpnConnections) HasConnectedAt() bool {
	if o != nil && !isNil(o.ConnectedAt) {
		return true
	}

	return false
}

// SetConnectedAt gets a reference to the given int32 and assigns it to the ConnectedAt field.
func (o *InlineResponse200263ClientVpnConnections) SetConnectedAt(v int32) {
	o.ConnectedAt = &v
}

// GetDisconnectedAt returns the DisconnectedAt field value if set, zero value otherwise.
func (o *InlineResponse200263ClientVpnConnections) GetDisconnectedAt() int32 {
	if o == nil || isNil(o.DisconnectedAt) {
		var ret int32
		return ret
	}
	return *o.DisconnectedAt
}

// GetDisconnectedAtOk returns a tuple with the DisconnectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200263ClientVpnConnections) GetDisconnectedAtOk() (*int32, bool) {
	if o == nil || isNil(o.DisconnectedAt) {
    return nil, false
	}
	return o.DisconnectedAt, true
}

// HasDisconnectedAt returns a boolean if a field has been set.
func (o *InlineResponse200263ClientVpnConnections) HasDisconnectedAt() bool {
	if o != nil && !isNil(o.DisconnectedAt) {
		return true
	}

	return false
}

// SetDisconnectedAt gets a reference to the given int32 and assigns it to the DisconnectedAt field.
func (o *InlineResponse200263ClientVpnConnections) SetDisconnectedAt(v int32) {
	o.DisconnectedAt = &v
}

func (o InlineResponse200263ClientVpnConnections) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RemoteIp) {
		toSerialize["remoteIp"] = o.RemoteIp
	}
	if !isNil(o.ConnectedAt) {
		toSerialize["connectedAt"] = o.ConnectedAt
	}
	if !isNil(o.DisconnectedAt) {
		toSerialize["disconnectedAt"] = o.DisconnectedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200263ClientVpnConnections struct {
	value *InlineResponse200263ClientVpnConnections
	isSet bool
}

func (v NullableInlineResponse200263ClientVpnConnections) Get() *InlineResponse200263ClientVpnConnections {
	return v.value
}

func (v *NullableInlineResponse200263ClientVpnConnections) Set(val *InlineResponse200263ClientVpnConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200263ClientVpnConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200263ClientVpnConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200263ClientVpnConnections(val *InlineResponse200263ClientVpnConnections) *NullableInlineResponse200263ClientVpnConnections {
	return &NullableInlineResponse200263ClientVpnConnections{value: val, isSet: true}
}

func (v NullableInlineResponse200263ClientVpnConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200263ClientVpnConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


