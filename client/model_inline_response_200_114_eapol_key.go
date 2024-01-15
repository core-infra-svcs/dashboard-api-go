/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200114EapolKey EAPOL Key settings.
type InlineResponse200114EapolKey struct {
	// Maximum number of EAPOL key retries.
	Retries *int32 `json:"retries,omitempty"`
	// EAPOL Key timeout in milliseconds.
	TimeoutInMs *int32 `json:"timeoutInMs,omitempty"`
}

// NewInlineResponse200114EapolKey instantiates a new InlineResponse200114EapolKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200114EapolKey() *InlineResponse200114EapolKey {
	this := InlineResponse200114EapolKey{}
	return &this
}

// NewInlineResponse200114EapolKeyWithDefaults instantiates a new InlineResponse200114EapolKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200114EapolKeyWithDefaults() *InlineResponse200114EapolKey {
	this := InlineResponse200114EapolKey{}
	return &this
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *InlineResponse200114EapolKey) GetRetries() int32 {
	if o == nil || isNil(o.Retries) {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200114EapolKey) GetRetriesOk() (*int32, bool) {
	if o == nil || isNil(o.Retries) {
    return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *InlineResponse200114EapolKey) HasRetries() bool {
	if o != nil && !isNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *InlineResponse200114EapolKey) SetRetries(v int32) {
	o.Retries = &v
}

// GetTimeoutInMs returns the TimeoutInMs field value if set, zero value otherwise.
func (o *InlineResponse200114EapolKey) GetTimeoutInMs() int32 {
	if o == nil || isNil(o.TimeoutInMs) {
		var ret int32
		return ret
	}
	return *o.TimeoutInMs
}

// GetTimeoutInMsOk returns a tuple with the TimeoutInMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200114EapolKey) GetTimeoutInMsOk() (*int32, bool) {
	if o == nil || isNil(o.TimeoutInMs) {
    return nil, false
	}
	return o.TimeoutInMs, true
}

// HasTimeoutInMs returns a boolean if a field has been set.
func (o *InlineResponse200114EapolKey) HasTimeoutInMs() bool {
	if o != nil && !isNil(o.TimeoutInMs) {
		return true
	}

	return false
}

// SetTimeoutInMs gets a reference to the given int32 and assigns it to the TimeoutInMs field.
func (o *InlineResponse200114EapolKey) SetTimeoutInMs(v int32) {
	o.TimeoutInMs = &v
}

func (o InlineResponse200114EapolKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Retries) {
		toSerialize["retries"] = o.Retries
	}
	if !isNil(o.TimeoutInMs) {
		toSerialize["timeoutInMs"] = o.TimeoutInMs
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200114EapolKey struct {
	value *InlineResponse200114EapolKey
	isSet bool
}

func (v NullableInlineResponse200114EapolKey) Get() *InlineResponse200114EapolKey {
	return v.value
}

func (v *NullableInlineResponse200114EapolKey) Set(val *InlineResponse200114EapolKey) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200114EapolKey) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200114EapolKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200114EapolKey(val *InlineResponse200114EapolKey) *NullableInlineResponse200114EapolKey {
	return &NullableInlineResponse200114EapolKey{value: val, isSet: true}
}

func (v NullableInlineResponse200114EapolKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200114EapolKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


