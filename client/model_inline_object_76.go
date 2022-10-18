/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineObject76 struct for InlineObject76
type InlineObject76 struct {
	// Product type to rollback (if the network is a combined network)
	Product *string `json:"product,omitempty"`
	// Scheduled time for the rollback
	Time *time.Time `json:"time,omitempty"`
	// Reasons for the rollback
	Reasons []NetworksNetworkIdFirmwareUpgradesRollbacksReasons `json:"reasons"`
	ToVersion *NetworksNetworkIdFirmwareUpgradesRollbacksToVersion `json:"toVersion,omitempty"`
}

// NewInlineObject76 instantiates a new InlineObject76 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject76(reasons []NetworksNetworkIdFirmwareUpgradesRollbacksReasons) *InlineObject76 {
	this := InlineObject76{}
	this.Reasons = reasons
	return &this
}

// NewInlineObject76WithDefaults instantiates a new InlineObject76 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject76WithDefaults() *InlineObject76 {
	this := InlineObject76{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *InlineObject76) GetProduct() string {
	if o == nil || o.Product == nil {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetProductOk() (*string, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *InlineObject76) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *InlineObject76) SetProduct(v string) {
	o.Product = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *InlineObject76) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetTimeOk() (*time.Time, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *InlineObject76) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *InlineObject76) SetTime(v time.Time) {
	o.Time = &v
}

// GetReasons returns the Reasons field value
func (o *InlineObject76) GetReasons() []NetworksNetworkIdFirmwareUpgradesRollbacksReasons {
	if o == nil {
		var ret []NetworksNetworkIdFirmwareUpgradesRollbacksReasons
		return ret
	}

	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetReasonsOk() (*[]NetworksNetworkIdFirmwareUpgradesRollbacksReasons, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reasons, true
}

// SetReasons sets field value
func (o *InlineObject76) SetReasons(v []NetworksNetworkIdFirmwareUpgradesRollbacksReasons) {
	o.Reasons = v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *InlineObject76) GetToVersion() NetworksNetworkIdFirmwareUpgradesRollbacksToVersion {
	if o == nil || o.ToVersion == nil {
		var ret NetworksNetworkIdFirmwareUpgradesRollbacksToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetToVersionOk() (*NetworksNetworkIdFirmwareUpgradesRollbacksToVersion, bool) {
	if o == nil || o.ToVersion == nil {
		return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *InlineObject76) HasToVersion() bool {
	if o != nil && o.ToVersion != nil {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given NetworksNetworkIdFirmwareUpgradesRollbacksToVersion and assigns it to the ToVersion field.
func (o *InlineObject76) SetToVersion(v NetworksNetworkIdFirmwareUpgradesRollbacksToVersion) {
	o.ToVersion = &v
}

func (o InlineObject76) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Product != nil {
		toSerialize["product"] = o.Product
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["reasons"] = o.Reasons
	}
	if o.ToVersion != nil {
		toSerialize["toVersion"] = o.ToVersion
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject76 struct {
	value *InlineObject76
	isSet bool
}

func (v NullableInlineObject76) Get() *InlineObject76 {
	return v.value
}

func (v *NullableInlineObject76) Set(val *InlineObject76) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject76) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject76) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject76(val *InlineObject76) *NullableInlineObject76 {
	return &NullableInlineObject76{value: val, isSet: true}
}

func (v NullableInlineObject76) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject76) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


