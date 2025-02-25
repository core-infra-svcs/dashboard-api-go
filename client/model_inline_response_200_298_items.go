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

// InlineResponse200298Items struct for InlineResponse200298Items
type InlineResponse200298Items struct {
	// The Id of the limited access role
	RoleId *string `json:"roleId,omitempty"`
	// The name of the limited access role
	Name *string `json:"name,omitempty"`
	// The scope of the limited access role
	Scope *string `json:"scope,omitempty"`
	// The tags of the limited access role
	Tags []string `json:"tags,omitempty"`
}

// NewInlineResponse200298Items instantiates a new InlineResponse200298Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200298Items() *InlineResponse200298Items {
	this := InlineResponse200298Items{}
	return &this
}

// NewInlineResponse200298ItemsWithDefaults instantiates a new InlineResponse200298Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200298ItemsWithDefaults() *InlineResponse200298Items {
	this := InlineResponse200298Items{}
	return &this
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *InlineResponse200298Items) GetRoleId() string {
	if o == nil || isNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298Items) GetRoleIdOk() (*string, bool) {
	if o == nil || isNil(o.RoleId) {
    return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *InlineResponse200298Items) HasRoleId() bool {
	if o != nil && !isNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *InlineResponse200298Items) SetRoleId(v string) {
	o.RoleId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200298Items) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298Items) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200298Items) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200298Items) SetName(v string) {
	o.Name = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineResponse200298Items) GetScope() string {
	if o == nil || isNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298Items) GetScopeOk() (*string, bool) {
	if o == nil || isNil(o.Scope) {
    return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineResponse200298Items) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *InlineResponse200298Items) SetScope(v string) {
	o.Scope = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200298Items) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200298Items) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200298Items) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse200298Items) SetTags(v []string) {
	o.Tags = v
}

func (o InlineResponse200298Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RoleId) {
		toSerialize["roleId"] = o.RoleId
	}
	if !isNil(o.Name) {
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

type NullableInlineResponse200298Items struct {
	value *InlineResponse200298Items
	isSet bool
}

func (v NullableInlineResponse200298Items) Get() *InlineResponse200298Items {
	return v.value
}

func (v *NullableInlineResponse200298Items) Set(val *InlineResponse200298Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200298Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200298Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200298Items(val *InlineResponse200298Items) *NullableInlineResponse200298Items {
	return &NullableInlineResponse200298Items{value: val, isSet: true}
}

func (v NullableInlineResponse200298Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200298Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


