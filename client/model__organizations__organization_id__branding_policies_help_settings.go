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

// OrganizationsOrganizationIdBrandingPoliciesHelpSettings     Settings for describing the modifications to various Help page features. Each property in this object accepts one of     'default or inherit' (do not modify functionality), 'hide' (remove the section from Dashboard), or 'show' (always show     the section on Dashboard). Some properties in this object also accept custom HTML used to replace the section on     Dashboard; see the documentation for each property to see the allowed values.  Each property defaults to 'default or inherit' when not provided.
type OrganizationsOrganizationIdBrandingPoliciesHelpSettings struct {
	//     The Help tab, under which all support information resides. If this tab is hidden, no other 'Help' branding     customizations will be visible. Can be one of 'default or inherit', 'hide' or 'show'. 
	HelpTab *string `json:"helpTab,omitempty"`
	//     The 'Help -> Get Help' subtab on which Cisco Meraki KB, Product Manuals, and Support/Case Information are displayed. Note     that if this subtab is hidden, branding customizations for the KB on 'Get help', Cisco Meraki product documentation,     and support contact info will not be visible. Can be one of 'default or inherit', 'hide' or 'show'. 
	GetHelpSubtab *string `json:"getHelpSubtab,omitempty"`
	//     The 'Help -> Community' subtab which provides a link to Meraki Community. Can be one of 'default or inherit', 'hide' or 'show'. 
	CommunitySubtab *string `json:"communitySubtab,omitempty"`
	//     The 'Help -> Cases' Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one     of 'default or inherit', 'hide' or 'show'. 
	CasesSubtab *string `json:"casesSubtab,omitempty"`
	//     The 'Help -> Data protection requests' Dashboard subtab on which requests to delete, restrict, or export end-user data can     be audited. Can be one of 'default or inherit', 'hide' or 'show'. 
	DataProtectionRequestsSubtab *string `json:"dataProtectionRequestsSubtab,omitempty"`
	//     The KB search box which appears on the Help page. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string. 
	GetHelpSubtabKnowledgeBaseSearch *string `json:"getHelpSubtabKnowledgeBaseSearch,omitempty"`
	//     The universal search box always visible on Dashboard will, by default, present results from the Meraki KB. This configures     whether these Meraki KB results should be returned. Can be one of 'default or inherit', 'hide' or 'show'. 
	UniversalSearchKnowledgeBaseSearch *string `json:"universalSearchKnowledgeBaseSearch,omitempty"`
	//     The 'Product Manuals' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string. 
	CiscoMerakiProductDocumentation *string `json:"ciscoMerakiProductDocumentation,omitempty"`
	//     The 'Contact Meraki Support' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string. 
	SupportContactInfo *string `json:"supportContactInfo,omitempty"`
	//     The 'Help -> New features' subtab where new Dashboard features are detailed. Can be one of 'default or inherit', 'hide' or 'show'. 
	NewFeaturesSubtab *string `json:"newFeaturesSubtab,omitempty"`
	//     The 'Help -> Firewall info' subtab where necessary upstream firewall rules for communication to the Cisco Meraki cloud are     listed. Can be one of 'default or inherit', 'hide' or 'show'. 
	FirewallInfoSubtab *string `json:"firewallInfoSubtab,omitempty"`
	//     The 'Help -> API docs' subtab where a detailed description of the Dashboard API is listed. Can be one of     'default or inherit', 'hide' or 'show'. 
	ApiDocsSubtab *string `json:"apiDocsSubtab,omitempty"`
	//     The 'Help -> Replacement info' subtab where important information regarding device replacements is detailed. Can be one of     'default or inherit', 'hide' or 'show'. 
	HardwareReplacementsSubtab *string `json:"hardwareReplacementsSubtab,omitempty"`
	//     The 'SM Forums' subtab which links to community-based support for Cisco Meraki Systems Manager. Only configurable for     organizations that contain Systems Manager networks. Can be one of 'default or inherit', 'hide' or 'show'. 
	SmForums *string `json:"smForums,omitempty"`
}

