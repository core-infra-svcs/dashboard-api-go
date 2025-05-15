# InlineObject190

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the new profile. Must be unique. | [optional] 
**IsIndoorDefault** | Pointer to **bool** | Set this profile as the default indoor rf profile. If the profile ID is one of &#39;indoor&#39; or &#39;outdoor&#39;,   then a new profile will be created from the respective ID and set as the default | [optional] 
**IsOutdoorDefault** | Pointer to **bool** | Set this profile as the default outdoor rf profile. If the profile ID is one of &#39;indoor&#39; or &#39;outdoor&#39;,   then a new profile will be created from the respective ID and set as the default | [optional] 
**ClientBalancingEnabled** | Pointer to **bool** | Steers client to best available access point. Can be either true or false. | [optional] 
**MinBitrateType** | Pointer to **string** | Minimum bitrate can be set to either &#39;band&#39; or &#39;ssid&#39;. | [optional] 
**BandSelectionType** | Pointer to **string** | Band selection can be set to either &#39;ssid&#39; or &#39;ap&#39;. | [optional] 
**ApBandSettings** | Pointer to [**NetworksNetworkIdWirelessRfProfilesRfProfileIdApBandSettings**](NetworksNetworkIdWirelessRfProfilesRfProfileIdApBandSettings.md) |  | [optional] 
**TwoFourGhzSettings** | Pointer to [**NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings**](NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings.md) |  | [optional] 
**FiveGhzSettings** | Pointer to [**NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings**](NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings.md) |  | [optional] 
**SixGhzSettings** | Pointer to [**NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings**](NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings.md) |  | [optional] 
**Transmission** | Pointer to [**InlineResponse200196Transmission**](InlineResponse200196Transmission.md) |  | [optional] 
**PerSsidSettings** | Pointer to [**NetworksNetworkIdWirelessRfProfilesPerSsidSettings**](NetworksNetworkIdWirelessRfProfilesPerSsidSettings.md) |  | [optional] 
**FlexRadios** | Pointer to [**NetworksNetworkIdWirelessRfProfilesFlexRadios**](NetworksNetworkIdWirelessRfProfilesFlexRadios.md) |  | [optional] 

## Methods

### NewInlineObject190

`func NewInlineObject190() *InlineObject190`

NewInlineObject190 instantiates a new InlineObject190 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject190WithDefaults

`func NewInlineObject190WithDefaults() *InlineObject190`

NewInlineObject190WithDefaults instantiates a new InlineObject190 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject190) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject190) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject190) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject190) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsIndoorDefault

`func (o *InlineObject190) GetIsIndoorDefault() bool`

GetIsIndoorDefault returns the IsIndoorDefault field if non-nil, zero value otherwise.

### GetIsIndoorDefaultOk

`func (o *InlineObject190) GetIsIndoorDefaultOk() (*bool, bool)`

GetIsIndoorDefaultOk returns a tuple with the IsIndoorDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIndoorDefault

`func (o *InlineObject190) SetIsIndoorDefault(v bool)`

SetIsIndoorDefault sets IsIndoorDefault field to given value.

### HasIsIndoorDefault

`func (o *InlineObject190) HasIsIndoorDefault() bool`

HasIsIndoorDefault returns a boolean if a field has been set.

### GetIsOutdoorDefault

`func (o *InlineObject190) GetIsOutdoorDefault() bool`

GetIsOutdoorDefault returns the IsOutdoorDefault field if non-nil, zero value otherwise.

### GetIsOutdoorDefaultOk

`func (o *InlineObject190) GetIsOutdoorDefaultOk() (*bool, bool)`

GetIsOutdoorDefaultOk returns a tuple with the IsOutdoorDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOutdoorDefault

`func (o *InlineObject190) SetIsOutdoorDefault(v bool)`

SetIsOutdoorDefault sets IsOutdoorDefault field to given value.

### HasIsOutdoorDefault

`func (o *InlineObject190) HasIsOutdoorDefault() bool`

HasIsOutdoorDefault returns a boolean if a field has been set.

### GetClientBalancingEnabled

`func (o *InlineObject190) GetClientBalancingEnabled() bool`

GetClientBalancingEnabled returns the ClientBalancingEnabled field if non-nil, zero value otherwise.

### GetClientBalancingEnabledOk

`func (o *InlineObject190) GetClientBalancingEnabledOk() (*bool, bool)`

GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBalancingEnabled

`func (o *InlineObject190) SetClientBalancingEnabled(v bool)`

SetClientBalancingEnabled sets ClientBalancingEnabled field to given value.

### HasClientBalancingEnabled

`func (o *InlineObject190) HasClientBalancingEnabled() bool`

HasClientBalancingEnabled returns a boolean if a field has been set.

### GetMinBitrateType

`func (o *InlineObject190) GetMinBitrateType() string`

GetMinBitrateType returns the MinBitrateType field if non-nil, zero value otherwise.

### GetMinBitrateTypeOk

`func (o *InlineObject190) GetMinBitrateTypeOk() (*string, bool)`

GetMinBitrateTypeOk returns a tuple with the MinBitrateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrateType

`func (o *InlineObject190) SetMinBitrateType(v string)`

SetMinBitrateType sets MinBitrateType field to given value.

### HasMinBitrateType

`func (o *InlineObject190) HasMinBitrateType() bool`

HasMinBitrateType returns a boolean if a field has been set.

### GetBandSelectionType

`func (o *InlineObject190) GetBandSelectionType() string`

GetBandSelectionType returns the BandSelectionType field if non-nil, zero value otherwise.

