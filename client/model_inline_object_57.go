/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject57 struct for InlineObject57
type InlineObject57 struct {
	// The name of the camera wireless profile.
	Name *string `json:"name,omitempty"`
	Ssid *NetworksNetworkIdCameraWirelessProfilesSsid `json:"ssid,omitempty"`
	Identity *NetworksNetworkIdCameraWirelessProfilesIdentity `json:"identity,omitempty"`
}

// NewInlineObject57 instantiates a new InlineObject57 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject57() *InlineObject57 {
	this := InlineObject57{}
	return &this
}

// NewInlineObject57WithDefaults instantiates a new InlineObject57 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject57WithDefaults() *InlineObject57 {
	this := InlineObject57{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject57) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject57) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject57) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject57) SetName(v string) {
	o.Name = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *InlineObject57) GetSsid() NetworksNetworkIdCameraWirelessProfilesSsid {
	if o == nil || o.Ssid == nil {
		var ret NetworksNetworkIdCameraWirelessProfilesSsid
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject57) GetSsidOk() (*NetworksNetworkIdCameraWirelessProfilesSsid, bool) {
	if o == nil || o.Ssid == nil {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *InlineObject57) HasSsid() bool {
	if o != nil && o.Ssid != nil {
		return true
	}

	return false
}

// SetSsid gets a reference to the given NetworksNetworkIdCameraWirelessProfilesSsid and assigns it to the Ssid field.
func (o *InlineObject57) SetSsid(v NetworksNetworkIdCameraWirelessProfilesSsid) {
	o.Ssid = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *InlineObject57) GetIdentity() NetworksNetworkIdCameraWirelessProfilesIdentity {
	if o == nil || o.Identity == nil {
		var ret NetworksNetworkIdCameraWirelessProfilesIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject57) GetIdentityOk() (*NetworksNetworkIdCameraWirelessProfilesIdentity, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *InlineObject57) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given NetworksNetworkIdCameraWirelessProfilesIdentity and assigns it to the Identity field.
func (o *InlineObject57) SetIdentity(v NetworksNetworkIdCameraWirelessProfilesIdentity) {
	o.Identity = &v
}

func (o InlineObject57) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Ssid != nil {
		toSerialize["ssid"] = o.Ssid
	}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject57 struct {
	value *InlineObject57
	isSet bool
}

func (v NullableInlineObject57) Get() *InlineObject57 {
	return v.value
}

func (v *NullableInlineObject57) Set(val *InlineObject57) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject57) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject57) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject57(val *InlineObject57) *NullableInlineObject57 {
	return &NullableInlineObject57{value: val, isSet: true}
}

func (v NullableInlineObject57) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject57) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


