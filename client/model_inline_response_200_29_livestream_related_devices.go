/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20029LivestreamRelatedDevices struct for InlineResponse20029LivestreamRelatedDevices
type InlineResponse20029LivestreamRelatedDevices struct {
	// The serial of the related device
	Serial *string `json:"serial,omitempty"`
	// The product type of the related device
	ProductType *string `json:"productType,omitempty"`
}

// NewInlineResponse20029LivestreamRelatedDevices instantiates a new InlineResponse20029LivestreamRelatedDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20029LivestreamRelatedDevices() *InlineResponse20029LivestreamRelatedDevices {
	this := InlineResponse20029LivestreamRelatedDevices{}
	return &this
}

// NewInlineResponse20029LivestreamRelatedDevicesWithDefaults instantiates a new InlineResponse20029LivestreamRelatedDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20029LivestreamRelatedDevicesWithDefaults() *InlineResponse20029LivestreamRelatedDevices {
	this := InlineResponse20029LivestreamRelatedDevices{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse20029LivestreamRelatedDevices) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029LivestreamRelatedDevices) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse20029LivestreamRelatedDevices) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse20029LivestreamRelatedDevices) SetSerial(v string) {
	o.Serial = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *InlineResponse20029LivestreamRelatedDevices) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029LivestreamRelatedDevices) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *InlineResponse20029LivestreamRelatedDevices) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *InlineResponse20029LivestreamRelatedDevices) SetProductType(v string) {
	o.ProductType = &v
}

func (o InlineResponse20029LivestreamRelatedDevices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20029LivestreamRelatedDevices struct {
	value *InlineResponse20029LivestreamRelatedDevices
	isSet bool
}

func (v NullableInlineResponse20029LivestreamRelatedDevices) Get() *InlineResponse20029LivestreamRelatedDevices {
	return v.value
}

func (v *NullableInlineResponse20029LivestreamRelatedDevices) Set(val *InlineResponse20029LivestreamRelatedDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20029LivestreamRelatedDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20029LivestreamRelatedDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20029LivestreamRelatedDevices(val *InlineResponse20029LivestreamRelatedDevices) *NullableInlineResponse20029LivestreamRelatedDevices {
	return &NullableInlineResponse20029LivestreamRelatedDevices{value: val, isSet: true}
}

func (v NullableInlineResponse20029LivestreamRelatedDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20029LivestreamRelatedDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


