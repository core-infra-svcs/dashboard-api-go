# InlineResponse200434DestinationsCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressRanges** | Pointer to **[]string** | Address ranges or addresses | [optional] 
**Ports** | Pointer to **[]string** | Port numbers or ranges | [optional] 
**Services** | Pointer to [**[]InlineResponse200434DestinationsCriteriaServices**](InlineResponse200434DestinationsCriteriaServices.md) | Protocol and port services | [optional] 
**ApplicationCategories** | Pointer to [**[]InlineResponse200434DestinationsCriteriaApplicationCategories**](InlineResponse200434DestinationsCriteriaApplicationCategories.md) | Application categories | [optional] 
**Applications** | Pointer to [**[]InlineResponse200434DestinationsCriteriaApplications**](InlineResponse200434DestinationsCriteriaApplications.md) | Applications | [optional] 
**PolicyObjects** | Pointer to [**[]InlineResponse200434SourcesCriteriaPolicyObjects**](InlineResponse200434SourcesCriteriaPolicyObjects.md) | Policy objects | [optional] 
**PolicyObjectGroups** | Pointer to [**[]InlineResponse200434SourcesCriteriaPolicyObjectGroups**](InlineResponse200434SourcesCriteriaPolicyObjectGroups.md) | Policy object groups | [optional] 
**ApplianceVlans** | Pointer to [**[]InlineResponse200434SourcesCriteriaApplianceVlans**](InlineResponse200434SourcesCriteriaApplianceVlans.md) | Appliance VLANs | [optional] 
**Countries** | Pointer to [**[]InlineResponse200434DestinationsCriteriaCountries**](InlineResponse200434DestinationsCriteriaCountries.md) | Countries | [optional] 
**Fqdns** | Pointer to **[]string** | Fully qualified domain names or wildcard patterns | [optional] 
**SiteSpecificVlans** | Pointer to [**[]InlineResponse200434SourcesCriteriaSiteSpecificVlans**](InlineResponse200434SourcesCriteriaSiteSpecificVlans.md) | Site-specific VLANs with address offset configuration. Maximum 100 items. | [optional] 
**AdaptivePolicyGroups** | Pointer to [**[]InlineResponse200434DestinationsCriteriaAdaptivePolicyGroups**](InlineResponse200434DestinationsCriteriaAdaptivePolicyGroups.md) | Adaptive policy groups | [optional] 

## Methods

### NewInlineResponse200434DestinationsCriteria

`func NewInlineResponse200434DestinationsCriteria() *InlineResponse200434DestinationsCriteria`

NewInlineResponse200434DestinationsCriteria instantiates a new InlineResponse200434DestinationsCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200434DestinationsCriteriaWithDefaults

`func NewInlineResponse200434DestinationsCriteriaWithDefaults() *InlineResponse200434DestinationsCriteria`

NewInlineResponse200434DestinationsCriteriaWithDefaults instantiates a new InlineResponse200434DestinationsCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressRanges

`func (o *InlineResponse200434DestinationsCriteria) GetAddressRanges() []string`

GetAddressRanges returns the AddressRanges field if non-nil, zero value otherwise.

### GetAddressRangesOk

`func (o *InlineResponse200434DestinationsCriteria) GetAddressRangesOk() (*[]string, bool)`

GetAddressRangesOk returns a tuple with the AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRanges

`func (o *InlineResponse200434DestinationsCriteria) SetAddressRanges(v []string)`

SetAddressRanges sets AddressRanges field to given value.

### HasAddressRanges

`func (o *InlineResponse200434DestinationsCriteria) HasAddressRanges() bool`

HasAddressRanges returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200434DestinationsCriteria) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200434DestinationsCriteria) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200434DestinationsCriteria) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200434DestinationsCriteria) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetServices

