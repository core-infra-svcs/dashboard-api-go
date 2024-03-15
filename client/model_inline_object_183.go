/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject183 struct for InlineObject183
type InlineObject183 struct {
	// Whether or not Hotspot 2.0 for this SSID is enabled
	Enabled *bool `json:"enabled,omitempty"`
	Operator *NetworksNetworkIdWirelessSsidsNumberHotspot20Operator `json:"operator,omitempty"`
	Venue *NetworksNetworkIdWirelessSsidsNumberHotspot20Venue `json:"venue,omitempty"`
	// The network type of this SSID ('Private network', 'Private network with guest access', 'Chargeable public network', 'Free public network', 'Personal device network', 'Emergency services only network', 'Test or experimental', 'Wildcard')
	NetworkAccessType *string `json:"networkAccessType,omitempty"`
	// An array of domain names
	Domains []string `json:"domains,omitempty"`
	// An array of roaming consortium OIs (hexadecimal number 3-5 octets in length)
	RoamConsortOis []string `json:"roamConsortOis,omitempty"`
	// An array of MCC/MNC pairs
	MccMncs []NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs `json:"mccMncs,omitempty"`
	// An array of NAI realms
	NaiRealms []NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms `json:"naiRealms,omitempty"`
}

// NewInlineObject183 instantiates a new InlineObject183 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject183() *InlineObject183 {
	this := InlineObject183{}
	return &this
}

