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

// OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems struct for OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems
type OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems struct {
	Network *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceNetwork `json:"network,omitempty"`
	// Name of the device.
	Name *string `json:"name,omitempty"`
	// Unique serial number for device.
	Serial *string `json:"serial,omitempty"`
	// Model number of the device.
	Model *string `json:"model,omitempty"`
	RfProfile *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile `json:"rfProfile,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems instantiates a new OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems() *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems {
	this := OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItemsWithDefaults instantiates a new OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItemsWithDefaults() *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems {
	this := OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) GetNetwork() OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) GetNetworkOk() (*OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceNetwork and assigns it to the Network field.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) SetNetwork(v OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceNetwork) {
	o.Network = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) SetName(v string) {
	o.Name = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) SetModel(v string) {
	o.Model = &v
}

// GetRfProfile returns the RfProfile field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) GetRfProfile() OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile {
	if o == nil || isNil(o.RfProfile) {
		var ret OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile
		return ret
	}
	return *o.RfProfile
}

// GetRfProfileOk returns a tuple with the RfProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) GetRfProfileOk() (*OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile, bool) {
	if o == nil || isNil(o.RfProfile) {
    return nil, false
	}
	return o.RfProfile, true
}

// HasRfProfile returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) HasRfProfile() bool {
	if o != nil && !isNil(o.RfProfile) {
		return true
	}

	return false
}

// SetRfProfile gets a reference to the given OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile and assigns it to the RfProfile field.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) SetRfProfile(v OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceRfProfile) {
	o.RfProfile = &v
}

func (o OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.RfProfile) {
		toSerialize["rfProfile"] = o.RfProfile
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems struct {
	value *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) Get() *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) Set(val *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems(val *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems {
	return &NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


