/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject76 struct for InlineObject76
type InlineObject76 struct {
	// The site-to-site VPN mode. Can be one of 'none', 'spoke' or 'hub'
	Mode string `json:"mode"`
	// The list of VPN hubs, in order of preference. In spoke mode, at least 1 hub is required.
	Hubs []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs `json:"hubs,omitempty"`
	// The list of subnets and their VPN presence.
	Subnets []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets `json:"subnets,omitempty"`
}

// NewInlineObject76 instantiates a new InlineObject76 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject76(mode string) *InlineObject76 {
	this := InlineObject76{}
	this.Mode = mode
	return &this
}

// NewInlineObject76WithDefaults instantiates a new InlineObject76 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject76WithDefaults() *InlineObject76 {
	this := InlineObject76{}
	return &this
}

// GetMode returns the Mode field value
func (o *InlineObject76) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetModeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *InlineObject76) SetMode(v string) {
	o.Mode = v
}

// GetHubs returns the Hubs field value if set, zero value otherwise.
func (o *InlineObject76) GetHubs() []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs {
	if o == nil || isNil(o.Hubs) {
		var ret []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs
		return ret
	}
	return o.Hubs
}

// GetHubsOk returns a tuple with the Hubs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetHubsOk() ([]NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs, bool) {
	if o == nil || isNil(o.Hubs) {
    return nil, false
	}
	return o.Hubs, true
}

// HasHubs returns a boolean if a field has been set.
func (o *InlineObject76) HasHubs() bool {
	if o != nil && !isNil(o.Hubs) {
		return true
	}

	return false
}

// SetHubs gets a reference to the given []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs and assigns it to the Hubs field.
func (o *InlineObject76) SetHubs(v []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) {
	o.Hubs = v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *InlineObject76) GetSubnets() []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets {
	if o == nil || isNil(o.Subnets) {
		var ret []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets
		return ret
	}
	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetSubnetsOk() ([]NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets, bool) {
	if o == nil || isNil(o.Subnets) {
    return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *InlineObject76) HasSubnets() bool {
	if o != nil && !isNil(o.Subnets) {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets and assigns it to the Subnets field.
func (o *InlineObject76) SetSubnets(v []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) {
	o.Subnets = v
}

func (o InlineObject76) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.Hubs) {
		toSerialize["hubs"] = o.Hubs
	}
	if !isNil(o.Subnets) {
		toSerialize["subnets"] = o.Subnets
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject76 struct {
	value *InlineObject76
	isSet bool
}

func (v NullableInlineObject76) Get() *InlineObject76 {
	return v.value
}

func (v *NullableInlineObject76) Set(val *InlineObject76) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject76) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject76) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject76(val *InlineObject76) *NullableInlineObject76 {
	return &NullableInlineObject76{value: val, isSet: true}
}

func (v NullableInlineObject76) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject76) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


