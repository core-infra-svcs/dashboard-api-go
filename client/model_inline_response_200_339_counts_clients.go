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

// InlineResponse200339CountsClients Wireless LAN controller client counts
type InlineResponse200339CountsClients struct {
	ByStatus *InlineResponse200339CountsClientsByStatus `json:"byStatus,omitempty"`
}

// NewInlineResponse200339CountsClients instantiates a new InlineResponse200339CountsClients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200339CountsClients() *InlineResponse200339CountsClients {
	this := InlineResponse200339CountsClients{}
	return &this
}

// NewInlineResponse200339CountsClientsWithDefaults instantiates a new InlineResponse200339CountsClients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200339CountsClientsWithDefaults() *InlineResponse200339CountsClients {
	this := InlineResponse200339CountsClients{}
	return &this
}

// GetByStatus returns the ByStatus field value if set, zero value otherwise.
func (o *InlineResponse200339CountsClients) GetByStatus() InlineResponse200339CountsClientsByStatus {
	if o == nil || isNil(o.ByStatus) {
		var ret InlineResponse200339CountsClientsByStatus
		return ret
	}
	return *o.ByStatus
}

// GetByStatusOk returns a tuple with the ByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200339CountsClients) GetByStatusOk() (*InlineResponse200339CountsClientsByStatus, bool) {
	if o == nil || isNil(o.ByStatus) {
    return nil, false
	}
	return o.ByStatus, true
}

// HasByStatus returns a boolean if a field has been set.
func (o *InlineResponse200339CountsClients) HasByStatus() bool {
	if o != nil && !isNil(o.ByStatus) {
		return true
	}

	return false
}

// SetByStatus gets a reference to the given InlineResponse200339CountsClientsByStatus and assigns it to the ByStatus field.
func (o *InlineResponse200339CountsClients) SetByStatus(v InlineResponse200339CountsClientsByStatus) {
	o.ByStatus = &v
}

func (o InlineResponse200339CountsClients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByStatus) {
		toSerialize["byStatus"] = o.ByStatus
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200339CountsClients struct {
	value *InlineResponse200339CountsClients
	isSet bool
}

func (v NullableInlineResponse200339CountsClients) Get() *InlineResponse200339CountsClients {
	return v.value
}

func (v *NullableInlineResponse200339CountsClients) Set(val *InlineResponse200339CountsClients) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200339CountsClients) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200339CountsClients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200339CountsClients(val *InlineResponse200339CountsClients) *NullableInlineResponse200339CountsClients {
	return &NullableInlineResponse200339CountsClients{value: val, isSet: true}
}

func (v NullableInlineResponse200339CountsClients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200339CountsClients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


