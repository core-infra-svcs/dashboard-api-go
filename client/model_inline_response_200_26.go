/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20026 struct for InlineResponse20026
type InlineResponse20026 struct {
	// DHCP Lease time for all MG in the network.
	DhcpLeaseTime *string `json:"dhcpLeaseTime,omitempty"`
	// DNS name servers mode for all MG in the network.
	DnsNameservers *string `json:"dnsNameservers,omitempty"`
	// List of fixed IPs representing the the DNS Name servers when the mode is 'custom'.
	DnsCustomNameservers []string `json:"dnsCustomNameservers,omitempty"`
}

// NewInlineResponse20026 instantiates a new InlineResponse20026 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20026() *InlineResponse20026 {
	this := InlineResponse20026{}
	return &this
}

// NewInlineResponse20026WithDefaults instantiates a new InlineResponse20026 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20026WithDefaults() *InlineResponse20026 {
	this := InlineResponse20026{}
	return &this
}

// GetDhcpLeaseTime returns the DhcpLeaseTime field value if set, zero value otherwise.
func (o *InlineResponse20026) GetDhcpLeaseTime() string {
	if o == nil || isNil(o.DhcpLeaseTime) {
		var ret string
		return ret
	}
	return *o.DhcpLeaseTime
}

// GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026) GetDhcpLeaseTimeOk() (*string, bool) {
	if o == nil || isNil(o.DhcpLeaseTime) {
    return nil, false
	}
	return o.DhcpLeaseTime, true
}

// HasDhcpLeaseTime returns a boolean if a field has been set.
func (o *InlineResponse20026) HasDhcpLeaseTime() bool {
	if o != nil && !isNil(o.DhcpLeaseTime) {
		return true
	}

	return false
}

// SetDhcpLeaseTime gets a reference to the given string and assigns it to the DhcpLeaseTime field.
func (o *InlineResponse20026) SetDhcpLeaseTime(v string) {
	o.DhcpLeaseTime = &v
}

// GetDnsNameservers returns the DnsNameservers field value if set, zero value otherwise.
func (o *InlineResponse20026) GetDnsNameservers() string {
	if o == nil || isNil(o.DnsNameservers) {
		var ret string
		return ret
	}
	return *o.DnsNameservers
}

// GetDnsNameserversOk returns a tuple with the DnsNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026) GetDnsNameserversOk() (*string, bool) {
	if o == nil || isNil(o.DnsNameservers) {
    return nil, false
	}
	return o.DnsNameservers, true
}

// HasDnsNameservers returns a boolean if a field has been set.
func (o *InlineResponse20026) HasDnsNameservers() bool {
	if o != nil && !isNil(o.DnsNameservers) {
		return true
	}

	return false
}

// SetDnsNameservers gets a reference to the given string and assigns it to the DnsNameservers field.
func (o *InlineResponse20026) SetDnsNameservers(v string) {
	o.DnsNameservers = &v
}

// GetDnsCustomNameservers returns the DnsCustomNameservers field value if set, zero value otherwise.
func (o *InlineResponse20026) GetDnsCustomNameservers() []string {
	if o == nil || isNil(o.DnsCustomNameservers) {
		var ret []string
		return ret
	}
	return o.DnsCustomNameservers
}

// GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026) GetDnsCustomNameserversOk() ([]string, bool) {
	if o == nil || isNil(o.DnsCustomNameservers) {
    return nil, false
	}
	return o.DnsCustomNameservers, true
}

// HasDnsCustomNameservers returns a boolean if a field has been set.
func (o *InlineResponse20026) HasDnsCustomNameservers() bool {
	if o != nil && !isNil(o.DnsCustomNameservers) {
		return true
	}

	return false
}

// SetDnsCustomNameservers gets a reference to the given []string and assigns it to the DnsCustomNameservers field.
func (o *InlineResponse20026) SetDnsCustomNameservers(v []string) {
	o.DnsCustomNameservers = v
}

func (o InlineResponse20026) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DhcpLeaseTime) {
		toSerialize["dhcpLeaseTime"] = o.DhcpLeaseTime
	}
	if !isNil(o.DnsNameservers) {
		toSerialize["dnsNameservers"] = o.DnsNameservers
	}
	if !isNil(o.DnsCustomNameservers) {
		toSerialize["dnsCustomNameservers"] = o.DnsCustomNameservers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20026 struct {
	value *InlineResponse20026
	isSet bool
}

func (v NullableInlineResponse20026) Get() *InlineResponse20026 {
	return v.value
}

func (v *NullableInlineResponse20026) Set(val *InlineResponse20026) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20026) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20026) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20026(val *InlineResponse20026) *NullableInlineResponse20026 {
	return &NullableInlineResponse20026{value: val, isSet: true}
}

func (v NullableInlineResponse20026) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20026) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


