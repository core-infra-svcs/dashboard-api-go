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

// InlineObject35 struct for InlineObject35
type InlineObject35 struct {
	// The status of the port
	Enabled *bool `json:"enabled,omitempty"`
	// Trunk port can Drop all Untagged traffic. When true, no VLAN is required. Access ports cannot have dropUntaggedTraffic set to true.
	DropUntaggedTraffic *bool `json:"dropUntaggedTraffic,omitempty"`
	// The type of the port: 'access' or 'trunk'.
	Type *string `json:"type,omitempty"`
	// Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode.
	Vlan *int32 `json:"vlan,omitempty"`
	// Comma-delimited list of the VLAN ID's allowed on the port, or 'all' to permit all VLAN's on the port.
	AllowedVlans *string `json:"allowedVlans,omitempty"`
	// The name of the policy. Only applicable to Access ports. Valid values are: 'open', '8021x-radius', 'mac-radius', 'hybris-radius' for MX64 or Z3 or any MX supporting the per port authentication feature. Otherwise, 'open' is the only valid value and 'open' is the default value if the field is missing.
	AccessPolicy *string `json:"accessPolicy,omitempty"`
}

// NewInlineObject35 instantiates a new InlineObject35 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject35() *InlineObject35 {
	this := InlineObject35{}
	return &this
}

// NewInlineObject35WithDefaults instantiates a new InlineObject35 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject35WithDefaults() *InlineObject35 {
	this := InlineObject35{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject35) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject35) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject35) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject35) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDropUntaggedTraffic returns the DropUntaggedTraffic field value if set, zero value otherwise.
func (o *InlineObject35) GetDropUntaggedTraffic() bool {
	if o == nil || o.DropUntaggedTraffic == nil {
		var ret bool
		return ret
	}
	return *o.DropUntaggedTraffic
}

// GetDropUntaggedTrafficOk returns a tuple with the DropUntaggedTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject35) GetDropUntaggedTrafficOk() (*bool, bool) {
	if o == nil || o.DropUntaggedTraffic == nil {
		return nil, false
	}
	return o.DropUntaggedTraffic, true
}

// HasDropUntaggedTraffic returns a boolean if a field has been set.
func (o *InlineObject35) HasDropUntaggedTraffic() bool {
	if o != nil && o.DropUntaggedTraffic != nil {
		return true
	}

	return false
}

// SetDropUntaggedTraffic gets a reference to the given bool and assigns it to the DropUntaggedTraffic field.
func (o *InlineObject35) SetDropUntaggedTraffic(v bool) {
	o.DropUntaggedTraffic = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineObject35) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject35) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject35) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineObject35) SetType(v string) {
	o.Type = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineObject35) GetVlan() int32 {
	if o == nil || o.Vlan == nil {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject35) GetVlanOk() (*int32, bool) {
	if o == nil || o.Vlan == nil {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineObject35) HasVlan() bool {
	if o != nil && o.Vlan != nil {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *InlineObject35) SetVlan(v int32) {
	o.Vlan = &v
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *InlineObject35) GetAllowedVlans() string {
	if o == nil || o.AllowedVlans == nil {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject35) GetAllowedVlansOk() (*string, bool) {
	if o == nil || o.AllowedVlans == nil {
		return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *InlineObject35) HasAllowedVlans() bool {
	if o != nil && o.AllowedVlans != nil {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *InlineObject35) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

// GetAccessPolicy returns the AccessPolicy field value if set, zero value otherwise.
func (o *InlineObject35) GetAccessPolicy() string {
	if o == nil || o.AccessPolicy == nil {
		var ret string
		return ret
	}
	return *o.AccessPolicy
}

// GetAccessPolicyOk returns a tuple with the AccessPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject35) GetAccessPolicyOk() (*string, bool) {
	if o == nil || o.AccessPolicy == nil {
		return nil, false
	}
	return o.AccessPolicy, true
}

// HasAccessPolicy returns a boolean if a field has been set.
func (o *InlineObject35) HasAccessPolicy() bool {
	if o != nil && o.AccessPolicy != nil {
		return true
	}

	return false
}

// SetAccessPolicy gets a reference to the given string and assigns it to the AccessPolicy field.
func (o *InlineObject35) SetAccessPolicy(v string) {
	o.AccessPolicy = &v
}

func (o InlineObject35) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.DropUntaggedTraffic != nil {
		toSerialize["dropUntaggedTraffic"] = o.DropUntaggedTraffic
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Vlan != nil {
		toSerialize["vlan"] = o.Vlan
	}
	if o.AllowedVlans != nil {
		toSerialize["allowedVlans"] = o.AllowedVlans
	}
	if o.AccessPolicy != nil {
		toSerialize["accessPolicy"] = o.AccessPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject35 struct {
	value *InlineObject35
	isSet bool
}

func (v NullableInlineObject35) Get() *InlineObject35 {
	return v.value
}

func (v *NullableInlineObject35) Set(val *InlineObject35) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject35) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject35) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject35(val *InlineObject35) *NullableInlineObject35 {
	return &NullableInlineObject35{value: val, isSet: true}
}

func (v NullableInlineObject35) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject35) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


