/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200312Items struct for InlineResponse200312Items
type InlineResponse200312Items struct {
	Network *InlineResponse200312Network `json:"network,omitempty"`
	// Access point Serial number
	Serial *string `json:"serial,omitempty"`
	Counts *InlineResponse200312Counts `json:"counts,omitempty"`
}

// NewInlineResponse200312Items instantiates a new InlineResponse200312Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200312Items() *InlineResponse200312Items {
	this := InlineResponse200312Items{}
	return &this
}

// NewInlineResponse200312ItemsWithDefaults instantiates a new InlineResponse200312Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200312ItemsWithDefaults() *InlineResponse200312Items {
	this := InlineResponse200312Items{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200312Items) GetNetwork() InlineResponse200312Network {
	if o == nil || isNil(o.Network) {
		var ret InlineResponse200312Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200312Items) GetNetworkOk() (*InlineResponse200312Network, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200312Items) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given InlineResponse200312Network and assigns it to the Network field.
func (o *InlineResponse200312Items) SetNetwork(v InlineResponse200312Network) {
	o.Network = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200312Items) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200312Items) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200312Items) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200312Items) SetSerial(v string) {
	o.Serial = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200312Items) GetCounts() InlineResponse200312Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200312Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200312Items) GetCountsOk() (*InlineResponse200312Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200312Items) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200312Counts and assigns it to the Counts field.
func (o *InlineResponse200312Items) SetCounts(v InlineResponse200312Counts) {
	o.Counts = &v
}

func (o InlineResponse200312Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200312Items struct {
	value *InlineResponse200312Items
	isSet bool
}

func (v NullableInlineResponse200312Items) Get() *InlineResponse200312Items {
	return v.value
}

func (v *NullableInlineResponse200312Items) Set(val *InlineResponse200312Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200312Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200312Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200312Items(val *InlineResponse200312Items) *NullableInlineResponse200312Items {
	return &NullableInlineResponse200312Items{value: val, isSet: true}
}

func (v NullableInlineResponse200312Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200312Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


