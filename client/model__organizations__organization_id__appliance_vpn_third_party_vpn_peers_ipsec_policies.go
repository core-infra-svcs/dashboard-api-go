/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies Custom IPSec policies for the VPN peer. If not included and a preset has not been chosen, the default preset for IPSec policies will be used.
type OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies struct {
	// This is the cipher algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: 'aes256', 'aes192', 'aes128', 'tripledes', 'des'
	IkeCipherAlgo *[]string `json:"ikeCipherAlgo,omitempty"`
	// This is the authentication algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: 'sha256', 'sha1', 'md5'
	IkeAuthAlgo *[]string `json:"ikeAuthAlgo,omitempty"`
	// [optional] This is the pseudo-random function to be used in IKE_SA. The value should be an array with one of the following algorithms: 'prfsha256', 'prfsha1', 'prfmd5', 'default'. The 'default' option can be used to default to the Authentication algorithm.
	IkePrfAlgo *[]string `json:"ikePrfAlgo,omitempty"`
	// This is the Diffie-Hellman group to be used in Phase 1. The value should be an array with one of the following algorithms: 'group14', 'group5', 'group2', 'group1'
	IkeDiffieHellmanGroup *[]string `json:"ikeDiffieHellmanGroup,omitempty"`
	// The lifetime of the Phase 1 SA in seconds.
	IkeLifetime *int32 `json:"ikeLifetime,omitempty"`
	// This is the cipher algorithms to be used in Phase 2. The value should be an array with one or more of the following algorithms: 'aes256', 'aes192', 'aes128', 'tripledes', 'des', 'null'
	ChildCipherAlgo *[]string `json:"childCipherAlgo,omitempty"`
	// This is the authentication algorithms to be used in Phase 2. The value should be an array with one of the following algorithms: 'sha256', 'sha1', 'md5'
	ChildAuthAlgo *[]string `json:"childAuthAlgo,omitempty"`
	// This is the Diffie-Hellman group to be used for Perfect Forward Secrecy in Phase 2. The value should be an array with one of the following values: 'disabled','group14', 'group5', 'group2', 'group1'
	ChildPfsGroup *[]string `json:"childPfsGroup,omitempty"`
	// The lifetime of the Phase 2 SA in seconds.
	ChildLifetime *int32 `json:"childLifetime,omitempty"`
}

// NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies instantiates a new OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies() *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies {
	this := OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies{}
	return &this
}

// NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPoliciesWithDefaults instantiates a new OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPoliciesWithDefaults() *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies {
	this := OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies{}
	return &this
}

// GetIkeCipherAlgo returns the IkeCipherAlgo field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetIkeCipherAlgo() []string {
	if o == nil || o.IkeCipherAlgo == nil {
		var ret []string
		return ret
	}
	return *o.IkeCipherAlgo
}

// GetIkeCipherAlgoOk returns a tuple with the IkeCipherAlgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetIkeCipherAlgoOk() (*[]string, bool) {
	if o == nil || o.IkeCipherAlgo == nil {
		return nil, false
	}
	return o.IkeCipherAlgo, true
}

// HasIkeCipherAlgo returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) HasIkeCipherAlgo() bool {
	if o != nil && o.IkeCipherAlgo != nil {
		return true
	}

	return false
}

// SetIkeCipherAlgo gets a reference to the given []string and assigns it to the IkeCipherAlgo field.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) SetIkeCipherAlgo(v []string) {
	o.IkeCipherAlgo = &v
}

// GetIkeAuthAlgo returns the IkeAuthAlgo field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetIkeAuthAlgo() []string {
	if o == nil || o.IkeAuthAlgo == nil {
		var ret []string
		return ret
	}
	return *o.IkeAuthAlgo
}

// GetIkeAuthAlgoOk returns a tuple with the IkeAuthAlgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetIkeAuthAlgoOk() (*[]string, bool) {
	if o == nil || o.IkeAuthAlgo == nil {
		return nil, false
	}
	return o.IkeAuthAlgo, true
}

