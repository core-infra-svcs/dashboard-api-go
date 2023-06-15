/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20030Products The network devices to be updated
type InlineResponse20030Products struct {
	Wireless *InlineResponse20030ProductsWireless `json:"wireless,omitempty"`
	Appliance *InlineResponse20030ProductsWireless `json:"appliance,omitempty"`
	Switch *InlineResponse20030ProductsWireless `json:"switch,omitempty"`
	Camera *InlineResponse20030ProductsWireless `json:"camera,omitempty"`
	CellularGateway *InlineResponse20030ProductsWireless `json:"cellularGateway,omitempty"`
	Sensor *InlineResponse20030ProductsWireless `json:"sensor,omitempty"`
	CloudGateway *InlineResponse20030ProductsWireless `json:"cloudGateway,omitempty"`
}

// NewInlineResponse20030Products instantiates a new InlineResponse20030Products object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20030Products() *InlineResponse20030Products {
	this := InlineResponse20030Products{}
	return &this
}

// NewInlineResponse20030ProductsWithDefaults instantiates a new InlineResponse20030Products object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20030ProductsWithDefaults() *InlineResponse20030Products {
	this := InlineResponse20030Products{}
	return &this
}

// GetWireless returns the Wireless field value if set, zero value otherwise.
func (o *InlineResponse20030Products) GetWireless() InlineResponse20030ProductsWireless {
	if o == nil || isNil(o.Wireless) {
		var ret InlineResponse20030ProductsWireless
		return ret
	}
	return *o.Wireless
}

// GetWirelessOk returns a tuple with the Wireless field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030Products) GetWirelessOk() (*InlineResponse20030ProductsWireless, bool) {
	if o == nil || isNil(o.Wireless) {
    return nil, false
	}
	return o.Wireless, true
}

// HasWireless returns a boolean if a field has been set.
func (o *InlineResponse20030Products) HasWireless() bool {
	if o != nil && !isNil(o.Wireless) {
		return true
	}

	return false
}

// SetWireless gets a reference to the given InlineResponse20030ProductsWireless and assigns it to the Wireless field.
func (o *InlineResponse20030Products) SetWireless(v InlineResponse20030ProductsWireless) {
	o.Wireless = &v
}

// GetAppliance returns the Appliance field value if set, zero value otherwise.
func (o *InlineResponse20030Products) GetAppliance() InlineResponse20030ProductsWireless {
	if o == nil || isNil(o.Appliance) {
		var ret InlineResponse20030ProductsWireless
		return ret
	}
	return *o.Appliance
}

// GetApplianceOk returns a tuple with the Appliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030Products) GetApplianceOk() (*InlineResponse20030ProductsWireless, bool) {
	if o == nil || isNil(o.Appliance) {
    return nil, false
	}
	return o.Appliance, true
}

// HasAppliance returns a boolean if a field has been set.
func (o *InlineResponse20030Products) HasAppliance() bool {
	if o != nil && !isNil(o.Appliance) {
		return true
	}

	return false
}

// SetAppliance gets a reference to the given InlineResponse20030ProductsWireless and assigns it to the Appliance field.
func (o *InlineResponse20030Products) SetAppliance(v InlineResponse20030ProductsWireless) {
	o.Appliance = &v
}

// GetSwitch returns the Switch field value if set, zero value otherwise.
func (o *InlineResponse20030Products) GetSwitch() InlineResponse20030ProductsWireless {
	if o == nil || isNil(o.Switch) {
		var ret InlineResponse20030ProductsWireless
		return ret
	}
	return *o.Switch
}

// GetSwitchOk returns a tuple with the Switch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030Products) GetSwitchOk() (*InlineResponse20030ProductsWireless, bool) {
	if o == nil || isNil(o.Switch) {
    return nil, false
	}
	return o.Switch, true
}

// HasSwitch returns a boolean if a field has been set.
func (o *InlineResponse20030Products) HasSwitch() bool {
	if o != nil && !isNil(o.Switch) {
		return true
	}

	return false
}

// SetSwitch gets a reference to the given InlineResponse20030ProductsWireless and assigns it to the Switch field.
func (o *InlineResponse20030Products) SetSwitch(v InlineResponse20030ProductsWireless) {
	o.Switch = &v
}

// GetCamera returns the Camera field value if set, zero value otherwise.
func (o *InlineResponse20030Products) GetCamera() InlineResponse20030ProductsWireless {
	if o == nil || isNil(o.Camera) {
		var ret InlineResponse20030ProductsWireless
		return ret
	}
	return *o.Camera
}

