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

// InlineResponse200286StatesExpiring Data for expiring licenses
type InlineResponse200286StatesExpiring struct {
	// The number of expiring licenses
	Count *int32 `json:"count,omitempty"`
	Critical *InlineResponse200286StatesExpiringCritical `json:"critical,omitempty"`
	Warning *InlineResponse200286StatesExpiringWarning `json:"warning,omitempty"`
}

// NewInlineResponse200286StatesExpiring instantiates a new InlineResponse200286StatesExpiring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200286StatesExpiring() *InlineResponse200286StatesExpiring {
	this := InlineResponse200286StatesExpiring{}
	return &this
}

// NewInlineResponse200286StatesExpiringWithDefaults instantiates a new InlineResponse200286StatesExpiring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200286StatesExpiringWithDefaults() *InlineResponse200286StatesExpiring {
	this := InlineResponse200286StatesExpiring{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineResponse200286StatesExpiring) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200286StatesExpiring) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineResponse200286StatesExpiring) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineResponse200286StatesExpiring) SetCount(v int32) {
	o.Count = &v
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *InlineResponse200286StatesExpiring) GetCritical() InlineResponse200286StatesExpiringCritical {
	if o == nil || isNil(o.Critical) {
		var ret InlineResponse200286StatesExpiringCritical
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200286StatesExpiring) GetCriticalOk() (*InlineResponse200286StatesExpiringCritical, bool) {
	if o == nil || isNil(o.Critical) {
    return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *InlineResponse200286StatesExpiring) HasCritical() bool {
	if o != nil && !isNil(o.Critical) {
		return true
	}

	return false
}

// SetCritical gets a reference to the given InlineResponse200286StatesExpiringCritical and assigns it to the Critical field.
func (o *InlineResponse200286StatesExpiring) SetCritical(v InlineResponse200286StatesExpiringCritical) {
	o.Critical = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *InlineResponse200286StatesExpiring) GetWarning() InlineResponse200286StatesExpiringWarning {
	if o == nil || isNil(o.Warning) {
		var ret InlineResponse200286StatesExpiringWarning
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200286StatesExpiring) GetWarningOk() (*InlineResponse200286StatesExpiringWarning, bool) {
	if o == nil || isNil(o.Warning) {
    return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *InlineResponse200286StatesExpiring) HasWarning() bool {
	if o != nil && !isNil(o.Warning) {
		return true
	}

	return false
}

// SetWarning gets a reference to the given InlineResponse200286StatesExpiringWarning and assigns it to the Warning field.
func (o *InlineResponse200286StatesExpiring) SetWarning(v InlineResponse200286StatesExpiringWarning) {
	o.Warning = &v
}

func (o InlineResponse200286StatesExpiring) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Critical) {
		toSerialize["critical"] = o.Critical
	}
	if !isNil(o.Warning) {
		toSerialize["warning"] = o.Warning
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200286StatesExpiring struct {
	value *InlineResponse200286StatesExpiring
	isSet bool
}

func (v NullableInlineResponse200286StatesExpiring) Get() *InlineResponse200286StatesExpiring {
	return v.value
}

func (v *NullableInlineResponse200286StatesExpiring) Set(val *InlineResponse200286StatesExpiring) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200286StatesExpiring) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200286StatesExpiring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200286StatesExpiring(val *InlineResponse200286StatesExpiring) *NullableInlineResponse200286StatesExpiring {
	return &NullableInlineResponse200286StatesExpiring{value: val, isSet: true}
}

func (v NullableInlineResponse200286StatesExpiring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200286StatesExpiring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


