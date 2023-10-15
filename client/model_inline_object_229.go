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

// InlineObject229 struct for InlineObject229
type InlineObject229 struct {
	// Name of a policy object, unique within the organization (alphanumeric, space, dash, or underscore characters only)
	Name string `json:"name"`
	// Category of a policy object (one of: adaptivePolicy, network)
	Category string `json:"category"`
	// Type of a policy object (one of: adaptivePolicyIpv4Cidr, cidr, fqdn, ipAndMask)
	Type string `json:"type"`
	// CIDR Value of a policy object (e.g. 10.11.12.1/24\")
	Cidr *string `json:"cidr,omitempty"`
	// Fully qualified domain name of policy object (e.g. \"example.com\")
	Fqdn *string `json:"fqdn,omitempty"`
	// Mask of a policy object (e.g. \"255.255.0.0\")
	Mask *string `json:"mask,omitempty"`
	// IP Address of a policy object (e.g. \"1.2.3.4\")
	Ip *string `json:"ip,omitempty"`
	// The IDs of policy object groups the policy object belongs to
	GroupIds []int32 `json:"groupIds,omitempty"`
}

// NewInlineObject229 instantiates a new InlineObject229 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject229(name string, category string, type_ string) *InlineObject229 {
	this := InlineObject229{}
	this.Name = name
	this.Category = category
	this.Type = type_
	return &this
}

// NewInlineObject229WithDefaults instantiates a new InlineObject229 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject229WithDefaults() *InlineObject229 {
	this := InlineObject229{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject229) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject229) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject229) SetName(v string) {
	o.Name = v
}

// GetCategory returns the Category field value
func (o *InlineObject229) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *InlineObject229) GetCategoryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *InlineObject229) SetCategory(v string) {
	o.Category = v
}

// GetType returns the Type field value
func (o *InlineObject229) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineObject229) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineObject229) SetType(v string) {
	o.Type = v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *InlineObject229) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject229) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *InlineObject229) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *InlineObject229) SetCidr(v string) {
	o.Cidr = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *InlineObject229) GetFqdn() string {
	if o == nil || isNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject229) GetFqdnOk() (*string, bool) {
	if o == nil || isNil(o.Fqdn) {
    return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *InlineObject229) HasFqdn() bool {
	if o != nil && !isNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *InlineObject229) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *InlineObject229) GetMask() string {
	if o == nil || isNil(o.Mask) {
		var ret string
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject229) GetMaskOk() (*string, bool) {
	if o == nil || isNil(o.Mask) {
    return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *InlineObject229) HasMask() bool {
	if o != nil && !isNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given string and assigns it to the Mask field.
func (o *InlineObject229) SetMask(v string) {
	o.Mask = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineObject229) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject229) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineObject229) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineObject229) SetIp(v string) {
	o.Ip = &v
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *InlineObject229) GetGroupIds() []int32 {
	if o == nil || isNil(o.GroupIds) {
		var ret []int32
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject229) GetGroupIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.GroupIds) {
    return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *InlineObject229) HasGroupIds() bool {
	if o != nil && !isNil(o.GroupIds) {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []int32 and assigns it to the GroupIds field.
func (o *InlineObject229) SetGroupIds(v []int32) {
	o.GroupIds = v
}

func (o InlineObject229) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !isNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !isNil(o.Mask) {
		toSerialize["mask"] = o.Mask
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.GroupIds) {
		toSerialize["groupIds"] = o.GroupIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject229 struct {
	value *InlineObject229
	isSet bool
}

func (v NullableInlineObject229) Get() *InlineObject229 {
	return v.value
}

func (v *NullableInlineObject229) Set(val *InlineObject229) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject229) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject229) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject229(val *InlineObject229) *NullableInlineObject229 {
	return &NullableInlineObject229{value: val, isSet: true}
}

func (v NullableInlineObject229) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject229) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


