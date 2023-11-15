/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject52 struct for InlineObject52
type InlineObject52 struct {
	// The name of the new static route
	Name string `json:"name"`
	// The subnet of the static route
	Subnet string `json:"subnet"`
	// The gateway IP (next hop) of the static route
	GatewayIp string `json:"gatewayIp"`
	// The gateway IP (next hop) VLAN ID of the static route
	GatewayVlanId *string `json:"gatewayVlanId,omitempty"`
}

// NewInlineObject52 instantiates a new InlineObject52 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject52(name string, subnet string, gatewayIp string) *InlineObject52 {
	this := InlineObject52{}
	this.Name = name
	this.Subnet = subnet
	this.GatewayIp = gatewayIp
	return &this
}

// NewInlineObject52WithDefaults instantiates a new InlineObject52 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject52WithDefaults() *InlineObject52 {
	this := InlineObject52{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject52) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject52) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject52) SetName(v string) {
	o.Name = v
}

// GetSubnet returns the Subnet field value
func (o *InlineObject52) GetSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value
// and a boolean to check if the value has been set.
func (o *InlineObject52) GetSubnetOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subnet, true
}

// SetSubnet sets field value
func (o *InlineObject52) SetSubnet(v string) {
	o.Subnet = v
}

// GetGatewayIp returns the GatewayIp field value
func (o *InlineObject52) GetGatewayIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value
// and a boolean to check if the value has been set.
func (o *InlineObject52) GetGatewayIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GatewayIp, true
}

// SetGatewayIp sets field value
func (o *InlineObject52) SetGatewayIp(v string) {
	o.GatewayIp = v
}

// GetGatewayVlanId returns the GatewayVlanId field value if set, zero value otherwise.
func (o *InlineObject52) GetGatewayVlanId() string {
	if o == nil || isNil(o.GatewayVlanId) {
		var ret string
		return ret
	}
	return *o.GatewayVlanId
}

// GetGatewayVlanIdOk returns a tuple with the GatewayVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject52) GetGatewayVlanIdOk() (*string, bool) {
	if o == nil || isNil(o.GatewayVlanId) {
    return nil, false
	}
	return o.GatewayVlanId, true
}

// HasGatewayVlanId returns a boolean if a field has been set.
func (o *InlineObject52) HasGatewayVlanId() bool {
	if o != nil && !isNil(o.GatewayVlanId) {
		return true
	}

	return false
}

// SetGatewayVlanId gets a reference to the given string and assigns it to the GatewayVlanId field.
func (o *InlineObject52) SetGatewayVlanId(v string) {
	o.GatewayVlanId = &v
}

func (o InlineObject52) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["subnet"] = o.Subnet
	}
	if true {
		toSerialize["gatewayIp"] = o.GatewayIp
	}
	if !isNil(o.GatewayVlanId) {
		toSerialize["gatewayVlanId"] = o.GatewayVlanId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject52 struct {
	value *InlineObject52
	isSet bool
}

func (v NullableInlineObject52) Get() *InlineObject52 {
	return v.value
}

func (v *NullableInlineObject52) Set(val *InlineObject52) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject52) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject52) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject52(val *InlineObject52) *NullableInlineObject52 {
	return &NullableInlineObject52{value: val, isSet: true}
}

func (v NullableInlineObject52) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject52) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


