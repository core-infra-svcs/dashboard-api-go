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

// InlineObject164 struct for InlineObject164
type InlineObject164 struct {
	// A list of the syslog servers for this network
	Servers []NetworksNetworkIdSyslogServersServers `json:"servers"`
}

// NewInlineObject164 instantiates a new InlineObject164 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject164(servers []NetworksNetworkIdSyslogServersServers) *InlineObject164 {
	this := InlineObject164{}
	this.Servers = servers
	return &this
}

// NewInlineObject164WithDefaults instantiates a new InlineObject164 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject164WithDefaults() *InlineObject164 {
	this := InlineObject164{}
	return &this
}

// GetServers returns the Servers field value
func (o *InlineObject164) GetServers() []NetworksNetworkIdSyslogServersServers {
	if o == nil {
		var ret []NetworksNetworkIdSyslogServersServers
		return ret
	}

	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value
// and a boolean to check if the value has been set.
func (o *InlineObject164) GetServersOk() ([]NetworksNetworkIdSyslogServersServers, bool) {
	if o == nil {
    return nil, false
	}
	return o.Servers, true
}

// SetServers sets field value
func (o *InlineObject164) SetServers(v []NetworksNetworkIdSyslogServersServers) {
	o.Servers = v
}

func (o InlineObject164) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject164 struct {
	value *InlineObject164
	isSet bool
}

func (v NullableInlineObject164) Get() *InlineObject164 {
	return v.value
}

func (v *NullableInlineObject164) Set(val *InlineObject164) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject164) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject164) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject164(val *InlineObject164) *NullableInlineObject164 {
	return &NullableInlineObject164{value: val, isSet: true}
}

func (v NullableInlineObject164) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject164) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


