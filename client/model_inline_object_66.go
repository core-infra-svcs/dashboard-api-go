/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject66 struct for InlineObject66
type InlineObject66 struct {
	// The list of connectivity monitoring destinations
	Destinations []NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations `json:"destinations,omitempty"`
}

// NewInlineObject66 instantiates a new InlineObject66 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject66() *InlineObject66 {
	this := InlineObject66{}
	return &this
}

// NewInlineObject66WithDefaults instantiates a new InlineObject66 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject66WithDefaults() *InlineObject66 {
	this := InlineObject66{}
	return &this
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *InlineObject66) GetDestinations() []NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations {
	if o == nil || isNil(o.Destinations) {
		var ret []NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject66) GetDestinationsOk() ([]NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations, bool) {
	if o == nil || isNil(o.Destinations) {
    return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *InlineObject66) HasDestinations() bool {
	if o != nil && !isNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations and assigns it to the Destinations field.
func (o *InlineObject66) SetDestinations(v []NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) {
	o.Destinations = v
}

func (o InlineObject66) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject66 struct {
	value *InlineObject66
	isSet bool
}

func (v NullableInlineObject66) Get() *InlineObject66 {
	return v.value
}

func (v *NullableInlineObject66) Set(val *InlineObject66) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject66) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject66) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject66(val *InlineObject66) *NullableInlineObject66 {
	return &NullableInlineObject66{value: val, isSet: true}
}

func (v NullableInlineObject66) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject66) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


