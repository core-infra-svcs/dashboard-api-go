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

// InlineResponse200240 struct for InlineResponse200240
type InlineResponse200240 struct {
	// The network id of the camera
	NetworkId *string `json:"networkId,omitempty"`
	// The serial number of the camera
	Serial *string `json:"serial,omitempty"`
	Boundaries *OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries `json:"boundaries,omitempty"`
}

// NewInlineResponse200240 instantiates a new InlineResponse200240 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200240() *InlineResponse200240 {
	this := InlineResponse200240{}
	return &this
}

// NewInlineResponse200240WithDefaults instantiates a new InlineResponse200240 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200240WithDefaults() *InlineResponse200240 {
	this := InlineResponse200240{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200240) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200240) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200240) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200240) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200240) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200240) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200240) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200240) SetSerial(v string) {
	o.Serial = &v
}

// GetBoundaries returns the Boundaries field value if set, zero value otherwise.
func (o *InlineResponse200240) GetBoundaries() OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries {
	if o == nil || isNil(o.Boundaries) {
		var ret OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries
		return ret
	}
	return *o.Boundaries
}

// GetBoundariesOk returns a tuple with the Boundaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200240) GetBoundariesOk() (*OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries, bool) {
	if o == nil || isNil(o.Boundaries) {
    return nil, false
	}
	return o.Boundaries, true
}

// HasBoundaries returns a boolean if a field has been set.
func (o *InlineResponse200240) HasBoundaries() bool {
	if o != nil && !isNil(o.Boundaries) {
		return true
	}

	return false
}

// SetBoundaries gets a reference to the given OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries and assigns it to the Boundaries field.
func (o *InlineResponse200240) SetBoundaries(v OrganizationsOrganizationIdCameraBoundariesLinesByDeviceBoundaries) {
	o.Boundaries = &v
}

func (o InlineResponse200240) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Boundaries) {
		toSerialize["boundaries"] = o.Boundaries
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200240 struct {
	value *InlineResponse200240
	isSet bool
}

func (v NullableInlineResponse200240) Get() *InlineResponse200240 {
	return v.value
}

func (v *NullableInlineResponse200240) Set(val *InlineResponse200240) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200240) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200240) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200240(val *InlineResponse200240) *NullableInlineResponse200240 {
	return &NullableInlineResponse200240{value: val, isSet: true}
}

func (v NullableInlineResponse200240) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200240) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


