/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject61 struct for InlineObject61
type InlineObject61 struct {
	BandwidthLimits *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular `json:"bandwidthLimits,omitempty"`
}

// NewInlineObject61 instantiates a new InlineObject61 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject61() *InlineObject61 {
	this := InlineObject61{}
	return &this
}

// NewInlineObject61WithDefaults instantiates a new InlineObject61 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject61WithDefaults() *InlineObject61 {
	this := InlineObject61{}
	return &this
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *InlineObject61) GetBandwidthLimits() NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular {
	if o == nil || o.BandwidthLimits == nil {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular
		return ret
	}
	return *o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject61) GetBandwidthLimitsOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular, bool) {
	if o == nil || o.BandwidthLimits == nil {
		return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *InlineObject61) HasBandwidthLimits() bool {
	if o != nil && o.BandwidthLimits != nil {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular and assigns it to the BandwidthLimits field.
func (o *InlineObject61) SetBandwidthLimits(v NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular) {
	o.BandwidthLimits = &v
}

func (o InlineObject61) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BandwidthLimits != nil {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject61 struct {
	value *InlineObject61
	isSet bool
}

func (v NullableInlineObject61) Get() *InlineObject61 {
	return v.value
}

func (v *NullableInlineObject61) Set(val *InlineObject61) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject61) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject61) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject61(val *InlineObject61) *NullableInlineObject61 {
	return &NullableInlineObject61{value: val, isSet: true}
}

func (v NullableInlineObject61) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject61) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


