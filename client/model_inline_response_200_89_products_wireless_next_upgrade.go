/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20089ProductsWirelessNextUpgrade Details of the next firmware upgrade on the device
type InlineResponse20089ProductsWirelessNextUpgrade struct {
	// Timestamp of the next scheduled firmware upgrade
	Time *time.Time `json:"time,omitempty"`
	ToVersion *InlineResponse20089ProductsWirelessNextUpgradeToVersion `json:"toVersion,omitempty"`
}

// NewInlineResponse20089ProductsWirelessNextUpgrade instantiates a new InlineResponse20089ProductsWirelessNextUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20089ProductsWirelessNextUpgrade() *InlineResponse20089ProductsWirelessNextUpgrade {
	this := InlineResponse20089ProductsWirelessNextUpgrade{}
	return &this
}

// NewInlineResponse20089ProductsWirelessNextUpgradeWithDefaults instantiates a new InlineResponse20089ProductsWirelessNextUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20089ProductsWirelessNextUpgradeWithDefaults() *InlineResponse20089ProductsWirelessNextUpgrade {
	this := InlineResponse20089ProductsWirelessNextUpgrade{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *InlineResponse20089ProductsWirelessNextUpgrade) GetTime() time.Time {
	if o == nil || isNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089ProductsWirelessNextUpgrade) GetTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.Time) {
    return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *InlineResponse20089ProductsWirelessNextUpgrade) HasTime() bool {
	if o != nil && !isNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *InlineResponse20089ProductsWirelessNextUpgrade) SetTime(v time.Time) {
	o.Time = &v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *InlineResponse20089ProductsWirelessNextUpgrade) GetToVersion() InlineResponse20089ProductsWirelessNextUpgradeToVersion {
	if o == nil || isNil(o.ToVersion) {
		var ret InlineResponse20089ProductsWirelessNextUpgradeToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089ProductsWirelessNextUpgrade) GetToVersionOk() (*InlineResponse20089ProductsWirelessNextUpgradeToVersion, bool) {
	if o == nil || isNil(o.ToVersion) {
    return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *InlineResponse20089ProductsWirelessNextUpgrade) HasToVersion() bool {
	if o != nil && !isNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given InlineResponse20089ProductsWirelessNextUpgradeToVersion and assigns it to the ToVersion field.
func (o *InlineResponse20089ProductsWirelessNextUpgrade) SetToVersion(v InlineResponse20089ProductsWirelessNextUpgradeToVersion) {
	o.ToVersion = &v
}

func (o InlineResponse20089ProductsWirelessNextUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !isNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20089ProductsWirelessNextUpgrade struct {
	value *InlineResponse20089ProductsWirelessNextUpgrade
	isSet bool
}

func (v NullableInlineResponse20089ProductsWirelessNextUpgrade) Get() *InlineResponse20089ProductsWirelessNextUpgrade {
	return v.value
}

func (v *NullableInlineResponse20089ProductsWirelessNextUpgrade) Set(val *InlineResponse20089ProductsWirelessNextUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20089ProductsWirelessNextUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20089ProductsWirelessNextUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20089ProductsWirelessNextUpgrade(val *InlineResponse20089ProductsWirelessNextUpgrade) *NullableInlineResponse20089ProductsWirelessNextUpgrade {
	return &NullableInlineResponse20089ProductsWirelessNextUpgrade{value: val, isSet: true}
}

func (v NullableInlineResponse20089ProductsWirelessNextUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20089ProductsWirelessNextUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


