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

// InlineResponse200305CountsByStatus The count data, indexed by active or inactive status
type InlineResponse200305CountsByStatus struct {
	Active *InlineResponse200305CountsByStatusActive `json:"active,omitempty"`
	Inactive *InlineResponse200305CountsByStatusInactive `json:"inactive,omitempty"`
}

// NewInlineResponse200305CountsByStatus instantiates a new InlineResponse200305CountsByStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200305CountsByStatus() *InlineResponse200305CountsByStatus {
	this := InlineResponse200305CountsByStatus{}
	return &this
}

// NewInlineResponse200305CountsByStatusWithDefaults instantiates a new InlineResponse200305CountsByStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200305CountsByStatusWithDefaults() *InlineResponse200305CountsByStatus {
	this := InlineResponse200305CountsByStatus{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *InlineResponse200305CountsByStatus) GetActive() InlineResponse200305CountsByStatusActive {
	if o == nil || isNil(o.Active) {
		var ret InlineResponse200305CountsByStatusActive
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200305CountsByStatus) GetActiveOk() (*InlineResponse200305CountsByStatusActive, bool) {
	if o == nil || isNil(o.Active) {
    return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *InlineResponse200305CountsByStatus) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given InlineResponse200305CountsByStatusActive and assigns it to the Active field.
func (o *InlineResponse200305CountsByStatus) SetActive(v InlineResponse200305CountsByStatusActive) {
	o.Active = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *InlineResponse200305CountsByStatus) GetInactive() InlineResponse200305CountsByStatusInactive {
	if o == nil || isNil(o.Inactive) {
		var ret InlineResponse200305CountsByStatusInactive
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200305CountsByStatus) GetInactiveOk() (*InlineResponse200305CountsByStatusInactive, bool) {
	if o == nil || isNil(o.Inactive) {
    return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *InlineResponse200305CountsByStatus) HasInactive() bool {
	if o != nil && !isNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given InlineResponse200305CountsByStatusInactive and assigns it to the Inactive field.
func (o *InlineResponse200305CountsByStatus) SetInactive(v InlineResponse200305CountsByStatusInactive) {
	o.Inactive = &v
}

func (o InlineResponse200305CountsByStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200305CountsByStatus struct {
	value *InlineResponse200305CountsByStatus
	isSet bool
}

func (v NullableInlineResponse200305CountsByStatus) Get() *InlineResponse200305CountsByStatus {
	return v.value
}

func (v *NullableInlineResponse200305CountsByStatus) Set(val *InlineResponse200305CountsByStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200305CountsByStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200305CountsByStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200305CountsByStatus(val *InlineResponse200305CountsByStatus) *NullableInlineResponse200305CountsByStatus {
	return &NullableInlineResponse200305CountsByStatus{value: val, isSet: true}
}

func (v NullableInlineResponse200305CountsByStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200305CountsByStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


