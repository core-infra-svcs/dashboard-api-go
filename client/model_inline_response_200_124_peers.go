/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200124Peers struct for InlineResponse200124Peers
type InlineResponse200124Peers struct {
	// The name of the VPN peer
	Name *string `json:"name,omitempty"`
	// [optional] The public IP of the VPN peer
	PublicIp *string `json:"publicIp,omitempty"`
	// [optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN.
	RemoteId *string `json:"remoteId,omitempty"`
	// [optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to.
	LocalId *string `json:"localId,omitempty"`
	// The shared secret with the VPN peer
	Secret *string `json:"secret,omitempty"`
	// The list of the private subnets of the VPN peer
	PrivateSubnets []string `json:"privateSubnets,omitempty"`
	IpsecPolicies *InlineResponse200124IpsecPolicies `json:"ipsecPolicies,omitempty"`
	// One of the following available presets: 'default', 'aws', 'azure'. If this is provided, the 'ipsecPolicies' parameter is ignored.
	IpsecPoliciesPreset *string `json:"ipsecPoliciesPreset,omitempty"`
	// [optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to '1' when omitted.
	IkeVersion *string `json:"ikeVersion,omitempty"`
	// A list of network tags that will connect with this peer. Use ['all'] for all networks. Use ['none'] for no networks. If not included, the default is ['all'].
	NetworkTags []string `json:"networkTags,omitempty"`
}

// NewInlineResponse200124Peers instantiates a new InlineResponse200124Peers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200124Peers() *InlineResponse200124Peers {
	this := InlineResponse200124Peers{}
	var ikeVersion string = "1"
	this.IkeVersion = &ikeVersion
	return &this
}

// NewInlineResponse200124PeersWithDefaults instantiates a new InlineResponse200124Peers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200124PeersWithDefaults() *InlineResponse200124Peers {
	this := InlineResponse200124Peers{}
	var ikeVersion string = "1"
	this.IkeVersion = &ikeVersion
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200124Peers) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Peers) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200124Peers) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200124Peers) SetName(v string) {
	o.Name = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *InlineResponse200124Peers) GetPublicIp() string {
	if o == nil || isNil(o.PublicIp) {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Peers) GetPublicIpOk() (*string, bool) {
	if o == nil || isNil(o.PublicIp) {
    return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *InlineResponse200124Peers) HasPublicIp() bool {
	if o != nil && !isNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *InlineResponse200124Peers) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *InlineResponse200124Peers) GetRemoteId() string {
	if o == nil || isNil(o.RemoteId) {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Peers) GetRemoteIdOk() (*string, bool) {
	if o == nil || isNil(o.RemoteId) {
    return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *InlineResponse200124Peers) HasRemoteId() bool {
	if o != nil && !isNil(o.RemoteId) {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *InlineResponse200124Peers) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetLocalId returns the LocalId field value if set, zero value otherwise.
func (o *InlineResponse200124Peers) GetLocalId() string {
	if o == nil || isNil(o.LocalId) {
		var ret string
		return ret
	}
	return *o.LocalId
}

// GetLocalIdOk returns a tuple with the LocalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Peers) GetLocalIdOk() (*string, bool) {
	if o == nil || isNil(o.LocalId) {
    return nil, false
	}
	return o.LocalId, true
}

// HasLocalId returns a boolean if a field has been set.
func (o *InlineResponse200124Peers) HasLocalId() bool {
	if o != nil && !isNil(o.LocalId) {
		return true
	}

	return false
}

// SetLocalId gets a reference to the given string and assigns it to the LocalId field.
func (o *InlineResponse200124Peers) SetLocalId(v string) {
	o.LocalId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *InlineResponse200124Peers) GetSecret() string {
	if o == nil || isNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Peers) GetSecretOk() (*string, bool) {
	if o == nil || isNil(o.Secret) {
    return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *InlineResponse200124Peers) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *InlineResponse200124Peers) SetSecret(v string) {
	o.Secret = &v
}

// GetPrivateSubnets returns the PrivateSubnets field value if set, zero value otherwise.
func (o *InlineResponse200124Peers) GetPrivateSubnets() []string {
	if o == nil || isNil(o.PrivateSubnets) {
		var ret []string
		return ret
	}
	return o.PrivateSubnets
}

// GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Peers) GetPrivateSubnetsOk() ([]string, bool) {
	if o == nil || isNil(o.PrivateSubnets) {
    return nil, false
	}
	return o.PrivateSubnets, true
}

// HasPrivateSubnets returns a boolean if a field has been set.
func (o *InlineResponse200124Peers) HasPrivateSubnets() bool {
	if o != nil && !isNil(o.PrivateSubnets) {
		return true
	}

	return false
}

// SetPrivateSubnets gets a reference to the given []string and assigns it to the PrivateSubnets field.
func (o *InlineResponse200124Peers) SetPrivateSubnets(v []string) {
	o.PrivateSubnets = v
}

// GetIpsecPolicies returns the IpsecPolicies field value if set, zero value otherwise.
func (o *InlineResponse200124Peers) GetIpsecPolicies() InlineResponse200124IpsecPolicies {
	if o == nil || isNil(o.IpsecPolicies) {
		var ret InlineResponse200124IpsecPolicies
		return ret
	}
	return *o.IpsecPolicies
}

// GetIpsecPoliciesOk returns a tuple with the IpsecPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Peers) GetIpsecPoliciesOk() (*InlineResponse200124IpsecPolicies, bool) {
	if o == nil || isNil(o.IpsecPolicies) {
    return nil, false
	}
	return o.IpsecPolicies, true
}

// HasIpsecPolicies returns a boolean if a field has been set.
func (o *InlineResponse200124Peers) HasIpsecPolicies() bool {
	if o != nil && !isNil(o.IpsecPolicies) {
		return true
	}

	return false
}

// SetIpsecPolicies gets a reference to the given InlineResponse200124IpsecPolicies and assigns it to the IpsecPolicies field.
func (o *InlineResponse200124Peers) SetIpsecPolicies(v InlineResponse200124IpsecPolicies) {
	o.IpsecPolicies = &v
}

// GetIpsecPoliciesPreset returns the IpsecPoliciesPreset field value if set, zero value otherwise.
func (o *InlineResponse200124Peers) GetIpsecPoliciesPreset() string {
	if o == nil || isNil(o.IpsecPoliciesPreset) {
		var ret string
		return ret
	}
	return *o.IpsecPoliciesPreset
}

// GetIpsecPoliciesPresetOk returns a tuple with the IpsecPoliciesPreset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Peers) GetIpsecPoliciesPresetOk() (*string, bool) {
	if o == nil || isNil(o.IpsecPoliciesPreset) {
    return nil, false
	}
	return o.IpsecPoliciesPreset, true
}

