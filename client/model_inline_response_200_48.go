/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20048 struct for InlineResponse20048
type InlineResponse20048 struct {
	// The list of connectivity monitoring destinations
	Destinations []InlineResponse20048Destinations `json:"destinations,omitempty"`
}

// NewInlineResponse20048 instantiates a new InlineResponse20048 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20048() *InlineResponse20048 {
	this := InlineResponse20048{}
	return &this
}

// NewInlineResponse20048WithDefaults instantiates a new InlineResponse20048 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20048WithDefaults() *InlineResponse20048 {
	this := InlineResponse20048{}
	return &this
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *InlineResponse20048) GetDestinations() []InlineResponse20048Destinations {
	if o == nil || isNil(o.Destinations) {
		var ret []InlineResponse20048Destinations
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20048) GetDestinationsOk() ([]InlineResponse20048Destinations, bool) {
	if o == nil || isNil(o.Destinations) {
    return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *InlineResponse20048) HasDestinations() bool {
	if o != nil && !isNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []InlineResponse20048Destinations and assigns it to the Destinations field.
func (o *InlineResponse20048) SetDestinations(v []InlineResponse20048Destinations) {
	o.Destinations = v
}

func (o InlineResponse20048) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20048 struct {
	value *InlineResponse20048
	isSet bool
}

func (v NullableInlineResponse20048) Get() *InlineResponse20048 {
	return v.value
}

func (v *NullableInlineResponse20048) Set(val *InlineResponse20048) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20048) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20048) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20048(val *InlineResponse20048) *NullableInlineResponse20048 {
	return &NullableInlineResponse20048{value: val, isSet: true}
}

func (v NullableInlineResponse20048) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20048) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


