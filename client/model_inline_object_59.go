/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject59 struct for InlineObject59
type InlineObject59 struct {
	// DHCP Lease time for all MG of the network. It can be '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'.
	DhcpLeaseTime *string `json:"dhcpLeaseTime,omitempty"`
	// DNS name servers mode for all MG of the network. It can take 4 different values: 'upstream_dns', 'google_dns', 'opendns', 'custom'.
	DnsNameservers *string `json:"dnsNameservers,omitempty"`
	// list of fixed IP representing the the DNS Name servers when the mode is 'custom'
	DnsCustomNameservers *[]string `json:"dnsCustomNameservers,omitempty"`
}

// NewInlineObject59 instantiates a new InlineObject59 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject59() *InlineObject59 {
	this := InlineObject59{}
	return &this
}

// NewInlineObject59WithDefaults instantiates a new InlineObject59 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject59WithDefaults() *InlineObject59 {
	this := InlineObject59{}
	return &this
}

// GetDhcpLeaseTime returns the DhcpLeaseTime field value if set, zero value otherwise.
func (o *InlineObject59) GetDhcpLeaseTime() string {
	if o == nil || o.DhcpLeaseTime == nil {
		var ret string
		return ret
	}
	return *o.DhcpLeaseTime
}

// GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject59) GetDhcpLeaseTimeOk() (*string, bool) {
	if o == nil || o.DhcpLeaseTime == nil {
		return nil, false
	}
	return o.DhcpLeaseTime, true
}

// HasDhcpLeaseTime returns a boolean if a field has been set.
func (o *InlineObject59) HasDhcpLeaseTime() bool {
	if o != nil && o.DhcpLeaseTime != nil {
		return true
	}

	return false
}

// SetDhcpLeaseTime gets a reference to the given string and assigns it to the DhcpLeaseTime field.
func (o *InlineObject59) SetDhcpLeaseTime(v string) {
	o.DhcpLeaseTime = &v
}

// GetDnsNameservers returns the DnsNameservers field value if set, zero value otherwise.
func (o *InlineObject59) GetDnsNameservers() string {
	if o == nil || o.DnsNameservers == nil {
		var ret string
		return ret
	}
	return *o.DnsNameservers
}

// GetDnsNameserversOk returns a tuple with the DnsNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject59) GetDnsNameserversOk() (*string, bool) {
	if o == nil || o.DnsNameservers == nil {
		return nil, false
	}
	return o.DnsNameservers, true
}

// HasDnsNameservers returns a boolean if a field has been set.
func (o *InlineObject59) HasDnsNameservers() bool {
	if o != nil && o.DnsNameservers != nil {
		return true
	}

	return false
}

// SetDnsNameservers gets a reference to the given string and assigns it to the DnsNameservers field.
func (o *InlineObject59) SetDnsNameservers(v string) {
	o.DnsNameservers = &v
}

// GetDnsCustomNameservers returns the DnsCustomNameservers field value if set, zero value otherwise.
func (o *InlineObject59) GetDnsCustomNameservers() []string {
	if o == nil || o.DnsCustomNameservers == nil {
		var ret []string
		return ret
	}
	return *o.DnsCustomNameservers
}

// GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject59) GetDnsCustomNameserversOk() (*[]string, bool) {
	if o == nil || o.DnsCustomNameservers == nil {
		return nil, false
	}
	return o.DnsCustomNameservers, true
}

// HasDnsCustomNameservers returns a boolean if a field has been set.
func (o *InlineObject59) HasDnsCustomNameservers() bool {
	if o != nil && o.DnsCustomNameservers != nil {
		return true
	}

	return false
}

// SetDnsCustomNameservers gets a reference to the given []string and assigns it to the DnsCustomNameservers field.
func (o *InlineObject59) SetDnsCustomNameservers(v []string) {
	o.DnsCustomNameservers = &v
}

func (o InlineObject59) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DhcpLeaseTime != nil {
		toSerialize["dhcpLeaseTime"] = o.DhcpLeaseTime
	}
	if o.DnsNameservers != nil {
		toSerialize["dnsNameservers"] = o.DnsNameservers
	}
	if o.DnsCustomNameservers != nil {
		toSerialize["dnsCustomNameservers"] = o.DnsCustomNameservers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject59 struct {
	value *InlineObject59
	isSet bool
}

func (v NullableInlineObject59) Get() *InlineObject59 {
	return v.value
}

func (v *NullableInlineObject59) Set(val *InlineObject59) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject59) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject59) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject59(val *InlineObject59) *NullableInlineObject59 {
	return &NullableInlineObject59{value: val, isSet: true}
}

func (v NullableInlineObject59) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject59) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


