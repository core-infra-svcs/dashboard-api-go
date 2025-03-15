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

// InlineObject271 struct for InlineObject271
type InlineObject271 struct {
	// A name for the group of network addresses, unique within the organization (alphanumeric, space, dash, or underscore characters only)
	Name string `json:"name"`
	// Category of a policy object group (one of: NetworkObjectGroup, GeoLocationGroup, PortObjectGroup, ApplicationGroup)
	Category *string `json:"category,omitempty"`
	// A list of Policy Object ID's that this NetworkObjectGroup should be associated to (note: these ID's will replace the existing associated Policy Objects)
	ObjectIds []int32 `json:"objectIds,omitempty"`
}

// NewInlineObject271 instantiates a new InlineObject271 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject271(name string) *InlineObject271 {
	this := InlineObject271{}
	this.Name = name
	return &this
}

// NewInlineObject271WithDefaults instantiates a new InlineObject271 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject271WithDefaults() *InlineObject271 {
	this := InlineObject271{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject271) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject271) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject271) SetName(v string) {
	o.Name = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InlineObject271) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject271) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InlineObject271) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InlineObject271) SetCategory(v string) {
	o.Category = &v
}

// GetObjectIds returns the ObjectIds field value if set, zero value otherwise.
func (o *InlineObject271) GetObjectIds() []int32 {
	if o == nil || isNil(o.ObjectIds) {
		var ret []int32
		return ret
	}
	return o.ObjectIds
}

// GetObjectIdsOk returns a tuple with the ObjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject271) GetObjectIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.ObjectIds) {
    return nil, false
	}
	return o.ObjectIds, true
}

// HasObjectIds returns a boolean if a field has been set.
func (o *InlineObject271) HasObjectIds() bool {
	if o != nil && !isNil(o.ObjectIds) {
		return true
	}

	return false
}

// SetObjectIds gets a reference to the given []int32 and assigns it to the ObjectIds field.
func (o *InlineObject271) SetObjectIds(v []int32) {
	o.ObjectIds = v
}

func (o InlineObject271) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject271 struct {
	value *InlineObject271
	isSet bool
}

func (v NullableInlineObject271) Get() *InlineObject271 {
	return v.value
}

func (v *NullableInlineObject271) Set(val *InlineObject271) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject271) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject271) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject271(val *InlineObject271) *NullableInlineObject271 {
	return &NullableInlineObject271{value: val, isSet: true}
}

func (v NullableInlineObject271) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject271) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


