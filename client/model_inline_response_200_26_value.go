/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20026Value Value of traffic filter
type InlineResponse20026Value struct {
	// Protocol of 'custom' type traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'
	Protocol *string `json:"protocol,omitempty"`
	Source InlineResponse20026ValueSource `json:"source"`
	Destination InlineResponse20026ValueDestination `json:"destination"`
}

// NewInlineResponse20026Value instantiates a new InlineResponse20026Value object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20026Value(source InlineResponse20026ValueSource, destination InlineResponse20026ValueDestination) *InlineResponse20026Value {
	this := InlineResponse20026Value{}
	this.Source = source
	this.Destination = destination
	return &this
}

// NewInlineResponse20026ValueWithDefaults instantiates a new InlineResponse20026Value object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20026ValueWithDefaults() *InlineResponse20026Value {
	this := InlineResponse20026Value{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *InlineResponse20026Value) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Value) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *InlineResponse20026Value) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *InlineResponse20026Value) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSource returns the Source field value
func (o *InlineResponse20026Value) GetSource() InlineResponse20026ValueSource {
	if o == nil {
		var ret InlineResponse20026ValueSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Value) GetSourceOk() (*InlineResponse20026ValueSource, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *InlineResponse20026Value) SetSource(v InlineResponse20026ValueSource) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *InlineResponse20026Value) GetDestination() InlineResponse20026ValueDestination {
	if o == nil {
		var ret InlineResponse20026ValueDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Value) GetDestinationOk() (*InlineResponse20026ValueDestination, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *InlineResponse20026Value) SetDestination(v InlineResponse20026ValueDestination) {
	o.Destination = v
}

func (o InlineResponse20026Value) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["destination"] = o.Destination
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20026Value struct {
	value *InlineResponse20026Value
	isSet bool
}

func (v NullableInlineResponse20026Value) Get() *InlineResponse20026Value {
	return v.value
}

func (v *NullableInlineResponse20026Value) Set(val *InlineResponse20026Value) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20026Value) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20026Value) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20026Value(val *InlineResponse20026Value) *NullableInlineResponse20026Value {
	return &NullableInlineResponse20026Value{value: val, isSet: true}
}

func (v NullableInlineResponse20026Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20026Value) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

