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

// InlineResponse200310 struct for InlineResponse200310
type InlineResponse200310 struct {
	Downstream *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream `json:"downstream,omitempty"`
	Upstream *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream `json:"upstream,omitempty"`
	Client *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientClient `json:"client,omitempty"`
	Network *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork `json:"network,omitempty"`
}

// NewInlineResponse200310 instantiates a new InlineResponse200310 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200310() *InlineResponse200310 {
	this := InlineResponse200310{}
	return &this
}

// NewInlineResponse200310WithDefaults instantiates a new InlineResponse200310 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200310WithDefaults() *InlineResponse200310 {
	this := InlineResponse200310{}
	return &this
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *InlineResponse200310) GetDownstream() OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream {
	if o == nil || isNil(o.Downstream) {
		var ret OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200310) GetDownstreamOk() (*OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream, bool) {
	if o == nil || isNil(o.Downstream) {
    return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *InlineResponse200310) HasDownstream() bool {
	if o != nil && !isNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream and assigns it to the Downstream field.
func (o *InlineResponse200310) SetDownstream(v OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) {
	o.Downstream = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *InlineResponse200310) GetUpstream() OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream {
	if o == nil || isNil(o.Upstream) {
		var ret OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200310) GetUpstreamOk() (*OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream, bool) {
	if o == nil || isNil(o.Upstream) {
    return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *InlineResponse200310) HasUpstream() bool {
	if o != nil && !isNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream and assigns it to the Upstream field.
func (o *InlineResponse200310) SetUpstream(v OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) {
	o.Upstream = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *InlineResponse200310) GetClient() OrganizationsOrganizationIdWirelessDevicesPacketLossByClientClient {
	if o == nil || isNil(o.Client) {
		var ret OrganizationsOrganizationIdWirelessDevicesPacketLossByClientClient
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200310) GetClientOk() (*OrganizationsOrganizationIdWirelessDevicesPacketLossByClientClient, bool) {
	if o == nil || isNil(o.Client) {
    return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *InlineResponse200310) HasClient() bool {
	if o != nil && !isNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given OrganizationsOrganizationIdWirelessDevicesPacketLossByClientClient and assigns it to the Client field.
func (o *InlineResponse200310) SetClient(v OrganizationsOrganizationIdWirelessDevicesPacketLossByClientClient) {
	o.Client = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200310) GetNetwork() OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200310) GetNetworkOk() (*OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200310) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork and assigns it to the Network field.
func (o *InlineResponse200310) SetNetwork(v OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) {
	o.Network = &v
}

func (o InlineResponse200310) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Downstream) {
		toSerialize["downstream"] = o.Downstream
	}
	if !isNil(o.Upstream) {
		toSerialize["upstream"] = o.Upstream
	}
	if !isNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200310 struct {
	value *InlineResponse200310
	isSet bool
}

func (v NullableInlineResponse200310) Get() *InlineResponse200310 {
	return v.value
}

func (v *NullableInlineResponse200310) Set(val *InlineResponse200310) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200310) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200310) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200310(val *InlineResponse200310) *NullableInlineResponse200310 {
	return &NullableInlineResponse200310{value: val, isSet: true}
}

func (v NullableInlineResponse200310) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200310) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


