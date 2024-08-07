/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20063BandwidthLimits A hash uplink keys and their configured settings for the Appliance
type InlineResponse20063BandwidthLimits struct {
	Wan1 *InlineResponse20063BandwidthLimitsWan1 `json:"wan1,omitempty"`
	Wan2 *InlineResponse20063BandwidthLimitsWan2 `json:"wan2,omitempty"`
	Cellular *InlineResponse20063BandwidthLimitsCellular `json:"cellular,omitempty"`
}

// NewInlineResponse20063BandwidthLimits instantiates a new InlineResponse20063BandwidthLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20063BandwidthLimits() *InlineResponse20063BandwidthLimits {
	this := InlineResponse20063BandwidthLimits{}
	return &this
}

// NewInlineResponse20063BandwidthLimitsWithDefaults instantiates a new InlineResponse20063BandwidthLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20063BandwidthLimitsWithDefaults() *InlineResponse20063BandwidthLimits {
	this := InlineResponse20063BandwidthLimits{}
	return &this
}

// GetWan1 returns the Wan1 field value if set, zero value otherwise.
func (o *InlineResponse20063BandwidthLimits) GetWan1() InlineResponse20063BandwidthLimitsWan1 {
	if o == nil || isNil(o.Wan1) {
		var ret InlineResponse20063BandwidthLimitsWan1
		return ret
	}
	return *o.Wan1
}

// GetWan1Ok returns a tuple with the Wan1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20063BandwidthLimits) GetWan1Ok() (*InlineResponse20063BandwidthLimitsWan1, bool) {
	if o == nil || isNil(o.Wan1) {
    return nil, false
	}
	return o.Wan1, true
}

// HasWan1 returns a boolean if a field has been set.
func (o *InlineResponse20063BandwidthLimits) HasWan1() bool {
	if o != nil && !isNil(o.Wan1) {
		return true
	}

	return false
}

// SetWan1 gets a reference to the given InlineResponse20063BandwidthLimitsWan1 and assigns it to the Wan1 field.
func (o *InlineResponse20063BandwidthLimits) SetWan1(v InlineResponse20063BandwidthLimitsWan1) {
	o.Wan1 = &v
}

// GetWan2 returns the Wan2 field value if set, zero value otherwise.
func (o *InlineResponse20063BandwidthLimits) GetWan2() InlineResponse20063BandwidthLimitsWan2 {
	if o == nil || isNil(o.Wan2) {
		var ret InlineResponse20063BandwidthLimitsWan2
		return ret
	}
	return *o.Wan2
}

// GetWan2Ok returns a tuple with the Wan2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20063BandwidthLimits) GetWan2Ok() (*InlineResponse20063BandwidthLimitsWan2, bool) {
	if o == nil || isNil(o.Wan2) {
    return nil, false
	}
	return o.Wan2, true
}

// HasWan2 returns a boolean if a field has been set.
func (o *InlineResponse20063BandwidthLimits) HasWan2() bool {
	if o != nil && !isNil(o.Wan2) {
		return true
	}

	return false
}

// SetWan2 gets a reference to the given InlineResponse20063BandwidthLimitsWan2 and assigns it to the Wan2 field.
func (o *InlineResponse20063BandwidthLimits) SetWan2(v InlineResponse20063BandwidthLimitsWan2) {
	o.Wan2 = &v
}

// GetCellular returns the Cellular field value if set, zero value otherwise.
func (o *InlineResponse20063BandwidthLimits) GetCellular() InlineResponse20063BandwidthLimitsCellular {
	if o == nil || isNil(o.Cellular) {
		var ret InlineResponse20063BandwidthLimitsCellular
		return ret
	}
	return *o.Cellular
}

// GetCellularOk returns a tuple with the Cellular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20063BandwidthLimits) GetCellularOk() (*InlineResponse20063BandwidthLimitsCellular, bool) {
	if o == nil || isNil(o.Cellular) {
    return nil, false
	}
	return o.Cellular, true
}

// HasCellular returns a boolean if a field has been set.
func (o *InlineResponse20063BandwidthLimits) HasCellular() bool {
	if o != nil && !isNil(o.Cellular) {
		return true
	}

	return false
}

// SetCellular gets a reference to the given InlineResponse20063BandwidthLimitsCellular and assigns it to the Cellular field.
func (o *InlineResponse20063BandwidthLimits) SetCellular(v InlineResponse20063BandwidthLimitsCellular) {
	o.Cellular = &v
}

func (o InlineResponse20063BandwidthLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Wan1) {
		toSerialize["wan1"] = o.Wan1
	}
	if !isNil(o.Wan2) {
		toSerialize["wan2"] = o.Wan2
	}
	if !isNil(o.Cellular) {
		toSerialize["cellular"] = o.Cellular
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20063BandwidthLimits struct {
	value *InlineResponse20063BandwidthLimits
	isSet bool
}

func (v NullableInlineResponse20063BandwidthLimits) Get() *InlineResponse20063BandwidthLimits {
	return v.value
}

func (v *NullableInlineResponse20063BandwidthLimits) Set(val *InlineResponse20063BandwidthLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20063BandwidthLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20063BandwidthLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20063BandwidthLimits(val *InlineResponse20063BandwidthLimits) *NullableInlineResponse20063BandwidthLimits {
	return &NullableInlineResponse20063BandwidthLimits{value: val, isSet: true}
}

func (v NullableInlineResponse20063BandwidthLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20063BandwidthLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


