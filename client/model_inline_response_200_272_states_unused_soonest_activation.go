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

// InlineResponse200272StatesUnusedSoonestActivation Information about the soonest forthcoming license activation
type InlineResponse200272StatesUnusedSoonestActivation struct {
	// The soonest license activation date
	ActivationDate *string `json:"activationDate,omitempty"`
	// The number of licenses that will activate on this date
	ToActivateCount *int32 `json:"toActivateCount,omitempty"`
}

// NewInlineResponse200272StatesUnusedSoonestActivation instantiates a new InlineResponse200272StatesUnusedSoonestActivation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200272StatesUnusedSoonestActivation() *InlineResponse200272StatesUnusedSoonestActivation {
	this := InlineResponse200272StatesUnusedSoonestActivation{}
	return &this
}

// NewInlineResponse200272StatesUnusedSoonestActivationWithDefaults instantiates a new InlineResponse200272StatesUnusedSoonestActivation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200272StatesUnusedSoonestActivationWithDefaults() *InlineResponse200272StatesUnusedSoonestActivation {
	this := InlineResponse200272StatesUnusedSoonestActivation{}
	return &this
}

// GetActivationDate returns the ActivationDate field value if set, zero value otherwise.
func (o *InlineResponse200272StatesUnusedSoonestActivation) GetActivationDate() string {
	if o == nil || isNil(o.ActivationDate) {
		var ret string
		return ret
	}
	return *o.ActivationDate
}

// GetActivationDateOk returns a tuple with the ActivationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200272StatesUnusedSoonestActivation) GetActivationDateOk() (*string, bool) {
	if o == nil || isNil(o.ActivationDate) {
    return nil, false
	}
	return o.ActivationDate, true
}

// HasActivationDate returns a boolean if a field has been set.
func (o *InlineResponse200272StatesUnusedSoonestActivation) HasActivationDate() bool {
	if o != nil && !isNil(o.ActivationDate) {
		return true
	}

	return false
}

// SetActivationDate gets a reference to the given string and assigns it to the ActivationDate field.
func (o *InlineResponse200272StatesUnusedSoonestActivation) SetActivationDate(v string) {
	o.ActivationDate = &v
}

// GetToActivateCount returns the ToActivateCount field value if set, zero value otherwise.
func (o *InlineResponse200272StatesUnusedSoonestActivation) GetToActivateCount() int32 {
	if o == nil || isNil(o.ToActivateCount) {
		var ret int32
		return ret
	}
	return *o.ToActivateCount
}

// GetToActivateCountOk returns a tuple with the ToActivateCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200272StatesUnusedSoonestActivation) GetToActivateCountOk() (*int32, bool) {
	if o == nil || isNil(o.ToActivateCount) {
    return nil, false
	}
	return o.ToActivateCount, true
}

// HasToActivateCount returns a boolean if a field has been set.
func (o *InlineResponse200272StatesUnusedSoonestActivation) HasToActivateCount() bool {
	if o != nil && !isNil(o.ToActivateCount) {
		return true
	}

	return false
}

// SetToActivateCount gets a reference to the given int32 and assigns it to the ToActivateCount field.
func (o *InlineResponse200272StatesUnusedSoonestActivation) SetToActivateCount(v int32) {
	o.ToActivateCount = &v
}

func (o InlineResponse200272StatesUnusedSoonestActivation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ActivationDate) {
		toSerialize["activationDate"] = o.ActivationDate
	}
	if !isNil(o.ToActivateCount) {
		toSerialize["toActivateCount"] = o.ToActivateCount
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200272StatesUnusedSoonestActivation struct {
	value *InlineResponse200272StatesUnusedSoonestActivation
	isSet bool
}

func (v NullableInlineResponse200272StatesUnusedSoonestActivation) Get() *InlineResponse200272StatesUnusedSoonestActivation {
	return v.value
}

func (v *NullableInlineResponse200272StatesUnusedSoonestActivation) Set(val *InlineResponse200272StatesUnusedSoonestActivation) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200272StatesUnusedSoonestActivation) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200272StatesUnusedSoonestActivation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200272StatesUnusedSoonestActivation(val *InlineResponse200272StatesUnusedSoonestActivation) *NullableInlineResponse200272StatesUnusedSoonestActivation {
	return &NullableInlineResponse200272StatesUnusedSoonestActivation{value: val, isSet: true}
}

func (v NullableInlineResponse200272StatesUnusedSoonestActivation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200272StatesUnusedSoonestActivation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

