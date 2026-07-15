# InlineResponse200241

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The name of the new profile. Must be unique. | [optional] 
**NetworkId** | Pointer to **string** | The network ID of the RF Profile | [optional] 
**Name** | Pointer to **string** | The name of the new profile. Must be unique. This param is required on creation. | [optional] 
**ClientBalancingEnabled** | Pointer to **bool** | Steers client to best available access point. Can be either true or false. Defaults to true. | [optional] 
**MinBitrateType** | Pointer to **string** | Minimum bitrate can be set to either &#39;band&#39; or &#39;ssid&#39;. Defaults to band. | [optional] 
**BandSelectionType** | Pointer to **string** | Band selection can be set to either &#39;ssid&#39; or &#39;ap&#39;. This param is required on creation. | [optional] 
**ApBandSettings** | Pointer to [**InlineResponse200241ApBandSettings**](InlineResponse200241ApBandSettings.md) |  | [optional] 
**TwoFourGhzSettings** | Pointer to [**InlineResponse200241TwoFourGhzSettings**](InlineResponse200241TwoFourGhzSettings.md) |  | [optional] 
**FiveGhzSettings** | Pointer to [**InlineResponse200241FiveGhzSettings**](InlineResponse200241FiveGhzSettings.md) |  | [optional] 
**SixGhzSettings** | Pointer to [**InlineResponse200241SixGhzSettings**](InlineResponse200241SixGhzSettings.md) |  | [optional] 
**Transmission** | Pointer to [**InlineResponse200241Transmission**](InlineResponse200241Transmission.md) |  | [optional] 
**PerSsidSettings** | Pointer to [**InlineResponse200241PerSsidSettings**](InlineResponse200241PerSsidSettings.md) |  | [optional] 
**IsIndoorDefault** | Pointer to **bool** | Set this profile as the default indoor rf profile. If the profile ID is one of &#39;indoor&#39; or &#39;outdoor&#39;, then a new profile will be created from the respective ID and set as the default | [optional] 
**IsOutdoorDefault** | Pointer to **bool** | Set this profile as the default outdoor rf profile. If the profile ID is one of &#39;indoor&#39; or &#39;outdoor&#39;, then a new profile will be created from the respective ID and set as the default | [optional] 

## Methods

### NewInlineResponse200241

`func NewInlineResponse200241() *InlineResponse200241`

NewInlineResponse200241 instantiates a new InlineResponse200241 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200241WithDefaults

`func NewInlineResponse200241WithDefaults() *InlineResponse200241`

NewInlineResponse200241WithDefaults instantiates a new InlineResponse200241 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200241) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200241) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200241) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200241) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200241) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200241) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200241) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200241) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200241) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200241) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200241) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200241) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClientBalancingEnabled

`func (o *InlineResponse200241) GetClientBalancingEnabled() bool`

GetClientBalancingEnabled returns the ClientBalancingEnabled field if non-nil, zero value otherwise.

### GetClientBalancingEnabledOk

`func (o *InlineResponse200241) GetClientBalancingEnabledOk() (*bool, bool)`

GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBalancingEnabled

`func (o *InlineResponse200241) SetClientBalancingEnabled(v bool)`

SetClientBalancingEnabled sets ClientBalancingEnabled field to given value.

### HasClientBalancingEnabled

`func (o *InlineResponse200241) HasClientBalancingEnabled() bool`

HasClientBalancingEnabled returns a boolean if a field has been set.

### GetMinBitrateType

`func (o *InlineResponse200241) GetMinBitrateType() string`

GetMinBitrateType returns the MinBitrateType field if non-nil, zero value otherwise.

### GetMinBitrateTypeOk

`func (o *InlineResponse200241) GetMinBitrateTypeOk() (*string, bool)`

GetMinBitrateTypeOk returns a tuple with the MinBitrateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrateType

`func (o *InlineResponse200241) SetMinBitrateType(v string)`

SetMinBitrateType sets MinBitrateType field to given value.

### HasMinBitrateType

`func (o *InlineResponse200241) HasMinBitrateType() bool`

HasMinBitrateType returns a boolean if a field has been set.

### GetBandSelectionType

`func (o *InlineResponse200241) GetBandSelectionType() string`

GetBandSelectionType returns the BandSelectionType field if non-nil, zero value otherwise.

### GetBandSelectionTypeOk

`func (o *InlineResponse200241) GetBandSelectionTypeOk() (*string, bool)`

GetBandSelectionTypeOk returns a tuple with the BandSelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSelectionType

`func (o *InlineResponse200241) SetBandSelectionType(v string)`

SetBandSelectionType sets BandSelectionType field to given value.

### HasBandSelectionType

`func (o *InlineResponse200241) HasBandSelectionType() bool`

HasBandSelectionType returns a boolean if a field has been set.

### GetApBandSettings

`func (o *InlineResponse200241) GetApBandSettings() InlineResponse200241ApBandSettings`

GetApBandSettings returns the ApBandSettings field if non-nil, zero value otherwise.

### GetApBandSettingsOk

`func (o *InlineResponse200241) GetApBandSettingsOk() (*InlineResponse200241ApBandSettings, bool)`

GetApBandSettingsOk returns a tuple with the ApBandSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApBandSettings

`func (o *InlineResponse200241) SetApBandSettings(v InlineResponse200241ApBandSettings)`

SetApBandSettings sets ApBandSettings field to given value.

### HasApBandSettings

`func (o *InlineResponse200241) HasApBandSettings() bool`

