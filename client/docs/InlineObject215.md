# InlineObject215

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new profile. Must be unique. This param is required on creation. | 
**ClientBalancingEnabled** | Pointer to **bool** | Steers client to best available access point. Can be either true or false. Defaults to true. | [optional] 
**MinBitrateType** | Pointer to **string** | Minimum bitrate can be set to either &#39;band&#39; or &#39;ssid&#39;. Defaults to band. | [optional] 
**BandSelectionType** | **string** | Band selection can be set to either &#39;ssid&#39; or &#39;ap&#39;. This param is required on creation. | 
**ApBandSettings** | Pointer to [**NetworksNetworkIdWirelessRfProfilesApBandSettings**](NetworksNetworkIdWirelessRfProfilesApBandSettings.md) |  | [optional] 
**TwoFourGhzSettings** | Pointer to [**InlineResponse200244TwoFourGhzSettings**](InlineResponse200244TwoFourGhzSettings.md) |  | [optional] 
**FiveGhzSettings** | Pointer to [**NetworksNetworkIdWirelessRfProfilesFiveGhzSettings**](NetworksNetworkIdWirelessRfProfilesFiveGhzSettings.md) |  | [optional] 
**SixGhzSettings** | Pointer to [**InlineResponse200244SixGhzSettings**](InlineResponse200244SixGhzSettings.md) |  | [optional] 
**Transmission** | Pointer to [**InlineResponse200244Transmission**](InlineResponse200244Transmission.md) |  | [optional] 
**PerSsidSettings** | Pointer to [**NetworksNetworkIdWirelessRfProfilesPerSsidSettings**](NetworksNetworkIdWirelessRfProfilesPerSsidSettings.md) |  | [optional] 
**FlexRadios** | Pointer to [**NetworksNetworkIdWirelessRfProfilesFlexRadios**](NetworksNetworkIdWirelessRfProfilesFlexRadios.md) |  | [optional] 
**Dot11be** | Pointer to [**NetworksNetworkIdWirelessRfProfilesDot11be**](NetworksNetworkIdWirelessRfProfilesDot11be.md) |  | [optional] 

## Methods

### NewInlineObject215

`func NewInlineObject215(name string, bandSelectionType string, ) *InlineObject215`

NewInlineObject215 instantiates a new InlineObject215 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject215WithDefaults

`func NewInlineObject215WithDefaults() *InlineObject215`

