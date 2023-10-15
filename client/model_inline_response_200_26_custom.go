/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20026Custom struct for InlineResponse20026Custom
type InlineResponse20026Custom struct {
	// Protocol.
	Protocol string `json:"protocol"`
	// Destination address; hostname required for DNS, IPv4 otherwise.
	Destination string `json:"destination"`
	// Destination port.
	Port string `json:"port"`
}

// NewInlineResponse20026Custom instantiates a new InlineResponse20026Custom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20026Custom(protocol string, destination string, port string) *InlineResponse20026Custom {
	this := InlineResponse20026Custom{}
	this.Protocol = protocol
	this.Destination = destination
	this.Port = port
	return &this
}

// NewInlineResponse20026CustomWithDefaults instantiates a new InlineResponse20026Custom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20026CustomWithDefaults() *InlineResponse20026Custom {
	this := InlineResponse20026Custom{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *InlineResponse20026Custom) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Custom) GetProtocolOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *InlineResponse20026Custom) SetProtocol(v string) {
	o.Protocol = v
}

// GetDestination returns the Destination field value
func (o *InlineResponse20026Custom) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Custom) GetDestinationOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *InlineResponse20026Custom) SetDestination(v string) {
	o.Destination = v
}

// GetPort returns the Port field value
func (o *InlineResponse20026Custom) GetPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Custom) GetPortOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *InlineResponse20026Custom) SetPort(v string) {
	o.Port = v
}

func (o InlineResponse20026Custom) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20026Custom struct {
	value *InlineResponse20026Custom
	isSet bool
}

func (v NullableInlineResponse20026Custom) Get() *InlineResponse20026Custom {
	return v.value
}

func (v *NullableInlineResponse20026Custom) Set(val *InlineResponse20026Custom) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20026Custom) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20026Custom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20026Custom(val *InlineResponse20026Custom) *NullableInlineResponse20026Custom {
	return &NullableInlineResponse20026Custom{value: val, isSet: true}
}

func (v NullableInlineResponse20026Custom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20026Custom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


