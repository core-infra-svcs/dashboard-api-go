# InlineResponse20097ProductsWireless

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to [**InlineResponse20097ProductsWirelessCurrentVersion**](InlineResponse20097ProductsWirelessCurrentVersion.md) |  | [optional] 
**LastUpgrade** | Pointer to [**InlineResponse20097ProductsWirelessLastUpgrade**](InlineResponse20097ProductsWirelessLastUpgrade.md) |  | [optional] 
**NextUpgrade** | Pointer to [**InlineResponse20097ProductsWirelessNextUpgrade**](InlineResponse20097ProductsWirelessNextUpgrade.md) |  | [optional] 
**IsUpgradeAvailable** | Pointer to **bool** | Whether or not an upgraded firmware version is available | [optional] 
**AvailableVersions** | Pointer to [**[]InlineResponse20097ProductsWirelessAvailableVersions**](InlineResponse20097ProductsWirelessAvailableVersions.md) | Firmware versions available for upgrade | [optional] 
**ParticipateInNextBetaRelease** | Pointer to **bool** | Whether or not the network wants beta firmware | [optional] 

## Methods

### NewInlineResponse20097ProductsWireless

`func NewInlineResponse20097ProductsWireless() *InlineResponse20097ProductsWireless`

NewInlineResponse20097ProductsWireless instantiates a new InlineResponse20097ProductsWireless object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20097ProductsWirelessWithDefaults

`func NewInlineResponse20097ProductsWirelessWithDefaults() *InlineResponse20097ProductsWireless`

NewInlineResponse20097ProductsWirelessWithDefaults instantiates a new InlineResponse20097ProductsWireless object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *InlineResponse20097ProductsWireless) GetCurrentVersion() InlineResponse20097ProductsWirelessCurrentVersion`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *InlineResponse20097ProductsWireless) GetCurrentVersionOk() (*InlineResponse20097ProductsWirelessCurrentVersion, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *InlineResponse20097ProductsWireless) SetCurrentVersion(v InlineResponse20097ProductsWirelessCurrentVersion)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *InlineResponse20097ProductsWireless) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetLastUpgrade

`func (o *InlineResponse20097ProductsWireless) GetLastUpgrade() InlineResponse20097ProductsWirelessLastUpgrade`

GetLastUpgrade returns the LastUpgrade field if non-nil, zero value otherwise.

### GetLastUpgradeOk

`func (o *InlineResponse20097ProductsWireless) GetLastUpgradeOk() (*InlineResponse20097ProductsWirelessLastUpgrade, bool)`

GetLastUpgradeOk returns a tuple with the LastUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpgrade

`func (o *InlineResponse20097ProductsWireless) SetLastUpgrade(v InlineResponse20097ProductsWirelessLastUpgrade)`

SetLastUpgrade sets LastUpgrade field to given value.

### HasLastUpgrade

`func (o *InlineResponse20097ProductsWireless) HasLastUpgrade() bool`

HasLastUpgrade returns a boolean if a field has been set.

### GetNextUpgrade

`func (o *InlineResponse20097ProductsWireless) GetNextUpgrade() InlineResponse20097ProductsWirelessNextUpgrade`

GetNextUpgrade returns the NextUpgrade field if non-nil, zero value otherwise.

### GetNextUpgradeOk

`func (o *InlineResponse20097ProductsWireless) GetNextUpgradeOk() (*InlineResponse20097ProductsWirelessNextUpgrade, bool)`

GetNextUpgradeOk returns a tuple with the NextUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpgrade

`func (o *InlineResponse20097ProductsWireless) SetNextUpgrade(v InlineResponse20097ProductsWirelessNextUpgrade)`

SetNextUpgrade sets NextUpgrade field to given value.

### HasNextUpgrade

`func (o *InlineResponse20097ProductsWireless) HasNextUpgrade() bool`

HasNextUpgrade returns a boolean if a field has been set.

### GetIsUpgradeAvailable

`func (o *InlineResponse20097ProductsWireless) GetIsUpgradeAvailable() bool`

GetIsUpgradeAvailable returns the IsUpgradeAvailable field if non-nil, zero value otherwise.

### GetIsUpgradeAvailableOk

`func (o *InlineResponse20097ProductsWireless) GetIsUpgradeAvailableOk() (*bool, bool)`

GetIsUpgradeAvailableOk returns a tuple with the IsUpgradeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgradeAvailable

`func (o *InlineResponse20097ProductsWireless) SetIsUpgradeAvailable(v bool)`

SetIsUpgradeAvailable sets IsUpgradeAvailable field to given value.

### HasIsUpgradeAvailable

`func (o *InlineResponse20097ProductsWireless) HasIsUpgradeAvailable() bool`

HasIsUpgradeAvailable returns a boolean if a field has been set.

### GetAvailableVersions

`func (o *InlineResponse20097ProductsWireless) GetAvailableVersions() []InlineResponse20097ProductsWirelessAvailableVersions`

GetAvailableVersions returns the AvailableVersions field if non-nil, zero value otherwise.

### GetAvailableVersionsOk

`func (o *InlineResponse20097ProductsWireless) GetAvailableVersionsOk() (*[]InlineResponse20097ProductsWirelessAvailableVersions, bool)`

GetAvailableVersionsOk returns a tuple with the AvailableVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVersions

`func (o *InlineResponse20097ProductsWireless) SetAvailableVersions(v []InlineResponse20097ProductsWirelessAvailableVersions)`

SetAvailableVersions sets AvailableVersions field to given value.

### HasAvailableVersions

`func (o *InlineResponse20097ProductsWireless) HasAvailableVersions() bool`

HasAvailableVersions returns a boolean if a field has been set.

### GetParticipateInNextBetaRelease

`func (o *InlineResponse20097ProductsWireless) GetParticipateInNextBetaRelease() bool`

GetParticipateInNextBetaRelease returns the ParticipateInNextBetaRelease field if non-nil, zero value otherwise.

### GetParticipateInNextBetaReleaseOk

`func (o *InlineResponse20097ProductsWireless) GetParticipateInNextBetaReleaseOk() (*bool, bool)`

GetParticipateInNextBetaReleaseOk returns a tuple with the ParticipateInNextBetaRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipateInNextBetaRelease

`func (o *InlineResponse20097ProductsWireless) SetParticipateInNextBetaRelease(v bool)`

SetParticipateInNextBetaRelease sets ParticipateInNextBetaRelease field to given value.

### HasParticipateInNextBetaRelease

`func (o *InlineResponse20097ProductsWireless) HasParticipateInNextBetaRelease() bool`

HasParticipateInNextBetaRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


