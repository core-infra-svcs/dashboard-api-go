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

// OrganizationsOrganizationIdAdaptivePolicyAclsRules struct for OrganizationsOrganizationIdAdaptivePolicyAclsRules
type OrganizationsOrganizationIdAdaptivePolicyAclsRules struct {
	// 'allow' or 'deny' traffic specified by this rule
	Policy *string `json:"policy,omitempty"`
	// The type of protocol
	Protocol *string `json:"protocol,omitempty"`
	// Source port
	SrcPort *string `json:"srcPort,omitempty"`
	// Destination port
	DstPort *string `json:"dstPort,omitempty"`
}

// NewOrganizationsOrganizationIdAdaptivePolicyAclsRules instantiates a new OrganizationsOrganizationIdAdaptivePolicyAclsRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdAdaptivePolicyAclsRules() *OrganizationsOrganizationIdAdaptivePolicyAclsRules {
	this := OrganizationsOrganizationIdAdaptivePolicyAclsRules{}
	return &this
}

// NewOrganizationsOrganizationIdAdaptivePolicyAclsRulesWithDefaults instantiates a new OrganizationsOrganizationIdAdaptivePolicyAclsRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdAdaptivePolicyAclsRulesWithDefaults() *OrganizationsOrganizationIdAdaptivePolicyAclsRules {
	this := OrganizationsOrganizationIdAdaptivePolicyAclsRules{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) GetPolicy() string {
	if o == nil || isNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) GetPolicyOk() (*string, bool) {
	if o == nil || isNil(o.Policy) {
    return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) HasPolicy() bool {
	if o != nil && !isNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) SetPolicy(v string) {
	o.Policy = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) GetSrcPort() string {
	if o == nil || isNil(o.SrcPort) {
		var ret string
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) GetSrcPortOk() (*string, bool) {
	if o == nil || isNil(o.SrcPort) {
    return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) HasSrcPort() bool {
	if o != nil && !isNil(o.SrcPort) {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given string and assigns it to the SrcPort field.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) SetSrcPort(v string) {
	o.SrcPort = &v
}

// GetDstPort returns the DstPort field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) GetDstPort() string {
	if o == nil || isNil(o.DstPort) {
		var ret string
		return ret
	}
	return *o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) GetDstPortOk() (*string, bool) {
	if o == nil || isNil(o.DstPort) {
    return nil, false
	}
	return o.DstPort, true
}

// HasDstPort returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) HasDstPort() bool {
	if o != nil && !isNil(o.DstPort) {
		return true
	}

	return false
}

// SetDstPort gets a reference to the given string and assigns it to the DstPort field.
func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) SetDstPort(v string) {
	o.DstPort = &v
}

func (o OrganizationsOrganizationIdAdaptivePolicyAclsRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !isNil(o.Protocol) {
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

type NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules struct {
	value *OrganizationsOrganizationIdAdaptivePolicyAclsRules
	isSet bool
}

func (v NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules) Get() *OrganizationsOrganizationIdAdaptivePolicyAclsRules {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules) Set(val *OrganizationsOrganizationIdAdaptivePolicyAclsRules) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdAdaptivePolicyAclsRules(val *OrganizationsOrganizationIdAdaptivePolicyAclsRules) *NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules {
	return &NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdAdaptivePolicyAclsRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


