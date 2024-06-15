/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdAssuranceAlertsScopeDevices struct for OrganizationsOrganizationIdAssuranceAlertsScopeDevices
type OrganizationsOrganizationIdAssuranceAlertsScopeDevices struct {
	// URL of affected device
	Url *string `json:"url,omitempty"`
	// Name of affected device
	Name *string `json:"name,omitempty"`
	// Order of affected device in array
	Order *int32 `json:"order,omitempty"`
	// Type of affected device
	ProductType *string `json:"productType,omitempty"`
	// Serial of affected device
	Serial *string `json:"serial,omitempty"`
	// MAC address of affected device
	Mac *string `json:"mac,omitempty"`
	// IMEI of affected device
	Imei *string `json:"imei,omitempty"`
	Lldp *OrganizationsOrganizationIdAssuranceAlertsScopeLldp `json:"lldp,omitempty"`
}

// NewOrganizationsOrganizationIdAssuranceAlertsScopeDevices instantiates a new OrganizationsOrganizationIdAssuranceAlertsScopeDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdAssuranceAlertsScopeDevices() *OrganizationsOrganizationIdAssuranceAlertsScopeDevices {
	this := OrganizationsOrganizationIdAssuranceAlertsScopeDevices{}
	return &this
}

// NewOrganizationsOrganizationIdAssuranceAlertsScopeDevicesWithDefaults instantiates a new OrganizationsOrganizationIdAssuranceAlertsScopeDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdAssuranceAlertsScopeDevicesWithDefaults() *OrganizationsOrganizationIdAssuranceAlertsScopeDevices {
	this := OrganizationsOrganizationIdAssuranceAlertsScopeDevices{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) SetUrl(v string) {
	o.Url = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) SetName(v string) {
	o.Name = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetOrder() int32 {
	if o == nil || isNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetOrderOk() (*int32, bool) {
	if o == nil || isNil(o.Order) {
    return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) HasOrder() bool {
	if o != nil && !isNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) SetOrder(v int32) {
	o.Order = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) SetProductType(v string) {
	o.ProductType = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) SetMac(v string) {
	o.Mac = &v
}

// GetImei returns the Imei field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetImei() string {
	if o == nil || isNil(o.Imei) {
		var ret string
		return ret
	}
	return *o.Imei
}

// GetImeiOk returns a tuple with the Imei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetImeiOk() (*string, bool) {
	if o == nil || isNil(o.Imei) {
    return nil, false
	}
	return o.Imei, true
}

// HasImei returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) HasImei() bool {
	if o != nil && !isNil(o.Imei) {
		return true
	}

	return false
}

// SetImei gets a reference to the given string and assigns it to the Imei field.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) SetImei(v string) {
	o.Imei = &v
}

// GetLldp returns the Lldp field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetLldp() OrganizationsOrganizationIdAssuranceAlertsScopeLldp {
	if o == nil || isNil(o.Lldp) {
		var ret OrganizationsOrganizationIdAssuranceAlertsScopeLldp
		return ret
	}
	return *o.Lldp
}

// GetLldpOk returns a tuple with the Lldp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetLldpOk() (*OrganizationsOrganizationIdAssuranceAlertsScopeLldp, bool) {
	if o == nil || isNil(o.Lldp) {
    return nil, false
	}
	return o.Lldp, true
}

// HasLldp returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) HasLldp() bool {
	if o != nil && !isNil(o.Lldp) {
		return true
	}

	return false
}

// SetLldp gets a reference to the given OrganizationsOrganizationIdAssuranceAlertsScopeLldp and assigns it to the Lldp field.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) SetLldp(v OrganizationsOrganizationIdAssuranceAlertsScopeLldp) {
	o.Lldp = &v
}

func (o OrganizationsOrganizationIdAssuranceAlertsScopeDevices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Imei) {
		toSerialize["imei"] = o.Imei
	}
	if !isNil(o.Lldp) {
		toSerialize["lldp"] = o.Lldp
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdAssuranceAlertsScopeDevices struct {
	value *OrganizationsOrganizationIdAssuranceAlertsScopeDevices
	isSet bool
}

func (v NullableOrganizationsOrganizationIdAssuranceAlertsScopeDevices) Get() *OrganizationsOrganizationIdAssuranceAlertsScopeDevices {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdAssuranceAlertsScopeDevices) Set(val *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdAssuranceAlertsScopeDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdAssuranceAlertsScopeDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdAssuranceAlertsScopeDevices(val *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) *NullableOrganizationsOrganizationIdAssuranceAlertsScopeDevices {
	return &NullableOrganizationsOrganizationIdAssuranceAlertsScopeDevices{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdAssuranceAlertsScopeDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdAssuranceAlertsScopeDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


