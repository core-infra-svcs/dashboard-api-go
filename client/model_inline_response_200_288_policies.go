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

// InlineResponse200288Policies struct for InlineResponse200288Policies
type InlineResponse200288Policies struct {
	// The Id of the Sentry Policy
	PolicyId *string `json:"policyId,omitempty"`
	// The Id of the Network the Sentry Policy is associated with. In a locale, this should be the Wireless Group if present, otherwise the Wired Group.
	NetworkId *string `json:"networkId,omitempty"`
	// The Id of the Systems Manager Network the Sentry Policy is assigned to
	SmNetworkId *string `json:"smNetworkId,omitempty"`
	// The tags of the Sentry Policy
	Tags []string `json:"tags,omitempty"`
	// The scope of the Sentry Policy
	Scope *string `json:"scope,omitempty"`
	// The number of the Group Policy
	GroupNumber *string `json:"groupNumber,omitempty"`
	// The Id of the Group Policy. This is associated with the network specified by the networkId.
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
	// The priority of the Sentry Policy
	Priority *string `json:"priority,omitempty"`
	// The creation time of the Sentry Policy
	CreatedAt *string `json:"createdAt,omitempty"`
	// The last update time of the Sentry Policy
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty"`
}

// NewInlineResponse200288Policies instantiates a new InlineResponse200288Policies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200288Policies() *InlineResponse200288Policies {
	this := InlineResponse200288Policies{}
	return &this
}

// NewInlineResponse200288PoliciesWithDefaults instantiates a new InlineResponse200288Policies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200288PoliciesWithDefaults() *InlineResponse200288Policies {
	this := InlineResponse200288Policies{}
	return &this
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *InlineResponse200288Policies) GetPolicyId() string {
	if o == nil || isNil(o.PolicyId) {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200288Policies) GetPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.PolicyId) {
    return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *InlineResponse200288Policies) HasPolicyId() bool {
	if o != nil && !isNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *InlineResponse200288Policies) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200288Policies) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200288Policies) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200288Policies) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200288Policies) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSmNetworkId returns the SmNetworkId field value if set, zero value otherwise.
func (o *InlineResponse200288Policies) GetSmNetworkId() string {
	if o == nil || isNil(o.SmNetworkId) {
		var ret string
		return ret
	}
	return *o.SmNetworkId
}

// GetSmNetworkIdOk returns a tuple with the SmNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200288Policies) GetSmNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.SmNetworkId) {
    return nil, false
	}
	return o.SmNetworkId, true
}

// HasSmNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200288Policies) HasSmNetworkId() bool {
	if o != nil && !isNil(o.SmNetworkId) {
		return true
	}

	return false
}

// SetSmNetworkId gets a reference to the given string and assigns it to the SmNetworkId field.
func (o *InlineResponse200288Policies) SetSmNetworkId(v string) {
	o.SmNetworkId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200288Policies) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200288Policies) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200288Policies) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse200288Policies) SetTags(v []string) {
	o.Tags = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineResponse200288Policies) GetScope() string {
	if o == nil || isNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200288Policies) GetScopeOk() (*string, bool) {
	if o == nil || isNil(o.Scope) {
    return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineResponse200288Policies) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *InlineResponse200288Policies) SetScope(v string) {
	o.Scope = &v
}

// GetGroupNumber returns the GroupNumber field value if set, zero value otherwise.
func (o *InlineResponse200288Policies) GetGroupNumber() string {
	if o == nil || isNil(o.GroupNumber) {
		var ret string
		return ret
	}
	return *o.GroupNumber
}

// GetGroupNumberOk returns a tuple with the GroupNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200288Policies) GetGroupNumberOk() (*string, bool) {
	if o == nil || isNil(o.GroupNumber) {
    return nil, false
	}
	return o.GroupNumber, true
}

// HasGroupNumber returns a boolean if a field has been set.
func (o *InlineResponse200288Policies) HasGroupNumber() bool {
	if o != nil && !isNil(o.GroupNumber) {
		return true
	}

	return false
}

// SetGroupNumber gets a reference to the given string and assigns it to the GroupNumber field.
func (o *InlineResponse200288Policies) SetGroupNumber(v string) {
	o.GroupNumber = &v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *InlineResponse200288Policies) GetGroupPolicyId() string {
	if o == nil || isNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200288Policies) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupPolicyId) {
    return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *InlineResponse200288Policies) HasGroupPolicyId() bool {
	if o != nil && !isNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *InlineResponse200288Policies) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *InlineResponse200288Policies) GetPriority() string {
	if o == nil || isNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200288Policies) GetPriorityOk() (*string, bool) {
	if o == nil || isNil(o.Priority) {
    return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *InlineResponse200288Policies) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *InlineResponse200288Policies) SetPriority(v string) {
	o.Priority = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse200288Policies) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200288Policies) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse200288Policies) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *InlineResponse200288Policies) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *InlineResponse200288Policies) GetLastUpdatedAt() string {
	if o == nil || isNil(o.LastUpdatedAt) {
		var ret string
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200288Policies) GetLastUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.LastUpdatedAt) {
    return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *InlineResponse200288Policies) HasLastUpdatedAt() bool {
	if o != nil && !isNil(o.LastUpdatedAt) {
		return true
	}

	return false
}

// SetLastUpdatedAt gets a reference to the given string and assigns it to the LastUpdatedAt field.
func (o *InlineResponse200288Policies) SetLastUpdatedAt(v string) {
	o.LastUpdatedAt = &v
}

func (o InlineResponse200288Policies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PolicyId) {
		toSerialize["policyId"] = o.PolicyId
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.SmNetworkId) {
		toSerialize["smNetworkId"] = o.SmNetworkId
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !isNil(o.GroupNumber) {
		toSerialize["groupNumber"] = o.GroupNumber
	}
	if !isNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.LastUpdatedAt) {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200288Policies struct {
	value *InlineResponse200288Policies
	isSet bool
}

func (v NullableInlineResponse200288Policies) Get() *InlineResponse200288Policies {
	return v.value
}

func (v *NullableInlineResponse200288Policies) Set(val *InlineResponse200288Policies) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200288Policies) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200288Policies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200288Policies(val *InlineResponse200288Policies) *NullableInlineResponse200288Policies {
	return &NullableInlineResponse200288Policies{value: val, isSet: true}
}

func (v NullableInlineResponse200288Policies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200288Policies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

