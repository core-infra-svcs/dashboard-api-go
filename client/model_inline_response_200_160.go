/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200160 struct for InlineResponse200160
type InlineResponse200160 struct {
	// Switch port schedule ID
	Id *string `json:"id,omitempty"`
	// Network ID
	NetworkId *string `json:"networkId,omitempty"`
	// Switch port schedule name
	Name *string `json:"name,omitempty"`
	PortSchedule *NetworksNetworkIdSwitchPortSchedulesPortSchedule `json:"portSchedule,omitempty"`
}

// NewInlineResponse200160 instantiates a new InlineResponse200160 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200160() *InlineResponse200160 {
	this := InlineResponse200160{}
	return &this
}

// NewInlineResponse200160WithDefaults instantiates a new InlineResponse200160 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200160WithDefaults() *InlineResponse200160 {
	this := InlineResponse200160{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200160) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200160) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200160) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200160) SetId(v string) {
	o.Id = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200160) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200160) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200160) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200160) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200160) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200160) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200160) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200160) SetName(v string) {
	o.Name = &v
}

// GetPortSchedule returns the PortSchedule field value if set, zero value otherwise.
func (o *InlineResponse200160) GetPortSchedule() NetworksNetworkIdSwitchPortSchedulesPortSchedule {
	if o == nil || isNil(o.PortSchedule) {
		var ret NetworksNetworkIdSwitchPortSchedulesPortSchedule
		return ret
	}
	return *o.PortSchedule
}

// GetPortScheduleOk returns a tuple with the PortSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200160) GetPortScheduleOk() (*NetworksNetworkIdSwitchPortSchedulesPortSchedule, bool) {
	if o == nil || isNil(o.PortSchedule) {
    return nil, false
	}
	return o.PortSchedule, true
}

// HasPortSchedule returns a boolean if a field has been set.
func (o *InlineResponse200160) HasPortSchedule() bool {
	if o != nil && !isNil(o.PortSchedule) {
		return true
	}

	return false
}

// SetPortSchedule gets a reference to the given NetworksNetworkIdSwitchPortSchedulesPortSchedule and assigns it to the PortSchedule field.
func (o *InlineResponse200160) SetPortSchedule(v NetworksNetworkIdSwitchPortSchedulesPortSchedule) {
	o.PortSchedule = &v
}

func (o InlineResponse200160) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.PortSchedule) {
		toSerialize["portSchedule"] = o.PortSchedule
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200160 struct {
	value *InlineResponse200160
	isSet bool
}

func (v NullableInlineResponse200160) Get() *InlineResponse200160 {
	return v.value
}

func (v *NullableInlineResponse200160) Set(val *InlineResponse200160) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200160) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200160) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200160(val *InlineResponse200160) *NullableInlineResponse200160 {
	return &NullableInlineResponse200160{value: val, isSet: true}
}

func (v NullableInlineResponse200160) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200160) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


