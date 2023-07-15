/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject72 struct for InlineObject72
type InlineObject72 struct {
	// Mask used for the subnet of all MGs in  this network.
	Mask *int32 `json:"mask,omitempty"`
	// CIDR of the pool of subnets. Each MG in this network will automatically pick a subnet from this pool.
	Cidr *string `json:"cidr,omitempty"`
}

// NewInlineObject72 instantiates a new InlineObject72 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject72() *InlineObject72 {
	this := InlineObject72{}
	return &this
}

// NewInlineObject72WithDefaults instantiates a new InlineObject72 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject72WithDefaults() *InlineObject72 {
	this := InlineObject72{}
	return &this
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *InlineObject72) GetMask() int32 {
	if o == nil || isNil(o.Mask) {
		var ret int32
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject72) GetMaskOk() (*int32, bool) {
	if o == nil || isNil(o.Mask) {
    return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *InlineObject72) HasMask() bool {
	if o != nil && !isNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given int32 and assigns it to the Mask field.
func (o *InlineObject72) SetMask(v int32) {
	o.Mask = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *InlineObject72) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject72) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *InlineObject72) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *InlineObject72) SetCidr(v string) {
	o.Cidr = &v
}

func (o InlineObject72) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mask) {
		toSerialize["mask"] = o.Mask
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject72 struct {
	value *InlineObject72
	isSet bool
}

func (v NullableInlineObject72) Get() *InlineObject72 {
	return v.value
}

func (v *NullableInlineObject72) Set(val *InlineObject72) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject72) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject72) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject72(val *InlineObject72) *NullableInlineObject72 {
	return &NullableInlineObject72{value: val, isSet: true}
}

func (v NullableInlineObject72) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject72) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


