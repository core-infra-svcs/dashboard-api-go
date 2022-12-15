/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject65 struct for InlineObject65
type InlineObject65 struct {
	// The name of the camera wireless profile.
	Name *string `json:"name,omitempty"`
	Ssid *NetworksNetworkIdCameraWirelessProfilesSsid `json:"ssid,omitempty"`
	Identity *NetworksNetworkIdCameraWirelessProfilesIdentity `json:"identity,omitempty"`
}

// NewInlineObject65 instantiates a new InlineObject65 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject65() *InlineObject65 {
	this := InlineObject65{}
	return &this
}

// NewInlineObject65WithDefaults instantiates a new InlineObject65 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject65WithDefaults() *InlineObject65 {
	this := InlineObject65{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject65) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject65) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject65) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject65) SetName(v string) {
	o.Name = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *InlineObject65) GetSsid() NetworksNetworkIdCameraWirelessProfilesSsid {
	if o == nil || isNil(o.Ssid) {
		var ret NetworksNetworkIdCameraWirelessProfilesSsid
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject65) GetSsidOk() (*NetworksNetworkIdCameraWirelessProfilesSsid, bool) {
	if o == nil || isNil(o.Ssid) {
    return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *InlineObject65) HasSsid() bool {
	if o != nil && !isNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given NetworksNetworkIdCameraWirelessProfilesSsid and assigns it to the Ssid field.
func (o *InlineObject65) SetSsid(v NetworksNetworkIdCameraWirelessProfilesSsid) {
	o.Ssid = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *InlineObject65) GetIdentity() NetworksNetworkIdCameraWirelessProfilesIdentity {
	if o == nil || isNil(o.Identity) {
		var ret NetworksNetworkIdCameraWirelessProfilesIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject65) GetIdentityOk() (*NetworksNetworkIdCameraWirelessProfilesIdentity, bool) {
	if o == nil || isNil(o.Identity) {
    return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *InlineObject65) HasIdentity() bool {
	if o != nil && !isNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given NetworksNetworkIdCameraWirelessProfilesIdentity and assigns it to the Identity field.
func (o *InlineObject65) SetIdentity(v NetworksNetworkIdCameraWirelessProfilesIdentity) {
	o.Identity = &v
}

func (o InlineObject65) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !isNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject65 struct {
	value *InlineObject65
	isSet bool
}

func (v NullableInlineObject65) Get() *InlineObject65 {
	return v.value
}

func (v *NullableInlineObject65) Set(val *InlineObject65) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject65) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject65) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject65(val *InlineObject65) *NullableInlineObject65 {
	return &NullableInlineObject65{value: val, isSet: true}
}

func (v NullableInlineObject65) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject65) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


