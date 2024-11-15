/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200196Rules struct for InlineResponse200196Rules
type InlineResponse200196Rules struct {
	// 'Deny' traffic specified by this rule
	Policy *string `json:"policy,omitempty"`
	// Type of the L7 firewall rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'
	Type *string `json:"type,omitempty"`
	// The value of what needs to get blocked. Format of the value varies depending on type of the firewall rule selected.
	Value *string `json:"value,omitempty"`
}

// NewInlineResponse200196Rules instantiates a new InlineResponse200196Rules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200196Rules() *InlineResponse200196Rules {
	this := InlineResponse200196Rules{}
	return &this
}

// NewInlineResponse200196RulesWithDefaults instantiates a new InlineResponse200196Rules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200196RulesWithDefaults() *InlineResponse200196Rules {
	this := InlineResponse200196Rules{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *InlineResponse200196Rules) GetPolicy() string {
	if o == nil || isNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200196Rules) GetPolicyOk() (*string, bool) {
	if o == nil || isNil(o.Policy) {
    return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *InlineResponse200196Rules) HasPolicy() bool {
	if o != nil && !isNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *InlineResponse200196Rules) SetPolicy(v string) {
	o.Policy = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200196Rules) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200196Rules) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200196Rules) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200196Rules) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InlineResponse200196Rules) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200196Rules) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineResponse200196Rules) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InlineResponse200196Rules) SetValue(v string) {
	o.Value = &v
}

func (o InlineResponse200196Rules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200196Rules struct {
	value *InlineResponse200196Rules
	isSet bool
}

func (v NullableInlineResponse200196Rules) Get() *InlineResponse200196Rules {
	return v.value
}

func (v *NullableInlineResponse200196Rules) Set(val *InlineResponse200196Rules) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200196Rules) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200196Rules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200196Rules(val *InlineResponse200196Rules) *NullableInlineResponse200196Rules {
	return &NullableInlineResponse200196Rules{value: val, isSet: true}
}

func (v NullableInlineResponse200196Rules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200196Rules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

