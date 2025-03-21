/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200343Details struct for InlineResponse200343Details
type InlineResponse200343Details struct {
	// Item name
	Name *string `json:"name,omitempty"`
	// Item value
	Value *string `json:"value,omitempty"`
}

// NewInlineResponse200343Details instantiates a new InlineResponse200343Details object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200343Details() *InlineResponse200343Details {
	this := InlineResponse200343Details{}
	return &this
}

// NewInlineResponse200343DetailsWithDefaults instantiates a new InlineResponse200343Details object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200343DetailsWithDefaults() *InlineResponse200343Details {
	this := InlineResponse200343Details{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200343Details) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200343Details) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200343Details) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200343Details) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InlineResponse200343Details) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200343Details) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineResponse200343Details) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InlineResponse200343Details) SetValue(v string) {
	o.Value = &v
}

func (o InlineResponse200343Details) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200343Details struct {
	value *InlineResponse200343Details
	isSet bool
}

func (v NullableInlineResponse200343Details) Get() *InlineResponse200343Details {
	return v.value
}

func (v *NullableInlineResponse200343Details) Set(val *InlineResponse200343Details) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200343Details) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200343Details) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200343Details(val *InlineResponse200343Details) *NullableInlineResponse200343Details {
	return &NullableInlineResponse200343Details{value: val, isSet: true}
}

func (v NullableInlineResponse200343Details) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200343Details) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


