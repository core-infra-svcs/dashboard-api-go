/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20035DhcpOptions struct for InlineResponse20035DhcpOptions
type InlineResponse20035DhcpOptions struct {
	// The code for DHCP option which should be from 2 to 254
	Code *string `json:"code,omitempty"`
	// The type of the DHCP option which should be one of ('text', 'ip', 'integer' or 'hex')
	Type *string `json:"type,omitempty"`
	// The value of the DHCP option
	Value *string `json:"value,omitempty"`
}

// NewInlineResponse20035DhcpOptions instantiates a new InlineResponse20035DhcpOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20035DhcpOptions() *InlineResponse20035DhcpOptions {
	this := InlineResponse20035DhcpOptions{}
	return &this
}

// NewInlineResponse20035DhcpOptionsWithDefaults instantiates a new InlineResponse20035DhcpOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20035DhcpOptionsWithDefaults() *InlineResponse20035DhcpOptions {
	this := InlineResponse20035DhcpOptions{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InlineResponse20035DhcpOptions) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035DhcpOptions) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InlineResponse20035DhcpOptions) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InlineResponse20035DhcpOptions) SetCode(v string) {
	o.Code = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20035DhcpOptions) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035DhcpOptions) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20035DhcpOptions) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20035DhcpOptions) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InlineResponse20035DhcpOptions) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035DhcpOptions) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineResponse20035DhcpOptions) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InlineResponse20035DhcpOptions) SetValue(v string) {
	o.Value = &v
}

func (o InlineResponse20035DhcpOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20035DhcpOptions struct {
	value *InlineResponse20035DhcpOptions
	isSet bool
}

func (v NullableInlineResponse20035DhcpOptions) Get() *InlineResponse20035DhcpOptions {
	return v.value
}

func (v *NullableInlineResponse20035DhcpOptions) Set(val *InlineResponse20035DhcpOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20035DhcpOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20035DhcpOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20035DhcpOptions(val *InlineResponse20035DhcpOptions) *NullableInlineResponse20035DhcpOptions {
	return &NullableInlineResponse20035DhcpOptions{value: val, isSet: true}
}

func (v NullableInlineResponse20035DhcpOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20035DhcpOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


