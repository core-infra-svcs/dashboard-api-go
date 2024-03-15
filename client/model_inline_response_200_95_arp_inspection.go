/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20095ArpInspection Dynamic ARP Inspection settings
type InlineResponse20095ArpInspection struct {
	// Enable or disable Dynamic ARP Inspection on the network. Default value is false.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse20095ArpInspection instantiates a new InlineResponse20095ArpInspection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20095ArpInspection() *InlineResponse20095ArpInspection {
	this := InlineResponse20095ArpInspection{}
	return &this
}

// NewInlineResponse20095ArpInspectionWithDefaults instantiates a new InlineResponse20095ArpInspection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20095ArpInspectionWithDefaults() *InlineResponse20095ArpInspection {
	this := InlineResponse20095ArpInspection{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20095ArpInspection) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095ArpInspection) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20095ArpInspection) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20095ArpInspection) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse20095ArpInspection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20095ArpInspection struct {
	value *InlineResponse20095ArpInspection
	isSet bool
}

func (v NullableInlineResponse20095ArpInspection) Get() *InlineResponse20095ArpInspection {
	return v.value
}

func (v *NullableInlineResponse20095ArpInspection) Set(val *InlineResponse20095ArpInspection) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20095ArpInspection) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20095ArpInspection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20095ArpInspection(val *InlineResponse20095ArpInspection) *NullableInlineResponse20095ArpInspection {
	return &NullableInlineResponse20095ArpInspection{value: val, isSet: true}
}

func (v NullableInlineResponse20095ArpInspection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20095ArpInspection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


