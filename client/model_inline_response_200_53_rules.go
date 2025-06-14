/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20053Rules struct for InlineResponse20053Rules
type InlineResponse20053Rules struct {
	// Description of the rule (optional)
	Comment *string `json:"comment,omitempty"`
	// 'allow' or 'deny' traffic specified by this rule
	Policy *string `json:"policy,omitempty"`
	// The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	Protocol *string `json:"protocol,omitempty"`
	// Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SrcPort *string `json:"srcPort,omitempty"`
	// Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)
	SrcCidr *string `json:"srcCidr,omitempty"`
	// Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	DestPort *string `json:"destPort,omitempty"`
	// Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestCidr *string `json:"destCidr,omitempty"`
	// Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
	SyslogEnabled *bool `json:"syslogEnabled,omitempty"`
}

// NewInlineResponse20053Rules instantiates a new InlineResponse20053Rules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20053Rules() *InlineResponse20053Rules {
	this := InlineResponse20053Rules{}
	return &this
}

// NewInlineResponse20053RulesWithDefaults instantiates a new InlineResponse20053Rules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20053RulesWithDefaults() *InlineResponse20053Rules {
	this := InlineResponse20053Rules{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *InlineResponse20053Rules) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Rules) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *InlineResponse20053Rules) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *InlineResponse20053Rules) SetComment(v string) {
	o.Comment = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *InlineResponse20053Rules) GetPolicy() string {
	if o == nil || isNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Rules) GetPolicyOk() (*string, bool) {
	if o == nil || isNil(o.Policy) {
    return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *InlineResponse20053Rules) HasPolicy() bool {
	if o != nil && !isNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *InlineResponse20053Rules) SetPolicy(v string) {
	o.Policy = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *InlineResponse20053Rules) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Rules) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *InlineResponse20053Rules) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *InlineResponse20053Rules) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *InlineResponse20053Rules) GetSrcPort() string {
	if o == nil || isNil(o.SrcPort) {
		var ret string
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Rules) GetSrcPortOk() (*string, bool) {
	if o == nil || isNil(o.SrcPort) {
    return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *InlineResponse20053Rules) HasSrcPort() bool {
	if o != nil && !isNil(o.SrcPort) {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given string and assigns it to the SrcPort field.
func (o *InlineResponse20053Rules) SetSrcPort(v string) {
	o.SrcPort = &v
}

// GetSrcCidr returns the SrcCidr field value if set, zero value otherwise.
func (o *InlineResponse20053Rules) GetSrcCidr() string {
	if o == nil || isNil(o.SrcCidr) {
		var ret string
		return ret
	}
	return *o.SrcCidr
}

// GetSrcCidrOk returns a tuple with the SrcCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Rules) GetSrcCidrOk() (*string, bool) {
	if o == nil || isNil(o.SrcCidr) {
    return nil, false
	}
	return o.SrcCidr, true
}

// HasSrcCidr returns a boolean if a field has been set.
func (o *InlineResponse20053Rules) HasSrcCidr() bool {
	if o != nil && !isNil(o.SrcCidr) {
		return true
	}

	return false
}

// SetSrcCidr gets a reference to the given string and assigns it to the SrcCidr field.
func (o *InlineResponse20053Rules) SetSrcCidr(v string) {
	o.SrcCidr = &v
}

// GetDestPort returns the DestPort field value if set, zero value otherwise.
func (o *InlineResponse20053Rules) GetDestPort() string {
	if o == nil || isNil(o.DestPort) {
		var ret string
		return ret
	}
	return *o.DestPort
}

// GetDestPortOk returns a tuple with the DestPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Rules) GetDestPortOk() (*string, bool) {
	if o == nil || isNil(o.DestPort) {
    return nil, false
	}
	return o.DestPort, true
}

// HasDestPort returns a boolean if a field has been set.
func (o *InlineResponse20053Rules) HasDestPort() bool {
	if o != nil && !isNil(o.DestPort) {
		return true
	}

	return false
}

// SetDestPort gets a reference to the given string and assigns it to the DestPort field.
func (o *InlineResponse20053Rules) SetDestPort(v string) {
	o.DestPort = &v
}

// GetDestCidr returns the DestCidr field value if set, zero value otherwise.
func (o *InlineResponse20053Rules) GetDestCidr() string {
	if o == nil || isNil(o.DestCidr) {
		var ret string
		return ret
	}
	return *o.DestCidr
}

// GetDestCidrOk returns a tuple with the DestCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Rules) GetDestCidrOk() (*string, bool) {
	if o == nil || isNil(o.DestCidr) {
    return nil, false
	}
	return o.DestCidr, true
}

// HasDestCidr returns a boolean if a field has been set.
func (o *InlineResponse20053Rules) HasDestCidr() bool {
	if o != nil && !isNil(o.DestCidr) {
		return true
	}

	return false
}

// SetDestCidr gets a reference to the given string and assigns it to the DestCidr field.
func (o *InlineResponse20053Rules) SetDestCidr(v string) {
	o.DestCidr = &v
}

// GetSyslogEnabled returns the SyslogEnabled field value if set, zero value otherwise.
func (o *InlineResponse20053Rules) GetSyslogEnabled() bool {
	if o == nil || isNil(o.SyslogEnabled) {
		var ret bool
		return ret
	}
	return *o.SyslogEnabled
}

// GetSyslogEnabledOk returns a tuple with the SyslogEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Rules) GetSyslogEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.SyslogEnabled) {
    return nil, false
	}
	return o.SyslogEnabled, true
}

// HasSyslogEnabled returns a boolean if a field has been set.
func (o *InlineResponse20053Rules) HasSyslogEnabled() bool {
	if o != nil && !isNil(o.SyslogEnabled) {
		return true
	}

	return false
}

// SetSyslogEnabled gets a reference to the given bool and assigns it to the SyslogEnabled field.
func (o *InlineResponse20053Rules) SetSyslogEnabled(v bool) {
	o.SyslogEnabled = &v
}

func (o InlineResponse20053Rules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !isNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.SrcPort) {
		toSerialize["srcPort"] = o.SrcPort
	}
	if !isNil(o.SrcCidr) {
		toSerialize["srcCidr"] = o.SrcCidr
	}
	if !isNil(o.DestPort) {
		toSerialize["destPort"] = o.DestPort
	}
	if !isNil(o.DestCidr) {
		toSerialize["destCidr"] = o.DestCidr
	}
	if !isNil(o.SyslogEnabled) {
		toSerialize["syslogEnabled"] = o.SyslogEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20053Rules struct {
	value *InlineResponse20053Rules
	isSet bool
}

func (v NullableInlineResponse20053Rules) Get() *InlineResponse20053Rules {
	return v.value
}

func (v *NullableInlineResponse20053Rules) Set(val *InlineResponse20053Rules) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20053Rules) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20053Rules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20053Rules(val *InlineResponse20053Rules) *NullableInlineResponse20053Rules {
	return &NullableInlineResponse20053Rules{value: val, isSet: true}
}

func (v NullableInlineResponse20053Rules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20053Rules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


