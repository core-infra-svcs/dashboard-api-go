/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20016 struct for InlineResponse20016
type InlineResponse20016 struct {
	// Event category
	Category *string `json:"category,omitempty"`
	// Event type
	Type *string `json:"type,omitempty"`
	// Description of the event
	Description *string `json:"description,omitempty"`
}

// NewInlineResponse20016 instantiates a new InlineResponse20016 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20016() *InlineResponse20016 {
	this := InlineResponse20016{}
	return &this
}

// NewInlineResponse20016WithDefaults instantiates a new InlineResponse20016 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20016WithDefaults() *InlineResponse20016 {
	this := InlineResponse20016{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InlineResponse20016) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InlineResponse20016) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InlineResponse20016) SetCategory(v string) {
	o.Category = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20016) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20016) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20016) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse20016) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20016) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse20016) SetDescription(v string) {
	o.Description = &v
}

func (o InlineResponse20016) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20016 struct {
	value *InlineResponse20016
	isSet bool
}

func (v NullableInlineResponse20016) Get() *InlineResponse20016 {
	return v.value
}

func (v *NullableInlineResponse20016) Set(val *InlineResponse20016) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20016) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20016) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20016(val *InlineResponse20016) *NullableInlineResponse20016 {
	return &NullableInlineResponse20016{value: val, isSet: true}
}

func (v NullableInlineResponse20016) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20016) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


