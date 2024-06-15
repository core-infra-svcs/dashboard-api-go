/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject80 struct for InlineObject80
type InlineObject80 struct {
	// The name of the camera wireless profile. This parameter is required.
	Name string `json:"name"`
	Ssid NetworksNetworkIdCameraWirelessProfilesSsid1 `json:"ssid"`
	Identity *NetworksNetworkIdCameraWirelessProfilesIdentity `json:"identity,omitempty"`
}

// NewInlineObject80 instantiates a new InlineObject80 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject80(name string, ssid NetworksNetworkIdCameraWirelessProfilesSsid1) *InlineObject80 {
	this := InlineObject80{}
	this.Name = name
	this.Ssid = ssid
	return &this
}

// NewInlineObject80WithDefaults instantiates a new InlineObject80 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject80WithDefaults() *InlineObject80 {
	this := InlineObject80{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject80) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject80) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject80) SetName(v string) {
	o.Name = v
}

// GetSsid returns the Ssid field value
func (o *InlineObject80) GetSsid() NetworksNetworkIdCameraWirelessProfilesSsid1 {
	if o == nil {
		var ret NetworksNetworkIdCameraWirelessProfilesSsid1
		return ret
	}

	return o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value
// and a boolean to check if the value has been set.
func (o *InlineObject80) GetSsidOk() (*NetworksNetworkIdCameraWirelessProfilesSsid1, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ssid, true
}

// SetSsid sets field value
func (o *InlineObject80) SetSsid(v NetworksNetworkIdCameraWirelessProfilesSsid1) {
	o.Ssid = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *InlineObject80) GetIdentity() NetworksNetworkIdCameraWirelessProfilesIdentity {
	if o == nil || isNil(o.Identity) {
		var ret NetworksNetworkIdCameraWirelessProfilesIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject80) GetIdentityOk() (*NetworksNetworkIdCameraWirelessProfilesIdentity, bool) {
	if o == nil || isNil(o.Identity) {
    return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *InlineObject80) HasIdentity() bool {
	if o != nil && !isNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given NetworksNetworkIdCameraWirelessProfilesIdentity and assigns it to the Identity field.
func (o *InlineObject80) SetIdentity(v NetworksNetworkIdCameraWirelessProfilesIdentity) {
	o.Identity = &v
}

func (o InlineObject80) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject80 struct {
	value *InlineObject80
	isSet bool
}

func (v NullableInlineObject80) Get() *InlineObject80 {
	return v.value
}

func (v *NullableInlineObject80) Set(val *InlineObject80) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject80) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject80) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject80(val *InlineObject80) *NullableInlineObject80 {
	return &NullableInlineObject80{value: val, isSet: true}
}

func (v NullableInlineObject80) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject80) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


