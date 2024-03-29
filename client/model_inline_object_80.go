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

// InlineObject80 struct for InlineObject80
type InlineObject80 struct {
	// DHCP Lease time for all MG of the network. Possible values are '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'.
	DhcpLeaseTime *string `json:"dhcpLeaseTime,omitempty"`
	// DNS name servers mode for all MG of the network. Possible values are: 'upstream_dns', 'google_dns', 'opendns', 'custom'.
	DnsNameservers *string `json:"dnsNameservers,omitempty"`
	// list of fixed IPs representing the the DNS Name servers when the mode is 'custom'
	DnsCustomNameservers []string `json:"dnsCustomNameservers,omitempty"`
}

// NewInlineObject80 instantiates a new InlineObject80 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject80() *InlineObject80 {
	this := InlineObject80{}
	return &this
}

// NewInlineObject80WithDefaults instantiates a new InlineObject80 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject80WithDefaults() *InlineObject80 {
	this := InlineObject80{}
	return &this
}

// GetDhcpLeaseTime returns the DhcpLeaseTime field value if set, zero value otherwise.
func (o *InlineObject80) GetDhcpLeaseTime() string {
	if o == nil || isNil(o.DhcpLeaseTime) {
		var ret string
		return ret
	}
	return *o.DhcpLeaseTime
}

// GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject80) GetDhcpLeaseTimeOk() (*string, bool) {
	if o == nil || isNil(o.DhcpLeaseTime) {
    return nil, false
	}
	return o.DhcpLeaseTime, true
}

// HasDhcpLeaseTime returns a boolean if a field has been set.
func (o *InlineObject80) HasDhcpLeaseTime() bool {
	if o != nil && !isNil(o.DhcpLeaseTime) {
		return true
	}

	return false
}

// SetDhcpLeaseTime gets a reference to the given string and assigns it to the DhcpLeaseTime field.
func (o *InlineObject80) SetDhcpLeaseTime(v string) {
	o.DhcpLeaseTime = &v
}

// GetDnsNameservers returns the DnsNameservers field value if set, zero value otherwise.
func (o *InlineObject80) GetDnsNameservers() string {
	if o == nil || isNil(o.DnsNameservers) {
		var ret string
		return ret
	}
	return *o.DnsNameservers
}

// GetDnsNameserversOk returns a tuple with the DnsNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject80) GetDnsNameserversOk() (*string, bool) {
	if o == nil || isNil(o.DnsNameservers) {
    return nil, false
	}
	return o.DnsNameservers, true
}

// HasDnsNameservers returns a boolean if a field has been set.
func (o *InlineObject80) HasDnsNameservers() bool {
	if o != nil && !isNil(o.DnsNameservers) {
		return true
	}

	return false
}

// SetDnsNameservers gets a reference to the given string and assigns it to the DnsNameservers field.
func (o *InlineObject80) SetDnsNameservers(v string) {
	o.DnsNameservers = &v
}

// GetDnsCustomNameservers returns the DnsCustomNameservers field value if set, zero value otherwise.
func (o *InlineObject80) GetDnsCustomNameservers() []string {
	if o == nil || isNil(o.DnsCustomNameservers) {
		var ret []string
		return ret
	}
	return o.DnsCustomNameservers
}

// GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject80) GetDnsCustomNameserversOk() ([]string, bool) {
	if o == nil || isNil(o.DnsCustomNameservers) {
    return nil, false
	}
	return o.DnsCustomNameservers, true
}

// HasDnsCustomNameservers returns a boolean if a field has been set.
func (o *InlineObject80) HasDnsCustomNameservers() bool {
	if o != nil && !isNil(o.DnsCustomNameservers) {
		return true
	}

	return false
}

// SetDnsCustomNameservers gets a reference to the given []string and assigns it to the DnsCustomNameservers field.
func (o *InlineObject80) SetDnsCustomNameservers(v []string) {
	o.DnsCustomNameservers = v
}

func (o InlineObject80) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject80 struct {
	value *InlineObject80
	isSet bool
}

func (v NullableInlineObject80) Get() *InlineObject80 {
	return v.value
}

func (v *NullableInlineObject80) Set(val *InlineObject80) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject80) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject80) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject80(val *InlineObject80) *NullableInlineObject80 {
	return &NullableInlineObject80{value: val, isSet: true}
}

func (v NullableInlineObject80) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject80) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


