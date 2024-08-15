/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200292CountsByStatus The count data, indexed by active or inactive status
type InlineResponse200292CountsByStatus struct {
	Active *InlineResponse200292CountsByStatusActive `json:"active,omitempty"`
	Inactive *InlineResponse200292CountsByStatusInactive `json:"inactive,omitempty"`
}

// NewInlineResponse200292CountsByStatus instantiates a new InlineResponse200292CountsByStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200292CountsByStatus() *InlineResponse200292CountsByStatus {
	this := InlineResponse200292CountsByStatus{}
	return &this
}

// NewInlineResponse200292CountsByStatusWithDefaults instantiates a new InlineResponse200292CountsByStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200292CountsByStatusWithDefaults() *InlineResponse200292CountsByStatus {
	this := InlineResponse200292CountsByStatus{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *InlineResponse200292CountsByStatus) GetActive() InlineResponse200292CountsByStatusActive {
	if o == nil || isNil(o.Active) {
		var ret InlineResponse200292CountsByStatusActive
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200292CountsByStatus) GetActiveOk() (*InlineResponse200292CountsByStatusActive, bool) {
	if o == nil || isNil(o.Active) {
    return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *InlineResponse200292CountsByStatus) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given InlineResponse200292CountsByStatusActive and assigns it to the Active field.
func (o *InlineResponse200292CountsByStatus) SetActive(v InlineResponse200292CountsByStatusActive) {
	o.Active = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *InlineResponse200292CountsByStatus) GetInactive() InlineResponse200292CountsByStatusInactive {
	if o == nil || isNil(o.Inactive) {
		var ret InlineResponse200292CountsByStatusInactive
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200292CountsByStatus) GetInactiveOk() (*InlineResponse200292CountsByStatusInactive, bool) {
	if o == nil || isNil(o.Inactive) {
    return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *InlineResponse200292CountsByStatus) HasInactive() bool {
	if o != nil && !isNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given InlineResponse200292CountsByStatusInactive and assigns it to the Inactive field.
func (o *InlineResponse200292CountsByStatus) SetInactive(v InlineResponse200292CountsByStatusInactive) {
	o.Inactive = &v
}

func (o InlineResponse200292CountsByStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200292CountsByStatus struct {
	value *InlineResponse200292CountsByStatus
	isSet bool
}

func (v NullableInlineResponse200292CountsByStatus) Get() *InlineResponse200292CountsByStatus {
	return v.value
}

func (v *NullableInlineResponse200292CountsByStatus) Set(val *InlineResponse200292CountsByStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200292CountsByStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200292CountsByStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200292CountsByStatus(val *InlineResponse200292CountsByStatus) *NullableInlineResponse200292CountsByStatus {
	return &NullableInlineResponse200292CountsByStatus{value: val, isSet: true}
}

func (v NullableInlineResponse200292CountsByStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200292CountsByStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


