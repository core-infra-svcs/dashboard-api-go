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

// InlineObject30 struct for InlineObject30
type InlineObject30 struct {
	// The list of connectivity monitoring destinations
	Destinations []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations `json:"destinations,omitempty"`
}

// NewInlineObject30 instantiates a new InlineObject30 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject30() *InlineObject30 {
	this := InlineObject30{}
	return &this
}

// NewInlineObject30WithDefaults instantiates a new InlineObject30 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject30WithDefaults() *InlineObject30 {
	this := InlineObject30{}
	return &this
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *InlineObject30) GetDestinations() []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations {
	if o == nil || isNil(o.Destinations) {
		var ret []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject30) GetDestinationsOk() ([]NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations, bool) {
	if o == nil || isNil(o.Destinations) {
    return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *InlineObject30) HasDestinations() bool {
	if o != nil && !isNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations and assigns it to the Destinations field.
func (o *InlineObject30) SetDestinations(v []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations) {
	o.Destinations = v
}

func (o InlineObject30) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject30 struct {
	value *InlineObject30
	isSet bool
}

func (v NullableInlineObject30) Get() *InlineObject30 {
	return v.value
}

func (v *NullableInlineObject30) Set(val *InlineObject30) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject30) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject30) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject30(val *InlineObject30) *NullableInlineObject30 {
	return &NullableInlineObject30{value: val, isSet: true}
}

func (v NullableInlineObject30) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject30) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


