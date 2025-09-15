# InlineObject192

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new profile. Must be unique. This param is required on creation. | 
**ClientBalancingEnabled** | Pointer to **bool** | Steers client to best available access point. Can be either true or false. Defaults to true. | [optional] 
**MinBitrateType** | Pointer to **string** | Minimum bitrate can be set to either &#39;band&#39; or &#39;ssid&#39;. Defaults to band. | [optional] 
**BandSelectionType** | **string** | Band selection can be set to either &#39;ssid&#39; or &#39;ap&#39;. This param is required on creation. | 
**ApBandSettings** | Pointer to [**NetworksNetworkIdWirelessRfProfilesApBandSettings**](NetworksNetworkIdWirelessRfProfilesApBandSettings.md) |  | [optional] 
**TwoFourGhzSettings** | Pointer to [**InlineResponse200200TwoFourGhzSettings**](InlineResponse200200TwoFourGhzSettings.md) |  | [optional] 
**FiveGhzSettings** | Pointer to [**InlineResponse200200FiveGhzSettings**](InlineResponse200200FiveGhzSettings.md) |  | [optional] 
**SixGhzSettings** | Pointer to [**NetworksNetworkIdWirelessRfProfilesSixGhzSettings**](NetworksNetworkIdWirelessRfProfilesSixGhzSettings.md) |  | [optional] 
**Transmission** | Pointer to [**InlineResponse200200Transmission**](InlineResponse200200Transmission.md) |  | [optional] 
**PerSsidSettings** | Pointer to [**NetworksNetworkIdWirelessRfProfilesPerSsidSettings**](NetworksNetworkIdWirelessRfProfilesPerSsidSettings.md) |  | [optional] 
**FlexRadios** | Pointer to [**NetworksNetworkIdWirelessRfProfilesFlexRadios**](NetworksNetworkIdWirelessRfProfilesFlexRadios.md) |  | [optional] 

## Methods

### NewInlineObject192

`func NewInlineObject192(name string, bandSelectionType string, ) *InlineObject192`

NewInlineObject192 instantiates a new InlineObject192 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject192WithDefaults

`func NewInlineObject192WithDefaults() *InlineObject192`

