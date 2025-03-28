/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200232MerakiVpnPeers struct for InlineResponse200232MerakiVpnPeers
type InlineResponse200232MerakiVpnPeers struct {
	// Network ID
	NetworkId *string `json:"networkId,omitempty"`
	// Network Name
	NetworkName *string `json:"networkName,omitempty"`
	// Reachability
	Reachability *string `json:"reachability,omitempty"`
}

// NewInlineResponse200232MerakiVpnPeers instantiates a new InlineResponse200232MerakiVpnPeers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200232MerakiVpnPeers() *InlineResponse200232MerakiVpnPeers {
	this := InlineResponse200232MerakiVpnPeers{}
	return &this
}

// NewInlineResponse200232MerakiVpnPeersWithDefaults instantiates a new InlineResponse200232MerakiVpnPeers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200232MerakiVpnPeersWithDefaults() *InlineResponse200232MerakiVpnPeers {
	this := InlineResponse200232MerakiVpnPeers{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200232MerakiVpnPeers) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200232MerakiVpnPeers) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200232MerakiVpnPeers) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200232MerakiVpnPeers) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetNetworkName returns the NetworkName field value if set, zero value otherwise.
func (o *InlineResponse200232MerakiVpnPeers) GetNetworkName() string {
	if o == nil || isNil(o.NetworkName) {
		var ret string
		return ret
	}
	return *o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200232MerakiVpnPeers) GetNetworkNameOk() (*string, bool) {
	if o == nil || isNil(o.NetworkName) {
    return nil, false
	}
	return o.NetworkName, true
}

// HasNetworkName returns a boolean if a field has been set.
func (o *InlineResponse200232MerakiVpnPeers) HasNetworkName() bool {
	if o != nil && !isNil(o.NetworkName) {
		return true
	}

	return false
}

// SetNetworkName gets a reference to the given string and assigns it to the NetworkName field.
func (o *InlineResponse200232MerakiVpnPeers) SetNetworkName(v string) {
	o.NetworkName = &v
}

// GetReachability returns the Reachability field value if set, zero value otherwise.
func (o *InlineResponse200232MerakiVpnPeers) GetReachability() string {
	if o == nil || isNil(o.Reachability) {
		var ret string
		return ret
	}
	return *o.Reachability
}

// GetReachabilityOk returns a tuple with the Reachability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200232MerakiVpnPeers) GetReachabilityOk() (*string, bool) {
	if o == nil || isNil(o.Reachability) {
    return nil, false
	}
	return o.Reachability, true
}

// HasReachability returns a boolean if a field has been set.
func (o *InlineResponse200232MerakiVpnPeers) HasReachability() bool {
	if o != nil && !isNil(o.Reachability) {
		return true
	}

	return false
}

// SetReachability gets a reference to the given string and assigns it to the Reachability field.
func (o *InlineResponse200232MerakiVpnPeers) SetReachability(v string) {
	o.Reachability = &v
}

func (o InlineResponse200232MerakiVpnPeers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.NetworkName) {
		toSerialize["networkName"] = o.NetworkName
	}
	if !isNil(o.Reachability) {
		toSerialize["reachability"] = o.Reachability
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200232MerakiVpnPeers struct {
	value *InlineResponse200232MerakiVpnPeers
	isSet bool
}

func (v NullableInlineResponse200232MerakiVpnPeers) Get() *InlineResponse200232MerakiVpnPeers {
	return v.value
}

func (v *NullableInlineResponse200232MerakiVpnPeers) Set(val *InlineResponse200232MerakiVpnPeers) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200232MerakiVpnPeers) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200232MerakiVpnPeers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200232MerakiVpnPeers(val *InlineResponse200232MerakiVpnPeers) *NullableInlineResponse200232MerakiVpnPeers {
	return &NullableInlineResponse200232MerakiVpnPeers{value: val, isSet: true}
}

func (v NullableInlineResponse200232MerakiVpnPeers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200232MerakiVpnPeers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


