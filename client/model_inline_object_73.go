/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject73 struct for InlineObject73
type InlineObject73 struct {
	BandwidthLimits *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular `json:"bandwidthLimits,omitempty"`
}

// NewInlineObject73 instantiates a new InlineObject73 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject73() *InlineObject73 {
	this := InlineObject73{}
	return &this
}

// NewInlineObject73WithDefaults instantiates a new InlineObject73 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject73WithDefaults() *InlineObject73 {
	this := InlineObject73{}
	return &this
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *InlineObject73) GetBandwidthLimits() NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular {
	if o == nil || isNil(o.BandwidthLimits) {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular
		return ret
	}
	return *o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject73) GetBandwidthLimitsOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular, bool) {
	if o == nil || isNil(o.BandwidthLimits) {
    return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *InlineObject73) HasBandwidthLimits() bool {
	if o != nil && !isNil(o.BandwidthLimits) {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular and assigns it to the BandwidthLimits field.
func (o *InlineObject73) SetBandwidthLimits(v NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular) {
	o.BandwidthLimits = &v
}

func (o InlineObject73) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BandwidthLimits) {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject73 struct {
	value *InlineObject73
	isSet bool
}

func (v NullableInlineObject73) Get() *InlineObject73 {
	return v.value
}

func (v *NullableInlineObject73) Set(val *InlineObject73) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject73) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject73) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject73(val *InlineObject73) *NullableInlineObject73 {
	return &NullableInlineObject73{value: val, isSet: true}
}

func (v NullableInlineObject73) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject73) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


