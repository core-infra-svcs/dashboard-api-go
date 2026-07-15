# InlineResponse200389SourcesCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressRanges** | Pointer to **[]string** | Address ranges or addresses | [optional] 
**Ports** | Pointer to **[]string** | Port numbers or ranges | [optional] 
**PolicyObjects** | Pointer to [**[]InlineResponse200389SourcesCriteriaPolicyObjects**](InlineResponse200389SourcesCriteriaPolicyObjects.md) | Policy objects | [optional] 
**PolicyObjectGroups** | Pointer to [**[]InlineResponse200389SourcesCriteriaPolicyObjectGroups**](InlineResponse200389SourcesCriteriaPolicyObjectGroups.md) | Policy object groups | [optional] 
**ApplianceVlans** | Pointer to [**[]InlineResponse200389SourcesCriteriaApplianceVlans**](InlineResponse200389SourcesCriteriaApplianceVlans.md) | Appliance VLANs | [optional] 
**SiteSpecificVlans** | Pointer to [**[]InlineResponse200389SourcesCriteriaSiteSpecificVlans**](InlineResponse200389SourcesCriteriaSiteSpecificVlans.md) | Site-specific VLANs with address offset configuration. Maximum 100 items. | [optional] 

## Methods

### NewInlineResponse200389SourcesCriteria

`func NewInlineResponse200389SourcesCriteria() *InlineResponse200389SourcesCriteria`

NewInlineResponse200389SourcesCriteria instantiates a new InlineResponse200389SourcesCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200389SourcesCriteriaWithDefaults

`func NewInlineResponse200389SourcesCriteriaWithDefaults() *InlineResponse200389SourcesCriteria`

NewInlineResponse200389SourcesCriteriaWithDefaults instantiates a new InlineResponse200389SourcesCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressRanges

`func (o *InlineResponse200389SourcesCriteria) GetAddressRanges() []string`

GetAddressRanges returns the AddressRanges field if non-nil, zero value otherwise.

### GetAddressRangesOk

`func (o *InlineResponse200389SourcesCriteria) GetAddressRangesOk() (*[]string, bool)`

GetAddressRangesOk returns a tuple with the AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRanges

`func (o *InlineResponse200389SourcesCriteria) SetAddressRanges(v []string)`

SetAddressRanges sets AddressRanges field to given value.

### HasAddressRanges

`func (o *InlineResponse200389SourcesCriteria) HasAddressRanges() bool`

HasAddressRanges returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200389SourcesCriteria) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200389SourcesCriteria) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200389SourcesCriteria) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200389SourcesCriteria) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *InlineResponse200389SourcesCriteria) GetPolicyObjects() []InlineResponse200389SourcesCriteriaPolicyObjects`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *InlineResponse200389SourcesCriteria) GetPolicyObjectsOk() (*[]InlineResponse200389SourcesCriteriaPolicyObjects, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *InlineResponse200389SourcesCriteria) SetPolicyObjects(v []InlineResponse200389SourcesCriteriaPolicyObjects)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *InlineResponse200389SourcesCriteria) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.

### GetPolicyObjectGroups

`func (o *InlineResponse200389SourcesCriteria) GetPolicyObjectGroups() []InlineResponse200389SourcesCriteriaPolicyObjectGroups`

GetPolicyObjectGroups returns the PolicyObjectGroups field if non-nil, zero value otherwise.

### GetPolicyObjectGroupsOk

`func (o *InlineResponse200389SourcesCriteria) GetPolicyObjectGroupsOk() (*[]InlineResponse200389SourcesCriteriaPolicyObjectGroups, bool)`

GetPolicyObjectGroupsOk returns a tuple with the PolicyObjectGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjectGroups

`func (o *InlineResponse200389SourcesCriteria) SetPolicyObjectGroups(v []InlineResponse200389SourcesCriteriaPolicyObjectGroups)`

SetPolicyObjectGroups sets PolicyObjectGroups field to given value.

### HasPolicyObjectGroups

`func (o *InlineResponse200389SourcesCriteria) HasPolicyObjectGroups() bool`

HasPolicyObjectGroups returns a boolean if a field has been set.

### GetApplianceVlans

`func (o *InlineResponse200389SourcesCriteria) GetApplianceVlans() []InlineResponse200389SourcesCriteriaApplianceVlans`

GetApplianceVlans returns the ApplianceVlans field if non-nil, zero value otherwise.

### GetApplianceVlansOk

`func (o *InlineResponse200389SourcesCriteria) GetApplianceVlansOk() (*[]InlineResponse200389SourcesCriteriaApplianceVlans, bool)`

GetApplianceVlansOk returns a tuple with the ApplianceVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceVlans

`func (o *InlineResponse200389SourcesCriteria) SetApplianceVlans(v []InlineResponse200389SourcesCriteriaApplianceVlans)`

SetApplianceVlans sets ApplianceVlans field to given value.

### HasApplianceVlans

`func (o *InlineResponse200389SourcesCriteria) HasApplianceVlans() bool`

HasApplianceVlans returns a boolean if a field has been set.

### GetSiteSpecificVlans

`func (o *InlineResponse200389SourcesCriteria) GetSiteSpecificVlans() []InlineResponse200389SourcesCriteriaSiteSpecificVlans`

GetSiteSpecificVlans returns the SiteSpecificVlans field if non-nil, zero value otherwise.

### GetSiteSpecificVlansOk

`func (o *InlineResponse200389SourcesCriteria) GetSiteSpecificVlansOk() (*[]InlineResponse200389SourcesCriteriaSiteSpecificVlans, bool)`

GetSiteSpecificVlansOk returns a tuple with the SiteSpecificVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteSpecificVlans

`func (o *InlineResponse200389SourcesCriteria) SetSiteSpecificVlans(v []InlineResponse200389SourcesCriteriaSiteSpecificVlans)`

SetSiteSpecificVlans sets SiteSpecificVlans field to given value.

### HasSiteSpecificVlans

`func (o *InlineResponse200389SourcesCriteria) HasSiteSpecificVlans() bool`

HasSiteSpecificVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


