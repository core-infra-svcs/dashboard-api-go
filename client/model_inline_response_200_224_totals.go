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

// InlineResponse200224Totals Totals by Severity
type InlineResponse200224Totals struct {
	// Informational Severity Count
	Informational *int32 `json:"informational,omitempty"`
	// Warning Severity Count
	Warning *int32 `json:"warning,omitempty"`
	// Critical Severity Count
	Critical *int32 `json:"critical,omitempty"`
}

// NewInlineResponse200224Totals instantiates a new InlineResponse200224Totals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200224Totals() *InlineResponse200224Totals {
	this := InlineResponse200224Totals{}
	return &this
}

// NewInlineResponse200224TotalsWithDefaults instantiates a new InlineResponse200224Totals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200224TotalsWithDefaults() *InlineResponse200224Totals {
	this := InlineResponse200224Totals{}
	return &this
}

// GetInformational returns the Informational field value if set, zero value otherwise.
func (o *InlineResponse200224Totals) GetInformational() int32 {
	if o == nil || isNil(o.Informational) {
		var ret int32
		return ret
	}
	return *o.Informational
}

// GetInformationalOk returns a tuple with the Informational field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200224Totals) GetInformationalOk() (*int32, bool) {
	if o == nil || isNil(o.Informational) {
    return nil, false
	}
	return o.Informational, true
}

// HasInformational returns a boolean if a field has been set.
func (o *InlineResponse200224Totals) HasInformational() bool {
	if o != nil && !isNil(o.Informational) {
		return true
	}

	return false
}

// SetInformational gets a reference to the given int32 and assigns it to the Informational field.
func (o *InlineResponse200224Totals) SetInformational(v int32) {
	o.Informational = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *InlineResponse200224Totals) GetWarning() int32 {
	if o == nil || isNil(o.Warning) {
		var ret int32
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200224Totals) GetWarningOk() (*int32, bool) {
	if o == nil || isNil(o.Warning) {
    return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *InlineResponse200224Totals) HasWarning() bool {
	if o != nil && !isNil(o.Warning) {
		return true
	}

	return false
}

// SetWarning gets a reference to the given int32 and assigns it to the Warning field.
func (o *InlineResponse200224Totals) SetWarning(v int32) {
	o.Warning = &v
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *InlineResponse200224Totals) GetCritical() int32 {
	if o == nil || isNil(o.Critical) {
		var ret int32
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200224Totals) GetCriticalOk() (*int32, bool) {
	if o == nil || isNil(o.Critical) {
    return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *InlineResponse200224Totals) HasCritical() bool {
	if o != nil && !isNil(o.Critical) {
		return true
	}

	return false
}

// SetCritical gets a reference to the given int32 and assigns it to the Critical field.
func (o *InlineResponse200224Totals) SetCritical(v int32) {
	o.Critical = &v
}

func (o InlineResponse200224Totals) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Informational) {
		toSerialize["informational"] = o.Informational
	}
	if !isNil(o.Warning) {
		toSerialize["warning"] = o.Warning
	}
	if !isNil(o.Critical) {
		toSerialize["critical"] = o.Critical
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200224Totals struct {
	value *InlineResponse200224Totals
	isSet bool
}

func (v NullableInlineResponse200224Totals) Get() *InlineResponse200224Totals {
	return v.value
}

func (v *NullableInlineResponse200224Totals) Set(val *InlineResponse200224Totals) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200224Totals) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200224Totals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200224Totals(val *InlineResponse200224Totals) *NullableInlineResponse200224Totals {
	return &NullableInlineResponse200224Totals{value: val, isSet: true}
}

func (v NullableInlineResponse200224Totals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200224Totals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


