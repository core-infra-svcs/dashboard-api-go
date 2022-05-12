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

// InlineObject95 struct for InlineObject95
type InlineObject95 struct {
	// 'allow' or 'block' new DHCP servers. Default value is 'allow'.
	DefaultPolicy *string `json:"defaultPolicy,omitempty"`
	// List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set to block. An empty array will clear the entries.
	AllowedServers *[]string `json:"allowedServers,omitempty"`
	// List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set to allow. An empty array will clear the entries.
	BlockedServers *[]string `json:"blockedServers,omitempty"`
}

// NewInlineObject95 instantiates a new InlineObject95 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject95() *InlineObject95 {
	this := InlineObject95{}
	return &this
}

// NewInlineObject95WithDefaults instantiates a new InlineObject95 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject95WithDefaults() *InlineObject95 {
	this := InlineObject95{}
	return &this
}

// GetDefaultPolicy returns the DefaultPolicy field value if set, zero value otherwise.
func (o *InlineObject95) GetDefaultPolicy() string {
	if o == nil || o.DefaultPolicy == nil {
		var ret string
		return ret
	}
	return *o.DefaultPolicy
}

// GetDefaultPolicyOk returns a tuple with the DefaultPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject95) GetDefaultPolicyOk() (*string, bool) {
	if o == nil || o.DefaultPolicy == nil {
		return nil, false
	}
	return o.DefaultPolicy, true
}

// HasDefaultPolicy returns a boolean if a field has been set.
func (o *InlineObject95) HasDefaultPolicy() bool {
	if o != nil && o.DefaultPolicy != nil {
		return true
	}

	return false
}

// SetDefaultPolicy gets a reference to the given string and assigns it to the DefaultPolicy field.
func (o *InlineObject95) SetDefaultPolicy(v string) {
	o.DefaultPolicy = &v
}

// GetAllowedServers returns the AllowedServers field value if set, zero value otherwise.
func (o *InlineObject95) GetAllowedServers() []string {
	if o == nil || o.AllowedServers == nil {
		var ret []string
		return ret
	}
	return *o.AllowedServers
}

// GetAllowedServersOk returns a tuple with the AllowedServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject95) GetAllowedServersOk() (*[]string, bool) {
	if o == nil || o.AllowedServers == nil {
		return nil, false
	}
	return o.AllowedServers, true
}

// HasAllowedServers returns a boolean if a field has been set.
func (o *InlineObject95) HasAllowedServers() bool {
	if o != nil && o.AllowedServers != nil {
		return true
	}

	return false
}

// SetAllowedServers gets a reference to the given []string and assigns it to the AllowedServers field.
func (o *InlineObject95) SetAllowedServers(v []string) {
	o.AllowedServers = &v
}

// GetBlockedServers returns the BlockedServers field value if set, zero value otherwise.
func (o *InlineObject95) GetBlockedServers() []string {
	if o == nil || o.BlockedServers == nil {
		var ret []string
		return ret
	}
	return *o.BlockedServers
}

// GetBlockedServersOk returns a tuple with the BlockedServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject95) GetBlockedServersOk() (*[]string, bool) {
	if o == nil || o.BlockedServers == nil {
		return nil, false
	}
	return o.BlockedServers, true
}

// HasBlockedServers returns a boolean if a field has been set.
func (o *InlineObject95) HasBlockedServers() bool {
	if o != nil && o.BlockedServers != nil {
		return true
	}

	return false
}

// SetBlockedServers gets a reference to the given []string and assigns it to the BlockedServers field.
func (o *InlineObject95) SetBlockedServers(v []string) {
	o.BlockedServers = &v
}

func (o InlineObject95) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultPolicy != nil {
		toSerialize["defaultPolicy"] = o.DefaultPolicy
	}
	if o.AllowedServers != nil {
		toSerialize["allowedServers"] = o.AllowedServers
	}
	if o.BlockedServers != nil {
		toSerialize["blockedServers"] = o.BlockedServers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject95 struct {
	value *InlineObject95
	isSet bool
}

func (v NullableInlineObject95) Get() *InlineObject95 {
	return v.value
}

func (v *NullableInlineObject95) Set(val *InlineObject95) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject95) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject95) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject95(val *InlineObject95) *NullableInlineObject95 {
	return &NullableInlineObject95{value: val, isSet: true}
}

func (v NullableInlineObject95) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject95) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

