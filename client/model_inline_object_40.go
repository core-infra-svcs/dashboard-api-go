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

// InlineObject40 struct for InlineObject40
type InlineObject40 struct {
	// The name of the static route
	Name *string `json:"name,omitempty"`
	// The subnet of the static route
	Subnet *string `json:"subnet,omitempty"`
	// The gateway IP (next hop) of the static route
	GatewayIp *string `json:"gatewayIp,omitempty"`
	// The gateway IP (next hop) VLAN ID of the static route
	GatewayVlanId *string `json:"gatewayVlanId,omitempty"`
	// The enabled state of the static route
	Enabled *bool `json:"enabled,omitempty"`
	// The DHCP fixed IP assignments on the static route. This should be an object that contains mappings from MAC addresses to objects that themselves each contain \"ip\" and \"name\" string fields. See the sample request/response for more details.
	FixedIpAssignments *map[string]interface{} `json:"fixedIpAssignments,omitempty"`
	// The DHCP reserved IP ranges on the static route
	ReservedIpRanges *[]NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges `json:"reservedIpRanges,omitempty"`
}

// NewInlineObject40 instantiates a new InlineObject40 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject40() *InlineObject40 {
	this := InlineObject40{}
	return &this
}

// NewInlineObject40WithDefaults instantiates a new InlineObject40 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject40WithDefaults() *InlineObject40 {
	this := InlineObject40{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject40) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject40) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject40) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject40) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineObject40) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject40) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineObject40) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineObject40) SetSubnet(v string) {
	o.Subnet = &v
}

// GetGatewayIp returns the GatewayIp field value if set, zero value otherwise.
func (o *InlineObject40) GetGatewayIp() string {
	if o == nil || o.GatewayIp == nil {
		var ret string
		return ret
	}
	return *o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject40) GetGatewayIpOk() (*string, bool) {
	if o == nil || o.GatewayIp == nil {
		return nil, false
	}
	return o.GatewayIp, true
}

// HasGatewayIp returns a boolean if a field has been set.
func (o *InlineObject40) HasGatewayIp() bool {
	if o != nil && o.GatewayIp != nil {
		return true
	}

	return false
}

// SetGatewayIp gets a reference to the given string and assigns it to the GatewayIp field.
func (o *InlineObject40) SetGatewayIp(v string) {
	o.GatewayIp = &v
}

// GetGatewayVlanId returns the GatewayVlanId field value if set, zero value otherwise.
func (o *InlineObject40) GetGatewayVlanId() string {
	if o == nil || o.GatewayVlanId == nil {
		var ret string
		return ret
	}
	return *o.GatewayVlanId
}

// GetGatewayVlanIdOk returns a tuple with the GatewayVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject40) GetGatewayVlanIdOk() (*string, bool) {
	if o == nil || o.GatewayVlanId == nil {
		return nil, false
	}
	return o.GatewayVlanId, true
}

// HasGatewayVlanId returns a boolean if a field has been set.
func (o *InlineObject40) HasGatewayVlanId() bool {
	if o != nil && o.GatewayVlanId != nil {
		return true
	}

	return false
}

// SetGatewayVlanId gets a reference to the given string and assigns it to the GatewayVlanId field.
func (o *InlineObject40) SetGatewayVlanId(v string) {
	o.GatewayVlanId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject40) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject40) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject40) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject40) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFixedIpAssignments returns the FixedIpAssignments field value if set, zero value otherwise.
func (o *InlineObject40) GetFixedIpAssignments() map[string]interface{} {
	if o == nil || o.FixedIpAssignments == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.FixedIpAssignments
}

// GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject40) GetFixedIpAssignmentsOk() (*map[string]interface{}, bool) {
	if o == nil || o.FixedIpAssignments == nil {
		return nil, false
	}
	return o.FixedIpAssignments, true
}

// HasFixedIpAssignments returns a boolean if a field has been set.
func (o *InlineObject40) HasFixedIpAssignments() bool {
	if o != nil && o.FixedIpAssignments != nil {
		return true
	}

	return false
}

// SetFixedIpAssignments gets a reference to the given map[string]interface{} and assigns it to the FixedIpAssignments field.
func (o *InlineObject40) SetFixedIpAssignments(v map[string]interface{}) {
	o.FixedIpAssignments = &v
}

// GetReservedIpRanges returns the ReservedIpRanges field value if set, zero value otherwise.
func (o *InlineObject40) GetReservedIpRanges() []NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges {
	if o == nil || o.ReservedIpRanges == nil {
		var ret []NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges
		return ret
	}
	return *o.ReservedIpRanges
}

// GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject40) GetReservedIpRangesOk() (*[]NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges, bool) {
	if o == nil || o.ReservedIpRanges == nil {
		return nil, false
	}
	return o.ReservedIpRanges, true
}

// HasReservedIpRanges returns a boolean if a field has been set.
func (o *InlineObject40) HasReservedIpRanges() bool {
	if o != nil && o.ReservedIpRanges != nil {
		return true
	}

	return false
}

// SetReservedIpRanges gets a reference to the given []NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges and assigns it to the ReservedIpRanges field.
func (o *InlineObject40) SetReservedIpRanges(v []NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) {
	o.ReservedIpRanges = &v
}

func (o InlineObject40) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Subnet != nil {
		toSerialize["subnet"] = o.Subnet
	}
	if o.GatewayIp != nil {
		toSerialize["gatewayIp"] = o.GatewayIp
	}
	if o.GatewayVlanId != nil {
		toSerialize["gatewayVlanId"] = o.GatewayVlanId
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.FixedIpAssignments != nil {
		toSerialize["fixedIpAssignments"] = o.FixedIpAssignments
	}
	if o.ReservedIpRanges != nil {
		toSerialize["reservedIpRanges"] = o.ReservedIpRanges
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject40 struct {
	value *InlineObject40
	isSet bool
}

func (v NullableInlineObject40) Get() *InlineObject40 {
	return v.value
}

func (v *NullableInlineObject40) Set(val *InlineObject40) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject40) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject40) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject40(val *InlineObject40) *NullableInlineObject40 {
	return &NullableInlineObject40{value: val, isSet: true}
}

func (v NullableInlineObject40) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject40) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


