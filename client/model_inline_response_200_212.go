/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200212 struct for InlineResponse200212
type InlineResponse200212 struct {
	// Network identifier
	NetworkId *string `json:"networkId,omitempty"`
	// Network name
	Name *string `json:"name,omitempty"`
	// Uplink usage
	ByUplink []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink `json:"byUplink,omitempty"`
}

// NewInlineResponse200212 instantiates a new InlineResponse200212 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200212() *InlineResponse200212 {
	this := InlineResponse200212{}
	return &this
}

// NewInlineResponse200212WithDefaults instantiates a new InlineResponse200212 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200212WithDefaults() *InlineResponse200212 {
	this := InlineResponse200212{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200212) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200212) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200212) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200212) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200212) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200212) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200212) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200212) SetName(v string) {
	o.Name = &v
}

// GetByUplink returns the ByUplink field value if set, zero value otherwise.
func (o *InlineResponse200212) GetByUplink() []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink {
	if o == nil || isNil(o.ByUplink) {
		var ret []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink
		return ret
	}
	return o.ByUplink
}

// GetByUplinkOk returns a tuple with the ByUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200212) GetByUplinkOk() ([]OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink, bool) {
	if o == nil || isNil(o.ByUplink) {
    return nil, false
	}
	return o.ByUplink, true
}

// HasByUplink returns a boolean if a field has been set.
func (o *InlineResponse200212) HasByUplink() bool {
	if o != nil && !isNil(o.ByUplink) {
		return true
	}

	return false
}

// SetByUplink gets a reference to the given []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink and assigns it to the ByUplink field.
func (o *InlineResponse200212) SetByUplink(v []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) {
	o.ByUplink = v
}

func (o InlineResponse200212) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ByUplink) {
		toSerialize["byUplink"] = o.ByUplink
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200212 struct {
	value *InlineResponse200212
	isSet bool
}

func (v NullableInlineResponse200212) Get() *InlineResponse200212 {
	return v.value
}

func (v *NullableInlineResponse200212) Set(val *InlineResponse200212) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200212) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200212) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200212(val *InlineResponse200212) *NullableInlineResponse200212 {
	return &NullableInlineResponse200212{value: val, isSet: true}
}

func (v NullableInlineResponse200212) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200212) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


