/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdHealthAlertsScopeDevices struct for NetworksNetworkIdHealthAlertsScopeDevices
type NetworksNetworkIdHealthAlertsScopeDevices struct {
	// URL to the device
	Url *string `json:"url,omitempty"`
	// Name of the device
	Name *string `json:"name,omitempty"`
	// Product type of the device
	ProductType *string `json:"productType,omitempty"`
	// Serial number of the device
	Serial *string `json:"serial,omitempty"`
	// The mac address of the device
	Mac *string `json:"mac,omitempty"`
	Lldp *NetworksNetworkIdHealthAlertsScopeLldp `json:"lldp,omitempty"`
	// Clients related to the device
	Clients []NetworksNetworkIdHealthAlertsScopeClients `json:"clients,omitempty"`
}

// NewNetworksNetworkIdHealthAlertsScopeDevices instantiates a new NetworksNetworkIdHealthAlertsScopeDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdHealthAlertsScopeDevices() *NetworksNetworkIdHealthAlertsScopeDevices {
	this := NetworksNetworkIdHealthAlertsScopeDevices{}
	return &this
}

// NewNetworksNetworkIdHealthAlertsScopeDevicesWithDefaults instantiates a new NetworksNetworkIdHealthAlertsScopeDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdHealthAlertsScopeDevicesWithDefaults() *NetworksNetworkIdHealthAlertsScopeDevices {
	this := NetworksNetworkIdHealthAlertsScopeDevices{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) SetUrl(v string) {
	o.Url = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) SetName(v string) {
	o.Name = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) SetProductType(v string) {
	o.ProductType = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) SetMac(v string) {
	o.Mac = &v
}

// GetLldp returns the Lldp field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetLldp() NetworksNetworkIdHealthAlertsScopeLldp {
	if o == nil || isNil(o.Lldp) {
		var ret NetworksNetworkIdHealthAlertsScopeLldp
		return ret
	}
	return *o.Lldp
}

// GetLldpOk returns a tuple with the Lldp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetLldpOk() (*NetworksNetworkIdHealthAlertsScopeLldp, bool) {
	if o == nil || isNil(o.Lldp) {
    return nil, false
	}
	return o.Lldp, true
}

// HasLldp returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) HasLldp() bool {
	if o != nil && !isNil(o.Lldp) {
		return true
	}

	return false
}

// SetLldp gets a reference to the given NetworksNetworkIdHealthAlertsScopeLldp and assigns it to the Lldp field.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) SetLldp(v NetworksNetworkIdHealthAlertsScopeLldp) {
	o.Lldp = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetClients() []NetworksNetworkIdHealthAlertsScopeClients {
	if o == nil || isNil(o.Clients) {
		var ret []NetworksNetworkIdHealthAlertsScopeClients
		return ret
	}
	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetClientsOk() ([]NetworksNetworkIdHealthAlertsScopeClients, bool) {
	if o == nil || isNil(o.Clients) {
    return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) HasClients() bool {
	if o != nil && !isNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given []NetworksNetworkIdHealthAlertsScopeClients and assigns it to the Clients field.
func (o *NetworksNetworkIdHealthAlertsScopeDevices) SetClients(v []NetworksNetworkIdHealthAlertsScopeClients) {
	o.Clients = v
}

func (o NetworksNetworkIdHealthAlertsScopeDevices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
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
	if !isNil(o.Lldp) {
		toSerialize["lldp"] = o.Lldp
	}
	if !isNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdHealthAlertsScopeDevices struct {
	value *NetworksNetworkIdHealthAlertsScopeDevices
	isSet bool
}

func (v NullableNetworksNetworkIdHealthAlertsScopeDevices) Get() *NetworksNetworkIdHealthAlertsScopeDevices {
	return v.value
}

func (v *NullableNetworksNetworkIdHealthAlertsScopeDevices) Set(val *NetworksNetworkIdHealthAlertsScopeDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdHealthAlertsScopeDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdHealthAlertsScopeDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdHealthAlertsScopeDevices(val *NetworksNetworkIdHealthAlertsScopeDevices) *NullableNetworksNetworkIdHealthAlertsScopeDevices {
	return &NullableNetworksNetworkIdHealthAlertsScopeDevices{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdHealthAlertsScopeDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdHealthAlertsScopeDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


