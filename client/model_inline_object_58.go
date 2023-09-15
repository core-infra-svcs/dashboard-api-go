/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject58 struct for InlineObject58
type InlineObject58 struct {
	BandwidthLimits *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits `json:"bandwidthLimits,omitempty"`
}

// NewInlineObject58 instantiates a new InlineObject58 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject58() *InlineObject58 {
	this := InlineObject58{}
	return &this
}

// NewInlineObject58WithDefaults instantiates a new InlineObject58 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject58WithDefaults() *InlineObject58 {
	this := InlineObject58{}
	return &this
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *InlineObject58) GetBandwidthLimits() NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits {
	if o == nil || isNil(o.BandwidthLimits) {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits
		return ret
	}
	return *o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject58) GetBandwidthLimitsOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits, bool) {
	if o == nil || isNil(o.BandwidthLimits) {
    return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *InlineObject58) HasBandwidthLimits() bool {
	if o != nil && !isNil(o.BandwidthLimits) {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits and assigns it to the BandwidthLimits field.
func (o *InlineObject58) SetBandwidthLimits(v NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) {
	o.BandwidthLimits = &v
}

func (o InlineObject58) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BandwidthLimits) {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject58 struct {
	value *InlineObject58
	isSet bool
}

func (v NullableInlineObject58) Get() *InlineObject58 {
	return v.value
}

func (v *NullableInlineObject58) Set(val *InlineObject58) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject58) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject58) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject58(val *InlineObject58) *NullableInlineObject58 {
	return &NullableInlineObject58{value: val, isSet: true}
}

func (v NullableInlineObject58) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject58) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


