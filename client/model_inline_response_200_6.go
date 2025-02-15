/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2006 struct for InlineResponse2006
type InlineResponse2006 struct {
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
	// MAC address of the device
	Mac *string `json:"mac,omitempty"`
	// LAN IP address of the device
	LanIp *string `json:"lanIp,omitempty"`
	// Firmware version of the device
	Firmware *string `json:"firmware,omitempty"`
	// The floor plan to associate to this device. null disassociates the device from the floorplan.
	FloorPlanId *string `json:"floorPlanId,omitempty"`
	// Additional device information
	Details []InlineResponse2006Details `json:"details,omitempty"`
	BeaconIdParams *InlineResponse2006BeaconIdParams `json:"beaconIdParams,omitempty"`
}

// NewInlineResponse2006 instantiates a new InlineResponse2006 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2006() *InlineResponse2006 {
	this := InlineResponse2006{}
	return &this
}

// NewInlineResponse2006WithDefaults instantiates a new InlineResponse2006 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2006WithDefaults() *InlineResponse2006 {
	this := InlineResponse2006{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse2006) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse2006) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse2006) SetName(v string) {
	o.Name = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *InlineResponse2006) GetLat() float32 {
	if o == nil || isNil(o.Lat) {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetLatOk() (*float32, bool) {
	if o == nil || isNil(o.Lat) {
    return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *InlineResponse2006) HasLat() bool {
	if o != nil && !isNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *InlineResponse2006) SetLat(v float32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *InlineResponse2006) GetLng() float32 {
	if o == nil || isNil(o.Lng) {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetLngOk() (*float32, bool) {
	if o == nil || isNil(o.Lng) {
    return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *InlineResponse2006) HasLng() bool {
	if o != nil && !isNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *InlineResponse2006) SetLng(v float32) {
	o.Lng = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *InlineResponse2006) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *InlineResponse2006) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *InlineResponse2006) SetAddress(v string) {
	o.Address = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineResponse2006) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineResponse2006) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineResponse2006) SetNotes(v string) {
	o.Notes = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse2006) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse2006) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse2006) SetTags(v []string) {
	o.Tags = v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse2006) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse2006) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse2006) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse2006) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse2006) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse2006) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse2006) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse2006) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse2006) SetModel(v string) {
	o.Model = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse2006) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse2006) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse2006) SetMac(v string) {
	o.Mac = &v
}

// GetLanIp returns the LanIp field value if set, zero value otherwise.
func (o *InlineResponse2006) GetLanIp() string {
	if o == nil || isNil(o.LanIp) {
		var ret string
		return ret
	}
	return *o.LanIp
}

// GetLanIpOk returns a tuple with the LanIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetLanIpOk() (*string, bool) {
	if o == nil || isNil(o.LanIp) {
    return nil, false
	}
	return o.LanIp, true
}

// HasLanIp returns a boolean if a field has been set.
func (o *InlineResponse2006) HasLanIp() bool {
	if o != nil && !isNil(o.LanIp) {
		return true
	}

	return false
}

// SetLanIp gets a reference to the given string and assigns it to the LanIp field.
func (o *InlineResponse2006) SetLanIp(v string) {
	o.LanIp = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *InlineResponse2006) GetFirmware() string {
	if o == nil || isNil(o.Firmware) {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetFirmwareOk() (*string, bool) {
	if o == nil || isNil(o.Firmware) {
    return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *InlineResponse2006) HasFirmware() bool {
	if o != nil && !isNil(o.Firmware) {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *InlineResponse2006) SetFirmware(v string) {
	o.Firmware = &v
}

// GetFloorPlanId returns the FloorPlanId field value if set, zero value otherwise.
func (o *InlineResponse2006) GetFloorPlanId() string {
	if o == nil || isNil(o.FloorPlanId) {
		var ret string
		return ret
	}
	return *o.FloorPlanId
}

// GetFloorPlanIdOk returns a tuple with the FloorPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetFloorPlanIdOk() (*string, bool) {
	if o == nil || isNil(o.FloorPlanId) {
    return nil, false
	}
	return o.FloorPlanId, true
}

// HasFloorPlanId returns a boolean if a field has been set.
func (o *InlineResponse2006) HasFloorPlanId() bool {
	if o != nil && !isNil(o.FloorPlanId) {
		return true
	}

	return false
}

// SetFloorPlanId gets a reference to the given string and assigns it to the FloorPlanId field.
func (o *InlineResponse2006) SetFloorPlanId(v string) {
	o.FloorPlanId = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *InlineResponse2006) GetDetails() []InlineResponse2006Details {
	if o == nil || isNil(o.Details) {
		var ret []InlineResponse2006Details
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetDetailsOk() ([]InlineResponse2006Details, bool) {
	if o == nil || isNil(o.Details) {
    return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *InlineResponse2006) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []InlineResponse2006Details and assigns it to the Details field.
func (o *InlineResponse2006) SetDetails(v []InlineResponse2006Details) {
	o.Details = v
}

// GetBeaconIdParams returns the BeaconIdParams field value if set, zero value otherwise.
func (o *InlineResponse2006) GetBeaconIdParams() InlineResponse2006BeaconIdParams {
	if o == nil || isNil(o.BeaconIdParams) {
		var ret InlineResponse2006BeaconIdParams
		return ret
	}
	return *o.BeaconIdParams
}

// GetBeaconIdParamsOk returns a tuple with the BeaconIdParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetBeaconIdParamsOk() (*InlineResponse2006BeaconIdParams, bool) {
	if o == nil || isNil(o.BeaconIdParams) {
    return nil, false
	}
	return o.BeaconIdParams, true
}

// HasBeaconIdParams returns a boolean if a field has been set.
func (o *InlineResponse2006) HasBeaconIdParams() bool {
	if o != nil && !isNil(o.BeaconIdParams) {
		return true
	}

	return false
}

// SetBeaconIdParams gets a reference to the given InlineResponse2006BeaconIdParams and assigns it to the BeaconIdParams field.
func (o *InlineResponse2006) SetBeaconIdParams(v InlineResponse2006BeaconIdParams) {
	o.BeaconIdParams = &v
}

func (o InlineResponse2006) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.LanIp) {
		toSerialize["lanIp"] = o.LanIp
	}
	if !isNil(o.Firmware) {
		toSerialize["firmware"] = o.Firmware
	}
	if !isNil(o.FloorPlanId) {
		toSerialize["floorPlanId"] = o.FloorPlanId
	}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !isNil(o.BeaconIdParams) {
		toSerialize["beaconIdParams"] = o.BeaconIdParams
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2006 struct {
	value *InlineResponse2006
	isSet bool
}

func (v NullableInlineResponse2006) Get() *InlineResponse2006 {
	return v.value
}

func (v *NullableInlineResponse2006) Set(val *InlineResponse2006) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2006) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2006) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2006(val *InlineResponse2006) *NullableInlineResponse2006 {
	return &NullableInlineResponse2006{value: val, isSet: true}
}

func (v NullableInlineResponse2006) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2006) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


