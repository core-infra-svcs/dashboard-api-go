# InlineResponse200389DestinationsCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressRanges** | Pointer to **[]string** | Address ranges or addresses | [optional] 
**Ports** | Pointer to **[]string** | Port numbers or ranges | [optional] 
**Services** | Pointer to [**[]InlineResponse200389DestinationsCriteriaServices**](InlineResponse200389DestinationsCriteriaServices.md) | Protocol and port services | [optional] 
**ApplicationCategories** | Pointer to [**[]InlineResponse200389DestinationsCriteriaApplicationCategories**](InlineResponse200389DestinationsCriteriaApplicationCategories.md) | Application categories | [optional] 
**Applications** | Pointer to [**[]InlineResponse200389DestinationsCriteriaApplications**](InlineResponse200389DestinationsCriteriaApplications.md) | Applications | [optional] 
**PolicyObjects** | Pointer to [**[]InlineResponse200389SourcesCriteriaPolicyObjects**](InlineResponse200389SourcesCriteriaPolicyObjects.md) | Policy objects | [optional] 
**PolicyObjectGroups** | Pointer to [**[]InlineResponse200389SourcesCriteriaPolicyObjectGroups**](InlineResponse200389SourcesCriteriaPolicyObjectGroups.md) | Policy object groups | [optional] 
**ApplianceVlans** | Pointer to [**[]InlineResponse200389SourcesCriteriaApplianceVlans**](InlineResponse200389SourcesCriteriaApplianceVlans.md) | Appliance VLANs | [optional] 
**Countries** | Pointer to [**[]InlineResponse200389DestinationsCriteriaCountries**](InlineResponse200389DestinationsCriteriaCountries.md) | Countries | [optional] 
**Fqdns** | Pointer to **[]string** | Fully qualified domain names or wildcard patterns | [optional] 
**SiteSpecificVlans** | Pointer to [**[]InlineResponse200389SourcesCriteriaSiteSpecificVlans**](InlineResponse200389SourcesCriteriaSiteSpecificVlans.md) | Site-specific VLANs with address offset configuration. Maximum 100 items. | [optional] 

## Methods

### NewInlineResponse200389DestinationsCriteria

`func NewInlineResponse200389DestinationsCriteria() *InlineResponse200389DestinationsCriteria`

NewInlineResponse200389DestinationsCriteria instantiates a new InlineResponse200389DestinationsCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200389DestinationsCriteriaWithDefaults

`func NewInlineResponse200389DestinationsCriteriaWithDefaults() *InlineResponse200389DestinationsCriteria`

NewInlineResponse200389DestinationsCriteriaWithDefaults instantiates a new InlineResponse200389DestinationsCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressRanges

`func (o *InlineResponse200389DestinationsCriteria) GetAddressRanges() []string`

GetAddressRanges returns the AddressRanges field if non-nil, zero value otherwise.

### GetAddressRangesOk

`func (o *InlineResponse200389DestinationsCriteria) GetAddressRangesOk() (*[]string, bool)`

GetAddressRangesOk returns a tuple with the AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRanges

`func (o *InlineResponse200389DestinationsCriteria) SetAddressRanges(v []string)`

SetAddressRanges sets AddressRanges field to given value.

### HasAddressRanges

`func (o *InlineResponse200389DestinationsCriteria) HasAddressRanges() bool`

HasAddressRanges returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200389DestinationsCriteria) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200389DestinationsCriteria) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200389DestinationsCriteria) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200389DestinationsCriteria) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetServices

