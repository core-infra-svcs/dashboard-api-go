/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject214 struct for InlineObject214
type InlineObject214 struct {
	// A name for the group of network addresses, unique within the organization (alphanumeric, space, dash, or underscore characters only)
	Name *string `json:"name,omitempty"`
	// A list of Policy Object ID's that this NetworkObjectGroup should be associated to (note: these ID's will replace the existing associated Policy Objects)
	ObjectIds []int32 `json:"objectIds,omitempty"`
}

// NewInlineObject214 instantiates a new InlineObject214 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject214() *InlineObject214 {
	this := InlineObject214{}
	return &this
}

// NewInlineObject214WithDefaults instantiates a new InlineObject214 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject214WithDefaults() *InlineObject214 {
	this := InlineObject214{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject214) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject214) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject214) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject214) SetName(v string) {
	o.Name = &v
}

// GetObjectIds returns the ObjectIds field value if set, zero value otherwise.
func (o *InlineObject214) GetObjectIds() []int32 {
	if o == nil || isNil(o.ObjectIds) {
		var ret []int32
		return ret
	}
	return o.ObjectIds
}

// GetObjectIdsOk returns a tuple with the ObjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject214) GetObjectIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.ObjectIds) {
    return nil, false
	}
	return o.ObjectIds, true
}

// HasObjectIds returns a boolean if a field has been set.
func (o *InlineObject214) HasObjectIds() bool {
	if o != nil && !isNil(o.ObjectIds) {
		return true
	}

	return false
}

// SetObjectIds gets a reference to the given []int32 and assigns it to the ObjectIds field.
func (o *InlineObject214) SetObjectIds(v []int32) {
	o.ObjectIds = v
}

func (o InlineObject214) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ObjectIds) {
		toSerialize["objectIds"] = o.ObjectIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject214 struct {
	value *InlineObject214
	isSet bool
}

func (v NullableInlineObject214) Get() *InlineObject214 {
	return v.value
}

func (v *NullableInlineObject214) Set(val *InlineObject214) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject214) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject214) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject214(val *InlineObject214) *NullableInlineObject214 {
	return &NullableInlineObject214{value: val, isSet: true}
}

func (v NullableInlineObject214) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject214) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


