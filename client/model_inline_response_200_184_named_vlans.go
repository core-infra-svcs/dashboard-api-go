/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200184NamedVlans Named VLAN settings for wireless networks.
type InlineResponse200184NamedVlans struct {
	PoolDhcpMonitoring *InlineResponse200184NamedVlansPoolDhcpMonitoring `json:"poolDhcpMonitoring,omitempty"`
}

// NewInlineResponse200184NamedVlans instantiates a new InlineResponse200184NamedVlans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200184NamedVlans() *InlineResponse200184NamedVlans {
	this := InlineResponse200184NamedVlans{}
	return &this
}

// NewInlineResponse200184NamedVlansWithDefaults instantiates a new InlineResponse200184NamedVlans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200184NamedVlansWithDefaults() *InlineResponse200184NamedVlans {
	this := InlineResponse200184NamedVlans{}
	return &this
}

// GetPoolDhcpMonitoring returns the PoolDhcpMonitoring field value if set, zero value otherwise.
func (o *InlineResponse200184NamedVlans) GetPoolDhcpMonitoring() InlineResponse200184NamedVlansPoolDhcpMonitoring {
	if o == nil || isNil(o.PoolDhcpMonitoring) {
		var ret InlineResponse200184NamedVlansPoolDhcpMonitoring
		return ret
	}
	return *o.PoolDhcpMonitoring
}

// GetPoolDhcpMonitoringOk returns a tuple with the PoolDhcpMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200184NamedVlans) GetPoolDhcpMonitoringOk() (*InlineResponse200184NamedVlansPoolDhcpMonitoring, bool) {
	if o == nil || isNil(o.PoolDhcpMonitoring) {
    return nil, false
	}
	return o.PoolDhcpMonitoring, true
}

// HasPoolDhcpMonitoring returns a boolean if a field has been set.
func (o *InlineResponse200184NamedVlans) HasPoolDhcpMonitoring() bool {
	if o != nil && !isNil(o.PoolDhcpMonitoring) {
		return true
	}

	return false
}

// SetPoolDhcpMonitoring gets a reference to the given InlineResponse200184NamedVlansPoolDhcpMonitoring and assigns it to the PoolDhcpMonitoring field.
func (o *InlineResponse200184NamedVlans) SetPoolDhcpMonitoring(v InlineResponse200184NamedVlansPoolDhcpMonitoring) {
	o.PoolDhcpMonitoring = &v
}

func (o InlineResponse200184NamedVlans) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PoolDhcpMonitoring) {
		toSerialize["poolDhcpMonitoring"] = o.PoolDhcpMonitoring
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200184NamedVlans struct {
	value *InlineResponse200184NamedVlans
	isSet bool
}

func (v NullableInlineResponse200184NamedVlans) Get() *InlineResponse200184NamedVlans {
	return v.value
}

func (v *NullableInlineResponse200184NamedVlans) Set(val *InlineResponse200184NamedVlans) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200184NamedVlans) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200184NamedVlans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200184NamedVlans(val *InlineResponse200184NamedVlans) *NullableInlineResponse200184NamedVlans {
	return &NullableInlineResponse200184NamedVlans{value: val, isSet: true}
}

func (v NullableInlineResponse200184NamedVlans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200184NamedVlans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


