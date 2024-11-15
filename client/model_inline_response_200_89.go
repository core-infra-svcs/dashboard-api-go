/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20089 struct for InlineResponse20089
type InlineResponse20089 struct {
	// Event category
	Category *string `json:"category,omitempty"`
	// Event type
	Type *string `json:"type,omitempty"`
	// Description of the event
	Description *string `json:"description,omitempty"`
}

// NewInlineResponse20089 instantiates a new InlineResponse20089 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20089() *InlineResponse20089 {
	this := InlineResponse20089{}
	return &this
}

// NewInlineResponse20089WithDefaults instantiates a new InlineResponse20089 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20089WithDefaults() *InlineResponse20089 {
	this := InlineResponse20089{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InlineResponse20089) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InlineResponse20089) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InlineResponse20089) SetCategory(v string) {
	o.Category = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20089) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20089) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20089) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse20089) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20089) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse20089) SetDescription(v string) {
	o.Description = &v
}

func (o InlineResponse20089) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20089 struct {
	value *InlineResponse20089
	isSet bool
}

func (v NullableInlineResponse20089) Get() *InlineResponse20089 {
	return v.value
}

func (v *NullableInlineResponse20089) Set(val *InlineResponse20089) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20089) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20089) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20089(val *InlineResponse20089) *NullableInlineResponse20089 {
	return &NullableInlineResponse20089{value: val, isSet: true}
}

func (v NullableInlineResponse20089) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20089) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


