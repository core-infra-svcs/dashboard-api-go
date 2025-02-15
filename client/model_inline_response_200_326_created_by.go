/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200326CreatedBy Information around who initiated the callback
type InlineResponse200326CreatedBy struct {
	// The ID of the user who initiated the callback
	AdminId *string `json:"adminId,omitempty"`
}

// NewInlineResponse200326CreatedBy instantiates a new InlineResponse200326CreatedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200326CreatedBy() *InlineResponse200326CreatedBy {
	this := InlineResponse200326CreatedBy{}
	return &this
}

// NewInlineResponse200326CreatedByWithDefaults instantiates a new InlineResponse200326CreatedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200326CreatedByWithDefaults() *InlineResponse200326CreatedBy {
	this := InlineResponse200326CreatedBy{}
	return &this
}

// GetAdminId returns the AdminId field value if set, zero value otherwise.
func (o *InlineResponse200326CreatedBy) GetAdminId() string {
	if o == nil || isNil(o.AdminId) {
		var ret string
		return ret
	}
	return *o.AdminId
}

// GetAdminIdOk returns a tuple with the AdminId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200326CreatedBy) GetAdminIdOk() (*string, bool) {
	if o == nil || isNil(o.AdminId) {
    return nil, false
	}
	return o.AdminId, true
}

// HasAdminId returns a boolean if a field has been set.
func (o *InlineResponse200326CreatedBy) HasAdminId() bool {
	if o != nil && !isNil(o.AdminId) {
		return true
	}

	return false
}

// SetAdminId gets a reference to the given string and assigns it to the AdminId field.
func (o *InlineResponse200326CreatedBy) SetAdminId(v string) {
	o.AdminId = &v
}

func (o InlineResponse200326CreatedBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AdminId) {
		toSerialize["adminId"] = o.AdminId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200326CreatedBy struct {
	value *InlineResponse200326CreatedBy
	isSet bool
}

func (v NullableInlineResponse200326CreatedBy) Get() *InlineResponse200326CreatedBy {
	return v.value
}

func (v *NullableInlineResponse200326CreatedBy) Set(val *InlineResponse200326CreatedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200326CreatedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200326CreatedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200326CreatedBy(val *InlineResponse200326CreatedBy) *NullableInlineResponse200326CreatedBy {
	return &NullableInlineResponse200326CreatedBy{value: val, isSet: true}
}

func (v NullableInlineResponse200326CreatedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200326CreatedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


