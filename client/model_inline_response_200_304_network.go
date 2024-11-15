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

// InlineResponse200304Network Identifying information of the switch's network.
type InlineResponse200304Network struct {
	// The name of the network.
	Name *string `json:"name,omitempty"`
	// The ID of the network.
	Id *string `json:"id,omitempty"`
}

// NewInlineResponse200304Network instantiates a new InlineResponse200304Network object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200304Network() *InlineResponse200304Network {
	this := InlineResponse200304Network{}
	return &this
}

// NewInlineResponse200304NetworkWithDefaults instantiates a new InlineResponse200304Network object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200304NetworkWithDefaults() *InlineResponse200304Network {
	this := InlineResponse200304Network{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200304Network) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200304Network) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200304Network) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200304Network) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200304Network) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200304Network) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200304Network) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200304Network) SetId(v string) {
	o.Id = &v
}

func (o InlineResponse200304Network) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200304Network struct {
	value *InlineResponse200304Network
	isSet bool
}

func (v NullableInlineResponse200304Network) Get() *InlineResponse200304Network {
	return v.value
}

func (v *NullableInlineResponse200304Network) Set(val *InlineResponse200304Network) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200304Network) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200304Network) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200304Network(val *InlineResponse200304Network) *NullableInlineResponse200304Network {
	return &NullableInlineResponse200304Network{value: val, isSet: true}
}

func (v NullableInlineResponse200304Network) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200304Network) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


