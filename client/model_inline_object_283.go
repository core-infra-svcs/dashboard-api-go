/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject283 struct for InlineObject283
type InlineObject283 struct {
	// A name for the group of network addresses, unique within the organization (alphanumeric, space, dash, or underscore characters only)
	Name string `json:"name"`
	// Category of a policy object group (one of: NetworkObjectGroup, GeoLocationGroup, PortObjectGroup, ApplicationGroup)
	Category *string `json:"category,omitempty"`
	// A list of Policy Object ID's that this NetworkObjectGroup should be associated to (note: these ID's will replace the existing associated Policy Objects)
	ObjectIds []int32 `json:"objectIds,omitempty"`
}

// NewInlineObject283 instantiates a new InlineObject283 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject283(name string) *InlineObject283 {
	this := InlineObject283{}
	this.Name = name
	return &this
}

// NewInlineObject283WithDefaults instantiates a new InlineObject283 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject283WithDefaults() *InlineObject283 {
	this := InlineObject283{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject283) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject283) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject283) SetName(v string) {
	o.Name = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InlineObject283) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject283) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InlineObject283) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InlineObject283) SetCategory(v string) {
	o.Category = &v
}

// GetObjectIds returns the ObjectIds field value if set, zero value otherwise.
func (o *InlineObject283) GetObjectIds() []int32 {
	if o == nil || isNil(o.ObjectIds) {
		var ret []int32
		return ret
	}
	return o.ObjectIds
}

// GetObjectIdsOk returns a tuple with the ObjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject283) GetObjectIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.ObjectIds) {
    return nil, false
	}
	return o.ObjectIds, true
}

// HasObjectIds returns a boolean if a field has been set.
func (o *InlineObject283) HasObjectIds() bool {
	if o != nil && !isNil(o.ObjectIds) {
		return true
	}

	return false
}

// SetObjectIds gets a reference to the given []int32 and assigns it to the ObjectIds field.
func (o *InlineObject283) SetObjectIds(v []int32) {
	o.ObjectIds = v
}

func (o InlineObject283) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.ObjectIds) {
		toSerialize["objectIds"] = o.ObjectIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject283 struct {
	value *InlineObject283
	isSet bool
}

func (v NullableInlineObject283) Get() *InlineObject283 {
	return v.value
}

func (v *NullableInlineObject283) Set(val *InlineObject283) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject283) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject283) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject283(val *InlineObject283) *NullableInlineObject283 {
	return &NullableInlineObject283{value: val, isSet: true}
}

func (v NullableInlineObject283) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject283) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