NewInlineObject215WithDefaults instantiates a new InlineObject215 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject215) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject215) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject215) SetName(v string)`

SetName sets Name field to given value.


### GetClientBalancingEnabled

`func (o *InlineObject215) GetClientBalancingEnabled() bool`

GetClientBalancingEnabled returns the ClientBalancingEnabled field if non-nil, zero value otherwise.

### GetClientBalancingEnabledOk

`func (o *InlineObject215) GetClientBalancingEnabledOk() (*bool, bool)`

GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBalancingEnabled

`func (o *InlineObject215) SetClientBalancingEnabled(v bool)`

SetClientBalancingEnabled sets ClientBalancingEnabled field to given value.

### HasClientBalancingEnabled

`func (o *InlineObject215) HasClientBalancingEnabled() bool`

HasClientBalancingEnabled returns a boolean if a field has been set.

### GetMinBitrateType

`func (o *InlineObject215) GetMinBitrateType() string`

GetMinBitrateType returns the MinBitrateType field if non-nil, zero value otherwise.

### GetMinBitrateTypeOk

`func (o *InlineObject215) GetMinBitrateTypeOk() (*string, bool)`

GetMinBitrateTypeOk returns a tuple with the MinBitrateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrateType

`func (o *InlineObject215) SetMinBitrateType(v string)`

SetMinBitrateType sets MinBitrateType field to given value.

### HasMinBitrateType

`func (o *InlineObject215) HasMinBitrateType() bool`

HasMinBitrateType returns a boolean if a field has been set.

### GetBandSelectionType

`func (o *InlineObject215) GetBandSelectionType() string`

GetBandSelectionType returns the BandSelectionType field if non-nil, zero value otherwise.

### GetBandSelectionTypeOk

`func (o *InlineObject215) GetBandSelectionTypeOk() (*string, bool)`

GetBandSelectionTypeOk returns a tuple with the BandSelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSelectionType

`func (o *InlineObject215) SetBandSelectionType(v string)`

SetBandSelectionType sets BandSelectionType field to given value.


### GetApBandSettings

`func (o *InlineObject215) GetApBandSettings() NetworksNetworkIdWirelessRfProfilesApBandSettings`

GetApBandSettings returns the ApBandSettings field if non-nil, zero value otherwise.

### GetApBandSettingsOk

`func (o *InlineObject215) GetApBandSettingsOk() (*NetworksNetworkIdWirelessRfProfilesApBandSettings, bool)`

GetApBandSettingsOk returns a tuple with the ApBandSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApBandSettings

`func (o *InlineObject215) SetApBandSettings(v NetworksNetworkIdWirelessRfProfilesApBandSettings)`

SetApBandSettings sets ApBandSettings field to given value.

### HasApBandSettings

`func (o *InlineObject215) HasApBandSettings() bool`

HasApBandSettings returns a boolean if a field has been set.

### GetTwoFourGhzSettings

`func (o *InlineObject215) GetTwoFourGhzSettings() InlineResponse200244TwoFourGhzSettings`

GetTwoFourGhzSettings returns the TwoFourGhzSettings field if non-nil, zero value otherwise.

### GetTwoFourGhzSettingsOk

`func (o *InlineObject215) GetTwoFourGhzSettingsOk() (*InlineResponse200244TwoFourGhzSettings, bool)`

GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFourGhzSettings

`func (o *InlineObject215) SetTwoFourGhzSettings(v InlineResponse200244TwoFourGhzSettings)`

SetTwoFourGhzSettings sets TwoFourGhzSettings field to given value.

### HasTwoFourGhzSettings

`func (o *InlineObject215) HasTwoFourGhzSettings() bool`

HasTwoFourGhzSettings returns a boolean if a field has been set.

### GetFiveGhzSettings

`func (o *InlineObject215) GetFiveGhzSettings() NetworksNetworkIdWirelessRfProfilesFiveGhzSettings`

GetFiveGhzSettings returns the FiveGhzSettings field if non-nil, zero value otherwise.

### GetFiveGhzSettingsOk

`func (o *InlineObject215) GetFiveGhzSettingsOk() (*NetworksNetworkIdWirelessRfProfilesFiveGhzSettings, bool)`

GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveGhzSettings

`func (o *InlineObject215) SetFiveGhzSettings(v NetworksNetworkIdWirelessRfProfilesFiveGhzSettings)`

SetFiveGhzSettings sets FiveGhzSettings field to given value.

### HasFiveGhzSettings

`func (o *InlineObject215) HasFiveGhzSettings() bool`

HasFiveGhzSettings returns a boolean if a field has been set.

### GetSixGhzSettings

`func (o *InlineObject215) GetSixGhzSettings() InlineResponse200244SixGhzSettings`

GetSixGhzSettings returns the SixGhzSettings field if non-nil, zero value otherwise.

### GetSixGhzSettingsOk

`func (o *InlineObject215) GetSixGhzSettingsOk() (*InlineResponse200244SixGhzSettings, bool)`

GetSixGhzSettingsOk returns a tuple with the SixGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSixGhzSettings

`func (o *InlineObject215) SetSixGhzSettings(v InlineResponse200244SixGhzSettings)`

SetSixGhzSettings sets SixGhzSettings field to given value.

### HasSixGhzSettings

`func (o *InlineObject215) HasSixGhzSettings() bool`

HasSixGhzSettings returns a boolean if a field has been set.

### GetTransmission

`func (o *InlineObject215) GetTransmission() InlineResponse200244Transmission`

GetTransmission returns the Transmission field if non-nil, zero value otherwise.

### GetTransmissionOk

`func (o *InlineObject215) GetTransmissionOk() (*InlineResponse200244Transmission, bool)`

GetTransmissionOk returns a tuple with the Transmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmission

`func (o *InlineObject215) SetTransmission(v InlineResponse200244Transmission)`

SetTransmission sets Transmission field to given value.

### HasTransmission

`func (o *InlineObject215) HasTransmission() bool`

HasTransmission returns a boolean if a field has been set.

### GetPerSsidSettings

`func (o *InlineObject215) GetPerSsidSettings() NetworksNetworkIdWirelessRfProfilesPerSsidSettings`

GetPerSsidSettings returns the PerSsidSettings field if non-nil, zero value otherwise.

### GetPerSsidSettingsOk

`func (o *InlineObject215) GetPerSsidSettingsOk() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings, bool)`

GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidSettings

`func (o *InlineObject215) SetPerSsidSettings(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings)`

SetPerSsidSettings sets PerSsidSettings field to given value.

### HasPerSsidSettings

`func (o *InlineObject215) HasPerSsidSettings() bool`

HasPerSsidSettings returns a boolean if a field has been set.

### GetFlexRadios

`func (o *InlineObject215) GetFlexRadios() NetworksNetworkIdWirelessRfProfilesFlexRadios`

GetFlexRadios returns the FlexRadios field if non-nil, zero value otherwise.

### GetFlexRadiosOk

`func (o *InlineObject215) GetFlexRadiosOk() (*NetworksNetworkIdWirelessRfProfilesFlexRadios, bool)`

GetFlexRadiosOk returns a tuple with the FlexRadios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexRadios

`func (o *InlineObject215) SetFlexRadios(v NetworksNetworkIdWirelessRfProfilesFlexRadios)`

SetFlexRadios sets FlexRadios field to given value.

### HasFlexRadios

`func (o *InlineObject215) HasFlexRadios() bool`

HasFlexRadios returns a boolean if a field has been set.

### GetDot11be

`func (o *InlineObject215) GetDot11be() NetworksNetworkIdWirelessRfProfilesDot11be`

GetDot11be returns the Dot11be field if non-nil, zero value otherwise.

### GetDot11beOk

`func (o *InlineObject215) GetDot11beOk() (*NetworksNetworkIdWirelessRfProfilesDot11be, bool)`

GetDot11beOk returns a tuple with the Dot11be field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot11be

`func (o *InlineObject215) SetDot11be(v NetworksNetworkIdWirelessRfProfilesDot11be)`

SetDot11be sets Dot11be field to given value.

### HasDot11be

`func (o *InlineObject215) HasDot11be() bool`

HasDot11be returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


