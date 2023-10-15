/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200142 struct for InlineResponse200142
type InlineResponse200142 struct {
	// Short name of the early access feature
	ShortName *string `json:"shortName,omitempty"`
	// Name of the early access feature
	Name *string `json:"name,omitempty"`
	Descriptions *OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions `json:"descriptions,omitempty"`
	// Topic of the early access feature
	Topic *string `json:"topic,omitempty"`
	// If this early access feature can only be opted in for the entire organization
	IsOrgScopedOnly *bool `json:"isOrgScopedOnly,omitempty"`
	// Link to the documentation of this early access feature
	DocumentationLink *string `json:"documentationLink,omitempty"`
	// Link to get support for this early access feature
	SupportLink *string `json:"supportLink,omitempty"`
}

// NewInlineResponse200142 instantiates a new InlineResponse200142 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200142() *InlineResponse200142 {
	this := InlineResponse200142{}
	return &this
}

// NewInlineResponse200142WithDefaults instantiates a new InlineResponse200142 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200142WithDefaults() *InlineResponse200142 {
	this := InlineResponse200142{}
	return &this
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *InlineResponse200142) GetShortName() string {
	if o == nil || isNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200142) GetShortNameOk() (*string, bool) {
	if o == nil || isNil(o.ShortName) {
    return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *InlineResponse200142) HasShortName() bool {
	if o != nil && !isNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *InlineResponse200142) SetShortName(v string) {
	o.ShortName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200142) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200142) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200142) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200142) SetName(v string) {
	o.Name = &v
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *InlineResponse200142) GetDescriptions() OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions {
	if o == nil || isNil(o.Descriptions) {
		var ret OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions
		return ret
	}
	return *o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200142) GetDescriptionsOk() (*OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions, bool) {
	if o == nil || isNil(o.Descriptions) {
    return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *InlineResponse200142) HasDescriptions() bool {
	if o != nil && !isNil(o.Descriptions) {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions and assigns it to the Descriptions field.
func (o *InlineResponse200142) SetDescriptions(v OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) {
	o.Descriptions = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *InlineResponse200142) GetTopic() string {
	if o == nil || isNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200142) GetTopicOk() (*string, bool) {
	if o == nil || isNil(o.Topic) {
    return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *InlineResponse200142) HasTopic() bool {
	if o != nil && !isNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *InlineResponse200142) SetTopic(v string) {
	o.Topic = &v
}

// GetIsOrgScopedOnly returns the IsOrgScopedOnly field value if set, zero value otherwise.
func (o *InlineResponse200142) GetIsOrgScopedOnly() bool {
	if o == nil || isNil(o.IsOrgScopedOnly) {
		var ret bool
		return ret
	}
	return *o.IsOrgScopedOnly
}

// GetIsOrgScopedOnlyOk returns a tuple with the IsOrgScopedOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200142) GetIsOrgScopedOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.IsOrgScopedOnly) {
    return nil, false
	}
	return o.IsOrgScopedOnly, true
}

// HasIsOrgScopedOnly returns a boolean if a field has been set.
func (o *InlineResponse200142) HasIsOrgScopedOnly() bool {
	if o != nil && !isNil(o.IsOrgScopedOnly) {
		return true
	}

	return false
}

// SetIsOrgScopedOnly gets a reference to the given bool and assigns it to the IsOrgScopedOnly field.
func (o *InlineResponse200142) SetIsOrgScopedOnly(v bool) {
	o.IsOrgScopedOnly = &v
}

// GetDocumentationLink returns the DocumentationLink field value if set, zero value otherwise.
func (o *InlineResponse200142) GetDocumentationLink() string {
	if o == nil || isNil(o.DocumentationLink) {
		var ret string
		return ret
	}
	return *o.DocumentationLink
}

// GetDocumentationLinkOk returns a tuple with the DocumentationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200142) GetDocumentationLinkOk() (*string, bool) {
	if o == nil || isNil(o.DocumentationLink) {
    return nil, false
	}
	return o.DocumentationLink, true
}

// HasDocumentationLink returns a boolean if a field has been set.
func (o *InlineResponse200142) HasDocumentationLink() bool {
	if o != nil && !isNil(o.DocumentationLink) {
		return true
	}

	return false
}

// SetDocumentationLink gets a reference to the given string and assigns it to the DocumentationLink field.
func (o *InlineResponse200142) SetDocumentationLink(v string) {
	o.DocumentationLink = &v
}

// GetSupportLink returns the SupportLink field value if set, zero value otherwise.
func (o *InlineResponse200142) GetSupportLink() string {
	if o == nil || isNil(o.SupportLink) {
		var ret string
		return ret
	}
	return *o.SupportLink
}

// GetSupportLinkOk returns a tuple with the SupportLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200142) GetSupportLinkOk() (*string, bool) {
	if o == nil || isNil(o.SupportLink) {
    return nil, false
	}
	return o.SupportLink, true
}

// HasSupportLink returns a boolean if a field has been set.
func (o *InlineResponse200142) HasSupportLink() bool {
	if o != nil && !isNil(o.SupportLink) {
		return true
	}

	return false
}

// SetSupportLink gets a reference to the given string and assigns it to the SupportLink field.
func (o *InlineResponse200142) SetSupportLink(v string) {
	o.SupportLink = &v
}

func (o InlineResponse200142) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Descriptions) {
		toSerialize["descriptions"] = o.Descriptions
	}
	if !isNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	if !isNil(o.IsOrgScopedOnly) {
		toSerialize["isOrgScopedOnly"] = o.IsOrgScopedOnly
	}
	if !isNil(o.DocumentationLink) {
		toSerialize["documentationLink"] = o.DocumentationLink
	}
	if !isNil(o.SupportLink) {
		toSerialize["supportLink"] = o.SupportLink
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200142 struct {
	value *InlineResponse200142
	isSet bool
}

func (v NullableInlineResponse200142) Get() *InlineResponse200142 {
	return v.value
}

func (v *NullableInlineResponse200142) Set(val *InlineResponse200142) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200142) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200142) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200142(val *InlineResponse200142) *NullableInlineResponse200142 {
	return &NullableInlineResponse200142{value: val, isSet: true}
}

func (v NullableInlineResponse200142) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200142) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


