/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20017Products The network devices to be updated
type InlineResponse20017Products struct {
	Wireless *InlineResponse20017ProductsWireless `json:"wireless,omitempty"`
	Appliance *InlineResponse20017ProductsWireless `json:"appliance,omitempty"`
	Switch *InlineResponse20017ProductsWireless `json:"switch,omitempty"`
	Camera *InlineResponse20017ProductsWireless `json:"camera,omitempty"`
	CellularGateway *InlineResponse20017ProductsWireless `json:"cellularGateway,omitempty"`
	Sensor *InlineResponse20017ProductsWireless `json:"sensor,omitempty"`
}

// NewInlineResponse20017Products instantiates a new InlineResponse20017Products object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20017Products() *InlineResponse20017Products {
	this := InlineResponse20017Products{}
	return &this
}

// NewInlineResponse20017ProductsWithDefaults instantiates a new InlineResponse20017Products object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20017ProductsWithDefaults() *InlineResponse20017Products {
	this := InlineResponse20017Products{}
	return &this
}

// GetWireless returns the Wireless field value if set, zero value otherwise.
func (o *InlineResponse20017Products) GetWireless() InlineResponse20017ProductsWireless {
	if o == nil || isNil(o.Wireless) {
		var ret InlineResponse20017ProductsWireless
		return ret
	}
	return *o.Wireless
}

// GetWirelessOk returns a tuple with the Wireless field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017Products) GetWirelessOk() (*InlineResponse20017ProductsWireless, bool) {
	if o == nil || isNil(o.Wireless) {
    return nil, false
	}
	return o.Wireless, true
}

// HasWireless returns a boolean if a field has been set.
func (o *InlineResponse20017Products) HasWireless() bool {
	if o != nil && !isNil(o.Wireless) {
		return true
	}

	return false
}

// SetWireless gets a reference to the given InlineResponse20017ProductsWireless and assigns it to the Wireless field.
func (o *InlineResponse20017Products) SetWireless(v InlineResponse20017ProductsWireless) {
	o.Wireless = &v
}

// GetAppliance returns the Appliance field value if set, zero value otherwise.
func (o *InlineResponse20017Products) GetAppliance() InlineResponse20017ProductsWireless {
	if o == nil || isNil(o.Appliance) {
		var ret InlineResponse20017ProductsWireless
		return ret
	}
	return *o.Appliance
}

// GetApplianceOk returns a tuple with the Appliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017Products) GetApplianceOk() (*InlineResponse20017ProductsWireless, bool) {
	if o == nil || isNil(o.Appliance) {
    return nil, false
	}
	return o.Appliance, true
}

// HasAppliance returns a boolean if a field has been set.
func (o *InlineResponse20017Products) HasAppliance() bool {
	if o != nil && !isNil(o.Appliance) {
		return true
	}

	return false
}

// SetAppliance gets a reference to the given InlineResponse20017ProductsWireless and assigns it to the Appliance field.
func (o *InlineResponse20017Products) SetAppliance(v InlineResponse20017ProductsWireless) {
	o.Appliance = &v
}

// GetSwitch returns the Switch field value if set, zero value otherwise.
func (o *InlineResponse20017Products) GetSwitch() InlineResponse20017ProductsWireless {
	if o == nil || isNil(o.Switch) {
		var ret InlineResponse20017ProductsWireless
		return ret
	}
	return *o.Switch
}

// GetSwitchOk returns a tuple with the Switch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017Products) GetSwitchOk() (*InlineResponse20017ProductsWireless, bool) {
	if o == nil || isNil(o.Switch) {
    return nil, false
	}
	return o.Switch, true
}

// HasSwitch returns a boolean if a field has been set.
func (o *InlineResponse20017Products) HasSwitch() bool {
	if o != nil && !isNil(o.Switch) {
		return true
	}

	return false
}

// SetSwitch gets a reference to the given InlineResponse20017ProductsWireless and assigns it to the Switch field.
func (o *InlineResponse20017Products) SetSwitch(v InlineResponse20017ProductsWireless) {
	o.Switch = &v
}

