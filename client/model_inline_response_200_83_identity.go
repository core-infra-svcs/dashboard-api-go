/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20083Identity EAP settings for identity requests.
type InlineResponse20083Identity struct {
	// Maximum number of EAP retries.
	Retries *int32 `json:"retries,omitempty"`
	// EAP timeout in seconds.
	Timeout *int32 `json:"timeout,omitempty"`
}

// NewInlineResponse20083Identity instantiates a new InlineResponse20083Identity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20083Identity() *InlineResponse20083Identity {
	this := InlineResponse20083Identity{}
	return &this
}

// NewInlineResponse20083IdentityWithDefaults instantiates a new InlineResponse20083Identity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20083IdentityWithDefaults() *InlineResponse20083Identity {
	this := InlineResponse20083Identity{}
	return &this
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *InlineResponse20083Identity) GetRetries() int32 {
	if o == nil || isNil(o.Retries) {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083Identity) GetRetriesOk() (*int32, bool) {
	if o == nil || isNil(o.Retries) {
    return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *InlineResponse20083Identity) HasRetries() bool {
	if o != nil && !isNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *InlineResponse20083Identity) SetRetries(v int32) {
	o.Retries = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *InlineResponse20083Identity) GetTimeout() int32 {
	if o == nil || isNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083Identity) GetTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.Timeout) {
    return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *InlineResponse20083Identity) HasTimeout() bool {
	if o != nil && !isNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *InlineResponse20083Identity) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o InlineResponse20083Identity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Retries) {
		toSerialize["retries"] = o.Retries
	}
	if !isNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20083Identity struct {
	value *InlineResponse20083Identity
	isSet bool
}

func (v NullableInlineResponse20083Identity) Get() *InlineResponse20083Identity {
	return v.value
}

func (v *NullableInlineResponse20083Identity) Set(val *InlineResponse20083Identity) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20083Identity) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20083Identity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20083Identity(val *InlineResponse20083Identity) *NullableInlineResponse20083Identity {
	return &NullableInlineResponse20083Identity{value: val, isSet: true}
}

func (v NullableInlineResponse20083Identity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20083Identity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


