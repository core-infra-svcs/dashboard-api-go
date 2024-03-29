/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject49 struct for InlineObject49
type InlineObject49 struct {
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

// NewInlineObject49 instantiates a new InlineObject49 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject49() *InlineObject49 {
	this := InlineObject49{}
	return &this
}

// NewInlineObject49WithDefaults instantiates a new InlineObject49 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject49WithDefaults() *InlineObject49 {
	this := InlineObject49{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject49) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject49) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject49) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject49) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDropUntaggedTraffic returns the DropUntaggedTraffic field value if set, zero value otherwise.
func (o *InlineObject49) GetDropUntaggedTraffic() bool {
	if o == nil || isNil(o.DropUntaggedTraffic) {
		var ret bool
		return ret
	}
	return *o.DropUntaggedTraffic
}

// GetDropUntaggedTrafficOk returns a tuple with the DropUntaggedTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject49) GetDropUntaggedTrafficOk() (*bool, bool) {
	if o == nil || isNil(o.DropUntaggedTraffic) {
    return nil, false
	}
	return o.DropUntaggedTraffic, true
}

// HasDropUntaggedTraffic returns a boolean if a field has been set.
func (o *InlineObject49) HasDropUntaggedTraffic() bool {
	if o != nil && !isNil(o.DropUntaggedTraffic) {
		return true
	}

	return false
}

// SetDropUntaggedTraffic gets a reference to the given bool and assigns it to the DropUntaggedTraffic field.
func (o *InlineObject49) SetDropUntaggedTraffic(v bool) {
	o.DropUntaggedTraffic = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineObject49) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject49) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject49) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineObject49) SetType(v string) {
	o.Type = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineObject49) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject49) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineObject49) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *InlineObject49) SetVlan(v int32) {
	o.Vlan = &v
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *InlineObject49) GetAllowedVlans() string {
	if o == nil || isNil(o.AllowedVlans) {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject49) GetAllowedVlansOk() (*string, bool) {
	if o == nil || isNil(o.AllowedVlans) {
    return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *InlineObject49) HasAllowedVlans() bool {
	if o != nil && !isNil(o.AllowedVlans) {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *InlineObject49) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

// GetAccessPolicy returns the AccessPolicy field value if set, zero value otherwise.
func (o *InlineObject49) GetAccessPolicy() string {
	if o == nil || isNil(o.AccessPolicy) {
		var ret string
		return ret
	}
	return *o.AccessPolicy
}

// GetAccessPolicyOk returns a tuple with the AccessPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject49) GetAccessPolicyOk() (*string, bool) {
	if o == nil || isNil(o.AccessPolicy) {
    return nil, false
	}
	return o.AccessPolicy, true
}

// HasAccessPolicy returns a boolean if a field has been set.
func (o *InlineObject49) HasAccessPolicy() bool {
	if o != nil && !isNil(o.AccessPolicy) {
		return true
	}

	return false
}

// SetAccessPolicy gets a reference to the given string and assigns it to the AccessPolicy field.
func (o *InlineObject49) SetAccessPolicy(v string) {
	o.AccessPolicy = &v
}

func (o InlineObject49) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.DropUntaggedTraffic) {
		toSerialize["dropUntaggedTraffic"] = o.DropUntaggedTraffic
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.AllowedVlans) {
		toSerialize["allowedVlans"] = o.AllowedVlans
	}
	if !isNil(o.AccessPolicy) {
		toSerialize["accessPolicy"] = o.AccessPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject49 struct {
	value *InlineObject49
	isSet bool
}

func (v NullableInlineObject49) Get() *InlineObject49 {
	return v.value
}

func (v *NullableInlineObject49) Set(val *InlineObject49) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject49) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject49) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject49(val *InlineObject49) *NullableInlineObject49 {
	return &NullableInlineObject49{value: val, isSet: true}
}

func (v NullableInlineObject49) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject49) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


