# InlineResponse200374DestinationsCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressRanges** | Pointer to **[]string** | Address ranges or addresses | [optional] 
**Ports** | Pointer to **[]string** | Port numbers or ranges | [optional] 
**Services** | Pointer to [**[]InlineResponse200374DestinationsCriteriaServices**](InlineResponse200374DestinationsCriteriaServices.md) | Protocol and port services | [optional] 
**ApplicationCategories** | Pointer to [**[]InlineResponse200374DestinationsCriteriaApplicationCategories**](InlineResponse200374DestinationsCriteriaApplicationCategories.md) | Application categories | [optional] 
**Applications** | Pointer to [**[]InlineResponse200374DestinationsCriteriaApplications**](InlineResponse200374DestinationsCriteriaApplications.md) | Applications | [optional] 
**PolicyObjects** | Pointer to [**[]InlineResponse200374SourcesCriteriaPolicyObjects**](InlineResponse200374SourcesCriteriaPolicyObjects.md) | Policy objects | [optional] 
**PolicyObjectGroups** | Pointer to [**[]InlineResponse200374SourcesCriteriaPolicyObjectGroups**](InlineResponse200374SourcesCriteriaPolicyObjectGroups.md) | Policy object groups | [optional] 
**ApplianceVlans** | Pointer to [**[]InlineResponse200374SourcesCriteriaApplianceVlans**](InlineResponse200374SourcesCriteriaApplianceVlans.md) | Appliance VLANs | [optional] 
**Countries** | Pointer to [**[]InlineResponse200374DestinationsCriteriaCountries**](InlineResponse200374DestinationsCriteriaCountries.md) | Countries | [optional] 
**Fqdns** | Pointer to **[]string** | Fully qualified domain names or wildcard patterns | [optional] 
**SiteSpecificVlans** | Pointer to [**[]InlineResponse200374SourcesCriteriaSiteSpecificVlans**](InlineResponse200374SourcesCriteriaSiteSpecificVlans.md) | Site-specific VLANs with address offset configuration. Maximum 100 items. | [optional] 

## Methods

### NewInlineResponse200374DestinationsCriteria

`func NewInlineResponse200374DestinationsCriteria() *InlineResponse200374DestinationsCriteria`

NewInlineResponse200374DestinationsCriteria instantiates a new InlineResponse200374DestinationsCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200374DestinationsCriteriaWithDefaults

`func NewInlineResponse200374DestinationsCriteriaWithDefaults() *InlineResponse200374DestinationsCriteria`

NewInlineResponse200374DestinationsCriteriaWithDefaults instantiates a new InlineResponse200374DestinationsCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressRanges

`func (o *InlineResponse200374DestinationsCriteria) GetAddressRanges() []string`

GetAddressRanges returns the AddressRanges field if non-nil, zero value otherwise.

### GetAddressRangesOk

`func (o *InlineResponse200374DestinationsCriteria) GetAddressRangesOk() (*[]string, bool)`

GetAddressRangesOk returns a tuple with the AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRanges

`func (o *InlineResponse200374DestinationsCriteria) SetAddressRanges(v []string)`

SetAddressRanges sets AddressRanges field to given value.

### HasAddressRanges

`func (o *InlineResponse200374DestinationsCriteria) HasAddressRanges() bool`

HasAddressRanges returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200374DestinationsCriteria) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200374DestinationsCriteria) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200374DestinationsCriteria) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200374DestinationsCriteria) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetServices

