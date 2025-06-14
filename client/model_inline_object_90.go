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

// InlineObject90 struct for InlineObject90
type InlineObject90 struct {
	BandwidthLimits *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular `json:"bandwidthLimits,omitempty"`
}

// NewInlineObject90 instantiates a new InlineObject90 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject90() *InlineObject90 {
	this := InlineObject90{}
	return &this
}

// NewInlineObject90WithDefaults instantiates a new InlineObject90 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject90WithDefaults() *InlineObject90 {
	this := InlineObject90{}
	return &this
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *InlineObject90) GetBandwidthLimits() NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular {
	if o == nil || isNil(o.BandwidthLimits) {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular
		return ret
	}
	return *o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject90) GetBandwidthLimitsOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular, bool) {
	if o == nil || isNil(o.BandwidthLimits) {
    return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *InlineObject90) HasBandwidthLimits() bool {
	if o != nil && !isNil(o.BandwidthLimits) {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular and assigns it to the BandwidthLimits field.
func (o *InlineObject90) SetBandwidthLimits(v NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular) {
	o.BandwidthLimits = &v
}

func (o InlineObject90) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BandwidthLimits) {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject90 struct {
	value *InlineObject90
	isSet bool
}

func (v NullableInlineObject90) Get() *InlineObject90 {
	return v.value
}

func (v *NullableInlineObject90) Set(val *InlineObject90) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject90) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject90) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject90(val *InlineObject90) *NullableInlineObject90 {
	return &NullableInlineObject90{value: val, isSet: true}
}

func (v NullableInlineObject90) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject90) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


