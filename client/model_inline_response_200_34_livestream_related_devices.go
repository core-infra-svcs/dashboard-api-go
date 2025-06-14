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

// InlineResponse20034LivestreamRelatedDevices struct for InlineResponse20034LivestreamRelatedDevices
type InlineResponse20034LivestreamRelatedDevices struct {
	// The serial of the related device
	Serial *string `json:"serial,omitempty"`
	// The product type of the related device
	ProductType *string `json:"productType,omitempty"`
}

// NewInlineResponse20034LivestreamRelatedDevices instantiates a new InlineResponse20034LivestreamRelatedDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20034LivestreamRelatedDevices() *InlineResponse20034LivestreamRelatedDevices {
	this := InlineResponse20034LivestreamRelatedDevices{}
	return &this
}

// NewInlineResponse20034LivestreamRelatedDevicesWithDefaults instantiates a new InlineResponse20034LivestreamRelatedDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20034LivestreamRelatedDevicesWithDefaults() *InlineResponse20034LivestreamRelatedDevices {
	this := InlineResponse20034LivestreamRelatedDevices{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse20034LivestreamRelatedDevices) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20034LivestreamRelatedDevices) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse20034LivestreamRelatedDevices) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse20034LivestreamRelatedDevices) SetSerial(v string) {
	o.Serial = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *InlineResponse20034LivestreamRelatedDevices) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20034LivestreamRelatedDevices) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *InlineResponse20034LivestreamRelatedDevices) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *InlineResponse20034LivestreamRelatedDevices) SetProductType(v string) {
	o.ProductType = &v
}

func (o InlineResponse20034LivestreamRelatedDevices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20034LivestreamRelatedDevices struct {
	value *InlineResponse20034LivestreamRelatedDevices
	isSet bool
}

func (v NullableInlineResponse20034LivestreamRelatedDevices) Get() *InlineResponse20034LivestreamRelatedDevices {
	return v.value
}

func (v *NullableInlineResponse20034LivestreamRelatedDevices) Set(val *InlineResponse20034LivestreamRelatedDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20034LivestreamRelatedDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20034LivestreamRelatedDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20034LivestreamRelatedDevices(val *InlineResponse20034LivestreamRelatedDevices) *NullableInlineResponse20034LivestreamRelatedDevices {
	return &NullableInlineResponse20034LivestreamRelatedDevices{value: val, isSet: true}
}

func (v NullableInlineResponse20034LivestreamRelatedDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20034LivestreamRelatedDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


