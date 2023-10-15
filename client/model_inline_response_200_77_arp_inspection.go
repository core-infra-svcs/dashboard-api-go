/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20077ArpInspection Dynamic ARP Inspection settings
type InlineResponse20077ArpInspection struct {
	// Enable or disable Dynamic ARP Inspection on the network. Default value is false.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse20077ArpInspection instantiates a new InlineResponse20077ArpInspection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20077ArpInspection() *InlineResponse20077ArpInspection {
	this := InlineResponse20077ArpInspection{}
	return &this
}

// NewInlineResponse20077ArpInspectionWithDefaults instantiates a new InlineResponse20077ArpInspection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20077ArpInspectionWithDefaults() *InlineResponse20077ArpInspection {
	this := InlineResponse20077ArpInspection{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20077ArpInspection) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20077ArpInspection) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20077ArpInspection) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20077ArpInspection) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse20077ArpInspection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20077ArpInspection struct {
	value *InlineResponse20077ArpInspection
	isSet bool
}

func (v NullableInlineResponse20077ArpInspection) Get() *InlineResponse20077ArpInspection {
	return v.value
}

func (v *NullableInlineResponse20077ArpInspection) Set(val *InlineResponse20077ArpInspection) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20077ArpInspection) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20077ArpInspection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20077ArpInspection(val *InlineResponse20077ArpInspection) *NullableInlineResponse20077ArpInspection {
	return &NullableInlineResponse20077ArpInspection{value: val, isSet: true}
}

func (v NullableInlineResponse20077ArpInspection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20077ArpInspection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


