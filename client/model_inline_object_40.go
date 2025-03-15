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

// InlineObject40 struct for InlineObject40
type InlineObject40 struct {
	// The list of connectivity monitoring destinations
	Destinations []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations `json:"destinations,omitempty"`
}

// NewInlineObject40 instantiates a new InlineObject40 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject40() *InlineObject40 {
	this := InlineObject40{}
	return &this
}

// NewInlineObject40WithDefaults instantiates a new InlineObject40 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject40WithDefaults() *InlineObject40 {
	this := InlineObject40{}
	return &this
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *InlineObject40) GetDestinations() []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations {
	if o == nil || isNil(o.Destinations) {
		var ret []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject40) GetDestinationsOk() ([]NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations, bool) {
	if o == nil || isNil(o.Destinations) {
    return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *InlineObject40) HasDestinations() bool {
	if o != nil && !isNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations and assigns it to the Destinations field.
func (o *InlineObject40) SetDestinations(v []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations) {
	o.Destinations = v
}

func (o InlineObject40) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject40 struct {
	value *InlineObject40
	isSet bool
}

func (v NullableInlineObject40) Get() *InlineObject40 {
	return v.value
}

func (v *NullableInlineObject40) Set(val *InlineObject40) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject40) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject40) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject40(val *InlineObject40) *NullableInlineObject40 {
	return &NullableInlineObject40{value: val, isSet: true}
}

func (v NullableInlineObject40) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject40) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


