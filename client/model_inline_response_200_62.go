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

// InlineResponse20062 struct for InlineResponse20062
type InlineResponse20062 struct {
	// Route ID
	Id *string `json:"id,omitempty"`
	// IP protocol version
	IpVersion *int32 `json:"ipVersion,omitempty"`
	// Network ID
	NetworkId *string `json:"networkId,omitempty"`
	// Whether the route is enabled or not
	Enabled *bool `json:"enabled,omitempty"`
	// Name of the route
	Name *string `json:"name,omitempty"`
	// Subnet of the route
	Subnet *string `json:"subnet,omitempty"`
	// Gateway IP address (next hop)
	GatewayIp *string `json:"gatewayIp,omitempty"`
	// Fixed DHCP IP assignments on the route
	FixedIpAssignments *map[string]NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments `json:"fixedIpAssignments,omitempty"`
	// DHCP reserved IP ranges
	ReservedIpRanges []NetworksNetworkIdApplianceStaticRoutesReservedIpRanges `json:"reservedIpRanges,omitempty"`
	// Gateway VLAN ID
	GatewayVlanId *int32 `json:"gatewayVlanId,omitempty"`
}

// NewInlineResponse20062 instantiates a new InlineResponse20062 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20062() *InlineResponse20062 {
	this := InlineResponse20062{}
	return &this
}

// NewInlineResponse20062WithDefaults instantiates a new InlineResponse20062 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20062WithDefaults() *InlineResponse20062 {
	this := InlineResponse20062{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20062) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20062) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20062) SetId(v string) {
	o.Id = &v
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise.
func (o *InlineResponse20062) GetIpVersion() int32 {
	if o == nil || isNil(o.IpVersion) {
		var ret int32
		return ret
	}
	return *o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetIpVersionOk() (*int32, bool) {
	if o == nil || isNil(o.IpVersion) {
    return nil, false
	}
	return o.IpVersion, true
}

// HasIpVersion returns a boolean if a field has been set.
func (o *InlineResponse20062) HasIpVersion() bool {
	if o != nil && !isNil(o.IpVersion) {
		return true
	}

	return false
}

// SetIpVersion gets a reference to the given int32 and assigns it to the IpVersion field.
func (o *InlineResponse20062) SetIpVersion(v int32) {
	o.IpVersion = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse20062) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse20062) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse20062) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20062) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20062) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20062) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20062) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20062) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20062) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineResponse20062) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineResponse20062) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineResponse20062) SetSubnet(v string) {
	o.Subnet = &v
}

// GetGatewayIp returns the GatewayIp field value if set, zero value otherwise.
func (o *InlineResponse20062) GetGatewayIp() string {
	if o == nil || isNil(o.GatewayIp) {
		var ret string
		return ret
	}
	return *o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetGatewayIpOk() (*string, bool) {
	if o == nil || isNil(o.GatewayIp) {
    return nil, false
	}
	return o.GatewayIp, true
}

// HasGatewayIp returns a boolean if a field has been set.
func (o *InlineResponse20062) HasGatewayIp() bool {
	if o != nil && !isNil(o.GatewayIp) {
		return true
	}

	return false
}

// SetGatewayIp gets a reference to the given string and assigns it to the GatewayIp field.
func (o *InlineResponse20062) SetGatewayIp(v string) {
	o.GatewayIp = &v
}

// GetFixedIpAssignments returns the FixedIpAssignments field value if set, zero value otherwise.
func (o *InlineResponse20062) GetFixedIpAssignments() map[string]NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments {
	if o == nil || isNil(o.FixedIpAssignments) {
		var ret map[string]NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments
		return ret
	}
	return *o.FixedIpAssignments
}

// GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetFixedIpAssignmentsOk() (*map[string]NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments, bool) {
	if o == nil || isNil(o.FixedIpAssignments) {
    return nil, false
	}
	return o.FixedIpAssignments, true
}

// HasFixedIpAssignments returns a boolean if a field has been set.
func (o *InlineResponse20062) HasFixedIpAssignments() bool {
	if o != nil && !isNil(o.FixedIpAssignments) {
		return true
	}

	return false
}

// SetFixedIpAssignments gets a reference to the given map[string]NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments and assigns it to the FixedIpAssignments field.
func (o *InlineResponse20062) SetFixedIpAssignments(v map[string]NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) {
	o.FixedIpAssignments = &v
}

// GetReservedIpRanges returns the ReservedIpRanges field value if set, zero value otherwise.
func (o *InlineResponse20062) GetReservedIpRanges() []NetworksNetworkIdApplianceStaticRoutesReservedIpRanges {
	if o == nil || isNil(o.ReservedIpRanges) {
		var ret []NetworksNetworkIdApplianceStaticRoutesReservedIpRanges
		return ret
	}
	return o.ReservedIpRanges
}

// GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetReservedIpRangesOk() ([]NetworksNetworkIdApplianceStaticRoutesReservedIpRanges, bool) {
	if o == nil || isNil(o.ReservedIpRanges) {
    return nil, false
	}
	return o.ReservedIpRanges, true
}

// HasReservedIpRanges returns a boolean if a field has been set.
func (o *InlineResponse20062) HasReservedIpRanges() bool {
	if o != nil && !isNil(o.ReservedIpRanges) {
		return true
	}

	return false
}

// SetReservedIpRanges gets a reference to the given []NetworksNetworkIdApplianceStaticRoutesReservedIpRanges and assigns it to the ReservedIpRanges field.
func (o *InlineResponse20062) SetReservedIpRanges(v []NetworksNetworkIdApplianceStaticRoutesReservedIpRanges) {
	o.ReservedIpRanges = v
}

// GetGatewayVlanId returns the GatewayVlanId field value if set, zero value otherwise.
func (o *InlineResponse20062) GetGatewayVlanId() int32 {
	if o == nil || isNil(o.GatewayVlanId) {
		var ret int32
		return ret
	}
	return *o.GatewayVlanId
}

// GetGatewayVlanIdOk returns a tuple with the GatewayVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetGatewayVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.GatewayVlanId) {
    return nil, false
	}
	return o.GatewayVlanId, true
}

// HasGatewayVlanId returns a boolean if a field has been set.
func (o *InlineResponse20062) HasGatewayVlanId() bool {
	if o != nil && !isNil(o.GatewayVlanId) {
		return true
	}

	return false
}

// SetGatewayVlanId gets a reference to the given int32 and assigns it to the GatewayVlanId field.
func (o *InlineResponse20062) SetGatewayVlanId(v int32) {
	o.GatewayVlanId = &v
}

func (o InlineResponse20062) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IpVersion) {
		toSerialize["ipVersion"] = o.IpVersion
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !isNil(o.GatewayIp) {
		toSerialize["gatewayIp"] = o.GatewayIp
	}
	if !isNil(o.FixedIpAssignments) {
		toSerialize["fixedIpAssignments"] = o.FixedIpAssignments
	}
	if !isNil(o.ReservedIpRanges) {
		toSerialize["reservedIpRanges"] = o.ReservedIpRanges
	}
	if !isNil(o.GatewayVlanId) {
		toSerialize["gatewayVlanId"] = o.GatewayVlanId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20062 struct {
	value *InlineResponse20062
	isSet bool
}

func (v NullableInlineResponse20062) Get() *InlineResponse20062 {
	return v.value
}

func (v *NullableInlineResponse20062) Set(val *InlineResponse20062) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20062) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20062) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20062(val *InlineResponse20062) *NullableInlineResponse20062 {
	return &NullableInlineResponse20062{value: val, isSet: true}
}

func (v NullableInlineResponse20062) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20062) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


