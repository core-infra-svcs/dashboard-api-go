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

// InlineResponse200218ExportedSubnets struct for InlineResponse200218ExportedSubnets
type InlineResponse200218ExportedSubnets struct {
	// Subnet
	Subnet *string `json:"subnet,omitempty"`
	// Name of the subnet
	Name *string `json:"name,omitempty"`
}

// NewInlineResponse200218ExportedSubnets instantiates a new InlineResponse200218ExportedSubnets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200218ExportedSubnets() *InlineResponse200218ExportedSubnets {
	this := InlineResponse200218ExportedSubnets{}
	return &this
}

// NewInlineResponse200218ExportedSubnetsWithDefaults instantiates a new InlineResponse200218ExportedSubnets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200218ExportedSubnetsWithDefaults() *InlineResponse200218ExportedSubnets {
	this := InlineResponse200218ExportedSubnets{}
	return &this
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineResponse200218ExportedSubnets) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200218ExportedSubnets) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineResponse200218ExportedSubnets) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineResponse200218ExportedSubnets) SetSubnet(v string) {
	o.Subnet = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200218ExportedSubnets) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200218ExportedSubnets) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200218ExportedSubnets) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200218ExportedSubnets) SetName(v string) {
	o.Name = &v
}

func (o InlineResponse200218ExportedSubnets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200218ExportedSubnets struct {
	value *InlineResponse200218ExportedSubnets
	isSet bool
}

func (v NullableInlineResponse200218ExportedSubnets) Get() *InlineResponse200218ExportedSubnets {
	return v.value
}

func (v *NullableInlineResponse200218ExportedSubnets) Set(val *InlineResponse200218ExportedSubnets) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200218ExportedSubnets) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200218ExportedSubnets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200218ExportedSubnets(val *InlineResponse200218ExportedSubnets) *NullableInlineResponse200218ExportedSubnets {
	return &NullableInlineResponse200218ExportedSubnets{value: val, isSet: true}
}

func (v NullableInlineResponse200218ExportedSubnets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200218ExportedSubnets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

