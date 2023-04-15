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

// InlineResponse20083 struct for InlineResponse20083
type InlineResponse20083 struct {
	// General EAP timeout in seconds.
	Timeout *int32 `json:"timeout,omitempty"`
	// Maximum number of general EAP retries.
	MaxRetries *int32 `json:"maxRetries,omitempty"`
	Identity *InlineResponse20083Identity `json:"identity,omitempty"`
	EapolKey *InlineResponse20083EapolKey `json:"eapolKey,omitempty"`
}

// NewInlineResponse20083 instantiates a new InlineResponse20083 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20083() *InlineResponse20083 {
	this := InlineResponse20083{}
	return &this
}

// NewInlineResponse20083WithDefaults instantiates a new InlineResponse20083 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20083WithDefaults() *InlineResponse20083 {
	this := InlineResponse20083{}
	return &this
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *InlineResponse20083) GetTimeout() int32 {
	if o == nil || isNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.Timeout) {
    return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *InlineResponse20083) HasTimeout() bool {
	if o != nil && !isNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *InlineResponse20083) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise.
func (o *InlineResponse20083) GetMaxRetries() int32 {
	if o == nil || isNil(o.MaxRetries) {
		var ret int32
		return ret
	}
	return *o.MaxRetries
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetMaxRetriesOk() (*int32, bool) {
	if o == nil || isNil(o.MaxRetries) {
    return nil, false
	}
	return o.MaxRetries, true
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *InlineResponse20083) HasMaxRetries() bool {
	if o != nil && !isNil(o.MaxRetries) {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given int32 and assigns it to the MaxRetries field.
func (o *InlineResponse20083) SetMaxRetries(v int32) {
	o.MaxRetries = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *InlineResponse20083) GetIdentity() InlineResponse20083Identity {
	if o == nil || isNil(o.Identity) {
		var ret InlineResponse20083Identity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetIdentityOk() (*InlineResponse20083Identity, bool) {
	if o == nil || isNil(o.Identity) {
    return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *InlineResponse20083) HasIdentity() bool {
	if o != nil && !isNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given InlineResponse20083Identity and assigns it to the Identity field.
func (o *InlineResponse20083) SetIdentity(v InlineResponse20083Identity) {
	o.Identity = &v
}

// GetEapolKey returns the EapolKey field value if set, zero value otherwise.
func (o *InlineResponse20083) GetEapolKey() InlineResponse20083EapolKey {
	if o == nil || isNil(o.EapolKey) {
		var ret InlineResponse20083EapolKey
		return ret
	}
	return *o.EapolKey
}

// GetEapolKeyOk returns a tuple with the EapolKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetEapolKeyOk() (*InlineResponse20083EapolKey, bool) {
	if o == nil || isNil(o.EapolKey) {
    return nil, false
	}
	return o.EapolKey, true
}

// HasEapolKey returns a boolean if a field has been set.
func (o *InlineResponse20083) HasEapolKey() bool {
	if o != nil && !isNil(o.EapolKey) {
		return true
	}

	return false
}

// SetEapolKey gets a reference to the given InlineResponse20083EapolKey and assigns it to the EapolKey field.
func (o *InlineResponse20083) SetEapolKey(v InlineResponse20083EapolKey) {
	o.EapolKey = &v
}

func (o InlineResponse20083) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20083 struct {
	value *InlineResponse20083
	isSet bool
}

func (v NullableInlineResponse20083) Get() *InlineResponse20083 {
	return v.value
}

func (v *NullableInlineResponse20083) Set(val *InlineResponse20083) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20083) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20083) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20083(val *InlineResponse20083) *NullableInlineResponse20083 {
	return &NullableInlineResponse20083{value: val, isSet: true}
}

func (v NullableInlineResponse20083) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20083) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


