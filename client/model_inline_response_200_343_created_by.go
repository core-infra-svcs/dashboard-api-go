/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200343CreatedBy Information around who initiated the callback
type InlineResponse200343CreatedBy struct {
	// The ID of the user who initiated the callback
	AdminId *string `json:"adminId,omitempty"`
}

// NewInlineResponse200343CreatedBy instantiates a new InlineResponse200343CreatedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200343CreatedBy() *InlineResponse200343CreatedBy {
	this := InlineResponse200343CreatedBy{}
	return &this
}

// NewInlineResponse200343CreatedByWithDefaults instantiates a new InlineResponse200343CreatedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200343CreatedByWithDefaults() *InlineResponse200343CreatedBy {
	this := InlineResponse200343CreatedBy{}
	return &this
}

// GetAdminId returns the AdminId field value if set, zero value otherwise.
func (o *InlineResponse200343CreatedBy) GetAdminId() string {
	if o == nil || isNil(o.AdminId) {
		var ret string
		return ret
	}
	return *o.AdminId
}

// GetAdminIdOk returns a tuple with the AdminId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200343CreatedBy) GetAdminIdOk() (*string, bool) {
	if o == nil || isNil(o.AdminId) {
    return nil, false
	}
	return o.AdminId, true
}

// HasAdminId returns a boolean if a field has been set.
func (o *InlineResponse200343CreatedBy) HasAdminId() bool {
	if o != nil && !isNil(o.AdminId) {
		return true
	}

	return false
}

// SetAdminId gets a reference to the given string and assigns it to the AdminId field.
func (o *InlineResponse200343CreatedBy) SetAdminId(v string) {
	o.AdminId = &v
}

func (o InlineResponse200343CreatedBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AdminId) {
		toSerialize["adminId"] = o.AdminId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200343CreatedBy struct {
	value *InlineResponse200343CreatedBy
	isSet bool
}

func (v NullableInlineResponse200343CreatedBy) Get() *InlineResponse200343CreatedBy {
	return v.value
}

func (v *NullableInlineResponse200343CreatedBy) Set(val *InlineResponse200343CreatedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200343CreatedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200343CreatedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200343CreatedBy(val *InlineResponse200343CreatedBy) *NullableInlineResponse200343CreatedBy {
	return &NullableInlineResponse200343CreatedBy{value: val, isSet: true}
}

func (v NullableInlineResponse200343CreatedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200343CreatedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


