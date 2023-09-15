/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20086 struct for InlineResponse20086
type InlineResponse20086 struct {
	// Switch stacks id
	Id *string `json:"id,omitempty"`
	// Switch stacks name
	Name *string `json:"name,omitempty"`
	// Serials of the switches in the switch stack
	Serials []string `json:"serials,omitempty"`
}

// NewInlineResponse20086 instantiates a new InlineResponse20086 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20086() *InlineResponse20086 {
	this := InlineResponse20086{}
	return &this
}

// NewInlineResponse20086WithDefaults instantiates a new InlineResponse20086 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20086WithDefaults() *InlineResponse20086 {
	this := InlineResponse20086{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20086) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20086) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20086) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20086) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20086) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20086) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20086) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20086) SetName(v string) {
	o.Name = &v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineResponse20086) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20086) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineResponse20086) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineResponse20086) SetSerials(v []string) {
	o.Serials = v
}

func (o InlineResponse20086) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20086 struct {
	value *InlineResponse20086
	isSet bool
}

func (v NullableInlineResponse20086) Get() *InlineResponse20086 {
	return v.value
}

func (v *NullableInlineResponse20086) Set(val *InlineResponse20086) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20086) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20086) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20086(val *InlineResponse20086) *NullableInlineResponse20086 {
	return &NullableInlineResponse20086{value: val, isSet: true}
}

func (v NullableInlineResponse20086) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20086) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


