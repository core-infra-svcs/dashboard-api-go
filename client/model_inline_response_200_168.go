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

// InlineResponse200168 struct for InlineResponse200168
type InlineResponse200168 struct {
	// The network ID
	NetworkId *string `json:"networkId,omitempty"`
	// Indicates whether or not clients are allowed to       connect to rogue SSIDs. (blocked by default)
	DefaultPolicy *string `json:"defaultPolicy,omitempty"`
}

// NewInlineResponse200168 instantiates a new InlineResponse200168 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200168() *InlineResponse200168 {
	this := InlineResponse200168{}
	return &this
}

// NewInlineResponse200168WithDefaults instantiates a new InlineResponse200168 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200168WithDefaults() *InlineResponse200168 {
	this := InlineResponse200168{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200168) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200168) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200168) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200168) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetDefaultPolicy returns the DefaultPolicy field value if set, zero value otherwise.
func (o *InlineResponse200168) GetDefaultPolicy() string {
	if o == nil || isNil(o.DefaultPolicy) {
		var ret string
		return ret
	}
	return *o.DefaultPolicy
}

// GetDefaultPolicyOk returns a tuple with the DefaultPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200168) GetDefaultPolicyOk() (*string, bool) {
	if o == nil || isNil(o.DefaultPolicy) {
    return nil, false
	}
	return o.DefaultPolicy, true
}

// HasDefaultPolicy returns a boolean if a field has been set.
func (o *InlineResponse200168) HasDefaultPolicy() bool {
	if o != nil && !isNil(o.DefaultPolicy) {
		return true
	}

	return false
}

// SetDefaultPolicy gets a reference to the given string and assigns it to the DefaultPolicy field.
func (o *InlineResponse200168) SetDefaultPolicy(v string) {
	o.DefaultPolicy = &v
}

func (o InlineResponse200168) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.DefaultPolicy) {
		toSerialize["defaultPolicy"] = o.DefaultPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200168 struct {
	value *InlineResponse200168
	isSet bool
}

func (v NullableInlineResponse200168) Get() *InlineResponse200168 {
	return v.value
}

func (v *NullableInlineResponse200168) Set(val *InlineResponse200168) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200168) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200168) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200168(val *InlineResponse200168) *NullableInlineResponse200168 {
	return &NullableInlineResponse200168{value: val, isSet: true}
}

func (v NullableInlineResponse200168) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200168) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


