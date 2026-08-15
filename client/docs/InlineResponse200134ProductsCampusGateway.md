# InlineResponse200134ProductsCampusGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to [**InlineResponse200134ProductsWirelessCurrentVersion**](InlineResponse200134ProductsWirelessCurrentVersion.md) |  | [optional] 
**LastUpgrade** | Pointer to [**InlineResponse200134ProductsWirelessLastUpgrade**](InlineResponse200134ProductsWirelessLastUpgrade.md) |  | [optional] 
**NextUpgrade** | Pointer to [**InlineResponse200134ProductsCampusGatewayNextUpgrade**](InlineResponse200134ProductsCampusGatewayNextUpgrade.md) |  | [optional] 
**IsUpgradeAvailable** | Pointer to **bool** | Whether or not an upgraded recommended firmware version is available | [optional] 
**AvailableVersions** | Pointer to [**[]InlineResponse200134ProductsWirelessAvailableVersions**](InlineResponse200134ProductsWirelessAvailableVersions.md) | Firmware versions available for upgrade | [optional] 
**ParticipateInNextBetaRelease** | Pointer to **bool** | Whether or not the network wants beta firmware | [optional] 

## Methods

### NewInlineResponse200134ProductsCampusGateway

`func NewInlineResponse200134ProductsCampusGateway() *InlineResponse200134ProductsCampusGateway`

NewInlineResponse200134ProductsCampusGateway instantiates a new InlineResponse200134ProductsCampusGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200134ProductsCampusGatewayWithDefaults

`func NewInlineResponse200134ProductsCampusGatewayWithDefaults() *InlineResponse200134ProductsCampusGateway`

NewInlineResponse200134ProductsCampusGatewayWithDefaults instantiates a new InlineResponse200134ProductsCampusGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *InlineResponse200134ProductsCampusGateway) GetCurrentVersion() InlineResponse200134ProductsWirelessCurrentVersion`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *InlineResponse200134ProductsCampusGateway) GetCurrentVersionOk() (*InlineResponse200134ProductsWirelessCurrentVersion, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *InlineResponse200134ProductsCampusGateway) SetCurrentVersion(v InlineResponse200134ProductsWirelessCurrentVersion)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *InlineResponse200134ProductsCampusGateway) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetLastUpgrade

`func (o *InlineResponse200134ProductsCampusGateway) GetLastUpgrade() InlineResponse200134ProductsWirelessLastUpgrade`

GetLastUpgrade returns the LastUpgrade field if non-nil, zero value otherwise.

### GetLastUpgradeOk

`func (o *InlineResponse200134ProductsCampusGateway) GetLastUpgradeOk() (*InlineResponse200134ProductsWirelessLastUpgrade, bool)`

GetLastUpgradeOk returns a tuple with the LastUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpgrade

`func (o *InlineResponse200134ProductsCampusGateway) SetLastUpgrade(v InlineResponse200134ProductsWirelessLastUpgrade)`

SetLastUpgrade sets LastUpgrade field to given value.

### HasLastUpgrade

`func (o *InlineResponse200134ProductsCampusGateway) HasLastUpgrade() bool`

HasLastUpgrade returns a boolean if a field has been set.

### GetNextUpgrade

`func (o *InlineResponse200134ProductsCampusGateway) GetNextUpgrade() InlineResponse200134ProductsCampusGatewayNextUpgrade`

GetNextUpgrade returns the NextUpgrade field if non-nil, zero value otherwise.

### GetNextUpgradeOk

`func (o *InlineResponse200134ProductsCampusGateway) GetNextUpgradeOk() (*InlineResponse200134ProductsCampusGatewayNextUpgrade, bool)`

GetNextUpgradeOk returns a tuple with the NextUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpgrade

`func (o *InlineResponse200134ProductsCampusGateway) SetNextUpgrade(v InlineResponse200134ProductsCampusGatewayNextUpgrade)`

SetNextUpgrade sets NextUpgrade field to given value.

### HasNextUpgrade

`func (o *InlineResponse200134ProductsCampusGateway) HasNextUpgrade() bool`

HasNextUpgrade returns a boolean if a field has been set.

### GetIsUpgradeAvailable

`func (o *InlineResponse200134ProductsCampusGateway) GetIsUpgradeAvailable() bool`

GetIsUpgradeAvailable returns the IsUpgradeAvailable field if non-nil, zero value otherwise.

### GetIsUpgradeAvailableOk

`func (o *InlineResponse200134ProductsCampusGateway) GetIsUpgradeAvailableOk() (*bool, bool)`

GetIsUpgradeAvailableOk returns a tuple with the IsUpgradeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgradeAvailable

`func (o *InlineResponse200134ProductsCampusGateway) SetIsUpgradeAvailable(v bool)`

SetIsUpgradeAvailable sets IsUpgradeAvailable field to given value.

### HasIsUpgradeAvailable

`func (o *InlineResponse200134ProductsCampusGateway) HasIsUpgradeAvailable() bool`

HasIsUpgradeAvailable returns a boolean if a field has been set.

### GetAvailableVersions

`func (o *InlineResponse200134ProductsCampusGateway) GetAvailableVersions() []InlineResponse200134ProductsWirelessAvailableVersions`

GetAvailableVersions returns the AvailableVersions field if non-nil, zero value otherwise.

### GetAvailableVersionsOk

`func (o *InlineResponse200134ProductsCampusGateway) GetAvailableVersionsOk() (*[]InlineResponse200134ProductsWirelessAvailableVersions, bool)`

GetAvailableVersionsOk returns a tuple with the AvailableVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVersions

`func (o *InlineResponse200134ProductsCampusGateway) SetAvailableVersions(v []InlineResponse200134ProductsWirelessAvailableVersions)`

SetAvailableVersions sets AvailableVersions field to given value.

### HasAvailableVersions

`func (o *InlineResponse200134ProductsCampusGateway) HasAvailableVersions() bool`

HasAvailableVersions returns a boolean if a field has been set.

### GetParticipateInNextBetaRelease

`func (o *InlineResponse200134ProductsCampusGateway) GetParticipateInNextBetaRelease() bool`

GetParticipateInNextBetaRelease returns the ParticipateInNextBetaRelease field if non-nil, zero value otherwise.

### GetParticipateInNextBetaReleaseOk

`func (o *InlineResponse200134ProductsCampusGateway) GetParticipateInNextBetaReleaseOk() (*bool, bool)`

GetParticipateInNextBetaReleaseOk returns a tuple with the ParticipateInNextBetaRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipateInNextBetaRelease

`func (o *InlineResponse200134ProductsCampusGateway) SetParticipateInNextBetaRelease(v bool)`

SetParticipateInNextBetaRelease sets ParticipateInNextBetaRelease field to given value.

### HasParticipateInNextBetaRelease

`func (o *InlineResponse200134ProductsCampusGateway) HasParticipateInNextBetaRelease() bool`

HasParticipateInNextBetaRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


