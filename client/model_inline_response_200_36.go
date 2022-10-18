/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20036 struct for InlineResponse20036
type InlineResponse20036 struct {
	// The IP address of the DCHP Server.
	DhcpServer *string `json:"dhcpServer,omitempty"`
	// The IP address of the DNS Server.
	DnsServer *string `json:"dnsServer,omitempty"`
	// The IP address of the Gateway.
	Gateway *string `json:"gateway,omitempty"`
	// The Meraki Id of the network adapter record.
	Id *string `json:"id,omitempty"`
	// The IP address of the network adapter.
	Ip *string `json:"ip,omitempty"`
	// The MAC associated with the network adapter.
	Mac *string `json:"mac,omitempty"`
	// The name of the newtwork adapter.
	Name *string `json:"name,omitempty"`
	// The subnet for the network adapter.
	Subnet *string `json:"subnet,omitempty"`
}

// NewInlineResponse20036 instantiates a new InlineResponse20036 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20036() *InlineResponse20036 {
	this := InlineResponse20036{}
	return &this
}

// NewInlineResponse20036WithDefaults instantiates a new InlineResponse20036 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20036WithDefaults() *InlineResponse20036 {
	this := InlineResponse20036{}
	return &this
}

// GetDhcpServer returns the DhcpServer field value if set, zero value otherwise.
func (o *InlineResponse20036) GetDhcpServer() string {
	if o == nil || o.DhcpServer == nil {
		var ret string
		return ret
	}
	return *o.DhcpServer
}

// GetDhcpServerOk returns a tuple with the DhcpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetDhcpServerOk() (*string, bool) {
	if o == nil || o.DhcpServer == nil {
		return nil, false
	}
	return o.DhcpServer, true
}

// HasDhcpServer returns a boolean if a field has been set.
func (o *InlineResponse20036) HasDhcpServer() bool {
	if o != nil && o.DhcpServer != nil {
		return true
	}

	return false
}

// SetDhcpServer gets a reference to the given string and assigns it to the DhcpServer field.
func (o *InlineResponse20036) SetDhcpServer(v string) {
	o.DhcpServer = &v
}

// GetDnsServer returns the DnsServer field value if set, zero value otherwise.
func (o *InlineResponse20036) GetDnsServer() string {
	if o == nil || o.DnsServer == nil {
		var ret string
		return ret
	}
	return *o.DnsServer
}

// GetDnsServerOk returns a tuple with the DnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetDnsServerOk() (*string, bool) {
	if o == nil || o.DnsServer == nil {
		return nil, false
	}
	return o.DnsServer, true
}

// HasDnsServer returns a boolean if a field has been set.
func (o *InlineResponse20036) HasDnsServer() bool {
	if o != nil && o.DnsServer != nil {
		return true
	}

	return false
}

// SetDnsServer gets a reference to the given string and assigns it to the DnsServer field.
func (o *InlineResponse20036) SetDnsServer(v string) {
	o.DnsServer = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *InlineResponse20036) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *InlineResponse20036) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *InlineResponse20036) SetGateway(v string) {
	o.Gateway = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20036) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20036) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20036) SetId(v string) {
	o.Id = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineResponse20036) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineResponse20036) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineResponse20036) SetIp(v string) {
	o.Ip = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse20036) GetMac() string {
	if o == nil || o.Mac == nil {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetMacOk() (*string, bool) {
	if o == nil || o.Mac == nil {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse20036) HasMac() bool {
	if o != nil && o.Mac != nil {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse20036) SetMac(v string) {
	o.Mac = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20036) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20036) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20036) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineResponse20036) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineResponse20036) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineResponse20036) SetSubnet(v string) {
	o.Subnet = &v
}

func (o InlineResponse20036) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DhcpServer != nil {
		toSerialize["dhcpServer"] = o.DhcpServer
	}
	if o.DnsServer != nil {
		toSerialize["dnsServer"] = o.DnsServer
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Mac != nil {
		toSerialize["mac"] = o.Mac
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Subnet != nil {
		toSerialize["subnet"] = o.Subnet
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20036 struct {
	value *InlineResponse20036
	isSet bool
}

func (v NullableInlineResponse20036) Get() *InlineResponse20036 {
	return v.value
}

func (v *NullableInlineResponse20036) Set(val *InlineResponse20036) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20036) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20036) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20036(val *InlineResponse20036) *NullableInlineResponse20036 {
	return &NullableInlineResponse20036{value: val, isSet: true}
}

func (v NullableInlineResponse20036) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20036) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


