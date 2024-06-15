/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject156 struct for InlineObject156
type InlineObject156 struct {
	// Name or description for layer 3 static route
	Name *string `json:"name,omitempty"`
	// The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24)
	Subnet *string `json:"subnet,omitempty"`
	// IP address of the next hop device to which the device sends its traffic for the subnet
	NextHopIp *string `json:"nextHopIp,omitempty"`
	// Option to advertise static route via OSPF
	AdvertiseViaOspfEnabled *bool `json:"advertiseViaOspfEnabled,omitempty"`
	// Option to prefer static route over OSPF routes
	PreferOverOspfRoutesEnabled *bool `json:"preferOverOspfRoutesEnabled,omitempty"`
}

// NewInlineObject156 instantiates a new InlineObject156 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject156() *InlineObject156 {
	this := InlineObject156{}
	return &this
}

// NewInlineObject156WithDefaults instantiates a new InlineObject156 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject156WithDefaults() *InlineObject156 {
	this := InlineObject156{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject156) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject156) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject156) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject156) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineObject156) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject156) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineObject156) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineObject156) SetSubnet(v string) {
	o.Subnet = &v
}

// GetNextHopIp returns the NextHopIp field value if set, zero value otherwise.
func (o *InlineObject156) GetNextHopIp() string {
	if o == nil || isNil(o.NextHopIp) {
		var ret string
		return ret
	}
	return *o.NextHopIp
}

// GetNextHopIpOk returns a tuple with the NextHopIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject156) GetNextHopIpOk() (*string, bool) {
	if o == nil || isNil(o.NextHopIp) {
    return nil, false
	}
	return o.NextHopIp, true
}

// HasNextHopIp returns a boolean if a field has been set.
func (o *InlineObject156) HasNextHopIp() bool {
	if o != nil && !isNil(o.NextHopIp) {
		return true
	}

	return false
}

// SetNextHopIp gets a reference to the given string and assigns it to the NextHopIp field.
func (o *InlineObject156) SetNextHopIp(v string) {
	o.NextHopIp = &v
}

// GetAdvertiseViaOspfEnabled returns the AdvertiseViaOspfEnabled field value if set, zero value otherwise.
func (o *InlineObject156) GetAdvertiseViaOspfEnabled() bool {
	if o == nil || isNil(o.AdvertiseViaOspfEnabled) {
		var ret bool
		return ret
	}
	return *o.AdvertiseViaOspfEnabled
}

// GetAdvertiseViaOspfEnabledOk returns a tuple with the AdvertiseViaOspfEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject156) GetAdvertiseViaOspfEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.AdvertiseViaOspfEnabled) {
    return nil, false
	}
	return o.AdvertiseViaOspfEnabled, true
}

// HasAdvertiseViaOspfEnabled returns a boolean if a field has been set.
func (o *InlineObject156) HasAdvertiseViaOspfEnabled() bool {
	if o != nil && !isNil(o.AdvertiseViaOspfEnabled) {
		return true
	}

	return false
}

// SetAdvertiseViaOspfEnabled gets a reference to the given bool and assigns it to the AdvertiseViaOspfEnabled field.
func (o *InlineObject156) SetAdvertiseViaOspfEnabled(v bool) {
	o.AdvertiseViaOspfEnabled = &v
}

// GetPreferOverOspfRoutesEnabled returns the PreferOverOspfRoutesEnabled field value if set, zero value otherwise.
func (o *InlineObject156) GetPreferOverOspfRoutesEnabled() bool {
	if o == nil || isNil(o.PreferOverOspfRoutesEnabled) {
		var ret bool
		return ret
	}
	return *o.PreferOverOspfRoutesEnabled
}

// GetPreferOverOspfRoutesEnabledOk returns a tuple with the PreferOverOspfRoutesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject156) GetPreferOverOspfRoutesEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.PreferOverOspfRoutesEnabled) {
    return nil, false
	}
	return o.PreferOverOspfRoutesEnabled, true
}

// HasPreferOverOspfRoutesEnabled returns a boolean if a field has been set.
func (o *InlineObject156) HasPreferOverOspfRoutesEnabled() bool {
	if o != nil && !isNil(o.PreferOverOspfRoutesEnabled) {
		return true
	}

	return false
}

// SetPreferOverOspfRoutesEnabled gets a reference to the given bool and assigns it to the PreferOverOspfRoutesEnabled field.
func (o *InlineObject156) SetPreferOverOspfRoutesEnabled(v bool) {
	o.PreferOverOspfRoutesEnabled = &v
}

func (o InlineObject156) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !isNil(o.NextHopIp) {
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

type NullableInlineObject156 struct {
	value *InlineObject156
	isSet bool
}

func (v NullableInlineObject156) Get() *InlineObject156 {
	return v.value
}

func (v *NullableInlineObject156) Set(val *InlineObject156) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject156) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject156) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject156(val *InlineObject156) *NullableInlineObject156 {
	return &NullableInlineObject156{value: val, isSet: true}
}

func (v NullableInlineObject156) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject156) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


