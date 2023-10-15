/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200121 struct for InlineResponse200121
type InlineResponse200121 struct {
	// Network identifier
	NetworkId *string `json:"networkId,omitempty"`
	// Network name
	Name *string `json:"name,omitempty"`
	// Uplink usage
	ByUplink []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink `json:"byUplink,omitempty"`
}

// NewInlineResponse200121 instantiates a new InlineResponse200121 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200121() *InlineResponse200121 {
	this := InlineResponse200121{}
	return &this
}

// NewInlineResponse200121WithDefaults instantiates a new InlineResponse200121 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200121WithDefaults() *InlineResponse200121 {
	this := InlineResponse200121{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200121) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200121) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200121) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200121) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200121) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200121) SetName(v string) {
	o.Name = &v
}

// GetByUplink returns the ByUplink field value if set, zero value otherwise.
func (o *InlineResponse200121) GetByUplink() []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink {
	if o == nil || isNil(o.ByUplink) {
		var ret []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink
		return ret
	}
	return o.ByUplink
}

// GetByUplinkOk returns a tuple with the ByUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121) GetByUplinkOk() ([]OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink, bool) {
	if o == nil || isNil(o.ByUplink) {
    return nil, false
	}
	return o.ByUplink, true
}

// HasByUplink returns a boolean if a field has been set.
func (o *InlineResponse200121) HasByUplink() bool {
	if o != nil && !isNil(o.ByUplink) {
		return true
	}

	return false
}

// SetByUplink gets a reference to the given []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink and assigns it to the ByUplink field.
func (o *InlineResponse200121) SetByUplink(v []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) {
	o.ByUplink = v
}

func (o InlineResponse200121) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200121 struct {
	value *InlineResponse200121
	isSet bool
}

func (v NullableInlineResponse200121) Get() *InlineResponse200121 {
	return v.value
}

func (v *NullableInlineResponse200121) Set(val *InlineResponse200121) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200121) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200121) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200121(val *InlineResponse200121) *NullableInlineResponse200121 {
	return &NullableInlineResponse200121{value: val, isSet: true}
}

func (v NullableInlineResponse200121) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200121) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


