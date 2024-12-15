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

// InlineResponse200189 struct for InlineResponse200189
type InlineResponse200189 struct {
	// The serial number for the device.
	Serial *string `json:"serial,omitempty"`
	// List of device serials that make up the mesh.
	MeshRoute []string `json:"meshRoute,omitempty"`
	LatestMeshPerformance *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance `json:"latestMeshPerformance,omitempty"`
}

// NewInlineResponse200189 instantiates a new InlineResponse200189 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200189() *InlineResponse200189 {
	this := InlineResponse200189{}
	return &this
}

// NewInlineResponse200189WithDefaults instantiates a new InlineResponse200189 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200189WithDefaults() *InlineResponse200189 {
	this := InlineResponse200189{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200189) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200189) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200189) SetSerial(v string) {
	o.Serial = &v
}

// GetMeshRoute returns the MeshRoute field value if set, zero value otherwise.
func (o *InlineResponse200189) GetMeshRoute() []string {
	if o == nil || isNil(o.MeshRoute) {
		var ret []string
		return ret
	}
	return o.MeshRoute
}

// GetMeshRouteOk returns a tuple with the MeshRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetMeshRouteOk() ([]string, bool) {
	if o == nil || isNil(o.MeshRoute) {
    return nil, false
	}
	return o.MeshRoute, true
}

// HasMeshRoute returns a boolean if a field has been set.
func (o *InlineResponse200189) HasMeshRoute() bool {
	if o != nil && !isNil(o.MeshRoute) {
		return true
	}

	return false
}

// SetMeshRoute gets a reference to the given []string and assigns it to the MeshRoute field.
func (o *InlineResponse200189) SetMeshRoute(v []string) {
	o.MeshRoute = v
}

// GetLatestMeshPerformance returns the LatestMeshPerformance field value if set, zero value otherwise.
func (o *InlineResponse200189) GetLatestMeshPerformance() NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance {
	if o == nil || isNil(o.LatestMeshPerformance) {
		var ret NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance
		return ret
	}
	return *o.LatestMeshPerformance
}

// GetLatestMeshPerformanceOk returns a tuple with the LatestMeshPerformance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetLatestMeshPerformanceOk() (*NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance, bool) {
	if o == nil || isNil(o.LatestMeshPerformance) {
    return nil, false
	}
	return o.LatestMeshPerformance, true
}

// HasLatestMeshPerformance returns a boolean if a field has been set.
func (o *InlineResponse200189) HasLatestMeshPerformance() bool {
	if o != nil && !isNil(o.LatestMeshPerformance) {
		return true
	}

	return false
}

// SetLatestMeshPerformance gets a reference to the given NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance and assigns it to the LatestMeshPerformance field.
func (o *InlineResponse200189) SetLatestMeshPerformance(v NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) {
	o.LatestMeshPerformance = &v
}

func (o InlineResponse200189) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.MeshRoute) {
		toSerialize["meshRoute"] = o.MeshRoute
	}
	if !isNil(o.LatestMeshPerformance) {
		toSerialize["latestMeshPerformance"] = o.LatestMeshPerformance
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200189 struct {
	value *InlineResponse200189
	isSet bool
}

func (v NullableInlineResponse200189) Get() *InlineResponse200189 {
	return v.value
}

func (v *NullableInlineResponse200189) Set(val *InlineResponse200189) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200189) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200189) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200189(val *InlineResponse200189) *NullableInlineResponse200189 {
	return &NullableInlineResponse200189{value: val, isSet: true}
}

func (v NullableInlineResponse200189) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200189) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


