/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200184 struct for InlineResponse200184
type InlineResponse200184 struct {
	// General EAP timeout in seconds.
	Timeout *int32 `json:"timeout,omitempty"`
	// Maximum number of general EAP retries.
	MaxRetries *int32 `json:"maxRetries,omitempty"`
	Identity *InlineResponse200184Identity `json:"identity,omitempty"`
	EapolKey *InlineResponse200184EapolKey `json:"eapolKey,omitempty"`
}

// NewInlineResponse200184 instantiates a new InlineResponse200184 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200184() *InlineResponse200184 {
	this := InlineResponse200184{}
	return &this
}

// NewInlineResponse200184WithDefaults instantiates a new InlineResponse200184 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200184WithDefaults() *InlineResponse200184 {
	this := InlineResponse200184{}
	return &this
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *InlineResponse200184) GetTimeout() int32 {
	if o == nil || isNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200184) GetTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.Timeout) {
    return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *InlineResponse200184) HasTimeout() bool {
	if o != nil && !isNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *InlineResponse200184) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise.
func (o *InlineResponse200184) GetMaxRetries() int32 {
	if o == nil || isNil(o.MaxRetries) {
		var ret int32
		return ret
	}
	return *o.MaxRetries
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200184) GetMaxRetriesOk() (*int32, bool) {
	if o == nil || isNil(o.MaxRetries) {
    return nil, false
	}
	return o.MaxRetries, true
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *InlineResponse200184) HasMaxRetries() bool {
	if o != nil && !isNil(o.MaxRetries) {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given int32 and assigns it to the MaxRetries field.
func (o *InlineResponse200184) SetMaxRetries(v int32) {
	o.MaxRetries = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *InlineResponse200184) GetIdentity() InlineResponse200184Identity {
	if o == nil || isNil(o.Identity) {
		var ret InlineResponse200184Identity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200184) GetIdentityOk() (*InlineResponse200184Identity, bool) {
	if o == nil || isNil(o.Identity) {
    return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *InlineResponse200184) HasIdentity() bool {
	if o != nil && !isNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given InlineResponse200184Identity and assigns it to the Identity field.
func (o *InlineResponse200184) SetIdentity(v InlineResponse200184Identity) {
	o.Identity = &v
}

// GetEapolKey returns the EapolKey field value if set, zero value otherwise.
func (o *InlineResponse200184) GetEapolKey() InlineResponse200184EapolKey {
	if o == nil || isNil(o.EapolKey) {
		var ret InlineResponse200184EapolKey
		return ret
	}
	return *o.EapolKey
}

// GetEapolKeyOk returns a tuple with the EapolKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200184) GetEapolKeyOk() (*InlineResponse200184EapolKey, bool) {
	if o == nil || isNil(o.EapolKey) {
    return nil, false
	}
	return o.EapolKey, true
}

// HasEapolKey returns a boolean if a field has been set.
func (o *InlineResponse200184) HasEapolKey() bool {
	if o != nil && !isNil(o.EapolKey) {
		return true
	}

	return false
}

// SetEapolKey gets a reference to the given InlineResponse200184EapolKey and assigns it to the EapolKey field.
func (o *InlineResponse200184) SetEapolKey(v InlineResponse200184EapolKey) {
	o.EapolKey = &v
}

func (o InlineResponse200184) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200184 struct {
	value *InlineResponse200184
	isSet bool
}

func (v NullableInlineResponse200184) Get() *InlineResponse200184 {
	return v.value
}

func (v *NullableInlineResponse200184) Set(val *InlineResponse200184) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200184) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200184) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200184(val *InlineResponse200184) *NullableInlineResponse200184 {
	return &NullableInlineResponse200184{value: val, isSet: true}
}

func (v NullableInlineResponse200184) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200184) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


