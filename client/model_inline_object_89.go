/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject89 struct for InlineObject89
type InlineObject89 struct {
	// The name of this target group
	Name *string `json:"name,omitempty"`
	// The scope and tag options of the target group. Comma separated values beginning with one of withAny, withAll, withoutAny, withoutAll, all, none, followed by tags. Default to none if empty.
	Scope *string `json:"scope,omitempty"`
}

// NewInlineObject89 instantiates a new InlineObject89 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject89() *InlineObject89 {
	this := InlineObject89{}
	return &this
}

// NewInlineObject89WithDefaults instantiates a new InlineObject89 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject89WithDefaults() *InlineObject89 {
	this := InlineObject89{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject89) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject89) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject89) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject89) SetName(v string) {
	o.Name = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineObject89) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject89) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineObject89) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *InlineObject89) SetScope(v string) {
	o.Scope = &v
}

func (o InlineObject89) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject89 struct {
	value *InlineObject89
	isSet bool
}

func (v NullableInlineObject89) Get() *InlineObject89 {
	return v.value
}

func (v *NullableInlineObject89) Set(val *InlineObject89) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject89) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject89) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject89(val *InlineObject89) *NullableInlineObject89 {
	return &NullableInlineObject89{value: val, isSet: true}
}

func (v NullableInlineObject89) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject89) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


