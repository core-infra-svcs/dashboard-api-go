/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2007InterfacesWan2 WAN 2 settings.
type InlineResponse2007InterfacesWan2 struct {
	// Enable or disable the interface.
	Enabled *bool `json:"enabled,omitempty"`
	VlanTagging *InlineResponse2007InterfacesWan1VlanTagging `json:"vlanTagging,omitempty"`
	Svis *InlineResponse2007InterfacesWan1Svis `json:"svis,omitempty"`
	Pppoe *InlineResponse2007InterfacesWan1Pppoe `json:"pppoe,omitempty"`
}

// NewInlineResponse2007InterfacesWan2 instantiates a new InlineResponse2007InterfacesWan2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2007InterfacesWan2() *InlineResponse2007InterfacesWan2 {
	this := InlineResponse2007InterfacesWan2{}
	return &this
}

// NewInlineResponse2007InterfacesWan2WithDefaults instantiates a new InlineResponse2007InterfacesWan2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2007InterfacesWan2WithDefaults() *InlineResponse2007InterfacesWan2 {
	this := InlineResponse2007InterfacesWan2{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse2007InterfacesWan2) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007InterfacesWan2) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse2007InterfacesWan2) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse2007InterfacesWan2) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVlanTagging returns the VlanTagging field value if set, zero value otherwise.
func (o *InlineResponse2007InterfacesWan2) GetVlanTagging() InlineResponse2007InterfacesWan1VlanTagging {
	if o == nil || isNil(o.VlanTagging) {
		var ret InlineResponse2007InterfacesWan1VlanTagging
		return ret
	}
	return *o.VlanTagging
}

// GetVlanTaggingOk returns a tuple with the VlanTagging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007InterfacesWan2) GetVlanTaggingOk() (*InlineResponse2007InterfacesWan1VlanTagging, bool) {
	if o == nil || isNil(o.VlanTagging) {
    return nil, false
	}
	return o.VlanTagging, true
}

// HasVlanTagging returns a boolean if a field has been set.
func (o *InlineResponse2007InterfacesWan2) HasVlanTagging() bool {
	if o != nil && !isNil(o.VlanTagging) {
		return true
	}

	return false
}

// SetVlanTagging gets a reference to the given InlineResponse2007InterfacesWan1VlanTagging and assigns it to the VlanTagging field.
func (o *InlineResponse2007InterfacesWan2) SetVlanTagging(v InlineResponse2007InterfacesWan1VlanTagging) {
	o.VlanTagging = &v
}

// GetSvis returns the Svis field value if set, zero value otherwise.
func (o *InlineResponse2007InterfacesWan2) GetSvis() InlineResponse2007InterfacesWan1Svis {
	if o == nil || isNil(o.Svis) {
		var ret InlineResponse2007InterfacesWan1Svis
		return ret
	}
	return *o.Svis
}

// GetSvisOk returns a tuple with the Svis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007InterfacesWan2) GetSvisOk() (*InlineResponse2007InterfacesWan1Svis, bool) {
	if o == nil || isNil(o.Svis) {
    return nil, false
	}
	return o.Svis, true
}

// HasSvis returns a boolean if a field has been set.
func (o *InlineResponse2007InterfacesWan2) HasSvis() bool {
	if o != nil && !isNil(o.Svis) {
		return true
	}

	return false
}

// SetSvis gets a reference to the given InlineResponse2007InterfacesWan1Svis and assigns it to the Svis field.
func (o *InlineResponse2007InterfacesWan2) SetSvis(v InlineResponse2007InterfacesWan1Svis) {
	o.Svis = &v
}

// GetPppoe returns the Pppoe field value if set, zero value otherwise.
func (o *InlineResponse2007InterfacesWan2) GetPppoe() InlineResponse2007InterfacesWan1Pppoe {
	if o == nil || isNil(o.Pppoe) {
		var ret InlineResponse2007InterfacesWan1Pppoe
		return ret
	}
	return *o.Pppoe
}

// GetPppoeOk returns a tuple with the Pppoe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007InterfacesWan2) GetPppoeOk() (*InlineResponse2007InterfacesWan1Pppoe, bool) {
	if o == nil || isNil(o.Pppoe) {
    return nil, false
	}
	return o.Pppoe, true
}

// HasPppoe returns a boolean if a field has been set.
func (o *InlineResponse2007InterfacesWan2) HasPppoe() bool {
	if o != nil && !isNil(o.Pppoe) {
		return true
	}

	return false
}

// SetPppoe gets a reference to the given InlineResponse2007InterfacesWan1Pppoe and assigns it to the Pppoe field.
func (o *InlineResponse2007InterfacesWan2) SetPppoe(v InlineResponse2007InterfacesWan1Pppoe) {
	o.Pppoe = &v
}

func (o InlineResponse2007InterfacesWan2) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse2007InterfacesWan2 struct {
	value *InlineResponse2007InterfacesWan2
	isSet bool
}

func (v NullableInlineResponse2007InterfacesWan2) Get() *InlineResponse2007InterfacesWan2 {
	return v.value
}

func (v *NullableInlineResponse2007InterfacesWan2) Set(val *InlineResponse2007InterfacesWan2) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2007InterfacesWan2) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2007InterfacesWan2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2007InterfacesWan2(val *InlineResponse2007InterfacesWan2) *NullableInlineResponse2007InterfacesWan2 {
	return &NullableInlineResponse2007InterfacesWan2{value: val, isSet: true}
}

func (v NullableInlineResponse2007InterfacesWan2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2007InterfacesWan2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


