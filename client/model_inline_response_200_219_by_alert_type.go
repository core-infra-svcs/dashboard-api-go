/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200219ByAlertType struct for InlineResponse200219ByAlertType
type InlineResponse200219ByAlertType struct {
	// Alert Type
	Type *string `json:"type,omitempty"`
	// Informational Severity Count
	Informational *int32 `json:"informational,omitempty"`
	// Warning Severity Count
	Warning *int32 `json:"warning,omitempty"`
	// Critical Severity Count
	Critical *int32 `json:"critical,omitempty"`
}

// NewInlineResponse200219ByAlertType instantiates a new InlineResponse200219ByAlertType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200219ByAlertType() *InlineResponse200219ByAlertType {
	this := InlineResponse200219ByAlertType{}
	return &this
}

// NewInlineResponse200219ByAlertTypeWithDefaults instantiates a new InlineResponse200219ByAlertType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200219ByAlertTypeWithDefaults() *InlineResponse200219ByAlertType {
	this := InlineResponse200219ByAlertType{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200219ByAlertType) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200219ByAlertType) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200219ByAlertType) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200219ByAlertType) SetType(v string) {
	o.Type = &v
}

// GetInformational returns the Informational field value if set, zero value otherwise.
func (o *InlineResponse200219ByAlertType) GetInformational() int32 {
	if o == nil || isNil(o.Informational) {
		var ret int32
		return ret
	}
	return *o.Informational
}

// GetInformationalOk returns a tuple with the Informational field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200219ByAlertType) GetInformationalOk() (*int32, bool) {
	if o == nil || isNil(o.Informational) {
    return nil, false
	}
	return o.Informational, true
}

// HasInformational returns a boolean if a field has been set.
func (o *InlineResponse200219ByAlertType) HasInformational() bool {
	if o != nil && !isNil(o.Informational) {
		return true
	}

	return false
}

// SetInformational gets a reference to the given int32 and assigns it to the Informational field.
func (o *InlineResponse200219ByAlertType) SetInformational(v int32) {
	o.Informational = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *InlineResponse200219ByAlertType) GetWarning() int32 {
	if o == nil || isNil(o.Warning) {
		var ret int32
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200219ByAlertType) GetWarningOk() (*int32, bool) {
	if o == nil || isNil(o.Warning) {
    return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *InlineResponse200219ByAlertType) HasWarning() bool {
	if o != nil && !isNil(o.Warning) {
		return true
	}

	return false
}

// SetWarning gets a reference to the given int32 and assigns it to the Warning field.
func (o *InlineResponse200219ByAlertType) SetWarning(v int32) {
	o.Warning = &v
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *InlineResponse200219ByAlertType) GetCritical() int32 {
	if o == nil || isNil(o.Critical) {
		var ret int32
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200219ByAlertType) GetCriticalOk() (*int32, bool) {
	if o == nil || isNil(o.Critical) {
    return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *InlineResponse200219ByAlertType) HasCritical() bool {
	if o != nil && !isNil(o.Critical) {
		return true
	}

	return false
}

// SetCritical gets a reference to the given int32 and assigns it to the Critical field.
func (o *InlineResponse200219ByAlertType) SetCritical(v int32) {
	o.Critical = &v
}

func (o InlineResponse200219ByAlertType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
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

type NullableInlineResponse200219ByAlertType struct {
	value *InlineResponse200219ByAlertType
	isSet bool
}

func (v NullableInlineResponse200219ByAlertType) Get() *InlineResponse200219ByAlertType {
	return v.value
}

func (v *NullableInlineResponse200219ByAlertType) Set(val *InlineResponse200219ByAlertType) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200219ByAlertType) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200219ByAlertType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200219ByAlertType(val *InlineResponse200219ByAlertType) *NullableInlineResponse200219ByAlertType {
	return &NullableInlineResponse200219ByAlertType{value: val, isSet: true}
}

func (v NullableInlineResponse200219ByAlertType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200219ByAlertType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