`func (o *InlineResponse200389DestinationsCriteria) GetServices() []InlineResponse200389DestinationsCriteriaServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *InlineResponse200389DestinationsCriteria) GetServicesOk() (*[]InlineResponse200389DestinationsCriteriaServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *InlineResponse200389DestinationsCriteria) SetServices(v []InlineResponse200389DestinationsCriteriaServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *InlineResponse200389DestinationsCriteria) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetApplicationCategories

`func (o *InlineResponse200389DestinationsCriteria) GetApplicationCategories() []InlineResponse200389DestinationsCriteriaApplicationCategories`

GetApplicationCategories returns the ApplicationCategories field if non-nil, zero value otherwise.

### GetApplicationCategoriesOk

`func (o *InlineResponse200389DestinationsCriteria) GetApplicationCategoriesOk() (*[]InlineResponse200389DestinationsCriteriaApplicationCategories, bool)`

GetApplicationCategoriesOk returns a tuple with the ApplicationCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCategories

`func (o *InlineResponse200389DestinationsCriteria) SetApplicationCategories(v []InlineResponse200389DestinationsCriteriaApplicationCategories)`

SetApplicationCategories sets ApplicationCategories field to given value.

### HasApplicationCategories

`func (o *InlineResponse200389DestinationsCriteria) HasApplicationCategories() bool`

HasApplicationCategories returns a boolean if a field has been set.

### GetApplications

`func (o *InlineResponse200389DestinationsCriteria) GetApplications() []InlineResponse200389DestinationsCriteriaApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *InlineResponse200389DestinationsCriteria) GetApplicationsOk() (*[]InlineResponse200389DestinationsCriteriaApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *InlineResponse200389DestinationsCriteria) SetApplications(v []InlineResponse200389DestinationsCriteriaApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *InlineResponse200389DestinationsCriteria) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *InlineResponse200389DestinationsCriteria) GetPolicyObjects() []InlineResponse200389SourcesCriteriaPolicyObjects`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *InlineResponse200389DestinationsCriteria) GetPolicyObjectsOk() (*[]InlineResponse200389SourcesCriteriaPolicyObjects, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *InlineResponse200389DestinationsCriteria) SetPolicyObjects(v []InlineResponse200389SourcesCriteriaPolicyObjects)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *InlineResponse200389DestinationsCriteria) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.

### GetPolicyObjectGroups

`func (o *InlineResponse200389DestinationsCriteria) GetPolicyObjectGroups() []InlineResponse200389SourcesCriteriaPolicyObjectGroups`

GetPolicyObjectGroups returns the PolicyObjectGroups field if non-nil, zero value otherwise.

### GetPolicyObjectGroupsOk

`func (o *InlineResponse200389DestinationsCriteria) GetPolicyObjectGroupsOk() (*[]InlineResponse200389SourcesCriteriaPolicyObjectGroups, bool)`

GetPolicyObjectGroupsOk returns a tuple with the PolicyObjectGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjectGroups

`func (o *InlineResponse200389DestinationsCriteria) SetPolicyObjectGroups(v []InlineResponse200389SourcesCriteriaPolicyObjectGroups)`

SetPolicyObjectGroups sets PolicyObjectGroups field to given value.

### HasPolicyObjectGroups

`func (o *InlineResponse200389DestinationsCriteria) HasPolicyObjectGroups() bool`

HasPolicyObjectGroups returns a boolean if a field has been set.

### GetApplianceVlans

`func (o *InlineResponse200389DestinationsCriteria) GetApplianceVlans() []InlineResponse200389SourcesCriteriaApplianceVlans`

GetApplianceVlans returns the ApplianceVlans field if non-nil, zero value otherwise.

### GetApplianceVlansOk

`func (o *InlineResponse200389DestinationsCriteria) GetApplianceVlansOk() (*[]InlineResponse200389SourcesCriteriaApplianceVlans, bool)`

GetApplianceVlansOk returns a tuple with the ApplianceVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceVlans

`func (o *InlineResponse200389DestinationsCriteria) SetApplianceVlans(v []InlineResponse200389SourcesCriteriaApplianceVlans)`

SetApplianceVlans sets ApplianceVlans field to given value.

### HasApplianceVlans

`func (o *InlineResponse200389DestinationsCriteria) HasApplianceVlans() bool`

HasApplianceVlans returns a boolean if a field has been set.

### GetCountries

`func (o *InlineResponse200389DestinationsCriteria) GetCountries() []InlineResponse200389DestinationsCriteriaCountries`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *InlineResponse200389DestinationsCriteria) GetCountriesOk() (*[]InlineResponse200389DestinationsCriteriaCountries, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *InlineResponse200389DestinationsCriteria) SetCountries(v []InlineResponse200389DestinationsCriteriaCountries)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *InlineResponse200389DestinationsCriteria) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetFqdns

`func (o *InlineResponse200389DestinationsCriteria) GetFqdns() []string`

GetFqdns returns the Fqdns field if non-nil, zero value otherwise.

### GetFqdnsOk

`func (o *InlineResponse200389DestinationsCriteria) GetFqdnsOk() (*[]string, bool)`

GetFqdnsOk returns a tuple with the Fqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdns

`func (o *InlineResponse200389DestinationsCriteria) SetFqdns(v []string)`

SetFqdns sets Fqdns field to given value.

### HasFqdns

`func (o *InlineResponse200389DestinationsCriteria) HasFqdns() bool`

HasFqdns returns a boolean if a field has been set.

### GetSiteSpecificVlans

`func (o *InlineResponse200389DestinationsCriteria) GetSiteSpecificVlans() []InlineResponse200389SourcesCriteriaSiteSpecificVlans`

GetSiteSpecificVlans returns the SiteSpecificVlans field if non-nil, zero value otherwise.

### GetSiteSpecificVlansOk

`func (o *InlineResponse200389DestinationsCriteria) GetSiteSpecificVlansOk() (*[]InlineResponse200389SourcesCriteriaSiteSpecificVlans, bool)`

GetSiteSpecificVlansOk returns a tuple with the SiteSpecificVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteSpecificVlans

`func (o *InlineResponse200389DestinationsCriteria) SetSiteSpecificVlans(v []InlineResponse200389SourcesCriteriaSiteSpecificVlans)`

SetSiteSpecificVlans sets SiteSpecificVlans field to given value.

### HasSiteSpecificVlans

`func (o *InlineResponse200389DestinationsCriteria) HasSiteSpecificVlans() bool`

HasSiteSpecificVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


