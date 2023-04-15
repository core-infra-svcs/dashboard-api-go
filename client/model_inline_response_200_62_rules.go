/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20062Rules struct for InlineResponse20062Rules
type InlineResponse20062Rules struct {
	// Description of the rule (optional)
	Comment *string `json:"comment,omitempty"`
	// 'allow' or 'deny' traffic specified by this rule
	Policy *string `json:"policy,omitempty"`
	// IP address version
	IpVersion *string `json:"ipVersion,omitempty"`
	// The type of protocol
	Protocol *string `json:"protocol,omitempty"`
	// Source IP address (in IP or CIDR notation)
	SrcCidr *string `json:"srcCidr,omitempty"`
	// Source port
	SrcPort *string `json:"srcPort,omitempty"`
	// Destination IP address (in IP or CIDR notation)
	DstCidr *string `json:"dstCidr,omitempty"`
	// Destination port
	DstPort *string `json:"dstPort,omitempty"`
	// ncoming traffic VLAN
	Vlan *string `json:"vlan,omitempty"`
}

// NewInlineResponse20062Rules instantiates a new InlineResponse20062Rules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20062Rules() *InlineResponse20062Rules {
	this := InlineResponse20062Rules{}
	return &this
}

// NewInlineResponse20062RulesWithDefaults instantiates a new InlineResponse20062Rules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20062RulesWithDefaults() *InlineResponse20062Rules {
	this := InlineResponse20062Rules{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *InlineResponse20062Rules) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062Rules) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *InlineResponse20062Rules) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *InlineResponse20062Rules) SetComment(v string) {
	o.Comment = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *InlineResponse20062Rules) GetPolicy() string {
	if o == nil || isNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062Rules) GetPolicyOk() (*string, bool) {
	if o == nil || isNil(o.Policy) {
    return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *InlineResponse20062Rules) HasPolicy() bool {
	if o != nil && !isNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *InlineResponse20062Rules) SetPolicy(v string) {
	o.Policy = &v
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise.
func (o *InlineResponse20062Rules) GetIpVersion() string {
	if o == nil || isNil(o.IpVersion) {
		var ret string
		return ret
	}
	return *o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062Rules) GetIpVersionOk() (*string, bool) {
	if o == nil || isNil(o.IpVersion) {
    return nil, false
	}
	return o.IpVersion, true
}

// HasIpVersion returns a boolean if a field has been set.
func (o *InlineResponse20062Rules) HasIpVersion() bool {
	if o != nil && !isNil(o.IpVersion) {
		return true
	}

	return false
}

// SetIpVersion gets a reference to the given string and assigns it to the IpVersion field.
func (o *InlineResponse20062Rules) SetIpVersion(v string) {
	o.IpVersion = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *InlineResponse20062Rules) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062Rules) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *InlineResponse20062Rules) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *InlineResponse20062Rules) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSrcCidr returns the SrcCidr field value if set, zero value otherwise.
func (o *InlineResponse20062Rules) GetSrcCidr() string {
	if o == nil || isNil(o.SrcCidr) {
		var ret string
		return ret
	}
	return *o.SrcCidr
}

// GetSrcCidrOk returns a tuple with the SrcCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062Rules) GetSrcCidrOk() (*string, bool) {
	if o == nil || isNil(o.SrcCidr) {
    return nil, false
	}
	return o.SrcCidr, true
}

// HasSrcCidr returns a boolean if a field has been set.
func (o *InlineResponse20062Rules) HasSrcCidr() bool {
	if o != nil && !isNil(o.SrcCidr) {
		return true
	}

	return false
}

// SetSrcCidr gets a reference to the given string and assigns it to the SrcCidr field.
func (o *InlineResponse20062Rules) SetSrcCidr(v string) {
	o.SrcCidr = &v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *InlineResponse20062Rules) GetSrcPort() string {
	if o == nil || isNil(o.SrcPort) {
		var ret string
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062Rules) GetSrcPortOk() (*string, bool) {
	if o == nil || isNil(o.SrcPort) {
    return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *InlineResponse20062Rules) HasSrcPort() bool {
	if o != nil && !isNil(o.SrcPort) {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given string and assigns it to the SrcPort field.
func (o *InlineResponse20062Rules) SetSrcPort(v string) {
	o.SrcPort = &v
}

// GetDstCidr returns the DstCidr field value if set, zero value otherwise.
func (o *InlineResponse20062Rules) GetDstCidr() string {
	if o == nil || isNil(o.DstCidr) {
		var ret string
		return ret
	}
	return *o.DstCidr
}

// GetDstCidrOk returns a tuple with the DstCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062Rules) GetDstCidrOk() (*string, bool) {
	if o == nil || isNil(o.DstCidr) {
    return nil, false
	}
	return o.DstCidr, true
}

// HasDstCidr returns a boolean if a field has been set.
func (o *InlineResponse20062Rules) HasDstCidr() bool {
	if o != nil && !isNil(o.DstCidr) {
		return true
	}

	return false
}

// SetDstCidr gets a reference to the given string and assigns it to the DstCidr field.
func (o *InlineResponse20062Rules) SetDstCidr(v string) {
	o.DstCidr = &v
}

// GetDstPort returns the DstPort field value if set, zero value otherwise.
func (o *InlineResponse20062Rules) GetDstPort() string {
	if o == nil || isNil(o.DstPort) {
		var ret string
		return ret
	}
	return *o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062Rules) GetDstPortOk() (*string, bool) {
	if o == nil || isNil(o.DstPort) {
    return nil, false
	}
	return o.DstPort, true
}

// HasDstPort returns a boolean if a field has been set.
func (o *InlineResponse20062Rules) HasDstPort() bool {
	if o != nil && !isNil(o.DstPort) {
		return true
	}

	return false
}

// SetDstPort gets a reference to the given string and assigns it to the DstPort field.
func (o *InlineResponse20062Rules) SetDstPort(v string) {
	o.DstPort = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineResponse20062Rules) GetVlan() string {
	if o == nil || isNil(o.Vlan) {
		var ret string
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062Rules) GetVlanOk() (*string, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineResponse20062Rules) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given string and assigns it to the Vlan field.
func (o *InlineResponse20062Rules) SetVlan(v string) {
	o.Vlan = &v
}

func (o InlineResponse20062Rules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !isNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !isNil(o.IpVersion) {
		toSerialize["ipVersion"] = o.IpVersion
	}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.SrcCidr) {
		toSerialize["srcCidr"] = o.SrcCidr
	}
	if !isNil(o.SrcPort) {
		toSerialize["srcPort"] = o.SrcPort
	}
	if !isNil(o.DstCidr) {
		toSerialize["dstCidr"] = o.DstCidr
	}
	if !isNil(o.DstPort) {
		toSerialize["dstPort"] = o.DstPort
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20062Rules struct {
	value *InlineResponse20062Rules
	isSet bool
}

func (v NullableInlineResponse20062Rules) Get() *InlineResponse20062Rules {
	return v.value
}

func (v *NullableInlineResponse20062Rules) Set(val *InlineResponse20062Rules) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20062Rules) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20062Rules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20062Rules(val *InlineResponse20062Rules) *NullableInlineResponse20062Rules {
	return &NullableInlineResponse20062Rules{value: val, isSet: true}
}

func (v NullableInlineResponse20062Rules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20062Rules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


