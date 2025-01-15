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

// OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules struct for OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules
type OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules struct {
	// A rule identifier of the format meraki:intrusion/snort/GID/<gid>/SID/<sid>. gid and sid can be obtained from either https://www.snort.org/rule-docs or as ruleIds from the security events in /organization/[orgId]/securityEvents
	RuleId string `json:"ruleId"`
	// Message is optional and is ignored on a PUT call. It is allowed in order for PUT to be compatible with GET
	Message *string `json:"message,omitempty"`
}

// NewOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules instantiates a new OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules(ruleId string) *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules {
	this := OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules{}
	this.RuleId = ruleId
	return &this
}

// NewOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRulesWithDefaults instantiates a new OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRulesWithDefaults() *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules {
	this := OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules{}
	return &this
}

// GetRuleId returns the RuleId field value
func (o *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) GetRuleIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value
func (o *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) SetRuleId(v string) {
	o.RuleId = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) SetMessage(v string) {
	o.Message = &v
}

func (o OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ruleId"] = o.RuleId
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules struct {
	value *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules
	isSet bool
}

func (v NullableOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) Get() *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) Set(val *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules(val *OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) *NullableOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules {
	return &NullableOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


