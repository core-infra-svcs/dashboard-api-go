/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20023ValueSource Source of 'custom' type traffic filter
type InlineResponse20023ValueSource struct {
	// E.g.: \"any\", \"0\" (also means \"any\"), \"8080\", \"1-1024\"
	Port *string `json:"port,omitempty"`
	// CIDR format address (e.g.\"192.168.10.1\", which is the same as \"192.168.10.1/32\"), or \"any\". Cannot be used in combination with the \"vlan\" property
	Cidr *string `json:"cidr,omitempty"`
	// VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the \"cidr\" property and is currently only available under a template network.
	Vlan *int32 `json:"vlan,omitempty"`
	// Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the \"vlan\" property and is currently only available under a template network.
	Host *int32 `json:"host,omitempty"`
}

// NewInlineResponse20023ValueSource instantiates a new InlineResponse20023ValueSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20023ValueSource() *InlineResponse20023ValueSource {
	this := InlineResponse20023ValueSource{}
	return &this
}

// NewInlineResponse20023ValueSourceWithDefaults instantiates a new InlineResponse20023ValueSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20023ValueSourceWithDefaults() *InlineResponse20023ValueSource {
	this := InlineResponse20023ValueSource{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *InlineResponse20023ValueSource) GetPort() string {
	if o == nil || isNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023ValueSource) GetPortOk() (*string, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *InlineResponse20023ValueSource) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *InlineResponse20023ValueSource) SetPort(v string) {
	o.Port = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *InlineResponse20023ValueSource) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023ValueSource) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *InlineResponse20023ValueSource) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *InlineResponse20023ValueSource) SetCidr(v string) {
	o.Cidr = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineResponse20023ValueSource) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023ValueSource) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineResponse20023ValueSource) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *InlineResponse20023ValueSource) SetVlan(v int32) {
	o.Vlan = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *InlineResponse20023ValueSource) GetHost() int32 {
	if o == nil || isNil(o.Host) {
		var ret int32
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023ValueSource) GetHostOk() (*int32, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *InlineResponse20023ValueSource) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given int32 and assigns it to the Host field.
func (o *InlineResponse20023ValueSource) SetHost(v int32) {
	o.Host = &v
}

func (o InlineResponse20023ValueSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20023ValueSource struct {
	value *InlineResponse20023ValueSource
	isSet bool
}

func (v NullableInlineResponse20023ValueSource) Get() *InlineResponse20023ValueSource {
	return v.value
}

func (v *NullableInlineResponse20023ValueSource) Set(val *InlineResponse20023ValueSource) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20023ValueSource) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20023ValueSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20023ValueSource(val *InlineResponse20023ValueSource) *NullableInlineResponse20023ValueSource {
	return &NullableInlineResponse20023ValueSource{value: val, isSet: true}
}

func (v NullableInlineResponse20023ValueSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20023ValueSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


