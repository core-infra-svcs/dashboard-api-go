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

// InlineObject154 struct for InlineObject154
type InlineObject154 struct {
	// Name or description for layer 3 static route
	Name *string `json:"name,omitempty"`
	// The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24)
	Subnet string `json:"subnet"`
	// IP address of the next hop device to which the device sends its traffic for the subnet
	NextHopIp string `json:"nextHopIp"`
	// Option to advertise static route via OSPF
	AdvertiseViaOspfEnabled *bool `json:"advertiseViaOspfEnabled,omitempty"`
	// Option to prefer static route over OSPF routes
	PreferOverOspfRoutesEnabled *bool `json:"preferOverOspfRoutesEnabled,omitempty"`
}

// NewInlineObject154 instantiates a new InlineObject154 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject154(subnet string, nextHopIp string) *InlineObject154 {
	this := InlineObject154{}
	this.Subnet = subnet
	this.NextHopIp = nextHopIp
	return &this
}

// NewInlineObject154WithDefaults instantiates a new InlineObject154 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject154WithDefaults() *InlineObject154 {
	this := InlineObject154{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject154) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject154) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject154) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject154) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value
func (o *InlineObject154) GetSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value
// and a boolean to check if the value has been set.
func (o *InlineObject154) GetSubnetOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subnet, true
}

// SetSubnet sets field value
func (o *InlineObject154) SetSubnet(v string) {
	o.Subnet = v
}

// GetNextHopIp returns the NextHopIp field value
func (o *InlineObject154) GetNextHopIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextHopIp
}

// GetNextHopIpOk returns a tuple with the NextHopIp field value
// and a boolean to check if the value has been set.
func (o *InlineObject154) GetNextHopIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NextHopIp, true
}

// SetNextHopIp sets field value
func (o *InlineObject154) SetNextHopIp(v string) {
	o.NextHopIp = v
}

// GetAdvertiseViaOspfEnabled returns the AdvertiseViaOspfEnabled field value if set, zero value otherwise.
func (o *InlineObject154) GetAdvertiseViaOspfEnabled() bool {
	if o == nil || isNil(o.AdvertiseViaOspfEnabled) {
		var ret bool
		return ret
	}
	return *o.AdvertiseViaOspfEnabled
}

// GetAdvertiseViaOspfEnabledOk returns a tuple with the AdvertiseViaOspfEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject154) GetAdvertiseViaOspfEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.AdvertiseViaOspfEnabled) {
    return nil, false
	}
	return o.AdvertiseViaOspfEnabled, true
}

// HasAdvertiseViaOspfEnabled returns a boolean if a field has been set.
func (o *InlineObject154) HasAdvertiseViaOspfEnabled() bool {
	if o != nil && !isNil(o.AdvertiseViaOspfEnabled) {
		return true
	}

	return false
}

// SetAdvertiseViaOspfEnabled gets a reference to the given bool and assigns it to the AdvertiseViaOspfEnabled field.
func (o *InlineObject154) SetAdvertiseViaOspfEnabled(v bool) {
	o.AdvertiseViaOspfEnabled = &v
}

// GetPreferOverOspfRoutesEnabled returns the PreferOverOspfRoutesEnabled field value if set, zero value otherwise.
func (o *InlineObject154) GetPreferOverOspfRoutesEnabled() bool {
	if o == nil || isNil(o.PreferOverOspfRoutesEnabled) {
		var ret bool
		return ret
	}
	return *o.PreferOverOspfRoutesEnabled
}

// GetPreferOverOspfRoutesEnabledOk returns a tuple with the PreferOverOspfRoutesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject154) GetPreferOverOspfRoutesEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.PreferOverOspfRoutesEnabled) {
    return nil, false
	}
	return o.PreferOverOspfRoutesEnabled, true
}

// HasPreferOverOspfRoutesEnabled returns a boolean if a field has been set.
func (o *InlineObject154) HasPreferOverOspfRoutesEnabled() bool {
	if o != nil && !isNil(o.PreferOverOspfRoutesEnabled) {
		return true
	}

	return false
}

// SetPreferOverOspfRoutesEnabled gets a reference to the given bool and assigns it to the PreferOverOspfRoutesEnabled field.
func (o *InlineObject154) SetPreferOverOspfRoutesEnabled(v bool) {
	o.PreferOverOspfRoutesEnabled = &v
}

func (o InlineObject154) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableInlineObject154 struct {
	value *InlineObject154
	isSet bool
}

func (v NullableInlineObject154) Get() *InlineObject154 {
	return v.value
}

func (v *NullableInlineObject154) Set(val *InlineObject154) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject154) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject154) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject154(val *InlineObject154) *NullableInlineObject154 {
	return &NullableInlineObject154{value: val, isSet: true}
}

func (v NullableInlineObject154) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject154) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


