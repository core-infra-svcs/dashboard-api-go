/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject82 struct for InlineObject82
type InlineObject82 struct {
	// The list of connectivity monitoring destinations
	Destinations []NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations `json:"destinations,omitempty"`
}

// NewInlineObject82 instantiates a new InlineObject82 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject82() *InlineObject82 {
	this := InlineObject82{}
	return &this
}

// NewInlineObject82WithDefaults instantiates a new InlineObject82 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject82WithDefaults() *InlineObject82 {
	this := InlineObject82{}
	return &this
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *InlineObject82) GetDestinations() []NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations {
	if o == nil || isNil(o.Destinations) {
		var ret []NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject82) GetDestinationsOk() ([]NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations, bool) {
	if o == nil || isNil(o.Destinations) {
    return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *InlineObject82) HasDestinations() bool {
	if o != nil && !isNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations and assigns it to the Destinations field.
func (o *InlineObject82) SetDestinations(v []NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) {
	o.Destinations = v
}

func (o InlineObject82) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject82 struct {
	value *InlineObject82
	isSet bool
}

func (v NullableInlineObject82) Get() *InlineObject82 {
	return v.value
}

func (v *NullableInlineObject82) Set(val *InlineObject82) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject82) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject82) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject82(val *InlineObject82) *NullableInlineObject82 {
	return &NullableInlineObject82{value: val, isSet: true}
}

func (v NullableInlineObject82) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject82) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


