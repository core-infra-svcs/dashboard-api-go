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

// InlineObject125 struct for InlineObject125
type InlineObject125 struct {
	// The name of this target group
	Name *string `json:"name,omitempty"`
	// The scope and tag options of the target group. Comma separated values beginning with one of withAny, withAll, withoutAny, withoutAll, all, none, followed by tags. Default to none if empty.
	Scope *string `json:"scope,omitempty"`
}

// NewInlineObject125 instantiates a new InlineObject125 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject125() *InlineObject125 {
	this := InlineObject125{}
	return &this
}

// NewInlineObject125WithDefaults instantiates a new InlineObject125 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject125WithDefaults() *InlineObject125 {
	this := InlineObject125{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject125) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject125) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject125) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject125) SetName(v string) {
	o.Name = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineObject125) GetScope() string {
	if o == nil || isNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject125) GetScopeOk() (*string, bool) {
	if o == nil || isNil(o.Scope) {
    return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineObject125) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *InlineObject125) SetScope(v string) {
	o.Scope = &v
}

func (o InlineObject125) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject125 struct {
	value *InlineObject125
	isSet bool
}

func (v NullableInlineObject125) Get() *InlineObject125 {
	return v.value
}

func (v *NullableInlineObject125) Set(val *InlineObject125) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject125) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject125) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject125(val *InlineObject125) *NullableInlineObject125 {
	return &NullableInlineObject125{value: val, isSet: true}
}

func (v NullableInlineObject125) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject125) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