// HasIkeAuthAlgo returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) HasIkeAuthAlgo() bool {
	if o != nil && o.IkeAuthAlgo != nil {
		return true
	}

	return false
}

// SetIkeAuthAlgo gets a reference to the given []string and assigns it to the IkeAuthAlgo field.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) SetIkeAuthAlgo(v []string) {
	o.IkeAuthAlgo = &v
}

// GetIkePrfAlgo returns the IkePrfAlgo field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetIkePrfAlgo() []string {
	if o == nil || o.IkePrfAlgo == nil {
		var ret []string
		return ret
	}
	return *o.IkePrfAlgo
}

// GetIkePrfAlgoOk returns a tuple with the IkePrfAlgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetIkePrfAlgoOk() (*[]string, bool) {
	if o == nil || o.IkePrfAlgo == nil {
		return nil, false
	}
	return o.IkePrfAlgo, true
}

// HasIkePrfAlgo returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) HasIkePrfAlgo() bool {
	if o != nil && o.IkePrfAlgo != nil {
		return true
	}

	return false
}

// SetIkePrfAlgo gets a reference to the given []string and assigns it to the IkePrfAlgo field.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) SetIkePrfAlgo(v []string) {
	o.IkePrfAlgo = &v
}

// GetIkeDiffieHellmanGroup returns the IkeDiffieHellmanGroup field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetIkeDiffieHellmanGroup() []string {
	if o == nil || o.IkeDiffieHellmanGroup == nil {
		var ret []string
		return ret
	}
	return *o.IkeDiffieHellmanGroup
}

// GetIkeDiffieHellmanGroupOk returns a tuple with the IkeDiffieHellmanGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetIkeDiffieHellmanGroupOk() (*[]string, bool) {
	if o == nil || o.IkeDiffieHellmanGroup == nil {
		return nil, false
	}
	return o.IkeDiffieHellmanGroup, true
}

// HasIkeDiffieHellmanGroup returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) HasIkeDiffieHellmanGroup() bool {
	if o != nil && o.IkeDiffieHellmanGroup != nil {
		return true
	}

	return false
}

// SetIkeDiffieHellmanGroup gets a reference to the given []string and assigns it to the IkeDiffieHellmanGroup field.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) SetIkeDiffieHellmanGroup(v []string) {
	o.IkeDiffieHellmanGroup = &v
}

// GetIkeLifetime returns the IkeLifetime field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetIkeLifetime() int32 {
	if o == nil || o.IkeLifetime == nil {
		var ret int32
		return ret
	}
	return *o.IkeLifetime
}

// GetIkeLifetimeOk returns a tuple with the IkeLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetIkeLifetimeOk() (*int32, bool) {
	if o == nil || o.IkeLifetime == nil {
		return nil, false
	}
	return o.IkeLifetime, true
}

// HasIkeLifetime returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) HasIkeLifetime() bool {
	if o != nil && o.IkeLifetime != nil {
		return true
	}

	return false
}

// SetIkeLifetime gets a reference to the given int32 and assigns it to the IkeLifetime field.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) SetIkeLifetime(v int32) {
	o.IkeLifetime = &v
}

// GetChildCipherAlgo returns the ChildCipherAlgo field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetChildCipherAlgo() []string {
	if o == nil || o.ChildCipherAlgo == nil {
		var ret []string
		return ret
	}
	return *o.ChildCipherAlgo
}

// GetChildCipherAlgoOk returns a tuple with the ChildCipherAlgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetChildCipherAlgoOk() (*[]string, bool) {
	if o == nil || o.ChildCipherAlgo == nil {
		return nil, false
	}
	return o.ChildCipherAlgo, true
}

// HasChildCipherAlgo returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) HasChildCipherAlgo() bool {
	if o != nil && o.ChildCipherAlgo != nil {
		return true
	}

	return false
}

// SetChildCipherAlgo gets a reference to the given []string and assigns it to the ChildCipherAlgo field.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) SetChildCipherAlgo(v []string) {
	o.ChildCipherAlgo = &v
}

