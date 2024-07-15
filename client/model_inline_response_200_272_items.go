/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200272Items struct for InlineResponse200272Items
type InlineResponse200272Items struct {
	// The Id of the Network
	NetworkId *string `json:"networkId,omitempty"`
	// Array of Sentry Group Policies for the Network
	Policies []InlineResponse200272Policies `json:"policies,omitempty"`
}

// NewInlineResponse200272Items instantiates a new InlineResponse200272Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200272Items() *InlineResponse200272Items {
	this := InlineResponse200272Items{}
	return &this
}

// NewInlineResponse200272ItemsWithDefaults instantiates a new InlineResponse200272Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200272ItemsWithDefaults() *InlineResponse200272Items {
	this := InlineResponse200272Items{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200272Items) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200272Items) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200272Items) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200272Items) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *InlineResponse200272Items) GetPolicies() []InlineResponse200272Policies {
	if o == nil || isNil(o.Policies) {
		var ret []InlineResponse200272Policies
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200272Items) GetPoliciesOk() ([]InlineResponse200272Policies, bool) {
	if o == nil || isNil(o.Policies) {
    return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *InlineResponse200272Items) HasPolicies() bool {
	if o != nil && !isNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []InlineResponse200272Policies and assigns it to the Policies field.
func (o *InlineResponse200272Items) SetPolicies(v []InlineResponse200272Policies) {
	o.Policies = v
}

func (o InlineResponse200272Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200272Items struct {
	value *InlineResponse200272Items
	isSet bool
}

func (v NullableInlineResponse200272Items) Get() *InlineResponse200272Items {
	return v.value
}

func (v *NullableInlineResponse200272Items) Set(val *InlineResponse200272Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200272Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200272Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200272Items(val *InlineResponse200272Items) *NullableInlineResponse200272Items {
	return &NullableInlineResponse200272Items{value: val, isSet: true}
}

func (v NullableInlineResponse200272Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200272Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

