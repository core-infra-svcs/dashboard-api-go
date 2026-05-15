# InlineResponse200125ProductsAppliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to [**InlineResponse200125ProductsWirelessCurrentVersion**](InlineResponse200125ProductsWirelessCurrentVersion.md) |  | [optional] 
**LastUpgrade** | Pointer to [**InlineResponse200125ProductsWirelessLastUpgrade**](InlineResponse200125ProductsWirelessLastUpgrade.md) |  | [optional] 
**NextUpgrade** | Pointer to [**InlineResponse200125ProductsApplianceNextUpgrade**](InlineResponse200125ProductsApplianceNextUpgrade.md) |  | [optional] 
**IsUpgradeAvailable** | Pointer to **bool** | Whether or not an upgraded recommended firmware version is available | [optional] 
**AvailableVersions** | Pointer to [**[]InlineResponse200125ProductsWirelessAvailableVersions**](InlineResponse200125ProductsWirelessAvailableVersions.md) | Firmware versions available for upgrade | [optional] 
**ParticipateInNextBetaRelease** | Pointer to **bool** | Whether or not the network wants beta firmware | [optional] 

## Methods

### NewInlineResponse200125ProductsAppliance

`func NewInlineResponse200125ProductsAppliance() *InlineResponse200125ProductsAppliance`

NewInlineResponse200125ProductsAppliance instantiates a new InlineResponse200125ProductsAppliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200125ProductsApplianceWithDefaults

`func NewInlineResponse200125ProductsApplianceWithDefaults() *InlineResponse200125ProductsAppliance`

NewInlineResponse200125ProductsApplianceWithDefaults instantiates a new InlineResponse200125ProductsAppliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *InlineResponse200125ProductsAppliance) GetCurrentVersion() InlineResponse200125ProductsWirelessCurrentVersion`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *InlineResponse200125ProductsAppliance) GetCurrentVersionOk() (*InlineResponse200125ProductsWirelessCurrentVersion, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *InlineResponse200125ProductsAppliance) SetCurrentVersion(v InlineResponse200125ProductsWirelessCurrentVersion)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *InlineResponse200125ProductsAppliance) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetLastUpgrade

`func (o *InlineResponse200125ProductsAppliance) GetLastUpgrade() InlineResponse200125ProductsWirelessLastUpgrade`

GetLastUpgrade returns the LastUpgrade field if non-nil, zero value otherwise.

### GetLastUpgradeOk

`func (o *InlineResponse200125ProductsAppliance) GetLastUpgradeOk() (*InlineResponse200125ProductsWirelessLastUpgrade, bool)`

GetLastUpgradeOk returns a tuple with the LastUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpgrade

`func (o *InlineResponse200125ProductsAppliance) SetLastUpgrade(v InlineResponse200125ProductsWirelessLastUpgrade)`

SetLastUpgrade sets LastUpgrade field to given value.

### HasLastUpgrade

`func (o *InlineResponse200125ProductsAppliance) HasLastUpgrade() bool`

HasLastUpgrade returns a boolean if a field has been set.

### GetNextUpgrade

`func (o *InlineResponse200125ProductsAppliance) GetNextUpgrade() InlineResponse200125ProductsApplianceNextUpgrade`

GetNextUpgrade returns the NextUpgrade field if non-nil, zero value otherwise.

### GetNextUpgradeOk

`func (o *InlineResponse200125ProductsAppliance) GetNextUpgradeOk() (*InlineResponse200125ProductsApplianceNextUpgrade, bool)`

GetNextUpgradeOk returns a tuple with the NextUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpgrade

`func (o *InlineResponse200125ProductsAppliance) SetNextUpgrade(v InlineResponse200125ProductsApplianceNextUpgrade)`

SetNextUpgrade sets NextUpgrade field to given value.

### HasNextUpgrade

`func (o *InlineResponse200125ProductsAppliance) HasNextUpgrade() bool`

HasNextUpgrade returns a boolean if a field has been set.

### GetIsUpgradeAvailable

`func (o *InlineResponse200125ProductsAppliance) GetIsUpgradeAvailable() bool`

GetIsUpgradeAvailable returns the IsUpgradeAvailable field if non-nil, zero value otherwise.

### GetIsUpgradeAvailableOk

`func (o *InlineResponse200125ProductsAppliance) GetIsUpgradeAvailableOk() (*bool, bool)`

GetIsUpgradeAvailableOk returns a tuple with the IsUpgradeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgradeAvailable

`func (o *InlineResponse200125ProductsAppliance) SetIsUpgradeAvailable(v bool)`

SetIsUpgradeAvailable sets IsUpgradeAvailable field to given value.

### HasIsUpgradeAvailable

`func (o *InlineResponse200125ProductsAppliance) HasIsUpgradeAvailable() bool`

HasIsUpgradeAvailable returns a boolean if a field has been set.

### GetAvailableVersions

`func (o *InlineResponse200125ProductsAppliance) GetAvailableVersions() []InlineResponse200125ProductsWirelessAvailableVersions`

GetAvailableVersions returns the AvailableVersions field if non-nil, zero value otherwise.

### GetAvailableVersionsOk

`func (o *InlineResponse200125ProductsAppliance) GetAvailableVersionsOk() (*[]InlineResponse200125ProductsWirelessAvailableVersions, bool)`

GetAvailableVersionsOk returns a tuple with the AvailableVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVersions

`func (o *InlineResponse200125ProductsAppliance) SetAvailableVersions(v []InlineResponse200125ProductsWirelessAvailableVersions)`

SetAvailableVersions sets AvailableVersions field to given value.

### HasAvailableVersions

`func (o *InlineResponse200125ProductsAppliance) HasAvailableVersions() bool`

HasAvailableVersions returns a boolean if a field has been set.

### GetParticipateInNextBetaRelease

`func (o *InlineResponse200125ProductsAppliance) GetParticipateInNextBetaRelease() bool`

GetParticipateInNextBetaRelease returns the ParticipateInNextBetaRelease field if non-nil, zero value otherwise.

### GetParticipateInNextBetaReleaseOk

`func (o *InlineResponse200125ProductsAppliance) GetParticipateInNextBetaReleaseOk() (*bool, bool)`

GetParticipateInNextBetaReleaseOk returns a tuple with the ParticipateInNextBetaRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipateInNextBetaRelease

`func (o *InlineResponse200125ProductsAppliance) SetParticipateInNextBetaRelease(v bool)`

SetParticipateInNextBetaRelease sets ParticipateInNextBetaRelease field to given value.

### HasParticipateInNextBetaRelease

`func (o *InlineResponse200125ProductsAppliance) HasParticipateInNextBetaRelease() bool`

HasParticipateInNextBetaRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


