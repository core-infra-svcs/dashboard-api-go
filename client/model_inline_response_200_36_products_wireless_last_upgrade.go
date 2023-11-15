/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20036ProductsWirelessLastUpgrade Details of the last firmware upgrade on the device
type InlineResponse20036ProductsWirelessLastUpgrade struct {
	// Timestamp of the last successful firmware upgrade
	Time *time.Time `json:"time,omitempty"`
	FromVersion *InlineResponse20036ProductsWirelessLastUpgradeFromVersion `json:"fromVersion,omitempty"`
	ToVersion *InlineResponse20036ProductsWirelessLastUpgradeToVersion `json:"toVersion,omitempty"`
}

// NewInlineResponse20036ProductsWirelessLastUpgrade instantiates a new InlineResponse20036ProductsWirelessLastUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20036ProductsWirelessLastUpgrade() *InlineResponse20036ProductsWirelessLastUpgrade {
	this := InlineResponse20036ProductsWirelessLastUpgrade{}
	return &this
}

// NewInlineResponse20036ProductsWirelessLastUpgradeWithDefaults instantiates a new InlineResponse20036ProductsWirelessLastUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20036ProductsWirelessLastUpgradeWithDefaults() *InlineResponse20036ProductsWirelessLastUpgrade {
	this := InlineResponse20036ProductsWirelessLastUpgrade{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *InlineResponse20036ProductsWirelessLastUpgrade) GetTime() time.Time {
	if o == nil || isNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036ProductsWirelessLastUpgrade) GetTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.Time) {
    return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *InlineResponse20036ProductsWirelessLastUpgrade) HasTime() bool {
	if o != nil && !isNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *InlineResponse20036ProductsWirelessLastUpgrade) SetTime(v time.Time) {
	o.Time = &v
}

// GetFromVersion returns the FromVersion field value if set, zero value otherwise.
func (o *InlineResponse20036ProductsWirelessLastUpgrade) GetFromVersion() InlineResponse20036ProductsWirelessLastUpgradeFromVersion {
	if o == nil || isNil(o.FromVersion) {
		var ret InlineResponse20036ProductsWirelessLastUpgradeFromVersion
		return ret
	}
	return *o.FromVersion
}

// GetFromVersionOk returns a tuple with the FromVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036ProductsWirelessLastUpgrade) GetFromVersionOk() (*InlineResponse20036ProductsWirelessLastUpgradeFromVersion, bool) {
	if o == nil || isNil(o.FromVersion) {
    return nil, false
	}
	return o.FromVersion, true
}

// HasFromVersion returns a boolean if a field has been set.
func (o *InlineResponse20036ProductsWirelessLastUpgrade) HasFromVersion() bool {
	if o != nil && !isNil(o.FromVersion) {
		return true
	}

	return false
}

// SetFromVersion gets a reference to the given InlineResponse20036ProductsWirelessLastUpgradeFromVersion and assigns it to the FromVersion field.
func (o *InlineResponse20036ProductsWirelessLastUpgrade) SetFromVersion(v InlineResponse20036ProductsWirelessLastUpgradeFromVersion) {
	o.FromVersion = &v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *InlineResponse20036ProductsWirelessLastUpgrade) GetToVersion() InlineResponse20036ProductsWirelessLastUpgradeToVersion {
	if o == nil || isNil(o.ToVersion) {
		var ret InlineResponse20036ProductsWirelessLastUpgradeToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036ProductsWirelessLastUpgrade) GetToVersionOk() (*InlineResponse20036ProductsWirelessLastUpgradeToVersion, bool) {
	if o == nil || isNil(o.ToVersion) {
    return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *InlineResponse20036ProductsWirelessLastUpgrade) HasToVersion() bool {
	if o != nil && !isNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given InlineResponse20036ProductsWirelessLastUpgradeToVersion and assigns it to the ToVersion field.
func (o *InlineResponse20036ProductsWirelessLastUpgrade) SetToVersion(v InlineResponse20036ProductsWirelessLastUpgradeToVersion) {
	o.ToVersion = &v
}

func (o InlineResponse20036ProductsWirelessLastUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !isNil(o.FromVersion) {
		toSerialize["fromVersion"] = o.FromVersion
	}
	if !isNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20036ProductsWirelessLastUpgrade struct {
	value *InlineResponse20036ProductsWirelessLastUpgrade
	isSet bool
}

func (v NullableInlineResponse20036ProductsWirelessLastUpgrade) Get() *InlineResponse20036ProductsWirelessLastUpgrade {
	return v.value
}

func (v *NullableInlineResponse20036ProductsWirelessLastUpgrade) Set(val *InlineResponse20036ProductsWirelessLastUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20036ProductsWirelessLastUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20036ProductsWirelessLastUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20036ProductsWirelessLastUpgrade(val *InlineResponse20036ProductsWirelessLastUpgrade) *NullableInlineResponse20036ProductsWirelessLastUpgrade {
	return &NullableInlineResponse20036ProductsWirelessLastUpgrade{value: val, isSet: true}
}

func (v NullableInlineResponse20036ProductsWirelessLastUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20036ProductsWirelessLastUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

