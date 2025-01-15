/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200195EapolKey EAPOL Key settings.
type InlineResponse200195EapolKey struct {
	// Maximum number of EAPOL key retries.
	Retries *int32 `json:"retries,omitempty"`
	// EAPOL Key timeout in milliseconds.
	TimeoutInMs *int32 `json:"timeoutInMs,omitempty"`
}

// NewInlineResponse200195EapolKey instantiates a new InlineResponse200195EapolKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200195EapolKey() *InlineResponse200195EapolKey {
	this := InlineResponse200195EapolKey{}
	return &this
}

// NewInlineResponse200195EapolKeyWithDefaults instantiates a new InlineResponse200195EapolKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200195EapolKeyWithDefaults() *InlineResponse200195EapolKey {
	this := InlineResponse200195EapolKey{}
	return &this
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *InlineResponse200195EapolKey) GetRetries() int32 {
	if o == nil || isNil(o.Retries) {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200195EapolKey) GetRetriesOk() (*int32, bool) {
	if o == nil || isNil(o.Retries) {
    return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *InlineResponse200195EapolKey) HasRetries() bool {
	if o != nil && !isNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *InlineResponse200195EapolKey) SetRetries(v int32) {
	o.Retries = &v
}

// GetTimeoutInMs returns the TimeoutInMs field value if set, zero value otherwise.
func (o *InlineResponse200195EapolKey) GetTimeoutInMs() int32 {
	if o == nil || isNil(o.TimeoutInMs) {
		var ret int32
		return ret
	}
	return *o.TimeoutInMs
}

// GetTimeoutInMsOk returns a tuple with the TimeoutInMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200195EapolKey) GetTimeoutInMsOk() (*int32, bool) {
	if o == nil || isNil(o.TimeoutInMs) {
    return nil, false
	}
	return o.TimeoutInMs, true
}

// HasTimeoutInMs returns a boolean if a field has been set.
func (o *InlineResponse200195EapolKey) HasTimeoutInMs() bool {
	if o != nil && !isNil(o.TimeoutInMs) {
		return true
	}

	return false
}

// SetTimeoutInMs gets a reference to the given int32 and assigns it to the TimeoutInMs field.
func (o *InlineResponse200195EapolKey) SetTimeoutInMs(v int32) {
	o.TimeoutInMs = &v
}

func (o InlineResponse200195EapolKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Retries) {
		toSerialize["retries"] = o.Retries
	}
	if !isNil(o.TimeoutInMs) {
		toSerialize["timeoutInMs"] = o.TimeoutInMs
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200195EapolKey struct {
	value *InlineResponse200195EapolKey
	isSet bool
}

func (v NullableInlineResponse200195EapolKey) Get() *InlineResponse200195EapolKey {
	return v.value
}

func (v *NullableInlineResponse200195EapolKey) Set(val *InlineResponse200195EapolKey) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200195EapolKey) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200195EapolKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200195EapolKey(val *InlineResponse200195EapolKey) *NullableInlineResponse200195EapolKey {
	return &NullableInlineResponse200195EapolKey{value: val, isSet: true}
}

func (v NullableInlineResponse200195EapolKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200195EapolKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


