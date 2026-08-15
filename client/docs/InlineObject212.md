# InlineObject212

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new profile. Must be unique. This param is required on creation. | 
**ClientBalancingEnabled** | Pointer to **bool** | Steers client to best available access point. Can be either true or false. Defaults to true. | [optional] 
**MinBitrateType** | Pointer to **string** | Minimum bitrate can be set to either &#39;band&#39; or &#39;ssid&#39;. Defaults to band. | [optional] 
**BandSelectionType** | **string** | Band selection can be set to either &#39;ssid&#39; or &#39;ap&#39;. This param is required on creation. | 
**ApBandSettings** | Pointer to [**NetworksNetworkIdWirelessRfProfilesApBandSettings**](NetworksNetworkIdWirelessRfProfilesApBandSettings.md) |  | [optional] 
**TwoFourGhzSettings** | Pointer to [**InlineResponse200242TwoFourGhzSettings**](InlineResponse200242TwoFourGhzSettings.md) |  | [optional] 
**FiveGhzSettings** | Pointer to [**InlineResponse200242FiveGhzSettings**](InlineResponse200242FiveGhzSettings.md) |  | [optional] 
**SixGhzSettings** | Pointer to [**InlineResponse200242SixGhzSettings**](InlineResponse200242SixGhzSettings.md) |  | [optional] 
**Transmission** | Pointer to [**InlineResponse200242Transmission**](InlineResponse200242Transmission.md) |  | [optional] 
**PerSsidSettings** | Pointer to [**NetworksNetworkIdWirelessRfProfilesPerSsidSettings**](NetworksNetworkIdWirelessRfProfilesPerSsidSettings.md) |  | [optional] 
**FlexRadios** | Pointer to [**NetworksNetworkIdWirelessRfProfilesFlexRadios**](NetworksNetworkIdWirelessRfProfilesFlexRadios.md) |  | [optional] 

## Methods

### NewInlineObject212

`func NewInlineObject212(name string, bandSelectionType string, ) *InlineObject212`

NewInlineObject212 instantiates a new InlineObject212 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject212WithDefaults

`func NewInlineObject212WithDefaults() *InlineObject212`

