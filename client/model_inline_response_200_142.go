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

// InlineResponse200142 struct for InlineResponse200142
type InlineResponse200142 struct {
	// Network identifier
	NetworkId *string `json:"networkId,omitempty"`
	// Network name
	Name *string `json:"name,omitempty"`
	// Uplink usage
	ByUplink []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink `json:"byUplink,omitempty"`
}

// NewInlineResponse200142 instantiates a new InlineResponse200142 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200142() *InlineResponse200142 {
	this := InlineResponse200142{}
	return &this
}

// NewInlineResponse200142WithDefaults instantiates a new InlineResponse200142 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200142WithDefaults() *InlineResponse200142 {
	this := InlineResponse200142{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200142) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200142) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200142) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200142) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200142) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200142) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200142) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200142) SetName(v string) {
	o.Name = &v
}

// GetByUplink returns the ByUplink field value if set, zero value otherwise.
func (o *InlineResponse200142) GetByUplink() []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink {
	if o == nil || isNil(o.ByUplink) {
		var ret []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink
		return ret
	}
	return o.ByUplink
}

// GetByUplinkOk returns a tuple with the ByUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200142) GetByUplinkOk() ([]OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink, bool) {
	if o == nil || isNil(o.ByUplink) {
    return nil, false
	}
	return o.ByUplink, true
}

// HasByUplink returns a boolean if a field has been set.
func (o *InlineResponse200142) HasByUplink() bool {
	if o != nil && !isNil(o.ByUplink) {
		return true
	}

	return false
}

// SetByUplink gets a reference to the given []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink and assigns it to the ByUplink field.
func (o *InlineResponse200142) SetByUplink(v []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) {
	o.ByUplink = v
}

func (o InlineResponse200142) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200142 struct {
	value *InlineResponse200142
	isSet bool
}

func (v NullableInlineResponse200142) Get() *InlineResponse200142 {
	return v.value
}

func (v *NullableInlineResponse200142) Set(val *InlineResponse200142) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200142) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200142) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200142(val *InlineResponse200142) *NullableInlineResponse200142 {
	return &NullableInlineResponse200142{value: val, isSet: true}
}

func (v NullableInlineResponse200142) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200142) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


