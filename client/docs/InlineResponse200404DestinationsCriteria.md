# InlineResponse200404DestinationsCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressRanges** | Pointer to **[]string** | Address ranges or addresses | [optional] 
**Ports** | Pointer to **[]string** | Port numbers or ranges | [optional] 
**Services** | Pointer to [**[]InlineResponse200404DestinationsCriteriaServices**](InlineResponse200404DestinationsCriteriaServices.md) | Protocol and port services | [optional] 
**ApplicationCategories** | Pointer to [**[]InlineResponse200404DestinationsCriteriaApplicationCategories**](InlineResponse200404DestinationsCriteriaApplicationCategories.md) | Application categories | [optional] 
**Applications** | Pointer to [**[]InlineResponse200404DestinationsCriteriaApplications**](InlineResponse200404DestinationsCriteriaApplications.md) | Applications | [optional] 
**PolicyObjects** | Pointer to [**[]InlineResponse200404SourcesCriteriaPolicyObjects**](InlineResponse200404SourcesCriteriaPolicyObjects.md) | Policy objects | [optional] 
**PolicyObjectGroups** | Pointer to [**[]InlineResponse200404SourcesCriteriaPolicyObjectGroups**](InlineResponse200404SourcesCriteriaPolicyObjectGroups.md) | Policy object groups | [optional] 
**ApplianceVlans** | Pointer to [**[]InlineResponse200404SourcesCriteriaApplianceVlans**](InlineResponse200404SourcesCriteriaApplianceVlans.md) | Appliance VLANs | [optional] 
**Countries** | Pointer to [**[]InlineResponse200404DestinationsCriteriaCountries**](InlineResponse200404DestinationsCriteriaCountries.md) | Countries | [optional] 
**Fqdns** | Pointer to **[]string** | Fully qualified domain names or wildcard patterns | [optional] 
**SiteSpecificVlans** | Pointer to [**[]InlineResponse200404SourcesCriteriaSiteSpecificVlans**](InlineResponse200404SourcesCriteriaSiteSpecificVlans.md) | Site-specific VLANs with address offset configuration. Maximum 100 items. | [optional] 

## Methods

### NewInlineResponse200404DestinationsCriteria

`func NewInlineResponse200404DestinationsCriteria() *InlineResponse200404DestinationsCriteria`

NewInlineResponse200404DestinationsCriteria instantiates a new InlineResponse200404DestinationsCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200404DestinationsCriteriaWithDefaults

`func NewInlineResponse200404DestinationsCriteriaWithDefaults() *InlineResponse200404DestinationsCriteria`

NewInlineResponse200404DestinationsCriteriaWithDefaults instantiates a new InlineResponse200404DestinationsCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressRanges

`func (o *InlineResponse200404DestinationsCriteria) GetAddressRanges() []string`

GetAddressRanges returns the AddressRanges field if non-nil, zero value otherwise.

### GetAddressRangesOk

`func (o *InlineResponse200404DestinationsCriteria) GetAddressRangesOk() (*[]string, bool)`

GetAddressRangesOk returns a tuple with the AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRanges

`func (o *InlineResponse200404DestinationsCriteria) SetAddressRanges(v []string)`

SetAddressRanges sets AddressRanges field to given value.

### HasAddressRanges

`func (o *InlineResponse200404DestinationsCriteria) HasAddressRanges() bool`

HasAddressRanges returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200404DestinationsCriteria) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200404DestinationsCriteria) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200404DestinationsCriteria) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200404DestinationsCriteria) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetServices

