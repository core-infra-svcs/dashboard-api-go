/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200127 struct for InlineResponse200127
type InlineResponse200127 struct {
	// Device Name
	Name *string `json:"name,omitempty"`
	// Device Serial Number
	Serial *string `json:"serial,omitempty"`
	// MAC Address
	Mac *string `json:"mac,omitempty"`
	// Public IP Address
	PublicIp *string `json:"publicIp,omitempty"`
	// Network ID
	NetworkId *string `json:"networkId,omitempty"`
	// Device Status
	Status *string `json:"status,omitempty"`
	// Device Last Reported Location
	LastReportedAt *string `json:"lastReportedAt,omitempty"`
	// LAN IP Address
	LanIp *string `json:"lanIp,omitempty"`
	// IP Gateway
	Gateway *string `json:"gateway,omitempty"`
	// IP Type
	IpType *string `json:"ipType,omitempty"`
	// Primary DNS
	PrimaryDns *string `json:"primaryDns,omitempty"`
	// Secondary DNS
	SecondaryDns *string `json:"secondaryDns,omitempty"`
	// Product Type
	ProductType *string `json:"productType,omitempty"`
	Components *OrganizationsOrganizationIdDevicesStatusesComponents `json:"components,omitempty"`
	// Model
	Model *string `json:"model,omitempty"`
	// Tags
	Tags []string `json:"tags,omitempty"`
}

// NewInlineResponse200127 instantiates a new InlineResponse200127 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200127() *InlineResponse200127 {
	this := InlineResponse200127{}
	return &this
}

// NewInlineResponse200127WithDefaults instantiates a new InlineResponse200127 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200127WithDefaults() *InlineResponse200127 {
	this := InlineResponse200127{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200127) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200127) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200127) SetName(v string) {
	o.Name = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200127) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200127) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200127) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200127) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200127) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200127) SetMac(v string) {
	o.Mac = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *InlineResponse200127) GetPublicIp() string {
	if o == nil || isNil(o.PublicIp) {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetPublicIpOk() (*string, bool) {
	if o == nil || isNil(o.PublicIp) {
    return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *InlineResponse200127) HasPublicIp() bool {
	if o != nil && !isNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *InlineResponse200127) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200127) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200127) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200127) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse200127) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse200127) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse200127) SetStatus(v string) {
	o.Status = &v
}

// GetLastReportedAt returns the LastReportedAt field value if set, zero value otherwise.
func (o *InlineResponse200127) GetLastReportedAt() string {
	if o == nil || isNil(o.LastReportedAt) {
		var ret string
		return ret
	}
	return *o.LastReportedAt
}

// GetLastReportedAtOk returns a tuple with the LastReportedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetLastReportedAtOk() (*string, bool) {
	if o == nil || isNil(o.LastReportedAt) {
    return nil, false
	}
	return o.LastReportedAt, true
}

// HasLastReportedAt returns a boolean if a field has been set.
func (o *InlineResponse200127) HasLastReportedAt() bool {
	if o != nil && !isNil(o.LastReportedAt) {
		return true
	}

	return false
}

// SetLastReportedAt gets a reference to the given string and assigns it to the LastReportedAt field.
func (o *InlineResponse200127) SetLastReportedAt(v string) {
	o.LastReportedAt = &v
}

// GetLanIp returns the LanIp field value if set, zero value otherwise.
func (o *InlineResponse200127) GetLanIp() string {
	if o == nil || isNil(o.LanIp) {
		var ret string
		return ret
	}
	return *o.LanIp
}

// GetLanIpOk returns a tuple with the LanIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetLanIpOk() (*string, bool) {
	if o == nil || isNil(o.LanIp) {
    return nil, false
	}
	return o.LanIp, true
}

// HasLanIp returns a boolean if a field has been set.
func (o *InlineResponse200127) HasLanIp() bool {
	if o != nil && !isNil(o.LanIp) {
		return true
	}

	return false
}

