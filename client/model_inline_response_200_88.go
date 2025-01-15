/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20088 struct for InlineResponse20088
type InlineResponse20088 struct {
	// Name of the device
	Name *string `json:"name,omitempty"`
	// Latitude of the device
	Lat *float32 `json:"lat,omitempty"`
	// Longitude of the device
	Lng *float32 `json:"lng,omitempty"`
	// Physical address of the device
	Address *string `json:"address,omitempty"`
	// Notes for the device, limited to 255 characters
	Notes *string `json:"notes,omitempty"`
	// List of tags assigned to the device
	Tags []string `json:"tags,omitempty"`
	// ID of the network the device belongs to
	NetworkId *string `json:"networkId,omitempty"`
	// Serial number of the device
	Serial *string `json:"serial,omitempty"`
	// Model of the device
	Model *string `json:"model,omitempty"`
	// IMEI of the device, if applicable
	Imei *string `json:"imei,omitempty"`
	// MAC address of the device
	Mac *string `json:"mac,omitempty"`
	// LAN IP address of the device
	LanIp *string `json:"lanIp,omitempty"`
	// Firmware version of the device
	Firmware *string `json:"firmware,omitempty"`
	// Product type of the device
	ProductType *string `json:"productType,omitempty"`
	// Additional device information
	Details []InlineResponse2006Details `json:"details,omitempty"`
}

// NewInlineResponse20088 instantiates a new InlineResponse20088 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20088() *InlineResponse20088 {
	this := InlineResponse20088{}
	return &this
}

// NewInlineResponse20088WithDefaults instantiates a new InlineResponse20088 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20088WithDefaults() *InlineResponse20088 {
	this := InlineResponse20088{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20088) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20088) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20088) SetName(v string) {
	o.Name = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *InlineResponse20088) GetLat() float32 {
	if o == nil || isNil(o.Lat) {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetLatOk() (*float32, bool) {
	if o == nil || isNil(o.Lat) {
    return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *InlineResponse20088) HasLat() bool {
	if o != nil && !isNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *InlineResponse20088) SetLat(v float32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *InlineResponse20088) GetLng() float32 {
	if o == nil || isNil(o.Lng) {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetLngOk() (*float32, bool) {
	if o == nil || isNil(o.Lng) {
    return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *InlineResponse20088) HasLng() bool {
	if o != nil && !isNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *InlineResponse20088) SetLng(v float32) {
	o.Lng = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *InlineResponse20088) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *InlineResponse20088) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *InlineResponse20088) SetAddress(v string) {
	o.Address = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineResponse20088) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineResponse20088) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineResponse20088) SetNotes(v string) {
	o.Notes = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse20088) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse20088) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse20088) SetTags(v []string) {
	o.Tags = v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse20088) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse20088) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse20088) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse20088) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse20088) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse20088) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse20088) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse20088) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse20088) SetModel(v string) {
	o.Model = &v
}

// GetImei returns the Imei field value if set, zero value otherwise.
func (o *InlineResponse20088) GetImei() string {
	if o == nil || isNil(o.Imei) {
		var ret string
		return ret
	}
	return *o.Imei
}

// GetImeiOk returns a tuple with the Imei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetImeiOk() (*string, bool) {
	if o == nil || isNil(o.Imei) {
    return nil, false
	}
	return o.Imei, true
}

// HasImei returns a boolean if a field has been set.
func (o *InlineResponse20088) HasImei() bool {
	if o != nil && !isNil(o.Imei) {
		return true
	}

	return false
}

// SetImei gets a reference to the given string and assigns it to the Imei field.
func (o *InlineResponse20088) SetImei(v string) {
	o.Imei = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse20088) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse20088) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse20088) SetMac(v string) {
	o.Mac = &v
}

// GetLanIp returns the LanIp field value if set, zero value otherwise.
func (o *InlineResponse20088) GetLanIp() string {
	if o == nil || isNil(o.LanIp) {
		var ret string
		return ret
	}
	return *o.LanIp
}

// GetLanIpOk returns a tuple with the LanIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetLanIpOk() (*string, bool) {
	if o == nil || isNil(o.LanIp) {
    return nil, false
	}
	return o.LanIp, true
}

// HasLanIp returns a boolean if a field has been set.
func (o *InlineResponse20088) HasLanIp() bool {
	if o != nil && !isNil(o.LanIp) {
		return true
	}

	return false
}

// SetLanIp gets a reference to the given string and assigns it to the LanIp field.
func (o *InlineResponse20088) SetLanIp(v string) {
	o.LanIp = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *InlineResponse20088) GetFirmware() string {
	if o == nil || isNil(o.Firmware) {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetFirmwareOk() (*string, bool) {
	if o == nil || isNil(o.Firmware) {
    return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *InlineResponse20088) HasFirmware() bool {
	if o != nil && !isNil(o.Firmware) {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *InlineResponse20088) SetFirmware(v string) {
	o.Firmware = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *InlineResponse20088) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *InlineResponse20088) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *InlineResponse20088) SetProductType(v string) {
	o.ProductType = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *InlineResponse20088) GetDetails() []InlineResponse2006Details {
	if o == nil || isNil(o.Details) {
		var ret []InlineResponse2006Details
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetDetailsOk() ([]InlineResponse2006Details, bool) {
	if o == nil || isNil(o.Details) {
    return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *InlineResponse20088) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []InlineResponse2006Details and assigns it to the Details field.
func (o *InlineResponse20088) SetDetails(v []InlineResponse2006Details) {
	o.Details = v
}

func (o InlineResponse20088) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !isNil(o.Lng) {
		toSerialize["lng"] = o.Lng
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Imei) {
		toSerialize["imei"] = o.Imei
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.LanIp) {
		toSerialize["lanIp"] = o.LanIp
	}
	if !isNil(o.Firmware) {
		toSerialize["firmware"] = o.Firmware
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20088 struct {
	value *InlineResponse20088
	isSet bool
}

func (v NullableInlineResponse20088) Get() *InlineResponse20088 {
	return v.value
}

func (v *NullableInlineResponse20088) Set(val *InlineResponse20088) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20088) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20088) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20088(val *InlineResponse20088) *NullableInlineResponse20088 {
	return &NullableInlineResponse20088{value: val, isSet: true}
}

func (v NullableInlineResponse20088) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20088) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


