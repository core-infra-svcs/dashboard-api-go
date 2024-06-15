/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200290CountsByStatus The count data, indexed by active or inactive status
type InlineResponse200290CountsByStatus struct {
	Active *InlineResponse200290CountsByStatusActive `json:"active,omitempty"`
	Inactive *InlineResponse200290CountsByStatusInactive `json:"inactive,omitempty"`
}

// NewInlineResponse200290CountsByStatus instantiates a new InlineResponse200290CountsByStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200290CountsByStatus() *InlineResponse200290CountsByStatus {
	this := InlineResponse200290CountsByStatus{}
	return &this
}

// NewInlineResponse200290CountsByStatusWithDefaults instantiates a new InlineResponse200290CountsByStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200290CountsByStatusWithDefaults() *InlineResponse200290CountsByStatus {
	this := InlineResponse200290CountsByStatus{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *InlineResponse200290CountsByStatus) GetActive() InlineResponse200290CountsByStatusActive {
	if o == nil || isNil(o.Active) {
		var ret InlineResponse200290CountsByStatusActive
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200290CountsByStatus) GetActiveOk() (*InlineResponse200290CountsByStatusActive, bool) {
	if o == nil || isNil(o.Active) {
    return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *InlineResponse200290CountsByStatus) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given InlineResponse200290CountsByStatusActive and assigns it to the Active field.
func (o *InlineResponse200290CountsByStatus) SetActive(v InlineResponse200290CountsByStatusActive) {
	o.Active = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *InlineResponse200290CountsByStatus) GetInactive() InlineResponse200290CountsByStatusInactive {
	if o == nil || isNil(o.Inactive) {
		var ret InlineResponse200290CountsByStatusInactive
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200290CountsByStatus) GetInactiveOk() (*InlineResponse200290CountsByStatusInactive, bool) {
	if o == nil || isNil(o.Inactive) {
    return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *InlineResponse200290CountsByStatus) HasInactive() bool {
	if o != nil && !isNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given InlineResponse200290CountsByStatusInactive and assigns it to the Inactive field.
func (o *InlineResponse200290CountsByStatus) SetInactive(v InlineResponse200290CountsByStatusInactive) {
	o.Inactive = &v
}

func (o InlineResponse200290CountsByStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200290CountsByStatus struct {
	value *InlineResponse200290CountsByStatus
	isSet bool
}

func (v NullableInlineResponse200290CountsByStatus) Get() *InlineResponse200290CountsByStatus {
	return v.value
}

func (v *NullableInlineResponse200290CountsByStatus) Set(val *InlineResponse200290CountsByStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200290CountsByStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200290CountsByStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200290CountsByStatus(val *InlineResponse200290CountsByStatus) *NullableInlineResponse200290CountsByStatus {
	return &NullableInlineResponse200290CountsByStatus{value: val, isSet: true}
}

func (v NullableInlineResponse200290CountsByStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200290CountsByStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


