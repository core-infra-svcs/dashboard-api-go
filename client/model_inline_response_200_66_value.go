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

// InlineResponse20066Value Value of traffic filter
type InlineResponse20066Value struct {
	// Protocol of 'custom' type traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'
	Protocol *string `json:"protocol,omitempty"`
	Source InlineResponse20066ValueSource `json:"source"`
	Destination NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination `json:"destination"`
}

// NewInlineResponse20066Value instantiates a new InlineResponse20066Value object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20066Value(source InlineResponse20066ValueSource, destination NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) *InlineResponse20066Value {
	this := InlineResponse20066Value{}
	this.Source = source
	this.Destination = destination
	return &this
}

// NewInlineResponse20066ValueWithDefaults instantiates a new InlineResponse20066Value object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20066ValueWithDefaults() *InlineResponse20066Value {
	this := InlineResponse20066Value{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *InlineResponse20066Value) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20066Value) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *InlineResponse20066Value) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *InlineResponse20066Value) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSource returns the Source field value
func (o *InlineResponse20066Value) GetSource() InlineResponse20066ValueSource {
	if o == nil {
		var ret InlineResponse20066ValueSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20066Value) GetSourceOk() (*InlineResponse20066ValueSource, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *InlineResponse20066Value) SetSource(v InlineResponse20066ValueSource) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *InlineResponse20066Value) GetDestination() NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination {
	if o == nil {
		var ret NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20066Value) GetDestinationOk() (*NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *InlineResponse20066Value) SetDestination(v NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) {
	o.Destination = v
}

func (o InlineResponse20066Value) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20066Value struct {
	value *InlineResponse20066Value
	isSet bool
}

func (v NullableInlineResponse20066Value) Get() *InlineResponse20066Value {
	return v.value
}

func (v *NullableInlineResponse20066Value) Set(val *InlineResponse20066Value) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20066Value) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20066Value) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20066Value(val *InlineResponse20066Value) *NullableInlineResponse20066Value {
	return &NullableInlineResponse20066Value{value: val, isSet: true}
}

func (v NullableInlineResponse20066Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20066Value) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

