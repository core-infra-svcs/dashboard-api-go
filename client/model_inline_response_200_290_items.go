/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200290Items struct for InlineResponse200290Items
type InlineResponse200290Items struct {
	// Unique serial number for device.
	Serial *string `json:"serial,omitempty"`
	// Name of device.
	Name *string `json:"name,omitempty"`
	Network *InlineResponse200290Network `json:"network,omitempty"`
	// Status information for wireless access points.
	BasicServiceSets []InlineResponse200290BasicServiceSets `json:"basicServiceSets,omitempty"`
}

// NewInlineResponse200290Items instantiates a new InlineResponse200290Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200290Items() *InlineResponse200290Items {
	this := InlineResponse200290Items{}
	return &this
}

// NewInlineResponse200290ItemsWithDefaults instantiates a new InlineResponse200290Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200290ItemsWithDefaults() *InlineResponse200290Items {
	this := InlineResponse200290Items{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200290Items) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200290Items) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200290Items) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200290Items) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200290Items) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200290Items) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200290Items) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200290Items) SetName(v string) {
	o.Name = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200290Items) GetNetwork() InlineResponse200290Network {
	if o == nil || isNil(o.Network) {
		var ret InlineResponse200290Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200290Items) GetNetworkOk() (*InlineResponse200290Network, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200290Items) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given InlineResponse200290Network and assigns it to the Network field.
func (o *InlineResponse200290Items) SetNetwork(v InlineResponse200290Network) {
	o.Network = &v
}

// GetBasicServiceSets returns the BasicServiceSets field value if set, zero value otherwise.
func (o *InlineResponse200290Items) GetBasicServiceSets() []InlineResponse200290BasicServiceSets {
	if o == nil || isNil(o.BasicServiceSets) {
		var ret []InlineResponse200290BasicServiceSets
		return ret
	}
	return o.BasicServiceSets
}

// GetBasicServiceSetsOk returns a tuple with the BasicServiceSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200290Items) GetBasicServiceSetsOk() ([]InlineResponse200290BasicServiceSets, bool) {
	if o == nil || isNil(o.BasicServiceSets) {
    return nil, false
	}
	return o.BasicServiceSets, true
}

// HasBasicServiceSets returns a boolean if a field has been set.
func (o *InlineResponse200290Items) HasBasicServiceSets() bool {
	if o != nil && !isNil(o.BasicServiceSets) {
		return true
	}

	return false
}

// SetBasicServiceSets gets a reference to the given []InlineResponse200290BasicServiceSets and assigns it to the BasicServiceSets field.
func (o *InlineResponse200290Items) SetBasicServiceSets(v []InlineResponse200290BasicServiceSets) {
	o.BasicServiceSets = v
}

func (o InlineResponse200290Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.BasicServiceSets) {
		toSerialize["basicServiceSets"] = o.BasicServiceSets
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200290Items struct {
	value *InlineResponse200290Items
	isSet bool
}

func (v NullableInlineResponse200290Items) Get() *InlineResponse200290Items {
	return v.value
}

func (v *NullableInlineResponse200290Items) Set(val *InlineResponse200290Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200290Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200290Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200290Items(val *InlineResponse200290Items) *NullableInlineResponse200290Items {
	return &NullableInlineResponse200290Items{value: val, isSet: true}
}

func (v NullableInlineResponse200290Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200290Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


