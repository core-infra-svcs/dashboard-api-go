/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200195 struct for InlineResponse200195
type InlineResponse200195 struct {
	// General EAP timeout in seconds.
	Timeout *int32 `json:"timeout,omitempty"`
	// Maximum number of general EAP retries.
	MaxRetries *int32 `json:"maxRetries,omitempty"`
	Identity *InlineResponse200195Identity `json:"identity,omitempty"`
	EapolKey *InlineResponse200195EapolKey `json:"eapolKey,omitempty"`
}

// NewInlineResponse200195 instantiates a new InlineResponse200195 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200195() *InlineResponse200195 {
	this := InlineResponse200195{}
	return &this
}

// NewInlineResponse200195WithDefaults instantiates a new InlineResponse200195 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200195WithDefaults() *InlineResponse200195 {
	this := InlineResponse200195{}
	return &this
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *InlineResponse200195) GetTimeout() int32 {
	if o == nil || isNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200195) GetTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.Timeout) {
    return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *InlineResponse200195) HasTimeout() bool {
	if o != nil && !isNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *InlineResponse200195) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise.
func (o *InlineResponse200195) GetMaxRetries() int32 {
	if o == nil || isNil(o.MaxRetries) {
		var ret int32
		return ret
	}
	return *o.MaxRetries
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200195) GetMaxRetriesOk() (*int32, bool) {
	if o == nil || isNil(o.MaxRetries) {
    return nil, false
	}
	return o.MaxRetries, true
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *InlineResponse200195) HasMaxRetries() bool {
	if o != nil && !isNil(o.MaxRetries) {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given int32 and assigns it to the MaxRetries field.
func (o *InlineResponse200195) SetMaxRetries(v int32) {
	o.MaxRetries = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *InlineResponse200195) GetIdentity() InlineResponse200195Identity {
	if o == nil || isNil(o.Identity) {
		var ret InlineResponse200195Identity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200195) GetIdentityOk() (*InlineResponse200195Identity, bool) {
	if o == nil || isNil(o.Identity) {
    return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *InlineResponse200195) HasIdentity() bool {
	if o != nil && !isNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given InlineResponse200195Identity and assigns it to the Identity field.
func (o *InlineResponse200195) SetIdentity(v InlineResponse200195Identity) {
	o.Identity = &v
}

// GetEapolKey returns the EapolKey field value if set, zero value otherwise.
func (o *InlineResponse200195) GetEapolKey() InlineResponse200195EapolKey {
	if o == nil || isNil(o.EapolKey) {
		var ret InlineResponse200195EapolKey
		return ret
	}
	return *o.EapolKey
}

// GetEapolKeyOk returns a tuple with the EapolKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200195) GetEapolKeyOk() (*InlineResponse200195EapolKey, bool) {
	if o == nil || isNil(o.EapolKey) {
    return nil, false
	}
	return o.EapolKey, true
}

// HasEapolKey returns a boolean if a field has been set.
func (o *InlineResponse200195) HasEapolKey() bool {
	if o != nil && !isNil(o.EapolKey) {
		return true
	}

	return false
}

// SetEapolKey gets a reference to the given InlineResponse200195EapolKey and assigns it to the EapolKey field.
func (o *InlineResponse200195) SetEapolKey(v InlineResponse200195EapolKey) {
	o.EapolKey = &v
}

func (o InlineResponse200195) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !isNil(o.MaxRetries) {
		toSerialize["maxRetries"] = o.MaxRetries
	}
	if !isNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !isNil(o.EapolKey) {
		toSerialize["eapolKey"] = o.EapolKey
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200195 struct {
	value *InlineResponse200195
	isSet bool
}

func (v NullableInlineResponse200195) Get() *InlineResponse200195 {
	return v.value
}

func (v *NullableInlineResponse200195) Set(val *InlineResponse200195) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200195) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200195) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200195(val *InlineResponse200195) *NullableInlineResponse200195 {
	return &NullableInlineResponse200195{value: val, isSet: true}
}

func (v NullableInlineResponse200195) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200195) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


