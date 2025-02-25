/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200300Items struct for InlineResponse200300Items
type InlineResponse200300Items struct {
	// The Id of the Network
	NetworkId *string `json:"networkId,omitempty"`
	// Array of Sentry Group Policies for the Network
	Policies []InlineResponse200300Policies `json:"policies,omitempty"`
}

// NewInlineResponse200300Items instantiates a new InlineResponse200300Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200300Items() *InlineResponse200300Items {
	this := InlineResponse200300Items{}
	return &this
}

// NewInlineResponse200300ItemsWithDefaults instantiates a new InlineResponse200300Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200300ItemsWithDefaults() *InlineResponse200300Items {
	this := InlineResponse200300Items{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200300Items) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200300Items) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200300Items) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200300Items) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *InlineResponse200300Items) GetPolicies() []InlineResponse200300Policies {
	if o == nil || isNil(o.Policies) {
		var ret []InlineResponse200300Policies
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200300Items) GetPoliciesOk() ([]InlineResponse200300Policies, bool) {
	if o == nil || isNil(o.Policies) {
    return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *InlineResponse200300Items) HasPolicies() bool {
	if o != nil && !isNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []InlineResponse200300Policies and assigns it to the Policies field.
func (o *InlineResponse200300Items) SetPolicies(v []InlineResponse200300Policies) {
	o.Policies = v
}

func (o InlineResponse200300Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200300Items struct {
	value *InlineResponse200300Items
	isSet bool
}

func (v NullableInlineResponse200300Items) Get() *InlineResponse200300Items {
	return v.value
}

func (v *NullableInlineResponse200300Items) Set(val *InlineResponse200300Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200300Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200300Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200300Items(val *InlineResponse200300Items) *NullableInlineResponse200300Items {
	return &NullableInlineResponse200300Items{value: val, isSet: true}
}

func (v NullableInlineResponse200300Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200300Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


