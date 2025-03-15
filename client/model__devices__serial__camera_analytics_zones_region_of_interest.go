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

// DevicesSerialCameraAnalyticsZonesRegionOfInterest The region of interest
type DevicesSerialCameraAnalyticsZonesRegionOfInterest struct {
	// The x0 coordinate
	X0 *string `json:"x0,omitempty"`
	// The y0 coordinate
	Y0 *string `json:"y0,omitempty"`
	// The x1 coordinate
	X1 *string `json:"x1,omitempty"`
	// The y1 coordinate
	Y1 *string `json:"y1,omitempty"`
}

// NewDevicesSerialCameraAnalyticsZonesRegionOfInterest instantiates a new DevicesSerialCameraAnalyticsZonesRegionOfInterest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialCameraAnalyticsZonesRegionOfInterest() *DevicesSerialCameraAnalyticsZonesRegionOfInterest {
	this := DevicesSerialCameraAnalyticsZonesRegionOfInterest{}
	return &this
}

// NewDevicesSerialCameraAnalyticsZonesRegionOfInterestWithDefaults instantiates a new DevicesSerialCameraAnalyticsZonesRegionOfInterest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialCameraAnalyticsZonesRegionOfInterestWithDefaults() *DevicesSerialCameraAnalyticsZonesRegionOfInterest {
	this := DevicesSerialCameraAnalyticsZonesRegionOfInterest{}
	return &this
}

// GetX0 returns the X0 field value if set, zero value otherwise.
func (o *DevicesSerialCameraAnalyticsZonesRegionOfInterest) GetX0() string {
	if o == nil || isNil(o.X0) {
		var ret string
		return ret
	}
	return *o.X0
}

// GetX0Ok returns a tuple with the X0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCameraAnalyticsZonesRegionOfInterest) GetX0Ok() (*string, bool) {
	if o == nil || isNil(o.X0) {
    return nil, false
	}
	return o.X0, true
}

// HasX0 returns a boolean if a field has been set.
func (o *DevicesSerialCameraAnalyticsZonesRegionOfInterest) HasX0() bool {
	if o != nil && !isNil(o.X0) {
		return true
	}

	return false
}

// SetX0 gets a reference to the given string and assigns it to the X0 field.
func (o *DevicesSerialCameraAnalyticsZonesRegionOfInterest) SetX0(v string) {
	o.X0 = &v
}

// GetY0 returns the Y0 field value if set, zero value otherwise.
func (o *DevicesSerialCameraAnalyticsZonesRegionOfInterest) GetY0() string {
	if o == nil || isNil(o.Y0) {
		var ret string
		return ret
	}
	return *o.Y0
}

// GetY0Ok returns a tuple with the Y0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCameraAnalyticsZonesRegionOfInterest) GetY0Ok() (*string, bool) {
	if o == nil || isNil(o.Y0) {
    return nil, false
	}
	return o.Y0, true
}

// HasY0 returns a boolean if a field has been set.
func (o *DevicesSerialCameraAnalyticsZonesRegionOfInterest) HasY0() bool {
	if o != nil && !isNil(o.Y0) {
		return true
	}

	return false
}

// SetY0 gets a reference to the given string and assigns it to the Y0 field.
func (o *DevicesSerialCameraAnalyticsZonesRegionOfInterest) SetY0(v string) {
	o.Y0 = &v
}

// GetX1 returns the X1 field value if set, zero value otherwise.
func (o *DevicesSerialCameraAnalyticsZonesRegionOfInterest) GetX1() string {
	if o == nil || isNil(o.X1) {
		var ret string
		return ret
	}
	return *o.X1
}

// GetX1Ok returns a tuple with the X1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCameraAnalyticsZonesRegionOfInterest) GetX1Ok() (*string, bool) {
	if o == nil || isNil(o.X1) {
    return nil, false
	}
	return o.X1, true
}

// HasX1 returns a boolean if a field has been set.
func (o *DevicesSerialCameraAnalyticsZonesRegionOfInterest) HasX1() bool {
	if o != nil && !isNil(o.X1) {
		return true
	}

	return false
}

// SetX1 gets a reference to the given string and assigns it to the X1 field.
func (o *DevicesSerialCameraAnalyticsZonesRegionOfInterest) SetX1(v string) {
	o.X1 = &v
}

// GetY1 returns the Y1 field value if set, zero value otherwise.
func (o *DevicesSerialCameraAnalyticsZonesRegionOfInterest) GetY1() string {
	if o == nil || isNil(o.Y1) {
		var ret string
		return ret
	}
	return *o.Y1
}

// GetY1Ok returns a tuple with the Y1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCameraAnalyticsZonesRegionOfInterest) GetY1Ok() (*string, bool) {
	if o == nil || isNil(o.Y1) {
    return nil, false
	}
	return o.Y1, true
}

// HasY1 returns a boolean if a field has been set.
func (o *DevicesSerialCameraAnalyticsZonesRegionOfInterest) HasY1() bool {
	if o != nil && !isNil(o.Y1) {
		return true
	}

	return false
}

// SetY1 gets a reference to the given string and assigns it to the Y1 field.
func (o *DevicesSerialCameraAnalyticsZonesRegionOfInterest) SetY1(v string) {
	o.Y1 = &v
}

func (o DevicesSerialCameraAnalyticsZonesRegionOfInterest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.X0) {
		toSerialize["x0"] = o.X0
	}
	if !isNil(o.Y0) {
		toSerialize["y0"] = o.Y0
	}
	if !isNil(o.X1) {
		toSerialize["x1"] = o.X1
	}
	if !isNil(o.Y1) {
		toSerialize["y1"] = o.Y1
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialCameraAnalyticsZonesRegionOfInterest struct {
	value *DevicesSerialCameraAnalyticsZonesRegionOfInterest
	isSet bool
}

func (v NullableDevicesSerialCameraAnalyticsZonesRegionOfInterest) Get() *DevicesSerialCameraAnalyticsZonesRegionOfInterest {
	return v.value
}

func (v *NullableDevicesSerialCameraAnalyticsZonesRegionOfInterest) Set(val *DevicesSerialCameraAnalyticsZonesRegionOfInterest) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialCameraAnalyticsZonesRegionOfInterest) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialCameraAnalyticsZonesRegionOfInterest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialCameraAnalyticsZonesRegionOfInterest(val *DevicesSerialCameraAnalyticsZonesRegionOfInterest) *NullableDevicesSerialCameraAnalyticsZonesRegionOfInterest {
	return &NullableDevicesSerialCameraAnalyticsZonesRegionOfInterest{value: val, isSet: true}
}

func (v NullableDevicesSerialCameraAnalyticsZonesRegionOfInterest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialCameraAnalyticsZonesRegionOfInterest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


