/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20030ProductsWirelessNextUpgrade Details of the next firmware upgrade on the device
type InlineResponse20030ProductsWirelessNextUpgrade struct {
	// Timestamp of the next scheduled firmware upgrade
	Time *time.Time `json:"time,omitempty"`
	ToVersion *InlineResponse20030ProductsWirelessNextUpgradeToVersion `json:"toVersion,omitempty"`
}

// NewInlineResponse20030ProductsWirelessNextUpgrade instantiates a new InlineResponse20030ProductsWirelessNextUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20030ProductsWirelessNextUpgrade() *InlineResponse20030ProductsWirelessNextUpgrade {
	this := InlineResponse20030ProductsWirelessNextUpgrade{}
	return &this
}

// NewInlineResponse20030ProductsWirelessNextUpgradeWithDefaults instantiates a new InlineResponse20030ProductsWirelessNextUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20030ProductsWirelessNextUpgradeWithDefaults() *InlineResponse20030ProductsWirelessNextUpgrade {
	this := InlineResponse20030ProductsWirelessNextUpgrade{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *InlineResponse20030ProductsWirelessNextUpgrade) GetTime() time.Time {
	if o == nil || isNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030ProductsWirelessNextUpgrade) GetTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.Time) {
    return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *InlineResponse20030ProductsWirelessNextUpgrade) HasTime() bool {
	if o != nil && !isNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *InlineResponse20030ProductsWirelessNextUpgrade) SetTime(v time.Time) {
	o.Time = &v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *InlineResponse20030ProductsWirelessNextUpgrade) GetToVersion() InlineResponse20030ProductsWirelessNextUpgradeToVersion {
	if o == nil || isNil(o.ToVersion) {
		var ret InlineResponse20030ProductsWirelessNextUpgradeToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030ProductsWirelessNextUpgrade) GetToVersionOk() (*InlineResponse20030ProductsWirelessNextUpgradeToVersion, bool) {
	if o == nil || isNil(o.ToVersion) {
    return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *InlineResponse20030ProductsWirelessNextUpgrade) HasToVersion() bool {
	if o != nil && !isNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given InlineResponse20030ProductsWirelessNextUpgradeToVersion and assigns it to the ToVersion field.
func (o *InlineResponse20030ProductsWirelessNextUpgrade) SetToVersion(v InlineResponse20030ProductsWirelessNextUpgradeToVersion) {
	o.ToVersion = &v
}

func (o InlineResponse20030ProductsWirelessNextUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !isNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20030ProductsWirelessNextUpgrade struct {
	value *InlineResponse20030ProductsWirelessNextUpgrade
	isSet bool
}

func (v NullableInlineResponse20030ProductsWirelessNextUpgrade) Get() *InlineResponse20030ProductsWirelessNextUpgrade {
	return v.value
}

func (v *NullableInlineResponse20030ProductsWirelessNextUpgrade) Set(val *InlineResponse20030ProductsWirelessNextUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20030ProductsWirelessNextUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20030ProductsWirelessNextUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20030ProductsWirelessNextUpgrade(val *InlineResponse20030ProductsWirelessNextUpgrade) *NullableInlineResponse20030ProductsWirelessNextUpgrade {
	return &NullableInlineResponse20030ProductsWirelessNextUpgrade{value: val, isSet: true}
}

func (v NullableInlineResponse20030ProductsWirelessNextUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20030ProductsWirelessNextUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


