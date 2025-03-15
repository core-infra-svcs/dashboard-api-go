/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSensorAlertsProfilesConditions struct for NetworksNetworkIdSensorAlertsProfilesConditions
type NetworksNetworkIdSensorAlertsProfilesConditions struct {
	// The type of sensor metric that will be monitored for changes.
	Metric string `json:"metric"`
	Threshold NetworksNetworkIdSensorAlertsProfilesThreshold `json:"threshold"`
	// If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature, humidity, realPower, apparentPower, powerFactor, voltage, current, and frequency thresholds.
	Direction *string `json:"direction,omitempty"`
	// Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, and 8 hours. Default is 0.
	Duration *int32 `json:"duration,omitempty"`
}

// NewNetworksNetworkIdSensorAlertsProfilesConditions instantiates a new NetworksNetworkIdSensorAlertsProfilesConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesConditions(metric string, threshold NetworksNetworkIdSensorAlertsProfilesThreshold) *NetworksNetworkIdSensorAlertsProfilesConditions {
	this := NetworksNetworkIdSensorAlertsProfilesConditions{}
	this.Metric = metric
	this.Threshold = threshold
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesConditionsWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesConditionsWithDefaults() *NetworksNetworkIdSensorAlertsProfilesConditions {
	this := NetworksNetworkIdSensorAlertsProfilesConditions{}
	return &this
}

// GetMetric returns the Metric field value
func (o *NetworksNetworkIdSensorAlertsProfilesConditions) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesConditions) GetMetricOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *NetworksNetworkIdSensorAlertsProfilesConditions) SetMetric(v string) {
	o.Metric = v
}

// GetThreshold returns the Threshold field value
func (o *NetworksNetworkIdSensorAlertsProfilesConditions) GetThreshold() NetworksNetworkIdSensorAlertsProfilesThreshold {
	if o == nil {
		var ret NetworksNetworkIdSensorAlertsProfilesThreshold
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesConditions) GetThresholdOk() (*NetworksNetworkIdSensorAlertsProfilesThreshold, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *NetworksNetworkIdSensorAlertsProfilesConditions) SetThreshold(v NetworksNetworkIdSensorAlertsProfilesThreshold) {
	o.Threshold = v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesConditions) GetDirection() string {
	if o == nil || isNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesConditions) GetDirectionOk() (*string, bool) {
	if o == nil || isNil(o.Direction) {
    return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesConditions) HasDirection() bool {
	if o != nil && !isNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *NetworksNetworkIdSensorAlertsProfilesConditions) SetDirection(v string) {
	o.Direction = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesConditions) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesConditions) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesConditions) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *NetworksNetworkIdSensorAlertsProfilesConditions) SetDuration(v int32) {
	o.Duration = &v
}

func (o NetworksNetworkIdSensorAlertsProfilesConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metric"] = o.Metric
	}
	if true {
		toSerialize["threshold"] = o.Threshold
	}
	if !isNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesConditions struct {
	value *NetworksNetworkIdSensorAlertsProfilesConditions
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesConditions) Get() *NetworksNetworkIdSensorAlertsProfilesConditions {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesConditions) Set(val *NetworksNetworkIdSensorAlertsProfilesConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesConditions(val *NetworksNetworkIdSensorAlertsProfilesConditions) *NullableNetworksNetworkIdSensorAlertsProfilesConditions {
	return &NullableNetworksNetworkIdSensorAlertsProfilesConditions{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


