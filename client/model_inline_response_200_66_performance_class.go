/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20066PerformanceClass Performance class setting for uplink preference rule
type InlineResponse20066PerformanceClass struct {
	// Type of this performance class. Must be one of: 'builtin' or 'custom'
	Type string `json:"type"`
	// Name of builtin performance class. Must be present when performanceClass type is 'builtin' and value must be one of: 'VoIP'
	BuiltinPerformanceClassName *string `json:"builtinPerformanceClassName,omitempty"`
	// ID of created custom performance class, must be present when performanceClass type is \"custom\"
	CustomPerformanceClassId *string `json:"customPerformanceClassId,omitempty"`
}

// NewInlineResponse20066PerformanceClass instantiates a new InlineResponse20066PerformanceClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20066PerformanceClass(type_ string) *InlineResponse20066PerformanceClass {
	this := InlineResponse20066PerformanceClass{}
	this.Type = type_
	return &this
}

// NewInlineResponse20066PerformanceClassWithDefaults instantiates a new InlineResponse20066PerformanceClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20066PerformanceClassWithDefaults() *InlineResponse20066PerformanceClass {
	this := InlineResponse20066PerformanceClass{}
	return &this
}

// GetType returns the Type field value
func (o *InlineResponse20066PerformanceClass) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20066PerformanceClass) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse20066PerformanceClass) SetType(v string) {
	o.Type = v
}

// GetBuiltinPerformanceClassName returns the BuiltinPerformanceClassName field value if set, zero value otherwise.
func (o *InlineResponse20066PerformanceClass) GetBuiltinPerformanceClassName() string {
	if o == nil || isNil(o.BuiltinPerformanceClassName) {
		var ret string
		return ret
	}
	return *o.BuiltinPerformanceClassName
}

// GetBuiltinPerformanceClassNameOk returns a tuple with the BuiltinPerformanceClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20066PerformanceClass) GetBuiltinPerformanceClassNameOk() (*string, bool) {
	if o == nil || isNil(o.BuiltinPerformanceClassName) {
    return nil, false
	}
	return o.BuiltinPerformanceClassName, true
}

// HasBuiltinPerformanceClassName returns a boolean if a field has been set.
func (o *InlineResponse20066PerformanceClass) HasBuiltinPerformanceClassName() bool {
	if o != nil && !isNil(o.BuiltinPerformanceClassName) {
		return true
	}

	return false
}

// SetBuiltinPerformanceClassName gets a reference to the given string and assigns it to the BuiltinPerformanceClassName field.
func (o *InlineResponse20066PerformanceClass) SetBuiltinPerformanceClassName(v string) {
	o.BuiltinPerformanceClassName = &v
}

// GetCustomPerformanceClassId returns the CustomPerformanceClassId field value if set, zero value otherwise.
func (o *InlineResponse20066PerformanceClass) GetCustomPerformanceClassId() string {
	if o == nil || isNil(o.CustomPerformanceClassId) {
		var ret string
		return ret
	}
	return *o.CustomPerformanceClassId
}

// GetCustomPerformanceClassIdOk returns a tuple with the CustomPerformanceClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20066PerformanceClass) GetCustomPerformanceClassIdOk() (*string, bool) {
	if o == nil || isNil(o.CustomPerformanceClassId) {
    return nil, false
	}
	return o.CustomPerformanceClassId, true
}

// HasCustomPerformanceClassId returns a boolean if a field has been set.
func (o *InlineResponse20066PerformanceClass) HasCustomPerformanceClassId() bool {
	if o != nil && !isNil(o.CustomPerformanceClassId) {
		return true
	}

	return false
}

// SetCustomPerformanceClassId gets a reference to the given string and assigns it to the CustomPerformanceClassId field.
func (o *InlineResponse20066PerformanceClass) SetCustomPerformanceClassId(v string) {
	o.CustomPerformanceClassId = &v
}

func (o InlineResponse20066PerformanceClass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.BuiltinPerformanceClassName) {
		toSerialize["builtinPerformanceClassName"] = o.BuiltinPerformanceClassName
	}
	if !isNil(o.CustomPerformanceClassId) {
		toSerialize["customPerformanceClassId"] = o.CustomPerformanceClassId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20066PerformanceClass struct {
	value *InlineResponse20066PerformanceClass
	isSet bool
}

func (v NullableInlineResponse20066PerformanceClass) Get() *InlineResponse20066PerformanceClass {
	return v.value
}

func (v *NullableInlineResponse20066PerformanceClass) Set(val *InlineResponse20066PerformanceClass) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20066PerformanceClass) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20066PerformanceClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20066PerformanceClass(val *InlineResponse20066PerformanceClass) *NullableInlineResponse20066PerformanceClass {
	return &NullableInlineResponse20066PerformanceClass{value: val, isSet: true}
}

func (v NullableInlineResponse20066PerformanceClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20066PerformanceClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


