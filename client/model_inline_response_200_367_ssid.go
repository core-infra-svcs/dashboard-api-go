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

// InlineResponse200367Ssid Wireless access point and network identifier.
type InlineResponse200367Ssid struct {
	// Name of wireless network.
	Name *string `json:"name,omitempty"`
	// Unique identifier for wireless network.
	Number *int32 `json:"number,omitempty"`
	// Status of wireless network.
	Enabled *bool `json:"enabled,omitempty"`
	// Availability of wireless network for devices to connect to.
	Advertised *bool `json:"advertised,omitempty"`
}

// NewInlineResponse200367Ssid instantiates a new InlineResponse200367Ssid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200367Ssid() *InlineResponse200367Ssid {
	this := InlineResponse200367Ssid{}
	return &this
}

// NewInlineResponse200367SsidWithDefaults instantiates a new InlineResponse200367Ssid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200367SsidWithDefaults() *InlineResponse200367Ssid {
	this := InlineResponse200367Ssid{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200367Ssid) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200367Ssid) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200367Ssid) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200367Ssid) SetName(v string) {
	o.Name = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *InlineResponse200367Ssid) GetNumber() int32 {
	if o == nil || isNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200367Ssid) GetNumberOk() (*int32, bool) {
	if o == nil || isNil(o.Number) {
    return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineResponse200367Ssid) HasNumber() bool {
	if o != nil && !isNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *InlineResponse200367Ssid) SetNumber(v int32) {
	o.Number = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200367Ssid) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200367Ssid) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200367Ssid) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200367Ssid) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAdvertised returns the Advertised field value if set, zero value otherwise.
func (o *InlineResponse200367Ssid) GetAdvertised() bool {
	if o == nil || isNil(o.Advertised) {
		var ret bool
		return ret
	}
	return *o.Advertised
}

// GetAdvertisedOk returns a tuple with the Advertised field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200367Ssid) GetAdvertisedOk() (*bool, bool) {
	if o == nil || isNil(o.Advertised) {
    return nil, false
	}
	return o.Advertised, true
}

// HasAdvertised returns a boolean if a field has been set.
func (o *InlineResponse200367Ssid) HasAdvertised() bool {
	if o != nil && !isNil(o.Advertised) {
		return true
	}

	return false
}

// SetAdvertised gets a reference to the given bool and assigns it to the Advertised field.
func (o *InlineResponse200367Ssid) SetAdvertised(v bool) {
	o.Advertised = &v
}

func (o InlineResponse200367Ssid) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Advertised) {
		toSerialize["advertised"] = o.Advertised
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200367Ssid struct {
	value *InlineResponse200367Ssid
	isSet bool
}

func (v NullableInlineResponse200367Ssid) Get() *InlineResponse200367Ssid {
	return v.value
}

func (v *NullableInlineResponse200367Ssid) Set(val *InlineResponse200367Ssid) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200367Ssid) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200367Ssid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200367Ssid(val *InlineResponse200367Ssid) *NullableInlineResponse200367Ssid {
	return &NullableInlineResponse200367Ssid{value: val, isSet: true}
}

func (v NullableInlineResponse200367Ssid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200367Ssid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


