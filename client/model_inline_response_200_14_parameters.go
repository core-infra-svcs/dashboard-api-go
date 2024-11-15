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

// InlineResponse20014Parameters struct for InlineResponse20014Parameters
type InlineResponse20014Parameters struct {
	// Name of the parameter
	Name *string `json:"name,omitempty"`
	// Value of the parameter
	Value *float32 `json:"value,omitempty"`
}

// NewInlineResponse20014Parameters instantiates a new InlineResponse20014Parameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20014Parameters() *InlineResponse20014Parameters {
	this := InlineResponse20014Parameters{}
	return &this
}

// NewInlineResponse20014ParametersWithDefaults instantiates a new InlineResponse20014Parameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20014ParametersWithDefaults() *InlineResponse20014Parameters {
	this := InlineResponse20014Parameters{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20014Parameters) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014Parameters) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20014Parameters) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20014Parameters) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InlineResponse20014Parameters) GetValue() float32 {
	if o == nil || isNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014Parameters) GetValueOk() (*float32, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineResponse20014Parameters) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *InlineResponse20014Parameters) SetValue(v float32) {
	o.Value = &v
}

func (o InlineResponse20014Parameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20014Parameters struct {
	value *InlineResponse20014Parameters
	isSet bool
}

func (v NullableInlineResponse20014Parameters) Get() *InlineResponse20014Parameters {
	return v.value
}

func (v *NullableInlineResponse20014Parameters) Set(val *InlineResponse20014Parameters) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20014Parameters) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20014Parameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20014Parameters(val *InlineResponse20014Parameters) *NullableInlineResponse20014Parameters {
	return &NullableInlineResponse20014Parameters{value: val, isSet: true}
}

func (v NullableInlineResponse20014Parameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20014Parameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


