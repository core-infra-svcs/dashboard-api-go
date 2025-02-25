/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200349Changes struct for InlineResponse200349Changes
type InlineResponse200349Changes struct {
	// The timestamp of current status of the interface
	Ts *time.Time `json:"ts,omitempty"`
	// The status of the interface
	Status *string `json:"status,omitempty"`
	// All warnings present on the port
	Warnings []string `json:"warnings,omitempty"`
	// All errors present on the port
	Errors []string `json:"errors,omitempty"`
}

// NewInlineResponse200349Changes instantiates a new InlineResponse200349Changes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200349Changes() *InlineResponse200349Changes {
	this := InlineResponse200349Changes{}
	return &this
}

// NewInlineResponse200349ChangesWithDefaults instantiates a new InlineResponse200349Changes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200349ChangesWithDefaults() *InlineResponse200349Changes {
	this := InlineResponse200349Changes{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse200349Changes) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200349Changes) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse200349Changes) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *InlineResponse200349Changes) SetTs(v time.Time) {
	o.Ts = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse200349Changes) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200349Changes) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse200349Changes) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse200349Changes) SetStatus(v string) {
	o.Status = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *InlineResponse200349Changes) GetWarnings() []string {
	if o == nil || isNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200349Changes) GetWarningsOk() ([]string, bool) {
	if o == nil || isNil(o.Warnings) {
    return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *InlineResponse200349Changes) HasWarnings() bool {
	if o != nil && !isNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *InlineResponse200349Changes) SetWarnings(v []string) {
	o.Warnings = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *InlineResponse200349Changes) GetErrors() []string {
	if o == nil || isNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200349Changes) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
    return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *InlineResponse200349Changes) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *InlineResponse200349Changes) SetErrors(v []string) {
	o.Errors = v
}

func (o InlineResponse200349Changes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200349Changes struct {
	value *InlineResponse200349Changes
	isSet bool
}

func (v NullableInlineResponse200349Changes) Get() *InlineResponse200349Changes {
	return v.value
}

func (v *NullableInlineResponse200349Changes) Set(val *InlineResponse200349Changes) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200349Changes) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200349Changes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200349Changes(val *InlineResponse200349Changes) *NullableInlineResponse200349Changes {
	return &NullableInlineResponse200349Changes{value: val, isSet: true}
}

func (v NullableInlineResponse200349Changes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200349Changes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


