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

// InlineResponse200200Rules struct for InlineResponse200200Rules
type InlineResponse200200Rules struct {
	// Description of the rule (optional)
	Comment *string `json:"comment,omitempty"`
	// 'allow' or 'deny' traffic specified by this rule
	Policy string `json:"policy"`
	// The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	Protocol string `json:"protocol"`
	// Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	DestPort *string `json:"destPort,omitempty"`
	// Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestCidr string `json:"destCidr"`
}

// NewInlineResponse200200Rules instantiates a new InlineResponse200200Rules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200200Rules(policy string, protocol string, destCidr string) *InlineResponse200200Rules {
	this := InlineResponse200200Rules{}
	this.Policy = policy
	this.Protocol = protocol
	this.DestCidr = destCidr
	return &this
}

// NewInlineResponse200200RulesWithDefaults instantiates a new InlineResponse200200Rules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200200RulesWithDefaults() *InlineResponse200200Rules {
	this := InlineResponse200200Rules{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *InlineResponse200200Rules) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200Rules) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *InlineResponse200200Rules) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *InlineResponse200200Rules) SetComment(v string) {
	o.Comment = &v
}

// GetPolicy returns the Policy field value
func (o *InlineResponse200200Rules) GetPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200200Rules) GetPolicyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *InlineResponse200200Rules) SetPolicy(v string) {
	o.Policy = v
}

// GetProtocol returns the Protocol field value
func (o *InlineResponse200200Rules) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200200Rules) GetProtocolOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *InlineResponse200200Rules) SetProtocol(v string) {
	o.Protocol = v
}

// GetDestPort returns the DestPort field value if set, zero value otherwise.
func (o *InlineResponse200200Rules) GetDestPort() string {
	if o == nil || isNil(o.DestPort) {
		var ret string
		return ret
	}
	return *o.DestPort
}

// GetDestPortOk returns a tuple with the DestPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200Rules) GetDestPortOk() (*string, bool) {
	if o == nil || isNil(o.DestPort) {
    return nil, false
	}
	return o.DestPort, true
}

// HasDestPort returns a boolean if a field has been set.
func (o *InlineResponse200200Rules) HasDestPort() bool {
	if o != nil && !isNil(o.DestPort) {
		return true
	}

	return false
}

// SetDestPort gets a reference to the given string and assigns it to the DestPort field.
func (o *InlineResponse200200Rules) SetDestPort(v string) {
	o.DestPort = &v
}

// GetDestCidr returns the DestCidr field value
func (o *InlineResponse200200Rules) GetDestCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestCidr
}

// GetDestCidrOk returns a tuple with the DestCidr field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200200Rules) GetDestCidrOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DestCidr, true
}

// SetDestCidr sets field value
func (o *InlineResponse200200Rules) SetDestCidr(v string) {
	o.DestCidr = v
}

func (o InlineResponse200200Rules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if true {
		toSerialize["policy"] = o.Policy
	}
	if true {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.DestPort) {
		toSerialize["destPort"] = o.DestPort
	}
	if true {
		toSerialize["destCidr"] = o.DestCidr
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200200Rules struct {
	value *InlineResponse200200Rules
	isSet bool
}

func (v NullableInlineResponse200200Rules) Get() *InlineResponse200200Rules {
	return v.value
}

func (v *NullableInlineResponse200200Rules) Set(val *InlineResponse200200Rules) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200200Rules) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200200Rules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200200Rules(val *InlineResponse200200Rules) *NullableInlineResponse200200Rules {
	return &NullableInlineResponse200200Rules{value: val, isSet: true}
}

func (v NullableInlineResponse200200Rules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200200Rules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