// GetChildAuthAlgo returns the ChildAuthAlgo field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetChildAuthAlgo() []string {
	if o == nil || o.ChildAuthAlgo == nil {
		var ret []string
		return ret
	}
	return *o.ChildAuthAlgo
}

// GetChildAuthAlgoOk returns a tuple with the ChildAuthAlgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetChildAuthAlgoOk() (*[]string, bool) {
	if o == nil || o.ChildAuthAlgo == nil {
		return nil, false
	}
	return o.ChildAuthAlgo, true
}

// HasChildAuthAlgo returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) HasChildAuthAlgo() bool {
	if o != nil && o.ChildAuthAlgo != nil {
		return true
	}

	return false
}

// SetChildAuthAlgo gets a reference to the given []string and assigns it to the ChildAuthAlgo field.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) SetChildAuthAlgo(v []string) {
	o.ChildAuthAlgo = &v
}

// GetChildPfsGroup returns the ChildPfsGroup field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetChildPfsGroup() []string {
	if o == nil || o.ChildPfsGroup == nil {
		var ret []string
		return ret
	}
	return *o.ChildPfsGroup
}

// GetChildPfsGroupOk returns a tuple with the ChildPfsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetChildPfsGroupOk() (*[]string, bool) {
	if o == nil || o.ChildPfsGroup == nil {
		return nil, false
	}
	return o.ChildPfsGroup, true
}

// HasChildPfsGroup returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) HasChildPfsGroup() bool {
	if o != nil && o.ChildPfsGroup != nil {
		return true
	}

	return false
}

// SetChildPfsGroup gets a reference to the given []string and assigns it to the ChildPfsGroup field.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) SetChildPfsGroup(v []string) {
	o.ChildPfsGroup = &v
}

// GetChildLifetime returns the ChildLifetime field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetChildLifetime() int32 {
	if o == nil || o.ChildLifetime == nil {
		var ret int32
		return ret
	}
	return *o.ChildLifetime
}

// GetChildLifetimeOk returns a tuple with the ChildLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) GetChildLifetimeOk() (*int32, bool) {
	if o == nil || o.ChildLifetime == nil {
		return nil, false
	}
	return o.ChildLifetime, true
}

// HasChildLifetime returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) HasChildLifetime() bool {
	if o != nil && o.ChildLifetime != nil {
		return true
	}

	return false
}

// SetChildLifetime gets a reference to the given int32 and assigns it to the ChildLifetime field.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) SetChildLifetime(v int32) {
	o.ChildLifetime = &v
}

func (o OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IkeCipherAlgo != nil {
		toSerialize["ikeCipherAlgo"] = o.IkeCipherAlgo
	}
	if o.IkeAuthAlgo != nil {
		toSerialize["ikeAuthAlgo"] = o.IkeAuthAlgo
	}
	if o.IkePrfAlgo != nil {
		toSerialize["ikePrfAlgo"] = o.IkePrfAlgo
	}
	if o.IkeDiffieHellmanGroup != nil {
		toSerialize["ikeDiffieHellmanGroup"] = o.IkeDiffieHellmanGroup
	}
	if o.IkeLifetime != nil {
		toSerialize["ikeLifetime"] = o.IkeLifetime
	}
	if o.ChildCipherAlgo != nil {
		toSerialize["childCipherAlgo"] = o.ChildCipherAlgo
	}
	if o.ChildAuthAlgo != nil {
		toSerialize["childAuthAlgo"] = o.ChildAuthAlgo
	}
	if o.ChildPfsGroup != nil {
		toSerialize["childPfsGroup"] = o.ChildPfsGroup
	}
	if o.ChildLifetime != nil {
		toSerialize["childLifetime"] = o.ChildLifetime
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies struct {
	value *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies
	isSet bool
}

func (v NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) Get() *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) Set(val *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies(val *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) *NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies {
	return &NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersIpsecPolicies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


