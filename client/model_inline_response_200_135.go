/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200135 struct for InlineResponse200135
type InlineResponse200135 struct {
	// The list of restrictions for the device
	Restrictions []InlineResponse200135Restrictions `json:"restrictions,omitempty"`
}

// NewInlineResponse200135 instantiates a new InlineResponse200135 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200135() *InlineResponse200135 {
	this := InlineResponse200135{}
	return &this
}

// NewInlineResponse200135WithDefaults instantiates a new InlineResponse200135 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200135WithDefaults() *InlineResponse200135 {
	this := InlineResponse200135{}
	return &this
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *InlineResponse200135) GetRestrictions() []InlineResponse200135Restrictions {
	if o == nil || isNil(o.Restrictions) {
		var ret []InlineResponse200135Restrictions
		return ret
	}
	return o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetRestrictionsOk() ([]InlineResponse200135Restrictions, bool) {
	if o == nil || isNil(o.Restrictions) {
    return nil, false
	}
	return o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *InlineResponse200135) HasRestrictions() bool {
	if o != nil && !isNil(o.Restrictions) {
		return true
	}

	return false
}

// SetRestrictions gets a reference to the given []InlineResponse200135Restrictions and assigns it to the Restrictions field.
func (o *InlineResponse200135) SetRestrictions(v []InlineResponse200135Restrictions) {
	o.Restrictions = v
}

func (o InlineResponse200135) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Restrictions) {
		toSerialize["restrictions"] = o.Restrictions
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200135 struct {
	value *InlineResponse200135
	isSet bool
}

func (v NullableInlineResponse200135) Get() *InlineResponse200135 {
	return v.value
}

func (v *NullableInlineResponse200135) Set(val *InlineResponse200135) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200135) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200135) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200135(val *InlineResponse200135) *NullableInlineResponse200135 {
	return &NullableInlineResponse200135{value: val, isSet: true}
}

func (v NullableInlineResponse200135) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200135) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


