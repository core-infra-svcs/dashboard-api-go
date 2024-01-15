/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs struct for NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs
type NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs struct {
	// The network ID of the hub.
	HubId string `json:"hubId"`
	// Only valid in 'spoke' mode. Indicates whether default route traffic should be sent to this hub.
	UseDefaultRoute *bool `json:"useDefaultRoute,omitempty"`
}

// NewNetworksNetworkIdApplianceVpnSiteToSiteVpnHubs instantiates a new NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceVpnSiteToSiteVpnHubs(hubId string) *NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs {
	this := NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs{}
	this.HubId = hubId
	return &this
}

// NewNetworksNetworkIdApplianceVpnSiteToSiteVpnHubsWithDefaults instantiates a new NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceVpnSiteToSiteVpnHubsWithDefaults() *NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs {
	this := NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs{}
	return &this
}

// GetHubId returns the HubId field value
func (o *NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) GetHubId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HubId
}

// GetHubIdOk returns a tuple with the HubId field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) GetHubIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HubId, true
}

// SetHubId sets field value
func (o *NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) SetHubId(v string) {
	o.HubId = v
}

// GetUseDefaultRoute returns the UseDefaultRoute field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) GetUseDefaultRoute() bool {
	if o == nil || isNil(o.UseDefaultRoute) {
		var ret bool
		return ret
	}
	return *o.UseDefaultRoute
}

// GetUseDefaultRouteOk returns a tuple with the UseDefaultRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) GetUseDefaultRouteOk() (*bool, bool) {
	if o == nil || isNil(o.UseDefaultRoute) {
    return nil, false
	}
	return o.UseDefaultRoute, true
}

// HasUseDefaultRoute returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) HasUseDefaultRoute() bool {
	if o != nil && !isNil(o.UseDefaultRoute) {
		return true
	}

	return false
}

// SetUseDefaultRoute gets a reference to the given bool and assigns it to the UseDefaultRoute field.
func (o *NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) SetUseDefaultRoute(v bool) {
	o.UseDefaultRoute = &v
}

func (o NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hubId"] = o.HubId
	}
	if !isNil(o.UseDefaultRoute) {
		toSerialize["useDefaultRoute"] = o.UseDefaultRoute
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnHubs struct {
	value *NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) Get() *NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) Set(val *NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceVpnSiteToSiteVpnHubs(val *NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) *NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnHubs {
	return &NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnHubs{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnHubs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


