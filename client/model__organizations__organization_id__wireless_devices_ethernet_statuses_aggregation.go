/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation Aggregation details object
type OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation struct {
	// Link Aggregation enabled flag
	Enabled *bool `json:"enabled,omitempty"`
	// Link Aggregation speed
	Speed *int32 `json:"speed,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation instantiates a new OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation {
	this := OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregationWithDefaults instantiates a new OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregationWithDefaults() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation {
	this := OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) GetSpeed() int32 {
	if o == nil || isNil(o.Speed) {
		var ret int32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) GetSpeedOk() (*int32, bool) {
	if o == nil || isNil(o.Speed) {
    return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) HasSpeed() bool {
	if o != nil && !isNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int32 and assigns it to the Speed field.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) SetSpeed(v int32) {
	o.Speed = &v
}

func (o OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation struct {
	value *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) Get() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) Set(val *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation(val *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation {
	return &NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


