/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200309Cdp struct for InlineResponse200309Cdp
type InlineResponse200309Cdp struct {
	// CDP RFC/official name of TLV
	Name *string `json:"name,omitempty"`
	// Value of the named TLV.
	Value *string `json:"value,omitempty"`
}

// NewInlineResponse200309Cdp instantiates a new InlineResponse200309Cdp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200309Cdp() *InlineResponse200309Cdp {
	this := InlineResponse200309Cdp{}
	return &this
}

// NewInlineResponse200309CdpWithDefaults instantiates a new InlineResponse200309Cdp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200309CdpWithDefaults() *InlineResponse200309Cdp {
	this := InlineResponse200309Cdp{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200309Cdp) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200309Cdp) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200309Cdp) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200309Cdp) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InlineResponse200309Cdp) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200309Cdp) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineResponse200309Cdp) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InlineResponse200309Cdp) SetValue(v string) {
	o.Value = &v
}

func (o InlineResponse200309Cdp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200309Cdp struct {
	value *InlineResponse200309Cdp
	isSet bool
}

func (v NullableInlineResponse200309Cdp) Get() *InlineResponse200309Cdp {
	return v.value
}

func (v *NullableInlineResponse200309Cdp) Set(val *InlineResponse200309Cdp) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200309Cdp) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200309Cdp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200309Cdp(val *InlineResponse200309Cdp) *NullableInlineResponse200309Cdp {
	return &NullableInlineResponse200309Cdp{value: val, isSet: true}
}

func (v NullableInlineResponse200309Cdp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200309Cdp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


