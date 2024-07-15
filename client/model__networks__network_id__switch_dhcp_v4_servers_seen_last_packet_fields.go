/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields DHCP-specific fields of the packet.
type NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields struct {
	// Operation code of the packet.
	Op *int32 `json:"op,omitempty"`
	// Hardware type code of the packet.
	Htype *int32 `json:"htype,omitempty"`
	// Hardware length of the packet.
	Hlen *int32 `json:"hlen,omitempty"`
	// Number of hops the packet took.
	Hops *int32 `json:"hops,omitempty"`
	// Transaction id of the packet.
	Xid *string `json:"xid,omitempty"`
	// Number of seconds since receiving the packet.
	Secs *int32 `json:"secs,omitempty"`
	// Packet flags.
	Flags *string `json:"flags,omitempty"`
	// Client IP address of the packet.
	Ciaddr *string `json:"ciaddr,omitempty"`
	// Assigned IP address of the packet.
	Yiaddr *string `json:"yiaddr,omitempty"`
	// Server IP address of the packet.
	Siaddr *string `json:"siaddr,omitempty"`
	// Gateway IP address of the packet.
	Giaddr *string `json:"giaddr,omitempty"`
	// Client hardware address of the packet.
	Chaddr *string `json:"chaddr,omitempty"`
	// Server identifier address of the packet.
	Sname *string `json:"sname,omitempty"`
	// Magic cookie of the packet.
	MagicCookie *string `json:"magicCookie,omitempty"`
	// Additional DHCP options of the packet.
	Options []NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions `json:"options,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsWithDefaults() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields{}
	return &this
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetOp() int32 {
	if o == nil || isNil(o.Op) {
		var ret int32
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetOpOk() (*int32, bool) {
	if o == nil || isNil(o.Op) {
    return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasOp() bool {
	if o != nil && !isNil(o.Op) {
		return true
	}

	return false
}

// SetOp gets a reference to the given int32 and assigns it to the Op field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetOp(v int32) {
	o.Op = &v
}

// GetHtype returns the Htype field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetHtype() int32 {
	if o == nil || isNil(o.Htype) {
		var ret int32
		return ret
	}
	return *o.Htype
}

// GetHtypeOk returns a tuple with the Htype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetHtypeOk() (*int32, bool) {
	if o == nil || isNil(o.Htype) {
    return nil, false
	}
	return o.Htype, true
}

// HasHtype returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasHtype() bool {
	if o != nil && !isNil(o.Htype) {
		return true
	}

	return false
}

// SetHtype gets a reference to the given int32 and assigns it to the Htype field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetHtype(v int32) {
	o.Htype = &v
}

// GetHlen returns the Hlen field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetHlen() int32 {
	if o == nil || isNil(o.Hlen) {
		var ret int32
		return ret
	}
	return *o.Hlen
}

// GetHlenOk returns a tuple with the Hlen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetHlenOk() (*int32, bool) {
	if o == nil || isNil(o.Hlen) {
    return nil, false
	}
	return o.Hlen, true
}

// HasHlen returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasHlen() bool {
	if o != nil && !isNil(o.Hlen) {
		return true
	}

	return false
}

// SetHlen gets a reference to the given int32 and assigns it to the Hlen field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetHlen(v int32) {
	o.Hlen = &v
}

// GetHops returns the Hops field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetHops() int32 {
	if o == nil || isNil(o.Hops) {
		var ret int32
		return ret
	}
	return *o.Hops
}

// GetHopsOk returns a tuple with the Hops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetHopsOk() (*int32, bool) {
	if o == nil || isNil(o.Hops) {
    return nil, false
	}
	return o.Hops, true
}

// HasHops returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasHops() bool {
	if o != nil && !isNil(o.Hops) {
		return true
	}

	return false
}

// SetHops gets a reference to the given int32 and assigns it to the Hops field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetHops(v int32) {
	o.Hops = &v
}

// GetXid returns the Xid field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetXid() string {
	if o == nil || isNil(o.Xid) {
		var ret string
		return ret
	}
	return *o.Xid
}

// GetXidOk returns a tuple with the Xid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetXidOk() (*string, bool) {
	if o == nil || isNil(o.Xid) {
    return nil, false
	}
	return o.Xid, true
}

// HasXid returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasXid() bool {
	if o != nil && !isNil(o.Xid) {
		return true
	}

	return false
}

// SetXid gets a reference to the given string and assigns it to the Xid field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetXid(v string) {
	o.Xid = &v
}

// GetSecs returns the Secs field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetSecs() int32 {
	if o == nil || isNil(o.Secs) {
		var ret int32
		return ret
	}
	return *o.Secs
}

// GetSecsOk returns a tuple with the Secs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetSecsOk() (*int32, bool) {
	if o == nil || isNil(o.Secs) {
    return nil, false
	}
	return o.Secs, true
}

// HasSecs returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasSecs() bool {
	if o != nil && !isNil(o.Secs) {
		return true
	}

	return false
}

// SetSecs gets a reference to the given int32 and assigns it to the Secs field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetSecs(v int32) {
	o.Secs = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetFlags() string {
	if o == nil || isNil(o.Flags) {
		var ret string
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetFlagsOk() (*string, bool) {
	if o == nil || isNil(o.Flags) {
    return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasFlags() bool {
	if o != nil && !isNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given string and assigns it to the Flags field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetFlags(v string) {
	o.Flags = &v
}

// GetCiaddr returns the Ciaddr field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetCiaddr() string {
	if o == nil || isNil(o.Ciaddr) {
		var ret string
		return ret
	}
	return *o.Ciaddr
}

// GetCiaddrOk returns a tuple with the Ciaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetCiaddrOk() (*string, bool) {
	if o == nil || isNil(o.Ciaddr) {
    return nil, false
	}
	return o.Ciaddr, true
}

// HasCiaddr returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasCiaddr() bool {
	if o != nil && !isNil(o.Ciaddr) {
		return true
	}

	return false
}

// SetCiaddr gets a reference to the given string and assigns it to the Ciaddr field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetCiaddr(v string) {
	o.Ciaddr = &v
}

// GetYiaddr returns the Yiaddr field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetYiaddr() string {
	if o == nil || isNil(o.Yiaddr) {
		var ret string
		return ret
	}
	return *o.Yiaddr
}

// GetYiaddrOk returns a tuple with the Yiaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetYiaddrOk() (*string, bool) {
	if o == nil || isNil(o.Yiaddr) {
    return nil, false
	}
	return o.Yiaddr, true
}

// HasYiaddr returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasYiaddr() bool {
	if o != nil && !isNil(o.Yiaddr) {
		return true
	}

	return false
}

// SetYiaddr gets a reference to the given string and assigns it to the Yiaddr field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetYiaddr(v string) {
	o.Yiaddr = &v
}

// GetSiaddr returns the Siaddr field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetSiaddr() string {
	if o == nil || isNil(o.Siaddr) {
		var ret string
		return ret
	}
	return *o.Siaddr
}

// GetSiaddrOk returns a tuple with the Siaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetSiaddrOk() (*string, bool) {
	if o == nil || isNil(o.Siaddr) {
    return nil, false
	}
	return o.Siaddr, true
}

// HasSiaddr returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasSiaddr() bool {
	if o != nil && !isNil(o.Siaddr) {
		return true
	}

	return false
}

// SetSiaddr gets a reference to the given string and assigns it to the Siaddr field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetSiaddr(v string) {
	o.Siaddr = &v
}

// GetGiaddr returns the Giaddr field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetGiaddr() string {
	if o == nil || isNil(o.Giaddr) {
		var ret string
		return ret
	}
	return *o.Giaddr
}

// GetGiaddrOk returns a tuple with the Giaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetGiaddrOk() (*string, bool) {
	if o == nil || isNil(o.Giaddr) {
    return nil, false
	}
	return o.Giaddr, true
}

// HasGiaddr returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasGiaddr() bool {
	if o != nil && !isNil(o.Giaddr) {
		return true
	}

	return false
}

// SetGiaddr gets a reference to the given string and assigns it to the Giaddr field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetGiaddr(v string) {
	o.Giaddr = &v
}

// GetChaddr returns the Chaddr field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetChaddr() string {
	if o == nil || isNil(o.Chaddr) {
		var ret string
		return ret
	}
	return *o.Chaddr
}

// GetChaddrOk returns a tuple with the Chaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetChaddrOk() (*string, bool) {
	if o == nil || isNil(o.Chaddr) {
    return nil, false
	}
	return o.Chaddr, true
}

// HasChaddr returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasChaddr() bool {
	if o != nil && !isNil(o.Chaddr) {
		return true
	}

	return false
}

// SetChaddr gets a reference to the given string and assigns it to the Chaddr field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetChaddr(v string) {
	o.Chaddr = &v
}

// GetSname returns the Sname field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetSname() string {
	if o == nil || isNil(o.Sname) {
		var ret string
		return ret
	}
	return *o.Sname
}

// GetSnameOk returns a tuple with the Sname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetSnameOk() (*string, bool) {
	if o == nil || isNil(o.Sname) {
    return nil, false
	}
	return o.Sname, true
}

// HasSname returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasSname() bool {
	if o != nil && !isNil(o.Sname) {
		return true
	}

	return false
}

// SetSname gets a reference to the given string and assigns it to the Sname field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetSname(v string) {
	o.Sname = &v
}

// GetMagicCookie returns the MagicCookie field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetMagicCookie() string {
	if o == nil || isNil(o.MagicCookie) {
		var ret string
		return ret
	}
	return *o.MagicCookie
}

// GetMagicCookieOk returns a tuple with the MagicCookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetMagicCookieOk() (*string, bool) {
	if o == nil || isNil(o.MagicCookie) {
    return nil, false
	}
	return o.MagicCookie, true
}

// HasMagicCookie returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasMagicCookie() bool {
	if o != nil && !isNil(o.MagicCookie) {
		return true
	}

	return false
}

// SetMagicCookie gets a reference to the given string and assigns it to the MagicCookie field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetMagicCookie(v string) {
	o.MagicCookie = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetOptions() []NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions {
	if o == nil || isNil(o.Options) {
		var ret []NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetOptionsOk() ([]NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions, bool) {
	if o == nil || isNil(o.Options) {
    return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasOptions() bool {
	if o != nil && !isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions and assigns it to the Options field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetOptions(v []NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) {
	o.Options = v
}

func (o NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Op) {
		toSerialize["op"] = o.Op
	}
	if !isNil(o.Htype) {
		toSerialize["htype"] = o.Htype
	}
	if !isNil(o.Hlen) {
		toSerialize["hlen"] = o.Hlen
	}
	if !isNil(o.Hops) {
		toSerialize["hops"] = o.Hops
	}
	if !isNil(o.Xid) {
		toSerialize["xid"] = o.Xid
	}
	if !isNil(o.Secs) {
		toSerialize["secs"] = o.Secs
	}
	if !isNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !isNil(o.Ciaddr) {
		toSerialize["ciaddr"] = o.Ciaddr
	}
	if !isNil(o.Yiaddr) {
		toSerialize["yiaddr"] = o.Yiaddr
	}
	if !isNil(o.Siaddr) {
		toSerialize["siaddr"] = o.Siaddr
	}
	if !isNil(o.Giaddr) {
		toSerialize["giaddr"] = o.Giaddr
	}
	if !isNil(o.Chaddr) {
		toSerialize["chaddr"] = o.Chaddr
	}
	if !isNil(o.Sname) {
		toSerialize["sname"] = o.Sname
	}
	if !isNil(o.MagicCookie) {
		toSerialize["magicCookie"] = o.MagicCookie
	}
	if !isNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields struct {
	value *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) Get() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) Set(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields {
	return &NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


