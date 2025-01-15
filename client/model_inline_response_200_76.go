/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20076 struct for InlineResponse20076
type InlineResponse20076 struct {
	// The ID of the camera wireless profile.
	Id *string `json:"id,omitempty"`
	// The name of the camera wireless profile.
	Name *string `json:"name,omitempty"`
	// The count of the applied devices.
	AppliedDeviceCount *int32 `json:"appliedDeviceCount,omitempty"`
	Ssid *NetworksNetworkIdCameraWirelessProfilesSsid `json:"ssid,omitempty"`
	Identity *NetworksNetworkIdCameraWirelessProfilesIdentity `json:"identity,omitempty"`
}

// NewInlineResponse20076 instantiates a new InlineResponse20076 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20076() *InlineResponse20076 {
	this := InlineResponse20076{}
	return &this
}

// NewInlineResponse20076WithDefaults instantiates a new InlineResponse20076 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20076WithDefaults() *InlineResponse20076 {
	this := InlineResponse20076{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20076) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20076) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20076) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20076) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20076) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20076) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20076) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20076) SetName(v string) {
	o.Name = &v
}

// GetAppliedDeviceCount returns the AppliedDeviceCount field value if set, zero value otherwise.
func (o *InlineResponse20076) GetAppliedDeviceCount() int32 {
	if o == nil || isNil(o.AppliedDeviceCount) {
		var ret int32
		return ret
	}
	return *o.AppliedDeviceCount
}

// GetAppliedDeviceCountOk returns a tuple with the AppliedDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20076) GetAppliedDeviceCountOk() (*int32, bool) {
	if o == nil || isNil(o.AppliedDeviceCount) {
    return nil, false
	}
	return o.AppliedDeviceCount, true
}

// HasAppliedDeviceCount returns a boolean if a field has been set.
func (o *InlineResponse20076) HasAppliedDeviceCount() bool {
	if o != nil && !isNil(o.AppliedDeviceCount) {
		return true
	}

	return false
}

// SetAppliedDeviceCount gets a reference to the given int32 and assigns it to the AppliedDeviceCount field.
func (o *InlineResponse20076) SetAppliedDeviceCount(v int32) {
	o.AppliedDeviceCount = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *InlineResponse20076) GetSsid() NetworksNetworkIdCameraWirelessProfilesSsid {
	if o == nil || isNil(o.Ssid) {
		var ret NetworksNetworkIdCameraWirelessProfilesSsid
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20076) GetSsidOk() (*NetworksNetworkIdCameraWirelessProfilesSsid, bool) {
	if o == nil || isNil(o.Ssid) {
    return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *InlineResponse20076) HasSsid() bool {
	if o != nil && !isNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given NetworksNetworkIdCameraWirelessProfilesSsid and assigns it to the Ssid field.
func (o *InlineResponse20076) SetSsid(v NetworksNetworkIdCameraWirelessProfilesSsid) {
	o.Ssid = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *InlineResponse20076) GetIdentity() NetworksNetworkIdCameraWirelessProfilesIdentity {
	if o == nil || isNil(o.Identity) {
		var ret NetworksNetworkIdCameraWirelessProfilesIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20076) GetIdentityOk() (*NetworksNetworkIdCameraWirelessProfilesIdentity, bool) {
	if o == nil || isNil(o.Identity) {
    return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *InlineResponse20076) HasIdentity() bool {
	if o != nil && !isNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given NetworksNetworkIdCameraWirelessProfilesIdentity and assigns it to the Identity field.
func (o *InlineResponse20076) SetIdentity(v NetworksNetworkIdCameraWirelessProfilesIdentity) {
	o.Identity = &v
}

func (o InlineResponse20076) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.AppliedDeviceCount) {
		toSerialize["appliedDeviceCount"] = o.AppliedDeviceCount
	}
	if !isNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !isNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20076 struct {
	value *InlineResponse20076
	isSet bool
}

func (v NullableInlineResponse20076) Get() *InlineResponse20076 {
	return v.value
}

func (v *NullableInlineResponse20076) Set(val *InlineResponse20076) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20076) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20076) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20076(val *InlineResponse20076) *NullableInlineResponse20076 {
	return &NullableInlineResponse20076{value: val, isSet: true}
}

func (v NullableInlineResponse20076) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20076) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


