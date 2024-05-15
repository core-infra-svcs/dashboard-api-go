/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject119 struct for InlineObject119
type InlineObject119 struct {
	// The wifiMacs of the endpoints to be rebooted.
	WifiMacs []string `json:"wifiMacs,omitempty"`
	// The ids of the endpoints to be rebooted.
	Ids []string `json:"ids,omitempty"`
	// The serials of the endpoints to be rebooted.
	Serials []string `json:"serials,omitempty"`
	// The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the endpoints to be rebooted.
	Scope []string `json:"scope,omitempty"`
	// The KextPaths of the endpoints to be rebooted. Available for macOS 11+
	KextPaths []string `json:"kextPaths,omitempty"`
	// Whether or not to notify the user before rebooting the endpoint. Available for macOS 11.3+
	NotifyUser *bool `json:"notifyUser,omitempty"`
	// Whether or not to rebuild the kernel cache when rebooting the endpoint. Available for macOS 11+
	RebuildKernelCache *bool `json:"rebuildKernelCache,omitempty"`
	// Whether or not the request requires network tethering. Available for macOS and supervised iOS or tvOS
	RequestRequiresNetworkTether *bool `json:"requestRequiresNetworkTether,omitempty"`
}

// NewInlineObject119 instantiates a new InlineObject119 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject119() *InlineObject119 {
	this := InlineObject119{}
	return &this
}

// NewInlineObject119WithDefaults instantiates a new InlineObject119 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject119WithDefaults() *InlineObject119 {
	this := InlineObject119{}
	return &this
}

// GetWifiMacs returns the WifiMacs field value if set, zero value otherwise.
func (o *InlineObject119) GetWifiMacs() []string {
	if o == nil || isNil(o.WifiMacs) {
		var ret []string
		return ret
	}
	return o.WifiMacs
}

// GetWifiMacsOk returns a tuple with the WifiMacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject119) GetWifiMacsOk() ([]string, bool) {
	if o == nil || isNil(o.WifiMacs) {
    return nil, false
	}
	return o.WifiMacs, true
}

// HasWifiMacs returns a boolean if a field has been set.
func (o *InlineObject119) HasWifiMacs() bool {
	if o != nil && !isNil(o.WifiMacs) {
		return true
	}

	return false
}

// SetWifiMacs gets a reference to the given []string and assigns it to the WifiMacs field.
func (o *InlineObject119) SetWifiMacs(v []string) {
	o.WifiMacs = v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InlineObject119) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject119) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InlineObject119) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *InlineObject119) SetIds(v []string) {
	o.Ids = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineObject119) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject119) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineObject119) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineObject119) SetSerials(v []string) {
	o.Serials = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineObject119) GetScope() []string {
	if o == nil || isNil(o.Scope) {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject119) GetScopeOk() ([]string, bool) {
	if o == nil || isNil(o.Scope) {
    return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineObject119) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *InlineObject119) SetScope(v []string) {
	o.Scope = v
}

// GetKextPaths returns the KextPaths field value if set, zero value otherwise.
func (o *InlineObject119) GetKextPaths() []string {
	if o == nil || isNil(o.KextPaths) {
		var ret []string
		return ret
	}
	return o.KextPaths
}

// GetKextPathsOk returns a tuple with the KextPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject119) GetKextPathsOk() ([]string, bool) {
	if o == nil || isNil(o.KextPaths) {
    return nil, false
	}
	return o.KextPaths, true
}

// HasKextPaths returns a boolean if a field has been set.
func (o *InlineObject119) HasKextPaths() bool {
	if o != nil && !isNil(o.KextPaths) {
		return true
	}

	return false
}

// SetKextPaths gets a reference to the given []string and assigns it to the KextPaths field.
func (o *InlineObject119) SetKextPaths(v []string) {
	o.KextPaths = v
}

// GetNotifyUser returns the NotifyUser field value if set, zero value otherwise.
func (o *InlineObject119) GetNotifyUser() bool {
	if o == nil || isNil(o.NotifyUser) {
		var ret bool
		return ret
	}
	return *o.NotifyUser
}

