/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20068BandwidthLimits A hash uplink keys and their configured settings for the Appliance
type InlineResponse20068BandwidthLimits struct {
	Wan1 *InlineResponse20068BandwidthLimitsWan1 `json:"wan1,omitempty"`
	Wan2 *InlineResponse20068BandwidthLimitsWan2 `json:"wan2,omitempty"`
	Cellular *InlineResponse20068BandwidthLimitsCellular `json:"cellular,omitempty"`
}

// NewInlineResponse20068BandwidthLimits instantiates a new InlineResponse20068BandwidthLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20068BandwidthLimits() *InlineResponse20068BandwidthLimits {
	this := InlineResponse20068BandwidthLimits{}
	return &this
}

// NewInlineResponse20068BandwidthLimitsWithDefaults instantiates a new InlineResponse20068BandwidthLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20068BandwidthLimitsWithDefaults() *InlineResponse20068BandwidthLimits {
	this := InlineResponse20068BandwidthLimits{}
	return &this
}

// GetWan1 returns the Wan1 field value if set, zero value otherwise.
func (o *InlineResponse20068BandwidthLimits) GetWan1() InlineResponse20068BandwidthLimitsWan1 {
	if o == nil || isNil(o.Wan1) {
		var ret InlineResponse20068BandwidthLimitsWan1
		return ret
	}
	return *o.Wan1
}

// GetWan1Ok returns a tuple with the Wan1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068BandwidthLimits) GetWan1Ok() (*InlineResponse20068BandwidthLimitsWan1, bool) {
	if o == nil || isNil(o.Wan1) {
    return nil, false
	}
	return o.Wan1, true
}

// HasWan1 returns a boolean if a field has been set.
func (o *InlineResponse20068BandwidthLimits) HasWan1() bool {
	if o != nil && !isNil(o.Wan1) {
		return true
	}

	return false
}

// SetWan1 gets a reference to the given InlineResponse20068BandwidthLimitsWan1 and assigns it to the Wan1 field.
func (o *InlineResponse20068BandwidthLimits) SetWan1(v InlineResponse20068BandwidthLimitsWan1) {
	o.Wan1 = &v
}

// GetWan2 returns the Wan2 field value if set, zero value otherwise.
func (o *InlineResponse20068BandwidthLimits) GetWan2() InlineResponse20068BandwidthLimitsWan2 {
	if o == nil || isNil(o.Wan2) {
		var ret InlineResponse20068BandwidthLimitsWan2
		return ret
	}
	return *o.Wan2
}

// GetWan2Ok returns a tuple with the Wan2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068BandwidthLimits) GetWan2Ok() (*InlineResponse20068BandwidthLimitsWan2, bool) {
	if o == nil || isNil(o.Wan2) {
    return nil, false
	}
	return o.Wan2, true
}

// HasWan2 returns a boolean if a field has been set.
func (o *InlineResponse20068BandwidthLimits) HasWan2() bool {
	if o != nil && !isNil(o.Wan2) {
		return true
	}

	return false
}

// SetWan2 gets a reference to the given InlineResponse20068BandwidthLimitsWan2 and assigns it to the Wan2 field.
func (o *InlineResponse20068BandwidthLimits) SetWan2(v InlineResponse20068BandwidthLimitsWan2) {
	o.Wan2 = &v
}

// GetCellular returns the Cellular field value if set, zero value otherwise.
func (o *InlineResponse20068BandwidthLimits) GetCellular() InlineResponse20068BandwidthLimitsCellular {
	if o == nil || isNil(o.Cellular) {
		var ret InlineResponse20068BandwidthLimitsCellular
		return ret
	}
	return *o.Cellular
}

// GetCellularOk returns a tuple with the Cellular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068BandwidthLimits) GetCellularOk() (*InlineResponse20068BandwidthLimitsCellular, bool) {
	if o == nil || isNil(o.Cellular) {
    return nil, false
	}
	return o.Cellular, true
}

// HasCellular returns a boolean if a field has been set.
func (o *InlineResponse20068BandwidthLimits) HasCellular() bool {
	if o != nil && !isNil(o.Cellular) {
		return true
	}

	return false
}

// SetCellular gets a reference to the given InlineResponse20068BandwidthLimitsCellular and assigns it to the Cellular field.
func (o *InlineResponse20068BandwidthLimits) SetCellular(v InlineResponse20068BandwidthLimitsCellular) {
	o.Cellular = &v
}

func (o InlineResponse20068BandwidthLimits) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20068BandwidthLimits struct {
	value *InlineResponse20068BandwidthLimits
	isSet bool
}

func (v NullableInlineResponse20068BandwidthLimits) Get() *InlineResponse20068BandwidthLimits {
	return v.value
}

func (v *NullableInlineResponse20068BandwidthLimits) Set(val *InlineResponse20068BandwidthLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20068BandwidthLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20068BandwidthLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20068BandwidthLimits(val *InlineResponse20068BandwidthLimits) *NullableInlineResponse20068BandwidthLimits {
	return &NullableInlineResponse20068BandwidthLimits{value: val, isSet: true}
}

func (v NullableInlineResponse20068BandwidthLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20068BandwidthLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


