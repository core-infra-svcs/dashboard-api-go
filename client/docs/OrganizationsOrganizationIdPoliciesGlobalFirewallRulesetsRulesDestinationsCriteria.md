# OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressRanges** | Pointer to **[]string** | Address ranges or addresses | [optional] 
**Ports** | Pointer to **[]string** | Port numbers or ranges | [optional] 
**Services** | Pointer to [**[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaServices**](OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaServices.md) | Protocol and port services | [optional] 
**ApplicationCategories** | Pointer to [**[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaApplicationCategories**](OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaApplicationCategories.md) | Application categories | [optional] 
**Applications** | Pointer to [**[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaApplications**](OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaApplications.md) | Applications | [optional] 
**PolicyObjects** | Pointer to [**[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjects**](OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjects.md) | Policy objects | [optional] 
**PolicyObjectGroups** | Pointer to [**[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjectGroups**](OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjectGroups.md) | Policy object groups | [optional] 
**ApplianceVlans** | Pointer to [**[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaApplianceVlans**](OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaApplianceVlans.md) | Appliance VLANs | [optional] 

## Methods

### NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria

`func NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria() *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria`

NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria instantiates a new OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaWithDefaults

`func NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaWithDefaults() *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria`

NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaWithDefaults instantiates a new OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressRanges

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) GetAddressRanges() []string`

GetAddressRanges returns the AddressRanges field if non-nil, zero value otherwise.

### GetAddressRangesOk

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) GetAddressRangesOk() (*[]string, bool)`

GetAddressRangesOk returns a tuple with the AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRanges

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) SetAddressRanges(v []string)`

SetAddressRanges sets AddressRanges field to given value.

### HasAddressRanges

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) HasAddressRanges() bool`

HasAddressRanges returns a boolean if a field has been set.

### GetPorts

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetServices

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) GetServices() []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) GetServicesOk() (*[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) SetServices(v []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetApplicationCategories

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) GetApplicationCategories() []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaApplicationCategories`

GetApplicationCategories returns the ApplicationCategories field if non-nil, zero value otherwise.

### GetApplicationCategoriesOk

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) GetApplicationCategoriesOk() (*[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaApplicationCategories, bool)`

GetApplicationCategoriesOk returns a tuple with the ApplicationCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCategories

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) SetApplicationCategories(v []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaApplicationCategories)`

SetApplicationCategories sets ApplicationCategories field to given value.

### HasApplicationCategories

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) HasApplicationCategories() bool`

HasApplicationCategories returns a boolean if a field has been set.

### GetApplications

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) GetApplications() []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) GetApplicationsOk() (*[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) SetApplications(v []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) GetPolicyObjects() []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjects`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) GetPolicyObjectsOk() (*[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjects, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) SetPolicyObjects(v []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjects)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.

### GetPolicyObjectGroups

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) GetPolicyObjectGroups() []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjectGroups`

GetPolicyObjectGroups returns the PolicyObjectGroups field if non-nil, zero value otherwise.

### GetPolicyObjectGroupsOk

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) GetPolicyObjectGroupsOk() (*[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjectGroups, bool)`

GetPolicyObjectGroupsOk returns a tuple with the PolicyObjectGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjectGroups

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) SetPolicyObjectGroups(v []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaPolicyObjectGroups)`

SetPolicyObjectGroups sets PolicyObjectGroups field to given value.

### HasPolicyObjectGroups

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) HasPolicyObjectGroups() bool`

HasPolicyObjectGroups returns a boolean if a field has been set.

### GetApplianceVlans

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) GetApplianceVlans() []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaApplianceVlans`

GetApplianceVlans returns the ApplianceVlans field if non-nil, zero value otherwise.

### GetApplianceVlansOk

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) GetApplianceVlansOk() (*[]OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaApplianceVlans, bool)`

GetApplianceVlansOk returns a tuple with the ApplianceVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceVlans

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) SetApplianceVlans(v []OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSourcesCriteriaApplianceVlans)`

SetApplianceVlans sets ApplianceVlans field to given value.

### HasApplianceVlans

`func (o *OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteria) HasApplianceVlans() bool`

HasApplianceVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


