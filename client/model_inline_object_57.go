/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject57 struct for InlineObject57
type InlineObject57 struct {
	GlobalBandwidthLimits *NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits `json:"globalBandwidthLimits,omitempty"`
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

// GetGlobalBandwidthLimits returns the GlobalBandwidthLimits field value if set, zero value otherwise.
func (o *InlineObject57) GetGlobalBandwidthLimits() NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits {
	if o == nil || isNil(o.GlobalBandwidthLimits) {
		var ret NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits
		return ret
	}
	return *o.GlobalBandwidthLimits
}

// GetGlobalBandwidthLimitsOk returns a tuple with the GlobalBandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject57) GetGlobalBandwidthLimitsOk() (*NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits, bool) {
	if o == nil || isNil(o.GlobalBandwidthLimits) {
    return nil, false
	}
	return o.GlobalBandwidthLimits, true
}

// HasGlobalBandwidthLimits returns a boolean if a field has been set.
func (o *InlineObject57) HasGlobalBandwidthLimits() bool {
	if o != nil && !isNil(o.GlobalBandwidthLimits) {
		return true
	}

	return false
}

// SetGlobalBandwidthLimits gets a reference to the given NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits and assigns it to the GlobalBandwidthLimits field.
func (o *InlineObject57) SetGlobalBandwidthLimits(v NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) {
	o.GlobalBandwidthLimits = &v
}

func (o InlineObject57) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GlobalBandwidthLimits) {
		toSerialize["globalBandwidthLimits"] = o.GlobalBandwidthLimits
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


