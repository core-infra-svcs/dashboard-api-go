/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject156 struct for InlineObject156
type InlineObject156 struct {
	// A list of the syslog servers for this network
	Servers []NetworksNetworkIdSyslogServersServers `json:"servers"`
}

// NewInlineObject156 instantiates a new InlineObject156 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject156(servers []NetworksNetworkIdSyslogServersServers) *InlineObject156 {
	this := InlineObject156{}
	this.Servers = servers
	return &this
}

// NewInlineObject156WithDefaults instantiates a new InlineObject156 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject156WithDefaults() *InlineObject156 {
	this := InlineObject156{}
	return &this
}

// GetServers returns the Servers field value
func (o *InlineObject156) GetServers() []NetworksNetworkIdSyslogServersServers {
	if o == nil {
		var ret []NetworksNetworkIdSyslogServersServers
		return ret
	}

	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value
// and a boolean to check if the value has been set.
func (o *InlineObject156) GetServersOk() ([]NetworksNetworkIdSyslogServersServers, bool) {
	if o == nil {
    return nil, false
	}
	return o.Servers, true
}

// SetServers sets field value
func (o *InlineObject156) SetServers(v []NetworksNetworkIdSyslogServersServers) {
	o.Servers = v
}

func (o InlineObject156) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject156 struct {
	value *InlineObject156
	isSet bool
}

func (v NullableInlineObject156) Get() *InlineObject156 {
	return v.value
}

func (v *NullableInlineObject156) Set(val *InlineObject156) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject156) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject156) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject156(val *InlineObject156) *NullableInlineObject156 {
	return &NullableInlineObject156{value: val, isSet: true}
}

func (v NullableInlineObject156) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject156) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


