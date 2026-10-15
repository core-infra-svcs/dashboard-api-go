# InlineResponse200422SourcesCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressRanges** | Pointer to **[]string** | Address ranges or addresses | [optional] 
**Ports** | Pointer to **[]string** | Port numbers or ranges | [optional] 
**PolicyObjects** | Pointer to [**[]InlineResponse200422SourcesCriteriaPolicyObjects**](InlineResponse200422SourcesCriteriaPolicyObjects.md) | Policy objects | [optional] 
**PolicyObjectGroups** | Pointer to [**[]InlineResponse200422SourcesCriteriaPolicyObjectGroups**](InlineResponse200422SourcesCriteriaPolicyObjectGroups.md) | Policy object groups | [optional] 
**ApplianceVlans** | Pointer to [**[]InlineResponse200422SourcesCriteriaApplianceVlans**](InlineResponse200422SourcesCriteriaApplianceVlans.md) | Appliance VLANs | [optional] 
**SiteSpecificVlans** | Pointer to [**[]InlineResponse200422SourcesCriteriaSiteSpecificVlans**](InlineResponse200422SourcesCriteriaSiteSpecificVlans.md) | Site-specific VLANs with address offset configuration. Maximum 100 items. | [optional] 

## Methods

### NewInlineResponse200422SourcesCriteria

`func NewInlineResponse200422SourcesCriteria() *InlineResponse200422SourcesCriteria`

NewInlineResponse200422SourcesCriteria instantiates a new InlineResponse200422SourcesCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200422SourcesCriteriaWithDefaults

`func NewInlineResponse200422SourcesCriteriaWithDefaults() *InlineResponse200422SourcesCriteria`

NewInlineResponse200422SourcesCriteriaWithDefaults instantiates a new InlineResponse200422SourcesCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressRanges

`func (o *InlineResponse200422SourcesCriteria) GetAddressRanges() []string`

GetAddressRanges returns the AddressRanges field if non-nil, zero value otherwise.

### GetAddressRangesOk

`func (o *InlineResponse200422SourcesCriteria) GetAddressRangesOk() (*[]string, bool)`

GetAddressRangesOk returns a tuple with the AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRanges

`func (o *InlineResponse200422SourcesCriteria) SetAddressRanges(v []string)`

SetAddressRanges sets AddressRanges field to given value.

### HasAddressRanges

`func (o *InlineResponse200422SourcesCriteria) HasAddressRanges() bool`

HasAddressRanges returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200422SourcesCriteria) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200422SourcesCriteria) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200422SourcesCriteria) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200422SourcesCriteria) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *InlineResponse200422SourcesCriteria) GetPolicyObjects() []InlineResponse200422SourcesCriteriaPolicyObjects`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *InlineResponse200422SourcesCriteria) GetPolicyObjectsOk() (*[]InlineResponse200422SourcesCriteriaPolicyObjects, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *InlineResponse200422SourcesCriteria) SetPolicyObjects(v []InlineResponse200422SourcesCriteriaPolicyObjects)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *InlineResponse200422SourcesCriteria) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.

### GetPolicyObjectGroups

`func (o *InlineResponse200422SourcesCriteria) GetPolicyObjectGroups() []InlineResponse200422SourcesCriteriaPolicyObjectGroups`

GetPolicyObjectGroups returns the PolicyObjectGroups field if non-nil, zero value otherwise.

### GetPolicyObjectGroupsOk

`func (o *InlineResponse200422SourcesCriteria) GetPolicyObjectGroupsOk() (*[]InlineResponse200422SourcesCriteriaPolicyObjectGroups, bool)`

GetPolicyObjectGroupsOk returns a tuple with the PolicyObjectGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjectGroups

`func (o *InlineResponse200422SourcesCriteria) SetPolicyObjectGroups(v []InlineResponse200422SourcesCriteriaPolicyObjectGroups)`

SetPolicyObjectGroups sets PolicyObjectGroups field to given value.

### HasPolicyObjectGroups

`func (o *InlineResponse200422SourcesCriteria) HasPolicyObjectGroups() bool`

HasPolicyObjectGroups returns a boolean if a field has been set.

### GetApplianceVlans

`func (o *InlineResponse200422SourcesCriteria) GetApplianceVlans() []InlineResponse200422SourcesCriteriaApplianceVlans`

GetApplianceVlans returns the ApplianceVlans field if non-nil, zero value otherwise.

### GetApplianceVlansOk

`func (o *InlineResponse200422SourcesCriteria) GetApplianceVlansOk() (*[]InlineResponse200422SourcesCriteriaApplianceVlans, bool)`

GetApplianceVlansOk returns a tuple with the ApplianceVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceVlans

`func (o *InlineResponse200422SourcesCriteria) SetApplianceVlans(v []InlineResponse200422SourcesCriteriaApplianceVlans)`

SetApplianceVlans sets ApplianceVlans field to given value.

### HasApplianceVlans

`func (o *InlineResponse200422SourcesCriteria) HasApplianceVlans() bool`

HasApplianceVlans returns a boolean if a field has been set.

### GetSiteSpecificVlans

`func (o *InlineResponse200422SourcesCriteria) GetSiteSpecificVlans() []InlineResponse200422SourcesCriteriaSiteSpecificVlans`

GetSiteSpecificVlans returns the SiteSpecificVlans field if non-nil, zero value otherwise.

### GetSiteSpecificVlansOk

`func (o *InlineResponse200422SourcesCriteria) GetSiteSpecificVlansOk() (*[]InlineResponse200422SourcesCriteriaSiteSpecificVlans, bool)`

GetSiteSpecificVlansOk returns a tuple with the SiteSpecificVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteSpecificVlans

`func (o *InlineResponse200422SourcesCriteria) SetSiteSpecificVlans(v []InlineResponse200422SourcesCriteriaSiteSpecificVlans)`

SetSiteSpecificVlans sets SiteSpecificVlans field to given value.

### HasSiteSpecificVlans

`func (o *InlineResponse200422SourcesCriteria) HasSiteSpecificVlans() bool`

HasSiteSpecificVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


