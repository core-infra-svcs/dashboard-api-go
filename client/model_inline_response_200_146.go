/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200146 struct for InlineResponse200146
type InlineResponse200146 struct {
	// Monitored media server id
	Id *string `json:"id,omitempty"`
	// The name of the VoIP provider
	Name *string `json:"name,omitempty"`
	// The IP address (IPv4 only) or hostname of the media server to monitor
	Address *string `json:"address,omitempty"`
	// Indicates that if the media server doesn't respond to ICMP pings, the nearest hop will be used in its stead
	BestEffortMonitoringEnabled *bool `json:"bestEffortMonitoringEnabled,omitempty"`
}

// NewInlineResponse200146 instantiates a new InlineResponse200146 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200146() *InlineResponse200146 {
	this := InlineResponse200146{}
	return &this
}

// NewInlineResponse200146WithDefaults instantiates a new InlineResponse200146 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200146WithDefaults() *InlineResponse200146 {
	this := InlineResponse200146{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200146) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200146) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200146) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200146) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200146) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200146) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200146) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200146) SetName(v string) {
	o.Name = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *InlineResponse200146) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200146) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *InlineResponse200146) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *InlineResponse200146) SetAddress(v string) {
	o.Address = &v
}

// GetBestEffortMonitoringEnabled returns the BestEffortMonitoringEnabled field value if set, zero value otherwise.
func (o *InlineResponse200146) GetBestEffortMonitoringEnabled() bool {
	if o == nil || isNil(o.BestEffortMonitoringEnabled) {
		var ret bool
		return ret
	}
	return *o.BestEffortMonitoringEnabled
}

// GetBestEffortMonitoringEnabledOk returns a tuple with the BestEffortMonitoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200146) GetBestEffortMonitoringEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.BestEffortMonitoringEnabled) {
    return nil, false
	}
	return o.BestEffortMonitoringEnabled, true
}

// HasBestEffortMonitoringEnabled returns a boolean if a field has been set.
func (o *InlineResponse200146) HasBestEffortMonitoringEnabled() bool {
	if o != nil && !isNil(o.BestEffortMonitoringEnabled) {
		return true
	}

	return false
}

// SetBestEffortMonitoringEnabled gets a reference to the given bool and assigns it to the BestEffortMonitoringEnabled field.
func (o *InlineResponse200146) SetBestEffortMonitoringEnabled(v bool) {
	o.BestEffortMonitoringEnabled = &v
}

func (o InlineResponse200146) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.BestEffortMonitoringEnabled) {
		toSerialize["bestEffortMonitoringEnabled"] = o.BestEffortMonitoringEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200146 struct {
	value *InlineResponse200146
	isSet bool
}

func (v NullableInlineResponse200146) Get() *InlineResponse200146 {
	return v.value
}

func (v *NullableInlineResponse200146) Set(val *InlineResponse200146) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200146) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200146) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200146(val *InlineResponse200146) *NullableInlineResponse200146 {
	return &NullableInlineResponse200146{value: val, isSet: true}
}

func (v NullableInlineResponse200146) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200146) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


