/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject139 struct for InlineObject139
type InlineObject139 struct {
	// Whether or not Hotspot 2.0 for this SSID is enabled
	Enabled *bool `json:"enabled,omitempty"`
	Operator *NetworksNetworkIdWirelessSsidsNumberHotspot20Operator `json:"operator,omitempty"`
	Venue *NetworksNetworkIdWirelessSsidsNumberHotspot20Venue `json:"venue,omitempty"`
	// The network type of this SSID ('Private network', 'Private network with guest access', 'Chargeable public network', 'Free public network', 'Personal device network', 'Emergency services only network', 'Test or experimental', 'Wildcard')
	NetworkAccessType *string `json:"networkAccessType,omitempty"`
	// An array of domain names
	Domains *[]string `json:"domains,omitempty"`
	// An array of roaming consortium OIs (hexadecimal number 3-5 octets in length)
	RoamConsortOis *[]string `json:"roamConsortOis,omitempty"`
	// An array of MCC/MNC pairs
	MccMncs *[]NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs `json:"mccMncs,omitempty"`
	// An array of NAI realms
	NaiRealms *[]NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms `json:"naiRealms,omitempty"`
}

// NewInlineObject139 instantiates a new InlineObject139 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject139() *InlineObject139 {
	this := InlineObject139{}
	return &this
}

// NewInlineObject139WithDefaults instantiates a new InlineObject139 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject139WithDefaults() *InlineObject139 {
	this := InlineObject139{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject139) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject139) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject139) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject139) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *InlineObject139) GetOperator() NetworksNetworkIdWirelessSsidsNumberHotspot20Operator {
	if o == nil || o.Operator == nil {
		var ret NetworksNetworkIdWirelessSsidsNumberHotspot20Operator
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject139) GetOperatorOk() (*NetworksNetworkIdWirelessSsidsNumberHotspot20Operator, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *InlineObject139) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given NetworksNetworkIdWirelessSsidsNumberHotspot20Operator and assigns it to the Operator field.
func (o *InlineObject139) SetOperator(v NetworksNetworkIdWirelessSsidsNumberHotspot20Operator) {
	o.Operator = &v
}

// GetVenue returns the Venue field value if set, zero value otherwise.
func (o *InlineObject139) GetVenue() NetworksNetworkIdWirelessSsidsNumberHotspot20Venue {
	if o == nil || o.Venue == nil {
		var ret NetworksNetworkIdWirelessSsidsNumberHotspot20Venue
		return ret
	}
	return *o.Venue
}

// GetVenueOk returns a tuple with the Venue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject139) GetVenueOk() (*NetworksNetworkIdWirelessSsidsNumberHotspot20Venue, bool) {
	if o == nil || o.Venue == nil {
		return nil, false
	}
	return o.Venue, true
}

// HasVenue returns a boolean if a field has been set.
func (o *InlineObject139) HasVenue() bool {
	if o != nil && o.Venue != nil {
		return true
	}

	return false
}

// SetVenue gets a reference to the given NetworksNetworkIdWirelessSsidsNumberHotspot20Venue and assigns it to the Venue field.
func (o *InlineObject139) SetVenue(v NetworksNetworkIdWirelessSsidsNumberHotspot20Venue) {
	o.Venue = &v
}

// GetNetworkAccessType returns the NetworkAccessType field value if set, zero value otherwise.
func (o *InlineObject139) GetNetworkAccessType() string {
	if o == nil || o.NetworkAccessType == nil {
		var ret string
		return ret
	}
	return *o.NetworkAccessType
}

// GetNetworkAccessTypeOk returns a tuple with the NetworkAccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject139) GetNetworkAccessTypeOk() (*string, bool) {
	if o == nil || o.NetworkAccessType == nil {
		return nil, false
	}
	return o.NetworkAccessType, true
}

// HasNetworkAccessType returns a boolean if a field has been set.
func (o *InlineObject139) HasNetworkAccessType() bool {
	if o != nil && o.NetworkAccessType != nil {
		return true
	}

	return false
}

