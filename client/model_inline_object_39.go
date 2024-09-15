/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject39 struct for InlineObject39
type InlineObject39 struct {
	// The list of connectivity monitoring destinations
	Destinations []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations `json:"destinations,omitempty"`
}

// NewInlineObject39 instantiates a new InlineObject39 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject39() *InlineObject39 {
	this := InlineObject39{}
	return &this
}

// NewInlineObject39WithDefaults instantiates a new InlineObject39 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject39WithDefaults() *InlineObject39 {
	this := InlineObject39{}
	return &this
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *InlineObject39) GetDestinations() []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations {
	if o == nil || isNil(o.Destinations) {
		var ret []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject39) GetDestinationsOk() ([]NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations, bool) {
	if o == nil || isNil(o.Destinations) {
    return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *InlineObject39) HasDestinations() bool {
	if o != nil && !isNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations and assigns it to the Destinations field.
func (o *InlineObject39) SetDestinations(v []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations) {
	o.Destinations = v
}

func (o InlineObject39) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject39 struct {
	value *InlineObject39
	isSet bool
}

func (v NullableInlineObject39) Get() *InlineObject39 {
	return v.value
}

func (v *NullableInlineObject39) Set(val *InlineObject39) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject39) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject39) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject39(val *InlineObject39) *NullableInlineObject39 {
	return &NullableInlineObject39{value: val, isSet: true}
}

func (v NullableInlineObject39) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject39) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