// GetNotifyUserOk returns a tuple with the NotifyUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject119) GetNotifyUserOk() (*bool, bool) {
	if o == nil || isNil(o.NotifyUser) {
    return nil, false
	}
	return o.NotifyUser, true
}

// HasNotifyUser returns a boolean if a field has been set.
func (o *InlineObject119) HasNotifyUser() bool {
	if o != nil && !isNil(o.NotifyUser) {
		return true
	}

	return false
}

// SetNotifyUser gets a reference to the given bool and assigns it to the NotifyUser field.
func (o *InlineObject119) SetNotifyUser(v bool) {
	o.NotifyUser = &v
}

// GetRebuildKernelCache returns the RebuildKernelCache field value if set, zero value otherwise.
func (o *InlineObject119) GetRebuildKernelCache() bool {
	if o == nil || isNil(o.RebuildKernelCache) {
		var ret bool
		return ret
	}
	return *o.RebuildKernelCache
}

// GetRebuildKernelCacheOk returns a tuple with the RebuildKernelCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject119) GetRebuildKernelCacheOk() (*bool, bool) {
	if o == nil || isNil(o.RebuildKernelCache) {
    return nil, false
	}
	return o.RebuildKernelCache, true
}

// HasRebuildKernelCache returns a boolean if a field has been set.
func (o *InlineObject119) HasRebuildKernelCache() bool {
	if o != nil && !isNil(o.RebuildKernelCache) {
		return true
	}

	return false
}

// SetRebuildKernelCache gets a reference to the given bool and assigns it to the RebuildKernelCache field.
func (o *InlineObject119) SetRebuildKernelCache(v bool) {
	o.RebuildKernelCache = &v
}

// GetRequestRequiresNetworkTether returns the RequestRequiresNetworkTether field value if set, zero value otherwise.
func (o *InlineObject119) GetRequestRequiresNetworkTether() bool {
	if o == nil || isNil(o.RequestRequiresNetworkTether) {
		var ret bool
		return ret
	}
	return *o.RequestRequiresNetworkTether
}

// GetRequestRequiresNetworkTetherOk returns a tuple with the RequestRequiresNetworkTether field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject119) GetRequestRequiresNetworkTetherOk() (*bool, bool) {
	if o == nil || isNil(o.RequestRequiresNetworkTether) {
    return nil, false
	}
	return o.RequestRequiresNetworkTether, true
}

// HasRequestRequiresNetworkTether returns a boolean if a field has been set.
func (o *InlineObject119) HasRequestRequiresNetworkTether() bool {
	if o != nil && !isNil(o.RequestRequiresNetworkTether) {
		return true
	}

	return false
}

// SetRequestRequiresNetworkTether gets a reference to the given bool and assigns it to the RequestRequiresNetworkTether field.
func (o *InlineObject119) SetRequestRequiresNetworkTether(v bool) {
	o.RequestRequiresNetworkTether = &v
}

func (o InlineObject119) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WifiMacs) {
		toSerialize["wifiMacs"] = o.WifiMacs
	}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	if !isNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !isNil(o.KextPaths) {
		toSerialize["kextPaths"] = o.KextPaths
	}
	if !isNil(o.NotifyUser) {
		toSerialize["notifyUser"] = o.NotifyUser
	}
	if !isNil(o.RebuildKernelCache) {
		toSerialize["rebuildKernelCache"] = o.RebuildKernelCache
	}
	if !isNil(o.RequestRequiresNetworkTether) {
		toSerialize["requestRequiresNetworkTether"] = o.RequestRequiresNetworkTether
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject119 struct {
	value *InlineObject119
	isSet bool
}

func (v NullableInlineObject119) Get() *InlineObject119 {
	return v.value
}

func (v *NullableInlineObject119) Set(val *InlineObject119) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject119) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject119) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject119(val *InlineObject119) *NullableInlineObject119 {
	return &NullableInlineObject119{value: val, isSet: true}
}

func (v NullableInlineObject119) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject119) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


