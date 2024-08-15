# InlineResponse200184

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The name of the new profile. Must be unique. | [optional] 
**NetworkId** | Pointer to **string** | The network ID of the RF Profile | [optional] 
**Name** | Pointer to **string** | The name of the new profile. Must be unique. This param is required on creation. | [optional] 
**ClientBalancingEnabled** | Pointer to **bool** | Steers client to best available access point. Can be either true or false. Defaults to true. | [optional] 
**MinBitrateType** | Pointer to **string** | Minimum bitrate can be set to either &#39;band&#39; or &#39;ssid&#39;. Defaults to band. | [optional] 
**BandSelectionType** | Pointer to **string** | Band selection can be set to either &#39;ssid&#39; or &#39;ap&#39;. This param is required on creation. | [optional] 
**ApBandSettings** | Pointer to [**InlineResponse200184ApBandSettings**](InlineResponse200184ApBandSettings.md) |  | [optional] 
**TwoFourGhzSettings** | Pointer to [**InlineResponse200184TwoFourGhzSettings**](InlineResponse200184TwoFourGhzSettings.md) |  | [optional] 
**FiveGhzSettings** | Pointer to [**InlineResponse200184FiveGhzSettings**](InlineResponse200184FiveGhzSettings.md) |  | [optional] 
**SixGhzSettings** | Pointer to [**InlineResponse200184SixGhzSettings**](InlineResponse200184SixGhzSettings.md) |  | [optional] 
**Transmission** | Pointer to [**InlineResponse200184Transmission**](InlineResponse200184Transmission.md) |  | [optional] 
**PerSsidSettings** | Pointer to [**InlineResponse200184PerSsidSettings**](InlineResponse200184PerSsidSettings.md) |  | [optional] 
**IsIndoorDefault** | Pointer to **bool** | Set this profile as the default indoor rf profile. If the profile ID is one of &#39;indoor&#39; or &#39;outdoor&#39;,   then a new profile will be created from the respective ID and set as the default | [optional] 
**IsOutdoorDefault** | Pointer to **bool** | Set this profile as the default outdoor rf profile. If the profile ID is one of &#39;indoor&#39; or &#39;outdoor&#39;,   then a new profile will be created from the respective ID and set as the default | [optional] 

## Methods

### NewInlineResponse200184

`func NewInlineResponse200184() *InlineResponse200184`

NewInlineResponse200184 instantiates a new InlineResponse200184 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200184WithDefaults

`func NewInlineResponse200184WithDefaults() *InlineResponse200184`

NewInlineResponse200184WithDefaults instantiates a new InlineResponse200184 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200184) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200184) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200184) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200184) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200184) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200184) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200184) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200184) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200184) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200184) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200184) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200184) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClientBalancingEnabled

`func (o *InlineResponse200184) GetClientBalancingEnabled() bool`

GetClientBalancingEnabled returns the ClientBalancingEnabled field if non-nil, zero value otherwise.

### GetClientBalancingEnabledOk

`func (o *InlineResponse200184) GetClientBalancingEnabledOk() (*bool, bool)`

GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBalancingEnabled

`func (o *InlineResponse200184) SetClientBalancingEnabled(v bool)`

SetClientBalancingEnabled sets ClientBalancingEnabled field to given value.

### HasClientBalancingEnabled

`func (o *InlineResponse200184) HasClientBalancingEnabled() bool`

HasClientBalancingEnabled returns a boolean if a field has been set.

### GetMinBitrateType

`func (o *InlineResponse200184) GetMinBitrateType() string`

GetMinBitrateType returns the MinBitrateType field if non-nil, zero value otherwise.

### GetMinBitrateTypeOk

`func (o *InlineResponse200184) GetMinBitrateTypeOk() (*string, bool)`

GetMinBitrateTypeOk returns a tuple with the MinBitrateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrateType

`func (o *InlineResponse200184) SetMinBitrateType(v string)`

SetMinBitrateType sets MinBitrateType field to given value.

### HasMinBitrateType

`func (o *InlineResponse200184) HasMinBitrateType() bool`

HasMinBitrateType returns a boolean if a field has been set.

### GetBandSelectionType

`func (o *InlineResponse200184) GetBandSelectionType() string`

GetBandSelectionType returns the BandSelectionType field if non-nil, zero value otherwise.

### GetBandSelectionTypeOk

`func (o *InlineResponse200184) GetBandSelectionTypeOk() (*string, bool)`

GetBandSelectionTypeOk returns a tuple with the BandSelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSelectionType

`func (o *InlineResponse200184) SetBandSelectionType(v string)`

SetBandSelectionType sets BandSelectionType field to given value.

### HasBandSelectionType

`func (o *InlineResponse200184) HasBandSelectionType() bool`

HasBandSelectionType returns a boolean if a field has been set.

### GetApBandSettings

`func (o *InlineResponse200184) GetApBandSettings() InlineResponse200184ApBandSettings`

GetApBandSettings returns the ApBandSettings field if non-nil, zero value otherwise.

### GetApBandSettingsOk

`func (o *InlineResponse200184) GetApBandSettingsOk() (*InlineResponse200184ApBandSettings, bool)`

GetApBandSettingsOk returns a tuple with the ApBandSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApBandSettings

`func (o *InlineResponse200184) SetApBandSettings(v InlineResponse200184ApBandSettings)`

SetApBandSettings sets ApBandSettings field to given value.

### HasApBandSettings

`func (o *InlineResponse200184) HasApBandSettings() bool`

