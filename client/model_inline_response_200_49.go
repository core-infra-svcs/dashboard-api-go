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

// InlineResponse20049 struct for InlineResponse20049
type InlineResponse20049 struct {
	DefaultDestinations *InlineResponse20049DefaultDestinations `json:"defaultDestinations,omitempty"`
	// Alert-specific configuration for each type. Only alerts that pertain to the network can be updated.
	Alerts []InlineResponse20049Alerts `json:"alerts,omitempty"`
	Muting *InlineResponse20049Muting `json:"muting,omitempty"`
}

// NewInlineResponse20049 instantiates a new InlineResponse20049 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20049() *InlineResponse20049 {
	this := InlineResponse20049{}
	return &this
}

// NewInlineResponse20049WithDefaults instantiates a new InlineResponse20049 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20049WithDefaults() *InlineResponse20049 {
	this := InlineResponse20049{}
	return &this
}

// GetDefaultDestinations returns the DefaultDestinations field value if set, zero value otherwise.
func (o *InlineResponse20049) GetDefaultDestinations() InlineResponse20049DefaultDestinations {
	if o == nil || isNil(o.DefaultDestinations) {
		var ret InlineResponse20049DefaultDestinations
		return ret
	}
	return *o.DefaultDestinations
}

// GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20049) GetDefaultDestinationsOk() (*InlineResponse20049DefaultDestinations, bool) {
	if o == nil || isNil(o.DefaultDestinations) {
    return nil, false
	}
	return o.DefaultDestinations, true
}

// HasDefaultDestinations returns a boolean if a field has been set.
func (o *InlineResponse20049) HasDefaultDestinations() bool {
	if o != nil && !isNil(o.DefaultDestinations) {
		return true
	}

	return false
}

// SetDefaultDestinations gets a reference to the given InlineResponse20049DefaultDestinations and assigns it to the DefaultDestinations field.
func (o *InlineResponse20049) SetDefaultDestinations(v InlineResponse20049DefaultDestinations) {
	o.DefaultDestinations = &v
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *InlineResponse20049) GetAlerts() []InlineResponse20049Alerts {
	if o == nil || isNil(o.Alerts) {
		var ret []InlineResponse20049Alerts
		return ret
	}
	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20049) GetAlertsOk() ([]InlineResponse20049Alerts, bool) {
	if o == nil || isNil(o.Alerts) {
    return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *InlineResponse20049) HasAlerts() bool {
	if o != nil && !isNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given []InlineResponse20049Alerts and assigns it to the Alerts field.
func (o *InlineResponse20049) SetAlerts(v []InlineResponse20049Alerts) {
	o.Alerts = v
}

// GetMuting returns the Muting field value if set, zero value otherwise.
func (o *InlineResponse20049) GetMuting() InlineResponse20049Muting {
	if o == nil || isNil(o.Muting) {
		var ret InlineResponse20049Muting
		return ret
	}
	return *o.Muting
}

// GetMutingOk returns a tuple with the Muting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20049) GetMutingOk() (*InlineResponse20049Muting, bool) {
	if o == nil || isNil(o.Muting) {
    return nil, false
	}
	return o.Muting, true
}

// HasMuting returns a boolean if a field has been set.
func (o *InlineResponse20049) HasMuting() bool {
	if o != nil && !isNil(o.Muting) {
		return true
	}

	return false
}

// SetMuting gets a reference to the given InlineResponse20049Muting and assigns it to the Muting field.
func (o *InlineResponse20049) SetMuting(v InlineResponse20049Muting) {
	o.Muting = &v
}

func (o InlineResponse20049) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20049 struct {
	value *InlineResponse20049
	isSet bool
}

func (v NullableInlineResponse20049) Get() *InlineResponse20049 {
	return v.value
}

func (v *NullableInlineResponse20049) Set(val *InlineResponse20049) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20049) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20049) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20049(val *InlineResponse20049) *NullableInlineResponse20049 {
	return &NullableInlineResponse20049{value: val, isSet: true}
}

func (v NullableInlineResponse20049) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20049) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


