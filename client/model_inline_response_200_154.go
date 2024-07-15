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

// InlineResponse200154 struct for InlineResponse200154
type InlineResponse200154 struct {
	// Boolean value to enable or disable OSPF routing. OSPF routing is disabled by default.
	Enabled *bool `json:"enabled,omitempty"`
	// Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.
	HelloTimerInSeconds *int32 `json:"helloTimerInSeconds,omitempty"`
	// Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535
	DeadTimerInSeconds *int32 `json:"deadTimerInSeconds,omitempty"`
	// OSPF areas
	Areas []InlineResponse200154Areas `json:"areas,omitempty"`
	V3 *InlineResponse200154V3 `json:"v3,omitempty"`
	// Boolean value to enable or disable MD5 authentication. MD5 authentication is disabled by default.
	Md5AuthenticationEnabled *bool `json:"md5AuthenticationEnabled,omitempty"`
	Md5AuthenticationKey *InlineResponse200154Md5AuthenticationKey `json:"md5AuthenticationKey,omitempty"`
}

// NewInlineResponse200154 instantiates a new InlineResponse200154 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200154() *InlineResponse200154 {
	this := InlineResponse200154{}
	return &this
}

// NewInlineResponse200154WithDefaults instantiates a new InlineResponse200154 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200154WithDefaults() *InlineResponse200154 {
	this := InlineResponse200154{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200154) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200154) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200154) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200154) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHelloTimerInSeconds returns the HelloTimerInSeconds field value if set, zero value otherwise.
func (o *InlineResponse200154) GetHelloTimerInSeconds() int32 {
	if o == nil || isNil(o.HelloTimerInSeconds) {
		var ret int32
		return ret
	}
	return *o.HelloTimerInSeconds
}

// GetHelloTimerInSecondsOk returns a tuple with the HelloTimerInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200154) GetHelloTimerInSecondsOk() (*int32, bool) {
	if o == nil || isNil(o.HelloTimerInSeconds) {
    return nil, false
	}
	return o.HelloTimerInSeconds, true
}

// HasHelloTimerInSeconds returns a boolean if a field has been set.
func (o *InlineResponse200154) HasHelloTimerInSeconds() bool {
	if o != nil && !isNil(o.HelloTimerInSeconds) {
		return true
	}

	return false
}

// SetHelloTimerInSeconds gets a reference to the given int32 and assigns it to the HelloTimerInSeconds field.
func (o *InlineResponse200154) SetHelloTimerInSeconds(v int32) {
	o.HelloTimerInSeconds = &v
}

// GetDeadTimerInSeconds returns the DeadTimerInSeconds field value if set, zero value otherwise.
func (o *InlineResponse200154) GetDeadTimerInSeconds() int32 {
	if o == nil || isNil(o.DeadTimerInSeconds) {
		var ret int32
		return ret
	}
	return *o.DeadTimerInSeconds
}

// GetDeadTimerInSecondsOk returns a tuple with the DeadTimerInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200154) GetDeadTimerInSecondsOk() (*int32, bool) {
	if o == nil || isNil(o.DeadTimerInSeconds) {
    return nil, false
	}
	return o.DeadTimerInSeconds, true
}

// HasDeadTimerInSeconds returns a boolean if a field has been set.
func (o *InlineResponse200154) HasDeadTimerInSeconds() bool {
	if o != nil && !isNil(o.DeadTimerInSeconds) {
		return true
	}

	return false
}

// SetDeadTimerInSeconds gets a reference to the given int32 and assigns it to the DeadTimerInSeconds field.
func (o *InlineResponse200154) SetDeadTimerInSeconds(v int32) {
	o.DeadTimerInSeconds = &v
}

