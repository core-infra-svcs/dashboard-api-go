/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20052Applications struct for InlineResponse20052Applications
type InlineResponse20052Applications struct {
	// The id of the application
	Id *string `json:"id,omitempty"`
	// The name of the application
	Name *string `json:"name,omitempty"`
}

// NewInlineResponse20052Applications instantiates a new InlineResponse20052Applications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20052Applications() *InlineResponse20052Applications {
	this := InlineResponse20052Applications{}
	return &this
}

// NewInlineResponse20052ApplicationsWithDefaults instantiates a new InlineResponse20052Applications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20052ApplicationsWithDefaults() *InlineResponse20052Applications {
	this := InlineResponse20052Applications{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20052Applications) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20052Applications) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20052Applications) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20052Applications) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20052Applications) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20052Applications) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20052Applications) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20052Applications) SetName(v string) {
	o.Name = &v
}

func (o InlineResponse20052Applications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20052Applications struct {
	value *InlineResponse20052Applications
	isSet bool
}

func (v NullableInlineResponse20052Applications) Get() *InlineResponse20052Applications {
	return v.value
}

func (v *NullableInlineResponse20052Applications) Set(val *InlineResponse20052Applications) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20052Applications) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20052Applications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20052Applications(val *InlineResponse20052Applications) *NullableInlineResponse20052Applications {
	return &NullableInlineResponse20052Applications{value: val, isSet: true}
}

func (v NullableInlineResponse20052Applications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20052Applications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

