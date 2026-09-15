# InlineObject86

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custom** | Pointer to [**[]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom**](NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom.md) | Custom VPN exclusion rules. Pass an empty array to clear existing rules. | [optional] 
**MajorApplications** | Pointer to [**[]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications**](NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications.md) | Major Application based VPN exclusion rules. Pass an empty array to clear existing rules. | [optional] 

## Methods

### NewInlineObject86

`func NewInlineObject86() *InlineObject86`

NewInlineObject86 instantiates a new InlineObject86 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject86WithDefaults

`func NewInlineObject86WithDefaults() *InlineObject86`

NewInlineObject86WithDefaults instantiates a new InlineObject86 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustom

`func (o *InlineObject86) GetCustom() []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *InlineObject86) GetCustomOk() (*[]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *InlineObject86) SetCustom(v []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *InlineObject86) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetMajorApplications

`func (o *InlineObject86) GetMajorApplications() []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications`

GetMajorApplications returns the MajorApplications field if non-nil, zero value otherwise.

### GetMajorApplicationsOk

`func (o *InlineObject86) GetMajorApplicationsOk() (*[]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications, bool)`

GetMajorApplicationsOk returns a tuple with the MajorApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorApplications

`func (o *InlineObject86) SetMajorApplications(v []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications)`

SetMajorApplications sets MajorApplications field to given value.

### HasMajorApplications

`func (o *InlineObject86) HasMajorApplications() bool`

HasMajorApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


