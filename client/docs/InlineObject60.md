# InlineObject60

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custom** | Pointer to [**[]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom**](NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom.md) | Custom VPN exclusion rules. Pass an empty array to clear existing rules. | [optional] 
**MajorApplications** | Pointer to [**[]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications**](NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications.md) | Major Application based VPN exclusion rules. Pass an empty array to clear existing rules. | [optional] 

## Methods

### NewInlineObject60

`func NewInlineObject60() *InlineObject60`

NewInlineObject60 instantiates a new InlineObject60 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject60WithDefaults

`func NewInlineObject60WithDefaults() *InlineObject60`

NewInlineObject60WithDefaults instantiates a new InlineObject60 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustom

`func (o *InlineObject60) GetCustom() []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *InlineObject60) GetCustomOk() (*[]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *InlineObject60) SetCustom(v []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *InlineObject60) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetMajorApplications

`func (o *InlineObject60) GetMajorApplications() []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications`

GetMajorApplications returns the MajorApplications field if non-nil, zero value otherwise.

### GetMajorApplicationsOk

`func (o *InlineObject60) GetMajorApplicationsOk() (*[]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications, bool)`

GetMajorApplicationsOk returns a tuple with the MajorApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorApplications

`func (o *InlineObject60) SetMajorApplications(v []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications)`

SetMajorApplications sets MajorApplications field to given value.

### HasMajorApplications

`func (o *InlineObject60) HasMajorApplications() bool`

HasMajorApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


