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

// InlineResponse200118 struct for InlineResponse200118
type InlineResponse200118 struct {
	Device *NetworksNetworkIdSensorRelationshipsDevice `json:"device,omitempty"`
	Relationships *NetworksNetworkIdSensorRelationshipsRelationships `json:"relationships,omitempty"`
}

// NewInlineResponse200118 instantiates a new InlineResponse200118 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200118() *InlineResponse200118 {
	this := InlineResponse200118{}
	return &this
}

// NewInlineResponse200118WithDefaults instantiates a new InlineResponse200118 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200118WithDefaults() *InlineResponse200118 {
	this := InlineResponse200118{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InlineResponse200118) GetDevice() NetworksNetworkIdSensorRelationshipsDevice {
	if o == nil || isNil(o.Device) {
		var ret NetworksNetworkIdSensorRelationshipsDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200118) GetDeviceOk() (*NetworksNetworkIdSensorRelationshipsDevice, bool) {
	if o == nil || isNil(o.Device) {
    return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InlineResponse200118) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NetworksNetworkIdSensorRelationshipsDevice and assigns it to the Device field.
func (o *InlineResponse200118) SetDevice(v NetworksNetworkIdSensorRelationshipsDevice) {
	o.Device = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InlineResponse200118) GetRelationships() NetworksNetworkIdSensorRelationshipsRelationships {
	if o == nil || isNil(o.Relationships) {
		var ret NetworksNetworkIdSensorRelationshipsRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200118) GetRelationshipsOk() (*NetworksNetworkIdSensorRelationshipsRelationships, bool) {
	if o == nil || isNil(o.Relationships) {
    return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InlineResponse200118) HasRelationships() bool {
	if o != nil && !isNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given NetworksNetworkIdSensorRelationshipsRelationships and assigns it to the Relationships field.
func (o *InlineResponse200118) SetRelationships(v NetworksNetworkIdSensorRelationshipsRelationships) {
	o.Relationships = &v
}

func (o InlineResponse200118) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !isNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200118 struct {
	value *InlineResponse200118
	isSet bool
}

func (v NullableInlineResponse200118) Get() *InlineResponse200118 {
	return v.value
}

func (v *NullableInlineResponse200118) Set(val *InlineResponse200118) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200118) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200118) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200118(val *InlineResponse200118) *NullableInlineResponse200118 {
	return &NullableInlineResponse200118{value: val, isSet: true}
}

func (v NullableInlineResponse200118) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200118) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


