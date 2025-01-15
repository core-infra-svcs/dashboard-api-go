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

// InlineResponse200288Items struct for InlineResponse200288Items
type InlineResponse200288Items struct {
	// The Id of the Network
	NetworkId *string `json:"networkId,omitempty"`
	// Array of Sentry Group Policies for the Network
	Policies []InlineResponse200288Policies `json:"policies,omitempty"`
}

// NewInlineResponse200288Items instantiates a new InlineResponse200288Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200288Items() *InlineResponse200288Items {
	this := InlineResponse200288Items{}
	return &this
}

// NewInlineResponse200288ItemsWithDefaults instantiates a new InlineResponse200288Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200288ItemsWithDefaults() *InlineResponse200288Items {
	this := InlineResponse200288Items{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200288Items) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200288Items) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200288Items) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200288Items) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *InlineResponse200288Items) GetPolicies() []InlineResponse200288Policies {
	if o == nil || isNil(o.Policies) {
		var ret []InlineResponse200288Policies
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200288Items) GetPoliciesOk() ([]InlineResponse200288Policies, bool) {
	if o == nil || isNil(o.Policies) {
    return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *InlineResponse200288Items) HasPolicies() bool {
	if o != nil && !isNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []InlineResponse200288Policies and assigns it to the Policies field.
func (o *InlineResponse200288Items) SetPolicies(v []InlineResponse200288Policies) {
	o.Policies = v
}

func (o InlineResponse200288Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200288Items struct {
	value *InlineResponse200288Items
	isSet bool
}

func (v NullableInlineResponse200288Items) Get() *InlineResponse200288Items {
	return v.value
}

func (v *NullableInlineResponse200288Items) Set(val *InlineResponse200288Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200288Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200288Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200288Items(val *InlineResponse200288Items) *NullableInlineResponse200288Items {
	return &NullableInlineResponse200288Items{value: val, isSet: true}
}

func (v NullableInlineResponse200288Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200288Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

