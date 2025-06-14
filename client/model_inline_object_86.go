/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject86 struct for InlineObject86
type InlineObject86 struct {
	// Name of the cluster
	Name *string `json:"name,omitempty"`
	// Uplink interface settings of the cluster
	Uplinks []NetworksNetworkIdCampusGatewayClustersClusterIdUplinks `json:"uplinks,omitempty"`
	// Tunnel interface settings of the cluster: Reuse uplink or specify tunnel interface
	Tunnels []NetworksNetworkIdCampusGatewayClustersClusterIdTunnels `json:"tunnels,omitempty"`
	Nameservers *NetworksNetworkIdCampusGatewayClustersNameservers `json:"nameservers,omitempty"`
	// Port channel settings of the cluster
	PortChannels []NetworksNetworkIdCampusGatewayClustersClusterIdPortChannels `json:"portChannels,omitempty"`
	// Devices in the cluster
	Devices []NetworksNetworkIdCampusGatewayClustersClusterIdDevices `json:"devices,omitempty"`
	// Notes about cluster with max size of 511 characters allowed
	Notes *string `json:"notes,omitempty"`
}

// NewInlineObject86 instantiates a new InlineObject86 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject86() *InlineObject86 {
	this := InlineObject86{}
	return &this
}

// NewInlineObject86WithDefaults instantiates a new InlineObject86 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject86WithDefaults() *InlineObject86 {
	this := InlineObject86{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject86) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject86) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject86) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject86) SetName(v string) {
	o.Name = &v
}

// GetUplinks returns the Uplinks field value if set, zero value otherwise.
func (o *InlineObject86) GetUplinks() []NetworksNetworkIdCampusGatewayClustersClusterIdUplinks {
	if o == nil || isNil(o.Uplinks) {
		var ret []NetworksNetworkIdCampusGatewayClustersClusterIdUplinks
		return ret
	}
	return o.Uplinks
}

// GetUplinksOk returns a tuple with the Uplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject86) GetUplinksOk() ([]NetworksNetworkIdCampusGatewayClustersClusterIdUplinks, bool) {
	if o == nil || isNil(o.Uplinks) {
    return nil, false
	}
	return o.Uplinks, true
}

// HasUplinks returns a boolean if a field has been set.
func (o *InlineObject86) HasUplinks() bool {
	if o != nil && !isNil(o.Uplinks) {
		return true
	}

	return false
}

// SetUplinks gets a reference to the given []NetworksNetworkIdCampusGatewayClustersClusterIdUplinks and assigns it to the Uplinks field.
func (o *InlineObject86) SetUplinks(v []NetworksNetworkIdCampusGatewayClustersClusterIdUplinks) {
	o.Uplinks = v
}

// GetTunnels returns the Tunnels field value if set, zero value otherwise.
func (o *InlineObject86) GetTunnels() []NetworksNetworkIdCampusGatewayClustersClusterIdTunnels {
	if o == nil || isNil(o.Tunnels) {
		var ret []NetworksNetworkIdCampusGatewayClustersClusterIdTunnels
		return ret
	}
	return o.Tunnels
}

// GetTunnelsOk returns a tuple with the Tunnels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject86) GetTunnelsOk() ([]NetworksNetworkIdCampusGatewayClustersClusterIdTunnels, bool) {
	if o == nil || isNil(o.Tunnels) {
    return nil, false
	}
	return o.Tunnels, true
}

// HasTunnels returns a boolean if a field has been set.
func (o *InlineObject86) HasTunnels() bool {
	if o != nil && !isNil(o.Tunnels) {
		return true
	}

	return false
}

// SetTunnels gets a reference to the given []NetworksNetworkIdCampusGatewayClustersClusterIdTunnels and assigns it to the Tunnels field.
func (o *InlineObject86) SetTunnels(v []NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) {
	o.Tunnels = v
}

