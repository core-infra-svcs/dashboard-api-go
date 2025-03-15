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

// InlineResponse200337 struct for InlineResponse200337
type InlineResponse200337 struct {
	// The serial number of the AP
	Serial *string `json:"serial,omitempty"`
	// The name of the AP
	Name *string `json:"name,omitempty"`
	Network *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork `json:"network,omitempty"`
	Power *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower `json:"power,omitempty"`
	// List of port details
	Ports []OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts `json:"ports,omitempty"`
	Aggregation *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation `json:"aggregation,omitempty"`
}

// NewInlineResponse200337 instantiates a new InlineResponse200337 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200337() *InlineResponse200337 {
	this := InlineResponse200337{}
	return &this
}

// NewInlineResponse200337WithDefaults instantiates a new InlineResponse200337 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200337WithDefaults() *InlineResponse200337 {
	this := InlineResponse200337{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200337) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200337) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200337) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200337) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200337) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200337) SetName(v string) {
	o.Name = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200337) GetNetwork() OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337) GetNetworkOk() (*OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200337) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork and assigns it to the Network field.
func (o *InlineResponse200337) SetNetwork(v OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork) {
	o.Network = &v
}

// GetPower returns the Power field value if set, zero value otherwise.
func (o *InlineResponse200337) GetPower() OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower {
	if o == nil || isNil(o.Power) {
		var ret OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower
		return ret
	}
	return *o.Power
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337) GetPowerOk() (*OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower, bool) {
	if o == nil || isNil(o.Power) {
    return nil, false
	}
	return o.Power, true
}

// HasPower returns a boolean if a field has been set.
func (o *InlineResponse200337) HasPower() bool {
	if o != nil && !isNil(o.Power) {
		return true
	}

	return false
}

// SetPower gets a reference to the given OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower and assigns it to the Power field.
func (o *InlineResponse200337) SetPower(v OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) {
	o.Power = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *InlineResponse200337) GetPorts() []OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts {
	if o == nil || isNil(o.Ports) {
		var ret []OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337) GetPortsOk() ([]OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts, bool) {
	if o == nil || isNil(o.Ports) {
    return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *InlineResponse200337) HasPorts() bool {
	if o != nil && !isNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts and assigns it to the Ports field.
func (o *InlineResponse200337) SetPorts(v []OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) {
	o.Ports = v
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *InlineResponse200337) GetAggregation() OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation {
	if o == nil || isNil(o.Aggregation) {
		var ret OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337) GetAggregationOk() (*OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation, bool) {
	if o == nil || isNil(o.Aggregation) {
    return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *InlineResponse200337) HasAggregation() bool {
	if o != nil && !isNil(o.Aggregation) {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation and assigns it to the Aggregation field.
func (o *InlineResponse200337) SetAggregation(v OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation) {
	o.Aggregation = &v
}

func (o InlineResponse200337) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Power) {
		toSerialize["power"] = o.Power
	}
	if !isNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !isNil(o.Aggregation) {
		toSerialize["aggregation"] = o.Aggregation
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200337 struct {
	value *InlineResponse200337
	isSet bool
}

func (v NullableInlineResponse200337) Get() *InlineResponse200337 {
	return v.value
}

func (v *NullableInlineResponse200337) Set(val *InlineResponse200337) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200337) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200337) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200337(val *InlineResponse200337) *NullableInlineResponse200337 {
	return &NullableInlineResponse200337{value: val, isSet: true}
}

func (v NullableInlineResponse200337) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200337) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