`func (o *InlineResponse200434DestinationsCriteria) GetServices() []InlineResponse200434DestinationsCriteriaServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *InlineResponse200434DestinationsCriteria) GetServicesOk() (*[]InlineResponse200434DestinationsCriteriaServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *InlineResponse200434DestinationsCriteria) SetServices(v []InlineResponse200434DestinationsCriteriaServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *InlineResponse200434DestinationsCriteria) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetApplicationCategories

`func (o *InlineResponse200434DestinationsCriteria) GetApplicationCategories() []InlineResponse200434DestinationsCriteriaApplicationCategories`

GetApplicationCategories returns the ApplicationCategories field if non-nil, zero value otherwise.

### GetApplicationCategoriesOk

`func (o *InlineResponse200434DestinationsCriteria) GetApplicationCategoriesOk() (*[]InlineResponse200434DestinationsCriteriaApplicationCategories, bool)`

GetApplicationCategoriesOk returns a tuple with the ApplicationCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCategories

`func (o *InlineResponse200434DestinationsCriteria) SetApplicationCategories(v []InlineResponse200434DestinationsCriteriaApplicationCategories)`

SetApplicationCategories sets ApplicationCategories field to given value.

### HasApplicationCategories

`func (o *InlineResponse200434DestinationsCriteria) HasApplicationCategories() bool`

HasApplicationCategories returns a boolean if a field has been set.

### GetApplications

`func (o *InlineResponse200434DestinationsCriteria) GetApplications() []InlineResponse200434DestinationsCriteriaApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *InlineResponse200434DestinationsCriteria) GetApplicationsOk() (*[]InlineResponse200434DestinationsCriteriaApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *InlineResponse200434DestinationsCriteria) SetApplications(v []InlineResponse200434DestinationsCriteriaApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *InlineResponse200434DestinationsCriteria) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *InlineResponse200434DestinationsCriteria) GetPolicyObjects() []InlineResponse200434SourcesCriteriaPolicyObjects`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *InlineResponse200434DestinationsCriteria) GetPolicyObjectsOk() (*[]InlineResponse200434SourcesCriteriaPolicyObjects, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *InlineResponse200434DestinationsCriteria) SetPolicyObjects(v []InlineResponse200434SourcesCriteriaPolicyObjects)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *InlineResponse200434DestinationsCriteria) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.

### GetPolicyObjectGroups

`func (o *InlineResponse200434DestinationsCriteria) GetPolicyObjectGroups() []InlineResponse200434SourcesCriteriaPolicyObjectGroups`

GetPolicyObjectGroups returns the PolicyObjectGroups field if non-nil, zero value otherwise.

### GetPolicyObjectGroupsOk

`func (o *InlineResponse200434DestinationsCriteria) GetPolicyObjectGroupsOk() (*[]InlineResponse200434SourcesCriteriaPolicyObjectGroups, bool)`

GetPolicyObjectGroupsOk returns a tuple with the PolicyObjectGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjectGroups

`func (o *InlineResponse200434DestinationsCriteria) SetPolicyObjectGroups(v []InlineResponse200434SourcesCriteriaPolicyObjectGroups)`

SetPolicyObjectGroups sets PolicyObjectGroups field to given value.

### HasPolicyObjectGroups

`func (o *InlineResponse200434DestinationsCriteria) HasPolicyObjectGroups() bool`

HasPolicyObjectGroups returns a boolean if a field has been set.

### GetApplianceVlans

`func (o *InlineResponse200434DestinationsCriteria) GetApplianceVlans() []InlineResponse200434SourcesCriteriaApplianceVlans`

GetApplianceVlans returns the ApplianceVlans field if non-nil, zero value otherwise.

### GetApplianceVlansOk

`func (o *InlineResponse200434DestinationsCriteria) GetApplianceVlansOk() (*[]InlineResponse200434SourcesCriteriaApplianceVlans, bool)`

GetApplianceVlansOk returns a tuple with the ApplianceVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceVlans

`func (o *InlineResponse200434DestinationsCriteria) SetApplianceVlans(v []InlineResponse200434SourcesCriteriaApplianceVlans)`

SetApplianceVlans sets ApplianceVlans field to given value.

### HasApplianceVlans

`func (o *InlineResponse200434DestinationsCriteria) HasApplianceVlans() bool`

HasApplianceVlans returns a boolean if a field has been set.

### GetCountries

`func (o *InlineResponse200434DestinationsCriteria) GetCountries() []InlineResponse200434DestinationsCriteriaCountries`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *InlineResponse200434DestinationsCriteria) GetCountriesOk() (*[]InlineResponse200434DestinationsCriteriaCountries, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *InlineResponse200434DestinationsCriteria) SetCountries(v []InlineResponse200434DestinationsCriteriaCountries)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *InlineResponse200434DestinationsCriteria) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetFqdns

`func (o *InlineResponse200434DestinationsCriteria) GetFqdns() []string`

GetFqdns returns the Fqdns field if non-nil, zero value otherwise.

### GetFqdnsOk

`func (o *InlineResponse200434DestinationsCriteria) GetFqdnsOk() (*[]string, bool)`

GetFqdnsOk returns a tuple with the Fqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdns

`func (o *InlineResponse200434DestinationsCriteria) SetFqdns(v []string)`

SetFqdns sets Fqdns field to given value.

### HasFqdns

`func (o *InlineResponse200434DestinationsCriteria) HasFqdns() bool`

HasFqdns returns a boolean if a field has been set.

### GetSiteSpecificVlans

`func (o *InlineResponse200434DestinationsCriteria) GetSiteSpecificVlans() []InlineResponse200434SourcesCriteriaSiteSpecificVlans`

GetSiteSpecificVlans returns the SiteSpecificVlans field if non-nil, zero value otherwise.

### GetSiteSpecificVlansOk

`func (o *InlineResponse200434DestinationsCriteria) GetSiteSpecificVlansOk() (*[]InlineResponse200434SourcesCriteriaSiteSpecificVlans, bool)`

GetSiteSpecificVlansOk returns a tuple with the SiteSpecificVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteSpecificVlans

`func (o *InlineResponse200434DestinationsCriteria) SetSiteSpecificVlans(v []InlineResponse200434SourcesCriteriaSiteSpecificVlans)`

SetSiteSpecificVlans sets SiteSpecificVlans field to given value.

### HasSiteSpecificVlans

`func (o *InlineResponse200434DestinationsCriteria) HasSiteSpecificVlans() bool`

HasSiteSpecificVlans returns a boolean if a field has been set.

### GetAdaptivePolicyGroups

`func (o *InlineResponse200434DestinationsCriteria) GetAdaptivePolicyGroups() []InlineResponse200434DestinationsCriteriaAdaptivePolicyGroups`

GetAdaptivePolicyGroups returns the AdaptivePolicyGroups field if non-nil, zero value otherwise.

### GetAdaptivePolicyGroupsOk

`func (o *InlineResponse200434DestinationsCriteria) GetAdaptivePolicyGroupsOk() (*[]InlineResponse200434DestinationsCriteriaAdaptivePolicyGroups, bool)`

GetAdaptivePolicyGroupsOk returns a tuple with the AdaptivePolicyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptivePolicyGroups

`func (o *InlineResponse200434DestinationsCriteria) SetAdaptivePolicyGroups(v []InlineResponse200434DestinationsCriteriaAdaptivePolicyGroups)`

SetAdaptivePolicyGroups sets AdaptivePolicyGroups field to given value.

### HasAdaptivePolicyGroups

`func (o *InlineResponse200434DestinationsCriteria) HasAdaptivePolicyGroups() bool`

HasAdaptivePolicyGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


