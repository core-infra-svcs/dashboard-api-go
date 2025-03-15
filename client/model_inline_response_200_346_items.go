/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200346Items struct for InlineResponse200346Items
type InlineResponse200346Items struct {
	// Unique serial number for device.
	Serial *string `json:"serial,omitempty"`
	// Name of device.
	Name *string `json:"name,omitempty"`
	Network *InlineResponse200346Network `json:"network,omitempty"`
	// Status information for wireless access points.
	BasicServiceSets []InlineResponse200346BasicServiceSets `json:"basicServiceSets,omitempty"`
}

// NewInlineResponse200346Items instantiates a new InlineResponse200346Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200346Items() *InlineResponse200346Items {
	this := InlineResponse200346Items{}
	return &this
}

// NewInlineResponse200346ItemsWithDefaults instantiates a new InlineResponse200346Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200346ItemsWithDefaults() *InlineResponse200346Items {
	this := InlineResponse200346Items{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200346Items) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200346Items) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200346Items) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200346Items) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200346Items) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200346Items) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200346Items) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200346Items) SetName(v string) {
	o.Name = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200346Items) GetNetwork() InlineResponse200346Network {
	if o == nil || isNil(o.Network) {
		var ret InlineResponse200346Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200346Items) GetNetworkOk() (*InlineResponse200346Network, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200346Items) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given InlineResponse200346Network and assigns it to the Network field.
func (o *InlineResponse200346Items) SetNetwork(v InlineResponse200346Network) {
	o.Network = &v
}

// GetBasicServiceSets returns the BasicServiceSets field value if set, zero value otherwise.
func (o *InlineResponse200346Items) GetBasicServiceSets() []InlineResponse200346BasicServiceSets {
	if o == nil || isNil(o.BasicServiceSets) {
		var ret []InlineResponse200346BasicServiceSets
		return ret
	}
	return o.BasicServiceSets
}

// GetBasicServiceSetsOk returns a tuple with the BasicServiceSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200346Items) GetBasicServiceSetsOk() ([]InlineResponse200346BasicServiceSets, bool) {
	if o == nil || isNil(o.BasicServiceSets) {
    return nil, false
	}
	return o.BasicServiceSets, true
}

// HasBasicServiceSets returns a boolean if a field has been set.
func (o *InlineResponse200346Items) HasBasicServiceSets() bool {
	if o != nil && !isNil(o.BasicServiceSets) {
		return true
	}

	return false
}

// SetBasicServiceSets gets a reference to the given []InlineResponse200346BasicServiceSets and assigns it to the BasicServiceSets field.
func (o *InlineResponse200346Items) SetBasicServiceSets(v []InlineResponse200346BasicServiceSets) {
	o.BasicServiceSets = v
}

func (o InlineResponse200346Items) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200346Items struct {
	value *InlineResponse200346Items
	isSet bool
}

func (v NullableInlineResponse200346Items) Get() *InlineResponse200346Items {
	return v.value
}

func (v *NullableInlineResponse200346Items) Set(val *InlineResponse200346Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200346Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200346Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200346Items(val *InlineResponse200346Items) *NullableInlineResponse200346Items {
	return &NullableInlineResponse200346Items{value: val, isSet: true}
}

func (v NullableInlineResponse200346Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200346Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


