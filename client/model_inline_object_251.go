/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject251 struct for InlineObject251
type InlineObject251 struct {
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

// NewInlineObject251 instantiates a new InlineObject251 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject251() *InlineObject251 {
	this := InlineObject251{}
	return &this
}

// NewInlineObject251WithDefaults instantiates a new InlineObject251 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject251WithDefaults() *InlineObject251 {
	this := InlineObject251{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject251) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject251) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject251) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject251) SetName(v string) {
	o.Name = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *InlineObject251) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject251) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *InlineObject251) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *InlineObject251) SetCidr(v string) {
	o.Cidr = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *InlineObject251) GetFqdn() string {
	if o == nil || isNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject251) GetFqdnOk() (*string, bool) {
	if o == nil || isNil(o.Fqdn) {
    return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *InlineObject251) HasFqdn() bool {
	if o != nil && !isNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *InlineObject251) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *InlineObject251) GetMask() string {
	if o == nil || isNil(o.Mask) {
		var ret string
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject251) GetMaskOk() (*string, bool) {
	if o == nil || isNil(o.Mask) {
    return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *InlineObject251) HasMask() bool {
	if o != nil && !isNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given string and assigns it to the Mask field.
func (o *InlineObject251) SetMask(v string) {
	o.Mask = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineObject251) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject251) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineObject251) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineObject251) SetIp(v string) {
	o.Ip = &v
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *InlineObject251) GetGroupIds() []string {
	if o == nil || isNil(o.GroupIds) {
		var ret []string
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject251) GetGroupIdsOk() ([]string, bool) {
	if o == nil || isNil(o.GroupIds) {
    return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *InlineObject251) HasGroupIds() bool {
	if o != nil && !isNil(o.GroupIds) {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *InlineObject251) SetGroupIds(v []string) {
	o.GroupIds = v
}

func (o InlineObject251) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject251 struct {
	value *InlineObject251
	isSet bool
}

func (v NullableInlineObject251) Get() *InlineObject251 {
	return v.value
}

func (v *NullableInlineObject251) Set(val *InlineObject251) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject251) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject251) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject251(val *InlineObject251) *NullableInlineObject251 {
	return &NullableInlineObject251{value: val, isSet: true}
}

func (v NullableInlineObject251) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject251) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


