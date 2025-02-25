/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200228 struct for InlineResponse200228
type InlineResponse200228 struct {
	// Network identifier
	NetworkId *string `json:"networkId,omitempty"`
	// Network name
	Name *string `json:"name,omitempty"`
	// Uplink usage
	ByUplink []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink `json:"byUplink,omitempty"`
}

// NewInlineResponse200228 instantiates a new InlineResponse200228 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200228() *InlineResponse200228 {
	this := InlineResponse200228{}
	return &this
}

// NewInlineResponse200228WithDefaults instantiates a new InlineResponse200228 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200228WithDefaults() *InlineResponse200228 {
	this := InlineResponse200228{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200228) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200228) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200228) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200228) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200228) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200228) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200228) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200228) SetName(v string) {
	o.Name = &v
}

// GetByUplink returns the ByUplink field value if set, zero value otherwise.
func (o *InlineResponse200228) GetByUplink() []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink {
	if o == nil || isNil(o.ByUplink) {
		var ret []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink
		return ret
	}
	return o.ByUplink
}

// GetByUplinkOk returns a tuple with the ByUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200228) GetByUplinkOk() ([]OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink, bool) {
	if o == nil || isNil(o.ByUplink) {
    return nil, false
	}
	return o.ByUplink, true
}

// HasByUplink returns a boolean if a field has been set.
func (o *InlineResponse200228) HasByUplink() bool {
	if o != nil && !isNil(o.ByUplink) {
		return true
	}

	return false
}

// SetByUplink gets a reference to the given []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink and assigns it to the ByUplink field.
func (o *InlineResponse200228) SetByUplink(v []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) {
	o.ByUplink = v
}

func (o InlineResponse200228) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200228 struct {
	value *InlineResponse200228
	isSet bool
}

func (v NullableInlineResponse200228) Get() *InlineResponse200228 {
	return v.value
}

func (v *NullableInlineResponse200228) Set(val *InlineResponse200228) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200228) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200228) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200228(val *InlineResponse200228) *NullableInlineResponse200228 {
	return &NullableInlineResponse200228{value: val, isSet: true}
}

func (v NullableInlineResponse200228) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200228) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


