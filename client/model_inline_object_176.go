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

// InlineObject176 struct for InlineObject176
type InlineObject176 struct {
	// The name of the VoIP provider
	Name string `json:"name"`
	// The IP address (IPv4 only) or hostname of the media server to monitor
	Address string `json:"address"`
	// Indicates that if the media server doesn't respond to ICMP pings, the nearest hop will be used in its stead.
	BestEffortMonitoringEnabled *bool `json:"bestEffortMonitoringEnabled,omitempty"`
}

// NewInlineObject176 instantiates a new InlineObject176 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject176(name string, address string) *InlineObject176 {
	this := InlineObject176{}
	this.Name = name
	this.Address = address
	return &this
}

// NewInlineObject176WithDefaults instantiates a new InlineObject176 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject176WithDefaults() *InlineObject176 {
	this := InlineObject176{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject176) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject176) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject176) SetName(v string) {
	o.Name = v
}

// GetAddress returns the Address field value
func (o *InlineObject176) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *InlineObject176) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *InlineObject176) SetAddress(v string) {
	o.Address = v
}

// GetBestEffortMonitoringEnabled returns the BestEffortMonitoringEnabled field value if set, zero value otherwise.
func (o *InlineObject176) GetBestEffortMonitoringEnabled() bool {
	if o == nil || o.BestEffortMonitoringEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BestEffortMonitoringEnabled
}

// GetBestEffortMonitoringEnabledOk returns a tuple with the BestEffortMonitoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject176) GetBestEffortMonitoringEnabledOk() (*bool, bool) {
	if o == nil || o.BestEffortMonitoringEnabled == nil {
		return nil, false
	}
	return o.BestEffortMonitoringEnabled, true
}

// HasBestEffortMonitoringEnabled returns a boolean if a field has been set.
func (o *InlineObject176) HasBestEffortMonitoringEnabled() bool {
	if o != nil && o.BestEffortMonitoringEnabled != nil {
		return true
	}

	return false
}

// SetBestEffortMonitoringEnabled gets a reference to the given bool and assigns it to the BestEffortMonitoringEnabled field.
func (o *InlineObject176) SetBestEffortMonitoringEnabled(v bool) {
	o.BestEffortMonitoringEnabled = &v
}

func (o InlineObject176) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if o.BestEffortMonitoringEnabled != nil {
		toSerialize["bestEffortMonitoringEnabled"] = o.BestEffortMonitoringEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject176 struct {
	value *InlineObject176
	isSet bool
}

func (v NullableInlineObject176) Get() *InlineObject176 {
	return v.value
}

func (v *NullableInlineObject176) Set(val *InlineObject176) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject176) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject176) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject176(val *InlineObject176) *NullableInlineObject176 {
	return &NullableInlineObject176{value: val, isSet: true}
}

func (v NullableInlineObject176) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject176) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