// NewOrganizationsOrganizationIdBrandingPoliciesHelpSettings instantiates a new OrganizationsOrganizationIdBrandingPoliciesHelpSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdBrandingPoliciesHelpSettings() *OrganizationsOrganizationIdBrandingPoliciesHelpSettings {
	this := OrganizationsOrganizationIdBrandingPoliciesHelpSettings{}
	return &this
}

// NewOrganizationsOrganizationIdBrandingPoliciesHelpSettingsWithDefaults instantiates a new OrganizationsOrganizationIdBrandingPoliciesHelpSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdBrandingPoliciesHelpSettingsWithDefaults() *OrganizationsOrganizationIdBrandingPoliciesHelpSettings {
	this := OrganizationsOrganizationIdBrandingPoliciesHelpSettings{}
	return &this
}

// GetHelpTab returns the HelpTab field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetHelpTab() string {
	if o == nil || o.HelpTab == nil {
		var ret string
		return ret
	}
	return *o.HelpTab
}

// GetHelpTabOk returns a tuple with the HelpTab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetHelpTabOk() (*string, bool) {
	if o == nil || o.HelpTab == nil {
		return nil, false
	}
	return o.HelpTab, true
}

// HasHelpTab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasHelpTab() bool {
	if o != nil && o.HelpTab != nil {
		return true
	}

	return false
}

// SetHelpTab gets a reference to the given string and assigns it to the HelpTab field.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetHelpTab(v string) {
	o.HelpTab = &v
}

// GetGetHelpSubtab returns the GetHelpSubtab field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetGetHelpSubtab() string {
	if o == nil || o.GetHelpSubtab == nil {
		var ret string
		return ret
	}
	return *o.GetHelpSubtab
}

// GetGetHelpSubtabOk returns a tuple with the GetHelpSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetGetHelpSubtabOk() (*string, bool) {
	if o == nil || o.GetHelpSubtab == nil {
		return nil, false
	}
	return o.GetHelpSubtab, true
}

// HasGetHelpSubtab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasGetHelpSubtab() bool {
	if o != nil && o.GetHelpSubtab != nil {
		return true
	}

	return false
}

// SetGetHelpSubtab gets a reference to the given string and assigns it to the GetHelpSubtab field.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetGetHelpSubtab(v string) {
	o.GetHelpSubtab = &v
}

// GetCommunitySubtab returns the CommunitySubtab field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetCommunitySubtab() string {
	if o == nil || o.CommunitySubtab == nil {
		var ret string
		return ret
	}
	return *o.CommunitySubtab
}

// GetCommunitySubtabOk returns a tuple with the CommunitySubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetCommunitySubtabOk() (*string, bool) {
	if o == nil || o.CommunitySubtab == nil {
		return nil, false
	}
	return o.CommunitySubtab, true
}

// HasCommunitySubtab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasCommunitySubtab() bool {
	if o != nil && o.CommunitySubtab != nil {
		return true
	}

	return false
}

// SetCommunitySubtab gets a reference to the given string and assigns it to the CommunitySubtab field.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetCommunitySubtab(v string) {
	o.CommunitySubtab = &v
}

// GetCasesSubtab returns the CasesSubtab field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetCasesSubtab() string {
	if o == nil || o.CasesSubtab == nil {
		var ret string
		return ret
	}
	return *o.CasesSubtab
}

// GetCasesSubtabOk returns a tuple with the CasesSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetCasesSubtabOk() (*string, bool) {
	if o == nil || o.CasesSubtab == nil {
		return nil, false
	}
	return o.CasesSubtab, true
}

// HasCasesSubtab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasCasesSubtab() bool {
	if o != nil && o.CasesSubtab != nil {
		return true
	}

	return false
}

// SetCasesSubtab gets a reference to the given string and assigns it to the CasesSubtab field.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetCasesSubtab(v string) {
	o.CasesSubtab = &v
}

// GetDataProtectionRequestsSubtab returns the DataProtectionRequestsSubtab field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetDataProtectionRequestsSubtab() string {
	if o == nil || o.DataProtectionRequestsSubtab == nil {
		var ret string
		return ret
	}
	return *o.DataProtectionRequestsSubtab
}

