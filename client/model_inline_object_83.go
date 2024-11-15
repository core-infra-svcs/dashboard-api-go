/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject83 struct for InlineObject83
type InlineObject83 struct {
	// DHCP Lease time for all MG of the network. Possible values are '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'.
	DhcpLeaseTime *string `json:"dhcpLeaseTime,omitempty"`
	// DNS name servers mode for all MG of the network. Possible values are: 'upstream_dns', 'google_dns', 'opendns', 'custom'.
	DnsNameservers *string `json:"dnsNameservers,omitempty"`
	// list of fixed IPs representing the the DNS Name servers when the mode is 'custom'
	DnsCustomNameservers []string `json:"dnsCustomNameservers,omitempty"`
}

// NewInlineObject83 instantiates a new InlineObject83 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject83() *InlineObject83 {
	this := InlineObject83{}
	return &this
}

// NewInlineObject83WithDefaults instantiates a new InlineObject83 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject83WithDefaults() *InlineObject83 {
	this := InlineObject83{}
	return &this
}

// GetDhcpLeaseTime returns the DhcpLeaseTime field value if set, zero value otherwise.
func (o *InlineObject83) GetDhcpLeaseTime() string {
	if o == nil || isNil(o.DhcpLeaseTime) {
		var ret string
		return ret
	}
	return *o.DhcpLeaseTime
}

// GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetDhcpLeaseTimeOk() (*string, bool) {
	if o == nil || isNil(o.DhcpLeaseTime) {
    return nil, false
	}
	return o.DhcpLeaseTime, true
}

// HasDhcpLeaseTime returns a boolean if a field has been set.
func (o *InlineObject83) HasDhcpLeaseTime() bool {
	if o != nil && !isNil(o.DhcpLeaseTime) {
		return true
	}

	return false
}

// SetDhcpLeaseTime gets a reference to the given string and assigns it to the DhcpLeaseTime field.
func (o *InlineObject83) SetDhcpLeaseTime(v string) {
	o.DhcpLeaseTime = &v
}

// GetDnsNameservers returns the DnsNameservers field value if set, zero value otherwise.
func (o *InlineObject83) GetDnsNameservers() string {
	if o == nil || isNil(o.DnsNameservers) {
		var ret string
		return ret
	}
	return *o.DnsNameservers
}

// GetDnsNameserversOk returns a tuple with the DnsNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetDnsNameserversOk() (*string, bool) {
	if o == nil || isNil(o.DnsNameservers) {
    return nil, false
	}
	return o.DnsNameservers, true
}

// HasDnsNameservers returns a boolean if a field has been set.
func (o *InlineObject83) HasDnsNameservers() bool {
	if o != nil && !isNil(o.DnsNameservers) {
		return true
	}

	return false
}

// SetDnsNameservers gets a reference to the given string and assigns it to the DnsNameservers field.
func (o *InlineObject83) SetDnsNameservers(v string) {
	o.DnsNameservers = &v
}

// GetDnsCustomNameservers returns the DnsCustomNameservers field value if set, zero value otherwise.
func (o *InlineObject83) GetDnsCustomNameservers() []string {
	if o == nil || isNil(o.DnsCustomNameservers) {
		var ret []string
		return ret
	}
	return o.DnsCustomNameservers
}

// GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetDnsCustomNameserversOk() ([]string, bool) {
	if o == nil || isNil(o.DnsCustomNameservers) {
    return nil, false
	}
	return o.DnsCustomNameservers, true
}

// HasDnsCustomNameservers returns a boolean if a field has been set.
func (o *InlineObject83) HasDnsCustomNameservers() bool {
	if o != nil && !isNil(o.DnsCustomNameservers) {
		return true
	}

	return false
}

// SetDnsCustomNameservers gets a reference to the given []string and assigns it to the DnsCustomNameservers field.
func (o *InlineObject83) SetDnsCustomNameservers(v []string) {
	o.DnsCustomNameservers = v
}

func (o InlineObject83) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject83 struct {
	value *InlineObject83
	isSet bool
}

func (v NullableInlineObject83) Get() *InlineObject83 {
	return v.value
}

func (v *NullableInlineObject83) Set(val *InlineObject83) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject83) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject83) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject83(val *InlineObject83) *NullableInlineObject83 {
	return &NullableInlineObject83{value: val, isSet: true}
}

func (v NullableInlineObject83) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject83) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


