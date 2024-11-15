/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice Meraki Device properties
type OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice struct {
	// Device name
	Name *string `json:"name,omitempty"`
	// Device model
	Model *string `json:"model,omitempty"`
	// Device serial number
	Serial *string `json:"serial,omitempty"`
	// Device URL
	Url *string `json:"url,omitempty"`
	// Device status
	Status *string `json:"status,omitempty"`
}

// NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice() *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice {
	this := OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice{}
	return &this
}

// NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryDeviceWithDefaults instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryDeviceWithDefaults() *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice {
	this := OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) SetName(v string) {
	o.Name = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) SetModel(v string) {
	o.Model = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) SetSerial(v string) {
	o.Serial = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) SetUrl(v string) {
	o.Url = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) SetStatus(v string) {
	o.Status = &v
}

func (o OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice struct {
	value *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) Get() *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) Set(val *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice(val *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) *NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice {
	return &NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

