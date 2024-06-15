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

// NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise Object containing the number of sensor alerts that occurred due to noise readings
type NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise struct {
	// Number of sensor alerts that occurred due to ambient noise readings
	Ambient *int32 `json:"ambient,omitempty"`
}

// NewNetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise instantiates a new NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise() *NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise {
	this := NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise{}
	return &this
}

// NewNetworksNetworkIdSensorAlertsOverviewByMetricCountsNoiseWithDefaults instantiates a new NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsOverviewByMetricCountsNoiseWithDefaults() *NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise {
	this := NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise{}
	return &this
}

// GetAmbient returns the Ambient field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise) GetAmbient() int32 {
	if o == nil || isNil(o.Ambient) {
		var ret int32
		return ret
	}
	return *o.Ambient
}

// GetAmbientOk returns a tuple with the Ambient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise) GetAmbientOk() (*int32, bool) {
	if o == nil || isNil(o.Ambient) {
    return nil, false
	}
	return o.Ambient, true
}

// HasAmbient returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise) HasAmbient() bool {
	if o != nil && !isNil(o.Ambient) {
		return true
	}

	return false
}

// SetAmbient gets a reference to the given int32 and assigns it to the Ambient field.
func (o *NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise) SetAmbient(v int32) {
	o.Ambient = &v
}

func (o NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ambient) {
		toSerialize["ambient"] = o.Ambient
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise struct {
	value *NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise) Get() *NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise) Set(val *NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise(val *NetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise) *NullableNetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise {
	return &NullableNetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsOverviewByMetricCountsNoise) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


