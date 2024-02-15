/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject143 struct for InlineObject143
type InlineObject143 struct {
	// Boolean value to enable or disable OSPF routing. OSPF routing is disabled by default.
	Enabled *bool `json:"enabled,omitempty"`
	// Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.
	HelloTimerInSeconds *int32 `json:"helloTimerInSeconds,omitempty"`
	// Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535
	DeadTimerInSeconds *int32 `json:"deadTimerInSeconds,omitempty"`
	// OSPF areas
	Areas []NetworksNetworkIdSwitchRoutingOspfAreas `json:"areas,omitempty"`
	V3 *NetworksNetworkIdSwitchRoutingOspfV3 `json:"v3,omitempty"`
	// Boolean value to enable or disable MD5 authentication. MD5 authentication is disabled by default.
	Md5AuthenticationEnabled *bool `json:"md5AuthenticationEnabled,omitempty"`
	Md5AuthenticationKey *NetworksNetworkIdSwitchRoutingOspfMd5AuthenticationKey `json:"md5AuthenticationKey,omitempty"`
}

// NewInlineObject143 instantiates a new InlineObject143 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject143() *InlineObject143 {
	this := InlineObject143{}
	return &this
}

// NewInlineObject143WithDefaults instantiates a new InlineObject143 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject143WithDefaults() *InlineObject143 {
	this := InlineObject143{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject143) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject143) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject143) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHelloTimerInSeconds returns the HelloTimerInSeconds field value if set, zero value otherwise.
func (o *InlineObject143) GetHelloTimerInSeconds() int32 {
	if o == nil || isNil(o.HelloTimerInSeconds) {
		var ret int32
		return ret
	}
	return *o.HelloTimerInSeconds
}

// GetHelloTimerInSecondsOk returns a tuple with the HelloTimerInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetHelloTimerInSecondsOk() (*int32, bool) {
	if o == nil || isNil(o.HelloTimerInSeconds) {
    return nil, false
	}
	return o.HelloTimerInSeconds, true
}

// HasHelloTimerInSeconds returns a boolean if a field has been set.
func (o *InlineObject143) HasHelloTimerInSeconds() bool {
	if o != nil && !isNil(o.HelloTimerInSeconds) {
		return true
	}

	return false
}

// SetHelloTimerInSeconds gets a reference to the given int32 and assigns it to the HelloTimerInSeconds field.
func (o *InlineObject143) SetHelloTimerInSeconds(v int32) {
	o.HelloTimerInSeconds = &v
}

// GetDeadTimerInSeconds returns the DeadTimerInSeconds field value if set, zero value otherwise.
func (o *InlineObject143) GetDeadTimerInSeconds() int32 {
	if o == nil || isNil(o.DeadTimerInSeconds) {
		var ret int32
		return ret
	}
	return *o.DeadTimerInSeconds
}

// GetDeadTimerInSecondsOk returns a tuple with the DeadTimerInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetDeadTimerInSecondsOk() (*int32, bool) {
	if o == nil || isNil(o.DeadTimerInSeconds) {
    return nil, false
	}
	return o.DeadTimerInSeconds, true
}

// HasDeadTimerInSeconds returns a boolean if a field has been set.
func (o *InlineObject143) HasDeadTimerInSeconds() bool {
	if o != nil && !isNil(o.DeadTimerInSeconds) {
		return true
	}

	return false
}

// SetDeadTimerInSeconds gets a reference to the given int32 and assigns it to the DeadTimerInSeconds field.
func (o *InlineObject143) SetDeadTimerInSeconds(v int32) {
	o.DeadTimerInSeconds = &v
}

