/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand struct for OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand
type OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand struct {
	// The band for the given metrics.
	Band *string `json:"band,omitempty"`
	Wifi *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi `json:"wifi,omitempty"`
	NonWifi *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNonWifi `json:"nonWifi,omitempty"`
	Total *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceTotal `json:"total,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand instantiates a new OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand() *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand {
	this := OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBandWithDefaults instantiates a new OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBandWithDefaults() *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand {
	this := OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand{}
	return &this
}

// GetBand returns the Band field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) GetBand() string {
	if o == nil || isNil(o.Band) {
		var ret string
		return ret
	}
	return *o.Band
}

// GetBandOk returns a tuple with the Band field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) GetBandOk() (*string, bool) {
	if o == nil || isNil(o.Band) {
    return nil, false
	}
	return o.Band, true
}

// HasBand returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) HasBand() bool {
	if o != nil && !isNil(o.Band) {
		return true
	}

	return false
}

// SetBand gets a reference to the given string and assigns it to the Band field.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) SetBand(v string) {
	o.Band = &v
}

// GetWifi returns the Wifi field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) GetWifi() OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi {
	if o == nil || isNil(o.Wifi) {
		var ret OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi
		return ret
	}
	return *o.Wifi
}

// GetWifiOk returns a tuple with the Wifi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) GetWifiOk() (*OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi, bool) {
	if o == nil || isNil(o.Wifi) {
    return nil, false
	}
	return o.Wifi, true
}

// HasWifi returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) HasWifi() bool {
	if o != nil && !isNil(o.Wifi) {
		return true
	}

	return false
}

// SetWifi gets a reference to the given OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi and assigns it to the Wifi field.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) SetWifi(v OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi) {
	o.Wifi = &v
}

// GetNonWifi returns the NonWifi field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) GetNonWifi() OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNonWifi {
	if o == nil || isNil(o.NonWifi) {
		var ret OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNonWifi
		return ret
	}
	return *o.NonWifi
}

// GetNonWifiOk returns a tuple with the NonWifi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) GetNonWifiOk() (*OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNonWifi, bool) {
	if o == nil || isNil(o.NonWifi) {
    return nil, false
	}
	return o.NonWifi, true
}

// HasNonWifi returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) HasNonWifi() bool {
	if o != nil && !isNil(o.NonWifi) {
		return true
	}

	return false
}

// SetNonWifi gets a reference to the given OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNonWifi and assigns it to the NonWifi field.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) SetNonWifi(v OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNonWifi) {
	o.NonWifi = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) GetTotal() OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceTotal {
	if o == nil || isNil(o.Total) {
		var ret OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceTotal
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) GetTotalOk() (*OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceTotal, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceTotal and assigns it to the Total field.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) SetTotal(v OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceTotal) {
	o.Total = &v
}

func (o OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Band) {
		toSerialize["band"] = o.Band
	}
	if !isNil(o.Wifi) {
		toSerialize["wifi"] = o.Wifi
	}
	if !isNil(o.NonWifi) {
		toSerialize["nonWifi"] = o.NonWifi
	}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand struct {
	value *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) Get() *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) Set(val *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand(val *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) *NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand {
	return &NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


