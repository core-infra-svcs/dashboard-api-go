/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject260 struct for InlineObject260
type InlineObject260 struct {
	// Name of a policy object, unique within the organization (alphanumeric, space, dash, or underscore characters only)
	Name *string `json:"name,omitempty"`
	// CIDR Value of a policy object (e.g. 10.11.12.1/24\")
	Cidr *string `json:"cidr,omitempty"`
	// Fully qualified domain name of policy object (e.g. \"example.com\")
	Fqdn *string `json:"fqdn,omitempty"`
	// Mask of a policy object (e.g. \"255.255.0.0\")
	Mask *string `json:"mask,omitempty"`
	// IP Address of a policy object (e.g. \"1.2.3.4\")
	Ip *string `json:"ip,omitempty"`
	// The IDs of policy object groups the policy object belongs to
	GroupIds []string `json:"groupIds,omitempty"`
}

// NewInlineObject260 instantiates a new InlineObject260 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject260() *InlineObject260 {
	this := InlineObject260{}
	return &this
}

// NewInlineObject260WithDefaults instantiates a new InlineObject260 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject260WithDefaults() *InlineObject260 {
	this := InlineObject260{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject260) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject260) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject260) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject260) SetName(v string) {
	o.Name = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *InlineObject260) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject260) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *InlineObject260) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *InlineObject260) SetCidr(v string) {
	o.Cidr = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *InlineObject260) GetFqdn() string {
	if o == nil || isNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject260) GetFqdnOk() (*string, bool) {
	if o == nil || isNil(o.Fqdn) {
    return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *InlineObject260) HasFqdn() bool {
	if o != nil && !isNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *InlineObject260) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *InlineObject260) GetMask() string {
	if o == nil || isNil(o.Mask) {
		var ret string
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject260) GetMaskOk() (*string, bool) {
	if o == nil || isNil(o.Mask) {
    return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *InlineObject260) HasMask() bool {
	if o != nil && !isNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given string and assigns it to the Mask field.
func (o *InlineObject260) SetMask(v string) {
	o.Mask = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineObject260) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject260) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineObject260) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineObject260) SetIp(v string) {
	o.Ip = &v
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *InlineObject260) GetGroupIds() []string {
	if o == nil || isNil(o.GroupIds) {
		var ret []string
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject260) GetGroupIdsOk() ([]string, bool) {
	if o == nil || isNil(o.GroupIds) {
    return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *InlineObject260) HasGroupIds() bool {
	if o != nil && !isNil(o.GroupIds) {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *InlineObject260) SetGroupIds(v []string) {
	o.GroupIds = v
}

func (o InlineObject260) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
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

type NullableInlineObject260 struct {
	value *InlineObject260
	isSet bool
}

func (v NullableInlineObject260) Get() *InlineObject260 {
	return v.value
}

func (v *NullableInlineObject260) Set(val *InlineObject260) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject260) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject260) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject260(val *InlineObject260) *NullableInlineObject260 {
	return &NullableInlineObject260{value: val, isSet: true}
}

func (v NullableInlineObject260) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject260) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


