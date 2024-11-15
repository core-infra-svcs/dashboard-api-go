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

// InlineResponse20074 struct for InlineResponse20074
type InlineResponse20074 struct {
	// Schedule id
	Id *string `json:"id,omitempty"`
	// Schedule name
	Name *string `json:"name,omitempty"`
}

// NewInlineResponse20074 instantiates a new InlineResponse20074 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20074() *InlineResponse20074 {
	this := InlineResponse20074{}
	return &this
}

// NewInlineResponse20074WithDefaults instantiates a new InlineResponse20074 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20074WithDefaults() *InlineResponse20074 {
	this := InlineResponse20074{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20074) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20074) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20074) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20074) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20074) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20074) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20074) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20074) SetName(v string) {
	o.Name = &v
}

func (o InlineResponse20074) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20074 struct {
	value *InlineResponse20074
	isSet bool
}

func (v NullableInlineResponse20074) Get() *InlineResponse20074 {
	return v.value
}

func (v *NullableInlineResponse20074) Set(val *InlineResponse20074) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20074) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20074) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20074(val *InlineResponse20074) *NullableInlineResponse20074 {
	return &NullableInlineResponse20074{value: val, isSet: true}
}

func (v NullableInlineResponse20074) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20074) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


