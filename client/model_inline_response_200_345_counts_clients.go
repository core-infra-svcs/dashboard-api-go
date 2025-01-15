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

// InlineResponse200345CountsClients Wireless LAN controller client counts
type InlineResponse200345CountsClients struct {
	ByStatus *InlineResponse200345CountsClientsByStatus `json:"byStatus,omitempty"`
}

// NewInlineResponse200345CountsClients instantiates a new InlineResponse200345CountsClients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200345CountsClients() *InlineResponse200345CountsClients {
	this := InlineResponse200345CountsClients{}
	return &this
}

// NewInlineResponse200345CountsClientsWithDefaults instantiates a new InlineResponse200345CountsClients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200345CountsClientsWithDefaults() *InlineResponse200345CountsClients {
	this := InlineResponse200345CountsClients{}
	return &this
}

// GetByStatus returns the ByStatus field value if set, zero value otherwise.
func (o *InlineResponse200345CountsClients) GetByStatus() InlineResponse200345CountsClientsByStatus {
	if o == nil || isNil(o.ByStatus) {
		var ret InlineResponse200345CountsClientsByStatus
		return ret
	}
	return *o.ByStatus
}

// GetByStatusOk returns a tuple with the ByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200345CountsClients) GetByStatusOk() (*InlineResponse200345CountsClientsByStatus, bool) {
	if o == nil || isNil(o.ByStatus) {
    return nil, false
	}
	return o.ByStatus, true
}

// HasByStatus returns a boolean if a field has been set.
func (o *InlineResponse200345CountsClients) HasByStatus() bool {
	if o != nil && !isNil(o.ByStatus) {
		return true
	}

	return false
}

// SetByStatus gets a reference to the given InlineResponse200345CountsClientsByStatus and assigns it to the ByStatus field.
func (o *InlineResponse200345CountsClients) SetByStatus(v InlineResponse200345CountsClientsByStatus) {
	o.ByStatus = &v
}

func (o InlineResponse200345CountsClients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByStatus) {
		toSerialize["byStatus"] = o.ByStatus
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200345CountsClients struct {
	value *InlineResponse200345CountsClients
	isSet bool
}

func (v NullableInlineResponse200345CountsClients) Get() *InlineResponse200345CountsClients {
	return v.value
}

func (v *NullableInlineResponse200345CountsClients) Set(val *InlineResponse200345CountsClients) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200345CountsClients) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200345CountsClients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200345CountsClients(val *InlineResponse200345CountsClients) *NullableInlineResponse200345CountsClients {
	return &NullableInlineResponse200345CountsClients{value: val, isSet: true}
}

func (v NullableInlineResponse200345CountsClients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200345CountsClients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

