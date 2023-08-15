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

// OrganizationsOrganizationIdBrandingPoliciesHelpSettings       Settings for describing the modifications to various Help page features. Each property in this object accepts one of       'default or inherit' (do not modify functionality), 'hide' (remove the section from Dashboard), or 'show' (always show       the section on Dashboard). Some properties in this object also accept custom HTML used to replace the section on       Dashboard; see the documentation for each property to see the allowed values. 
type OrganizationsOrganizationIdBrandingPoliciesHelpSettings struct {
	//       The Help tab, under which all support information resides. If this tab is hidden, no other 'Help' branding       customizations will be visible. Can be one of 'default or inherit', 'hide' or 'show'. 
	HelpTab *string `json:"helpTab,omitempty"`
	//       The 'Help -> Get Help' subtab on which Cisco Meraki KB, Product Manuals, and Support/Case Information are displayed. Note       that if this subtab is hidden, branding customizations for the KB on 'Get help', Cisco Meraki product documentation,       and support contact info will not be visible. Can be one of 'default or inherit', 'hide' or 'show'. 
	GetHelpSubtab *string `json:"getHelpSubtab,omitempty"`
	//       The 'Help -> Community' subtab which provides a link to Meraki Community. Can be one of 'default or inherit', 'hide' or 'show'. 
	CommunitySubtab *string `json:"communitySubtab,omitempty"`
	//       The 'Help -> Cases' Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one       of 'default or inherit', 'hide' or 'show'. 
	CasesSubtab *string `json:"casesSubtab,omitempty"`
	//       The 'Help -> Data protection requests' Dashboard subtab on which requests to delete, restrict, or export end-user data can       be audited. Can be one of 'default or inherit', 'hide' or 'show'. 
	DataProtectionRequestsSubtab *string `json:"dataProtectionRequestsSubtab,omitempty"`
	//       The KB search box which appears on the Help page. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string. 
	GetHelpSubtabKnowledgeBaseSearch *string `json:"getHelpSubtabKnowledgeBaseSearch,omitempty"`
	//       The universal search box always visible on Dashboard will, by default, present results from the Meraki KB. This configures       whether these Meraki KB results should be returned. Can be one of 'default or inherit', 'hide' or 'show'. 
	UniversalSearchKnowledgeBaseSearch *string `json:"universalSearchKnowledgeBaseSearch,omitempty"`
	//       The 'Product Manuals' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string. 
	CiscoMerakiProductDocumentation *string `json:"ciscoMerakiProductDocumentation,omitempty"`
	//       The 'Contact Meraki Support' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string. 
	SupportContactInfo *string `json:"supportContactInfo,omitempty"`
	//       The 'Help -> New features' subtab where new Dashboard features are detailed. Can be one of 'default or inherit', 'hide' or 'show'. 
	NewFeaturesSubtab *string `json:"newFeaturesSubtab,omitempty"`
	//       The 'Help -> Firewall info' subtab where necessary upstream firewall rules for communication to the Cisco Meraki cloud are       listed. Can be one of 'default or inherit', 'hide' or 'show'. 
	FirewallInfoSubtab *string `json:"firewallInfoSubtab,omitempty"`
	//       The 'Help -> API docs' subtab where a detailed description of the Dashboard API is listed. Can be one of       'default or inherit', 'hide' or 'show'. 
	ApiDocsSubtab *string `json:"apiDocsSubtab,omitempty"`
	//       The 'Help -> Replacement info' subtab where important information regarding device replacements is detailed. Can be one of       'default or inherit', 'hide' or 'show'. 
	HardwareReplacementsSubtab *string `json:"hardwareReplacementsSubtab,omitempty"`
	//       The 'SM Forums' subtab which links to community-based support for Cisco Meraki Systems Manager. Only configurable for       organizations that contain Systems Manager networks. Can be one of 'default or inherit', 'hide' or 'show'. 
	SmForums *string `json:"smForums,omitempty"`
	//       The 'Help Widget' is a support widget which provides access to live chat, documentation links, Sales contact info,       and other contact avenues to reach Meraki Support. Can be one of 'default or inherit', 'hide' or 'show'. 
	HelpWidget *string `json:"helpWidget,omitempty"`
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
	if o == nil || isNil(o.HelpTab) {
		var ret string
		return ret
	}
	return *o.HelpTab
}

