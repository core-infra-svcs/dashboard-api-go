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

// InlineResponse200197Rules struct for InlineResponse200197Rules
type InlineResponse200197Rules struct {
	// Desctiption of the bonjour forwarding rule
	Description *string `json:"description,omitempty"`
	// The ID of the service VLAN. Required
	VlanId *string `json:"vlanId,omitempty"`
	// A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AFP', 'AirPlay', 'Apple screen share', 'BitTorrent', 'Chromecast', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners', 'Spotify' and 'SSH'
	Services []string `json:"services,omitempty"`
}

// NewInlineResponse200197Rules instantiates a new InlineResponse200197Rules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200197Rules() *InlineResponse200197Rules {
	this := InlineResponse200197Rules{}
	return &this
}

// NewInlineResponse200197RulesWithDefaults instantiates a new InlineResponse200197Rules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200197RulesWithDefaults() *InlineResponse200197Rules {
	this := InlineResponse200197Rules{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse200197Rules) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200197Rules) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse200197Rules) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse200197Rules) SetDescription(v string) {
	o.Description = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *InlineResponse200197Rules) GetVlanId() string {
	if o == nil || isNil(o.VlanId) {
		var ret string
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200197Rules) GetVlanIdOk() (*string, bool) {
	if o == nil || isNil(o.VlanId) {
    return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *InlineResponse200197Rules) HasVlanId() bool {
	if o != nil && !isNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given string and assigns it to the VlanId field.
func (o *InlineResponse200197Rules) SetVlanId(v string) {
	o.VlanId = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *InlineResponse200197Rules) GetServices() []string {
	if o == nil || isNil(o.Services) {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200197Rules) GetServicesOk() ([]string, bool) {
	if o == nil || isNil(o.Services) {
    return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *InlineResponse200197Rules) HasServices() bool {
	if o != nil && !isNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *InlineResponse200197Rules) SetServices(v []string) {
	o.Services = v
}

func (o InlineResponse200197Rules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !isNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200197Rules struct {
	value *InlineResponse200197Rules
	isSet bool
}

func (v NullableInlineResponse200197Rules) Get() *InlineResponse200197Rules {
	return v.value
}

func (v *NullableInlineResponse200197Rules) Set(val *InlineResponse200197Rules) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200197Rules) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200197Rules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200197Rules(val *InlineResponse200197Rules) *NullableInlineResponse200197Rules {
	return &NullableInlineResponse200197Rules{value: val, isSet: true}
}

func (v NullableInlineResponse200197Rules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200197Rules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


