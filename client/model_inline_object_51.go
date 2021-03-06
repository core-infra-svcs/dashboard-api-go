/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject51 struct for InlineObject51
type InlineObject51 struct {
	// The site-to-site VPN mode. Can be one of 'none', 'spoke' or 'hub'
	Mode string `json:"mode"`
	// The list of VPN hubs, in order of preference. In spoke mode, at least 1 hub is required.
	Hubs *[]NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs `json:"hubs,omitempty"`
	// The list of subnets and their VPN presence.
	Subnets *[]NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets `json:"subnets,omitempty"`
}

// NewInlineObject51 instantiates a new InlineObject51 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject51(mode string) *InlineObject51 {
	this := InlineObject51{}
	this.Mode = mode
	return &this
}

// NewInlineObject51WithDefaults instantiates a new InlineObject51 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject51WithDefaults() *InlineObject51 {
	this := InlineObject51{}
	return &this
}

// GetMode returns the Mode field value
func (o *InlineObject51) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *InlineObject51) GetModeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *InlineObject51) SetMode(v string) {
	o.Mode = v
}

// GetHubs returns the Hubs field value if set, zero value otherwise.
func (o *InlineObject51) GetHubs() []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs {
	if o == nil || o.Hubs == nil {
		var ret []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs
		return ret
	}
	return *o.Hubs
}

// GetHubsOk returns a tuple with the Hubs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject51) GetHubsOk() (*[]NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs, bool) {
	if o == nil || o.Hubs == nil {
		return nil, false
	}
	return o.Hubs, true
}

// HasHubs returns a boolean if a field has been set.
func (o *InlineObject51) HasHubs() bool {
	if o != nil && o.Hubs != nil {
		return true
	}

	return false
}

// SetHubs gets a reference to the given []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs and assigns it to the Hubs field.
func (o *InlineObject51) SetHubs(v []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) {
	o.Hubs = &v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *InlineObject51) GetSubnets() []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets {
	if o == nil || o.Subnets == nil {
		var ret []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets
		return ret
	}
	return *o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject51) GetSubnetsOk() (*[]NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets, bool) {
	if o == nil || o.Subnets == nil {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *InlineObject51) HasSubnets() bool {
	if o != nil && o.Subnets != nil {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets and assigns it to the Subnets field.
func (o *InlineObject51) SetSubnets(v []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) {
	o.Subnets = &v
}

func (o InlineObject51) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if o.Hubs != nil {
		toSerialize["hubs"] = o.Hubs
	}
	if o.Subnets != nil {
		toSerialize["subnets"] = o.Subnets
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject51 struct {
	value *InlineObject51
	isSet bool
}

func (v NullableInlineObject51) Get() *InlineObject51 {
	return v.value
}

func (v *NullableInlineObject51) Set(val *InlineObject51) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject51) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject51) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject51(val *InlineObject51) *NullableInlineObject51 {
	return &NullableInlineObject51{value: val, isSet: true}
}

func (v NullableInlineObject51) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject51) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


