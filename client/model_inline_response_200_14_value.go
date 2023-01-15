/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20014Value Value of traffic filter
type InlineResponse20014Value struct {
	// Protocol of 'custom' type traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'
	Protocol *string `json:"protocol,omitempty"`
	Source InlineResponse20014ValueSource `json:"source"`
	Destination InlineResponse20014ValueDestination `json:"destination"`
}

// NewInlineResponse20014Value instantiates a new InlineResponse20014Value object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20014Value(source InlineResponse20014ValueSource, destination InlineResponse20014ValueDestination) *InlineResponse20014Value {
	this := InlineResponse20014Value{}
	this.Source = source
	this.Destination = destination
	return &this
}

// NewInlineResponse20014ValueWithDefaults instantiates a new InlineResponse20014Value object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20014ValueWithDefaults() *InlineResponse20014Value {
	this := InlineResponse20014Value{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *InlineResponse20014Value) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014Value) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *InlineResponse20014Value) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *InlineResponse20014Value) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSource returns the Source field value
func (o *InlineResponse20014Value) GetSource() InlineResponse20014ValueSource {
	if o == nil {
		var ret InlineResponse20014ValueSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20014Value) GetSourceOk() (*InlineResponse20014ValueSource, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *InlineResponse20014Value) SetSource(v InlineResponse20014ValueSource) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *InlineResponse20014Value) GetDestination() InlineResponse20014ValueDestination {
	if o == nil {
		var ret InlineResponse20014ValueDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20014Value) GetDestinationOk() (*InlineResponse20014ValueDestination, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *InlineResponse20014Value) SetDestination(v InlineResponse20014ValueDestination) {
	o.Destination = v
}

func (o InlineResponse20014Value) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20014Value struct {
	value *InlineResponse20014Value
	isSet bool
}

func (v NullableInlineResponse20014Value) Get() *InlineResponse20014Value {
	return v.value
}

func (v *NullableInlineResponse20014Value) Set(val *InlineResponse20014Value) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20014Value) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20014Value) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20014Value(val *InlineResponse20014Value) *NullableInlineResponse20014Value {
	return &NullableInlineResponse20014Value{value: val, isSet: true}
}

func (v NullableInlineResponse20014Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20014Value) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


