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

// OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork struct for OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork
type OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork struct {
	// Network identifier
	NetworkId *string `json:"networkId,omitempty"`
	// Number of useful information bits delivered over a network per unit of time
	Goodput *int32 `json:"goodput,omitempty"`
	// Duration of the response, in milliseconds
	ResponseDuration *int32 `json:"responseDuration,omitempty"`
}

// NewOrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork instantiates a new OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork() *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork {
	this := OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork{}
	return &this
}

// NewOrganizationsOrganizationIdInsightApplicationsThresholdsByNetworkWithDefaults instantiates a new OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdInsightApplicationsThresholdsByNetworkWithDefaults() *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork {
	this := OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetGoodput returns the Goodput field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) GetGoodput() int32 {
	if o == nil || isNil(o.Goodput) {
		var ret int32
		return ret
	}
	return *o.Goodput
}

// GetGoodputOk returns a tuple with the Goodput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) GetGoodputOk() (*int32, bool) {
	if o == nil || isNil(o.Goodput) {
    return nil, false
	}
	return o.Goodput, true
}

// HasGoodput returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) HasGoodput() bool {
	if o != nil && !isNil(o.Goodput) {
		return true
	}

	return false
}

// SetGoodput gets a reference to the given int32 and assigns it to the Goodput field.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) SetGoodput(v int32) {
	o.Goodput = &v
}

// GetResponseDuration returns the ResponseDuration field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) GetResponseDuration() int32 {
	if o == nil || isNil(o.ResponseDuration) {
		var ret int32
		return ret
	}
	return *o.ResponseDuration
}

// GetResponseDurationOk returns a tuple with the ResponseDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) GetResponseDurationOk() (*int32, bool) {
	if o == nil || isNil(o.ResponseDuration) {
    return nil, false
	}
	return o.ResponseDuration, true
}

// HasResponseDuration returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) HasResponseDuration() bool {
	if o != nil && !isNil(o.ResponseDuration) {
		return true
	}

	return false
}

// SetResponseDuration gets a reference to the given int32 and assigns it to the ResponseDuration field.
func (o *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) SetResponseDuration(v int32) {
	o.ResponseDuration = &v
}

func (o OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Goodput) {
		toSerialize["goodput"] = o.Goodput
	}
	if !isNil(o.ResponseDuration) {
		toSerialize["responseDuration"] = o.ResponseDuration
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork struct {
	value *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork
	isSet bool
}

func (v NullableOrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) Get() *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) Set(val *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork(val *OrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) *NullableOrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork {
	return &NullableOrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdInsightApplicationsThresholdsByNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


