/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications struct for NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications
type NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications struct {
	// Id of the major application, or a list of NBAR Application Category or Application selections
	Id *string `json:"id,omitempty"`
	// Name of the major application or application category selected
	Name *string `json:"name,omitempty"`
	// app type (major or nbar)
	Type *string `json:"type,omitempty"`
}

// NewNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications instantiates a new NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications() *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications {
	this := NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications{}
	return &this
}

// NewNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplicationsWithDefaults instantiates a new NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplicationsWithDefaults() *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications {
	this := NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) SetType(v string) {
	o.Type = &v
}

func (o NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications struct {
	value *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) Get() *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) Set(val *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications(val *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) *NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications {
	return &NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


