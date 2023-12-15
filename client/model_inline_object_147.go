/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject147 struct for InlineObject147
type InlineObject147 struct {
	// A list of the syslog servers for this network
	Servers []NetworksNetworkIdSyslogServersServers `json:"servers"`
}

// NewInlineObject147 instantiates a new InlineObject147 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject147(servers []NetworksNetworkIdSyslogServersServers) *InlineObject147 {
	this := InlineObject147{}
	this.Servers = servers
	return &this
}

// NewInlineObject147WithDefaults instantiates a new InlineObject147 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject147WithDefaults() *InlineObject147 {
	this := InlineObject147{}
	return &this
}

// GetServers returns the Servers field value
func (o *InlineObject147) GetServers() []NetworksNetworkIdSyslogServersServers {
	if o == nil {
		var ret []NetworksNetworkIdSyslogServersServers
		return ret
	}

	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value
// and a boolean to check if the value has been set.
func (o *InlineObject147) GetServersOk() ([]NetworksNetworkIdSyslogServersServers, bool) {
	if o == nil {
    return nil, false
	}
	return o.Servers, true
}

// SetServers sets field value
func (o *InlineObject147) SetServers(v []NetworksNetworkIdSyslogServersServers) {
	o.Servers = v
}

func (o InlineObject147) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject147 struct {
	value *InlineObject147
	isSet bool
}

func (v NullableInlineObject147) Get() *InlineObject147 {
	return v.value
}

func (v *NullableInlineObject147) Set(val *InlineObject147) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject147) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject147) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject147(val *InlineObject147) *NullableInlineObject147 {
	return &NullableInlineObject147{value: val, isSet: true}
}

func (v NullableInlineObject147) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject147) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


