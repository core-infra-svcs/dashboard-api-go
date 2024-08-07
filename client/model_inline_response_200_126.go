/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200126 struct for InlineResponse200126
type InlineResponse200126 struct {
	// The list of restrictions for the device
	Restrictions []InlineResponse200126Restrictions `json:"restrictions,omitempty"`
}

// NewInlineResponse200126 instantiates a new InlineResponse200126 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200126() *InlineResponse200126 {
	this := InlineResponse200126{}
	return &this
}

// NewInlineResponse200126WithDefaults instantiates a new InlineResponse200126 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200126WithDefaults() *InlineResponse200126 {
	this := InlineResponse200126{}
	return &this
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *InlineResponse200126) GetRestrictions() []InlineResponse200126Restrictions {
	if o == nil || isNil(o.Restrictions) {
		var ret []InlineResponse200126Restrictions
		return ret
	}
	return o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200126) GetRestrictionsOk() ([]InlineResponse200126Restrictions, bool) {
	if o == nil || isNil(o.Restrictions) {
    return nil, false
	}
	return o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *InlineResponse200126) HasRestrictions() bool {
	if o != nil && !isNil(o.Restrictions) {
		return true
	}

	return false
}

// SetRestrictions gets a reference to the given []InlineResponse200126Restrictions and assigns it to the Restrictions field.
func (o *InlineResponse200126) SetRestrictions(v []InlineResponse200126Restrictions) {
	o.Restrictions = v
}

func (o InlineResponse200126) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Restrictions) {
		toSerialize["restrictions"] = o.Restrictions
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200126 struct {
	value *InlineResponse200126
	isSet bool
}

func (v NullableInlineResponse200126) Get() *InlineResponse200126 {
	return v.value
}

func (v *NullableInlineResponse200126) Set(val *InlineResponse200126) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200126) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200126) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200126(val *InlineResponse200126) *NullableInlineResponse200126 {
	return &NullableInlineResponse200126{value: val, isSet: true}
}

func (v NullableInlineResponse200126) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200126) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