NewInlineObject212WithDefaults instantiates a new InlineObject212 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject212) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject212) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject212) SetName(v string)`

SetName sets Name field to given value.


### GetClientBalancingEnabled

`func (o *InlineObject212) GetClientBalancingEnabled() bool`

GetClientBalancingEnabled returns the ClientBalancingEnabled field if non-nil, zero value otherwise.

### GetClientBalancingEnabledOk

`func (o *InlineObject212) GetClientBalancingEnabledOk() (*bool, bool)`

GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBalancingEnabled

`func (o *InlineObject212) SetClientBalancingEnabled(v bool)`

SetClientBalancingEnabled sets ClientBalancingEnabled field to given value.

### HasClientBalancingEnabled

`func (o *InlineObject212) HasClientBalancingEnabled() bool`

HasClientBalancingEnabled returns a boolean if a field has been set.

### GetMinBitrateType

`func (o *InlineObject212) GetMinBitrateType() string`

GetMinBitrateType returns the MinBitrateType field if non-nil, zero value otherwise.

### GetMinBitrateTypeOk

`func (o *InlineObject212) GetMinBitrateTypeOk() (*string, bool)`

GetMinBitrateTypeOk returns a tuple with the MinBitrateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrateType

`func (o *InlineObject212) SetMinBitrateType(v string)`

SetMinBitrateType sets MinBitrateType field to given value.

### HasMinBitrateType

`func (o *InlineObject212) HasMinBitrateType() bool`

HasMinBitrateType returns a boolean if a field has been set.

### GetBandSelectionType

`func (o *InlineObject212) GetBandSelectionType() string`

GetBandSelectionType returns the BandSelectionType field if non-nil, zero value otherwise.

### GetBandSelectionTypeOk

`func (o *InlineObject212) GetBandSelectionTypeOk() (*string, bool)`

GetBandSelectionTypeOk returns a tuple with the BandSelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSelectionType

`func (o *InlineObject212) SetBandSelectionType(v string)`

SetBandSelectionType sets BandSelectionType field to given value.


### GetApBandSettings

`func (o *InlineObject212) GetApBandSettings() NetworksNetworkIdWirelessRfProfilesApBandSettings`

GetApBandSettings returns the ApBandSettings field if non-nil, zero value otherwise.

### GetApBandSettingsOk

`func (o *InlineObject212) GetApBandSettingsOk() (*NetworksNetworkIdWirelessRfProfilesApBandSettings, bool)`

GetApBandSettingsOk returns a tuple with the ApBandSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApBandSettings

`func (o *InlineObject212) SetApBandSettings(v NetworksNetworkIdWirelessRfProfilesApBandSettings)`

SetApBandSettings sets ApBandSettings field to given value.

### HasApBandSettings

`func (o *InlineObject212) HasApBandSettings() bool`

HasApBandSettings returns a boolean if a field has been set.

### GetTwoFourGhzSettings

`func (o *InlineObject212) GetTwoFourGhzSettings() InlineResponse200242TwoFourGhzSettings`

GetTwoFourGhzSettings returns the TwoFourGhzSettings field if non-nil, zero value otherwise.

### GetTwoFourGhzSettingsOk

`func (o *InlineObject212) GetTwoFourGhzSettingsOk() (*InlineResponse200242TwoFourGhzSettings, bool)`

GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFourGhzSettings

`func (o *InlineObject212) SetTwoFourGhzSettings(v InlineResponse200242TwoFourGhzSettings)`

SetTwoFourGhzSettings sets TwoFourGhzSettings field to given value.

### HasTwoFourGhzSettings

`func (o *InlineObject212) HasTwoFourGhzSettings() bool`

HasTwoFourGhzSettings returns a boolean if a field has been set.

### GetFiveGhzSettings

`func (o *InlineObject212) GetFiveGhzSettings() InlineResponse200242FiveGhzSettings`

GetFiveGhzSettings returns the FiveGhzSettings field if non-nil, zero value otherwise.

### GetFiveGhzSettingsOk

`func (o *InlineObject212) GetFiveGhzSettingsOk() (*InlineResponse200242FiveGhzSettings, bool)`

GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveGhzSettings

`func (o *InlineObject212) SetFiveGhzSettings(v InlineResponse200242FiveGhzSettings)`

SetFiveGhzSettings sets FiveGhzSettings field to given value.

### HasFiveGhzSettings

`func (o *InlineObject212) HasFiveGhzSettings() bool`

HasFiveGhzSettings returns a boolean if a field has been set.

### GetSixGhzSettings

`func (o *InlineObject212) GetSixGhzSettings() InlineResponse200242SixGhzSettings`

GetSixGhzSettings returns the SixGhzSettings field if non-nil, zero value otherwise.

### GetSixGhzSettingsOk

`func (o *InlineObject212) GetSixGhzSettingsOk() (*InlineResponse200242SixGhzSettings, bool)`

GetSixGhzSettingsOk returns a tuple with the SixGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSixGhzSettings

`func (o *InlineObject212) SetSixGhzSettings(v InlineResponse200242SixGhzSettings)`

SetSixGhzSettings sets SixGhzSettings field to given value.

### HasSixGhzSettings

`func (o *InlineObject212) HasSixGhzSettings() bool`

HasSixGhzSettings returns a boolean if a field has been set.

### GetTransmission

`func (o *InlineObject212) GetTransmission() InlineResponse200242Transmission`

GetTransmission returns the Transmission field if non-nil, zero value otherwise.

### GetTransmissionOk

`func (o *InlineObject212) GetTransmissionOk() (*InlineResponse200242Transmission, bool)`

GetTransmissionOk returns a tuple with the Transmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmission

`func (o *InlineObject212) SetTransmission(v InlineResponse200242Transmission)`

SetTransmission sets Transmission field to given value.

### HasTransmission

`func (o *InlineObject212) HasTransmission() bool`

HasTransmission returns a boolean if a field has been set.

### GetPerSsidSettings

`func (o *InlineObject212) GetPerSsidSettings() NetworksNetworkIdWirelessRfProfilesPerSsidSettings`

GetPerSsidSettings returns the PerSsidSettings field if non-nil, zero value otherwise.

### GetPerSsidSettingsOk

`func (o *InlineObject212) GetPerSsidSettingsOk() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings, bool)`

GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidSettings

`func (o *InlineObject212) SetPerSsidSettings(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings)`

SetPerSsidSettings sets PerSsidSettings field to given value.

### HasPerSsidSettings

`func (o *InlineObject212) HasPerSsidSettings() bool`

HasPerSsidSettings returns a boolean if a field has been set.

### GetFlexRadios

`func (o *InlineObject212) GetFlexRadios() NetworksNetworkIdWirelessRfProfilesFlexRadios`

GetFlexRadios returns the FlexRadios field if non-nil, zero value otherwise.

### GetFlexRadiosOk

`func (o *InlineObject212) GetFlexRadiosOk() (*NetworksNetworkIdWirelessRfProfilesFlexRadios, bool)`

GetFlexRadiosOk returns a tuple with the FlexRadios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexRadios

`func (o *InlineObject212) SetFlexRadios(v NetworksNetworkIdWirelessRfProfilesFlexRadios)`

SetFlexRadios sets FlexRadios field to given value.

### HasFlexRadios

`func (o *InlineObject212) HasFlexRadios() bool`

HasFlexRadios returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