// GetAreas returns the Areas field value if set, zero value otherwise.
func (o *InlineObject143) GetAreas() []NetworksNetworkIdSwitchRoutingOspfAreas {
	if o == nil || isNil(o.Areas) {
		var ret []NetworksNetworkIdSwitchRoutingOspfAreas
		return ret
	}
	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetAreasOk() ([]NetworksNetworkIdSwitchRoutingOspfAreas, bool) {
	if o == nil || isNil(o.Areas) {
    return nil, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *InlineObject143) HasAreas() bool {
	if o != nil && !isNil(o.Areas) {
		return true
	}

	return false
}

// SetAreas gets a reference to the given []NetworksNetworkIdSwitchRoutingOspfAreas and assigns it to the Areas field.
func (o *InlineObject143) SetAreas(v []NetworksNetworkIdSwitchRoutingOspfAreas) {
	o.Areas = v
}

// GetV3 returns the V3 field value if set, zero value otherwise.
func (o *InlineObject143) GetV3() NetworksNetworkIdSwitchRoutingOspfV3 {
	if o == nil || isNil(o.V3) {
		var ret NetworksNetworkIdSwitchRoutingOspfV3
		return ret
	}
	return *o.V3
}

// GetV3Ok returns a tuple with the V3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetV3Ok() (*NetworksNetworkIdSwitchRoutingOspfV3, bool) {
	if o == nil || isNil(o.V3) {
    return nil, false
	}
	return o.V3, true
}

// HasV3 returns a boolean if a field has been set.
func (o *InlineObject143) HasV3() bool {
	if o != nil && !isNil(o.V3) {
		return true
	}

	return false
}

// SetV3 gets a reference to the given NetworksNetworkIdSwitchRoutingOspfV3 and assigns it to the V3 field.
func (o *InlineObject143) SetV3(v NetworksNetworkIdSwitchRoutingOspfV3) {
	o.V3 = &v
}

// GetMd5AuthenticationEnabled returns the Md5AuthenticationEnabled field value if set, zero value otherwise.
func (o *InlineObject143) GetMd5AuthenticationEnabled() bool {
	if o == nil || isNil(o.Md5AuthenticationEnabled) {
		var ret bool
		return ret
	}
	return *o.Md5AuthenticationEnabled
}

// GetMd5AuthenticationEnabledOk returns a tuple with the Md5AuthenticationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetMd5AuthenticationEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Md5AuthenticationEnabled) {
    return nil, false
	}
	return o.Md5AuthenticationEnabled, true
}

// HasMd5AuthenticationEnabled returns a boolean if a field has been set.
func (o *InlineObject143) HasMd5AuthenticationEnabled() bool {
	if o != nil && !isNil(o.Md5AuthenticationEnabled) {
		return true
	}

	return false
}

// SetMd5AuthenticationEnabled gets a reference to the given bool and assigns it to the Md5AuthenticationEnabled field.
func (o *InlineObject143) SetMd5AuthenticationEnabled(v bool) {
	o.Md5AuthenticationEnabled = &v
}

// GetMd5AuthenticationKey returns the Md5AuthenticationKey field value if set, zero value otherwise.
func (o *InlineObject143) GetMd5AuthenticationKey() NetworksNetworkIdSwitchRoutingOspfMd5AuthenticationKey {
	if o == nil || isNil(o.Md5AuthenticationKey) {
		var ret NetworksNetworkIdSwitchRoutingOspfMd5AuthenticationKey
		return ret
	}
	return *o.Md5AuthenticationKey
}

// GetMd5AuthenticationKeyOk returns a tuple with the Md5AuthenticationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetMd5AuthenticationKeyOk() (*NetworksNetworkIdSwitchRoutingOspfMd5AuthenticationKey, bool) {
	if o == nil || isNil(o.Md5AuthenticationKey) {
    return nil, false
	}
	return o.Md5AuthenticationKey, true
}

// HasMd5AuthenticationKey returns a boolean if a field has been set.
func (o *InlineObject143) HasMd5AuthenticationKey() bool {
	if o != nil && !isNil(o.Md5AuthenticationKey) {
		return true
	}

	return false
}

// SetMd5AuthenticationKey gets a reference to the given NetworksNetworkIdSwitchRoutingOspfMd5AuthenticationKey and assigns it to the Md5AuthenticationKey field.
func (o *InlineObject143) SetMd5AuthenticationKey(v NetworksNetworkIdSwitchRoutingOspfMd5AuthenticationKey) {
	o.Md5AuthenticationKey = &v
}

func (o InlineObject143) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.HelloTimerInSeconds) {
		toSerialize["helloTimerInSeconds"] = o.HelloTimerInSeconds
	}
	if !isNil(o.DeadTimerInSeconds) {
		toSerialize["deadTimerInSeconds"] = o.DeadTimerInSeconds
	}
	if !isNil(o.Areas) {
		toSerialize["areas"] = o.Areas
	}
	if !isNil(o.V3) {
		toSerialize["v3"] = o.V3
	}
	if !isNil(o.Md5AuthenticationEnabled) {
		toSerialize["md5AuthenticationEnabled"] = o.Md5AuthenticationEnabled
	}
	if !isNil(o.Md5AuthenticationKey) {
		toSerialize["md5AuthenticationKey"] = o.Md5AuthenticationKey
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject143 struct {
	value *InlineObject143
	isSet bool
}

func (v NullableInlineObject143) Get() *InlineObject143 {
	return v.value
}

func (v *NullableInlineObject143) Set(val *InlineObject143) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject143) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject143) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject143(val *InlineObject143) *NullableInlineObject143 {
	return &NullableInlineObject143{value: val, isSet: true}
}

func (v NullableInlineObject143) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject143) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


