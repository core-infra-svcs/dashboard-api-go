# InlineObject22

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RfProfileId** | Pointer to **string** | The ID of an RF profile to assign to the device. If the value of this parameter is null, the appropriate basic RF profile (indoor or outdoor) will be assigned to the device. Assigning an RF profile will clear ALL manually configured overrides on the device (channel width, channel, power). | [optional] 
**TwoFourGhzSettings** | Pointer to [**DevicesSerialWirelessRadioSettingsTwoFourGhzSettings**](DevicesSerialWirelessRadioSettingsTwoFourGhzSettings.md) |  | [optional] 
**FiveGhzSettings** | Pointer to [**DevicesSerialWirelessRadioSettingsFiveGhzSettings**](DevicesSerialWirelessRadioSettingsFiveGhzSettings.md) |  | [optional] 

## Methods

### NewInlineObject22

`func NewInlineObject22() *InlineObject22`

NewInlineObject22 instantiates a new InlineObject22 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject22WithDefaults

`func NewInlineObject22WithDefaults() *InlineObject22`

NewInlineObject22WithDefaults instantiates a new InlineObject22 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRfProfileId

`func (o *InlineObject22) GetRfProfileId() string`

GetRfProfileId returns the RfProfileId field if non-nil, zero value otherwise.

### GetRfProfileIdOk

`func (o *InlineObject22) GetRfProfileIdOk() (*string, bool)`

GetRfProfileIdOk returns a tuple with the RfProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfProfileId

`func (o *InlineObject22) SetRfProfileId(v string)`

SetRfProfileId sets RfProfileId field to given value.

### HasRfProfileId

`func (o *InlineObject22) HasRfProfileId() bool`

HasRfProfileId returns a boolean if a field has been set.

### GetTwoFourGhzSettings

`func (o *InlineObject22) GetTwoFourGhzSettings() DevicesSerialWirelessRadioSettingsTwoFourGhzSettings`

GetTwoFourGhzSettings returns the TwoFourGhzSettings field if non-nil, zero value otherwise.

### GetTwoFourGhzSettingsOk

`func (o *InlineObject22) GetTwoFourGhzSettingsOk() (*DevicesSerialWirelessRadioSettingsTwoFourGhzSettings, bool)`

GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFourGhzSettings

`func (o *InlineObject22) SetTwoFourGhzSettings(v DevicesSerialWirelessRadioSettingsTwoFourGhzSettings)`

SetTwoFourGhzSettings sets TwoFourGhzSettings field to given value.

### HasTwoFourGhzSettings

`func (o *InlineObject22) HasTwoFourGhzSettings() bool`

HasTwoFourGhzSettings returns a boolean if a field has been set.

### GetFiveGhzSettings

`func (o *InlineObject22) GetFiveGhzSettings() DevicesSerialWirelessRadioSettingsFiveGhzSettings`

GetFiveGhzSettings returns the FiveGhzSettings field if non-nil, zero value otherwise.

### GetFiveGhzSettingsOk

`func (o *InlineObject22) GetFiveGhzSettingsOk() (*DevicesSerialWirelessRadioSettingsFiveGhzSettings, bool)`

GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveGhzSettings

`func (o *InlineObject22) SetFiveGhzSettings(v DevicesSerialWirelessRadioSettingsFiveGhzSettings)`

SetFiveGhzSettings sets FiveGhzSettings field to given value.

### HasFiveGhzSettings

`func (o *InlineObject22) HasFiveGhzSettings() bool`

HasFiveGhzSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


