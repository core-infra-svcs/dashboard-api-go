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

// OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots struct for OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots
type OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots struct {
	// Which slot the AC power supply occupies. Possible values are: 0, 1, 2.
	Number *int32 `json:"number,omitempty"`
	// The power supply unit serial number.
	Serial *string `json:"serial,omitempty"`
	// The power supply unit model.
	Model *string `json:"model,omitempty"`
	// Status of the power supply unit. Possible values are: connected, not connected, powering.
	Status *string `json:"status,omitempty"`
}

// NewOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots instantiates a new OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots() *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots {
	this := OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots{}
	return &this
}

// NewOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlotsWithDefaults instantiates a new OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlotsWithDefaults() *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots {
	this := OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) GetNumber() int32 {
	if o == nil || isNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) GetNumberOk() (*int32, bool) {
	if o == nil || isNil(o.Number) {
    return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) HasNumber() bool {
	if o != nil && !isNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) SetNumber(v int32) {
	o.Number = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) SetModel(v string) {
	o.Model = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) SetStatus(v string) {
	o.Status = &v
}

func (o OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots struct {
	value *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots
	isSet bool
}

func (v NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) Get() *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) Set(val *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots(val *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) *NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots {
	return &NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