// GetHelpTabOk returns a tuple with the HelpTab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetHelpTabOk() (*string, bool) {
	if o == nil || isNil(o.HelpTab) {
    return nil, false
	}
	return o.HelpTab, true
}

// HasHelpTab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasHelpTab() bool {
	if o != nil && !isNil(o.HelpTab) {
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
	if o == nil || isNil(o.GetHelpSubtab) {
		var ret string
		return ret
	}
	return *o.GetHelpSubtab
}

// GetGetHelpSubtabOk returns a tuple with the GetHelpSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetGetHelpSubtabOk() (*string, bool) {
	if o == nil || isNil(o.GetHelpSubtab) {
    return nil, false
	}
	return o.GetHelpSubtab, true
}

// HasGetHelpSubtab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasGetHelpSubtab() bool {
	if o != nil && !isNil(o.GetHelpSubtab) {
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
	if o == nil || isNil(o.CommunitySubtab) {
		var ret string
		return ret
	}
	return *o.CommunitySubtab
}

// GetCommunitySubtabOk returns a tuple with the CommunitySubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetCommunitySubtabOk() (*string, bool) {
	if o == nil || isNil(o.CommunitySubtab) {
    return nil, false
	}
	return o.CommunitySubtab, true
}

// HasCommunitySubtab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasCommunitySubtab() bool {
	if o != nil && !isNil(o.CommunitySubtab) {
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
	if o == nil || isNil(o.CasesSubtab) {
		var ret string
		return ret
	}
	return *o.CasesSubtab
}

// GetCasesSubtabOk returns a tuple with the CasesSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetCasesSubtabOk() (*string, bool) {
	if o == nil || isNil(o.CasesSubtab) {
    return nil, false
	}
	return o.CasesSubtab, true
}

// HasCasesSubtab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasCasesSubtab() bool {
	if o != nil && !isNil(o.CasesSubtab) {
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
	if o == nil || isNil(o.DataProtectionRequestsSubtab) {
		var ret string
		return ret
	}
	return *o.DataProtectionRequestsSubtab
}

// GetDataProtectionRequestsSubtabOk returns a tuple with the DataProtectionRequestsSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetDataProtectionRequestsSubtabOk() (*string, bool) {
	if o == nil || isNil(o.DataProtectionRequestsSubtab) {
    return nil, false
	}
	return o.DataProtectionRequestsSubtab, true
}

// HasDataProtectionRequestsSubtab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasDataProtectionRequestsSubtab() bool {
	if o != nil && !isNil(o.DataProtectionRequestsSubtab) {
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
	if o == nil || isNil(o.GetHelpSubtabKnowledgeBaseSearch) {
		var ret string
		return ret
	}
	return *o.GetHelpSubtabKnowledgeBaseSearch
}

// GetGetHelpSubtabKnowledgeBaseSearchOk returns a tuple with the GetHelpSubtabKnowledgeBaseSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetGetHelpSubtabKnowledgeBaseSearchOk() (*string, bool) {
	if o == nil || isNil(o.GetHelpSubtabKnowledgeBaseSearch) {
    return nil, false
	}
	return o.GetHelpSubtabKnowledgeBaseSearch, true
}

// HasGetHelpSubtabKnowledgeBaseSearch returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasGetHelpSubtabKnowledgeBaseSearch() bool {
	if o != nil && !isNil(o.GetHelpSubtabKnowledgeBaseSearch) {
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
	if o == nil || isNil(o.UniversalSearchKnowledgeBaseSearch) {
		var ret string
		return ret
	}
	return *o.UniversalSearchKnowledgeBaseSearch
}