`func (o *InlineResponse200404DestinationsCriteria) GetServices() []InlineResponse200404DestinationsCriteriaServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *InlineResponse200404DestinationsCriteria) GetServicesOk() (*[]InlineResponse200404DestinationsCriteriaServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *InlineResponse200404DestinationsCriteria) SetServices(v []InlineResponse200404DestinationsCriteriaServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *InlineResponse200404DestinationsCriteria) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetApplicationCategories

`func (o *InlineResponse200404DestinationsCriteria) GetApplicationCategories() []InlineResponse200404DestinationsCriteriaApplicationCategories`

GetApplicationCategories returns the ApplicationCategories field if non-nil, zero value otherwise.

### GetApplicationCategoriesOk

`func (o *InlineResponse200404DestinationsCriteria) GetApplicationCategoriesOk() (*[]InlineResponse200404DestinationsCriteriaApplicationCategories, bool)`

GetApplicationCategoriesOk returns a tuple with the ApplicationCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCategories

`func (o *InlineResponse200404DestinationsCriteria) SetApplicationCategories(v []InlineResponse200404DestinationsCriteriaApplicationCategories)`

SetApplicationCategories sets ApplicationCategories field to given value.

### HasApplicationCategories

`func (o *InlineResponse200404DestinationsCriteria) HasApplicationCategories() bool`

HasApplicationCategories returns a boolean if a field has been set.

### GetApplications

`func (o *InlineResponse200404DestinationsCriteria) GetApplications() []InlineResponse200404DestinationsCriteriaApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *InlineResponse200404DestinationsCriteria) GetApplicationsOk() (*[]InlineResponse200404DestinationsCriteriaApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *InlineResponse200404DestinationsCriteria) SetApplications(v []InlineResponse200404DestinationsCriteriaApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *InlineResponse200404DestinationsCriteria) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *InlineResponse200404DestinationsCriteria) GetPolicyObjects() []InlineResponse200404SourcesCriteriaPolicyObjects`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *InlineResponse200404DestinationsCriteria) GetPolicyObjectsOk() (*[]InlineResponse200404SourcesCriteriaPolicyObjects, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *InlineResponse200404DestinationsCriteria) SetPolicyObjects(v []InlineResponse200404SourcesCriteriaPolicyObjects)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *InlineResponse200404DestinationsCriteria) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.

### GetPolicyObjectGroups

`func (o *InlineResponse200404DestinationsCriteria) GetPolicyObjectGroups() []InlineResponse200404SourcesCriteriaPolicyObjectGroups`

GetPolicyObjectGroups returns the PolicyObjectGroups field if non-nil, zero value otherwise.

### GetPolicyObjectGroupsOk

`func (o *InlineResponse200404DestinationsCriteria) GetPolicyObjectGroupsOk() (*[]InlineResponse200404SourcesCriteriaPolicyObjectGroups, bool)`

GetPolicyObjectGroupsOk returns a tuple with the PolicyObjectGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjectGroups

`func (o *InlineResponse200404DestinationsCriteria) SetPolicyObjectGroups(v []InlineResponse200404SourcesCriteriaPolicyObjectGroups)`

SetPolicyObjectGroups sets PolicyObjectGroups field to given value.

### HasPolicyObjectGroups

`func (o *InlineResponse200404DestinationsCriteria) HasPolicyObjectGroups() bool`

HasPolicyObjectGroups returns a boolean if a field has been set.

### GetApplianceVlans

`func (o *InlineResponse200404DestinationsCriteria) GetApplianceVlans() []InlineResponse200404SourcesCriteriaApplianceVlans`

GetApplianceVlans returns the ApplianceVlans field if non-nil, zero value otherwise.

### GetApplianceVlansOk

`func (o *InlineResponse200404DestinationsCriteria) GetApplianceVlansOk() (*[]InlineResponse200404SourcesCriteriaApplianceVlans, bool)`

GetApplianceVlansOk returns a tuple with the ApplianceVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceVlans

`func (o *InlineResponse200404DestinationsCriteria) SetApplianceVlans(v []InlineResponse200404SourcesCriteriaApplianceVlans)`

SetApplianceVlans sets ApplianceVlans field to given value.

### HasApplianceVlans

`func (o *InlineResponse200404DestinationsCriteria) HasApplianceVlans() bool`

HasApplianceVlans returns a boolean if a field has been set.

### GetCountries

`func (o *InlineResponse200404DestinationsCriteria) GetCountries() []InlineResponse200404DestinationsCriteriaCountries`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *InlineResponse200404DestinationsCriteria) GetCountriesOk() (*[]InlineResponse200404DestinationsCriteriaCountries, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *InlineResponse200404DestinationsCriteria) SetCountries(v []InlineResponse200404DestinationsCriteriaCountries)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *InlineResponse200404DestinationsCriteria) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetFqdns

`func (o *InlineResponse200404DestinationsCriteria) GetFqdns() []string`

GetFqdns returns the Fqdns field if non-nil, zero value otherwise.

### GetFqdnsOk

`func (o *InlineResponse200404DestinationsCriteria) GetFqdnsOk() (*[]string, bool)`

GetFqdnsOk returns a tuple with the Fqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdns

`func (o *InlineResponse200404DestinationsCriteria) SetFqdns(v []string)`

SetFqdns sets Fqdns field to given value.

### HasFqdns

`func (o *InlineResponse200404DestinationsCriteria) HasFqdns() bool`

HasFqdns returns a boolean if a field has been set.

### GetSiteSpecificVlans

`func (o *InlineResponse200404DestinationsCriteria) GetSiteSpecificVlans() []InlineResponse200404SourcesCriteriaSiteSpecificVlans`

GetSiteSpecificVlans returns the SiteSpecificVlans field if non-nil, zero value otherwise.

### GetSiteSpecificVlansOk

`func (o *InlineResponse200404DestinationsCriteria) GetSiteSpecificVlansOk() (*[]InlineResponse200404SourcesCriteriaSiteSpecificVlans, bool)`

GetSiteSpecificVlansOk returns a tuple with the SiteSpecificVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteSpecificVlans

`func (o *InlineResponse200404DestinationsCriteria) SetSiteSpecificVlans(v []InlineResponse200404SourcesCriteriaSiteSpecificVlans)`

SetSiteSpecificVlans sets SiteSpecificVlans field to given value.

### HasSiteSpecificVlans

`func (o *InlineResponse200404DestinationsCriteria) HasSiteSpecificVlans() bool`

HasSiteSpecificVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


