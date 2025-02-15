/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200359Counts Wireless LAN controller client and access point counts
type InlineResponse200359Counts struct {
	Clients *InlineResponse200359CountsClients `json:"clients,omitempty"`
	Connections *InlineResponse200359CountsConnections `json:"connections,omitempty"`
}

// NewInlineResponse200359Counts instantiates a new InlineResponse200359Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200359Counts() *InlineResponse200359Counts {
	this := InlineResponse200359Counts{}
	return &this
}

// NewInlineResponse200359CountsWithDefaults instantiates a new InlineResponse200359Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200359CountsWithDefaults() *InlineResponse200359Counts {
	this := InlineResponse200359Counts{}
	return &this
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *InlineResponse200359Counts) GetClients() InlineResponse200359CountsClients {
	if o == nil || isNil(o.Clients) {
		var ret InlineResponse200359CountsClients
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200359Counts) GetClientsOk() (*InlineResponse200359CountsClients, bool) {
	if o == nil || isNil(o.Clients) {
    return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *InlineResponse200359Counts) HasClients() bool {
	if o != nil && !isNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given InlineResponse200359CountsClients and assigns it to the Clients field.
func (o *InlineResponse200359Counts) SetClients(v InlineResponse200359CountsClients) {
	o.Clients = &v
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *InlineResponse200359Counts) GetConnections() InlineResponse200359CountsConnections {
	if o == nil || isNil(o.Connections) {
		var ret InlineResponse200359CountsConnections
		return ret
	}
	return *o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200359Counts) GetConnectionsOk() (*InlineResponse200359CountsConnections, bool) {
	if o == nil || isNil(o.Connections) {
    return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *InlineResponse200359Counts) HasConnections() bool {
	if o != nil && !isNil(o.Connections) {
		return true
	}

	return false
}

// SetConnections gets a reference to the given InlineResponse200359CountsConnections and assigns it to the Connections field.
func (o *InlineResponse200359Counts) SetConnections(v InlineResponse200359CountsConnections) {
	o.Connections = &v
}

func (o InlineResponse200359Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	if !isNil(o.Connections) {
		toSerialize["connections"] = o.Connections
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200359Counts struct {
	value *InlineResponse200359Counts
	isSet bool
}

func (v NullableInlineResponse200359Counts) Get() *InlineResponse200359Counts {
	return v.value
}

func (v *NullableInlineResponse200359Counts) Set(val *InlineResponse200359Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200359Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200359Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200359Counts(val *InlineResponse200359Counts) *NullableInlineResponse200359Counts {
	return &NullableInlineResponse200359Counts{value: val, isSet: true}
}

func (v NullableInlineResponse200359Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200359Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


