/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20038 struct for InlineResponse20038
type InlineResponse20038 struct {
	// Enable or disable warm spare for a switch
	Enabled *bool `json:"enabled,omitempty"`
	// Serial number of the primary switch
	PrimarySerial *string `json:"primarySerial,omitempty"`
	// Serial number of the warm spare switch
	SpareSerial *string `json:"spareSerial,omitempty"`
}

// NewInlineResponse20038 instantiates a new InlineResponse20038 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20038() *InlineResponse20038 {
	this := InlineResponse20038{}
	return &this
}

// NewInlineResponse20038WithDefaults instantiates a new InlineResponse20038 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20038WithDefaults() *InlineResponse20038 {
	this := InlineResponse20038{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20038) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20038) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20038) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20038) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrimarySerial returns the PrimarySerial field value if set, zero value otherwise.
func (o *InlineResponse20038) GetPrimarySerial() string {
	if o == nil || isNil(o.PrimarySerial) {
		var ret string
		return ret
	}
	return *o.PrimarySerial
}

// GetPrimarySerialOk returns a tuple with the PrimarySerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20038) GetPrimarySerialOk() (*string, bool) {
	if o == nil || isNil(o.PrimarySerial) {
    return nil, false
	}
	return o.PrimarySerial, true
}

// HasPrimarySerial returns a boolean if a field has been set.
func (o *InlineResponse20038) HasPrimarySerial() bool {
	if o != nil && !isNil(o.PrimarySerial) {
		return true
	}

	return false
}

// SetPrimarySerial gets a reference to the given string and assigns it to the PrimarySerial field.
func (o *InlineResponse20038) SetPrimarySerial(v string) {
	o.PrimarySerial = &v
}

// GetSpareSerial returns the SpareSerial field value if set, zero value otherwise.
func (o *InlineResponse20038) GetSpareSerial() string {
	if o == nil || isNil(o.SpareSerial) {
		var ret string
		return ret
	}
	return *o.SpareSerial
}

// GetSpareSerialOk returns a tuple with the SpareSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20038) GetSpareSerialOk() (*string, bool) {
	if o == nil || isNil(o.SpareSerial) {
    return nil, false
	}
	return o.SpareSerial, true
}

// HasSpareSerial returns a boolean if a field has been set.
func (o *InlineResponse20038) HasSpareSerial() bool {
	if o != nil && !isNil(o.SpareSerial) {
		return true
	}

	return false
}

// SetSpareSerial gets a reference to the given string and assigns it to the SpareSerial field.
func (o *InlineResponse20038) SetSpareSerial(v string) {
	o.SpareSerial = &v
}

func (o InlineResponse20038) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.PrimarySerial) {
		toSerialize["primarySerial"] = o.PrimarySerial
	}
	if !isNil(o.SpareSerial) {
		toSerialize["spareSerial"] = o.SpareSerial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20038 struct {
	value *InlineResponse20038
	isSet bool
}

func (v NullableInlineResponse20038) Get() *InlineResponse20038 {
	return v.value
}

func (v *NullableInlineResponse20038) Set(val *InlineResponse20038) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20038) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20038) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20038(val *InlineResponse20038) *NullableInlineResponse20038 {
	return &NullableInlineResponse20038{value: val, isSet: true}
}

func (v NullableInlineResponse20038) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20038) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


