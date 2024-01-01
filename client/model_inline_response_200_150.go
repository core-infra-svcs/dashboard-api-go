/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200150 struct for InlineResponse200150
type InlineResponse200150 struct {
	// Serial of the device
	Serial *string `json:"serial,omitempty"`
	// Name assigned to the device
	Name *string `json:"name,omitempty"`
	// Status of the device upgrade
	DeviceStatus *string `json:"deviceStatus,omitempty"`
	Upgrade *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade `json:"upgrade,omitempty"`
}

// NewInlineResponse200150 instantiates a new InlineResponse200150 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200150() *InlineResponse200150 {
	this := InlineResponse200150{}
	return &this
}

// NewInlineResponse200150WithDefaults instantiates a new InlineResponse200150 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200150WithDefaults() *InlineResponse200150 {
	this := InlineResponse200150{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200150) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200150) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200150) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200150) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200150) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200150) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200150) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200150) SetName(v string) {
	o.Name = &v
}

// GetDeviceStatus returns the DeviceStatus field value if set, zero value otherwise.
func (o *InlineResponse200150) GetDeviceStatus() string {
	if o == nil || isNil(o.DeviceStatus) {
		var ret string
		return ret
	}
	return *o.DeviceStatus
}

// GetDeviceStatusOk returns a tuple with the DeviceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200150) GetDeviceStatusOk() (*string, bool) {
	if o == nil || isNil(o.DeviceStatus) {
    return nil, false
	}
	return o.DeviceStatus, true
}

// HasDeviceStatus returns a boolean if a field has been set.
func (o *InlineResponse200150) HasDeviceStatus() bool {
	if o != nil && !isNil(o.DeviceStatus) {
		return true
	}

	return false
}

// SetDeviceStatus gets a reference to the given string and assigns it to the DeviceStatus field.
func (o *InlineResponse200150) SetDeviceStatus(v string) {
	o.DeviceStatus = &v
}

// GetUpgrade returns the Upgrade field value if set, zero value otherwise.
func (o *InlineResponse200150) GetUpgrade() OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade {
	if o == nil || isNil(o.Upgrade) {
		var ret OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade
		return ret
	}
	return *o.Upgrade
}

// GetUpgradeOk returns a tuple with the Upgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200150) GetUpgradeOk() (*OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade, bool) {
	if o == nil || isNil(o.Upgrade) {
    return nil, false
	}
	return o.Upgrade, true
}

// HasUpgrade returns a boolean if a field has been set.
func (o *InlineResponse200150) HasUpgrade() bool {
	if o != nil && !isNil(o.Upgrade) {
		return true
	}

	return false
}

// SetUpgrade gets a reference to the given OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade and assigns it to the Upgrade field.
func (o *InlineResponse200150) SetUpgrade(v OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) {
	o.Upgrade = &v
}

func (o InlineResponse200150) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.DeviceStatus) {
		toSerialize["deviceStatus"] = o.DeviceStatus
	}
	if !isNil(o.Upgrade) {
		toSerialize["upgrade"] = o.Upgrade
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200150 struct {
	value *InlineResponse200150
	isSet bool
}

func (v NullableInlineResponse200150) Get() *InlineResponse200150 {
	return v.value
}

func (v *NullableInlineResponse200150) Set(val *InlineResponse200150) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200150) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200150) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200150(val *InlineResponse200150) *NullableInlineResponse200150 {
	return &NullableInlineResponse200150{value: val, isSet: true}
}

func (v NullableInlineResponse200150) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200150) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


