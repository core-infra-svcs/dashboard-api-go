# InlineResponse200422DestinationsCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressRanges** | Pointer to **[]string** | Address ranges or addresses | [optional] 
**Ports** | Pointer to **[]string** | Port numbers or ranges | [optional] 
**Services** | Pointer to [**[]InlineResponse200422DestinationsCriteriaServices**](InlineResponse200422DestinationsCriteriaServices.md) | Protocol and port services | [optional] 
**ApplicationCategories** | Pointer to [**[]InlineResponse200422DestinationsCriteriaApplicationCategories**](InlineResponse200422DestinationsCriteriaApplicationCategories.md) | Application categories | [optional] 
**Applications** | Pointer to [**[]InlineResponse200422DestinationsCriteriaApplications**](InlineResponse200422DestinationsCriteriaApplications.md) | Applications | [optional] 
**PolicyObjects** | Pointer to [**[]InlineResponse200422SourcesCriteriaPolicyObjects**](InlineResponse200422SourcesCriteriaPolicyObjects.md) | Policy objects | [optional] 
**PolicyObjectGroups** | Pointer to [**[]InlineResponse200422SourcesCriteriaPolicyObjectGroups**](InlineResponse200422SourcesCriteriaPolicyObjectGroups.md) | Policy object groups | [optional] 
**ApplianceVlans** | Pointer to [**[]InlineResponse200422SourcesCriteriaApplianceVlans**](InlineResponse200422SourcesCriteriaApplianceVlans.md) | Appliance VLANs | [optional] 
**Countries** | Pointer to [**[]InlineResponse200422DestinationsCriteriaCountries**](InlineResponse200422DestinationsCriteriaCountries.md) | Countries | [optional] 
**Fqdns** | Pointer to **[]string** | Fully qualified domain names or wildcard patterns | [optional] 
**SiteSpecificVlans** | Pointer to [**[]InlineResponse200422SourcesCriteriaSiteSpecificVlans**](InlineResponse200422SourcesCriteriaSiteSpecificVlans.md) | Site-specific VLANs with address offset configuration. Maximum 100 items. | [optional] 
**AdaptivePolicyGroups** | Pointer to [**[]InlineResponse200422DestinationsCriteriaAdaptivePolicyGroups**](InlineResponse200422DestinationsCriteriaAdaptivePolicyGroups.md) | Adaptive policy groups | [optional] 

## Methods

### NewInlineResponse200422DestinationsCriteria

`func NewInlineResponse200422DestinationsCriteria() *InlineResponse200422DestinationsCriteria`

NewInlineResponse200422DestinationsCriteria instantiates a new InlineResponse200422DestinationsCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200422DestinationsCriteriaWithDefaults

`func NewInlineResponse200422DestinationsCriteriaWithDefaults() *InlineResponse200422DestinationsCriteria`

NewInlineResponse200422DestinationsCriteriaWithDefaults instantiates a new InlineResponse200422DestinationsCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressRanges

`func (o *InlineResponse200422DestinationsCriteria) GetAddressRanges() []string`

GetAddressRanges returns the AddressRanges field if non-nil, zero value otherwise.

### GetAddressRangesOk

`func (o *InlineResponse200422DestinationsCriteria) GetAddressRangesOk() (*[]string, bool)`

GetAddressRangesOk returns a tuple with the AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRanges

`func (o *InlineResponse200422DestinationsCriteria) SetAddressRanges(v []string)`

SetAddressRanges sets AddressRanges field to given value.

### HasAddressRanges

`func (o *InlineResponse200422DestinationsCriteria) HasAddressRanges() bool`

HasAddressRanges returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200422DestinationsCriteria) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200422DestinationsCriteria) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200422DestinationsCriteria) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200422DestinationsCriteria) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetServices

`func (o *InlineResponse200422DestinationsCriteria) GetServices() []InlineResponse200422DestinationsCriteriaServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *InlineResponse200422DestinationsCriteria) GetServicesOk() (*[]InlineResponse200422DestinationsCriteriaServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *InlineResponse200422DestinationsCriteria) SetServices(v []InlineResponse200422DestinationsCriteriaServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *InlineResponse200422DestinationsCriteria) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetApplicationCategories

`func (o *InlineResponse200422DestinationsCriteria) GetApplicationCategories() []InlineResponse200422DestinationsCriteriaApplicationCategories`

GetApplicationCategories returns the ApplicationCategories field if non-nil, zero value otherwise.

### GetApplicationCategoriesOk

`func (o *InlineResponse200422DestinationsCriteria) GetApplicationCategoriesOk() (*[]InlineResponse200422DestinationsCriteriaApplicationCategories, bool)`

GetApplicationCategoriesOk returns a tuple with the ApplicationCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCategories

`func (o *InlineResponse200422DestinationsCriteria) SetApplicationCategories(v []InlineResponse200422DestinationsCriteriaApplicationCategories)`

SetApplicationCategories sets ApplicationCategories field to given value.

### HasApplicationCategories

`func (o *InlineResponse200422DestinationsCriteria) HasApplicationCategories() bool`

HasApplicationCategories returns a boolean if a field has been set.

### GetApplications

`func (o *InlineResponse200422DestinationsCriteria) GetApplications() []InlineResponse200422DestinationsCriteriaApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *InlineResponse200422DestinationsCriteria) GetApplicationsOk() (*[]InlineResponse200422DestinationsCriteriaApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *InlineResponse200422DestinationsCriteria) SetApplications(v []InlineResponse200422DestinationsCriteriaApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *InlineResponse200422DestinationsCriteria) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *InlineResponse200422DestinationsCriteria) GetPolicyObjects() []InlineResponse200422SourcesCriteriaPolicyObjects`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *InlineResponse200422DestinationsCriteria) GetPolicyObjectsOk() (*[]InlineResponse200422SourcesCriteriaPolicyObjects, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *InlineResponse200422DestinationsCriteria) SetPolicyObjects(v []InlineResponse200422SourcesCriteriaPolicyObjects)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *InlineResponse200422DestinationsCriteria) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.

### GetPolicyObjectGroups

`func (o *InlineResponse200422DestinationsCriteria) GetPolicyObjectGroups() []InlineResponse200422SourcesCriteriaPolicyObjectGroups`

GetPolicyObjectGroups returns the PolicyObjectGroups field if non-nil, zero value otherwise.

### GetPolicyObjectGroupsOk

`func (o *InlineResponse200422DestinationsCriteria) GetPolicyObjectGroupsOk() (*[]InlineResponse200422SourcesCriteriaPolicyObjectGroups, bool)`

GetPolicyObjectGroupsOk returns a tuple with the PolicyObjectGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjectGroups

`func (o *InlineResponse200422DestinationsCriteria) SetPolicyObjectGroups(v []InlineResponse200422SourcesCriteriaPolicyObjectGroups)`

SetPolicyObjectGroups sets PolicyObjectGroups field to given value.

### HasPolicyObjectGroups

`func (o *InlineResponse200422DestinationsCriteria) HasPolicyObjectGroups() bool`

HasPolicyObjectGroups returns a boolean if a field has been set.

### GetApplianceVlans

`func (o *InlineResponse200422DestinationsCriteria) GetApplianceVlans() []InlineResponse200422SourcesCriteriaApplianceVlans`

GetApplianceVlans returns the ApplianceVlans field if non-nil, zero value otherwise.

### GetApplianceVlansOk

`func (o *InlineResponse200422DestinationsCriteria) GetApplianceVlansOk() (*[]InlineResponse200422SourcesCriteriaApplianceVlans, bool)`

GetApplianceVlansOk returns a tuple with the ApplianceVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceVlans

`func (o *InlineResponse200422DestinationsCriteria) SetApplianceVlans(v []InlineResponse200422SourcesCriteriaApplianceVlans)`

SetApplianceVlans sets ApplianceVlans field to given value.

### HasApplianceVlans

`func (o *InlineResponse200422DestinationsCriteria) HasApplianceVlans() bool`

HasApplianceVlans returns a boolean if a field has been set.

### GetCountries

`func (o *InlineResponse200422DestinationsCriteria) GetCountries() []InlineResponse200422DestinationsCriteriaCountries`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *InlineResponse200422DestinationsCriteria) GetCountriesOk() (*[]InlineResponse200422DestinationsCriteriaCountries, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *InlineResponse200422DestinationsCriteria) SetCountries(v []InlineResponse200422DestinationsCriteriaCountries)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *InlineResponse200422DestinationsCriteria) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetFqdns

`func (o *InlineResponse200422DestinationsCriteria) GetFqdns() []string`

GetFqdns returns the Fqdns field if non-nil, zero value otherwise.

### GetFqdnsOk

`func (o *InlineResponse200422DestinationsCriteria) GetFqdnsOk() (*[]string, bool)`

GetFqdnsOk returns a tuple with the Fqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdns

`func (o *InlineResponse200422DestinationsCriteria) SetFqdns(v []string)`

SetFqdns sets Fqdns field to given value.

### HasFqdns

`func (o *InlineResponse200422DestinationsCriteria) HasFqdns() bool`

HasFqdns returns a boolean if a field has been set.

### GetSiteSpecificVlans

`func (o *InlineResponse200422DestinationsCriteria) GetSiteSpecificVlans() []InlineResponse200422SourcesCriteriaSiteSpecificVlans`

GetSiteSpecificVlans returns the SiteSpecificVlans field if non-nil, zero value otherwise.

### GetSiteSpecificVlansOk

`func (o *InlineResponse200422DestinationsCriteria) GetSiteSpecificVlansOk() (*[]InlineResponse200422SourcesCriteriaSiteSpecificVlans, bool)`

GetSiteSpecificVlansOk returns a tuple with the SiteSpecificVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteSpecificVlans

`func (o *InlineResponse200422DestinationsCriteria) SetSiteSpecificVlans(v []InlineResponse200422SourcesCriteriaSiteSpecificVlans)`

SetSiteSpecificVlans sets SiteSpecificVlans field to given value.

### HasSiteSpecificVlans

`func (o *InlineResponse200422DestinationsCriteria) HasSiteSpecificVlans() bool`

HasSiteSpecificVlans returns a boolean if a field has been set.

### GetAdaptivePolicyGroups

`func (o *InlineResponse200422DestinationsCriteria) GetAdaptivePolicyGroups() []InlineResponse200422DestinationsCriteriaAdaptivePolicyGroups`

GetAdaptivePolicyGroups returns the AdaptivePolicyGroups field if non-nil, zero value otherwise.

### GetAdaptivePolicyGroupsOk

`func (o *InlineResponse200422DestinationsCriteria) GetAdaptivePolicyGroupsOk() (*[]InlineResponse200422DestinationsCriteriaAdaptivePolicyGroups, bool)`

GetAdaptivePolicyGroupsOk returns a tuple with the AdaptivePolicyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptivePolicyGroups

`func (o *InlineResponse200422DestinationsCriteria) SetAdaptivePolicyGroups(v []InlineResponse200422DestinationsCriteriaAdaptivePolicyGroups)`

SetAdaptivePolicyGroups sets AdaptivePolicyGroups field to given value.

### HasAdaptivePolicyGroups

`func (o *InlineResponse200422DestinationsCriteria) HasAdaptivePolicyGroups() bool`

HasAdaptivePolicyGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