// GetAreas returns the Areas field value if set, zero value otherwise.
func (o *InlineResponse200154) GetAreas() []InlineResponse200154Areas {
	if o == nil || isNil(o.Areas) {
		var ret []InlineResponse200154Areas
		return ret
	}
	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200154) GetAreasOk() ([]InlineResponse200154Areas, bool) {
	if o == nil || isNil(o.Areas) {
    return nil, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *InlineResponse200154) HasAreas() bool {
	if o != nil && !isNil(o.Areas) {
		return true
	}

	return false
}

// SetAreas gets a reference to the given []InlineResponse200154Areas and assigns it to the Areas field.
func (o *InlineResponse200154) SetAreas(v []InlineResponse200154Areas) {
	o.Areas = v
}

// GetV3 returns the V3 field value if set, zero value otherwise.
func (o *InlineResponse200154) GetV3() InlineResponse200154V3 {
	if o == nil || isNil(o.V3) {
		var ret InlineResponse200154V3
		return ret
	}
	return *o.V3
}

// GetV3Ok returns a tuple with the V3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200154) GetV3Ok() (*InlineResponse200154V3, bool) {
	if o == nil || isNil(o.V3) {
    return nil, false
	}
	return o.V3, true
}

// HasV3 returns a boolean if a field has been set.
func (o *InlineResponse200154) HasV3() bool {
	if o != nil && !isNil(o.V3) {
		return true
	}

	return false
}

// SetV3 gets a reference to the given InlineResponse200154V3 and assigns it to the V3 field.
func (o *InlineResponse200154) SetV3(v InlineResponse200154V3) {
	o.V3 = &v
}

// GetMd5AuthenticationEnabled returns the Md5AuthenticationEnabled field value if set, zero value otherwise.
func (o *InlineResponse200154) GetMd5AuthenticationEnabled() bool {
	if o == nil || isNil(o.Md5AuthenticationEnabled) {
		var ret bool
		return ret
	}
	return *o.Md5AuthenticationEnabled
}

// GetMd5AuthenticationEnabledOk returns a tuple with the Md5AuthenticationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200154) GetMd5AuthenticationEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Md5AuthenticationEnabled) {
    return nil, false
	}
	return o.Md5AuthenticationEnabled, true
}

// HasMd5AuthenticationEnabled returns a boolean if a field has been set.
func (o *InlineResponse200154) HasMd5AuthenticationEnabled() bool {
	if o != nil && !isNil(o.Md5AuthenticationEnabled) {
		return true
	}

	return false
}

// SetMd5AuthenticationEnabled gets a reference to the given bool and assigns it to the Md5AuthenticationEnabled field.
func (o *InlineResponse200154) SetMd5AuthenticationEnabled(v bool) {
	o.Md5AuthenticationEnabled = &v
}

// GetMd5AuthenticationKey returns the Md5AuthenticationKey field value if set, zero value otherwise.
func (o *InlineResponse200154) GetMd5AuthenticationKey() InlineResponse200154Md5AuthenticationKey {
	if o == nil || isNil(o.Md5AuthenticationKey) {
		var ret InlineResponse200154Md5AuthenticationKey
		return ret
	}
	return *o.Md5AuthenticationKey
}

// GetMd5AuthenticationKeyOk returns a tuple with the Md5AuthenticationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200154) GetMd5AuthenticationKeyOk() (*InlineResponse200154Md5AuthenticationKey, bool) {
	if o == nil || isNil(o.Md5AuthenticationKey) {
    return nil, false
	}
	return o.Md5AuthenticationKey, true
}

// HasMd5AuthenticationKey returns a boolean if a field has been set.
func (o *InlineResponse200154) HasMd5AuthenticationKey() bool {
	if o != nil && !isNil(o.Md5AuthenticationKey) {
		return true
	}

	return false
}

// SetMd5AuthenticationKey gets a reference to the given InlineResponse200154Md5AuthenticationKey and assigns it to the Md5AuthenticationKey field.
func (o *InlineResponse200154) SetMd5AuthenticationKey(v InlineResponse200154Md5AuthenticationKey) {
	o.Md5AuthenticationKey = &v
}

func (o InlineResponse200154) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200154 struct {
	value *InlineResponse200154
	isSet bool
}

func (v NullableInlineResponse200154) Get() *InlineResponse200154 {
	return v.value
}

func (v *NullableInlineResponse200154) Set(val *InlineResponse200154) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200154) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200154) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200154(val *InlineResponse200154) *NullableInlineResponse200154 {
	return &NullableInlineResponse200154{value: val, isSet: true}
}

func (v NullableInlineResponse200154) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200154) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


