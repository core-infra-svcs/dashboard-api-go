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

// InlineResponse200109 struct for InlineResponse200109
type InlineResponse200109 struct {
	Device *NetworksNetworkIdSensorRelationshipsDevice `json:"device,omitempty"`
	Relationships *NetworksNetworkIdSensorRelationshipsRelationships `json:"relationships,omitempty"`
}

// NewInlineResponse200109 instantiates a new InlineResponse200109 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200109() *InlineResponse200109 {
	this := InlineResponse200109{}
	return &this
}

// NewInlineResponse200109WithDefaults instantiates a new InlineResponse200109 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200109WithDefaults() *InlineResponse200109 {
	this := InlineResponse200109{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InlineResponse200109) GetDevice() NetworksNetworkIdSensorRelationshipsDevice {
	if o == nil || isNil(o.Device) {
		var ret NetworksNetworkIdSensorRelationshipsDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200109) GetDeviceOk() (*NetworksNetworkIdSensorRelationshipsDevice, bool) {
	if o == nil || isNil(o.Device) {
    return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InlineResponse200109) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NetworksNetworkIdSensorRelationshipsDevice and assigns it to the Device field.
func (o *InlineResponse200109) SetDevice(v NetworksNetworkIdSensorRelationshipsDevice) {
	o.Device = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InlineResponse200109) GetRelationships() NetworksNetworkIdSensorRelationshipsRelationships {
	if o == nil || isNil(o.Relationships) {
		var ret NetworksNetworkIdSensorRelationshipsRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200109) GetRelationshipsOk() (*NetworksNetworkIdSensorRelationshipsRelationships, bool) {
	if o == nil || isNil(o.Relationships) {
    return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InlineResponse200109) HasRelationships() bool {
	if o != nil && !isNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given NetworksNetworkIdSensorRelationshipsRelationships and assigns it to the Relationships field.
func (o *InlineResponse200109) SetRelationships(v NetworksNetworkIdSensorRelationshipsRelationships) {
	o.Relationships = &v
}

func (o InlineResponse200109) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !isNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200109 struct {
	value *InlineResponse200109
	isSet bool
}

func (v NullableInlineResponse200109) Get() *InlineResponse200109 {
	return v.value
}

func (v *NullableInlineResponse200109) Set(val *InlineResponse200109) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200109) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200109) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200109(val *InlineResponse200109) *NullableInlineResponse200109 {
	return &NullableInlineResponse200109{value: val, isSet: true}
}

func (v NullableInlineResponse200109) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200109) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


