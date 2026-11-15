# InlineObject90

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custom** | Pointer to [**[]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom**](NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom.md) | Custom VPN exclusion rules. Pass an empty array to clear existing rules. | [optional] 
**MajorApplications** | Pointer to [**[]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications**](NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications.md) | Major Application based VPN exclusion rules. Pass an empty array to clear existing rules. | [optional] 
**ApplicationCategoryMapping** | Pointer to [**NetworksNetworkIdApplianceTrafficShapingVpnExclusionsApplicationCategoryMapping**](NetworksNetworkIdApplianceTrafficShapingVpnExclusionsApplicationCategoryMapping.md) |  | [optional] 

## Methods

### NewInlineObject90

`func NewInlineObject90() *InlineObject90`

NewInlineObject90 instantiates a new InlineObject90 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject90WithDefaults

`func NewInlineObject90WithDefaults() *InlineObject90`

NewInlineObject90WithDefaults instantiates a new InlineObject90 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustom

`func (o *InlineObject90) GetCustom() []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *InlineObject90) GetCustomOk() (*[]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *InlineObject90) SetCustom(v []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *InlineObject90) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetMajorApplications

`func (o *InlineObject90) GetMajorApplications() []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications`

GetMajorApplications returns the MajorApplications field if non-nil, zero value otherwise.

### GetMajorApplicationsOk

`func (o *InlineObject90) GetMajorApplicationsOk() (*[]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications, bool)`

GetMajorApplicationsOk returns a tuple with the MajorApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorApplications

`func (o *InlineObject90) SetMajorApplications(v []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications)`

SetMajorApplications sets MajorApplications field to given value.

### HasMajorApplications

`func (o *InlineObject90) HasMajorApplications() bool`

HasMajorApplications returns a boolean if a field has been set.

### GetApplicationCategoryMapping

`func (o *InlineObject90) GetApplicationCategoryMapping() NetworksNetworkIdApplianceTrafficShapingVpnExclusionsApplicationCategoryMapping`

GetApplicationCategoryMapping returns the ApplicationCategoryMapping field if non-nil, zero value otherwise.

### GetApplicationCategoryMappingOk

`func (o *InlineObject90) GetApplicationCategoryMappingOk() (*NetworksNetworkIdApplianceTrafficShapingVpnExclusionsApplicationCategoryMapping, bool)`

GetApplicationCategoryMappingOk returns a tuple with the ApplicationCategoryMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCategoryMapping

`func (o *InlineObject90) SetApplicationCategoryMapping(v NetworksNetworkIdApplianceTrafficShapingVpnExclusionsApplicationCategoryMapping)`

SetApplicationCategoryMapping sets ApplicationCategoryMapping field to given value.

### HasApplicationCategoryMapping

`func (o *InlineObject90) HasApplicationCategoryMapping() bool`

HasApplicationCategoryMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


