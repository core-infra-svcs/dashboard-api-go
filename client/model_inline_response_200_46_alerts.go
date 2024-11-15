/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20046Alerts struct for InlineResponse20046Alerts
type InlineResponse20046Alerts struct {
	// The type of alert
	Type string `json:"type"`
	// A boolean depicting if the alert is turned on or off
	Enabled *bool `json:"enabled,omitempty"`
	AlertDestinations *InlineResponse20046AlertDestinations `json:"alertDestinations,omitempty"`
	Filters *InlineResponse20046Filters `json:"filters,omitempty"`
}

// NewInlineResponse20046Alerts instantiates a new InlineResponse20046Alerts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20046Alerts(type_ string) *InlineResponse20046Alerts {
	this := InlineResponse20046Alerts{}
	this.Type = type_
	return &this
}

// NewInlineResponse20046AlertsWithDefaults instantiates a new InlineResponse20046Alerts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20046AlertsWithDefaults() *InlineResponse20046Alerts {
	this := InlineResponse20046Alerts{}
	return &this
}

// GetType returns the Type field value
func (o *InlineResponse20046Alerts) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Alerts) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse20046Alerts) SetType(v string) {
	o.Type = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20046Alerts) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Alerts) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20046Alerts) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20046Alerts) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAlertDestinations returns the AlertDestinations field value if set, zero value otherwise.
func (o *InlineResponse20046Alerts) GetAlertDestinations() InlineResponse20046AlertDestinations {
	if o == nil || isNil(o.AlertDestinations) {
		var ret InlineResponse20046AlertDestinations
		return ret
	}
	return *o.AlertDestinations
}

// GetAlertDestinationsOk returns a tuple with the AlertDestinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Alerts) GetAlertDestinationsOk() (*InlineResponse20046AlertDestinations, bool) {
	if o == nil || isNil(o.AlertDestinations) {
    return nil, false
	}
	return o.AlertDestinations, true
}

// HasAlertDestinations returns a boolean if a field has been set.
func (o *InlineResponse20046Alerts) HasAlertDestinations() bool {
	if o != nil && !isNil(o.AlertDestinations) {
		return true
	}

	return false
}

// SetAlertDestinations gets a reference to the given InlineResponse20046AlertDestinations and assigns it to the AlertDestinations field.
func (o *InlineResponse20046Alerts) SetAlertDestinations(v InlineResponse20046AlertDestinations) {
	o.AlertDestinations = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *InlineResponse20046Alerts) GetFilters() InlineResponse20046Filters {
	if o == nil || isNil(o.Filters) {
		var ret InlineResponse20046Filters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Alerts) GetFiltersOk() (*InlineResponse20046Filters, bool) {
	if o == nil || isNil(o.Filters) {
    return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *InlineResponse20046Alerts) HasFilters() bool {
	if o != nil && !isNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given InlineResponse20046Filters and assigns it to the Filters field.
func (o *InlineResponse20046Alerts) SetFilters(v InlineResponse20046Filters) {
	o.Filters = &v
}

func (o InlineResponse20046Alerts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.AlertDestinations) {
		toSerialize["alertDestinations"] = o.AlertDestinations
	}
	if !isNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20046Alerts struct {
	value *InlineResponse20046Alerts
	isSet bool
}

func (v NullableInlineResponse20046Alerts) Get() *InlineResponse20046Alerts {
	return v.value
}

func (v *NullableInlineResponse20046Alerts) Set(val *InlineResponse20046Alerts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20046Alerts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20046Alerts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20046Alerts(val *InlineResponse20046Alerts) *NullableInlineResponse20046Alerts {
	return &NullableInlineResponse20046Alerts{value: val, isSet: true}
}

func (v NullableInlineResponse20046Alerts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20046Alerts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


