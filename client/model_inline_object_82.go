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

// InlineObject82 struct for InlineObject82
type InlineObject82 struct {
	// The name of the camera wireless profile. This parameter is required.
	Name string `json:"name"`
	Ssid NetworksNetworkIdCameraWirelessProfilesSsid1 `json:"ssid"`
	Identity *NetworksNetworkIdCameraWirelessProfilesIdentity `json:"identity,omitempty"`
}

// NewInlineObject82 instantiates a new InlineObject82 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject82(name string, ssid NetworksNetworkIdCameraWirelessProfilesSsid1) *InlineObject82 {
	this := InlineObject82{}
	this.Name = name
	this.Ssid = ssid
	return &this
}

// NewInlineObject82WithDefaults instantiates a new InlineObject82 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject82WithDefaults() *InlineObject82 {
	this := InlineObject82{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject82) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject82) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject82) SetName(v string) {
	o.Name = v
}

// GetSsid returns the Ssid field value
func (o *InlineObject82) GetSsid() NetworksNetworkIdCameraWirelessProfilesSsid1 {
	if o == nil {
		var ret NetworksNetworkIdCameraWirelessProfilesSsid1
		return ret
	}

	return o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value
// and a boolean to check if the value has been set.
func (o *InlineObject82) GetSsidOk() (*NetworksNetworkIdCameraWirelessProfilesSsid1, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ssid, true
}

// SetSsid sets field value
func (o *InlineObject82) SetSsid(v NetworksNetworkIdCameraWirelessProfilesSsid1) {
	o.Ssid = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *InlineObject82) GetIdentity() NetworksNetworkIdCameraWirelessProfilesIdentity {
	if o == nil || isNil(o.Identity) {
		var ret NetworksNetworkIdCameraWirelessProfilesIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject82) GetIdentityOk() (*NetworksNetworkIdCameraWirelessProfilesIdentity, bool) {
	if o == nil || isNil(o.Identity) {
    return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *InlineObject82) HasIdentity() bool {
	if o != nil && !isNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given NetworksNetworkIdCameraWirelessProfilesIdentity and assigns it to the Identity field.
func (o *InlineObject82) SetIdentity(v NetworksNetworkIdCameraWirelessProfilesIdentity) {
	o.Identity = &v
}

func (o InlineObject82) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["ssid"] = o.Ssid
	}
	if !isNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject82 struct {
	value *InlineObject82
	isSet bool
}

func (v NullableInlineObject82) Get() *InlineObject82 {
	return v.value
}

func (v *NullableInlineObject82) Set(val *InlineObject82) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject82) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject82) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject82(val *InlineObject82) *NullableInlineObject82 {
	return &NullableInlineObject82{value: val, isSet: true}
}

func (v NullableInlineObject82) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject82) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


