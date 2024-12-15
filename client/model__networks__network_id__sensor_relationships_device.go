/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSensorRelationshipsDevice A sensor or gateway device in the network
type NetworksNetworkIdSensorRelationshipsDevice struct {
	// The name of the device
	Name *string `json:"name,omitempty"`
	// The serial of the device
	Serial *string `json:"serial,omitempty"`
	// The product type of the device
	ProductType *string `json:"productType,omitempty"`
}

// NewNetworksNetworkIdSensorRelationshipsDevice instantiates a new NetworksNetworkIdSensorRelationshipsDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorRelationshipsDevice() *NetworksNetworkIdSensorRelationshipsDevice {
	this := NetworksNetworkIdSensorRelationshipsDevice{}
	return &this
}

// NewNetworksNetworkIdSensorRelationshipsDeviceWithDefaults instantiates a new NetworksNetworkIdSensorRelationshipsDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorRelationshipsDeviceWithDefaults() *NetworksNetworkIdSensorRelationshipsDevice {
	this := NetworksNetworkIdSensorRelationshipsDevice{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorRelationshipsDevice) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorRelationshipsDevice) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorRelationshipsDevice) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdSensorRelationshipsDevice) SetName(v string) {
	o.Name = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorRelationshipsDevice) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorRelationshipsDevice) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorRelationshipsDevice) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *NetworksNetworkIdSensorRelationshipsDevice) SetSerial(v string) {
	o.Serial = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorRelationshipsDevice) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorRelationshipsDevice) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorRelationshipsDevice) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *NetworksNetworkIdSensorRelationshipsDevice) SetProductType(v string) {
	o.ProductType = &v
}

func (o NetworksNetworkIdSensorRelationshipsDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorRelationshipsDevice struct {
	value *NetworksNetworkIdSensorRelationshipsDevice
	isSet bool
}

func (v NullableNetworksNetworkIdSensorRelationshipsDevice) Get() *NetworksNetworkIdSensorRelationshipsDevice {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorRelationshipsDevice) Set(val *NetworksNetworkIdSensorRelationshipsDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorRelationshipsDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorRelationshipsDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorRelationshipsDevice(val *NetworksNetworkIdSensorRelationshipsDevice) *NullableNetworksNetworkIdSensorRelationshipsDevice {
	return &NullableNetworksNetworkIdSensorRelationshipsDevice{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorRelationshipsDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorRelationshipsDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