`func (o *InlineResponse200374DestinationsCriteria) GetServices() []InlineResponse200374DestinationsCriteriaServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *InlineResponse200374DestinationsCriteria) GetServicesOk() (*[]InlineResponse200374DestinationsCriteriaServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *InlineResponse200374DestinationsCriteria) SetServices(v []InlineResponse200374DestinationsCriteriaServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *InlineResponse200374DestinationsCriteria) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetApplicationCategories

`func (o *InlineResponse200374DestinationsCriteria) GetApplicationCategories() []InlineResponse200374DestinationsCriteriaApplicationCategories`

GetApplicationCategories returns the ApplicationCategories field if non-nil, zero value otherwise.

### GetApplicationCategoriesOk

`func (o *InlineResponse200374DestinationsCriteria) GetApplicationCategoriesOk() (*[]InlineResponse200374DestinationsCriteriaApplicationCategories, bool)`

GetApplicationCategoriesOk returns a tuple with the ApplicationCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCategories

`func (o *InlineResponse200374DestinationsCriteria) SetApplicationCategories(v []InlineResponse200374DestinationsCriteriaApplicationCategories)`

SetApplicationCategories sets ApplicationCategories field to given value.

### HasApplicationCategories

`func (o *InlineResponse200374DestinationsCriteria) HasApplicationCategories() bool`

HasApplicationCategories returns a boolean if a field has been set.

### GetApplications

`func (o *InlineResponse200374DestinationsCriteria) GetApplications() []InlineResponse200374DestinationsCriteriaApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *InlineResponse200374DestinationsCriteria) GetApplicationsOk() (*[]InlineResponse200374DestinationsCriteriaApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *InlineResponse200374DestinationsCriteria) SetApplications(v []InlineResponse200374DestinationsCriteriaApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *InlineResponse200374DestinationsCriteria) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *InlineResponse200374DestinationsCriteria) GetPolicyObjects() []InlineResponse200374SourcesCriteriaPolicyObjects`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *InlineResponse200374DestinationsCriteria) GetPolicyObjectsOk() (*[]InlineResponse200374SourcesCriteriaPolicyObjects, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *InlineResponse200374DestinationsCriteria) SetPolicyObjects(v []InlineResponse200374SourcesCriteriaPolicyObjects)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *InlineResponse200374DestinationsCriteria) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.

### GetPolicyObjectGroups

`func (o *InlineResponse200374DestinationsCriteria) GetPolicyObjectGroups() []InlineResponse200374SourcesCriteriaPolicyObjectGroups`

GetPolicyObjectGroups returns the PolicyObjectGroups field if non-nil, zero value otherwise.

### GetPolicyObjectGroupsOk

`func (o *InlineResponse200374DestinationsCriteria) GetPolicyObjectGroupsOk() (*[]InlineResponse200374SourcesCriteriaPolicyObjectGroups, bool)`

GetPolicyObjectGroupsOk returns a tuple with the PolicyObjectGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjectGroups

`func (o *InlineResponse200374DestinationsCriteria) SetPolicyObjectGroups(v []InlineResponse200374SourcesCriteriaPolicyObjectGroups)`

SetPolicyObjectGroups sets PolicyObjectGroups field to given value.

### HasPolicyObjectGroups

`func (o *InlineResponse200374DestinationsCriteria) HasPolicyObjectGroups() bool`

HasPolicyObjectGroups returns a boolean if a field has been set.

### GetApplianceVlans

`func (o *InlineResponse200374DestinationsCriteria) GetApplianceVlans() []InlineResponse200374SourcesCriteriaApplianceVlans`

GetApplianceVlans returns the ApplianceVlans field if non-nil, zero value otherwise.

### GetApplianceVlansOk

`func (o *InlineResponse200374DestinationsCriteria) GetApplianceVlansOk() (*[]InlineResponse200374SourcesCriteriaApplianceVlans, bool)`

GetApplianceVlansOk returns a tuple with the ApplianceVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceVlans

`func (o *InlineResponse200374DestinationsCriteria) SetApplianceVlans(v []InlineResponse200374SourcesCriteriaApplianceVlans)`

SetApplianceVlans sets ApplianceVlans field to given value.

### HasApplianceVlans

`func (o *InlineResponse200374DestinationsCriteria) HasApplianceVlans() bool`

HasApplianceVlans returns a boolean if a field has been set.

### GetCountries

`func (o *InlineResponse200374DestinationsCriteria) GetCountries() []InlineResponse200374DestinationsCriteriaCountries`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *InlineResponse200374DestinationsCriteria) GetCountriesOk() (*[]InlineResponse200374DestinationsCriteriaCountries, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *InlineResponse200374DestinationsCriteria) SetCountries(v []InlineResponse200374DestinationsCriteriaCountries)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *InlineResponse200374DestinationsCriteria) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetFqdns

`func (o *InlineResponse200374DestinationsCriteria) GetFqdns() []string`

GetFqdns returns the Fqdns field if non-nil, zero value otherwise.

### GetFqdnsOk

`func (o *InlineResponse200374DestinationsCriteria) GetFqdnsOk() (*[]string, bool)`

GetFqdnsOk returns a tuple with the Fqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdns

`func (o *InlineResponse200374DestinationsCriteria) SetFqdns(v []string)`

SetFqdns sets Fqdns field to given value.

### HasFqdns

`func (o *InlineResponse200374DestinationsCriteria) HasFqdns() bool`

HasFqdns returns a boolean if a field has been set.

### GetSiteSpecificVlans

`func (o *InlineResponse200374DestinationsCriteria) GetSiteSpecificVlans() []InlineResponse200374SourcesCriteriaSiteSpecificVlans`

GetSiteSpecificVlans returns the SiteSpecificVlans field if non-nil, zero value otherwise.

### GetSiteSpecificVlansOk

`func (o *InlineResponse200374DestinationsCriteria) GetSiteSpecificVlansOk() (*[]InlineResponse200374SourcesCriteriaSiteSpecificVlans, bool)`

GetSiteSpecificVlansOk returns a tuple with the SiteSpecificVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteSpecificVlans

`func (o *InlineResponse200374DestinationsCriteria) SetSiteSpecificVlans(v []InlineResponse200374SourcesCriteriaSiteSpecificVlans)`

SetSiteSpecificVlans sets SiteSpecificVlans field to given value.

### HasSiteSpecificVlans

`func (o *InlineResponse200374DestinationsCriteria) HasSiteSpecificVlans() bool`

HasSiteSpecificVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


