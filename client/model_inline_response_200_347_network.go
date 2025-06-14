/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200347Network Access point network
type InlineResponse200347Network struct {
	// Access point network ID
	Id *string `json:"id,omitempty"`
}

// NewInlineResponse200347Network instantiates a new InlineResponse200347Network object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200347Network() *InlineResponse200347Network {
	this := InlineResponse200347Network{}
	return &this
}

// NewInlineResponse200347NetworkWithDefaults instantiates a new InlineResponse200347Network object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200347NetworkWithDefaults() *InlineResponse200347Network {
	this := InlineResponse200347Network{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200347Network) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200347Network) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200347Network) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200347Network) SetId(v string) {
	o.Id = &v
}

func (o InlineResponse200347Network) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200347Network struct {
	value *InlineResponse200347Network
	isSet bool
}

func (v NullableInlineResponse200347Network) Get() *InlineResponse200347Network {
	return v.value
}

func (v *NullableInlineResponse200347Network) Set(val *InlineResponse200347Network) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200347Network) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200347Network) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200347Network(val *InlineResponse200347Network) *NullableInlineResponse200347Network {
	return &NullableInlineResponse200347Network{value: val, isSet: true}
}

func (v NullableInlineResponse200347Network) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200347Network) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