// NewInlineObject183WithDefaults instantiates a new InlineObject183 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject183WithDefaults() *InlineObject183 {
	this := InlineObject183{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject183) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject183) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject183) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *InlineObject183) GetOperator() NetworksNetworkIdWirelessSsidsNumberHotspot20Operator {
	if o == nil || isNil(o.Operator) {
		var ret NetworksNetworkIdWirelessSsidsNumberHotspot20Operator
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetOperatorOk() (*NetworksNetworkIdWirelessSsidsNumberHotspot20Operator, bool) {
	if o == nil || isNil(o.Operator) {
    return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *InlineObject183) HasOperator() bool {
	if o != nil && !isNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given NetworksNetworkIdWirelessSsidsNumberHotspot20Operator and assigns it to the Operator field.
func (o *InlineObject183) SetOperator(v NetworksNetworkIdWirelessSsidsNumberHotspot20Operator) {
	o.Operator = &v
}

// GetVenue returns the Venue field value if set, zero value otherwise.
func (o *InlineObject183) GetVenue() NetworksNetworkIdWirelessSsidsNumberHotspot20Venue {
	if o == nil || isNil(o.Venue) {
		var ret NetworksNetworkIdWirelessSsidsNumberHotspot20Venue
		return ret
	}
	return *o.Venue
}

// GetVenueOk returns a tuple with the Venue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetVenueOk() (*NetworksNetworkIdWirelessSsidsNumberHotspot20Venue, bool) {
	if o == nil || isNil(o.Venue) {
    return nil, false
	}
	return o.Venue, true
}

// HasVenue returns a boolean if a field has been set.
func (o *InlineObject183) HasVenue() bool {
	if o != nil && !isNil(o.Venue) {
		return true
	}

	return false
}

// SetVenue gets a reference to the given NetworksNetworkIdWirelessSsidsNumberHotspot20Venue and assigns it to the Venue field.
func (o *InlineObject183) SetVenue(v NetworksNetworkIdWirelessSsidsNumberHotspot20Venue) {
	o.Venue = &v
}

// GetNetworkAccessType returns the NetworkAccessType field value if set, zero value otherwise.
func (o *InlineObject183) GetNetworkAccessType() string {
	if o == nil || isNil(o.NetworkAccessType) {
		var ret string
		return ret
	}
	return *o.NetworkAccessType
}

// GetNetworkAccessTypeOk returns a tuple with the NetworkAccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetNetworkAccessTypeOk() (*string, bool) {
	if o == nil || isNil(o.NetworkAccessType) {
    return nil, false
	}
	return o.NetworkAccessType, true
}

// HasNetworkAccessType returns a boolean if a field has been set.
func (o *InlineObject183) HasNetworkAccessType() bool {
	if o != nil && !isNil(o.NetworkAccessType) {
		return true
	}

	return false
}

// SetNetworkAccessType gets a reference to the given string and assigns it to the NetworkAccessType field.
func (o *InlineObject183) SetNetworkAccessType(v string) {
	o.NetworkAccessType = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *InlineObject183) GetDomains() []string {
	if o == nil || isNil(o.Domains) {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetDomainsOk() ([]string, bool) {
	if o == nil || isNil(o.Domains) {
    return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *InlineObject183) HasDomains() bool {
	if o != nil && !isNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *InlineObject183) SetDomains(v []string) {
	o.Domains = v
}

// GetRoamConsortOis returns the RoamConsortOis field value if set, zero value otherwise.
func (o *InlineObject183) GetRoamConsortOis() []string {
	if o == nil || isNil(o.RoamConsortOis) {
		var ret []string
		return ret
	}
	return o.RoamConsortOis
}

// GetRoamConsortOisOk returns a tuple with the RoamConsortOis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetRoamConsortOisOk() ([]string, bool) {
	if o == nil || isNil(o.RoamConsortOis) {
    return nil, false
	}
	return o.RoamConsortOis, true
}

// HasRoamConsortOis returns a boolean if a field has been set.
func (o *InlineObject183) HasRoamConsortOis() bool {
	if o != nil && !isNil(o.RoamConsortOis) {
		return true
	}

	return false
}

// SetRoamConsortOis gets a reference to the given []string and assigns it to the RoamConsortOis field.
func (o *InlineObject183) SetRoamConsortOis(v []string) {
	o.RoamConsortOis = v
}

// GetMccMncs returns the MccMncs field value if set, zero value otherwise.
func (o *InlineObject183) GetMccMncs() []NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs {
	if o == nil || isNil(o.MccMncs) {
		var ret []NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs
		return ret
	}
	return o.MccMncs
}

// GetMccMncsOk returns a tuple with the MccMncs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetMccMncsOk() ([]NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs, bool) {
	if o == nil || isNil(o.MccMncs) {
    return nil, false
	}
	return o.MccMncs, true
}

// HasMccMncs returns a boolean if a field has been set.
func (o *InlineObject183) HasMccMncs() bool {
	if o != nil && !isNil(o.MccMncs) {
		return true
	}

	return false
}

// SetMccMncs gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs and assigns it to the MccMncs field.
func (o *InlineObject183) SetMccMncs(v []NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) {
	o.MccMncs = v
}

// GetNaiRealms returns the NaiRealms field value if set, zero value otherwise.
func (o *InlineObject183) GetNaiRealms() []NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms {
	if o == nil || isNil(o.NaiRealms) {
		var ret []NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms
		return ret
	}
	return o.NaiRealms
}

// GetNaiRealmsOk returns a tuple with the NaiRealms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetNaiRealmsOk() ([]NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms, bool) {
	if o == nil || isNil(o.NaiRealms) {
    return nil, false
	}
	return o.NaiRealms, true
}

// HasNaiRealms returns a boolean if a field has been set.
func (o *InlineObject183) HasNaiRealms() bool {
	if o != nil && !isNil(o.NaiRealms) {
		return true
	}

	return false
}

// SetNaiRealms gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms and assigns it to the NaiRealms field.
func (o *InlineObject183) SetNaiRealms(v []NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) {
	o.NaiRealms = v
}

func (o InlineObject183) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !isNil(o.Venue) {
		toSerialize["venue"] = o.Venue
	}
	if !isNil(o.NetworkAccessType) {
		toSerialize["networkAccessType"] = o.NetworkAccessType
	}
	if !isNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}
	if !isNil(o.RoamConsortOis) {
		toSerialize["roamConsortOis"] = o.RoamConsortOis
	}
	if !isNil(o.MccMncs) {
		toSerialize["mccMncs"] = o.MccMncs
	}
	if !isNil(o.NaiRealms) {
		toSerialize["naiRealms"] = o.NaiRealms
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject183 struct {
	value *InlineObject183
	isSet bool
}

func (v NullableInlineObject183) Get() *InlineObject183 {
	return v.value
}

func (v *NullableInlineObject183) Set(val *InlineObject183) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject183) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject183) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject183(val *InlineObject183) *NullableInlineObject183 {
	return &NullableInlineObject183{value: val, isSet: true}
}

func (v NullableInlineObject183) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject183) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


