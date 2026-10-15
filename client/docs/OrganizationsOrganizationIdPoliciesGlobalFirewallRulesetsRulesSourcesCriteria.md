# OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressRanges** | Pointer to **[]string** | Address ranges or addresses | [optional] 
**Ports** | Pointer to **[]string** | Port numbers or ranges | [optional] 
**PolicyObjects** | Pointer to [**[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjects**](OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjects.md) | Policy objects | [optional] 
**PolicyObjectGroups** | Pointer to [**[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjectGroups**](OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjectGroups.md) | Policy object groups | [optional] 
**ApplianceVlans** | Pointer to [**[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaApplianceVlans**](OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaApplianceVlans.md) | Appliance VLANs | [optional] 
**SiteSpecificVlans** | Pointer to [**[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaSiteSpecificVlans**](OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaSiteSpecificVlans.md) | Site-specific VLANs with address offset configuration. | [optional] 

## Methods

### NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria

`func NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria() *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria`

NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria instantiates a new OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaWithDefaults

`func NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaWithDefaults() *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria`

NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaWithDefaults instantiates a new OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressRanges

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) GetAddressRanges() []string`

GetAddressRanges returns the AddressRanges field if non-nil, zero value otherwise.

### GetAddressRangesOk

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) GetAddressRangesOk() (*[]string, bool)`

GetAddressRangesOk returns a tuple with the AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRanges

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) SetAddressRanges(v []string)`

SetAddressRanges sets AddressRanges field to given value.

### HasAddressRanges

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) HasAddressRanges() bool`

HasAddressRanges returns a boolean if a field has been set.

### GetPorts

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) GetPolicyObjects() []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjects`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) GetPolicyObjectsOk() (*[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjects, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) SetPolicyObjects(v []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjects)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.

### GetPolicyObjectGroups

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) GetPolicyObjectGroups() []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjectGroups`

GetPolicyObjectGroups returns the PolicyObjectGroups field if non-nil, zero value otherwise.

### GetPolicyObjectGroupsOk

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) GetPolicyObjectGroupsOk() (*[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjectGroups, bool)`

GetPolicyObjectGroupsOk returns a tuple with the PolicyObjectGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjectGroups

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) SetPolicyObjectGroups(v []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjectGroups)`

SetPolicyObjectGroups sets PolicyObjectGroups field to given value.

### HasPolicyObjectGroups

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) HasPolicyObjectGroups() bool`

HasPolicyObjectGroups returns a boolean if a field has been set.

### GetApplianceVlans

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) GetApplianceVlans() []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaApplianceVlans`

GetApplianceVlans returns the ApplianceVlans field if non-nil, zero value otherwise.

### GetApplianceVlansOk

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) GetApplianceVlansOk() (*[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaApplianceVlans, bool)`

GetApplianceVlansOk returns a tuple with the ApplianceVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceVlans

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) SetApplianceVlans(v []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaApplianceVlans)`

SetApplianceVlans sets ApplianceVlans field to given value.

### HasApplianceVlans

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) HasApplianceVlans() bool`

HasApplianceVlans returns a boolean if a field has been set.

### GetSiteSpecificVlans

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) GetSiteSpecificVlans() []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaSiteSpecificVlans`

GetSiteSpecificVlans returns the SiteSpecificVlans field if non-nil, zero value otherwise.

### GetSiteSpecificVlansOk

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) GetSiteSpecificVlansOk() (*[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaSiteSpecificVlans, bool)`

GetSiteSpecificVlansOk returns a tuple with the SiteSpecificVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteSpecificVlans

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) SetSiteSpecificVlans(v []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaSiteSpecificVlans)`

SetSiteSpecificVlans sets SiteSpecificVlans field to given value.

### HasSiteSpecificVlans

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteria) HasSiteSpecificVlans() bool`

HasSiteSpecificVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


