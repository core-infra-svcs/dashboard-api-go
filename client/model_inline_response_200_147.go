/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200147 struct for InlineResponse200147
type InlineResponse200147 struct {
	// The ID for the link aggregation.
	Id *string `json:"id,omitempty"`
	// The ID for the link aggregation.
	SwitchPorts []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts `json:"switchPorts,omitempty"`
}

// NewInlineResponse200147 instantiates a new InlineResponse200147 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200147() *InlineResponse200147 {
	this := InlineResponse200147{}
	return &this
}

// NewInlineResponse200147WithDefaults instantiates a new InlineResponse200147 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200147WithDefaults() *InlineResponse200147 {
	this := InlineResponse200147{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200147) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200147) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200147) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200147) SetId(v string) {
	o.Id = &v
}

// GetSwitchPorts returns the SwitchPorts field value if set, zero value otherwise.
func (o *InlineResponse200147) GetSwitchPorts() []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts {
	if o == nil || isNil(o.SwitchPorts) {
		var ret []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts
		return ret
	}
	return o.SwitchPorts
}

// GetSwitchPortsOk returns a tuple with the SwitchPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200147) GetSwitchPortsOk() ([]NetworksNetworkIdSwitchLinkAggregationsSwitchPorts, bool) {
	if o == nil || isNil(o.SwitchPorts) {
    return nil, false
	}
	return o.SwitchPorts, true
}

// HasSwitchPorts returns a boolean if a field has been set.
func (o *InlineResponse200147) HasSwitchPorts() bool {
	if o != nil && !isNil(o.SwitchPorts) {
		return true
	}

	return false
}

// SetSwitchPorts gets a reference to the given []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts and assigns it to the SwitchPorts field.
func (o *InlineResponse200147) SetSwitchPorts(v []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) {
	o.SwitchPorts = v
}

func (o InlineResponse200147) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.SwitchPorts) {
		toSerialize["switchPorts"] = o.SwitchPorts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200147 struct {
	value *InlineResponse200147
	isSet bool
}

func (v NullableInlineResponse200147) Get() *InlineResponse200147 {
	return v.value
}

func (v *NullableInlineResponse200147) Set(val *InlineResponse200147) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200147) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200147) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200147(val *InlineResponse200147) *NullableInlineResponse200147 {
	return &NullableInlineResponse200147{value: val, isSet: true}
}

func (v NullableInlineResponse200147) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200147) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


