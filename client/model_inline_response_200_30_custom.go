/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20030Custom struct for InlineResponse20030Custom
type InlineResponse20030Custom struct {
	// Protocol.
	Protocol string `json:"protocol"`
	// Destination address; hostname required for DNS, IPv4 otherwise.
	Destination string `json:"destination"`
	// Destination port.
	Port string `json:"port"`
}

// NewInlineResponse20030Custom instantiates a new InlineResponse20030Custom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20030Custom(protocol string, destination string, port string) *InlineResponse20030Custom {
	this := InlineResponse20030Custom{}
	this.Protocol = protocol
	this.Destination = destination
	this.Port = port
	return &this
}

// NewInlineResponse20030CustomWithDefaults instantiates a new InlineResponse20030Custom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20030CustomWithDefaults() *InlineResponse20030Custom {
	this := InlineResponse20030Custom{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *InlineResponse20030Custom) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030Custom) GetProtocolOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *InlineResponse20030Custom) SetProtocol(v string) {
	o.Protocol = v
}

// GetDestination returns the Destination field value
func (o *InlineResponse20030Custom) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030Custom) GetDestinationOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *InlineResponse20030Custom) SetDestination(v string) {
	o.Destination = v
}

// GetPort returns the Port field value
func (o *InlineResponse20030Custom) GetPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030Custom) GetPortOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *InlineResponse20030Custom) SetPort(v string) {
	o.Port = v
}

func (o InlineResponse20030Custom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["protocol"] = o.Protocol
	}
	if true {
		toSerialize["destination"] = o.Destination
	}
	if true {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20030Custom struct {
	value *InlineResponse20030Custom
	isSet bool
}

func (v NullableInlineResponse20030Custom) Get() *InlineResponse20030Custom {
	return v.value
}

func (v *NullableInlineResponse20030Custom) Set(val *InlineResponse20030Custom) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20030Custom) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20030Custom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20030Custom(val *InlineResponse20030Custom) *NullableInlineResponse20030Custom {
	return &NullableInlineResponse20030Custom{value: val, isSet: true}
}

func (v NullableInlineResponse20030Custom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20030Custom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


