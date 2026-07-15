# InlineResponse200134ProductsAppliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to [**InlineResponse200134ProductsWirelessCurrentVersion**](InlineResponse200134ProductsWirelessCurrentVersion.md) |  | [optional] 
**LastUpgrade** | Pointer to [**InlineResponse200134ProductsWirelessLastUpgrade**](InlineResponse200134ProductsWirelessLastUpgrade.md) |  | [optional] 
**NextUpgrade** | Pointer to [**InlineResponse200134ProductsApplianceNextUpgrade**](InlineResponse200134ProductsApplianceNextUpgrade.md) |  | [optional] 
**IsUpgradeAvailable** | Pointer to **bool** | Whether or not an upgraded recommended firmware version is available | [optional] 
**AvailableVersions** | Pointer to [**[]InlineResponse200134ProductsWirelessAvailableVersions**](InlineResponse200134ProductsWirelessAvailableVersions.md) | Firmware versions available for upgrade | [optional] 
**ParticipateInNextBetaRelease** | Pointer to **bool** | Whether or not the network wants beta firmware | [optional] 

## Methods

### NewInlineResponse200134ProductsAppliance

`func NewInlineResponse200134ProductsAppliance() *InlineResponse200134ProductsAppliance`

NewInlineResponse200134ProductsAppliance instantiates a new InlineResponse200134ProductsAppliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200134ProductsApplianceWithDefaults

`func NewInlineResponse200134ProductsApplianceWithDefaults() *InlineResponse200134ProductsAppliance`

NewInlineResponse200134ProductsApplianceWithDefaults instantiates a new InlineResponse200134ProductsAppliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *InlineResponse200134ProductsAppliance) GetCurrentVersion() InlineResponse200134ProductsWirelessCurrentVersion`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *InlineResponse200134ProductsAppliance) GetCurrentVersionOk() (*InlineResponse200134ProductsWirelessCurrentVersion, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *InlineResponse200134ProductsAppliance) SetCurrentVersion(v InlineResponse200134ProductsWirelessCurrentVersion)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *InlineResponse200134ProductsAppliance) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetLastUpgrade

`func (o *InlineResponse200134ProductsAppliance) GetLastUpgrade() InlineResponse200134ProductsWirelessLastUpgrade`

GetLastUpgrade returns the LastUpgrade field if non-nil, zero value otherwise.

### GetLastUpgradeOk

`func (o *InlineResponse200134ProductsAppliance) GetLastUpgradeOk() (*InlineResponse200134ProductsWirelessLastUpgrade, bool)`

GetLastUpgradeOk returns a tuple with the LastUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpgrade

`func (o *InlineResponse200134ProductsAppliance) SetLastUpgrade(v InlineResponse200134ProductsWirelessLastUpgrade)`

SetLastUpgrade sets LastUpgrade field to given value.

### HasLastUpgrade

`func (o *InlineResponse200134ProductsAppliance) HasLastUpgrade() bool`

HasLastUpgrade returns a boolean if a field has been set.

### GetNextUpgrade

`func (o *InlineResponse200134ProductsAppliance) GetNextUpgrade() InlineResponse200134ProductsApplianceNextUpgrade`

GetNextUpgrade returns the NextUpgrade field if non-nil, zero value otherwise.

### GetNextUpgradeOk

`func (o *InlineResponse200134ProductsAppliance) GetNextUpgradeOk() (*InlineResponse200134ProductsApplianceNextUpgrade, bool)`

GetNextUpgradeOk returns a tuple with the NextUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpgrade

`func (o *InlineResponse200134ProductsAppliance) SetNextUpgrade(v InlineResponse200134ProductsApplianceNextUpgrade)`

SetNextUpgrade sets NextUpgrade field to given value.

### HasNextUpgrade

`func (o *InlineResponse200134ProductsAppliance) HasNextUpgrade() bool`

HasNextUpgrade returns a boolean if a field has been set.

### GetIsUpgradeAvailable

`func (o *InlineResponse200134ProductsAppliance) GetIsUpgradeAvailable() bool`

GetIsUpgradeAvailable returns the IsUpgradeAvailable field if non-nil, zero value otherwise.

### GetIsUpgradeAvailableOk

`func (o *InlineResponse200134ProductsAppliance) GetIsUpgradeAvailableOk() (*bool, bool)`

GetIsUpgradeAvailableOk returns a tuple with the IsUpgradeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgradeAvailable

`func (o *InlineResponse200134ProductsAppliance) SetIsUpgradeAvailable(v bool)`

SetIsUpgradeAvailable sets IsUpgradeAvailable field to given value.

### HasIsUpgradeAvailable

`func (o *InlineResponse200134ProductsAppliance) HasIsUpgradeAvailable() bool`

HasIsUpgradeAvailable returns a boolean if a field has been set.

### GetAvailableVersions

`func (o *InlineResponse200134ProductsAppliance) GetAvailableVersions() []InlineResponse200134ProductsWirelessAvailableVersions`

GetAvailableVersions returns the AvailableVersions field if non-nil, zero value otherwise.

### GetAvailableVersionsOk

`func (o *InlineResponse200134ProductsAppliance) GetAvailableVersionsOk() (*[]InlineResponse200134ProductsWirelessAvailableVersions, bool)`

GetAvailableVersionsOk returns a tuple with the AvailableVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVersions

`func (o *InlineResponse200134ProductsAppliance) SetAvailableVersions(v []InlineResponse200134ProductsWirelessAvailableVersions)`

SetAvailableVersions sets AvailableVersions field to given value.

### HasAvailableVersions

`func (o *InlineResponse200134ProductsAppliance) HasAvailableVersions() bool`

HasAvailableVersions returns a boolean if a field has been set.

### GetParticipateInNextBetaRelease

`func (o *InlineResponse200134ProductsAppliance) GetParticipateInNextBetaRelease() bool`

GetParticipateInNextBetaRelease returns the ParticipateInNextBetaRelease field if non-nil, zero value otherwise.

### GetParticipateInNextBetaReleaseOk

`func (o *InlineResponse200134ProductsAppliance) GetParticipateInNextBetaReleaseOk() (*bool, bool)`

GetParticipateInNextBetaReleaseOk returns a tuple with the ParticipateInNextBetaRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipateInNextBetaRelease

`func (o *InlineResponse200134ProductsAppliance) SetParticipateInNextBetaRelease(v bool)`

SetParticipateInNextBetaRelease sets ParticipateInNextBetaRelease field to given value.

### HasParticipateInNextBetaRelease

`func (o *InlineResponse200134ProductsAppliance) HasParticipateInNextBetaRelease() bool`

HasParticipateInNextBetaRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