// HasIpsecPoliciesPreset returns a boolean if a field has been set.
func (o *InlineResponse200124Peers) HasIpsecPoliciesPreset() bool {
	if o != nil && !isNil(o.IpsecPoliciesPreset) {
		return true
	}

	return false
}

// SetIpsecPoliciesPreset gets a reference to the given string and assigns it to the IpsecPoliciesPreset field.
func (o *InlineResponse200124Peers) SetIpsecPoliciesPreset(v string) {
	o.IpsecPoliciesPreset = &v
}

// GetIkeVersion returns the IkeVersion field value if set, zero value otherwise.
func (o *InlineResponse200124Peers) GetIkeVersion() string {
	if o == nil || isNil(o.IkeVersion) {
		var ret string
		return ret
	}
	return *o.IkeVersion
}

// GetIkeVersionOk returns a tuple with the IkeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Peers) GetIkeVersionOk() (*string, bool) {
	if o == nil || isNil(o.IkeVersion) {
    return nil, false
	}
	return o.IkeVersion, true
}

// HasIkeVersion returns a boolean if a field has been set.
func (o *InlineResponse200124Peers) HasIkeVersion() bool {
	if o != nil && !isNil(o.IkeVersion) {
		return true
	}

	return false
}

// SetIkeVersion gets a reference to the given string and assigns it to the IkeVersion field.
func (o *InlineResponse200124Peers) SetIkeVersion(v string) {
	o.IkeVersion = &v
}

// GetNetworkTags returns the NetworkTags field value if set, zero value otherwise.
func (o *InlineResponse200124Peers) GetNetworkTags() []string {
	if o == nil || isNil(o.NetworkTags) {
		var ret []string
		return ret
	}
	return o.NetworkTags
}

// GetNetworkTagsOk returns a tuple with the NetworkTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Peers) GetNetworkTagsOk() ([]string, bool) {
	if o == nil || isNil(o.NetworkTags) {
    return nil, false
	}
	return o.NetworkTags, true
}

// HasNetworkTags returns a boolean if a field has been set.
func (o *InlineResponse200124Peers) HasNetworkTags() bool {
	if o != nil && !isNil(o.NetworkTags) {
		return true
	}

	return false
}

// SetNetworkTags gets a reference to the given []string and assigns it to the NetworkTags field.
func (o *InlineResponse200124Peers) SetNetworkTags(v []string) {
	o.NetworkTags = v
}

func (o InlineResponse200124Peers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.PublicIp) {
		toSerialize["publicIp"] = o.PublicIp
	}
	if !isNil(o.RemoteId) {
		toSerialize["remoteId"] = o.RemoteId
	}
	if !isNil(o.LocalId) {
		toSerialize["localId"] = o.LocalId
	}
	if !isNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !isNil(o.PrivateSubnets) {
		toSerialize["privateSubnets"] = o.PrivateSubnets
	}
	if !isNil(o.IpsecPolicies) {
		toSerialize["ipsecPolicies"] = o.IpsecPolicies
	}
	if !isNil(o.IpsecPoliciesPreset) {
		toSerialize["ipsecPoliciesPreset"] = o.IpsecPoliciesPreset
	}
	if !isNil(o.IkeVersion) {
		toSerialize["ikeVersion"] = o.IkeVersion
	}
	if !isNil(o.NetworkTags) {
		toSerialize["networkTags"] = o.NetworkTags
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200124Peers struct {
	value *InlineResponse200124Peers
	isSet bool
}

func (v NullableInlineResponse200124Peers) Get() *InlineResponse200124Peers {
	return v.value
}

func (v *NullableInlineResponse200124Peers) Set(val *InlineResponse200124Peers) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200124Peers) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200124Peers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200124Peers(val *InlineResponse200124Peers) *NullableInlineResponse200124Peers {
	return &NullableInlineResponse200124Peers{value: val, isSet: true}
}

func (v NullableInlineResponse200124Peers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200124Peers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

