/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineObject94 struct for InlineObject94
type InlineObject94 struct {
	// Product type to rollback (if the network is a combined network)
	Product *string `json:"product,omitempty"`
	// Scheduled time for the rollback
	Time *time.Time `json:"time,omitempty"`
	// Reasons for the rollback
	Reasons []NetworksNetworkIdFirmwareUpgradesRollbacksReasons `json:"reasons"`
	ToVersion *NetworksNetworkIdFirmwareUpgradesRollbacksToVersion `json:"toVersion,omitempty"`
}

// NewInlineObject94 instantiates a new InlineObject94 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject94(reasons []NetworksNetworkIdFirmwareUpgradesRollbacksReasons) *InlineObject94 {
	this := InlineObject94{}
	this.Reasons = reasons
	return &this
}

// NewInlineObject94WithDefaults instantiates a new InlineObject94 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject94WithDefaults() *InlineObject94 {
	this := InlineObject94{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *InlineObject94) GetProduct() string {
	if o == nil || isNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject94) GetProductOk() (*string, bool) {
	if o == nil || isNil(o.Product) {
    return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *InlineObject94) HasProduct() bool {
	if o != nil && !isNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *InlineObject94) SetProduct(v string) {
	o.Product = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *InlineObject94) GetTime() time.Time {
	if o == nil || isNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject94) GetTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.Time) {
    return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *InlineObject94) HasTime() bool {
	if o != nil && !isNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *InlineObject94) SetTime(v time.Time) {
	o.Time = &v
}

// GetReasons returns the Reasons field value
func (o *InlineObject94) GetReasons() []NetworksNetworkIdFirmwareUpgradesRollbacksReasons {
	if o == nil {
		var ret []NetworksNetworkIdFirmwareUpgradesRollbacksReasons
		return ret
	}

	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value
// and a boolean to check if the value has been set.
func (o *InlineObject94) GetReasonsOk() ([]NetworksNetworkIdFirmwareUpgradesRollbacksReasons, bool) {
	if o == nil {
    return nil, false
	}
	return o.Reasons, true
}

// SetReasons sets field value
func (o *InlineObject94) SetReasons(v []NetworksNetworkIdFirmwareUpgradesRollbacksReasons) {
	o.Reasons = v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *InlineObject94) GetToVersion() NetworksNetworkIdFirmwareUpgradesRollbacksToVersion {
	if o == nil || isNil(o.ToVersion) {
		var ret NetworksNetworkIdFirmwareUpgradesRollbacksToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject94) GetToVersionOk() (*NetworksNetworkIdFirmwareUpgradesRollbacksToVersion, bool) {
	if o == nil || isNil(o.ToVersion) {
    return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *InlineObject94) HasToVersion() bool {
	if o != nil && !isNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given NetworksNetworkIdFirmwareUpgradesRollbacksToVersion and assigns it to the ToVersion field.
func (o *InlineObject94) SetToVersion(v NetworksNetworkIdFirmwareUpgradesRollbacksToVersion) {
	o.ToVersion = &v
}

func (o InlineObject94) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !isNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["reasons"] = o.Reasons
	}
	if !isNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject94 struct {
	value *InlineObject94
	isSet bool
}

func (v NullableInlineObject94) Get() *InlineObject94 {
	return v.value
}

func (v *NullableInlineObject94) Set(val *InlineObject94) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject94) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject94) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject94(val *InlineObject94) *NullableInlineObject94 {
	return &NullableInlineObject94{value: val, isSet: true}
}

func (v NullableInlineObject94) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject94) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


