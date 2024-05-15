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

// InlineObject61 struct for InlineObject61
type InlineObject61 struct {
	// Name of the route
	Name string `json:"name"`
	// Subnet of the route
	Subnet string `json:"subnet"`
	// Gateway IP address (next hop)
	GatewayIp string `json:"gatewayIp"`
	// Gateway VLAN ID
	GatewayVlanId *string `json:"gatewayVlanId,omitempty"`
}

// NewInlineObject61 instantiates a new InlineObject61 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject61(name string, subnet string, gatewayIp string) *InlineObject61 {
	this := InlineObject61{}
	this.Name = name
	this.Subnet = subnet
	this.GatewayIp = gatewayIp
	return &this
}

// NewInlineObject61WithDefaults instantiates a new InlineObject61 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject61WithDefaults() *InlineObject61 {
	this := InlineObject61{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject61) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject61) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject61) SetName(v string) {
	o.Name = v
}

// GetSubnet returns the Subnet field value
func (o *InlineObject61) GetSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value
// and a boolean to check if the value has been set.
func (o *InlineObject61) GetSubnetOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subnet, true
}

// SetSubnet sets field value
func (o *InlineObject61) SetSubnet(v string) {
	o.Subnet = v
}

// GetGatewayIp returns the GatewayIp field value
func (o *InlineObject61) GetGatewayIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value
// and a boolean to check if the value has been set.
func (o *InlineObject61) GetGatewayIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GatewayIp, true
}

// SetGatewayIp sets field value
func (o *InlineObject61) SetGatewayIp(v string) {
	o.GatewayIp = v
}

// GetGatewayVlanId returns the GatewayVlanId field value if set, zero value otherwise.
func (o *InlineObject61) GetGatewayVlanId() string {
	if o == nil || isNil(o.GatewayVlanId) {
		var ret string
		return ret
	}
	return *o.GatewayVlanId
}

// GetGatewayVlanIdOk returns a tuple with the GatewayVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject61) GetGatewayVlanIdOk() (*string, bool) {
	if o == nil || isNil(o.GatewayVlanId) {
    return nil, false
	}
	return o.GatewayVlanId, true
}

// HasGatewayVlanId returns a boolean if a field has been set.
func (o *InlineObject61) HasGatewayVlanId() bool {
	if o != nil && !isNil(o.GatewayVlanId) {
		return true
	}

	return false
}

// SetGatewayVlanId gets a reference to the given string and assigns it to the GatewayVlanId field.
func (o *InlineObject61) SetGatewayVlanId(v string) {
	o.GatewayVlanId = &v
}

func (o InlineObject61) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject61 struct {
	value *InlineObject61
	isSet bool
}

func (v NullableInlineObject61) Get() *InlineObject61 {
	return v.value
}

func (v *NullableInlineObject61) Set(val *InlineObject61) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject61) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject61) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject61(val *InlineObject61) *NullableInlineObject61 {
	return &NullableInlineObject61{value: val, isSet: true}
}

func (v NullableInlineObject61) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject61) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


