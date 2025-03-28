/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject224 struct for InlineObject224
type InlineObject224 struct {
	// Name of profile
	Name string `json:"name"`
	// The hostname patterns to match for redirection. For more information on Split DNS hostname pattern formatting, please consult the Split DNS KB.
	Hostnames []string `json:"hostnames"`
	Nameservers OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers `json:"nameservers"`
}

// NewInlineObject224 instantiates a new InlineObject224 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject224(name string, hostnames []string, nameservers OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers) *InlineObject224 {
	this := InlineObject224{}
	this.Name = name
	this.Hostnames = hostnames
	this.Nameservers = nameservers
	return &this
}

// NewInlineObject224WithDefaults instantiates a new InlineObject224 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject224WithDefaults() *InlineObject224 {
	this := InlineObject224{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject224) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject224) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject224) SetName(v string) {
	o.Name = v
}

// GetHostnames returns the Hostnames field value
func (o *InlineObject224) GetHostnames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Hostnames
}

// GetHostnamesOk returns a tuple with the Hostnames field value
// and a boolean to check if the value has been set.
func (o *InlineObject224) GetHostnamesOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Hostnames, true
}

// SetHostnames sets field value
func (o *InlineObject224) SetHostnames(v []string) {
	o.Hostnames = v
}

// GetNameservers returns the Nameservers field value
func (o *InlineObject224) GetNameservers() OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers {
	if o == nil {
		var ret OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers
		return ret
	}

	return o.Nameservers
}

// GetNameserversOk returns a tuple with the Nameservers field value
// and a boolean to check if the value has been set.
func (o *InlineObject224) GetNameserversOk() (*OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Nameservers, true
}

// SetNameservers sets field value
func (o *InlineObject224) SetNameservers(v OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers) {
	o.Nameservers = v
}

func (o InlineObject224) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["hostnames"] = o.Hostnames
	}
	if true {
		toSerialize["nameservers"] = o.Nameservers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject224 struct {
	value *InlineObject224
	isSet bool
}

func (v NullableInlineObject224) Get() *InlineObject224 {
	return v.value
}

func (v *NullableInlineObject224) Set(val *InlineObject224) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject224) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject224) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject224(val *InlineObject224) *NullableInlineObject224 {
	return &NullableInlineObject224{value: val, isSet: true}
}

func (v NullableInlineObject224) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject224) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


