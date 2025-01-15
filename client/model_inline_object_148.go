/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject148 struct for InlineObject148
type InlineObject148 struct {
	// The VLAN of the incoming packet. A null value will match any VLAN.
	Vlan *int32 `json:"vlan,omitempty"`
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
	// DSCP tag that should be assigned to incoming packet. Set this to -1 to trust incoming DSCP. Default value is 0
	Dscp *int32 `json:"dscp,omitempty"`
}

// NewInlineObject148 instantiates a new InlineObject148 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject148() *InlineObject148 {
	this := InlineObject148{}
	return &this
}

// NewInlineObject148WithDefaults instantiates a new InlineObject148 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject148WithDefaults() *InlineObject148 {
	this := InlineObject148{}
	return &this
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineObject148) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject148) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineObject148) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *InlineObject148) SetVlan(v int32) {
	o.Vlan = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *InlineObject148) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject148) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *InlineObject148) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *InlineObject148) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *InlineObject148) GetSrcPort() int32 {
	if o == nil || isNil(o.SrcPort) {
		var ret int32
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject148) GetSrcPortOk() (*int32, bool) {
	if o == nil || isNil(o.SrcPort) {
    return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *InlineObject148) HasSrcPort() bool {
	if o != nil && !isNil(o.SrcPort) {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given int32 and assigns it to the SrcPort field.
func (o *InlineObject148) SetSrcPort(v int32) {
	o.SrcPort = &v
}

// GetSrcPortRange returns the SrcPortRange field value if set, zero value otherwise.
func (o *InlineObject148) GetSrcPortRange() string {
	if o == nil || isNil(o.SrcPortRange) {
		var ret string
		return ret
	}
	return *o.SrcPortRange
}

// GetSrcPortRangeOk returns a tuple with the SrcPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject148) GetSrcPortRangeOk() (*string, bool) {
	if o == nil || isNil(o.SrcPortRange) {
    return nil, false
	}
	return o.SrcPortRange, true
}

// HasSrcPortRange returns a boolean if a field has been set.
func (o *InlineObject148) HasSrcPortRange() bool {
	if o != nil && !isNil(o.SrcPortRange) {
		return true
	}

	return false
}

// SetSrcPortRange gets a reference to the given string and assigns it to the SrcPortRange field.
func (o *InlineObject148) SetSrcPortRange(v string) {
	o.SrcPortRange = &v
}

// GetDstPort returns the DstPort field value if set, zero value otherwise.
func (o *InlineObject148) GetDstPort() int32 {
	if o == nil || isNil(o.DstPort) {
		var ret int32
		return ret
	}
	return *o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject148) GetDstPortOk() (*int32, bool) {
	if o == nil || isNil(o.DstPort) {
    return nil, false
	}
	return o.DstPort, true
}

// HasDstPort returns a boolean if a field has been set.
func (o *InlineObject148) HasDstPort() bool {
	if o != nil && !isNil(o.DstPort) {
		return true
	}

	return false
}

// SetDstPort gets a reference to the given int32 and assigns it to the DstPort field.
func (o *InlineObject148) SetDstPort(v int32) {
	o.DstPort = &v
}

// GetDstPortRange returns the DstPortRange field value if set, zero value otherwise.
func (o *InlineObject148) GetDstPortRange() string {
	if o == nil || isNil(o.DstPortRange) {
		var ret string
		return ret
	}
	return *o.DstPortRange
}

// GetDstPortRangeOk returns a tuple with the DstPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject148) GetDstPortRangeOk() (*string, bool) {
	if o == nil || isNil(o.DstPortRange) {
    return nil, false
	}
	return o.DstPortRange, true
}

// HasDstPortRange returns a boolean if a field has been set.
func (o *InlineObject148) HasDstPortRange() bool {
	if o != nil && !isNil(o.DstPortRange) {
		return true
	}

	return false
}

// SetDstPortRange gets a reference to the given string and assigns it to the DstPortRange field.
func (o *InlineObject148) SetDstPortRange(v string) {
	o.DstPortRange = &v
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *InlineObject148) GetDscp() int32 {
	if o == nil || isNil(o.Dscp) {
		var ret int32
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject148) GetDscpOk() (*int32, bool) {
	if o == nil || isNil(o.Dscp) {
    return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *InlineObject148) HasDscp() bool {
	if o != nil && !isNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given int32 and assigns it to the Dscp field.
func (o *InlineObject148) SetDscp(v int32) {
	o.Dscp = &v
}

func (o InlineObject148) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vlan) {
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

type NullableInlineObject148 struct {
	value *InlineObject148
	isSet bool
}

func (v NullableInlineObject148) Get() *InlineObject148 {
	return v.value
}

func (v *NullableInlineObject148) Set(val *InlineObject148) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject148) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject148) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject148(val *InlineObject148) *NullableInlineObject148 {
	return &NullableInlineObject148{value: val, isSet: true}
}

func (v NullableInlineObject148) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject148) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


