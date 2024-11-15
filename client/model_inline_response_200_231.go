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

// InlineResponse200231 struct for InlineResponse200231
type InlineResponse200231 struct {
	// Permission scope id
	Id *string `json:"id,omitempty"`
	// Name of permission scope
	Name *string `json:"name,omitempty"`
	// Permission scope level
	Level *string `json:"level,omitempty"`
}

// NewInlineResponse200231 instantiates a new InlineResponse200231 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200231() *InlineResponse200231 {
	this := InlineResponse200231{}
	return &this
}

// NewInlineResponse200231WithDefaults instantiates a new InlineResponse200231 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200231WithDefaults() *InlineResponse200231 {
	this := InlineResponse200231{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200231) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200231) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200231) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200231) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200231) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200231) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200231) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200231) SetName(v string) {
	o.Name = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *InlineResponse200231) GetLevel() string {
	if o == nil || isNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200231) GetLevelOk() (*string, bool) {
	if o == nil || isNil(o.Level) {
    return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *InlineResponse200231) HasLevel() bool {
	if o != nil && !isNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *InlineResponse200231) SetLevel(v string) {
	o.Level = &v
}

func (o InlineResponse200231) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200231 struct {
	value *InlineResponse200231
	isSet bool
}

func (v NullableInlineResponse200231) Get() *InlineResponse200231 {
	return v.value
}

func (v *NullableInlineResponse200231) Set(val *InlineResponse200231) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200231) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200231) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200231(val *InlineResponse200231) *NullableInlineResponse200231 {
	return &NullableInlineResponse200231{value: val, isSet: true}
}

func (v NullableInlineResponse200231) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200231) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