### GetBandSelectionTypeOk

`func (o *InlineObject190) GetBandSelectionTypeOk() (*string, bool)`

GetBandSelectionTypeOk returns a tuple with the BandSelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSelectionType

`func (o *InlineObject190) SetBandSelectionType(v string)`

SetBandSelectionType sets BandSelectionType field to given value.

### HasBandSelectionType

`func (o *InlineObject190) HasBandSelectionType() bool`

HasBandSelectionType returns a boolean if a field has been set.

### GetApBandSettings

`func (o *InlineObject190) GetApBandSettings() NetworksNetworkIdWirelessRfProfilesRfProfileIdApBandSettings`

GetApBandSettings returns the ApBandSettings field if non-nil, zero value otherwise.

### GetApBandSettingsOk

`func (o *InlineObject190) GetApBandSettingsOk() (*NetworksNetworkIdWirelessRfProfilesRfProfileIdApBandSettings, bool)`

GetApBandSettingsOk returns a tuple with the ApBandSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApBandSettings

`func (o *InlineObject190) SetApBandSettings(v NetworksNetworkIdWirelessRfProfilesRfProfileIdApBandSettings)`

SetApBandSettings sets ApBandSettings field to given value.

### HasApBandSettings

`func (o *InlineObject190) HasApBandSettings() bool`

HasApBandSettings returns a boolean if a field has been set.

### GetTwoFourGhzSettings

`func (o *InlineObject190) GetTwoFourGhzSettings() NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings`

GetTwoFourGhzSettings returns the TwoFourGhzSettings field if non-nil, zero value otherwise.

### GetTwoFourGhzSettingsOk

`func (o *InlineObject190) GetTwoFourGhzSettingsOk() (*NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings, bool)`

GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFourGhzSettings

`func (o *InlineObject190) SetTwoFourGhzSettings(v NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings)`

SetTwoFourGhzSettings sets TwoFourGhzSettings field to given value.

### HasTwoFourGhzSettings

`func (o *InlineObject190) HasTwoFourGhzSettings() bool`

HasTwoFourGhzSettings returns a boolean if a field has been set.

### GetFiveGhzSettings

`func (o *InlineObject190) GetFiveGhzSettings() NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings`

GetFiveGhzSettings returns the FiveGhzSettings field if non-nil, zero value otherwise.

### GetFiveGhzSettingsOk

`func (o *InlineObject190) GetFiveGhzSettingsOk() (*NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings, bool)`

GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveGhzSettings

`func (o *InlineObject190) SetFiveGhzSettings(v NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings)`

SetFiveGhzSettings sets FiveGhzSettings field to given value.

### HasFiveGhzSettings

`func (o *InlineObject190) HasFiveGhzSettings() bool`

HasFiveGhzSettings returns a boolean if a field has been set.

### GetSixGhzSettings

`func (o *InlineObject190) GetSixGhzSettings() NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings`

GetSixGhzSettings returns the SixGhzSettings field if non-nil, zero value otherwise.

### GetSixGhzSettingsOk

`func (o *InlineObject190) GetSixGhzSettingsOk() (*NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings, bool)`

GetSixGhzSettingsOk returns a tuple with the SixGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSixGhzSettings

`func (o *InlineObject190) SetSixGhzSettings(v NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings)`

SetSixGhzSettings sets SixGhzSettings field to given value.

### HasSixGhzSettings

`func (o *InlineObject190) HasSixGhzSettings() bool`

HasSixGhzSettings returns a boolean if a field has been set.

### GetTransmission

`func (o *InlineObject190) GetTransmission() InlineResponse200196Transmission`

GetTransmission returns the Transmission field if non-nil, zero value otherwise.

### GetTransmissionOk

`func (o *InlineObject190) GetTransmissionOk() (*InlineResponse200196Transmission, bool)`

GetTransmissionOk returns a tuple with the Transmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmission

`func (o *InlineObject190) SetTransmission(v InlineResponse200196Transmission)`

SetTransmission sets Transmission field to given value.

### HasTransmission

`func (o *InlineObject190) HasTransmission() bool`

HasTransmission returns a boolean if a field has been set.

### GetPerSsidSettings

`func (o *InlineObject190) GetPerSsidSettings() NetworksNetworkIdWirelessRfProfilesPerSsidSettings`

GetPerSsidSettings returns the PerSsidSettings field if non-nil, zero value otherwise.

### GetPerSsidSettingsOk

`func (o *InlineObject190) GetPerSsidSettingsOk() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings, bool)`

GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidSettings

`func (o *InlineObject190) SetPerSsidSettings(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings)`

SetPerSsidSettings sets PerSsidSettings field to given value.

### HasPerSsidSettings

`func (o *InlineObject190) HasPerSsidSettings() bool`

HasPerSsidSettings returns a boolean if a field has been set.

### GetFlexRadios

`func (o *InlineObject190) GetFlexRadios() NetworksNetworkIdWirelessRfProfilesFlexRadios`

GetFlexRadios returns the FlexRadios field if non-nil, zero value otherwise.

### GetFlexRadiosOk

`func (o *InlineObject190) GetFlexRadiosOk() (*NetworksNetworkIdWirelessRfProfilesFlexRadios, bool)`

GetFlexRadiosOk returns a tuple with the FlexRadios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexRadios

`func (o *InlineObject190) SetFlexRadios(v NetworksNetworkIdWirelessRfProfilesFlexRadios)`

SetFlexRadios sets FlexRadios field to given value.

### HasFlexRadios

`func (o *InlineObject190) HasFlexRadios() bool`

HasFlexRadios returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


