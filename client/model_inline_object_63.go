/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject63 struct for InlineObject63
type InlineObject63 struct {
	// The policy to assign. Can be 'Whitelisted', 'Blocked', 'Normal' or 'Group policy'. Required.
	DevicePolicy string `json:"devicePolicy"`
	// [optional] If 'devicePolicy' is set to 'Group policy' this param is used to specify the group policy ID.
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
}

// NewInlineObject63 instantiates a new InlineObject63 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject63(devicePolicy string) *InlineObject63 {
	this := InlineObject63{}
	this.DevicePolicy = devicePolicy
	return &this
}

// NewInlineObject63WithDefaults instantiates a new InlineObject63 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject63WithDefaults() *InlineObject63 {
	this := InlineObject63{}
	return &this
}

// GetDevicePolicy returns the DevicePolicy field value
func (o *InlineObject63) GetDevicePolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevicePolicy
}

// GetDevicePolicyOk returns a tuple with the DevicePolicy field value
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetDevicePolicyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DevicePolicy, true
}

// SetDevicePolicy sets field value
func (o *InlineObject63) SetDevicePolicy(v string) {
	o.DevicePolicy = v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *InlineObject63) GetGroupPolicyId() string {
	if o == nil || o.GroupPolicyId == nil {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || o.GroupPolicyId == nil {
		return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *InlineObject63) HasGroupPolicyId() bool {
	if o != nil && o.GroupPolicyId != nil {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *InlineObject63) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

func (o InlineObject63) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["devicePolicy"] = o.DevicePolicy
	}
	if o.GroupPolicyId != nil {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject63 struct {
	value *InlineObject63
	isSet bool
}

func (v NullableInlineObject63) Get() *InlineObject63 {
	return v.value
}

func (v *NullableInlineObject63) Set(val *InlineObject63) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject63) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject63) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject63(val *InlineObject63) *NullableInlineObject63 {
	return &NullableInlineObject63{value: val, isSet: true}
}

func (v NullableInlineObject63) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject63) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


