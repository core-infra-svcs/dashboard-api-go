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

// InlineResponse20010InterfacesWan1 WAN 1 settings.
type InlineResponse20010InterfacesWan1 struct {
	// Enable or disable the interface.
	Enabled *bool `json:"enabled,omitempty"`
	VlanTagging *InlineResponse20010InterfacesWan1VlanTagging `json:"vlanTagging,omitempty"`
	Svis *InlineResponse20010InterfacesWan1Svis `json:"svis,omitempty"`
	Pppoe *InlineResponse20010InterfacesWan1Pppoe `json:"pppoe,omitempty"`
}

// NewInlineResponse20010InterfacesWan1 instantiates a new InlineResponse20010InterfacesWan1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20010InterfacesWan1() *InlineResponse20010InterfacesWan1 {
	this := InlineResponse20010InterfacesWan1{}
	return &this
}

// NewInlineResponse20010InterfacesWan1WithDefaults instantiates a new InlineResponse20010InterfacesWan1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20010InterfacesWan1WithDefaults() *InlineResponse20010InterfacesWan1 {
	this := InlineResponse20010InterfacesWan1{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20010InterfacesWan1) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010InterfacesWan1) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20010InterfacesWan1) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20010InterfacesWan1) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVlanTagging returns the VlanTagging field value if set, zero value otherwise.
func (o *InlineResponse20010InterfacesWan1) GetVlanTagging() InlineResponse20010InterfacesWan1VlanTagging {
	if o == nil || isNil(o.VlanTagging) {
		var ret InlineResponse20010InterfacesWan1VlanTagging
		return ret
	}
	return *o.VlanTagging
}

// GetVlanTaggingOk returns a tuple with the VlanTagging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010InterfacesWan1) GetVlanTaggingOk() (*InlineResponse20010InterfacesWan1VlanTagging, bool) {
	if o == nil || isNil(o.VlanTagging) {
    return nil, false
	}
	return o.VlanTagging, true
}

// HasVlanTagging returns a boolean if a field has been set.
func (o *InlineResponse20010InterfacesWan1) HasVlanTagging() bool {
	if o != nil && !isNil(o.VlanTagging) {
		return true
	}

	return false
}

// SetVlanTagging gets a reference to the given InlineResponse20010InterfacesWan1VlanTagging and assigns it to the VlanTagging field.
func (o *InlineResponse20010InterfacesWan1) SetVlanTagging(v InlineResponse20010InterfacesWan1VlanTagging) {
	o.VlanTagging = &v
}

// GetSvis returns the Svis field value if set, zero value otherwise.
func (o *InlineResponse20010InterfacesWan1) GetSvis() InlineResponse20010InterfacesWan1Svis {
	if o == nil || isNil(o.Svis) {
		var ret InlineResponse20010InterfacesWan1Svis
		return ret
	}
	return *o.Svis
}

// GetSvisOk returns a tuple with the Svis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010InterfacesWan1) GetSvisOk() (*InlineResponse20010InterfacesWan1Svis, bool) {
	if o == nil || isNil(o.Svis) {
    return nil, false
	}
	return o.Svis, true
}

// HasSvis returns a boolean if a field has been set.
func (o *InlineResponse20010InterfacesWan1) HasSvis() bool {
	if o != nil && !isNil(o.Svis) {
		return true
	}

	return false
}

// SetSvis gets a reference to the given InlineResponse20010InterfacesWan1Svis and assigns it to the Svis field.
func (o *InlineResponse20010InterfacesWan1) SetSvis(v InlineResponse20010InterfacesWan1Svis) {
	o.Svis = &v
}

// GetPppoe returns the Pppoe field value if set, zero value otherwise.
func (o *InlineResponse20010InterfacesWan1) GetPppoe() InlineResponse20010InterfacesWan1Pppoe {
	if o == nil || isNil(o.Pppoe) {
		var ret InlineResponse20010InterfacesWan1Pppoe
		return ret
	}
	return *o.Pppoe
}

// GetPppoeOk returns a tuple with the Pppoe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010InterfacesWan1) GetPppoeOk() (*InlineResponse20010InterfacesWan1Pppoe, bool) {
	if o == nil || isNil(o.Pppoe) {
    return nil, false
	}
	return o.Pppoe, true
}

// HasPppoe returns a boolean if a field has been set.
func (o *InlineResponse20010InterfacesWan1) HasPppoe() bool {
	if o != nil && !isNil(o.Pppoe) {
		return true
	}

	return false
}

// SetPppoe gets a reference to the given InlineResponse20010InterfacesWan1Pppoe and assigns it to the Pppoe field.
func (o *InlineResponse20010InterfacesWan1) SetPppoe(v InlineResponse20010InterfacesWan1Pppoe) {
	o.Pppoe = &v
}

func (o InlineResponse20010InterfacesWan1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.VlanTagging) {
		toSerialize["vlanTagging"] = o.VlanTagging
	}
	if !isNil(o.Svis) {
		toSerialize["svis"] = o.Svis
	}
	if !isNil(o.Pppoe) {
		toSerialize["pppoe"] = o.Pppoe
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20010InterfacesWan1 struct {
	value *InlineResponse20010InterfacesWan1
	isSet bool
}

func (v NullableInlineResponse20010InterfacesWan1) Get() *InlineResponse20010InterfacesWan1 {
	return v.value
}

func (v *NullableInlineResponse20010InterfacesWan1) Set(val *InlineResponse20010InterfacesWan1) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20010InterfacesWan1) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20010InterfacesWan1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20010InterfacesWan1(val *InlineResponse20010InterfacesWan1) *NullableInlineResponse20010InterfacesWan1 {
	return &NullableInlineResponse20010InterfacesWan1{value: val, isSet: true}
}

func (v NullableInlineResponse20010InterfacesWan1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20010InterfacesWan1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


