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

// InlineResponse2017 struct for InlineResponse2017
type InlineResponse2017 struct {
	// The list of clients to provision
	Clients []InlineResponse2017Clients `json:"clients,omitempty"`
	// The name of the client's policy
	DevicePolicy *string `json:"devicePolicy,omitempty"`
	// The group policy identifier of the client
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
}

// NewInlineResponse2017 instantiates a new InlineResponse2017 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2017() *InlineResponse2017 {
	this := InlineResponse2017{}
	return &this
}

// NewInlineResponse2017WithDefaults instantiates a new InlineResponse2017 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2017WithDefaults() *InlineResponse2017 {
	this := InlineResponse2017{}
	return &this
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *InlineResponse2017) GetClients() []InlineResponse2017Clients {
	if o == nil || isNil(o.Clients) {
		var ret []InlineResponse2017Clients
		return ret
	}
	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017) GetClientsOk() ([]InlineResponse2017Clients, bool) {
	if o == nil || isNil(o.Clients) {
    return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *InlineResponse2017) HasClients() bool {
	if o != nil && !isNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given []InlineResponse2017Clients and assigns it to the Clients field.
func (o *InlineResponse2017) SetClients(v []InlineResponse2017Clients) {
	o.Clients = v
}

// GetDevicePolicy returns the DevicePolicy field value if set, zero value otherwise.
func (o *InlineResponse2017) GetDevicePolicy() string {
	if o == nil || isNil(o.DevicePolicy) {
		var ret string
		return ret
	}
	return *o.DevicePolicy
}

// GetDevicePolicyOk returns a tuple with the DevicePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017) GetDevicePolicyOk() (*string, bool) {
	if o == nil || isNil(o.DevicePolicy) {
    return nil, false
	}
	return o.DevicePolicy, true
}

// HasDevicePolicy returns a boolean if a field has been set.
func (o *InlineResponse2017) HasDevicePolicy() bool {
	if o != nil && !isNil(o.DevicePolicy) {
		return true
	}

	return false
}

// SetDevicePolicy gets a reference to the given string and assigns it to the DevicePolicy field.
func (o *InlineResponse2017) SetDevicePolicy(v string) {
	o.DevicePolicy = &v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *InlineResponse2017) GetGroupPolicyId() string {
	if o == nil || isNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupPolicyId) {
    return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *InlineResponse2017) HasGroupPolicyId() bool {
	if o != nil && !isNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *InlineResponse2017) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

func (o InlineResponse2017) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	if !isNil(o.DevicePolicy) {
		toSerialize["devicePolicy"] = o.DevicePolicy
	}
	if !isNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2017 struct {
	value *InlineResponse2017
	isSet bool
}

func (v NullableInlineResponse2017) Get() *InlineResponse2017 {
	return v.value
}

func (v *NullableInlineResponse2017) Set(val *InlineResponse2017) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2017) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2017) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2017(val *InlineResponse2017) *NullableInlineResponse2017 {
	return &NullableInlineResponse2017{value: val, isSet: true}
}

func (v NullableInlineResponse2017) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2017) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


