/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems struct for OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems
type OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems struct {
	Device *OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice `json:"device,omitempty"`
	// Whether eSIM is currently active SIM on Device
	Active *bool `json:"active,omitempty"`
	// eSIM EID
	Eid *string `json:"eid,omitempty"`
	// Last update of eSIM
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty"`
	Network *OrganizationsOrganizationIdCellularGatewayEsimsInventoryNetwork `json:"network,omitempty"`
	// eSIM Profile Information
	Profiles []OrganizationsOrganizationIdCellularGatewayEsimsInventoryProfiles `json:"profiles,omitempty"`
}

// NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryItems instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryItems() *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems {
	this := OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems{}
	return &this
}

// NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryItemsWithDefaults instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryItemsWithDefaults() *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems {
	this := OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetDevice() OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice {
	if o == nil || isNil(o.Device) {
		var ret OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetDeviceOk() (*OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice, bool) {
	if o == nil || isNil(o.Device) {
    return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice and assigns it to the Device field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) SetDevice(v OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice) {
	o.Device = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
    return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) SetActive(v bool) {
	o.Active = &v
}

// GetEid returns the Eid field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetEid() string {
	if o == nil || isNil(o.Eid) {
		var ret string
		return ret
	}
	return *o.Eid
}

// GetEidOk returns a tuple with the Eid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetEidOk() (*string, bool) {
	if o == nil || isNil(o.Eid) {
    return nil, false
	}
	return o.Eid, true
}

// HasEid returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) HasEid() bool {
	if o != nil && !isNil(o.Eid) {
		return true
	}

	return false
}

// SetEid gets a reference to the given string and assigns it to the Eid field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) SetEid(v string) {
	o.Eid = &v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetLastUpdatedAt() string {
	if o == nil || isNil(o.LastUpdatedAt) {
		var ret string
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetLastUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.LastUpdatedAt) {
    return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) HasLastUpdatedAt() bool {
	if o != nil && !isNil(o.LastUpdatedAt) {
		return true
	}

	return false
}

// SetLastUpdatedAt gets a reference to the given string and assigns it to the LastUpdatedAt field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) SetLastUpdatedAt(v string) {
	o.LastUpdatedAt = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetNetwork() OrganizationsOrganizationIdCellularGatewayEsimsInventoryNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdCellularGatewayEsimsInventoryNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetNetworkOk() (*OrganizationsOrganizationIdCellularGatewayEsimsInventoryNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdCellularGatewayEsimsInventoryNetwork and assigns it to the Network field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) SetNetwork(v OrganizationsOrganizationIdCellularGatewayEsimsInventoryNetwork) {
	o.Network = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetProfiles() []OrganizationsOrganizationIdCellularGatewayEsimsInventoryProfiles {
	if o == nil || isNil(o.Profiles) {
		var ret []OrganizationsOrganizationIdCellularGatewayEsimsInventoryProfiles
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetProfilesOk() ([]OrganizationsOrganizationIdCellularGatewayEsimsInventoryProfiles, bool) {
	if o == nil || isNil(o.Profiles) {
    return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) HasProfiles() bool {
	if o != nil && !isNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []OrganizationsOrganizationIdCellularGatewayEsimsInventoryProfiles and assigns it to the Profiles field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) SetProfiles(v []OrganizationsOrganizationIdCellularGatewayEsimsInventoryProfiles) {
	o.Profiles = v
}

func (o OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.Eid) {
		toSerialize["eid"] = o.Eid
	}
	if !isNil(o.LastUpdatedAt) {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Profiles) {
		toSerialize["profiles"] = o.Profiles
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryItems struct {
	value *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) Get() *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) Set(val *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryItems(val *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) *NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryItems {
	return &NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryItems{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


