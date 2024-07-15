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

// OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies struct for OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies
type OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies struct {
	// Slot the power supply is in
	Slot *int32 `json:"slot,omitempty"`
	// Serial of the power supply
	Serial *string `json:"serial,omitempty"`
	// Model of the power supply
	Model *string `json:"model,omitempty"`
	// Status of the power supply
	Status *string `json:"status,omitempty"`
	Poe *OrganizationsOrganizationIdDevicesStatusesComponentsPoe `json:"poe,omitempty"`
}

// NewOrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies instantiates a new OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies() *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies {
	this := OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies{}
	return &this
}

// NewOrganizationsOrganizationIdDevicesStatusesComponentsPowerSuppliesWithDefaults instantiates a new OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdDevicesStatusesComponentsPowerSuppliesWithDefaults() *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies {
	this := OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies{}
	return &this
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) GetSlot() int32 {
	if o == nil || isNil(o.Slot) {
		var ret int32
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) GetSlotOk() (*int32, bool) {
	if o == nil || isNil(o.Slot) {
    return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) HasSlot() bool {
	if o != nil && !isNil(o.Slot) {
		return true
	}

	return false
}

// SetSlot gets a reference to the given int32 and assigns it to the Slot field.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) SetSlot(v int32) {
	o.Slot = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) SetModel(v string) {
	o.Model = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) SetStatus(v string) {
	o.Status = &v
}

// GetPoe returns the Poe field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) GetPoe() OrganizationsOrganizationIdDevicesStatusesComponentsPoe {
	if o == nil || isNil(o.Poe) {
		var ret OrganizationsOrganizationIdDevicesStatusesComponentsPoe
		return ret
	}
	return *o.Poe
}

// GetPoeOk returns a tuple with the Poe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) GetPoeOk() (*OrganizationsOrganizationIdDevicesStatusesComponentsPoe, bool) {
	if o == nil || isNil(o.Poe) {
    return nil, false
	}
	return o.Poe, true
}

// HasPoe returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) HasPoe() bool {
	if o != nil && !isNil(o.Poe) {
		return true
	}

	return false
}

// SetPoe gets a reference to the given OrganizationsOrganizationIdDevicesStatusesComponentsPoe and assigns it to the Poe field.
func (o *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) SetPoe(v OrganizationsOrganizationIdDevicesStatusesComponentsPoe) {
	o.Poe = &v
}

func (o OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Slot) {
		toSerialize["slot"] = o.Slot
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
	if !isNil(o.Poe) {
		toSerialize["poe"] = o.Poe
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies struct {
	value *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies
	isSet bool
}

func (v NullableOrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) Get() *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) Set(val *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies(val *OrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) *NullableOrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies {
	return &NullableOrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdDevicesStatusesComponentsPowerSupplies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


