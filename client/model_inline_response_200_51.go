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

// InlineResponse20051 struct for InlineResponse20051
type InlineResponse20051 struct {
	// The list of connectivity monitoring destinations
	Destinations []InlineResponse20051Destinations `json:"destinations,omitempty"`
}

// NewInlineResponse20051 instantiates a new InlineResponse20051 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20051() *InlineResponse20051 {
	this := InlineResponse20051{}
	return &this
}

// NewInlineResponse20051WithDefaults instantiates a new InlineResponse20051 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20051WithDefaults() *InlineResponse20051 {
	this := InlineResponse20051{}
	return &this
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *InlineResponse20051) GetDestinations() []InlineResponse20051Destinations {
	if o == nil || isNil(o.Destinations) {
		var ret []InlineResponse20051Destinations
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20051) GetDestinationsOk() ([]InlineResponse20051Destinations, bool) {
	if o == nil || isNil(o.Destinations) {
    return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *InlineResponse20051) HasDestinations() bool {
	if o != nil && !isNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []InlineResponse20051Destinations and assigns it to the Destinations field.
func (o *InlineResponse20051) SetDestinations(v []InlineResponse20051Destinations) {
	o.Destinations = v
}

func (o InlineResponse20051) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20051 struct {
	value *InlineResponse20051
	isSet bool
}

func (v NullableInlineResponse20051) Get() *InlineResponse20051 {
	return v.value
}

func (v *NullableInlineResponse20051) Set(val *InlineResponse20051) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20051) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20051) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20051(val *InlineResponse20051) *NullableInlineResponse20051 {
	return &NullableInlineResponse20051{value: val, isSet: true}
}

func (v NullableInlineResponse20051) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20051) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


