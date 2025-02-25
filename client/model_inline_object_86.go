/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject86 struct for InlineObject86
type InlineObject86 struct {
	BandwidthLimits *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular `json:"bandwidthLimits,omitempty"`
}

// NewInlineObject86 instantiates a new InlineObject86 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject86() *InlineObject86 {
	this := InlineObject86{}
	return &this
}

// NewInlineObject86WithDefaults instantiates a new InlineObject86 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject86WithDefaults() *InlineObject86 {
	this := InlineObject86{}
	return &this
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *InlineObject86) GetBandwidthLimits() NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular {
	if o == nil || isNil(o.BandwidthLimits) {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular
		return ret
	}
	return *o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject86) GetBandwidthLimitsOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular, bool) {
	if o == nil || isNil(o.BandwidthLimits) {
    return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *InlineObject86) HasBandwidthLimits() bool {
	if o != nil && !isNil(o.BandwidthLimits) {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular and assigns it to the BandwidthLimits field.
func (o *InlineObject86) SetBandwidthLimits(v NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular) {
	o.BandwidthLimits = &v
}

func (o InlineObject86) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BandwidthLimits) {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject86 struct {
	value *InlineObject86
	isSet bool
}

func (v NullableInlineObject86) Get() *InlineObject86 {
	return v.value
}

func (v *NullableInlineObject86) Set(val *InlineObject86) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject86) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject86) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject86(val *InlineObject86) *NullableInlineObject86 {
	return &NullableInlineObject86{value: val, isSet: true}
}

func (v NullableInlineObject86) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject86) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


