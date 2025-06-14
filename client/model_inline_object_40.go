/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject40 struct for InlineObject40
type InlineObject40 struct {
	DefaultDestinations *InlineResponse20050DefaultDestinations `json:"defaultDestinations,omitempty"`
	// Alert-specific configuration for each type. Only alerts that pertain to the network can be updated.
	Alerts []InlineResponse20050Alerts `json:"alerts,omitempty"`
	Muting *InlineResponse20050Muting `json:"muting,omitempty"`
}

// NewInlineObject40 instantiates a new InlineObject40 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject40() *InlineObject40 {
	this := InlineObject40{}
	return &this
}

// NewInlineObject40WithDefaults instantiates a new InlineObject40 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject40WithDefaults() *InlineObject40 {
	this := InlineObject40{}
	return &this
}

// GetDefaultDestinations returns the DefaultDestinations field value if set, zero value otherwise.
func (o *InlineObject40) GetDefaultDestinations() InlineResponse20050DefaultDestinations {
	if o == nil || isNil(o.DefaultDestinations) {
		var ret InlineResponse20050DefaultDestinations
		return ret
	}
	return *o.DefaultDestinations
}

// GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject40) GetDefaultDestinationsOk() (*InlineResponse20050DefaultDestinations, bool) {
	if o == nil || isNil(o.DefaultDestinations) {
    return nil, false
	}
	return o.DefaultDestinations, true
}

// HasDefaultDestinations returns a boolean if a field has been set.
func (o *InlineObject40) HasDefaultDestinations() bool {
	if o != nil && !isNil(o.DefaultDestinations) {
		return true
	}

	return false
}

// SetDefaultDestinations gets a reference to the given InlineResponse20050DefaultDestinations and assigns it to the DefaultDestinations field.
func (o *InlineObject40) SetDefaultDestinations(v InlineResponse20050DefaultDestinations) {
	o.DefaultDestinations = &v
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *InlineObject40) GetAlerts() []InlineResponse20050Alerts {
	if o == nil || isNil(o.Alerts) {
		var ret []InlineResponse20050Alerts
		return ret
	}
	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject40) GetAlertsOk() ([]InlineResponse20050Alerts, bool) {
	if o == nil || isNil(o.Alerts) {
    return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *InlineObject40) HasAlerts() bool {
	if o != nil && !isNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given []InlineResponse20050Alerts and assigns it to the Alerts field.
func (o *InlineObject40) SetAlerts(v []InlineResponse20050Alerts) {
	o.Alerts = v
}

// GetMuting returns the Muting field value if set, zero value otherwise.
func (o *InlineObject40) GetMuting() InlineResponse20050Muting {
	if o == nil || isNil(o.Muting) {
		var ret InlineResponse20050Muting
		return ret
	}
	return *o.Muting
}

// GetMutingOk returns a tuple with the Muting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject40) GetMutingOk() (*InlineResponse20050Muting, bool) {
	if o == nil || isNil(o.Muting) {
    return nil, false
	}
	return o.Muting, true
}

// HasMuting returns a boolean if a field has been set.
func (o *InlineObject40) HasMuting() bool {
	if o != nil && !isNil(o.Muting) {
		return true
	}

	return false
}

// SetMuting gets a reference to the given InlineResponse20050Muting and assigns it to the Muting field.
func (o *InlineObject40) SetMuting(v InlineResponse20050Muting) {
	o.Muting = &v
}

func (o InlineObject40) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject40 struct {
	value *InlineObject40
	isSet bool
}

func (v NullableInlineObject40) Get() *InlineObject40 {
	return v.value
}

func (v *NullableInlineObject40) Set(val *InlineObject40) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject40) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject40) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject40(val *InlineObject40) *NullableInlineObject40 {
	return &NullableInlineObject40{value: val, isSet: true}
}

func (v NullableInlineObject40) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject40) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