HasApBandSettings returns a boolean if a field has been set.

### GetTwoFourGhzSettings

`func (o *InlineResponse200241) GetTwoFourGhzSettings() InlineResponse200241TwoFourGhzSettings`

GetTwoFourGhzSettings returns the TwoFourGhzSettings field if non-nil, zero value otherwise.

### GetTwoFourGhzSettingsOk

`func (o *InlineResponse200241) GetTwoFourGhzSettingsOk() (*InlineResponse200241TwoFourGhzSettings, bool)`

GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFourGhzSettings

`func (o *InlineResponse200241) SetTwoFourGhzSettings(v InlineResponse200241TwoFourGhzSettings)`

SetTwoFourGhzSettings sets TwoFourGhzSettings field to given value.

### HasTwoFourGhzSettings

`func (o *InlineResponse200241) HasTwoFourGhzSettings() bool`

HasTwoFourGhzSettings returns a boolean if a field has been set.

### GetFiveGhzSettings

`func (o *InlineResponse200241) GetFiveGhzSettings() InlineResponse200241FiveGhzSettings`

GetFiveGhzSettings returns the FiveGhzSettings field if non-nil, zero value otherwise.

### GetFiveGhzSettingsOk

`func (o *InlineResponse200241) GetFiveGhzSettingsOk() (*InlineResponse200241FiveGhzSettings, bool)`

GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveGhzSettings

`func (o *InlineResponse200241) SetFiveGhzSettings(v InlineResponse200241FiveGhzSettings)`

SetFiveGhzSettings sets FiveGhzSettings field to given value.

### HasFiveGhzSettings

`func (o *InlineResponse200241) HasFiveGhzSettings() bool`

HasFiveGhzSettings returns a boolean if a field has been set.

### GetSixGhzSettings

`func (o *InlineResponse200241) GetSixGhzSettings() InlineResponse200241SixGhzSettings`

GetSixGhzSettings returns the SixGhzSettings field if non-nil, zero value otherwise.

### GetSixGhzSettingsOk

`func (o *InlineResponse200241) GetSixGhzSettingsOk() (*InlineResponse200241SixGhzSettings, bool)`

GetSixGhzSettingsOk returns a tuple with the SixGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSixGhzSettings

`func (o *InlineResponse200241) SetSixGhzSettings(v InlineResponse200241SixGhzSettings)`

SetSixGhzSettings sets SixGhzSettings field to given value.

### HasSixGhzSettings

`func (o *InlineResponse200241) HasSixGhzSettings() bool`

HasSixGhzSettings returns a boolean if a field has been set.

### GetTransmission

`func (o *InlineResponse200241) GetTransmission() InlineResponse200241Transmission`

GetTransmission returns the Transmission field if non-nil, zero value otherwise.

### GetTransmissionOk

`func (o *InlineResponse200241) GetTransmissionOk() (*InlineResponse200241Transmission, bool)`

GetTransmissionOk returns a tuple with the Transmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmission

`func (o *InlineResponse200241) SetTransmission(v InlineResponse200241Transmission)`

SetTransmission sets Transmission field to given value.

### HasTransmission

`func (o *InlineResponse200241) HasTransmission() bool`

HasTransmission returns a boolean if a field has been set.

### GetPerSsidSettings

`func (o *InlineResponse200241) GetPerSsidSettings() InlineResponse200241PerSsidSettings`

GetPerSsidSettings returns the PerSsidSettings field if non-nil, zero value otherwise.

### GetPerSsidSettingsOk

`func (o *InlineResponse200241) GetPerSsidSettingsOk() (*InlineResponse200241PerSsidSettings, bool)`

GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidSettings

`func (o *InlineResponse200241) SetPerSsidSettings(v InlineResponse200241PerSsidSettings)`

SetPerSsidSettings sets PerSsidSettings field to given value.

### HasPerSsidSettings

`func (o *InlineResponse200241) HasPerSsidSettings() bool`

HasPerSsidSettings returns a boolean if a field has been set.

### GetIsIndoorDefault

`func (o *InlineResponse200241) GetIsIndoorDefault() bool`

GetIsIndoorDefault returns the IsIndoorDefault field if non-nil, zero value otherwise.

### GetIsIndoorDefaultOk

`func (o *InlineResponse200241) GetIsIndoorDefaultOk() (*bool, bool)`

GetIsIndoorDefaultOk returns a tuple with the IsIndoorDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIndoorDefault

`func (o *InlineResponse200241) SetIsIndoorDefault(v bool)`

SetIsIndoorDefault sets IsIndoorDefault field to given value.

### HasIsIndoorDefault

`func (o *InlineResponse200241) HasIsIndoorDefault() bool`

HasIsIndoorDefault returns a boolean if a field has been set.

### GetIsOutdoorDefault

`func (o *InlineResponse200241) GetIsOutdoorDefault() bool`

GetIsOutdoorDefault returns the IsOutdoorDefault field if non-nil, zero value otherwise.

### GetIsOutdoorDefaultOk

`func (o *InlineResponse200241) GetIsOutdoorDefaultOk() (*bool, bool)`

GetIsOutdoorDefaultOk returns a tuple with the IsOutdoorDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOutdoorDefault

`func (o *InlineResponse200241) SetIsOutdoorDefault(v bool)`

SetIsOutdoorDefault sets IsOutdoorDefault field to given value.

### HasIsOutdoorDefault

`func (o *InlineResponse200241) HasIsOutdoorDefault() bool`

HasIsOutdoorDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


