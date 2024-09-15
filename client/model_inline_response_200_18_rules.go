/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20018Rules struct for InlineResponse20018Rules
type InlineResponse20018Rules struct {
	// A descriptive name for the rule
	Name *string `json:"name,omitempty"`
	// The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
	LanIp *string `json:"lanIp,omitempty"`
	// A port or port ranges that will be forwarded to the host on the LAN
	PublicPort *string `json:"publicPort,omitempty"`
	// A port or port ranges that will receive the forwarded traffic from the WAN
	LocalPort *string `json:"localPort,omitempty"`
	// An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges.
	AllowedIps []string `json:"allowedIps,omitempty"`
	// TCP or UDP
	Protocol *string `json:"protocol,omitempty"`
	// `any` or `restricted`. Specify the right to make inbound connections on the specified ports or port ranges. If `restricted`, a list of allowed IPs is mandatory.
	Access *string `json:"access,omitempty"`
}

// NewInlineResponse20018Rules instantiates a new InlineResponse20018Rules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20018Rules() *InlineResponse20018Rules {
	this := InlineResponse20018Rules{}
	return &this
}

// NewInlineResponse20018RulesWithDefaults instantiates a new InlineResponse20018Rules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20018RulesWithDefaults() *InlineResponse20018Rules {
	this := InlineResponse20018Rules{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20018Rules) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20018Rules) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20018Rules) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20018Rules) SetName(v string) {
	o.Name = &v
}

// GetLanIp returns the LanIp field value if set, zero value otherwise.
func (o *InlineResponse20018Rules) GetLanIp() string {
	if o == nil || isNil(o.LanIp) {
		var ret string
		return ret
	}
	return *o.LanIp
}

// GetLanIpOk returns a tuple with the LanIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20018Rules) GetLanIpOk() (*string, bool) {
	if o == nil || isNil(o.LanIp) {
    return nil, false
	}
	return o.LanIp, true
}

// HasLanIp returns a boolean if a field has been set.
func (o *InlineResponse20018Rules) HasLanIp() bool {
	if o != nil && !isNil(o.LanIp) {
		return true
	}

	return false
}

// SetLanIp gets a reference to the given string and assigns it to the LanIp field.
func (o *InlineResponse20018Rules) SetLanIp(v string) {
	o.LanIp = &v
}

// GetPublicPort returns the PublicPort field value if set, zero value otherwise.
func (o *InlineResponse20018Rules) GetPublicPort() string {
	if o == nil || isNil(o.PublicPort) {
		var ret string
		return ret
	}
	return *o.PublicPort
}

// GetPublicPortOk returns a tuple with the PublicPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20018Rules) GetPublicPortOk() (*string, bool) {
	if o == nil || isNil(o.PublicPort) {
    return nil, false
	}
	return o.PublicPort, true
}

// HasPublicPort returns a boolean if a field has been set.
func (o *InlineResponse20018Rules) HasPublicPort() bool {
	if o != nil && !isNil(o.PublicPort) {
		return true
	}

	return false
}

// SetPublicPort gets a reference to the given string and assigns it to the PublicPort field.
func (o *InlineResponse20018Rules) SetPublicPort(v string) {
	o.PublicPort = &v
}

// GetLocalPort returns the LocalPort field value if set, zero value otherwise.
func (o *InlineResponse20018Rules) GetLocalPort() string {
	if o == nil || isNil(o.LocalPort) {
		var ret string
		return ret
	}
	return *o.LocalPort
}

// GetLocalPortOk returns a tuple with the LocalPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20018Rules) GetLocalPortOk() (*string, bool) {
	if o == nil || isNil(o.LocalPort) {
    return nil, false
	}
	return o.LocalPort, true
}

// HasLocalPort returns a boolean if a field has been set.
func (o *InlineResponse20018Rules) HasLocalPort() bool {
	if o != nil && !isNil(o.LocalPort) {
		return true
	}

	return false
}

// SetLocalPort gets a reference to the given string and assigns it to the LocalPort field.
func (o *InlineResponse20018Rules) SetLocalPort(v string) {
	o.LocalPort = &v
}

// GetAllowedIps returns the AllowedIps field value if set, zero value otherwise.
func (o *InlineResponse20018Rules) GetAllowedIps() []string {
	if o == nil || isNil(o.AllowedIps) {
		var ret []string
		return ret
	}
	return o.AllowedIps
}

// GetAllowedIpsOk returns a tuple with the AllowedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20018Rules) GetAllowedIpsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedIps) {
    return nil, false
	}
	return o.AllowedIps, true
}

// HasAllowedIps returns a boolean if a field has been set.
func (o *InlineResponse20018Rules) HasAllowedIps() bool {
	if o != nil && !isNil(o.AllowedIps) {
		return true
	}

	return false
}

// SetAllowedIps gets a reference to the given []string and assigns it to the AllowedIps field.
func (o *InlineResponse20018Rules) SetAllowedIps(v []string) {
	o.AllowedIps = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *InlineResponse20018Rules) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20018Rules) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *InlineResponse20018Rules) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *InlineResponse20018Rules) SetProtocol(v string) {
	o.Protocol = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *InlineResponse20018Rules) GetAccess() string {
	if o == nil || isNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20018Rules) GetAccessOk() (*string, bool) {
	if o == nil || isNil(o.Access) {
    return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *InlineResponse20018Rules) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *InlineResponse20018Rules) SetAccess(v string) {
	o.Access = &v
}

func (o InlineResponse20018Rules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.LanIp) {
		toSerialize["lanIp"] = o.LanIp
	}
	if !isNil(o.PublicPort) {
		toSerialize["publicPort"] = o.PublicPort
	}
	if !isNil(o.LocalPort) {
		toSerialize["localPort"] = o.LocalPort
	}
	if !isNil(o.AllowedIps) {
		toSerialize["allowedIps"] = o.AllowedIps
	}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20018Rules struct {
	value *InlineResponse20018Rules
	isSet bool
}

func (v NullableInlineResponse20018Rules) Get() *InlineResponse20018Rules {
	return v.value
}

func (v *NullableInlineResponse20018Rules) Set(val *InlineResponse20018Rules) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20018Rules) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20018Rules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20018Rules(val *InlineResponse20018Rules) *NullableInlineResponse20018Rules {
	return &NullableInlineResponse20018Rules{value: val, isSet: true}
}

func (v NullableInlineResponse20018Rules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20018Rules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