// GetDataProtectionRequestsSubtabOk returns a tuple with the DataProtectionRequestsSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetDataProtectionRequestsSubtabOk() (*string, bool) {
	if o == nil || o.DataProtectionRequestsSubtab == nil {
		return nil, false
	}
	return o.DataProtectionRequestsSubtab, true
}

// HasDataProtectionRequestsSubtab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasDataProtectionRequestsSubtab() bool {
	if o != nil && o.DataProtectionRequestsSubtab != nil {
		return true
	}

	return false
}

// SetDataProtectionRequestsSubtab gets a reference to the given string and assigns it to the DataProtectionRequestsSubtab field.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetDataProtectionRequestsSubtab(v string) {
	o.DataProtectionRequestsSubtab = &v
}

// GetGetHelpSubtabKnowledgeBaseSearch returns the GetHelpSubtabKnowledgeBaseSearch field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetGetHelpSubtabKnowledgeBaseSearch() string {
	if o == nil || o.GetHelpSubtabKnowledgeBaseSearch == nil {
		var ret string
		return ret
	}
	return *o.GetHelpSubtabKnowledgeBaseSearch
}

// GetGetHelpSubtabKnowledgeBaseSearchOk returns a tuple with the GetHelpSubtabKnowledgeBaseSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetGetHelpSubtabKnowledgeBaseSearchOk() (*string, bool) {
	if o == nil || o.GetHelpSubtabKnowledgeBaseSearch == nil {
		return nil, false
	}
	return o.GetHelpSubtabKnowledgeBaseSearch, true
}

// HasGetHelpSubtabKnowledgeBaseSearch returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasGetHelpSubtabKnowledgeBaseSearch() bool {
	if o != nil && o.GetHelpSubtabKnowledgeBaseSearch != nil {
		return true
	}

	return false
}

// SetGetHelpSubtabKnowledgeBaseSearch gets a reference to the given string and assigns it to the GetHelpSubtabKnowledgeBaseSearch field.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetGetHelpSubtabKnowledgeBaseSearch(v string) {
	o.GetHelpSubtabKnowledgeBaseSearch = &v
}

// GetUniversalSearchKnowledgeBaseSearch returns the UniversalSearchKnowledgeBaseSearch field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetUniversalSearchKnowledgeBaseSearch() string {
	if o == nil || o.UniversalSearchKnowledgeBaseSearch == nil {
		var ret string
		return ret
	}
	return *o.UniversalSearchKnowledgeBaseSearch
}

// GetUniversalSearchKnowledgeBaseSearchOk returns a tuple with the UniversalSearchKnowledgeBaseSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetUniversalSearchKnowledgeBaseSearchOk() (*string, bool) {
	if o == nil || o.UniversalSearchKnowledgeBaseSearch == nil {
		return nil, false
	}
	return o.UniversalSearchKnowledgeBaseSearch, true
}

// HasUniversalSearchKnowledgeBaseSearch returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasUniversalSearchKnowledgeBaseSearch() bool {
	if o != nil && o.UniversalSearchKnowledgeBaseSearch != nil {
		return true
	}

	return false
}

// SetUniversalSearchKnowledgeBaseSearch gets a reference to the given string and assigns it to the UniversalSearchKnowledgeBaseSearch field.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetUniversalSearchKnowledgeBaseSearch(v string) {
	o.UniversalSearchKnowledgeBaseSearch = &v
}

// GetCiscoMerakiProductDocumentation returns the CiscoMerakiProductDocumentation field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetCiscoMerakiProductDocumentation() string {
	if o == nil || o.CiscoMerakiProductDocumentation == nil {
		var ret string
		return ret
	}
	return *o.CiscoMerakiProductDocumentation
}

// GetCiscoMerakiProductDocumentationOk returns a tuple with the CiscoMerakiProductDocumentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetCiscoMerakiProductDocumentationOk() (*string, bool) {
	if o == nil || o.CiscoMerakiProductDocumentation == nil {
		return nil, false
	}
	return o.CiscoMerakiProductDocumentation, true
}

