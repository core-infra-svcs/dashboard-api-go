/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialCellularGatewayPortForwardingRulesRules struct for DevicesSerialCellularGatewayPortForwardingRulesRules
type DevicesSerialCellularGatewayPortForwardingRulesRules struct {
	// A descriptive name for the rule
	Name *string `json:"name,omitempty"`
	// The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
	LanIp string `json:"lanIp"`
	// A port or port ranges that will be forwarded to the host on the LAN
	PublicPort string `json:"publicPort"`
	// A port or port ranges that will receive the forwarded traffic from the WAN
	LocalPort string `json:"localPort"`
	// An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges.
	AllowedIps []string `json:"allowedIps,omitempty"`
	// TCP or UDP
	Protocol string `json:"protocol"`
	// `any` or `restricted`. Specify the right to make inbound connections on the specified ports or port ranges. If `restricted`, a list of allowed IPs is mandatory.
	Access string `json:"access"`
}

// NewDevicesSerialCellularGatewayPortForwardingRulesRules instantiates a new DevicesSerialCellularGatewayPortForwardingRulesRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialCellularGatewayPortForwardingRulesRules(lanIp string, publicPort string, localPort string, protocol string, access string) *DevicesSerialCellularGatewayPortForwardingRulesRules {
	this := DevicesSerialCellularGatewayPortForwardingRulesRules{}
	this.LanIp = lanIp
	this.PublicPort = publicPort
	this.LocalPort = localPort
	this.Protocol = protocol
	this.Access = access
	return &this
}

// NewDevicesSerialCellularGatewayPortForwardingRulesRulesWithDefaults instantiates a new DevicesSerialCellularGatewayPortForwardingRulesRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialCellularGatewayPortForwardingRulesRulesWithDefaults() *DevicesSerialCellularGatewayPortForwardingRulesRules {
	this := DevicesSerialCellularGatewayPortForwardingRulesRules{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) SetName(v string) {
	o.Name = &v
}

// GetLanIp returns the LanIp field value
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetLanIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LanIp
}

// GetLanIpOk returns a tuple with the LanIp field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetLanIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LanIp, true
}

// SetLanIp sets field value
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) SetLanIp(v string) {
	o.LanIp = v
}

// GetPublicPort returns the PublicPort field value
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetPublicPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicPort
}

// GetPublicPortOk returns a tuple with the PublicPort field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetPublicPortOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PublicPort, true
}

// SetPublicPort sets field value
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) SetPublicPort(v string) {
	o.PublicPort = v
}

// GetLocalPort returns the LocalPort field value
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetLocalPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocalPort
}

// GetLocalPortOk returns a tuple with the LocalPort field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetLocalPortOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LocalPort, true
}

// SetLocalPort sets field value
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) SetLocalPort(v string) {
	o.LocalPort = v
}

// GetAllowedIps returns the AllowedIps field value if set, zero value otherwise.
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetAllowedIps() []string {
	if o == nil || isNil(o.AllowedIps) {
		var ret []string
		return ret
	}
	return o.AllowedIps
}

// GetAllowedIpsOk returns a tuple with the AllowedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetAllowedIpsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedIps) {
    return nil, false
	}
	return o.AllowedIps, true
}

// HasAllowedIps returns a boolean if a field has been set.
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) HasAllowedIps() bool {
	if o != nil && !isNil(o.AllowedIps) {
		return true
	}

	return false
}

// SetAllowedIps gets a reference to the given []string and assigns it to the AllowedIps field.
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) SetAllowedIps(v []string) {
	o.AllowedIps = v
}

// GetProtocol returns the Protocol field value
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetProtocolOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) SetProtocol(v string) {
	o.Protocol = v
}

// GetAccess returns the Access field value
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetAccessOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) SetAccess(v string) {
	o.Access = v
}

func (o DevicesSerialCellularGatewayPortForwardingRulesRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["lanIp"] = o.LanIp
	}
	if true {
		toSerialize["publicPort"] = o.PublicPort
	}
	if true {
		toSerialize["localPort"] = o.LocalPort
	}
	if !isNil(o.AllowedIps) {
		toSerialize["allowedIps"] = o.AllowedIps
	}
	if true {
		toSerialize["protocol"] = o.Protocol
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialCellularGatewayPortForwardingRulesRules struct {
	value *DevicesSerialCellularGatewayPortForwardingRulesRules
	isSet bool
}

func (v NullableDevicesSerialCellularGatewayPortForwardingRulesRules) Get() *DevicesSerialCellularGatewayPortForwardingRulesRules {
	return v.value
}

func (v *NullableDevicesSerialCellularGatewayPortForwardingRulesRules) Set(val *DevicesSerialCellularGatewayPortForwardingRulesRules) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialCellularGatewayPortForwardingRulesRules) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialCellularGatewayPortForwardingRulesRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialCellularGatewayPortForwardingRulesRules(val *DevicesSerialCellularGatewayPortForwardingRulesRules) *NullableDevicesSerialCellularGatewayPortForwardingRulesRules {
	return &NullableDevicesSerialCellularGatewayPortForwardingRulesRules{value: val, isSet: true}
}

func (v NullableDevicesSerialCellularGatewayPortForwardingRulesRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialCellularGatewayPortForwardingRulesRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


