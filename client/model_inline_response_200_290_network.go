/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200290Network Group of devices and settings.
type InlineResponse200290Network struct {
	// Unique identifier for network.
	Id *string `json:"id,omitempty"`
	// Name of network.
	Name *string `json:"name,omitempty"`
}

// NewInlineResponse200290Network instantiates a new InlineResponse200290Network object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200290Network() *InlineResponse200290Network {
	this := InlineResponse200290Network{}
	return &this
}

// NewInlineResponse200290NetworkWithDefaults instantiates a new InlineResponse200290Network object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200290NetworkWithDefaults() *InlineResponse200290Network {
	this := InlineResponse200290Network{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200290Network) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200290Network) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200290Network) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200290Network) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200290Network) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200290Network) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200290Network) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200290Network) SetName(v string) {
	o.Name = &v
}

func (o InlineResponse200290Network) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200290Network struct {
	value *InlineResponse200290Network
	isSet bool
}

func (v NullableInlineResponse200290Network) Get() *InlineResponse200290Network {
	return v.value
}

func (v *NullableInlineResponse200290Network) Set(val *InlineResponse200290Network) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200290Network) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200290Network) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200290Network(val *InlineResponse200290Network) *NullableInlineResponse200290Network {
	return &NullableInlineResponse200290Network{value: val, isSet: true}
}

func (v NullableInlineResponse200290Network) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200290Network) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