// GetCamera returns the Camera field value if set, zero value otherwise.
func (o *InlineResponse20017Products) GetCamera() InlineResponse20017ProductsWireless {
	if o == nil || isNil(o.Camera) {
		var ret InlineResponse20017ProductsWireless
		return ret
	}
	return *o.Camera
}

// GetCameraOk returns a tuple with the Camera field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017Products) GetCameraOk() (*InlineResponse20017ProductsWireless, bool) {
	if o == nil || isNil(o.Camera) {
    return nil, false
	}
	return o.Camera, true
}

// HasCamera returns a boolean if a field has been set.
func (o *InlineResponse20017Products) HasCamera() bool {
	if o != nil && !isNil(o.Camera) {
		return true
	}

	return false
}

// SetCamera gets a reference to the given InlineResponse20017ProductsWireless and assigns it to the Camera field.
func (o *InlineResponse20017Products) SetCamera(v InlineResponse20017ProductsWireless) {
	o.Camera = &v
}

// GetCellularGateway returns the CellularGateway field value if set, zero value otherwise.
func (o *InlineResponse20017Products) GetCellularGateway() InlineResponse20017ProductsWireless {
	if o == nil || isNil(o.CellularGateway) {
		var ret InlineResponse20017ProductsWireless
		return ret
	}
	return *o.CellularGateway
}

// GetCellularGatewayOk returns a tuple with the CellularGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017Products) GetCellularGatewayOk() (*InlineResponse20017ProductsWireless, bool) {
	if o == nil || isNil(o.CellularGateway) {
    return nil, false
	}
	return o.CellularGateway, true
}

// HasCellularGateway returns a boolean if a field has been set.
func (o *InlineResponse20017Products) HasCellularGateway() bool {
	if o != nil && !isNil(o.CellularGateway) {
		return true
	}

	return false
}

// SetCellularGateway gets a reference to the given InlineResponse20017ProductsWireless and assigns it to the CellularGateway field.
func (o *InlineResponse20017Products) SetCellularGateway(v InlineResponse20017ProductsWireless) {
	o.CellularGateway = &v
}

// GetSensor returns the Sensor field value if set, zero value otherwise.
func (o *InlineResponse20017Products) GetSensor() InlineResponse20017ProductsWireless {
	if o == nil || isNil(o.Sensor) {
		var ret InlineResponse20017ProductsWireless
		return ret
	}
	return *o.Sensor
}

// GetSensorOk returns a tuple with the Sensor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017Products) GetSensorOk() (*InlineResponse20017ProductsWireless, bool) {
	if o == nil || isNil(o.Sensor) {
    return nil, false
	}
	return o.Sensor, true
}

// HasSensor returns a boolean if a field has been set.
func (o *InlineResponse20017Products) HasSensor() bool {
	if o != nil && !isNil(o.Sensor) {
		return true
	}

	return false
}

// SetSensor gets a reference to the given InlineResponse20017ProductsWireless and assigns it to the Sensor field.
func (o *InlineResponse20017Products) SetSensor(v InlineResponse20017ProductsWireless) {
	o.Sensor = &v
}

func (o InlineResponse20017Products) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Wireless) {
		toSerialize["wireless"] = o.Wireless
	}
	if !isNil(o.Appliance) {
		toSerialize["appliance"] = o.Appliance
	}
	if !isNil(o.Switch) {
		toSerialize["switch"] = o.Switch
	}
	if !isNil(o.Camera) {
		toSerialize["camera"] = o.Camera
	}
	if !isNil(o.CellularGateway) {
		toSerialize["cellularGateway"] = o.CellularGateway
	}
	if !isNil(o.Sensor) {
		toSerialize["sensor"] = o.Sensor
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20017Products struct {
	value *InlineResponse20017Products
	isSet bool
}

func (v NullableInlineResponse20017Products) Get() *InlineResponse20017Products {
	return v.value
}

func (v *NullableInlineResponse20017Products) Set(val *InlineResponse20017Products) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20017Products) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20017Products) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20017Products(val *InlineResponse20017Products) *NullableInlineResponse20017Products {
	return &NullableInlineResponse20017Products{value: val, isSet: true}
}

func (v NullableInlineResponse20017Products) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20017Products) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


