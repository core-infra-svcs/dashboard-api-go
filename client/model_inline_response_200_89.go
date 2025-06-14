/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20089 struct for InlineResponse20089
type InlineResponse20089 struct {
	// The MAC address of the client
	Mac *string `json:"mac,omitempty"`
	// The name of the client's policy
	DevicePolicy *string `json:"devicePolicy,omitempty"`
	// The group policy identifier of the client
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
}

// NewInlineResponse20089 instantiates a new InlineResponse20089 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20089() *InlineResponse20089 {
	this := InlineResponse20089{}
	return &this
}

// NewInlineResponse20089WithDefaults instantiates a new InlineResponse20089 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20089WithDefaults() *InlineResponse20089 {
	this := InlineResponse20089{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse20089) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse20089) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse20089) SetMac(v string) {
	o.Mac = &v
}

// GetDevicePolicy returns the DevicePolicy field value if set, zero value otherwise.
func (o *InlineResponse20089) GetDevicePolicy() string {
	if o == nil || isNil(o.DevicePolicy) {
		var ret string
		return ret
	}
	return *o.DevicePolicy
}

// GetDevicePolicyOk returns a tuple with the DevicePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089) GetDevicePolicyOk() (*string, bool) {
	if o == nil || isNil(o.DevicePolicy) {
    return nil, false
	}
	return o.DevicePolicy, true
}

// HasDevicePolicy returns a boolean if a field has been set.
func (o *InlineResponse20089) HasDevicePolicy() bool {
	if o != nil && !isNil(o.DevicePolicy) {
		return true
	}

	return false
}

// SetDevicePolicy gets a reference to the given string and assigns it to the DevicePolicy field.
func (o *InlineResponse20089) SetDevicePolicy(v string) {
	o.DevicePolicy = &v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *InlineResponse20089) GetGroupPolicyId() string {
	if o == nil || isNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupPolicyId) {
    return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *InlineResponse20089) HasGroupPolicyId() bool {
	if o != nil && !isNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *InlineResponse20089) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

func (o InlineResponse20089) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.DevicePolicy) {
		toSerialize["devicePolicy"] = o.DevicePolicy
	}
	if !isNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20089 struct {
	value *InlineResponse20089
	isSet bool
}

func (v NullableInlineResponse20089) Get() *InlineResponse20089 {
	return v.value
}

func (v *NullableInlineResponse20089) Set(val *InlineResponse20089) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20089) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20089) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20089(val *InlineResponse20089) *NullableInlineResponse20089 {
	return &NullableInlineResponse20089{value: val, isSet: true}
}

func (v NullableInlineResponse20089) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20089) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


