/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200291 struct for InlineResponse200291
type InlineResponse200291 struct {
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
	Uplinks []OrganizationsOrganizationIdUplinksStatusesUplinks `json:"uplinks,omitempty"`
}

// NewInlineResponse200291 instantiates a new InlineResponse200291 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200291() *InlineResponse200291 {
	this := InlineResponse200291{}
	return &this
}

// NewInlineResponse200291WithDefaults instantiates a new InlineResponse200291 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200291WithDefaults() *InlineResponse200291 {
	this := InlineResponse200291{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200291) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200291) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200291) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200291) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200291) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200291) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200291) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200291) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse200291) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200291) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse200291) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse200291) SetModel(v string) {
	o.Model = &v
}

// GetLastReportedAt returns the LastReportedAt field value if set, zero value otherwise.
func (o *InlineResponse200291) GetLastReportedAt() time.Time {
	if o == nil || isNil(o.LastReportedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastReportedAt
}

// GetLastReportedAtOk returns a tuple with the LastReportedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200291) GetLastReportedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastReportedAt) {
    return nil, false
	}
	return o.LastReportedAt, true
}

// HasLastReportedAt returns a boolean if a field has been set.
func (o *InlineResponse200291) HasLastReportedAt() bool {
	if o != nil && !isNil(o.LastReportedAt) {
		return true
	}

	return false
}

// SetLastReportedAt gets a reference to the given time.Time and assigns it to the LastReportedAt field.
func (o *InlineResponse200291) SetLastReportedAt(v time.Time) {
	o.LastReportedAt = &v
}

// GetHighAvailability returns the HighAvailability field value if set, zero value otherwise.
func (o *InlineResponse200291) GetHighAvailability() OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability {
	if o == nil || isNil(o.HighAvailability) {
		var ret OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability
		return ret
	}
	return *o.HighAvailability
}

// GetHighAvailabilityOk returns a tuple with the HighAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200291) GetHighAvailabilityOk() (*OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability, bool) {
	if o == nil || isNil(o.HighAvailability) {
    return nil, false
	}
	return o.HighAvailability, true
}

// HasHighAvailability returns a boolean if a field has been set.
func (o *InlineResponse200291) HasHighAvailability() bool {
	if o != nil && !isNil(o.HighAvailability) {
		return true
	}

	return false
}

// SetHighAvailability gets a reference to the given OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability and assigns it to the HighAvailability field.
func (o *InlineResponse200291) SetHighAvailability(v OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability) {
	o.HighAvailability = &v
}

// GetUplinks returns the Uplinks field value if set, zero value otherwise.
func (o *InlineResponse200291) GetUplinks() []OrganizationsOrganizationIdUplinksStatusesUplinks {
	if o == nil || isNil(o.Uplinks) {
		var ret []OrganizationsOrganizationIdUplinksStatusesUplinks
		return ret
	}
	return o.Uplinks
}

// GetUplinksOk returns a tuple with the Uplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200291) GetUplinksOk() ([]OrganizationsOrganizationIdUplinksStatusesUplinks, bool) {
	if o == nil || isNil(o.Uplinks) {
    return nil, false
	}
	return o.Uplinks, true
}

// HasUplinks returns a boolean if a field has been set.
func (o *InlineResponse200291) HasUplinks() bool {
	if o != nil && !isNil(o.Uplinks) {
		return true
	}

	return false
}

// SetUplinks gets a reference to the given []OrganizationsOrganizationIdUplinksStatusesUplinks and assigns it to the Uplinks field.
func (o *InlineResponse200291) SetUplinks(v []OrganizationsOrganizationIdUplinksStatusesUplinks) {
	o.Uplinks = v
}

func (o InlineResponse200291) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200291 struct {
	value *InlineResponse200291
	isSet bool
}

func (v NullableInlineResponse200291) Get() *InlineResponse200291 {
	return v.value
}

func (v *NullableInlineResponse200291) Set(val *InlineResponse200291) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200291) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200291) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200291(val *InlineResponse200291) *NullableInlineResponse200291 {
	return &NullableInlineResponse200291{value: val, isSet: true}
}

func (v NullableInlineResponse200291) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200291) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