// SetNetworkAccessType gets a reference to the given string and assigns it to the NetworkAccessType field.
func (o *InlineObject139) SetNetworkAccessType(v string) {
	o.NetworkAccessType = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *InlineObject139) GetDomains() []string {
	if o == nil || o.Domains == nil {
		var ret []string
		return ret
	}
	return *o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject139) GetDomainsOk() (*[]string, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *InlineObject139) HasDomains() bool {
	if o != nil && o.Domains != nil {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *InlineObject139) SetDomains(v []string) {
	o.Domains = &v
}

// GetRoamConsortOis returns the RoamConsortOis field value if set, zero value otherwise.
func (o *InlineObject139) GetRoamConsortOis() []string {
	if o == nil || o.RoamConsortOis == nil {
		var ret []string
		return ret
	}
	return *o.RoamConsortOis
}

// GetRoamConsortOisOk returns a tuple with the RoamConsortOis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject139) GetRoamConsortOisOk() (*[]string, bool) {
	if o == nil || o.RoamConsortOis == nil {
		return nil, false
	}
	return o.RoamConsortOis, true
}

// HasRoamConsortOis returns a boolean if a field has been set.
func (o *InlineObject139) HasRoamConsortOis() bool {
	if o != nil && o.RoamConsortOis != nil {
		return true
	}

	return false
}

// SetRoamConsortOis gets a reference to the given []string and assigns it to the RoamConsortOis field.
func (o *InlineObject139) SetRoamConsortOis(v []string) {
	o.RoamConsortOis = &v
}

// GetMccMncs returns the MccMncs field value if set, zero value otherwise.
func (o *InlineObject139) GetMccMncs() []NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs {
	if o == nil || o.MccMncs == nil {
		var ret []NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs
		return ret
	}
	return *o.MccMncs
}

// GetMccMncsOk returns a tuple with the MccMncs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject139) GetMccMncsOk() (*[]NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs, bool) {
	if o == nil || o.MccMncs == nil {
		return nil, false
	}
	return o.MccMncs, true
}

// HasMccMncs returns a boolean if a field has been set.
func (o *InlineObject139) HasMccMncs() bool {
	if o != nil && o.MccMncs != nil {
		return true
	}

	return false
}

// SetMccMncs gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs and assigns it to the MccMncs field.
func (o *InlineObject139) SetMccMncs(v []NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) {
	o.MccMncs = &v
}

// GetNaiRealms returns the NaiRealms field value if set, zero value otherwise.
func (o *InlineObject139) GetNaiRealms() []NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms {
	if o == nil || o.NaiRealms == nil {
		var ret []NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms
		return ret
	}
	return *o.NaiRealms
}

// GetNaiRealmsOk returns a tuple with the NaiRealms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject139) GetNaiRealmsOk() (*[]NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms, bool) {
	if o == nil || o.NaiRealms == nil {
		return nil, false
	}
	return o.NaiRealms, true
}

// HasNaiRealms returns a boolean if a field has been set.
func (o *InlineObject139) HasNaiRealms() bool {
	if o != nil && o.NaiRealms != nil {
		return true
	}

	return false
}

// SetNaiRealms gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms and assigns it to the NaiRealms field.
func (o *InlineObject139) SetNaiRealms(v []NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) {
	o.NaiRealms = &v
}

func (o InlineObject139) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.Venue != nil {
		toSerialize["venue"] = o.Venue
	}
	if o.NetworkAccessType != nil {
		toSerialize["networkAccessType"] = o.NetworkAccessType
	}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	if o.RoamConsortOis != nil {
		toSerialize["roamConsortOis"] = o.RoamConsortOis
	}
	if o.MccMncs != nil {
		toSerialize["mccMncs"] = o.MccMncs
	}
	if o.NaiRealms != nil {
		toSerialize["naiRealms"] = o.NaiRealms
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject139 struct {
	value *InlineObject139
	isSet bool
}

func (v NullableInlineObject139) Get() *InlineObject139 {
	return v.value
}

func (v *NullableInlineObject139) Set(val *InlineObject139) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject139) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject139) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject139(val *InlineObject139) *NullableInlineObject139 {
	return &NullableInlineObject139{value: val, isSet: true}
}

func (v NullableInlineObject139) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject139) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