// SetLanIp gets a reference to the given string and assigns it to the LanIp field.
func (o *InlineResponse200127) SetLanIp(v string) {
	o.LanIp = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *InlineResponse200127) GetGateway() string {
	if o == nil || isNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetGatewayOk() (*string, bool) {
	if o == nil || isNil(o.Gateway) {
    return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *InlineResponse200127) HasGateway() bool {
	if o != nil && !isNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *InlineResponse200127) SetGateway(v string) {
	o.Gateway = &v
}

// GetIpType returns the IpType field value if set, zero value otherwise.
func (o *InlineResponse200127) GetIpType() string {
	if o == nil || isNil(o.IpType) {
		var ret string
		return ret
	}
	return *o.IpType
}

// GetIpTypeOk returns a tuple with the IpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetIpTypeOk() (*string, bool) {
	if o == nil || isNil(o.IpType) {
    return nil, false
	}
	return o.IpType, true
}

// HasIpType returns a boolean if a field has been set.
func (o *InlineResponse200127) HasIpType() bool {
	if o != nil && !isNil(o.IpType) {
		return true
	}

	return false
}

// SetIpType gets a reference to the given string and assigns it to the IpType field.
func (o *InlineResponse200127) SetIpType(v string) {
	o.IpType = &v
}

// GetPrimaryDns returns the PrimaryDns field value if set, zero value otherwise.
func (o *InlineResponse200127) GetPrimaryDns() string {
	if o == nil || isNil(o.PrimaryDns) {
		var ret string
		return ret
	}
	return *o.PrimaryDns
}

// GetPrimaryDnsOk returns a tuple with the PrimaryDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetPrimaryDnsOk() (*string, bool) {
	if o == nil || isNil(o.PrimaryDns) {
    return nil, false
	}
	return o.PrimaryDns, true
}

// HasPrimaryDns returns a boolean if a field has been set.
func (o *InlineResponse200127) HasPrimaryDns() bool {
	if o != nil && !isNil(o.PrimaryDns) {
		return true
	}

	return false
}

// SetPrimaryDns gets a reference to the given string and assigns it to the PrimaryDns field.
func (o *InlineResponse200127) SetPrimaryDns(v string) {
	o.PrimaryDns = &v
}

// GetSecondaryDns returns the SecondaryDns field value if set, zero value otherwise.
func (o *InlineResponse200127) GetSecondaryDns() string {
	if o == nil || isNil(o.SecondaryDns) {
		var ret string
		return ret
	}
	return *o.SecondaryDns
}

// GetSecondaryDnsOk returns a tuple with the SecondaryDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetSecondaryDnsOk() (*string, bool) {
	if o == nil || isNil(o.SecondaryDns) {
    return nil, false
	}
	return o.SecondaryDns, true
}

// HasSecondaryDns returns a boolean if a field has been set.
func (o *InlineResponse200127) HasSecondaryDns() bool {
	if o != nil && !isNil(o.SecondaryDns) {
		return true
	}

	return false
}

// SetSecondaryDns gets a reference to the given string and assigns it to the SecondaryDns field.
func (o *InlineResponse200127) SetSecondaryDns(v string) {
	o.SecondaryDns = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *InlineResponse200127) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *InlineResponse200127) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *InlineResponse200127) SetProductType(v string) {
	o.ProductType = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *InlineResponse200127) GetComponents() OrganizationsOrganizationIdDevicesStatusesComponents {
	if o == nil || isNil(o.Components) {
		var ret OrganizationsOrganizationIdDevicesStatusesComponents
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetComponentsOk() (*OrganizationsOrganizationIdDevicesStatusesComponents, bool) {
	if o == nil || isNil(o.Components) {
    return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *InlineResponse200127) HasComponents() bool {
	if o != nil && !isNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given OrganizationsOrganizationIdDevicesStatusesComponents and assigns it to the Components field.
func (o *InlineResponse200127) SetComponents(v OrganizationsOrganizationIdDevicesStatusesComponents) {
	o.Components = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse200127) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse200127) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse200127) SetModel(v string) {
	o.Model = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200127) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200127) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse200127) SetTags(v []string) {
	o.Tags = v
}

func (o InlineResponse200127) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.PublicIp) {
		toSerialize["publicIp"] = o.PublicIp
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.LastReportedAt) {
		toSerialize["lastReportedAt"] = o.LastReportedAt
	}
	if !isNil(o.LanIp) {
		toSerialize["lanIp"] = o.LanIp
	}
	if !isNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !isNil(o.IpType) {
		toSerialize["ipType"] = o.IpType
	}
	if !isNil(o.PrimaryDns) {
		toSerialize["primaryDns"] = o.PrimaryDns
	}
	if !isNil(o.SecondaryDns) {
		toSerialize["secondaryDns"] = o.SecondaryDns
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !isNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200127 struct {
	value *InlineResponse200127
	isSet bool
}

func (v NullableInlineResponse200127) Get() *InlineResponse200127 {
	return v.value
}

func (v *NullableInlineResponse200127) Set(val *InlineResponse200127) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200127) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200127) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200127(val *InlineResponse200127) *NullableInlineResponse200127 {
	return &NullableInlineResponse200127{value: val, isSet: true}
}

func (v NullableInlineResponse200127) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200127) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


