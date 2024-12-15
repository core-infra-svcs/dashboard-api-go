/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200115 struct for InlineResponse200115
type InlineResponse200115 struct {
	Device *NetworksNetworkIdSensorRelationshipsDevice `json:"device,omitempty"`
	Relationships *NetworksNetworkIdSensorRelationshipsRelationships `json:"relationships,omitempty"`
}

// NewInlineResponse200115 instantiates a new InlineResponse200115 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200115() *InlineResponse200115 {
	this := InlineResponse200115{}
	return &this
}

// NewInlineResponse200115WithDefaults instantiates a new InlineResponse200115 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200115WithDefaults() *InlineResponse200115 {
	this := InlineResponse200115{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InlineResponse200115) GetDevice() NetworksNetworkIdSensorRelationshipsDevice {
	if o == nil || isNil(o.Device) {
		var ret NetworksNetworkIdSensorRelationshipsDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200115) GetDeviceOk() (*NetworksNetworkIdSensorRelationshipsDevice, bool) {
	if o == nil || isNil(o.Device) {
    return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InlineResponse200115) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NetworksNetworkIdSensorRelationshipsDevice and assigns it to the Device field.
func (o *InlineResponse200115) SetDevice(v NetworksNetworkIdSensorRelationshipsDevice) {
	o.Device = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InlineResponse200115) GetRelationships() NetworksNetworkIdSensorRelationshipsRelationships {
	if o == nil || isNil(o.Relationships) {
		var ret NetworksNetworkIdSensorRelationshipsRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200115) GetRelationshipsOk() (*NetworksNetworkIdSensorRelationshipsRelationships, bool) {
	if o == nil || isNil(o.Relationships) {
    return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InlineResponse200115) HasRelationships() bool {
	if o != nil && !isNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given NetworksNetworkIdSensorRelationshipsRelationships and assigns it to the Relationships field.
func (o *InlineResponse200115) SetRelationships(v NetworksNetworkIdSensorRelationshipsRelationships) {
	o.Relationships = &v
}

func (o InlineResponse200115) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !isNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200115 struct {
	value *InlineResponse200115
	isSet bool
}

func (v NullableInlineResponse200115) Get() *InlineResponse200115 {
	return v.value
}

func (v *NullableInlineResponse200115) Set(val *InlineResponse200115) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200115) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200115) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200115(val *InlineResponse200115) *NullableInlineResponse200115 {
	return &NullableInlineResponse200115{value: val, isSet: true}
}

func (v NullableInlineResponse200115) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200115) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


