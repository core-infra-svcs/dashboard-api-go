/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject80 struct for InlineObject80
type InlineObject80 struct {
	// The policy to assign. Can be 'Whitelisted', 'Blocked', 'Normal' or 'Group policy'. Required.
	DevicePolicy string `json:"devicePolicy"`
	// [optional] If 'devicePolicy' is set to 'Group policy' this param is used to specify the group policy ID.
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
}

// NewInlineObject80 instantiates a new InlineObject80 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject80(devicePolicy string) *InlineObject80 {
	this := InlineObject80{}
	this.DevicePolicy = devicePolicy
	return &this
}

// NewInlineObject80WithDefaults instantiates a new InlineObject80 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject80WithDefaults() *InlineObject80 {
	this := InlineObject80{}
	return &this
}

// GetDevicePolicy returns the DevicePolicy field value
func (o *InlineObject80) GetDevicePolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevicePolicy
}

// GetDevicePolicyOk returns a tuple with the DevicePolicy field value
// and a boolean to check if the value has been set.
func (o *InlineObject80) GetDevicePolicyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DevicePolicy, true
}

// SetDevicePolicy sets field value
func (o *InlineObject80) SetDevicePolicy(v string) {
	o.DevicePolicy = v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *InlineObject80) GetGroupPolicyId() string {
	if o == nil || isNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject80) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupPolicyId) {
    return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *InlineObject80) HasGroupPolicyId() bool {
	if o != nil && !isNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *InlineObject80) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

func (o InlineObject80) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["devicePolicy"] = o.DevicePolicy
	}
	if !isNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject80 struct {
	value *InlineObject80
	isSet bool
}

func (v NullableInlineObject80) Get() *InlineObject80 {
	return v.value
}

func (v *NullableInlineObject80) Set(val *InlineObject80) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject80) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject80) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject80(val *InlineObject80) *NullableInlineObject80 {
	return &NullableInlineObject80{value: val, isSet: true}
}

func (v NullableInlineObject80) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject80) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


