/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject39 struct for InlineObject39
type InlineObject39 struct {
	// The name of the new static route
	Name string `json:"name"`
	// The subnet of the static route
	Subnet string `json:"subnet"`
	// The gateway IP (next hop) of the static route
	GatewayIp string `json:"gatewayIp"`
	// The gateway IP (next hop) VLAN ID of the static route
	GatewayVlanId *string `json:"gatewayVlanId,omitempty"`
}

// NewInlineObject39 instantiates a new InlineObject39 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject39(name string, subnet string, gatewayIp string) *InlineObject39 {
	this := InlineObject39{}
	this.Name = name
	this.Subnet = subnet
	this.GatewayIp = gatewayIp
	return &this
}

// NewInlineObject39WithDefaults instantiates a new InlineObject39 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject39WithDefaults() *InlineObject39 {
	this := InlineObject39{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject39) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject39) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject39) SetName(v string) {
	o.Name = v
}

// GetSubnet returns the Subnet field value
func (o *InlineObject39) GetSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value
// and a boolean to check if the value has been set.
func (o *InlineObject39) GetSubnetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subnet, true
}

// SetSubnet sets field value
func (o *InlineObject39) SetSubnet(v string) {
	o.Subnet = v
}

// GetGatewayIp returns the GatewayIp field value
func (o *InlineObject39) GetGatewayIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value
// and a boolean to check if the value has been set.
func (o *InlineObject39) GetGatewayIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GatewayIp, true
}

// SetGatewayIp sets field value
func (o *InlineObject39) SetGatewayIp(v string) {
	o.GatewayIp = v
}

// GetGatewayVlanId returns the GatewayVlanId field value if set, zero value otherwise.
func (o *InlineObject39) GetGatewayVlanId() string {
	if o == nil || o.GatewayVlanId == nil {
		var ret string
		return ret
	}
	return *o.GatewayVlanId
}

// GetGatewayVlanIdOk returns a tuple with the GatewayVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject39) GetGatewayVlanIdOk() (*string, bool) {
	if o == nil || o.GatewayVlanId == nil {
		return nil, false
	}
	return o.GatewayVlanId, true
}

// HasGatewayVlanId returns a boolean if a field has been set.
func (o *InlineObject39) HasGatewayVlanId() bool {
	if o != nil && o.GatewayVlanId != nil {
		return true
	}

	return false
}

// SetGatewayVlanId gets a reference to the given string and assigns it to the GatewayVlanId field.
func (o *InlineObject39) SetGatewayVlanId(v string) {
	o.GatewayVlanId = &v
}

func (o InlineObject39) MarshalJSON() ([]byte, error) {
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
	if o.GatewayVlanId != nil {
		toSerialize["gatewayVlanId"] = o.GatewayVlanId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject39 struct {
	value *InlineObject39
	isSet bool
}

func (v NullableInlineObject39) Get() *InlineObject39 {
	return v.value
}

func (v *NullableInlineObject39) Set(val *InlineObject39) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject39) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject39) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject39(val *InlineObject39) *NullableInlineObject39 {
	return &NullableInlineObject39{value: val, isSet: true}
}

func (v NullableInlineObject39) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject39) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


