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

// InlineResponse200114 struct for InlineResponse200114
type InlineResponse200114 struct {
	// General EAP timeout in seconds.
	Timeout *int32 `json:"timeout,omitempty"`
	// Maximum number of general EAP retries.
	MaxRetries *int32 `json:"maxRetries,omitempty"`
	Identity *InlineResponse200114Identity `json:"identity,omitempty"`
	EapolKey *InlineResponse200114EapolKey `json:"eapolKey,omitempty"`
}

// NewInlineResponse200114 instantiates a new InlineResponse200114 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200114() *InlineResponse200114 {
	this := InlineResponse200114{}
	return &this
}

// NewInlineResponse200114WithDefaults instantiates a new InlineResponse200114 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200114WithDefaults() *InlineResponse200114 {
	this := InlineResponse200114{}
	return &this
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *InlineResponse200114) GetTimeout() int32 {
	if o == nil || isNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200114) GetTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.Timeout) {
    return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *InlineResponse200114) HasTimeout() bool {
	if o != nil && !isNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *InlineResponse200114) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise.
func (o *InlineResponse200114) GetMaxRetries() int32 {
	if o == nil || isNil(o.MaxRetries) {
		var ret int32
		return ret
	}
	return *o.MaxRetries
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200114) GetMaxRetriesOk() (*int32, bool) {
	if o == nil || isNil(o.MaxRetries) {
    return nil, false
	}
	return o.MaxRetries, true
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *InlineResponse200114) HasMaxRetries() bool {
	if o != nil && !isNil(o.MaxRetries) {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given int32 and assigns it to the MaxRetries field.
func (o *InlineResponse200114) SetMaxRetries(v int32) {
	o.MaxRetries = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *InlineResponse200114) GetIdentity() InlineResponse200114Identity {
	if o == nil || isNil(o.Identity) {
		var ret InlineResponse200114Identity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200114) GetIdentityOk() (*InlineResponse200114Identity, bool) {
	if o == nil || isNil(o.Identity) {
    return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *InlineResponse200114) HasIdentity() bool {
	if o != nil && !isNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given InlineResponse200114Identity and assigns it to the Identity field.
func (o *InlineResponse200114) SetIdentity(v InlineResponse200114Identity) {
	o.Identity = &v
}

// GetEapolKey returns the EapolKey field value if set, zero value otherwise.
func (o *InlineResponse200114) GetEapolKey() InlineResponse200114EapolKey {
	if o == nil || isNil(o.EapolKey) {
		var ret InlineResponse200114EapolKey
		return ret
	}
	return *o.EapolKey
}

// GetEapolKeyOk returns a tuple with the EapolKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200114) GetEapolKeyOk() (*InlineResponse200114EapolKey, bool) {
	if o == nil || isNil(o.EapolKey) {
    return nil, false
	}
	return o.EapolKey, true
}

// HasEapolKey returns a boolean if a field has been set.
func (o *InlineResponse200114) HasEapolKey() bool {
	if o != nil && !isNil(o.EapolKey) {
		return true
	}

	return false
}

// SetEapolKey gets a reference to the given InlineResponse200114EapolKey and assigns it to the EapolKey field.
func (o *InlineResponse200114) SetEapolKey(v InlineResponse200114EapolKey) {
	o.EapolKey = &v
}

func (o InlineResponse200114) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200114 struct {
	value *InlineResponse200114
	isSet bool
}

func (v NullableInlineResponse200114) Get() *InlineResponse200114 {
	return v.value
}

func (v *NullableInlineResponse200114) Set(val *InlineResponse200114) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200114) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200114) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200114(val *InlineResponse200114) *NullableInlineResponse200114 {
	return &NullableInlineResponse200114{value: val, isSet: true}
}

func (v NullableInlineResponse200114) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200114) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


