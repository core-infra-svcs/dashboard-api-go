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

// InlineResponse200235Apns struct for InlineResponse200235Apns
type InlineResponse200235Apns struct {
	// APN name
	Name *string `json:"name,omitempty"`
}

// NewInlineResponse200235Apns instantiates a new InlineResponse200235Apns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200235Apns() *InlineResponse200235Apns {
	this := InlineResponse200235Apns{}
	return &this
}

// NewInlineResponse200235ApnsWithDefaults instantiates a new InlineResponse200235Apns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200235ApnsWithDefaults() *InlineResponse200235Apns {
	this := InlineResponse200235Apns{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200235Apns) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200235Apns) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200235Apns) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200235Apns) SetName(v string) {
	o.Name = &v
}

func (o InlineResponse200235Apns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200235Apns struct {
	value *InlineResponse200235Apns
	isSet bool
}

func (v NullableInlineResponse200235Apns) Get() *InlineResponse200235Apns {
	return v.value
}

func (v *NullableInlineResponse200235Apns) Set(val *InlineResponse200235Apns) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200235Apns) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200235Apns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200235Apns(val *InlineResponse200235Apns) *NullableInlineResponse200235Apns {
	return &NullableInlineResponse200235Apns{value: val, isSet: true}
}

func (v NullableInlineResponse200235Apns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200235Apns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


