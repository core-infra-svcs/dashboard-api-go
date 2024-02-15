/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberHotspot20Venue Venue settings for this SSID
type NetworksNetworkIdWirelessSsidsNumberHotspot20Venue struct {
	// Venue name
	Name *string `json:"name,omitempty"`
	// Venue type ('Unspecified', 'Unspecified Assembly', 'Arena', 'Stadium', 'Passenger Terminal', 'Amphitheater', 'Amusement Park', 'Place of Worship', 'Convention Center', 'Library', 'Museum', 'Restaurant', 'Theater', 'Bar', 'Coffee Shop', 'Zoo or Aquarium', 'Emergency Coordination Center', 'Unspecified Business', 'Doctor or Dentist office', 'Bank', 'Fire Station', 'Police Station', 'Post Office', 'Professional Office', 'Research and Development Facility', 'Attorney Office', 'Unspecified Educational', 'School, Primary', 'School, Secondary', 'University or College', 'Unspecified Factory and Industrial', 'Factory', 'Unspecified Institutional', 'Hospital', 'Long-Term Care Facility', 'Alcohol and Drug Rehabilitation Center', 'Group Home', 'Prison or Jail', 'Unspecified Mercantile', 'Retail Store', 'Grocery Market', 'Automotive Service Station', 'Shopping Mall', 'Gas Station', 'Unspecified Residential', 'Private Residence', 'Hotel or Motel', 'Dormitory', 'Boarding House', 'Unspecified Storage', 'Unspecified Utility and Miscellaneous', 'Unspecified Vehicular', 'Automobile or Truck', 'Airplane', 'Bus', 'Ferry', 'Ship or Boat', 'Train', 'Motor Bike', 'Unspecified Outdoor', 'Muni-mesh Network', 'City Park', 'Rest Area', 'Traffic Control', 'Bus Stop', 'Kiosk')
	Type *string `json:"type,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberHotspot20Venue instantiates a new NetworksNetworkIdWirelessSsidsNumberHotspot20Venue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberHotspot20Venue() *NetworksNetworkIdWirelessSsidsNumberHotspot20Venue {
	this := NetworksNetworkIdWirelessSsidsNumberHotspot20Venue{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberHotspot20VenueWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberHotspot20Venue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberHotspot20VenueWithDefaults() *NetworksNetworkIdWirelessSsidsNumberHotspot20Venue {
	this := NetworksNetworkIdWirelessSsidsNumberHotspot20Venue{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Venue) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Venue) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Venue) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Venue) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Venue) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Venue) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Venue) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Venue) SetType(v string) {
	o.Type = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberHotspot20Venue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Venue struct {
	value *NetworksNetworkIdWirelessSsidsNumberHotspot20Venue
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Venue) Get() *NetworksNetworkIdWirelessSsidsNumberHotspot20Venue {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Venue) Set(val *NetworksNetworkIdWirelessSsidsNumberHotspot20Venue) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Venue) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Venue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberHotspot20Venue(val *NetworksNetworkIdWirelessSsidsNumberHotspot20Venue) *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Venue {
	return &NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Venue{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Venue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Venue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


