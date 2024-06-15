# InlineObject70

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custom** | Pointer to [**[]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom**](NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom.md) | Custom VPN exclusion rules. Pass an empty array to clear existing rules. | [optional] 
**MajorApplications** | Pointer to [**[]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications**](NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications.md) | Major Application based VPN exclusion rules. Pass an empty array to clear existing rules. | [optional] 

## Methods

### NewInlineObject70

`func NewInlineObject70() *InlineObject70`

NewInlineObject70 instantiates a new InlineObject70 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject70WithDefaults

`func NewInlineObject70WithDefaults() *InlineObject70`

NewInlineObject70WithDefaults instantiates a new InlineObject70 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustom

`func (o *InlineObject70) GetCustom() []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *InlineObject70) GetCustomOk() (*[]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *InlineObject70) SetCustom(v []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *InlineObject70) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetMajorApplications

`func (o *InlineObject70) GetMajorApplications() []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications`

GetMajorApplications returns the MajorApplications field if non-nil, zero value otherwise.

### GetMajorApplicationsOk

`func (o *InlineObject70) GetMajorApplicationsOk() (*[]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications, bool)`

GetMajorApplicationsOk returns a tuple with the MajorApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorApplications

`func (o *InlineObject70) SetMajorApplications(v []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications)`

SetMajorApplications sets MajorApplications field to given value.

### HasMajorApplications

`func (o *InlineObject70) HasMajorApplications() bool`

HasMajorApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


