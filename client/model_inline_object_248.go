/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject248 struct for InlineObject248
type InlineObject248 struct {
	// The name of the Limited Access Role
	Name string `json:"name"`
	// The scope of the Limited Access Role
	Scope *string `json:"scope,omitempty"`
	// The tags of the Limited Access Role
	Tags []string `json:"tags,omitempty"`
}

// NewInlineObject248 instantiates a new InlineObject248 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject248(name string) *InlineObject248 {
	this := InlineObject248{}
	this.Name = name
	return &this
}

// NewInlineObject248WithDefaults instantiates a new InlineObject248 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject248WithDefaults() *InlineObject248 {
	this := InlineObject248{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject248) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject248) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject248) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineObject248) GetScope() string {
	if o == nil || isNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject248) GetScopeOk() (*string, bool) {
	if o == nil || isNil(o.Scope) {
    return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineObject248) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *InlineObject248) SetScope(v string) {
	o.Scope = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject248) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject248) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject248) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineObject248) SetTags(v []string) {
	o.Tags = v
}

func (o InlineObject248) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject248 struct {
	value *InlineObject248
	isSet bool
}

func (v NullableInlineObject248) Get() *InlineObject248 {
	return v.value
}

func (v *NullableInlineObject248) Set(val *InlineObject248) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject248) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject248) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject248(val *InlineObject248) *NullableInlineObject248 {
	return &NullableInlineObject248{value: val, isSet: true}
}

func (v NullableInlineObject248) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject248) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


