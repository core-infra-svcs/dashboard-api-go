/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject29 struct for InlineObject29
type InlineObject29 struct {
	DefaultDestinations *NetworksNetworkIdAlertsSettingsDefaultDestinations `json:"defaultDestinations,omitempty"`
	// Alert-specific configuration for each type. Only alerts that pertain to the network can be updated.
	Alerts []NetworksNetworkIdAlertsSettingsAlerts `json:"alerts,omitempty"`
	Muting *NetworksNetworkIdAlertsSettingsMuting `json:"muting,omitempty"`
}

// NewInlineObject29 instantiates a new InlineObject29 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject29() *InlineObject29 {
	this := InlineObject29{}
	return &this
}

// NewInlineObject29WithDefaults instantiates a new InlineObject29 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject29WithDefaults() *InlineObject29 {
	this := InlineObject29{}
	return &this
}

// GetDefaultDestinations returns the DefaultDestinations field value if set, zero value otherwise.
func (o *InlineObject29) GetDefaultDestinations() NetworksNetworkIdAlertsSettingsDefaultDestinations {
	if o == nil || isNil(o.DefaultDestinations) {
		var ret NetworksNetworkIdAlertsSettingsDefaultDestinations
		return ret
	}
	return *o.DefaultDestinations
}

// GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject29) GetDefaultDestinationsOk() (*NetworksNetworkIdAlertsSettingsDefaultDestinations, bool) {
	if o == nil || isNil(o.DefaultDestinations) {
    return nil, false
	}
	return o.DefaultDestinations, true
}

// HasDefaultDestinations returns a boolean if a field has been set.
func (o *InlineObject29) HasDefaultDestinations() bool {
	if o != nil && !isNil(o.DefaultDestinations) {
		return true
	}

	return false
}

// SetDefaultDestinations gets a reference to the given NetworksNetworkIdAlertsSettingsDefaultDestinations and assigns it to the DefaultDestinations field.
func (o *InlineObject29) SetDefaultDestinations(v NetworksNetworkIdAlertsSettingsDefaultDestinations) {
	o.DefaultDestinations = &v
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *InlineObject29) GetAlerts() []NetworksNetworkIdAlertsSettingsAlerts {
	if o == nil || isNil(o.Alerts) {
		var ret []NetworksNetworkIdAlertsSettingsAlerts
		return ret
	}
	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject29) GetAlertsOk() ([]NetworksNetworkIdAlertsSettingsAlerts, bool) {
	if o == nil || isNil(o.Alerts) {
    return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *InlineObject29) HasAlerts() bool {
	if o != nil && !isNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given []NetworksNetworkIdAlertsSettingsAlerts and assigns it to the Alerts field.
func (o *InlineObject29) SetAlerts(v []NetworksNetworkIdAlertsSettingsAlerts) {
	o.Alerts = v
}

// GetMuting returns the Muting field value if set, zero value otherwise.
func (o *InlineObject29) GetMuting() NetworksNetworkIdAlertsSettingsMuting {
	if o == nil || isNil(o.Muting) {
		var ret NetworksNetworkIdAlertsSettingsMuting
		return ret
	}
	return *o.Muting
}

// GetMutingOk returns a tuple with the Muting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject29) GetMutingOk() (*NetworksNetworkIdAlertsSettingsMuting, bool) {
	if o == nil || isNil(o.Muting) {
    return nil, false
	}
	return o.Muting, true
}

// HasMuting returns a boolean if a field has been set.
func (o *InlineObject29) HasMuting() bool {
	if o != nil && !isNil(o.Muting) {
		return true
	}

	return false
}

// SetMuting gets a reference to the given NetworksNetworkIdAlertsSettingsMuting and assigns it to the Muting field.
func (o *InlineObject29) SetMuting(v NetworksNetworkIdAlertsSettingsMuting) {
	o.Muting = &v
}

func (o InlineObject29) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DefaultDestinations) {
		toSerialize["defaultDestinations"] = o.DefaultDestinations
	}
	if !isNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !isNil(o.Muting) {
		toSerialize["muting"] = o.Muting
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject29 struct {
	value *InlineObject29
	isSet bool
}

func (v NullableInlineObject29) Get() *InlineObject29 {
	return v.value
}

func (v *NullableInlineObject29) Set(val *InlineObject29) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject29) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject29) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject29(val *InlineObject29) *NullableInlineObject29 {
	return &NullableInlineObject29{value: val, isSet: true}
}

func (v NullableInlineObject29) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject29) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


