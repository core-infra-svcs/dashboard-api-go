/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20036 struct for InlineResponse20036
type InlineResponse20036 struct {
	// The identifier of a layer 3 static route
	StaticRouteId *string `json:"staticRouteId,omitempty"`
	// The name or description of the layer 3 static route
	Name *string `json:"name,omitempty"`
	// The IP address of the subnetwork specified in CIDR notation (ex. 1.2.3.0/24)
	Subnet string `json:"subnet"`
	//  The IP address of the router to which traffic for this destination network should be sent
	NextHopIp string `json:"nextHopIp"`
	// Option to advertise static routes via OSPF
	AdvertiseViaOspfEnabled *bool `json:"advertiseViaOspfEnabled,omitempty"`
	// Option to prefer static routes over OSPF routes
	PreferOverOspfRoutesEnabled *bool `json:"preferOverOspfRoutesEnabled,omitempty"`
}

// NewInlineResponse20036 instantiates a new InlineResponse20036 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20036(subnet string, nextHopIp string) *InlineResponse20036 {
	this := InlineResponse20036{}
	this.Subnet = subnet
	this.NextHopIp = nextHopIp
	return &this
}

// NewInlineResponse20036WithDefaults instantiates a new InlineResponse20036 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20036WithDefaults() *InlineResponse20036 {
	this := InlineResponse20036{}
	return &this
}

// GetStaticRouteId returns the StaticRouteId field value if set, zero value otherwise.
func (o *InlineResponse20036) GetStaticRouteId() string {
	if o == nil || isNil(o.StaticRouteId) {
		var ret string
		return ret
	}
	return *o.StaticRouteId
}

// GetStaticRouteIdOk returns a tuple with the StaticRouteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetStaticRouteIdOk() (*string, bool) {
	if o == nil || isNil(o.StaticRouteId) {
    return nil, false
	}
	return o.StaticRouteId, true
}

// HasStaticRouteId returns a boolean if a field has been set.
func (o *InlineResponse20036) HasStaticRouteId() bool {
	if o != nil && !isNil(o.StaticRouteId) {
		return true
	}

	return false
}

// SetStaticRouteId gets a reference to the given string and assigns it to the StaticRouteId field.
func (o *InlineResponse20036) SetStaticRouteId(v string) {
	o.StaticRouteId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20036) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20036) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20036) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value
func (o *InlineResponse20036) GetSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetSubnetOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subnet, true
}

// SetSubnet sets field value
func (o *InlineResponse20036) SetSubnet(v string) {
	o.Subnet = v
}

// GetNextHopIp returns the NextHopIp field value
func (o *InlineResponse20036) GetNextHopIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextHopIp
}

// GetNextHopIpOk returns a tuple with the NextHopIp field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetNextHopIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NextHopIp, true
}

// SetNextHopIp sets field value
func (o *InlineResponse20036) SetNextHopIp(v string) {
	o.NextHopIp = v
}

// GetAdvertiseViaOspfEnabled returns the AdvertiseViaOspfEnabled field value if set, zero value otherwise.
func (o *InlineResponse20036) GetAdvertiseViaOspfEnabled() bool {
	if o == nil || isNil(o.AdvertiseViaOspfEnabled) {
		var ret bool
		return ret
	}
	return *o.AdvertiseViaOspfEnabled
}

// GetAdvertiseViaOspfEnabledOk returns a tuple with the AdvertiseViaOspfEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetAdvertiseViaOspfEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.AdvertiseViaOspfEnabled) {
    return nil, false
	}
	return o.AdvertiseViaOspfEnabled, true
}

// HasAdvertiseViaOspfEnabled returns a boolean if a field has been set.
func (o *InlineResponse20036) HasAdvertiseViaOspfEnabled() bool {
	if o != nil && !isNil(o.AdvertiseViaOspfEnabled) {
		return true
	}

	return false
}

// SetAdvertiseViaOspfEnabled gets a reference to the given bool and assigns it to the AdvertiseViaOspfEnabled field.
func (o *InlineResponse20036) SetAdvertiseViaOspfEnabled(v bool) {
	o.AdvertiseViaOspfEnabled = &v
}

// GetPreferOverOspfRoutesEnabled returns the PreferOverOspfRoutesEnabled field value if set, zero value otherwise.
func (o *InlineResponse20036) GetPreferOverOspfRoutesEnabled() bool {
	if o == nil || isNil(o.PreferOverOspfRoutesEnabled) {
		var ret bool
		return ret
	}
	return *o.PreferOverOspfRoutesEnabled
}

// GetPreferOverOspfRoutesEnabledOk returns a tuple with the PreferOverOspfRoutesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetPreferOverOspfRoutesEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.PreferOverOspfRoutesEnabled) {
    return nil, false
	}
	return o.PreferOverOspfRoutesEnabled, true
}

// HasPreferOverOspfRoutesEnabled returns a boolean if a field has been set.
func (o *InlineResponse20036) HasPreferOverOspfRoutesEnabled() bool {
	if o != nil && !isNil(o.PreferOverOspfRoutesEnabled) {
		return true
	}

	return false
}

// SetPreferOverOspfRoutesEnabled gets a reference to the given bool and assigns it to the PreferOverOspfRoutesEnabled field.
func (o *InlineResponse20036) SetPreferOverOspfRoutesEnabled(v bool) {
	o.PreferOverOspfRoutesEnabled = &v
}

func (o InlineResponse20036) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StaticRouteId) {
		toSerialize["staticRouteId"] = o.StaticRouteId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["subnet"] = o.Subnet
	}
	if true {
		toSerialize["nextHopIp"] = o.NextHopIp
	}
	if !isNil(o.AdvertiseViaOspfEnabled) {
		toSerialize["advertiseViaOspfEnabled"] = o.AdvertiseViaOspfEnabled
	}
	if !isNil(o.PreferOverOspfRoutesEnabled) {
		toSerialize["preferOverOspfRoutesEnabled"] = o.PreferOverOspfRoutesEnabled
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


