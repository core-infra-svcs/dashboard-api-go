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

// InlineResponse200335ByInterface struct for InlineResponse200335ByInterface
type InlineResponse200335ByInterface struct {
	// The name of the wireless LAN controller interface
	Name *string `json:"name,omitempty"`
	Usage *InlineResponse200335Usage `json:"usage,omitempty"`
}

// NewInlineResponse200335ByInterface instantiates a new InlineResponse200335ByInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200335ByInterface() *InlineResponse200335ByInterface {
	this := InlineResponse200335ByInterface{}
	return &this
}

// NewInlineResponse200335ByInterfaceWithDefaults instantiates a new InlineResponse200335ByInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200335ByInterfaceWithDefaults() *InlineResponse200335ByInterface {
	this := InlineResponse200335ByInterface{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200335ByInterface) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200335ByInterface) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200335ByInterface) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200335ByInterface) SetName(v string) {
	o.Name = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200335ByInterface) GetUsage() InlineResponse200335Usage {
	if o == nil || isNil(o.Usage) {
		var ret InlineResponse200335Usage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200335ByInterface) GetUsageOk() (*InlineResponse200335Usage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200335ByInterface) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given InlineResponse200335Usage and assigns it to the Usage field.
func (o *InlineResponse200335ByInterface) SetUsage(v InlineResponse200335Usage) {
	o.Usage = &v
}

func (o InlineResponse200335ByInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200335ByInterface struct {
	value *InlineResponse200335ByInterface
	isSet bool
}

func (v NullableInlineResponse200335ByInterface) Get() *InlineResponse200335ByInterface {
	return v.value
}

func (v *NullableInlineResponse200335ByInterface) Set(val *InlineResponse200335ByInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200335ByInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200335ByInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200335ByInterface(val *InlineResponse200335ByInterface) *NullableInlineResponse200335ByInterface {
	return &NullableInlineResponse200335ByInterface{value: val, isSet: true}
}

func (v NullableInlineResponse200335ByInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200335ByInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

