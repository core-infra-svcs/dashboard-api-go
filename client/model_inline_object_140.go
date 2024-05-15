/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject140 struct for InlineObject140
type InlineObject140 struct {
	// The VLAN of the incoming packet. A null value will match any VLAN.
	Vlan int32 `json:"vlan"`
	// The protocol of the incoming packet. Default value is \"ANY\"
	Protocol *string `json:"protocol,omitempty"`
	// The source port of the incoming packet. Applicable only if protocol is TCP or UDP.
	SrcPort *int32 `json:"srcPort,omitempty"`
	// The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP.
	SrcPortRange *string `json:"srcPortRange,omitempty"`
	// The destination port of the incoming packet. Applicable only if protocol is TCP or UDP.
	DstPort *int32 `json:"dstPort,omitempty"`
	// The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP.
	DstPortRange *string `json:"dstPortRange,omitempty"`
	// DSCP tag for the incoming packet. Set this to -1 to trust incoming DSCP. Default value is 0
	Dscp *int32 `json:"dscp,omitempty"`
}

// NewInlineObject140 instantiates a new InlineObject140 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject140(vlan int32) *InlineObject140 {
	this := InlineObject140{}
	this.Vlan = vlan
	return &this
}

// NewInlineObject140WithDefaults instantiates a new InlineObject140 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject140WithDefaults() *InlineObject140 {
	this := InlineObject140{}
	return &this
}

// GetVlan returns the Vlan field value
func (o *InlineObject140) GetVlan() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value
// and a boolean to check if the value has been set.
func (o *InlineObject140) GetVlanOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Vlan, true
}

// SetVlan sets field value
func (o *InlineObject140) SetVlan(v int32) {
	o.Vlan = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *InlineObject140) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject140) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *InlineObject140) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *InlineObject140) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *InlineObject140) GetSrcPort() int32 {
	if o == nil || isNil(o.SrcPort) {
		var ret int32
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject140) GetSrcPortOk() (*int32, bool) {
	if o == nil || isNil(o.SrcPort) {
    return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *InlineObject140) HasSrcPort() bool {
	if o != nil && !isNil(o.SrcPort) {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given int32 and assigns it to the SrcPort field.
func (o *InlineObject140) SetSrcPort(v int32) {
	o.SrcPort = &v
}

// GetSrcPortRange returns the SrcPortRange field value if set, zero value otherwise.
func (o *InlineObject140) GetSrcPortRange() string {
	if o == nil || isNil(o.SrcPortRange) {
		var ret string
		return ret
	}
	return *o.SrcPortRange
}

// GetSrcPortRangeOk returns a tuple with the SrcPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject140) GetSrcPortRangeOk() (*string, bool) {
	if o == nil || isNil(o.SrcPortRange) {
    return nil, false
	}
	return o.SrcPortRange, true
}

// HasSrcPortRange returns a boolean if a field has been set.
func (o *InlineObject140) HasSrcPortRange() bool {
	if o != nil && !isNil(o.SrcPortRange) {
		return true
	}

	return false
}

// SetSrcPortRange gets a reference to the given string and assigns it to the SrcPortRange field.
func (o *InlineObject140) SetSrcPortRange(v string) {
	o.SrcPortRange = &v
}

// GetDstPort returns the DstPort field value if set, zero value otherwise.
func (o *InlineObject140) GetDstPort() int32 {
	if o == nil || isNil(o.DstPort) {
		var ret int32
		return ret
	}
	return *o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject140) GetDstPortOk() (*int32, bool) {
	if o == nil || isNil(o.DstPort) {
    return nil, false
	}
	return o.DstPort, true
}

// HasDstPort returns a boolean if a field has been set.
func (o *InlineObject140) HasDstPort() bool {
	if o != nil && !isNil(o.DstPort) {
		return true
	}

	return false
}

// SetDstPort gets a reference to the given int32 and assigns it to the DstPort field.
func (o *InlineObject140) SetDstPort(v int32) {
	o.DstPort = &v
}

// GetDstPortRange returns the DstPortRange field value if set, zero value otherwise.
func (o *InlineObject140) GetDstPortRange() string {
	if o == nil || isNil(o.DstPortRange) {
		var ret string
		return ret
	}
	return *o.DstPortRange
}

// GetDstPortRangeOk returns a tuple with the DstPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject140) GetDstPortRangeOk() (*string, bool) {
	if o == nil || isNil(o.DstPortRange) {
    return nil, false
	}
	return o.DstPortRange, true
}

// HasDstPortRange returns a boolean if a field has been set.
func (o *InlineObject140) HasDstPortRange() bool {
	if o != nil && !isNil(o.DstPortRange) {
		return true
	}

	return false
}

// SetDstPortRange gets a reference to the given string and assigns it to the DstPortRange field.
func (o *InlineObject140) SetDstPortRange(v string) {
	o.DstPortRange = &v
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *InlineObject140) GetDscp() int32 {
	if o == nil || isNil(o.Dscp) {
		var ret int32
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject140) GetDscpOk() (*int32, bool) {
	if o == nil || isNil(o.Dscp) {
    return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *InlineObject140) HasDscp() bool {
	if o != nil && !isNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given int32 and assigns it to the Dscp field.
func (o *InlineObject140) SetDscp(v int32) {
	o.Dscp = &v
}

func (o InlineObject140) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.SrcPort) {
		toSerialize["srcPort"] = o.SrcPort
	}
	if !isNil(o.SrcPortRange) {
		toSerialize["srcPortRange"] = o.SrcPortRange
	}
	if !isNil(o.DstPort) {
		toSerialize["dstPort"] = o.DstPort
	}
	if !isNil(o.DstPortRange) {
		toSerialize["dstPortRange"] = o.DstPortRange
	}
	if !isNil(o.Dscp) {
		toSerialize["dscp"] = o.Dscp
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject140 struct {
	value *InlineObject140
	isSet bool
}

func (v NullableInlineObject140) Get() *InlineObject140 {
	return v.value
}

func (v *NullableInlineObject140) Set(val *InlineObject140) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject140) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject140) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject140(val *InlineObject140) *NullableInlineObject140 {
	return &NullableInlineObject140{value: val, isSet: true}
}

func (v NullableInlineObject140) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject140) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


