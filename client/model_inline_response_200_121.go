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

// InlineResponse200121 struct for InlineResponse200121
type InlineResponse200121 struct {
	// The Meraki Ids of the set of devices.
	Ids []string `json:"ids,omitempty"`
	// The network to which the devices was moved.
	NewNetwork *string `json:"newNetwork,omitempty"`
}

// NewInlineResponse200121 instantiates a new InlineResponse200121 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200121() *InlineResponse200121 {
	this := InlineResponse200121{}
	return &this
}

// NewInlineResponse200121WithDefaults instantiates a new InlineResponse200121 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200121WithDefaults() *InlineResponse200121 {
	this := InlineResponse200121{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InlineResponse200121) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InlineResponse200121) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *InlineResponse200121) SetIds(v []string) {
	o.Ids = v
}

// GetNewNetwork returns the NewNetwork field value if set, zero value otherwise.
func (o *InlineResponse200121) GetNewNetwork() string {
	if o == nil || isNil(o.NewNetwork) {
		var ret string
		return ret
	}
	return *o.NewNetwork
}

// GetNewNetworkOk returns a tuple with the NewNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121) GetNewNetworkOk() (*string, bool) {
	if o == nil || isNil(o.NewNetwork) {
    return nil, false
	}
	return o.NewNetwork, true
}

// HasNewNetwork returns a boolean if a field has been set.
func (o *InlineResponse200121) HasNewNetwork() bool {
	if o != nil && !isNil(o.NewNetwork) {
		return true
	}

	return false
}

// SetNewNetwork gets a reference to the given string and assigns it to the NewNetwork field.
func (o *InlineResponse200121) SetNewNetwork(v string) {
	o.NewNetwork = &v
}

func (o InlineResponse200121) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !isNil(o.NewNetwork) {
		toSerialize["newNetwork"] = o.NewNetwork
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200121 struct {
	value *InlineResponse200121
	isSet bool
}

func (v NullableInlineResponse200121) Get() *InlineResponse200121 {
	return v.value
}

func (v *NullableInlineResponse200121) Set(val *InlineResponse200121) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200121) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200121) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200121(val *InlineResponse200121) *NullableInlineResponse200121 {
	return &NullableInlineResponse200121{value: val, isSet: true}
}

func (v NullableInlineResponse200121) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200121) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


