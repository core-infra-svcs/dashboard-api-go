/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdAdaptivePolicyAclsRules1 struct for OrganizationsOrganizationIdAdaptivePolicyAclsRules1
type OrganizationsOrganizationIdAdaptivePolicyAclsRules1 struct {
	// 'allow' or 'deny' traffic specified by this rule.
	Policy string `json:"policy"`
	// The type of protocol (must be 'tcp', 'udp', 'icmp' or 'any').
	Protocol string `json:"protocol"`
	// Source port. Must be in the format of single port: '1', port list: '1,2' or port range: '1-10', and in the range of 1-65535, or 'any'. Default is 'any'.
	SrcPort *string `json:"srcPort,omitempty"`
	// Destination port. Must be in the format of single port: '1', port list: '1,2' or port range: '1-10', and in the range of 1-65535, or 'any'. Default is 'any'.
	DstPort *string `json:"dstPort,omitempty"`
}

// NewOrganizationsOrganizationIdAdaptivePolicyAclsRules1 instantiates a new OrganizationsOrganizationIdAdaptivePolicyAclsRules1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdAdaptivePolicyAclsRules1(policy string, protocol string) *OrganizationsOrganizationIdAdaptivePolicyAclsRules1 {
	this := OrganizationsOrganizationIdAdaptivePolicyAclsRules1{}
	this.Policy = policy
	this.Protocol = protocol
	return &this
}

// NewOrganizationsOrganizationIdAdaptivePolicyAclsRules1WithDefaults instantiates a new OrganizationsOrganizationIdAdaptivePolicyAclsRules1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdAdaptivePolicyAclsRules1WithDefaults() *OrganizationsOrganizationIdAdaptivePolicyAclsRules1 {
	this := OrganizationsOrganizationIdAdaptivePolicyAclsRules1{}
	return &this
}

// GetPolicy returns the Policy field value
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetPolicyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) SetPolicy(v string) {
	o.Policy = v
}

// GetProtocol returns the Protocol field value
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetProtocolOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) SetProtocol(v string) {
	o.Protocol = v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetSrcPort() string {
	if o == nil || isNil(o.SrcPort) {
		var ret string
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetSrcPortOk() (*string, bool) {
	if o == nil || isNil(o.SrcPort) {
    return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) HasSrcPort() bool {
	if o != nil && !isNil(o.SrcPort) {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given string and assigns it to the SrcPort field.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) SetSrcPort(v string) {
	o.SrcPort = &v
}

// GetDstPort returns the DstPort field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetDstPort() string {
	if o == nil || isNil(o.DstPort) {
		var ret string
		return ret
	}
	return *o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetDstPortOk() (*string, bool) {
	if o == nil || isNil(o.DstPort) {
    return nil, false
	}
	return o.DstPort, true
}

// HasDstPort returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) HasDstPort() bool {
	if o != nil && !isNil(o.DstPort) {
		return true
	}

	return false
}

// SetDstPort gets a reference to the given string and assigns it to the DstPort field.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) SetDstPort(v string) {
	o.DstPort = &v
}

func (o OrganizationsOrganizationIdAdaptivePolicyAclsRules1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["policy"] = o.Policy
	}
	if true {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.SrcPort) {
		toSerialize["srcPort"] = o.SrcPort
	}
	if !isNil(o.DstPort) {
		toSerialize["dstPort"] = o.DstPort
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules1 struct {
	value *OrganizationsOrganizationIdAdaptivePolicyAclsRules1
	isSet bool
}

func (v NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules1) Get() *OrganizationsOrganizationIdAdaptivePolicyAclsRules1 {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules1) Set(val *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules1) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdAdaptivePolicyAclsRules1(val *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) *NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules1 {
	return &NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules1{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


