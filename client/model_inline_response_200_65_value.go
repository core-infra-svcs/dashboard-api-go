/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20065Value Value of traffic filter
type InlineResponse20065Value struct {
	// Protocol of 'custom' type traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'
	Protocol *string `json:"protocol,omitempty"`
	Source InlineResponse20065ValueSource `json:"source"`
	Destination NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination `json:"destination"`
}

// NewInlineResponse20065Value instantiates a new InlineResponse20065Value object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20065Value(source InlineResponse20065ValueSource, destination NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) *InlineResponse20065Value {
	this := InlineResponse20065Value{}
	this.Source = source
	this.Destination = destination
	return &this
}

// NewInlineResponse20065ValueWithDefaults instantiates a new InlineResponse20065Value object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20065ValueWithDefaults() *InlineResponse20065Value {
	this := InlineResponse20065Value{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *InlineResponse20065Value) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065Value) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *InlineResponse20065Value) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *InlineResponse20065Value) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSource returns the Source field value
func (o *InlineResponse20065Value) GetSource() InlineResponse20065ValueSource {
	if o == nil {
		var ret InlineResponse20065ValueSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20065Value) GetSourceOk() (*InlineResponse20065ValueSource, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *InlineResponse20065Value) SetSource(v InlineResponse20065ValueSource) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *InlineResponse20065Value) GetDestination() NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination {
	if o == nil {
		var ret NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20065Value) GetDestinationOk() (*NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *InlineResponse20065Value) SetDestination(v NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) {
	o.Destination = v
}

func (o InlineResponse20065Value) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20065Value struct {
	value *InlineResponse20065Value
	isSet bool
}

func (v NullableInlineResponse20065Value) Get() *InlineResponse20065Value {
	return v.value
}

func (v *NullableInlineResponse20065Value) Set(val *InlineResponse20065Value) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20065Value) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20065Value) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20065Value(val *InlineResponse20065Value) *NullableInlineResponse20065Value {
	return &NullableInlineResponse20065Value{value: val, isSet: true}
}

func (v NullableInlineResponse20065Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20065Value) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