NewInlineObject192WithDefaults instantiates a new InlineObject192 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject192) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject192) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject192) SetName(v string)`

SetName sets Name field to given value.


### GetClientBalancingEnabled

`func (o *InlineObject192) GetClientBalancingEnabled() bool`

GetClientBalancingEnabled returns the ClientBalancingEnabled field if non-nil, zero value otherwise.

### GetClientBalancingEnabledOk

`func (o *InlineObject192) GetClientBalancingEnabledOk() (*bool, bool)`

GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBalancingEnabled

`func (o *InlineObject192) SetClientBalancingEnabled(v bool)`

SetClientBalancingEnabled sets ClientBalancingEnabled field to given value.

### HasClientBalancingEnabled

`func (o *InlineObject192) HasClientBalancingEnabled() bool`

HasClientBalancingEnabled returns a boolean if a field has been set.

### GetMinBitrateType

`func (o *InlineObject192) GetMinBitrateType() string`

GetMinBitrateType returns the MinBitrateType field if non-nil, zero value otherwise.

### GetMinBitrateTypeOk

`func (o *InlineObject192) GetMinBitrateTypeOk() (*string, bool)`

GetMinBitrateTypeOk returns a tuple with the MinBitrateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrateType

`func (o *InlineObject192) SetMinBitrateType(v string)`

SetMinBitrateType sets MinBitrateType field to given value.

### HasMinBitrateType

`func (o *InlineObject192) HasMinBitrateType() bool`

HasMinBitrateType returns a boolean if a field has been set.

### GetBandSelectionType

`func (o *InlineObject192) GetBandSelectionType() string`

GetBandSelectionType returns the BandSelectionType field if non-nil, zero value otherwise.

### GetBandSelectionTypeOk

`func (o *InlineObject192) GetBandSelectionTypeOk() (*string, bool)`

GetBandSelectionTypeOk returns a tuple with the BandSelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSelectionType

`func (o *InlineObject192) SetBandSelectionType(v string)`

SetBandSelectionType sets BandSelectionType field to given value.


### GetApBandSettings

`func (o *InlineObject192) GetApBandSettings() NetworksNetworkIdWirelessRfProfilesApBandSettings`

GetApBandSettings returns the ApBandSettings field if non-nil, zero value otherwise.

### GetApBandSettingsOk

`func (o *InlineObject192) GetApBandSettingsOk() (*NetworksNetworkIdWirelessRfProfilesApBandSettings, bool)`

GetApBandSettingsOk returns a tuple with the ApBandSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApBandSettings

`func (o *InlineObject192) SetApBandSettings(v NetworksNetworkIdWirelessRfProfilesApBandSettings)`

SetApBandSettings sets ApBandSettings field to given value.

### HasApBandSettings

`func (o *InlineObject192) HasApBandSettings() bool`

HasApBandSettings returns a boolean if a field has been set.

### GetTwoFourGhzSettings

`func (o *InlineObject192) GetTwoFourGhzSettings() InlineResponse200200TwoFourGhzSettings`

GetTwoFourGhzSettings returns the TwoFourGhzSettings field if non-nil, zero value otherwise.

### GetTwoFourGhzSettingsOk

`func (o *InlineObject192) GetTwoFourGhzSettingsOk() (*InlineResponse200200TwoFourGhzSettings, bool)`

GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFourGhzSettings

`func (o *InlineObject192) SetTwoFourGhzSettings(v InlineResponse200200TwoFourGhzSettings)`

SetTwoFourGhzSettings sets TwoFourGhzSettings field to given value.

### HasTwoFourGhzSettings

`func (o *InlineObject192) HasTwoFourGhzSettings() bool`

HasTwoFourGhzSettings returns a boolean if a field has been set.

### GetFiveGhzSettings

`func (o *InlineObject192) GetFiveGhzSettings() InlineResponse200200FiveGhzSettings`

GetFiveGhzSettings returns the FiveGhzSettings field if non-nil, zero value otherwise.

### GetFiveGhzSettingsOk

`func (o *InlineObject192) GetFiveGhzSettingsOk() (*InlineResponse200200FiveGhzSettings, bool)`

GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveGhzSettings

`func (o *InlineObject192) SetFiveGhzSettings(v InlineResponse200200FiveGhzSettings)`

SetFiveGhzSettings sets FiveGhzSettings field to given value.

### HasFiveGhzSettings

`func (o *InlineObject192) HasFiveGhzSettings() bool`

HasFiveGhzSettings returns a boolean if a field has been set.

### GetSixGhzSettings

`func (o *InlineObject192) GetSixGhzSettings() NetworksNetworkIdWirelessRfProfilesSixGhzSettings`

GetSixGhzSettings returns the SixGhzSettings field if non-nil, zero value otherwise.

### GetSixGhzSettingsOk

`func (o *InlineObject192) GetSixGhzSettingsOk() (*NetworksNetworkIdWirelessRfProfilesSixGhzSettings, bool)`

GetSixGhzSettingsOk returns a tuple with the SixGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSixGhzSettings

`func (o *InlineObject192) SetSixGhzSettings(v NetworksNetworkIdWirelessRfProfilesSixGhzSettings)`

SetSixGhzSettings sets SixGhzSettings field to given value.

### HasSixGhzSettings

`func (o *InlineObject192) HasSixGhzSettings() bool`

HasSixGhzSettings returns a boolean if a field has been set.

### GetTransmission

`func (o *InlineObject192) GetTransmission() InlineResponse200200Transmission`

GetTransmission returns the Transmission field if non-nil, zero value otherwise.

### GetTransmissionOk

`func (o *InlineObject192) GetTransmissionOk() (*InlineResponse200200Transmission, bool)`

GetTransmissionOk returns a tuple with the Transmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmission

`func (o *InlineObject192) SetTransmission(v InlineResponse200200Transmission)`

SetTransmission sets Transmission field to given value.

### HasTransmission

`func (o *InlineObject192) HasTransmission() bool`

HasTransmission returns a boolean if a field has been set.

### GetPerSsidSettings

`func (o *InlineObject192) GetPerSsidSettings() NetworksNetworkIdWirelessRfProfilesPerSsidSettings`

GetPerSsidSettings returns the PerSsidSettings field if non-nil, zero value otherwise.

### GetPerSsidSettingsOk

`func (o *InlineObject192) GetPerSsidSettingsOk() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings, bool)`

GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidSettings

`func (o *InlineObject192) SetPerSsidSettings(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings)`

SetPerSsidSettings sets PerSsidSettings field to given value.

### HasPerSsidSettings

`func (o *InlineObject192) HasPerSsidSettings() bool`

HasPerSsidSettings returns a boolean if a field has been set.

### GetFlexRadios

`func (o *InlineObject192) GetFlexRadios() NetworksNetworkIdWirelessRfProfilesFlexRadios`

GetFlexRadios returns the FlexRadios field if non-nil, zero value otherwise.

### GetFlexRadiosOk

`func (o *InlineObject192) GetFlexRadiosOk() (*NetworksNetworkIdWirelessRfProfilesFlexRadios, bool)`

GetFlexRadiosOk returns a tuple with the FlexRadios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexRadios

`func (o *InlineObject192) SetFlexRadios(v NetworksNetworkIdWirelessRfProfilesFlexRadios)`

SetFlexRadios sets FlexRadios field to given value.

### HasFlexRadios

`func (o *InlineObject192) HasFlexRadios() bool`

HasFlexRadios returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


