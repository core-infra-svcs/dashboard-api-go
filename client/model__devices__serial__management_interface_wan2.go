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

// DevicesSerialManagementInterfaceWan2 WAN 2 settings (only for MX devices)
type DevicesSerialManagementInterfaceWan2 struct {
	// Enable or disable the interface (only for MX devices). Valid values are 'enabled', 'disabled', and 'not configured'.
	WanEnabled *string `json:"wanEnabled,omitempty"`
	// Configure the interface to have static IP settings or use DHCP.
	UsingStaticIp *bool `json:"usingStaticIp,omitempty"`
	// The IP the device should use on the WAN.
	StaticIp *string `json:"staticIp,omitempty"`
	// The IP of the gateway on the WAN.
	StaticGatewayIp *string `json:"staticGatewayIp,omitempty"`
	// The subnet mask for the WAN.
	StaticSubnetMask *string `json:"staticSubnetMask,omitempty"`
	// Up to two DNS IPs.
	StaticDns []string `json:"staticDns,omitempty"`
	// The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.
	Vlan *int32 `json:"vlan,omitempty"`
}

// NewDevicesSerialManagementInterfaceWan2 instantiates a new DevicesSerialManagementInterfaceWan2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialManagementInterfaceWan2() *DevicesSerialManagementInterfaceWan2 {
	this := DevicesSerialManagementInterfaceWan2{}
	return &this
}

// NewDevicesSerialManagementInterfaceWan2WithDefaults instantiates a new DevicesSerialManagementInterfaceWan2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialManagementInterfaceWan2WithDefaults() *DevicesSerialManagementInterfaceWan2 {
	this := DevicesSerialManagementInterfaceWan2{}
	return &this
}

// GetWanEnabled returns the WanEnabled field value if set, zero value otherwise.
func (o *DevicesSerialManagementInterfaceWan2) GetWanEnabled() string {
	if o == nil || isNil(o.WanEnabled) {
		var ret string
		return ret
	}
	return *o.WanEnabled
}

// GetWanEnabledOk returns a tuple with the WanEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialManagementInterfaceWan2) GetWanEnabledOk() (*string, bool) {
	if o == nil || isNil(o.WanEnabled) {
    return nil, false
	}
	return o.WanEnabled, true
}

// HasWanEnabled returns a boolean if a field has been set.
func (o *DevicesSerialManagementInterfaceWan2) HasWanEnabled() bool {
	if o != nil && !isNil(o.WanEnabled) {
		return true
	}

	return false
}

// SetWanEnabled gets a reference to the given string and assigns it to the WanEnabled field.
func (o *DevicesSerialManagementInterfaceWan2) SetWanEnabled(v string) {
	o.WanEnabled = &v
}

// GetUsingStaticIp returns the UsingStaticIp field value if set, zero value otherwise.
func (o *DevicesSerialManagementInterfaceWan2) GetUsingStaticIp() bool {
	if o == nil || isNil(o.UsingStaticIp) {
		var ret bool
		return ret
	}
	return *o.UsingStaticIp
}

// GetUsingStaticIpOk returns a tuple with the UsingStaticIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialManagementInterfaceWan2) GetUsingStaticIpOk() (*bool, bool) {
	if o == nil || isNil(o.UsingStaticIp) {
    return nil, false
	}
	return o.UsingStaticIp, true
}

// HasUsingStaticIp returns a boolean if a field has been set.
func (o *DevicesSerialManagementInterfaceWan2) HasUsingStaticIp() bool {
	if o != nil && !isNil(o.UsingStaticIp) {
		return true
	}

	return false
}

// SetUsingStaticIp gets a reference to the given bool and assigns it to the UsingStaticIp field.
func (o *DevicesSerialManagementInterfaceWan2) SetUsingStaticIp(v bool) {
	o.UsingStaticIp = &v
}

// GetStaticIp returns the StaticIp field value if set, zero value otherwise.
func (o *DevicesSerialManagementInterfaceWan2) GetStaticIp() string {
	if o == nil || isNil(o.StaticIp) {
		var ret string
		return ret
	}
	return *o.StaticIp
}

// GetStaticIpOk returns a tuple with the StaticIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialManagementInterfaceWan2) GetStaticIpOk() (*string, bool) {
	if o == nil || isNil(o.StaticIp) {
    return nil, false
	}
	return o.StaticIp, true
}

// HasStaticIp returns a boolean if a field has been set.
func (o *DevicesSerialManagementInterfaceWan2) HasStaticIp() bool {
	if o != nil && !isNil(o.StaticIp) {
		return true
	}

	return false
}

// SetStaticIp gets a reference to the given string and assigns it to the StaticIp field.
func (o *DevicesSerialManagementInterfaceWan2) SetStaticIp(v string) {
	o.StaticIp = &v
}

