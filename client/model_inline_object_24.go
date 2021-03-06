/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject24 struct for InlineObject24
type InlineObject24 struct {
	DefaultDestinations *NetworksNetworkIdAlertsSettingsDefaultDestinations `json:"defaultDestinations,omitempty"`
	// Alert-specific configuration for each type. Only alerts that pertain to the network can be updated.
	Alerts *[]NetworksNetworkIdAlertsSettingsAlerts `json:"alerts,omitempty"`
}

// NewInlineObject24 instantiates a new InlineObject24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject24() *InlineObject24 {
	this := InlineObject24{}
	return &this
}

// NewInlineObject24WithDefaults instantiates a new InlineObject24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject24WithDefaults() *InlineObject24 {
	this := InlineObject24{}
	return &this
}

// GetDefaultDestinations returns the DefaultDestinations field value if set, zero value otherwise.
func (o *InlineObject24) GetDefaultDestinations() NetworksNetworkIdAlertsSettingsDefaultDestinations {
	if o == nil || o.DefaultDestinations == nil {
		var ret NetworksNetworkIdAlertsSettingsDefaultDestinations
		return ret
	}
	return *o.DefaultDestinations
}

// GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject24) GetDefaultDestinationsOk() (*NetworksNetworkIdAlertsSettingsDefaultDestinations, bool) {
	if o == nil || o.DefaultDestinations == nil {
		return nil, false
	}
	return o.DefaultDestinations, true
}

// HasDefaultDestinations returns a boolean if a field has been set.
func (o *InlineObject24) HasDefaultDestinations() bool {
	if o != nil && o.DefaultDestinations != nil {
		return true
	}

	return false
}

// SetDefaultDestinations gets a reference to the given NetworksNetworkIdAlertsSettingsDefaultDestinations and assigns it to the DefaultDestinations field.
func (o *InlineObject24) SetDefaultDestinations(v NetworksNetworkIdAlertsSettingsDefaultDestinations) {
	o.DefaultDestinations = &v
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *InlineObject24) GetAlerts() []NetworksNetworkIdAlertsSettingsAlerts {
	if o == nil || o.Alerts == nil {
		var ret []NetworksNetworkIdAlertsSettingsAlerts
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject24) GetAlertsOk() (*[]NetworksNetworkIdAlertsSettingsAlerts, bool) {
	if o == nil || o.Alerts == nil {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *InlineObject24) HasAlerts() bool {
	if o != nil && o.Alerts != nil {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given []NetworksNetworkIdAlertsSettingsAlerts and assigns it to the Alerts field.
func (o *InlineObject24) SetAlerts(v []NetworksNetworkIdAlertsSettingsAlerts) {
	o.Alerts = &v
}

func (o InlineObject24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultDestinations != nil {
		toSerialize["defaultDestinations"] = o.DefaultDestinations
	}
	if o.Alerts != nil {
		toSerialize["alerts"] = o.Alerts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject24 struct {
	value *InlineObject24
	isSet bool
}

func (v NullableInlineObject24) Get() *InlineObject24 {
	return v.value
}

func (v *NullableInlineObject24) Set(val *InlineObject24) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject24) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject24(val *InlineObject24) *NullableInlineObject24 {
	return &NullableInlineObject24{value: val, isSet: true}
}

func (v NullableInlineObject24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


