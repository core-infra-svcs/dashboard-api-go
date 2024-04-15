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

// InlineResponse200124 struct for InlineResponse200124
type InlineResponse200124 struct {
	// The list of restrictions for the device
	Restrictions []InlineResponse200124Restrictions `json:"restrictions,omitempty"`
}

// NewInlineResponse200124 instantiates a new InlineResponse200124 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200124() *InlineResponse200124 {
	this := InlineResponse200124{}
	return &this
}

// NewInlineResponse200124WithDefaults instantiates a new InlineResponse200124 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200124WithDefaults() *InlineResponse200124 {
	this := InlineResponse200124{}
	return &this
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *InlineResponse200124) GetRestrictions() []InlineResponse200124Restrictions {
	if o == nil || isNil(o.Restrictions) {
		var ret []InlineResponse200124Restrictions
		return ret
	}
	return o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124) GetRestrictionsOk() ([]InlineResponse200124Restrictions, bool) {
	if o == nil || isNil(o.Restrictions) {
    return nil, false
	}
	return o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *InlineResponse200124) HasRestrictions() bool {
	if o != nil && !isNil(o.Restrictions) {
		return true
	}

	return false
}

// SetRestrictions gets a reference to the given []InlineResponse200124Restrictions and assigns it to the Restrictions field.
func (o *InlineResponse200124) SetRestrictions(v []InlineResponse200124Restrictions) {
	o.Restrictions = v
}

func (o InlineResponse200124) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Restrictions) {
		toSerialize["restrictions"] = o.Restrictions
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200124 struct {
	value *InlineResponse200124
	isSet bool
}

func (v NullableInlineResponse200124) Get() *InlineResponse200124 {
	return v.value
}

func (v *NullableInlineResponse200124) Set(val *InlineResponse200124) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200124) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200124) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200124(val *InlineResponse200124) *NullableInlineResponse200124 {
	return &NullableInlineResponse200124{value: val, isSet: true}
}

func (v NullableInlineResponse200124) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200124) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