// HasCiscoMerakiProductDocumentation returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasCiscoMerakiProductDocumentation() bool {
	if o != nil && o.CiscoMerakiProductDocumentation != nil {
		return true
	}

	return false
}

// SetCiscoMerakiProductDocumentation gets a reference to the given string and assigns it to the CiscoMerakiProductDocumentation field.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetCiscoMerakiProductDocumentation(v string) {
	o.CiscoMerakiProductDocumentation = &v
}

// GetSupportContactInfo returns the SupportContactInfo field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetSupportContactInfo() string {
	if o == nil || o.SupportContactInfo == nil {
		var ret string
		return ret
	}
	return *o.SupportContactInfo
}

// GetSupportContactInfoOk returns a tuple with the SupportContactInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetSupportContactInfoOk() (*string, bool) {
	if o == nil || o.SupportContactInfo == nil {
		return nil, false
	}
	return o.SupportContactInfo, true
}

// HasSupportContactInfo returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasSupportContactInfo() bool {
	if o != nil && o.SupportContactInfo != nil {
		return true
	}

	return false
}

// SetSupportContactInfo gets a reference to the given string and assigns it to the SupportContactInfo field.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetSupportContactInfo(v string) {
	o.SupportContactInfo = &v
}

// GetNewFeaturesSubtab returns the NewFeaturesSubtab field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetNewFeaturesSubtab() string {
	if o == nil || o.NewFeaturesSubtab == nil {
		var ret string
		return ret
	}
	return *o.NewFeaturesSubtab
}

// GetNewFeaturesSubtabOk returns a tuple with the NewFeaturesSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetNewFeaturesSubtabOk() (*string, bool) {
	if o == nil || o.NewFeaturesSubtab == nil {
		return nil, false
	}
	return o.NewFeaturesSubtab, true
}

// HasNewFeaturesSubtab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasNewFeaturesSubtab() bool {
	if o != nil && o.NewFeaturesSubtab != nil {
		return true
	}

	return false
}

// SetNewFeaturesSubtab gets a reference to the given string and assigns it to the NewFeaturesSubtab field.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetNewFeaturesSubtab(v string) {
	o.NewFeaturesSubtab = &v
}

// GetFirewallInfoSubtab returns the FirewallInfoSubtab field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetFirewallInfoSubtab() string {
	if o == nil || o.FirewallInfoSubtab == nil {
		var ret string
		return ret
	}
	return *o.FirewallInfoSubtab
}

// GetFirewallInfoSubtabOk returns a tuple with the FirewallInfoSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetFirewallInfoSubtabOk() (*string, bool) {
	if o == nil || o.FirewallInfoSubtab == nil {
		return nil, false
	}
	return o.FirewallInfoSubtab, true
}

// HasFirewallInfoSubtab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasFirewallInfoSubtab() bool {
	if o != nil && o.FirewallInfoSubtab != nil {
		return true
	}

	return false
}

// SetFirewallInfoSubtab gets a reference to the given string and assigns it to the FirewallInfoSubtab field.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetFirewallInfoSubtab(v string) {
	o.FirewallInfoSubtab = &v
}

// GetApiDocsSubtab returns the ApiDocsSubtab field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetApiDocsSubtab() string {
	if o == nil || o.ApiDocsSubtab == nil {
		var ret string
		return ret
	}
	return *o.ApiDocsSubtab
}

// GetApiDocsSubtabOk returns a tuple with the ApiDocsSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetApiDocsSubtabOk() (*string, bool) {
	if o == nil || o.ApiDocsSubtab == nil {
		return nil, false
	}
	return o.ApiDocsSubtab, true
}

// HasApiDocsSubtab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasApiDocsSubtab() bool {
	if o != nil && o.ApiDocsSubtab != nil {
		return true
	}

	return false
}

// SetApiDocsSubtab gets a reference to the given string and assigns it to the ApiDocsSubtab field.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetApiDocsSubtab(v string) {
	o.ApiDocsSubtab = &v
}

// GetHardwareReplacementsSubtab returns the HardwareReplacementsSubtab field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetHardwareReplacementsSubtab() string {
	if o == nil || o.HardwareReplacementsSubtab == nil {
		var ret string
		return ret
	}
	return *o.HardwareReplacementsSubtab
}

