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

// InlineResponse200239 struct for InlineResponse200239
type InlineResponse200239 struct {
	// The network id of the camera
	NetworkId *string `json:"networkId,omitempty"`
	// The serial number of the camera
	Serial *string `json:"serial,omitempty"`
	Boundaries *OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries `json:"boundaries,omitempty"`
}

// NewInlineResponse200239 instantiates a new InlineResponse200239 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200239() *InlineResponse200239 {
	this := InlineResponse200239{}
	return &this
}

// NewInlineResponse200239WithDefaults instantiates a new InlineResponse200239 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200239WithDefaults() *InlineResponse200239 {
	this := InlineResponse200239{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200239) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200239) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200239) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200239) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200239) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200239) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200239) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200239) SetSerial(v string) {
	o.Serial = &v
}

// GetBoundaries returns the Boundaries field value if set, zero value otherwise.
func (o *InlineResponse200239) GetBoundaries() OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries {
	if o == nil || isNil(o.Boundaries) {
		var ret OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries
		return ret
	}
	return *o.Boundaries
}

// GetBoundariesOk returns a tuple with the Boundaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200239) GetBoundariesOk() (*OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries, bool) {
	if o == nil || isNil(o.Boundaries) {
    return nil, false
	}
	return o.Boundaries, true
}

// HasBoundaries returns a boolean if a field has been set.
func (o *InlineResponse200239) HasBoundaries() bool {
	if o != nil && !isNil(o.Boundaries) {
		return true
	}

	return false
}

// SetBoundaries gets a reference to the given OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries and assigns it to the Boundaries field.
func (o *InlineResponse200239) SetBoundaries(v OrganizationsOrganizationIdCameraBoundariesAreasByDeviceBoundaries) {
	o.Boundaries = &v
}

func (o InlineResponse200239) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200239 struct {
	value *InlineResponse200239
	isSet bool
}

func (v NullableInlineResponse200239) Get() *InlineResponse200239 {
	return v.value
}

func (v *NullableInlineResponse200239) Set(val *InlineResponse200239) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200239) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200239) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200239(val *InlineResponse200239) *NullableInlineResponse200239 {
	return &NullableInlineResponse200239{value: val, isSet: true}
}

func (v NullableInlineResponse200239) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200239) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