// GetCameraOk returns a tuple with the Camera field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030Products) GetCameraOk() (*InlineResponse20030ProductsWireless, bool) {
	if o == nil || isNil(o.Camera) {
    return nil, false
	}
	return o.Camera, true
}

// HasCamera returns a boolean if a field has been set.
func (o *InlineResponse20030Products) HasCamera() bool {
	if o != nil && !isNil(o.Camera) {
		return true
	}

	return false
}

// SetCamera gets a reference to the given InlineResponse20030ProductsWireless and assigns it to the Camera field.
func (o *InlineResponse20030Products) SetCamera(v InlineResponse20030ProductsWireless) {
	o.Camera = &v
}

// GetCellularGateway returns the CellularGateway field value if set, zero value otherwise.
func (o *InlineResponse20030Products) GetCellularGateway() InlineResponse20030ProductsWireless {
	if o == nil || isNil(o.CellularGateway) {
		var ret InlineResponse20030ProductsWireless
		return ret
	}
	return *o.CellularGateway
}

// GetCellularGatewayOk returns a tuple with the CellularGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030Products) GetCellularGatewayOk() (*InlineResponse20030ProductsWireless, bool) {
	if o == nil || isNil(o.CellularGateway) {
    return nil, false
	}
	return o.CellularGateway, true
}

// HasCellularGateway returns a boolean if a field has been set.
func (o *InlineResponse20030Products) HasCellularGateway() bool {
	if o != nil && !isNil(o.CellularGateway) {
		return true
	}

	return false
}

// SetCellularGateway gets a reference to the given InlineResponse20030ProductsWireless and assigns it to the CellularGateway field.
func (o *InlineResponse20030Products) SetCellularGateway(v InlineResponse20030ProductsWireless) {
	o.CellularGateway = &v
}

// GetSensor returns the Sensor field value if set, zero value otherwise.
func (o *InlineResponse20030Products) GetSensor() InlineResponse20030ProductsWireless {
	if o == nil || isNil(o.Sensor) {
		var ret InlineResponse20030ProductsWireless
		return ret
	}
	return *o.Sensor
}

// GetSensorOk returns a tuple with the Sensor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030Products) GetSensorOk() (*InlineResponse20030ProductsWireless, bool) {
	if o == nil || isNil(o.Sensor) {
    return nil, false
	}
	return o.Sensor, true
}

// HasSensor returns a boolean if a field has been set.
func (o *InlineResponse20030Products) HasSensor() bool {
	if o != nil && !isNil(o.Sensor) {
		return true
	}

	return false
}

// SetSensor gets a reference to the given InlineResponse20030ProductsWireless and assigns it to the Sensor field.
func (o *InlineResponse20030Products) SetSensor(v InlineResponse20030ProductsWireless) {
	o.Sensor = &v
}

// GetCloudGateway returns the CloudGateway field value if set, zero value otherwise.
func (o *InlineResponse20030Products) GetCloudGateway() InlineResponse20030ProductsWireless {
	if o == nil || isNil(o.CloudGateway) {
		var ret InlineResponse20030ProductsWireless
		return ret
	}
	return *o.CloudGateway
}

// GetCloudGatewayOk returns a tuple with the CloudGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030Products) GetCloudGatewayOk() (*InlineResponse20030ProductsWireless, bool) {
	if o == nil || isNil(o.CloudGateway) {
    return nil, false
	}
	return o.CloudGateway, true
}

// HasCloudGateway returns a boolean if a field has been set.
func (o *InlineResponse20030Products) HasCloudGateway() bool {
	if o != nil && !isNil(o.CloudGateway) {
		return true
	}

	return false
}

// SetCloudGateway gets a reference to the given InlineResponse20030ProductsWireless and assigns it to the CloudGateway field.
func (o *InlineResponse20030Products) SetCloudGateway(v InlineResponse20030ProductsWireless) {
	o.CloudGateway = &v
}

func (o InlineResponse20030Products) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.CloudGateway) {
		toSerialize["cloudGateway"] = o.CloudGateway
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20030Products struct {
	value *InlineResponse20030Products
	isSet bool
}

func (v NullableInlineResponse20030Products) Get() *InlineResponse20030Products {
	return v.value
}

func (v *NullableInlineResponse20030Products) Set(val *InlineResponse20030Products) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20030Products) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20030Products) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20030Products(val *InlineResponse20030Products) *NullableInlineResponse20030Products {
	return &NullableInlineResponse20030Products{value: val, isSet: true}
}

func (v NullableInlineResponse20030Products) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20030Products) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

