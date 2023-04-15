/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20044 struct for InlineResponse20044
type InlineResponse20044 struct {
	// The Meraki Ids of the set of devices.
	Ids []string `json:"ids,omitempty"`
	// The network to which the devices was moved.
	NewNetwork *string `json:"newNetwork,omitempty"`
}

// NewInlineResponse20044 instantiates a new InlineResponse20044 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20044() *InlineResponse20044 {
	this := InlineResponse20044{}
	return &this
}

// NewInlineResponse20044WithDefaults instantiates a new InlineResponse20044 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20044WithDefaults() *InlineResponse20044 {
	this := InlineResponse20044{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InlineResponse20044) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InlineResponse20044) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *InlineResponse20044) SetIds(v []string) {
	o.Ids = v
}

// GetNewNetwork returns the NewNetwork field value if set, zero value otherwise.
func (o *InlineResponse20044) GetNewNetwork() string {
	if o == nil || isNil(o.NewNetwork) {
		var ret string
		return ret
	}
	return *o.NewNetwork
}

// GetNewNetworkOk returns a tuple with the NewNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044) GetNewNetworkOk() (*string, bool) {
	if o == nil || isNil(o.NewNetwork) {
    return nil, false
	}
	return o.NewNetwork, true
}

// HasNewNetwork returns a boolean if a field has been set.
func (o *InlineResponse20044) HasNewNetwork() bool {
	if o != nil && !isNil(o.NewNetwork) {
		return true
	}

	return false
}

// SetNewNetwork gets a reference to the given string and assigns it to the NewNetwork field.
func (o *InlineResponse20044) SetNewNetwork(v string) {
	o.NewNetwork = &v
}

func (o InlineResponse20044) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !isNil(o.NewNetwork) {
		toSerialize["newNetwork"] = o.NewNetwork
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20044 struct {
	value *InlineResponse20044
	isSet bool
}

func (v NullableInlineResponse20044) Get() *InlineResponse20044 {
	return v.value
}

func (v *NullableInlineResponse20044) Set(val *InlineResponse20044) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20044) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20044) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20044(val *InlineResponse20044) *NullableInlineResponse20044 {
	return &NullableInlineResponse20044{value: val, isSet: true}
}

func (v NullableInlineResponse20044) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20044) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