// GetHardwareReplacementsSubtabOk returns a tuple with the HardwareReplacementsSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetHardwareReplacementsSubtabOk() (*string, bool) {
	if o == nil || o.HardwareReplacementsSubtab == nil {
		return nil, false
	}
	return o.HardwareReplacementsSubtab, true
}

// HasHardwareReplacementsSubtab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasHardwareReplacementsSubtab() bool {
	if o != nil && o.HardwareReplacementsSubtab != nil {
		return true
	}

	return false
}

// SetHardwareReplacementsSubtab gets a reference to the given string and assigns it to the HardwareReplacementsSubtab field.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetHardwareReplacementsSubtab(v string) {
	o.HardwareReplacementsSubtab = &v
}

// GetSmForums returns the SmForums field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetSmForums() string {
	if o == nil || o.SmForums == nil {
		var ret string
		return ret
	}
	return *o.SmForums
}

// GetSmForumsOk returns a tuple with the SmForums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetSmForumsOk() (*string, bool) {
	if o == nil || o.SmForums == nil {
		return nil, false
	}
	return o.SmForums, true
}

// HasSmForums returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasSmForums() bool {
	if o != nil && o.SmForums != nil {
		return true
	}

	return false
}

// SetSmForums gets a reference to the given string and assigns it to the SmForums field.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetSmForums(v string) {
	o.SmForums = &v
}

func (o OrganizationsOrganizationIdBrandingPoliciesHelpSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HelpTab != nil {
		toSerialize["helpTab"] = o.HelpTab
	}
	if o.GetHelpSubtab != nil {
		toSerialize["getHelpSubtab"] = o.GetHelpSubtab
	}
	if o.CommunitySubtab != nil {
		toSerialize["communitySubtab"] = o.CommunitySubtab
	}
	if o.CasesSubtab != nil {
		toSerialize["casesSubtab"] = o.CasesSubtab
	}
	if o.DataProtectionRequestsSubtab != nil {
		toSerialize["dataProtectionRequestsSubtab"] = o.DataProtectionRequestsSubtab
	}
	if o.GetHelpSubtabKnowledgeBaseSearch != nil {
		toSerialize["getHelpSubtabKnowledgeBaseSearch"] = o.GetHelpSubtabKnowledgeBaseSearch
	}
	if o.UniversalSearchKnowledgeBaseSearch != nil {
		toSerialize["universalSearchKnowledgeBaseSearch"] = o.UniversalSearchKnowledgeBaseSearch
	}
	if o.CiscoMerakiProductDocumentation != nil {
		toSerialize["ciscoMerakiProductDocumentation"] = o.CiscoMerakiProductDocumentation
	}
	if o.SupportContactInfo != nil {
		toSerialize["supportContactInfo"] = o.SupportContactInfo
	}
	if o.NewFeaturesSubtab != nil {
		toSerialize["newFeaturesSubtab"] = o.NewFeaturesSubtab
	}
	if o.FirewallInfoSubtab != nil {
		toSerialize["firewallInfoSubtab"] = o.FirewallInfoSubtab
	}
	if o.ApiDocsSubtab != nil {
		toSerialize["apiDocsSubtab"] = o.ApiDocsSubtab
	}
	if o.HardwareReplacementsSubtab != nil {
		toSerialize["hardwareReplacementsSubtab"] = o.HardwareReplacementsSubtab
	}
	if o.SmForums != nil {
		toSerialize["smForums"] = o.SmForums
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdBrandingPoliciesHelpSettings struct {
	value *OrganizationsOrganizationIdBrandingPoliciesHelpSettings
	isSet bool
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesHelpSettings) Get() *OrganizationsOrganizationIdBrandingPoliciesHelpSettings {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesHelpSettings) Set(val *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesHelpSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesHelpSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdBrandingPoliciesHelpSettings(val *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) *NullableOrganizationsOrganizationIdBrandingPoliciesHelpSettings {
	return &NullableOrganizationsOrganizationIdBrandingPoliciesHelpSettings{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesHelpSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesHelpSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


