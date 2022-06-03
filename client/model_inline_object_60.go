/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject60 struct for InlineObject60
type InlineObject60 struct {
	// Mask used for the subnet of all MGs in  this network.
	Mask *int32 `json:"mask,omitempty"`
	// CIDR of the pool of subnets. Each MG in this network will automatically pick a subnet from this pool.
	Cidr *string `json:"cidr,omitempty"`
}

// NewInlineObject60 instantiates a new InlineObject60 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject60() *InlineObject60 {
	this := InlineObject60{}
	return &this
}

// NewInlineObject60WithDefaults instantiates a new InlineObject60 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject60WithDefaults() *InlineObject60 {
	this := InlineObject60{}
	return &this
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *InlineObject60) GetMask() int32 {
	if o == nil || o.Mask == nil {
		var ret int32
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject60) GetMaskOk() (*int32, bool) {
	if o == nil || o.Mask == nil {
		return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *InlineObject60) HasMask() bool {
	if o != nil && o.Mask != nil {
		return true
	}

	return false
}

// SetMask gets a reference to the given int32 and assigns it to the Mask field.
func (o *InlineObject60) SetMask(v int32) {
	o.Mask = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *InlineObject60) GetCidr() string {
	if o == nil || o.Cidr == nil {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject60) GetCidrOk() (*string, bool) {
	if o == nil || o.Cidr == nil {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *InlineObject60) HasCidr() bool {
	if o != nil && o.Cidr != nil {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *InlineObject60) SetCidr(v string) {
	o.Cidr = &v
}

func (o InlineObject60) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mask != nil {
		toSerialize["mask"] = o.Mask
	}
	if o.Cidr != nil {
		toSerialize["cidr"] = o.Cidr
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject60 struct {
	value *InlineObject60
	isSet bool
}

func (v NullableInlineObject60) Get() *InlineObject60 {
	return v.value
}

func (v *NullableInlineObject60) Set(val *InlineObject60) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject60) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject60) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject60(val *InlineObject60) *NullableInlineObject60 {
	return &NullableInlineObject60{value: val, isSet: true}
}

func (v NullableInlineObject60) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject60) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