// GetNameservers returns the Nameservers field value if set, zero value otherwise.
func (o *InlineObject86) GetNameservers() NetworksNetworkIdCampusGatewayClustersNameservers {
	if o == nil || isNil(o.Nameservers) {
		var ret NetworksNetworkIdCampusGatewayClustersNameservers
		return ret
	}
	return *o.Nameservers
}

// GetNameserversOk returns a tuple with the Nameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject86) GetNameserversOk() (*NetworksNetworkIdCampusGatewayClustersNameservers, bool) {
	if o == nil || isNil(o.Nameservers) {
    return nil, false
	}
	return o.Nameservers, true
}

// HasNameservers returns a boolean if a field has been set.
func (o *InlineObject86) HasNameservers() bool {
	if o != nil && !isNil(o.Nameservers) {
		return true
	}

	return false
}

// SetNameservers gets a reference to the given NetworksNetworkIdCampusGatewayClustersNameservers and assigns it to the Nameservers field.
func (o *InlineObject86) SetNameservers(v NetworksNetworkIdCampusGatewayClustersNameservers) {
	o.Nameservers = &v
}

// GetPortChannels returns the PortChannels field value if set, zero value otherwise.
func (o *InlineObject86) GetPortChannels() []NetworksNetworkIdCampusGatewayClustersClusterIdPortChannels {
	if o == nil || isNil(o.PortChannels) {
		var ret []NetworksNetworkIdCampusGatewayClustersClusterIdPortChannels
		return ret
	}
	return o.PortChannels
}

// GetPortChannelsOk returns a tuple with the PortChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject86) GetPortChannelsOk() ([]NetworksNetworkIdCampusGatewayClustersClusterIdPortChannels, bool) {
	if o == nil || isNil(o.PortChannels) {
    return nil, false
	}
	return o.PortChannels, true
}

// HasPortChannels returns a boolean if a field has been set.
func (o *InlineObject86) HasPortChannels() bool {
	if o != nil && !isNil(o.PortChannels) {
		return true
	}

	return false
}

// SetPortChannels gets a reference to the given []NetworksNetworkIdCampusGatewayClustersClusterIdPortChannels and assigns it to the PortChannels field.
func (o *InlineObject86) SetPortChannels(v []NetworksNetworkIdCampusGatewayClustersClusterIdPortChannels) {
	o.PortChannels = v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *InlineObject86) GetDevices() []NetworksNetworkIdCampusGatewayClustersClusterIdDevices {
	if o == nil || isNil(o.Devices) {
		var ret []NetworksNetworkIdCampusGatewayClustersClusterIdDevices
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject86) GetDevicesOk() ([]NetworksNetworkIdCampusGatewayClustersClusterIdDevices, bool) {
	if o == nil || isNil(o.Devices) {
    return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *InlineObject86) HasDevices() bool {
	if o != nil && !isNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []NetworksNetworkIdCampusGatewayClustersClusterIdDevices and assigns it to the Devices field.
func (o *InlineObject86) SetDevices(v []NetworksNetworkIdCampusGatewayClustersClusterIdDevices) {
	o.Devices = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineObject86) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject86) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineObject86) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineObject86) SetNotes(v string) {
	o.Notes = &v
}

func (o InlineObject86) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Uplinks) {
		toSerialize["uplinks"] = o.Uplinks
	}
	if !isNil(o.Tunnels) {
		toSerialize["tunnels"] = o.Tunnels
	}
	if !isNil(o.Nameservers) {
		toSerialize["nameservers"] = o.Nameservers
	}
	if !isNil(o.PortChannels) {
		toSerialize["portChannels"] = o.PortChannels
	}
	if !isNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject86 struct {
	value *InlineObject86
	isSet bool
}

func (v NullableInlineObject86) Get() *InlineObject86 {
	return v.value
}

func (v *NullableInlineObject86) Set(val *InlineObject86) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject86) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject86) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject86(val *InlineObject86) *NullableInlineObject86 {
	return &NullableInlineObject86{value: val, isSet: true}
}

func (v NullableInlineObject86) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject86) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


