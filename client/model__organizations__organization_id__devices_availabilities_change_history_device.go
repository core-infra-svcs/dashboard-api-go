/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice Device information
type OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice struct {
	// Device serial number
	Serial *string `json:"serial,omitempty"`
	// Device name
	Name *string `json:"name,omitempty"`
	// Device product type.
	ProductType *string `json:"productType,omitempty"`
	// Device model
	Model *string `json:"model,omitempty"`
}

// NewOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice instantiates a new OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice() *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice {
	this := OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice{}
	return &this
}

// NewOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDeviceWithDefaults instantiates a new OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDeviceWithDefaults() *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice {
	this := OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) SetName(v string) {
	o.Name = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) SetProductType(v string) {
	o.ProductType = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) SetModel(v string) {
	o.Model = &v
}

func (o OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice struct {
	value *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice
	isSet bool
}

func (v NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) Get() *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) Set(val *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice(val *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) *NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice {
	return &NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

