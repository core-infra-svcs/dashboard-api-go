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

// InlineResponse20073 struct for InlineResponse20073
type InlineResponse20073 struct {
	// ID of the client
	Id *string `json:"id,omitempty"`
	// MAC address of the client
	Mac *string `json:"mac,omitempty"`
	// Network ID
	NetworkId *string `json:"networkId,omitempty"`
	// Name of the client
	Name *string `json:"name,omitempty"`
	// Bluetooth device name
	DeviceName *string `json:"deviceName,omitempty"`
	// Name of the manufacturer
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Epoch timestamp of the device's last appearance
	LastSeen *int32 `json:"lastSeen,omitempty"`
	// Seen by device MAC
	SeenByDeviceMac *string `json:"seenByDeviceMac,omitempty"`
	// Device in sight alert
	InSightAlert *bool `json:"inSightAlert,omitempty"`
	// Device out of sight alert
	OutOfSightAlert *bool `json:"outOfSightAlert,omitempty"`
	// A list of tags applied to the device
	Tags []string `json:"tags,omitempty"`
}

// NewInlineResponse20073 instantiates a new InlineResponse20073 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20073() *InlineResponse20073 {
	this := InlineResponse20073{}
	return &this
}

// NewInlineResponse20073WithDefaults instantiates a new InlineResponse20073 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20073WithDefaults() *InlineResponse20073 {
	this := InlineResponse20073{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20073) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20073) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20073) SetId(v string) {
	o.Id = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse20073) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse20073) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse20073) SetMac(v string) {
	o.Mac = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse20073) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse20073) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse20073) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20073) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20073) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20073) SetName(v string) {
	o.Name = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *InlineResponse20073) GetDeviceName() string {
	if o == nil || isNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetDeviceNameOk() (*string, bool) {
	if o == nil || isNil(o.DeviceName) {
    return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *InlineResponse20073) HasDeviceName() bool {
	if o != nil && !isNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *InlineResponse20073) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *InlineResponse20073) GetManufacturer() string {
	if o == nil || isNil(o.Manufacturer) {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetManufacturerOk() (*string, bool) {
	if o == nil || isNil(o.Manufacturer) {
    return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *InlineResponse20073) HasManufacturer() bool {
	if o != nil && !isNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *InlineResponse20073) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *InlineResponse20073) GetLastSeen() int32 {
	if o == nil || isNil(o.LastSeen) {
		var ret int32
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetLastSeenOk() (*int32, bool) {
	if o == nil || isNil(o.LastSeen) {
    return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *InlineResponse20073) HasLastSeen() bool {
	if o != nil && !isNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given int32 and assigns it to the LastSeen field.
func (o *InlineResponse20073) SetLastSeen(v int32) {
	o.LastSeen = &v
}

// GetSeenByDeviceMac returns the SeenByDeviceMac field value if set, zero value otherwise.
func (o *InlineResponse20073) GetSeenByDeviceMac() string {
	if o == nil || isNil(o.SeenByDeviceMac) {
		var ret string
		return ret
	}
	return *o.SeenByDeviceMac
}

// GetSeenByDeviceMacOk returns a tuple with the SeenByDeviceMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetSeenByDeviceMacOk() (*string, bool) {
	if o == nil || isNil(o.SeenByDeviceMac) {
    return nil, false
	}
	return o.SeenByDeviceMac, true
}

// HasSeenByDeviceMac returns a boolean if a field has been set.
func (o *InlineResponse20073) HasSeenByDeviceMac() bool {
	if o != nil && !isNil(o.SeenByDeviceMac) {
		return true
	}

	return false
}

// SetSeenByDeviceMac gets a reference to the given string and assigns it to the SeenByDeviceMac field.
func (o *InlineResponse20073) SetSeenByDeviceMac(v string) {
	o.SeenByDeviceMac = &v
}

// GetInSightAlert returns the InSightAlert field value if set, zero value otherwise.
func (o *InlineResponse20073) GetInSightAlert() bool {
	if o == nil || isNil(o.InSightAlert) {
		var ret bool
		return ret
	}
	return *o.InSightAlert
}

// GetInSightAlertOk returns a tuple with the InSightAlert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetInSightAlertOk() (*bool, bool) {
	if o == nil || isNil(o.InSightAlert) {
    return nil, false
	}
	return o.InSightAlert, true
}

// HasInSightAlert returns a boolean if a field has been set.
func (o *InlineResponse20073) HasInSightAlert() bool {
	if o != nil && !isNil(o.InSightAlert) {
		return true
	}

	return false
}

// SetInSightAlert gets a reference to the given bool and assigns it to the InSightAlert field.
func (o *InlineResponse20073) SetInSightAlert(v bool) {
	o.InSightAlert = &v
}

// GetOutOfSightAlert returns the OutOfSightAlert field value if set, zero value otherwise.
func (o *InlineResponse20073) GetOutOfSightAlert() bool {
	if o == nil || isNil(o.OutOfSightAlert) {
		var ret bool
		return ret
	}
	return *o.OutOfSightAlert
}

// GetOutOfSightAlertOk returns a tuple with the OutOfSightAlert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetOutOfSightAlertOk() (*bool, bool) {
	if o == nil || isNil(o.OutOfSightAlert) {
    return nil, false
	}
	return o.OutOfSightAlert, true
}

// HasOutOfSightAlert returns a boolean if a field has been set.
func (o *InlineResponse20073) HasOutOfSightAlert() bool {
	if o != nil && !isNil(o.OutOfSightAlert) {
		return true
	}

	return false
}

// SetOutOfSightAlert gets a reference to the given bool and assigns it to the OutOfSightAlert field.
func (o *InlineResponse20073) SetOutOfSightAlert(v bool) {
	o.OutOfSightAlert = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse20073) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse20073) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse20073) SetTags(v []string) {
	o.Tags = v
}

func (o InlineResponse20073) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.DeviceName) {
		toSerialize["deviceName"] = o.DeviceName
	}
	if !isNil(o.Manufacturer) {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if !isNil(o.LastSeen) {
		toSerialize["lastSeen"] = o.LastSeen
	}
	if !isNil(o.SeenByDeviceMac) {
		toSerialize["seenByDeviceMac"] = o.SeenByDeviceMac
	}
	if !isNil(o.InSightAlert) {
		toSerialize["inSightAlert"] = o.InSightAlert
	}
	if !isNil(o.OutOfSightAlert) {
		toSerialize["outOfSightAlert"] = o.OutOfSightAlert
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20073 struct {
	value *InlineResponse20073
	isSet bool
}

func (v NullableInlineResponse20073) Get() *InlineResponse20073 {
	return v.value
}

func (v *NullableInlineResponse20073) Set(val *InlineResponse20073) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20073) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20073) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20073(val *InlineResponse20073) *NullableInlineResponse20073 {
	return &NullableInlineResponse20073{value: val, isSet: true}
}

func (v NullableInlineResponse20073) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20073) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