// GetUniversalSearchKnowledgeBaseSearchOk returns a tuple with the UniversalSearchKnowledgeBaseSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetUniversalSearchKnowledgeBaseSearchOk() (*string, bool) {
	if o == nil || isNil(o.UniversalSearchKnowledgeBaseSearch) {
    return nil, false
	}
	return o.UniversalSearchKnowledgeBaseSearch, true
}

// HasUniversalSearchKnowledgeBaseSearch returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasUniversalSearchKnowledgeBaseSearch() bool {
	if o != nil && !isNil(o.UniversalSearchKnowledgeBaseSearch) {
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
	if o == nil || isNil(o.CiscoMerakiProductDocumentation) {
		var ret string
		return ret
	}
	return *o.CiscoMerakiProductDocumentation
}

// GetCiscoMerakiProductDocumentationOk returns a tuple with the CiscoMerakiProductDocumentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetCiscoMerakiProductDocumentationOk() (*string, bool) {
	if o == nil || isNil(o.CiscoMerakiProductDocumentation) {
    return nil, false
	}
	return o.CiscoMerakiProductDocumentation, true
}

// HasCiscoMerakiProductDocumentation returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasCiscoMerakiProductDocumentation() bool {
	if o != nil && !isNil(o.CiscoMerakiProductDocumentation) {
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
	if o == nil || isNil(o.SupportContactInfo) {
		var ret string
		return ret
	}
	return *o.SupportContactInfo
}

// GetSupportContactInfoOk returns a tuple with the SupportContactInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetSupportContactInfoOk() (*string, bool) {
	if o == nil || isNil(o.SupportContactInfo) {
    return nil, false
	}
	return o.SupportContactInfo, true
}

// HasSupportContactInfo returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasSupportContactInfo() bool {
	if o != nil && !isNil(o.SupportContactInfo) {
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
	if o == nil || isNil(o.NewFeaturesSubtab) {
		var ret string
		return ret
	}
	return *o.NewFeaturesSubtab
}

// GetNewFeaturesSubtabOk returns a tuple with the NewFeaturesSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetNewFeaturesSubtabOk() (*string, bool) {
	if o == nil || isNil(o.NewFeaturesSubtab) {
    return nil, false
	}
	return o.NewFeaturesSubtab, true
}

// HasNewFeaturesSubtab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasNewFeaturesSubtab() bool {
	if o != nil && !isNil(o.NewFeaturesSubtab) {
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
	if o == nil || isNil(o.FirewallInfoSubtab) {
		var ret string
		return ret
	}
	return *o.FirewallInfoSubtab
}

// GetFirewallInfoSubtabOk returns a tuple with the FirewallInfoSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetFirewallInfoSubtabOk() (*string, bool) {
	if o == nil || isNil(o.FirewallInfoSubtab) {
    return nil, false
	}
	return o.FirewallInfoSubtab, true
}

// HasFirewallInfoSubtab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasFirewallInfoSubtab() bool {
	if o != nil && !isNil(o.FirewallInfoSubtab) {
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
	if o == nil || isNil(o.ApiDocsSubtab) {
		var ret string
		return ret
	}
	return *o.ApiDocsSubtab
}

// GetApiDocsSubtabOk returns a tuple with the ApiDocsSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetApiDocsSubtabOk() (*string, bool) {
	if o == nil || isNil(o.ApiDocsSubtab) {
    return nil, false
	}
	return o.ApiDocsSubtab, true
}

// HasApiDocsSubtab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasApiDocsSubtab() bool {
	if o != nil && !isNil(o.ApiDocsSubtab) {
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
	if o == nil || isNil(o.HardwareReplacementsSubtab) {
		var ret string
		return ret
	}
	return *o.HardwareReplacementsSubtab
}

// GetHardwareReplacementsSubtabOk returns a tuple with the HardwareReplacementsSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetHardwareReplacementsSubtabOk() (*string, bool) {
	if o == nil || isNil(o.HardwareReplacementsSubtab) {
    return nil, false
	}
	return o.HardwareReplacementsSubtab, true
}

