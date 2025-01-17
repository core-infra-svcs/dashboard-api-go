/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2008InterfacesWan1 WAN 1 settings.
type InlineResponse2008InterfacesWan1 struct {
	// Enable or disable the interface.
	Enabled *bool `json:"enabled,omitempty"`
	VlanTagging *InlineResponse2008InterfacesWan1VlanTagging `json:"vlanTagging,omitempty"`
	Svis *InlineResponse2008InterfacesWan1Svis `json:"svis,omitempty"`
	Pppoe *InlineResponse2008InterfacesWan1Pppoe `json:"pppoe,omitempty"`
}

// NewInlineResponse2008InterfacesWan1 instantiates a new InlineResponse2008InterfacesWan1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2008InterfacesWan1() *InlineResponse2008InterfacesWan1 {
	this := InlineResponse2008InterfacesWan1{}
	return &this
}

// NewInlineResponse2008InterfacesWan1WithDefaults instantiates a new InlineResponse2008InterfacesWan1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2008InterfacesWan1WithDefaults() *InlineResponse2008InterfacesWan1 {
	this := InlineResponse2008InterfacesWan1{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse2008InterfacesWan1) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008InterfacesWan1) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse2008InterfacesWan1) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse2008InterfacesWan1) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVlanTagging returns the VlanTagging field value if set, zero value otherwise.
func (o *InlineResponse2008InterfacesWan1) GetVlanTagging() InlineResponse2008InterfacesWan1VlanTagging {
	if o == nil || isNil(o.VlanTagging) {
		var ret InlineResponse2008InterfacesWan1VlanTagging
		return ret
	}
	return *o.VlanTagging
}

// GetVlanTaggingOk returns a tuple with the VlanTagging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008InterfacesWan1) GetVlanTaggingOk() (*InlineResponse2008InterfacesWan1VlanTagging, bool) {
	if o == nil || isNil(o.VlanTagging) {
    return nil, false
	}
	return o.VlanTagging, true
}

// HasVlanTagging returns a boolean if a field has been set.
func (o *InlineResponse2008InterfacesWan1) HasVlanTagging() bool {
	if o != nil && !isNil(o.VlanTagging) {
		return true
	}

	return false
}

// SetVlanTagging gets a reference to the given InlineResponse2008InterfacesWan1VlanTagging and assigns it to the VlanTagging field.
func (o *InlineResponse2008InterfacesWan1) SetVlanTagging(v InlineResponse2008InterfacesWan1VlanTagging) {
	o.VlanTagging = &v
}

// GetSvis returns the Svis field value if set, zero value otherwise.
func (o *InlineResponse2008InterfacesWan1) GetSvis() InlineResponse2008InterfacesWan1Svis {
	if o == nil || isNil(o.Svis) {
		var ret InlineResponse2008InterfacesWan1Svis
		return ret
	}
	return *o.Svis
}

// GetSvisOk returns a tuple with the Svis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008InterfacesWan1) GetSvisOk() (*InlineResponse2008InterfacesWan1Svis, bool) {
	if o == nil || isNil(o.Svis) {
    return nil, false
	}
	return o.Svis, true
}

// HasSvis returns a boolean if a field has been set.
func (o *InlineResponse2008InterfacesWan1) HasSvis() bool {
	if o != nil && !isNil(o.Svis) {
		return true
	}

	return false
}

// SetSvis gets a reference to the given InlineResponse2008InterfacesWan1Svis and assigns it to the Svis field.
func (o *InlineResponse2008InterfacesWan1) SetSvis(v InlineResponse2008InterfacesWan1Svis) {
	o.Svis = &v
}

// GetPppoe returns the Pppoe field value if set, zero value otherwise.
func (o *InlineResponse2008InterfacesWan1) GetPppoe() InlineResponse2008InterfacesWan1Pppoe {
	if o == nil || isNil(o.Pppoe) {
		var ret InlineResponse2008InterfacesWan1Pppoe
		return ret
	}
	return *o.Pppoe
}

// GetPppoeOk returns a tuple with the Pppoe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008InterfacesWan1) GetPppoeOk() (*InlineResponse2008InterfacesWan1Pppoe, bool) {
	if o == nil || isNil(o.Pppoe) {
    return nil, false
	}
	return o.Pppoe, true
}

// HasPppoe returns a boolean if a field has been set.
func (o *InlineResponse2008InterfacesWan1) HasPppoe() bool {
	if o != nil && !isNil(o.Pppoe) {
		return true
	}

	return false
}

// SetPppoe gets a reference to the given InlineResponse2008InterfacesWan1Pppoe and assigns it to the Pppoe field.
func (o *InlineResponse2008InterfacesWan1) SetPppoe(v InlineResponse2008InterfacesWan1Pppoe) {
	o.Pppoe = &v
}

func (o InlineResponse2008InterfacesWan1) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse2008InterfacesWan1 struct {
	value *InlineResponse2008InterfacesWan1
	isSet bool
}

func (v NullableInlineResponse2008InterfacesWan1) Get() *InlineResponse2008InterfacesWan1 {
	return v.value
}

func (v *NullableInlineResponse2008InterfacesWan1) Set(val *InlineResponse2008InterfacesWan1) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2008InterfacesWan1) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2008InterfacesWan1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2008InterfacesWan1(val *InlineResponse2008InterfacesWan1) *NullableInlineResponse2008InterfacesWan1 {
	return &NullableInlineResponse2008InterfacesWan1{value: val, isSet: true}
}

func (v NullableInlineResponse2008InterfacesWan1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2008InterfacesWan1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


