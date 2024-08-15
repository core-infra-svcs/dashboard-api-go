/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200183 struct for InlineResponse200183
type InlineResponse200183 struct {
	// The serial number for the device.
	Serial *string `json:"serial,omitempty"`
	// List of device serials that make up the mesh.
	MeshRoute []string `json:"meshRoute,omitempty"`
	LatestMeshPerformance *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance `json:"latestMeshPerformance,omitempty"`
}

// NewInlineResponse200183 instantiates a new InlineResponse200183 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200183() *InlineResponse200183 {
	this := InlineResponse200183{}
	return &this
}

// NewInlineResponse200183WithDefaults instantiates a new InlineResponse200183 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200183WithDefaults() *InlineResponse200183 {
	this := InlineResponse200183{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200183) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200183) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200183) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200183) SetSerial(v string) {
	o.Serial = &v
}

// GetMeshRoute returns the MeshRoute field value if set, zero value otherwise.
func (o *InlineResponse200183) GetMeshRoute() []string {
	if o == nil || isNil(o.MeshRoute) {
		var ret []string
		return ret
	}
	return o.MeshRoute
}

// GetMeshRouteOk returns a tuple with the MeshRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200183) GetMeshRouteOk() ([]string, bool) {
	if o == nil || isNil(o.MeshRoute) {
    return nil, false
	}
	return o.MeshRoute, true
}

// HasMeshRoute returns a boolean if a field has been set.
func (o *InlineResponse200183) HasMeshRoute() bool {
	if o != nil && !isNil(o.MeshRoute) {
		return true
	}

	return false
}

// SetMeshRoute gets a reference to the given []string and assigns it to the MeshRoute field.
func (o *InlineResponse200183) SetMeshRoute(v []string) {
	o.MeshRoute = v
}

// GetLatestMeshPerformance returns the LatestMeshPerformance field value if set, zero value otherwise.
func (o *InlineResponse200183) GetLatestMeshPerformance() NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance {
	if o == nil || isNil(o.LatestMeshPerformance) {
		var ret NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance
		return ret
	}
	return *o.LatestMeshPerformance
}

// GetLatestMeshPerformanceOk returns a tuple with the LatestMeshPerformance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200183) GetLatestMeshPerformanceOk() (*NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance, bool) {
	if o == nil || isNil(o.LatestMeshPerformance) {
    return nil, false
	}
	return o.LatestMeshPerformance, true
}

// HasLatestMeshPerformance returns a boolean if a field has been set.
func (o *InlineResponse200183) HasLatestMeshPerformance() bool {
	if o != nil && !isNil(o.LatestMeshPerformance) {
		return true
	}

	return false
}

// SetLatestMeshPerformance gets a reference to the given NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance and assigns it to the LatestMeshPerformance field.
func (o *InlineResponse200183) SetLatestMeshPerformance(v NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) {
	o.LatestMeshPerformance = &v
}

func (o InlineResponse200183) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200183 struct {
	value *InlineResponse200183
	isSet bool
}

func (v NullableInlineResponse200183) Get() *InlineResponse200183 {
	return v.value
}

func (v *NullableInlineResponse200183) Set(val *InlineResponse200183) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200183) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200183) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200183(val *InlineResponse200183) *NullableInlineResponse200183 {
	return &NullableInlineResponse200183{value: val, isSet: true}
}

func (v NullableInlineResponse200183) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200183) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


