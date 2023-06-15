/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20022ValueDestination Destination of 'custom' type traffic filter
type InlineResponse20022ValueDestination struct {
	// E.g.: \"any\", \"0\" (also means \"any\"), \"8080\", \"1-1024\"
	Port *string `json:"port,omitempty"`
	// CIDR format address (e.g.\"192.168.10.1\", which is the same as \"192.168.10.1/32\"), or \"any\"
	Cidr *string `json:"cidr,omitempty"`
}

// NewInlineResponse20022ValueDestination instantiates a new InlineResponse20022ValueDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20022ValueDestination() *InlineResponse20022ValueDestination {
	this := InlineResponse20022ValueDestination{}
	return &this
}

// NewInlineResponse20022ValueDestinationWithDefaults instantiates a new InlineResponse20022ValueDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20022ValueDestinationWithDefaults() *InlineResponse20022ValueDestination {
	this := InlineResponse20022ValueDestination{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *InlineResponse20022ValueDestination) GetPort() string {
	if o == nil || isNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20022ValueDestination) GetPortOk() (*string, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *InlineResponse20022ValueDestination) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *InlineResponse20022ValueDestination) SetPort(v string) {
	o.Port = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *InlineResponse20022ValueDestination) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20022ValueDestination) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *InlineResponse20022ValueDestination) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *InlineResponse20022ValueDestination) SetCidr(v string) {
	o.Cidr = &v
}

func (o InlineResponse20022ValueDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20022ValueDestination struct {
	value *InlineResponse20022ValueDestination
	isSet bool
}

func (v NullableInlineResponse20022ValueDestination) Get() *InlineResponse20022ValueDestination {
	return v.value
}

func (v *NullableInlineResponse20022ValueDestination) Set(val *InlineResponse20022ValueDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20022ValueDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20022ValueDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20022ValueDestination(val *InlineResponse20022ValueDestination) *NullableInlineResponse20022ValueDestination {
	return &NullableInlineResponse20022ValueDestination{value: val, isSet: true}
}

func (v NullableInlineResponse20022ValueDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20022ValueDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