// HasHardwareReplacementsSubtab returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasHardwareReplacementsSubtab() bool {
	if o != nil && !isNil(o.HardwareReplacementsSubtab) {
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
	if o == nil || isNil(o.SmForums) {
		var ret string
		return ret
	}
	return *o.SmForums
}

// GetSmForumsOk returns a tuple with the SmForums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetSmForumsOk() (*string, bool) {
	if o == nil || isNil(o.SmForums) {
    return nil, false
	}
	return o.SmForums, true
}

// HasSmForums returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasSmForums() bool {
	if o != nil && !isNil(o.SmForums) {
		return true
	}

	return false
}

// SetSmForums gets a reference to the given string and assigns it to the SmForums field.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetSmForums(v string) {
	o.SmForums = &v
}

// GetHelpWidget returns the HelpWidget field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetHelpWidget() string {
	if o == nil || isNil(o.HelpWidget) {
		var ret string
		return ret
	}
	return *o.HelpWidget
}

// GetHelpWidgetOk returns a tuple with the HelpWidget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetHelpWidgetOk() (*string, bool) {
	if o == nil || isNil(o.HelpWidget) {
    return nil, false
	}
	return o.HelpWidget, true
}

// HasHelpWidget returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasHelpWidget() bool {
	if o != nil && !isNil(o.HelpWidget) {
		return true
	}

	return false
}

// SetHelpWidget gets a reference to the given string and assigns it to the HelpWidget field.
func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetHelpWidget(v string) {
	o.HelpWidget = &v
}

func (o OrganizationsOrganizationIdBrandingPoliciesHelpSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HelpTab) {
		toSerialize["helpTab"] = o.HelpTab
	}
	if !isNil(o.GetHelpSubtab) {
		toSerialize["getHelpSubtab"] = o.GetHelpSubtab
	}
	if !isNil(o.CommunitySubtab) {
		toSerialize["communitySubtab"] = o.CommunitySubtab
	}
	if !isNil(o.CasesSubtab) {
		toSerialize["casesSubtab"] = o.CasesSubtab
	}
	if !isNil(o.DataProtectionRequestsSubtab) {
		toSerialize["dataProtectionRequestsSubtab"] = o.DataProtectionRequestsSubtab
	}
	if !isNil(o.GetHelpSubtabKnowledgeBaseSearch) {
		toSerialize["getHelpSubtabKnowledgeBaseSearch"] = o.GetHelpSubtabKnowledgeBaseSearch
	}
	if !isNil(o.UniversalSearchKnowledgeBaseSearch) {
		toSerialize["universalSearchKnowledgeBaseSearch"] = o.UniversalSearchKnowledgeBaseSearch
	}
	if !isNil(o.CiscoMerakiProductDocumentation) {
		toSerialize["ciscoMerakiProductDocumentation"] = o.CiscoMerakiProductDocumentation
	}
	if !isNil(o.SupportContactInfo) {
		toSerialize["supportContactInfo"] = o.SupportContactInfo
	}
	if !isNil(o.NewFeaturesSubtab) {
		toSerialize["newFeaturesSubtab"] = o.NewFeaturesSubtab
	}
	if !isNil(o.FirewallInfoSubtab) {
		toSerialize["firewallInfoSubtab"] = o.FirewallInfoSubtab
	}
	if !isNil(o.ApiDocsSubtab) {
		toSerialize["apiDocsSubtab"] = o.ApiDocsSubtab
	}
	if !isNil(o.HardwareReplacementsSubtab) {
		toSerialize["hardwareReplacementsSubtab"] = o.HardwareReplacementsSubtab
	}
	if !isNil(o.SmForums) {
		toSerialize["smForums"] = o.SmForums
	}
	if !isNil(o.HelpWidget) {
		toSerialize["helpWidget"] = o.HelpWidget
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


