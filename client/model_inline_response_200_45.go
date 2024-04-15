/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20045 struct for InlineResponse20045
type InlineResponse20045 struct {
	// The list of connectivity monitoring destinations
	Destinations []InlineResponse20045Destinations `json:"destinations,omitempty"`
}

// NewInlineResponse20045 instantiates a new InlineResponse20045 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20045() *InlineResponse20045 {
	this := InlineResponse20045{}
	return &this
}

// NewInlineResponse20045WithDefaults instantiates a new InlineResponse20045 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20045WithDefaults() *InlineResponse20045 {
	this := InlineResponse20045{}
	return &this
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *InlineResponse20045) GetDestinations() []InlineResponse20045Destinations {
	if o == nil || isNil(o.Destinations) {
		var ret []InlineResponse20045Destinations
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20045) GetDestinationsOk() ([]InlineResponse20045Destinations, bool) {
	if o == nil || isNil(o.Destinations) {
    return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *InlineResponse20045) HasDestinations() bool {
	if o != nil && !isNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []InlineResponse20045Destinations and assigns it to the Destinations field.
func (o *InlineResponse20045) SetDestinations(v []InlineResponse20045Destinations) {
	o.Destinations = v
}

func (o InlineResponse20045) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20045 struct {
	value *InlineResponse20045
	isSet bool
}

func (v NullableInlineResponse20045) Get() *InlineResponse20045 {
	return v.value
}

func (v *NullableInlineResponse20045) Set(val *InlineResponse20045) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20045) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20045) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20045(val *InlineResponse20045) *NullableInlineResponse20045 {
	return &NullableInlineResponse20045{value: val, isSet: true}
}

func (v NullableInlineResponse20045) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20045) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


