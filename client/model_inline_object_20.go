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

// InlineObject20 struct for InlineObject20
type InlineObject20 struct {
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

// NewInlineObject20 instantiates a new InlineObject20 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject20() *InlineObject20 {
	this := InlineObject20{}
	return &this
}

// NewInlineObject20WithDefaults instantiates a new InlineObject20 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject20WithDefaults() *InlineObject20 {
	this := InlineObject20{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject20) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject20) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject20) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject20) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineObject20) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject20) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineObject20) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineObject20) SetSubnet(v string) {
	o.Subnet = &v
}

// GetNextHopIp returns the NextHopIp field value if set, zero value otherwise.
func (o *InlineObject20) GetNextHopIp() string {
	if o == nil || o.NextHopIp == nil {
		var ret string
		return ret
	}
	return *o.NextHopIp
}

// GetNextHopIpOk returns a tuple with the NextHopIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject20) GetNextHopIpOk() (*string, bool) {
	if o == nil || o.NextHopIp == nil {
		return nil, false
	}
	return o.NextHopIp, true
}

// HasNextHopIp returns a boolean if a field has been set.
func (o *InlineObject20) HasNextHopIp() bool {
	if o != nil && o.NextHopIp != nil {
		return true
	}

	return false
}

// SetNextHopIp gets a reference to the given string and assigns it to the NextHopIp field.
func (o *InlineObject20) SetNextHopIp(v string) {
	o.NextHopIp = &v
}

// GetAdvertiseViaOspfEnabled returns the AdvertiseViaOspfEnabled field value if set, zero value otherwise.
func (o *InlineObject20) GetAdvertiseViaOspfEnabled() bool {
	if o == nil || o.AdvertiseViaOspfEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AdvertiseViaOspfEnabled
}

// GetAdvertiseViaOspfEnabledOk returns a tuple with the AdvertiseViaOspfEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject20) GetAdvertiseViaOspfEnabledOk() (*bool, bool) {
	if o == nil || o.AdvertiseViaOspfEnabled == nil {
		return nil, false
	}
	return o.AdvertiseViaOspfEnabled, true
}

// HasAdvertiseViaOspfEnabled returns a boolean if a field has been set.
func (o *InlineObject20) HasAdvertiseViaOspfEnabled() bool {
	if o != nil && o.AdvertiseViaOspfEnabled != nil {
		return true
	}

	return false
}

// SetAdvertiseViaOspfEnabled gets a reference to the given bool and assigns it to the AdvertiseViaOspfEnabled field.
func (o *InlineObject20) SetAdvertiseViaOspfEnabled(v bool) {
	o.AdvertiseViaOspfEnabled = &v
}

// GetPreferOverOspfRoutesEnabled returns the PreferOverOspfRoutesEnabled field value if set, zero value otherwise.
func (o *InlineObject20) GetPreferOverOspfRoutesEnabled() bool {
	if o == nil || o.PreferOverOspfRoutesEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PreferOverOspfRoutesEnabled
}

// GetPreferOverOspfRoutesEnabledOk returns a tuple with the PreferOverOspfRoutesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject20) GetPreferOverOspfRoutesEnabledOk() (*bool, bool) {
	if o == nil || o.PreferOverOspfRoutesEnabled == nil {
		return nil, false
	}
	return o.PreferOverOspfRoutesEnabled, true
}

// HasPreferOverOspfRoutesEnabled returns a boolean if a field has been set.
func (o *InlineObject20) HasPreferOverOspfRoutesEnabled() bool {
	if o != nil && o.PreferOverOspfRoutesEnabled != nil {
		return true
	}

	return false
}

// SetPreferOverOspfRoutesEnabled gets a reference to the given bool and assigns it to the PreferOverOspfRoutesEnabled field.
func (o *InlineObject20) SetPreferOverOspfRoutesEnabled(v bool) {
	o.PreferOverOspfRoutesEnabled = &v
}

func (o InlineObject20) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Subnet != nil {
		toSerialize["subnet"] = o.Subnet
	}
	if o.NextHopIp != nil {
		toSerialize["nextHopIp"] = o.NextHopIp
	}
	if o.AdvertiseViaOspfEnabled != nil {
		toSerialize["advertiseViaOspfEnabled"] = o.AdvertiseViaOspfEnabled
	}
	if o.PreferOverOspfRoutesEnabled != nil {
		toSerialize["preferOverOspfRoutesEnabled"] = o.PreferOverOspfRoutesEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject20 struct {
	value *InlineObject20
	isSet bool
}

func (v NullableInlineObject20) Get() *InlineObject20 {
	return v.value
}

func (v *NullableInlineObject20) Set(val *InlineObject20) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject20) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject20) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject20(val *InlineObject20) *NullableInlineObject20 {
	return &NullableInlineObject20{value: val, isSet: true}
}

func (v NullableInlineObject20) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject20) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