// GetStaticGatewayIp returns the StaticGatewayIp field value if set, zero value otherwise.
func (o *DevicesSerialManagementInterfaceWan2) GetStaticGatewayIp() string {
	if o == nil || isNil(o.StaticGatewayIp) {
		var ret string
		return ret
	}
	return *o.StaticGatewayIp
}

// GetStaticGatewayIpOk returns a tuple with the StaticGatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialManagementInterfaceWan2) GetStaticGatewayIpOk() (*string, bool) {
	if o == nil || isNil(o.StaticGatewayIp) {
    return nil, false
	}
	return o.StaticGatewayIp, true
}

// HasStaticGatewayIp returns a boolean if a field has been set.
func (o *DevicesSerialManagementInterfaceWan2) HasStaticGatewayIp() bool {
	if o != nil && !isNil(o.StaticGatewayIp) {
		return true
	}

	return false
}

// SetStaticGatewayIp gets a reference to the given string and assigns it to the StaticGatewayIp field.
func (o *DevicesSerialManagementInterfaceWan2) SetStaticGatewayIp(v string) {
	o.StaticGatewayIp = &v
}

// GetStaticSubnetMask returns the StaticSubnetMask field value if set, zero value otherwise.
func (o *DevicesSerialManagementInterfaceWan2) GetStaticSubnetMask() string {
	if o == nil || isNil(o.StaticSubnetMask) {
		var ret string
		return ret
	}
	return *o.StaticSubnetMask
}

// GetStaticSubnetMaskOk returns a tuple with the StaticSubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialManagementInterfaceWan2) GetStaticSubnetMaskOk() (*string, bool) {
	if o == nil || isNil(o.StaticSubnetMask) {
    return nil, false
	}
	return o.StaticSubnetMask, true
}

// HasStaticSubnetMask returns a boolean if a field has been set.
func (o *DevicesSerialManagementInterfaceWan2) HasStaticSubnetMask() bool {
	if o != nil && !isNil(o.StaticSubnetMask) {
		return true
	}

	return false
}

// SetStaticSubnetMask gets a reference to the given string and assigns it to the StaticSubnetMask field.
func (o *DevicesSerialManagementInterfaceWan2) SetStaticSubnetMask(v string) {
	o.StaticSubnetMask = &v
}

// GetStaticDns returns the StaticDns field value if set, zero value otherwise.
func (o *DevicesSerialManagementInterfaceWan2) GetStaticDns() []string {
	if o == nil || isNil(o.StaticDns) {
		var ret []string
		return ret
	}
	return o.StaticDns
}

// GetStaticDnsOk returns a tuple with the StaticDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialManagementInterfaceWan2) GetStaticDnsOk() ([]string, bool) {
	if o == nil || isNil(o.StaticDns) {
    return nil, false
	}
	return o.StaticDns, true
}

// HasStaticDns returns a boolean if a field has been set.
func (o *DevicesSerialManagementInterfaceWan2) HasStaticDns() bool {
	if o != nil && !isNil(o.StaticDns) {
		return true
	}

	return false
}

// SetStaticDns gets a reference to the given []string and assigns it to the StaticDns field.
func (o *DevicesSerialManagementInterfaceWan2) SetStaticDns(v []string) {
	o.StaticDns = v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *DevicesSerialManagementInterfaceWan2) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialManagementInterfaceWan2) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *DevicesSerialManagementInterfaceWan2) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *DevicesSerialManagementInterfaceWan2) SetVlan(v int32) {
	o.Vlan = &v
}

func (o DevicesSerialManagementInterfaceWan2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WanEnabled) {
		toSerialize["wanEnabled"] = o.WanEnabled
	}
	if !isNil(o.UsingStaticIp) {
		toSerialize["usingStaticIp"] = o.UsingStaticIp
	}
	if !isNil(o.StaticIp) {
		toSerialize["staticIp"] = o.StaticIp
	}
	if !isNil(o.StaticGatewayIp) {
		toSerialize["staticGatewayIp"] = o.StaticGatewayIp
	}
	if !isNil(o.StaticSubnetMask) {
		toSerialize["staticSubnetMask"] = o.StaticSubnetMask
	}
	if !isNil(o.StaticDns) {
		toSerialize["staticDns"] = o.StaticDns
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialManagementInterfaceWan2 struct {
	value *DevicesSerialManagementInterfaceWan2
	isSet bool
}

func (v NullableDevicesSerialManagementInterfaceWan2) Get() *DevicesSerialManagementInterfaceWan2 {
	return v.value
}

func (v *NullableDevicesSerialManagementInterfaceWan2) Set(val *DevicesSerialManagementInterfaceWan2) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialManagementInterfaceWan2) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialManagementInterfaceWan2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialManagementInterfaceWan2(val *DevicesSerialManagementInterfaceWan2) *NullableDevicesSerialManagementInterfaceWan2 {
	return &NullableDevicesSerialManagementInterfaceWan2{value: val, isSet: true}
}

func (v NullableDevicesSerialManagementInterfaceWan2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialManagementInterfaceWan2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


