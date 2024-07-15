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

// OrganizationsOrganizationIdAssuranceAlertsScope Scope of the alert (which devices and networks are affected)
type OrganizationsOrganizationIdAssuranceAlertsScope struct {
	// Description of affected devices
	Devices []OrganizationsOrganizationIdAssuranceAlertsScopeDevices `json:"devices,omitempty"`
	// Applications affected by the alert
	Applications []map[string]interface{} `json:"applications,omitempty"`
	// Peers affected by the alert
	Peers []map[string]interface{} `json:"peers,omitempty"`
}

// NewOrganizationsOrganizationIdAssuranceAlertsScope instantiates a new OrganizationsOrganizationIdAssuranceAlertsScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdAssuranceAlertsScope() *OrganizationsOrganizationIdAssuranceAlertsScope {
	this := OrganizationsOrganizationIdAssuranceAlertsScope{}
	return &this
}

// NewOrganizationsOrganizationIdAssuranceAlertsScopeWithDefaults instantiates a new OrganizationsOrganizationIdAssuranceAlertsScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdAssuranceAlertsScopeWithDefaults() *OrganizationsOrganizationIdAssuranceAlertsScope {
	this := OrganizationsOrganizationIdAssuranceAlertsScope{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAssuranceAlertsScope) GetDevices() []OrganizationsOrganizationIdAssuranceAlertsScopeDevices {
	if o == nil || isNil(o.Devices) {
		var ret []OrganizationsOrganizationIdAssuranceAlertsScopeDevices
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScope) GetDevicesOk() ([]OrganizationsOrganizationIdAssuranceAlertsScopeDevices, bool) {
	if o == nil || isNil(o.Devices) {
    return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScope) HasDevices() bool {
	if o != nil && !isNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []OrganizationsOrganizationIdAssuranceAlertsScopeDevices and assigns it to the Devices field.
func (o *OrganizationsOrganizationIdAssuranceAlertsScope) SetDevices(v []OrganizationsOrganizationIdAssuranceAlertsScopeDevices) {
	o.Devices = v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAssuranceAlertsScope) GetApplications() []map[string]interface{} {
	if o == nil || isNil(o.Applications) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScope) GetApplicationsOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.Applications) {
    return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScope) HasApplications() bool {
	if o != nil && !isNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []map[string]interface{} and assigns it to the Applications field.
func (o *OrganizationsOrganizationIdAssuranceAlertsScope) SetApplications(v []map[string]interface{}) {
	o.Applications = v
}

// GetPeers returns the Peers field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAssuranceAlertsScope) GetPeers() []map[string]interface{} {
	if o == nil || isNil(o.Peers) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScope) GetPeersOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.Peers) {
    return nil, false
	}
	return o.Peers, true
}

// HasPeers returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScope) HasPeers() bool {
	if o != nil && !isNil(o.Peers) {
		return true
	}

	return false
}

// SetPeers gets a reference to the given []map[string]interface{} and assigns it to the Peers field.
func (o *OrganizationsOrganizationIdAssuranceAlertsScope) SetPeers(v []map[string]interface{}) {
	o.Peers = v
}

func (o OrganizationsOrganizationIdAssuranceAlertsScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !isNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	if !isNil(o.Peers) {
		toSerialize["peers"] = o.Peers
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdAssuranceAlertsScope struct {
	value *OrganizationsOrganizationIdAssuranceAlertsScope
	isSet bool
}

func (v NullableOrganizationsOrganizationIdAssuranceAlertsScope) Get() *OrganizationsOrganizationIdAssuranceAlertsScope {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdAssuranceAlertsScope) Set(val *OrganizationsOrganizationIdAssuranceAlertsScope) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdAssuranceAlertsScope) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdAssuranceAlertsScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdAssuranceAlertsScope(val *OrganizationsOrganizationIdAssuranceAlertsScope) *NullableOrganizationsOrganizationIdAssuranceAlertsScope {
	return &NullableOrganizationsOrganizationIdAssuranceAlertsScope{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdAssuranceAlertsScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdAssuranceAlertsScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