HasApBandSettings returns a boolean if a field has been set.

### GetTwoFourGhzSettings

`func (o *InlineResponse200184) GetTwoFourGhzSettings() InlineResponse200184TwoFourGhzSettings`

GetTwoFourGhzSettings returns the TwoFourGhzSettings field if non-nil, zero value otherwise.

### GetTwoFourGhzSettingsOk

`func (o *InlineResponse200184) GetTwoFourGhzSettingsOk() (*InlineResponse200184TwoFourGhzSettings, bool)`

GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFourGhzSettings

`func (o *InlineResponse200184) SetTwoFourGhzSettings(v InlineResponse200184TwoFourGhzSettings)`

SetTwoFourGhzSettings sets TwoFourGhzSettings field to given value.

### HasTwoFourGhzSettings

`func (o *InlineResponse200184) HasTwoFourGhzSettings() bool`

HasTwoFourGhzSettings returns a boolean if a field has been set.

### GetFiveGhzSettings

`func (o *InlineResponse200184) GetFiveGhzSettings() InlineResponse200184FiveGhzSettings`

GetFiveGhzSettings returns the FiveGhzSettings field if non-nil, zero value otherwise.

### GetFiveGhzSettingsOk

`func (o *InlineResponse200184) GetFiveGhzSettingsOk() (*InlineResponse200184FiveGhzSettings, bool)`

GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveGhzSettings

`func (o *InlineResponse200184) SetFiveGhzSettings(v InlineResponse200184FiveGhzSettings)`

SetFiveGhzSettings sets FiveGhzSettings field to given value.

### HasFiveGhzSettings

`func (o *InlineResponse200184) HasFiveGhzSettings() bool`

HasFiveGhzSettings returns a boolean if a field has been set.

### GetSixGhzSettings

`func (o *InlineResponse200184) GetSixGhzSettings() InlineResponse200184SixGhzSettings`

GetSixGhzSettings returns the SixGhzSettings field if non-nil, zero value otherwise.

### GetSixGhzSettingsOk

`func (o *InlineResponse200184) GetSixGhzSettingsOk() (*InlineResponse200184SixGhzSettings, bool)`

GetSixGhzSettingsOk returns a tuple with the SixGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSixGhzSettings

`func (o *InlineResponse200184) SetSixGhzSettings(v InlineResponse200184SixGhzSettings)`

SetSixGhzSettings sets SixGhzSettings field to given value.

### HasSixGhzSettings

`func (o *InlineResponse200184) HasSixGhzSettings() bool`

HasSixGhzSettings returns a boolean if a field has been set.

### GetTransmission

`func (o *InlineResponse200184) GetTransmission() InlineResponse200184Transmission`

GetTransmission returns the Transmission field if non-nil, zero value otherwise.

### GetTransmissionOk

`func (o *InlineResponse200184) GetTransmissionOk() (*InlineResponse200184Transmission, bool)`

GetTransmissionOk returns a tuple with the Transmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmission

`func (o *InlineResponse200184) SetTransmission(v InlineResponse200184Transmission)`

SetTransmission sets Transmission field to given value.

### HasTransmission

`func (o *InlineResponse200184) HasTransmission() bool`

HasTransmission returns a boolean if a field has been set.

### GetPerSsidSettings

`func (o *InlineResponse200184) GetPerSsidSettings() InlineResponse200184PerSsidSettings`

GetPerSsidSettings returns the PerSsidSettings field if non-nil, zero value otherwise.

### GetPerSsidSettingsOk

`func (o *InlineResponse200184) GetPerSsidSettingsOk() (*InlineResponse200184PerSsidSettings, bool)`

GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidSettings

`func (o *InlineResponse200184) SetPerSsidSettings(v InlineResponse200184PerSsidSettings)`

SetPerSsidSettings sets PerSsidSettings field to given value.

### HasPerSsidSettings

`func (o *InlineResponse200184) HasPerSsidSettings() bool`

HasPerSsidSettings returns a boolean if a field has been set.

### GetIsIndoorDefault

`func (o *InlineResponse200184) GetIsIndoorDefault() bool`

GetIsIndoorDefault returns the IsIndoorDefault field if non-nil, zero value otherwise.

### GetIsIndoorDefaultOk

`func (o *InlineResponse200184) GetIsIndoorDefaultOk() (*bool, bool)`

GetIsIndoorDefaultOk returns a tuple with the IsIndoorDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIndoorDefault

`func (o *InlineResponse200184) SetIsIndoorDefault(v bool)`

SetIsIndoorDefault sets IsIndoorDefault field to given value.

### HasIsIndoorDefault

`func (o *InlineResponse200184) HasIsIndoorDefault() bool`

HasIsIndoorDefault returns a boolean if a field has been set.

### GetIsOutdoorDefault

`func (o *InlineResponse200184) GetIsOutdoorDefault() bool`

GetIsOutdoorDefault returns the IsOutdoorDefault field if non-nil, zero value otherwise.

### GetIsOutdoorDefaultOk

`func (o *InlineResponse200184) GetIsOutdoorDefaultOk() (*bool, bool)`

GetIsOutdoorDefaultOk returns a tuple with the IsOutdoorDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOutdoorDefault

`func (o *InlineResponse200184) SetIsOutdoorDefault(v bool)`

SetIsOutdoorDefault sets IsOutdoorDefault field to given value.

### HasIsOutdoorDefault

`func (o *InlineResponse200184) HasIsOutdoorDefault() bool`

HasIsOutdoorDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


