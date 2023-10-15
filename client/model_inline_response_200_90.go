/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20090 struct for InlineResponse20090
type InlineResponse20090 struct {
	// IName of the VLAN profile
	Iname *string `json:"iname,omitempty"`
	// Name of the profile, string length must be from 1 to 255 characters
	Name *string `json:"name,omitempty"`
	// Boolean indicating the default VLAN Profile for any device that does not have a profile explicitly assigned
	IsDefault *bool `json:"isDefault,omitempty"`
	// An array of named VLANs
	VlanNames []NetworksNetworkIdVlanProfilesVlanNames `json:"vlanNames,omitempty"`
	// An array of named VLANs
	VlanGroups []NetworksNetworkIdVlanProfilesVlanGroups `json:"vlanGroups,omitempty"`
}

// NewInlineResponse20090 instantiates a new InlineResponse20090 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20090() *InlineResponse20090 {
	this := InlineResponse20090{}
	return &this
}

// NewInlineResponse20090WithDefaults instantiates a new InlineResponse20090 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20090WithDefaults() *InlineResponse20090 {
	this := InlineResponse20090{}
	return &this
}

// GetIname returns the Iname field value if set, zero value otherwise.
func (o *InlineResponse20090) GetIname() string {
	if o == nil || isNil(o.Iname) {
		var ret string
		return ret
	}
	return *o.Iname
}

// GetInameOk returns a tuple with the Iname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090) GetInameOk() (*string, bool) {
	if o == nil || isNil(o.Iname) {
    return nil, false
	}
	return o.Iname, true
}

// HasIname returns a boolean if a field has been set.
func (o *InlineResponse20090) HasIname() bool {
	if o != nil && !isNil(o.Iname) {
		return true
	}

	return false
}

// SetIname gets a reference to the given string and assigns it to the Iname field.
func (o *InlineResponse20090) SetIname(v string) {
	o.Iname = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20090) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20090) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20090) SetName(v string) {
	o.Name = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *InlineResponse20090) GetIsDefault() bool {
	if o == nil || isNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090) GetIsDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.IsDefault) {
    return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *InlineResponse20090) HasIsDefault() bool {
	if o != nil && !isNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *InlineResponse20090) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetVlanNames returns the VlanNames field value if set, zero value otherwise.
func (o *InlineResponse20090) GetVlanNames() []NetworksNetworkIdVlanProfilesVlanNames {
	if o == nil || isNil(o.VlanNames) {
		var ret []NetworksNetworkIdVlanProfilesVlanNames
		return ret
	}
	return o.VlanNames
}

// GetVlanNamesOk returns a tuple with the VlanNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090) GetVlanNamesOk() ([]NetworksNetworkIdVlanProfilesVlanNames, bool) {
	if o == nil || isNil(o.VlanNames) {
    return nil, false
	}
	return o.VlanNames, true
}

// HasVlanNames returns a boolean if a field has been set.
func (o *InlineResponse20090) HasVlanNames() bool {
	if o != nil && !isNil(o.VlanNames) {
		return true
	}

	return false
}

// SetVlanNames gets a reference to the given []NetworksNetworkIdVlanProfilesVlanNames and assigns it to the VlanNames field.
func (o *InlineResponse20090) SetVlanNames(v []NetworksNetworkIdVlanProfilesVlanNames) {
	o.VlanNames = v
}

// GetVlanGroups returns the VlanGroups field value if set, zero value otherwise.
func (o *InlineResponse20090) GetVlanGroups() []NetworksNetworkIdVlanProfilesVlanGroups {
	if o == nil || isNil(o.VlanGroups) {
		var ret []NetworksNetworkIdVlanProfilesVlanGroups
		return ret
	}
	return o.VlanGroups
}

// GetVlanGroupsOk returns a tuple with the VlanGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090) GetVlanGroupsOk() ([]NetworksNetworkIdVlanProfilesVlanGroups, bool) {
	if o == nil || isNil(o.VlanGroups) {
    return nil, false
	}
	return o.VlanGroups, true
}

// HasVlanGroups returns a boolean if a field has been set.
func (o *InlineResponse20090) HasVlanGroups() bool {
	if o != nil && !isNil(o.VlanGroups) {
		return true
	}

	return false
}

// SetVlanGroups gets a reference to the given []NetworksNetworkIdVlanProfilesVlanGroups and assigns it to the VlanGroups field.
func (o *InlineResponse20090) SetVlanGroups(v []NetworksNetworkIdVlanProfilesVlanGroups) {
	o.VlanGroups = v
}

func (o InlineResponse20090) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Iname) {
		toSerialize["iname"] = o.Iname
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !isNil(o.VlanNames) {
		toSerialize["vlanNames"] = o.VlanNames
	}
	if !isNil(o.VlanGroups) {
		toSerialize["vlanGroups"] = o.VlanGroups
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20090 struct {
	value *InlineResponse20090
	isSet bool
}

func (v NullableInlineResponse20090) Get() *InlineResponse20090 {
	return v.value
}

func (v *NullableInlineResponse20090) Set(val *InlineResponse20090) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20090) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20090) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20090(val *InlineResponse20090) *NullableInlineResponse20090 {
	return &NullableInlineResponse20090{value: val, isSet: true}
}

func (v NullableInlineResponse20090) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20090) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


