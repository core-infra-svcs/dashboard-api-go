/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200209 struct for InlineResponse200209
type InlineResponse200209 struct {
	// Network identifier
	NetworkId *string `json:"networkId,omitempty"`
	// The uplink serial
	Serial *string `json:"serial,omitempty"`
	// The uplink model
	Model *string `json:"model,omitempty"`
	// Last reported time for the device
	LastReportedAt *time.Time `json:"lastReportedAt,omitempty"`
	HighAvailability *OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability `json:"highAvailability,omitempty"`
	// Uplinks
	Uplinks []OrganizationsOrganizationIdApplianceUplinkStatusesUplinks `json:"uplinks,omitempty"`
}

// NewInlineResponse200209 instantiates a new InlineResponse200209 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200209() *InlineResponse200209 {
	this := InlineResponse200209{}
	return &this
}

// NewInlineResponse200209WithDefaults instantiates a new InlineResponse200209 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200209WithDefaults() *InlineResponse200209 {
	this := InlineResponse200209{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200209) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200209) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200209) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200209) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200209) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200209) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200209) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200209) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse200209) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200209) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse200209) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse200209) SetModel(v string) {
	o.Model = &v
}

// GetLastReportedAt returns the LastReportedAt field value if set, zero value otherwise.
func (o *InlineResponse200209) GetLastReportedAt() time.Time {
	if o == nil || isNil(o.LastReportedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastReportedAt
}

// GetLastReportedAtOk returns a tuple with the LastReportedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200209) GetLastReportedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastReportedAt) {
    return nil, false
	}
	return o.LastReportedAt, true
}

// HasLastReportedAt returns a boolean if a field has been set.
func (o *InlineResponse200209) HasLastReportedAt() bool {
	if o != nil && !isNil(o.LastReportedAt) {
		return true
	}

	return false
}

// SetLastReportedAt gets a reference to the given time.Time and assigns it to the LastReportedAt field.
func (o *InlineResponse200209) SetLastReportedAt(v time.Time) {
	o.LastReportedAt = &v
}

// GetHighAvailability returns the HighAvailability field value if set, zero value otherwise.
func (o *InlineResponse200209) GetHighAvailability() OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability {
	if o == nil || isNil(o.HighAvailability) {
		var ret OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability
		return ret
	}
	return *o.HighAvailability
}

// GetHighAvailabilityOk returns a tuple with the HighAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200209) GetHighAvailabilityOk() (*OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability, bool) {
	if o == nil || isNil(o.HighAvailability) {
    return nil, false
	}
	return o.HighAvailability, true
}

// HasHighAvailability returns a boolean if a field has been set.
func (o *InlineResponse200209) HasHighAvailability() bool {
	if o != nil && !isNil(o.HighAvailability) {
		return true
	}

	return false
}

// SetHighAvailability gets a reference to the given OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability and assigns it to the HighAvailability field.
func (o *InlineResponse200209) SetHighAvailability(v OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) {
	o.HighAvailability = &v
}

// GetUplinks returns the Uplinks field value if set, zero value otherwise.
func (o *InlineResponse200209) GetUplinks() []OrganizationsOrganizationIdApplianceUplinkStatusesUplinks {
	if o == nil || isNil(o.Uplinks) {
		var ret []OrganizationsOrganizationIdApplianceUplinkStatusesUplinks
		return ret
	}
	return o.Uplinks
}

// GetUplinksOk returns a tuple with the Uplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200209) GetUplinksOk() ([]OrganizationsOrganizationIdApplianceUplinkStatusesUplinks, bool) {
	if o == nil || isNil(o.Uplinks) {
    return nil, false
	}
	return o.Uplinks, true
}

// HasUplinks returns a boolean if a field has been set.
func (o *InlineResponse200209) HasUplinks() bool {
	if o != nil && !isNil(o.Uplinks) {
		return true
	}

	return false
}

// SetUplinks gets a reference to the given []OrganizationsOrganizationIdApplianceUplinkStatusesUplinks and assigns it to the Uplinks field.
func (o *InlineResponse200209) SetUplinks(v []OrganizationsOrganizationIdApplianceUplinkStatusesUplinks) {
	o.Uplinks = v
}

func (o InlineResponse200209) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.LastReportedAt) {
		toSerialize["lastReportedAt"] = o.LastReportedAt
	}
	if !isNil(o.HighAvailability) {
		toSerialize["highAvailability"] = o.HighAvailability
	}
	if !isNil(o.Uplinks) {
		toSerialize["uplinks"] = o.Uplinks
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200209 struct {
	value *InlineResponse200209
	isSet bool
}

func (v NullableInlineResponse200209) Get() *InlineResponse200209 {
	return v.value
}

func (v *NullableInlineResponse200209) Set(val *InlineResponse200209) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200209) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200209) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200209(val *InlineResponse200209) *NullableInlineResponse200209 {
	return &NullableInlineResponse200209{value: val, isSet: true}
}

func (v NullableInlineResponse200209) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200209) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


