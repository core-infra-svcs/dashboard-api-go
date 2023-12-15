/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFirmwareUpgradesProducts Contains information about the network to update
type NetworksNetworkIdFirmwareUpgradesProducts struct {
	Wireless *NetworksNetworkIdFirmwareUpgradesProductsWireless `json:"wireless,omitempty"`
	Appliance *NetworksNetworkIdFirmwareUpgradesProductsWireless `json:"appliance,omitempty"`
	Switch *NetworksNetworkIdFirmwareUpgradesProductsWireless `json:"switch,omitempty"`
	Camera *NetworksNetworkIdFirmwareUpgradesProductsWireless `json:"camera,omitempty"`
	CellularGateway *NetworksNetworkIdFirmwareUpgradesProductsWireless `json:"cellularGateway,omitempty"`
	Sensor *NetworksNetworkIdFirmwareUpgradesProductsWireless `json:"sensor,omitempty"`
	SwitchCatalyst *NetworksNetworkIdFirmwareUpgradesProductsWireless `json:"switchCatalyst,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesProducts instantiates a new NetworksNetworkIdFirmwareUpgradesProducts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesProducts() *NetworksNetworkIdFirmwareUpgradesProducts {
	this := NetworksNetworkIdFirmwareUpgradesProducts{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesProductsWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesProducts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesProductsWithDefaults() *NetworksNetworkIdFirmwareUpgradesProducts {
	this := NetworksNetworkIdFirmwareUpgradesProducts{}
	return &this
}

// GetWireless returns the Wireless field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) GetWireless() NetworksNetworkIdFirmwareUpgradesProductsWireless {
	if o == nil || isNil(o.Wireless) {
		var ret NetworksNetworkIdFirmwareUpgradesProductsWireless
		return ret
	}
	return *o.Wireless
}

// GetWirelessOk returns a tuple with the Wireless field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) GetWirelessOk() (*NetworksNetworkIdFirmwareUpgradesProductsWireless, bool) {
	if o == nil || isNil(o.Wireless) {
    return nil, false
	}
	return o.Wireless, true
}

// HasWireless returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) HasWireless() bool {
	if o != nil && !isNil(o.Wireless) {
		return true
	}

	return false
}

// SetWireless gets a reference to the given NetworksNetworkIdFirmwareUpgradesProductsWireless and assigns it to the Wireless field.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) SetWireless(v NetworksNetworkIdFirmwareUpgradesProductsWireless) {
	o.Wireless = &v
}

// GetAppliance returns the Appliance field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) GetAppliance() NetworksNetworkIdFirmwareUpgradesProductsWireless {
	if o == nil || isNil(o.Appliance) {
		var ret NetworksNetworkIdFirmwareUpgradesProductsWireless
		return ret
	}
	return *o.Appliance
}

// GetApplianceOk returns a tuple with the Appliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) GetApplianceOk() (*NetworksNetworkIdFirmwareUpgradesProductsWireless, bool) {
	if o == nil || isNil(o.Appliance) {
    return nil, false
	}
	return o.Appliance, true
}

// HasAppliance returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) HasAppliance() bool {
	if o != nil && !isNil(o.Appliance) {
		return true
	}

	return false
}

// SetAppliance gets a reference to the given NetworksNetworkIdFirmwareUpgradesProductsWireless and assigns it to the Appliance field.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) SetAppliance(v NetworksNetworkIdFirmwareUpgradesProductsWireless) {
	o.Appliance = &v
}

// GetSwitch returns the Switch field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) GetSwitch() NetworksNetworkIdFirmwareUpgradesProductsWireless {
	if o == nil || isNil(o.Switch) {
		var ret NetworksNetworkIdFirmwareUpgradesProductsWireless
		return ret
	}
	return *o.Switch
}

// GetSwitchOk returns a tuple with the Switch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) GetSwitchOk() (*NetworksNetworkIdFirmwareUpgradesProductsWireless, bool) {
	if o == nil || isNil(o.Switch) {
    return nil, false
	}
	return o.Switch, true
}

// HasSwitch returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) HasSwitch() bool {
	if o != nil && !isNil(o.Switch) {
		return true
	}

	return false
}

// SetSwitch gets a reference to the given NetworksNetworkIdFirmwareUpgradesProductsWireless and assigns it to the Switch field.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) SetSwitch(v NetworksNetworkIdFirmwareUpgradesProductsWireless) {
	o.Switch = &v
}

// GetCamera returns the Camera field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) GetCamera() NetworksNetworkIdFirmwareUpgradesProductsWireless {
	if o == nil || isNil(o.Camera) {
		var ret NetworksNetworkIdFirmwareUpgradesProductsWireless
		return ret
	}
	return *o.Camera
}

// GetCameraOk returns a tuple with the Camera field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) GetCameraOk() (*NetworksNetworkIdFirmwareUpgradesProductsWireless, bool) {
	if o == nil || isNil(o.Camera) {
    return nil, false
	}
	return o.Camera, true
}

// HasCamera returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) HasCamera() bool {
	if o != nil && !isNil(o.Camera) {
		return true
	}

	return false
}

// SetCamera gets a reference to the given NetworksNetworkIdFirmwareUpgradesProductsWireless and assigns it to the Camera field.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) SetCamera(v NetworksNetworkIdFirmwareUpgradesProductsWireless) {
	o.Camera = &v
}

// GetCellularGateway returns the CellularGateway field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) GetCellularGateway() NetworksNetworkIdFirmwareUpgradesProductsWireless {
	if o == nil || isNil(o.CellularGateway) {
		var ret NetworksNetworkIdFirmwareUpgradesProductsWireless
		return ret
	}
	return *o.CellularGateway
}

// GetCellularGatewayOk returns a tuple with the CellularGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) GetCellularGatewayOk() (*NetworksNetworkIdFirmwareUpgradesProductsWireless, bool) {
	if o == nil || isNil(o.CellularGateway) {
    return nil, false
	}
	return o.CellularGateway, true
}

// HasCellularGateway returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) HasCellularGateway() bool {
	if o != nil && !isNil(o.CellularGateway) {
		return true
	}

	return false
}

// SetCellularGateway gets a reference to the given NetworksNetworkIdFirmwareUpgradesProductsWireless and assigns it to the CellularGateway field.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) SetCellularGateway(v NetworksNetworkIdFirmwareUpgradesProductsWireless) {
	o.CellularGateway = &v
}

// GetSensor returns the Sensor field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) GetSensor() NetworksNetworkIdFirmwareUpgradesProductsWireless {
	if o == nil || isNil(o.Sensor) {
		var ret NetworksNetworkIdFirmwareUpgradesProductsWireless
		return ret
	}
	return *o.Sensor
}

// GetSensorOk returns a tuple with the Sensor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) GetSensorOk() (*NetworksNetworkIdFirmwareUpgradesProductsWireless, bool) {
	if o == nil || isNil(o.Sensor) {
    return nil, false
	}
	return o.Sensor, true
}

// HasSensor returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) HasSensor() bool {
	if o != nil && !isNil(o.Sensor) {
		return true
	}

	return false
}

// SetSensor gets a reference to the given NetworksNetworkIdFirmwareUpgradesProductsWireless and assigns it to the Sensor field.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) SetSensor(v NetworksNetworkIdFirmwareUpgradesProductsWireless) {
	o.Sensor = &v
}

// GetSwitchCatalyst returns the SwitchCatalyst field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) GetSwitchCatalyst() NetworksNetworkIdFirmwareUpgradesProductsWireless {
	if o == nil || isNil(o.SwitchCatalyst) {
		var ret NetworksNetworkIdFirmwareUpgradesProductsWireless
		return ret
	}
	return *o.SwitchCatalyst
}

// GetSwitchCatalystOk returns a tuple with the SwitchCatalyst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) GetSwitchCatalystOk() (*NetworksNetworkIdFirmwareUpgradesProductsWireless, bool) {
	if o == nil || isNil(o.SwitchCatalyst) {
    return nil, false
	}
	return o.SwitchCatalyst, true
}

// HasSwitchCatalyst returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) HasSwitchCatalyst() bool {
	if o != nil && !isNil(o.SwitchCatalyst) {
		return true
	}

	return false
}

// SetSwitchCatalyst gets a reference to the given NetworksNetworkIdFirmwareUpgradesProductsWireless and assigns it to the SwitchCatalyst field.
func (o *NetworksNetworkIdFirmwareUpgradesProducts) SetSwitchCatalyst(v NetworksNetworkIdFirmwareUpgradesProductsWireless) {
	o.SwitchCatalyst = &v
}

func (o NetworksNetworkIdFirmwareUpgradesProducts) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.SwitchCatalyst) {
		toSerialize["switchCatalyst"] = o.SwitchCatalyst
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesProducts struct {
	value *NetworksNetworkIdFirmwareUpgradesProducts
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesProducts) Get() *NetworksNetworkIdFirmwareUpgradesProducts {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesProducts) Set(val *NetworksNetworkIdFirmwareUpgradesProducts) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesProducts) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesProducts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesProducts(val *NetworksNetworkIdFirmwareUpgradesProducts) *NullableNetworksNetworkIdFirmwareUpgradesProducts {
	return &NullableNetworksNetworkIdFirmwareUpgradesProducts{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesProducts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesProducts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


