# InlineObject83

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the new profile. Must be unique. | [optional] 
**MotionBasedRetentionEnabled** | Pointer to **bool** | Deletes footage older than 3 days in which no motion was detected. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras. | [optional] 
**RestrictedBandwidthModeEnabled** | Pointer to **bool** | Disable features that require additional bandwidth such as Motion Recap. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras. | [optional] 
**AudioRecordingEnabled** | Pointer to **bool** | Whether or not to record audio. Can be either true or false. Defaults to false. | [optional] 
**CloudArchiveEnabled** | Pointer to **bool** | Create redundant video backup using Cloud Archive. Can be either true or false. Defaults to false. | [optional] 
**MotionDetectorVersion** | Pointer to **int32** | The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2. | [optional] 
**SmartRetention** | Pointer to [**NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention**](NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention.md) |  | [optional] 
**ScheduleId** | Pointer to **NullableString** | Schedule for which this camera will record video, or &#39;null&#39; to always record. | [optional] 
**MaxRetentionDays** | Pointer to **NullableInt32** | The maximum number of days for which the data will be stored, or &#39;null&#39; to keep data until storage space runs out. If the former, it can be in the range of one to ninety days. | [optional] 
**VideoSettings** | Pointer to [**NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings**](NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings.md) |  | [optional] 

## Methods

### NewInlineObject83

`func NewInlineObject83() *InlineObject83`

NewInlineObject83 instantiates a new InlineObject83 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject83WithDefaults

`func NewInlineObject83WithDefaults() *InlineObject83`

NewInlineObject83WithDefaults instantiates a new InlineObject83 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject83) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject83) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject83) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject83) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMotionBasedRetentionEnabled

`func (o *InlineObject83) GetMotionBasedRetentionEnabled() bool`

GetMotionBasedRetentionEnabled returns the MotionBasedRetentionEnabled field if non-nil, zero value otherwise.

### GetMotionBasedRetentionEnabledOk

`func (o *InlineObject83) GetMotionBasedRetentionEnabledOk() (*bool, bool)`

GetMotionBasedRetentionEnabledOk returns a tuple with the MotionBasedRetentionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionBasedRetentionEnabled

`func (o *InlineObject83) SetMotionBasedRetentionEnabled(v bool)`

SetMotionBasedRetentionEnabled sets MotionBasedRetentionEnabled field to given value.

### HasMotionBasedRetentionEnabled

`func (o *InlineObject83) HasMotionBasedRetentionEnabled() bool`

HasMotionBasedRetentionEnabled returns a boolean if a field has been set.

### GetRestrictedBandwidthModeEnabled

`func (o *InlineObject83) GetRestrictedBandwidthModeEnabled() bool`

GetRestrictedBandwidthModeEnabled returns the RestrictedBandwidthModeEnabled field if non-nil, zero value otherwise.

### GetRestrictedBandwidthModeEnabledOk

`func (o *InlineObject83) GetRestrictedBandwidthModeEnabledOk() (*bool, bool)`

GetRestrictedBandwidthModeEnabledOk returns a tuple with the RestrictedBandwidthModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedBandwidthModeEnabled

`func (o *InlineObject83) SetRestrictedBandwidthModeEnabled(v bool)`

SetRestrictedBandwidthModeEnabled sets RestrictedBandwidthModeEnabled field to given value.

### HasRestrictedBandwidthModeEnabled

`func (o *InlineObject83) HasRestrictedBandwidthModeEnabled() bool`

HasRestrictedBandwidthModeEnabled returns a boolean if a field has been set.

### GetAudioRecordingEnabled

`func (o *InlineObject83) GetAudioRecordingEnabled() bool`

GetAudioRecordingEnabled returns the AudioRecordingEnabled field if non-nil, zero value otherwise.

### GetAudioRecordingEnabledOk

`func (o *InlineObject83) GetAudioRecordingEnabledOk() (*bool, bool)`

GetAudioRecordingEnabledOk returns a tuple with the AudioRecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioRecordingEnabled

`func (o *InlineObject83) SetAudioRecordingEnabled(v bool)`

SetAudioRecordingEnabled sets AudioRecordingEnabled field to given value.

### HasAudioRecordingEnabled

`func (o *InlineObject83) HasAudioRecordingEnabled() bool`

HasAudioRecordingEnabled returns a boolean if a field has been set.

### GetCloudArchiveEnabled

`func (o *InlineObject83) GetCloudArchiveEnabled() bool`

GetCloudArchiveEnabled returns the CloudArchiveEnabled field if non-nil, zero value otherwise.

### GetCloudArchiveEnabledOk

`func (o *InlineObject83) GetCloudArchiveEnabledOk() (*bool, bool)`

GetCloudArchiveEnabledOk returns a tuple with the CloudArchiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudArchiveEnabled

`func (o *InlineObject83) SetCloudArchiveEnabled(v bool)`

SetCloudArchiveEnabled sets CloudArchiveEnabled field to given value.

### HasCloudArchiveEnabled

`func (o *InlineObject83) HasCloudArchiveEnabled() bool`

HasCloudArchiveEnabled returns a boolean if a field has been set.

### GetMotionDetectorVersion

`func (o *InlineObject83) GetMotionDetectorVersion() int32`

GetMotionDetectorVersion returns the MotionDetectorVersion field if non-nil, zero value otherwise.

### GetMotionDetectorVersionOk

`func (o *InlineObject83) GetMotionDetectorVersionOk() (*int32, bool)`

GetMotionDetectorVersionOk returns a tuple with the MotionDetectorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionDetectorVersion

`func (o *InlineObject83) SetMotionDetectorVersion(v int32)`

SetMotionDetectorVersion sets MotionDetectorVersion field to given value.

### HasMotionDetectorVersion

`func (o *InlineObject83) HasMotionDetectorVersion() bool`

HasMotionDetectorVersion returns a boolean if a field has been set.

### GetSmartRetention

`func (o *InlineObject83) GetSmartRetention() NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention`

GetSmartRetention returns the SmartRetention field if non-nil, zero value otherwise.

### GetSmartRetentionOk

`func (o *InlineObject83) GetSmartRetentionOk() (*NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention, bool)`

GetSmartRetentionOk returns a tuple with the SmartRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartRetention

`func (o *InlineObject83) SetSmartRetention(v NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention)`

SetSmartRetention sets SmartRetention field to given value.

### HasSmartRetention

`func (o *InlineObject83) HasSmartRetention() bool`

HasSmartRetention returns a boolean if a field has been set.

### GetScheduleId

`func (o *InlineObject83) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *InlineObject83) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *InlineObject83) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *InlineObject83) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *InlineObject83) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *InlineObject83) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetMaxRetentionDays

`func (o *InlineObject83) GetMaxRetentionDays() int32`

GetMaxRetentionDays returns the MaxRetentionDays field if non-nil, zero value otherwise.

### GetMaxRetentionDaysOk

`func (o *InlineObject83) GetMaxRetentionDaysOk() (*int32, bool)`

GetMaxRetentionDaysOk returns a tuple with the MaxRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetentionDays

`func (o *InlineObject83) SetMaxRetentionDays(v int32)`

SetMaxRetentionDays sets MaxRetentionDays field to given value.

### HasMaxRetentionDays

`func (o *InlineObject83) HasMaxRetentionDays() bool`

HasMaxRetentionDays returns a boolean if a field has been set.

### SetMaxRetentionDaysNil

`func (o *InlineObject83) SetMaxRetentionDaysNil(b bool)`

 SetMaxRetentionDaysNil sets the value for MaxRetentionDays to be an explicit nil

### UnsetMaxRetentionDays
`func (o *InlineObject83) UnsetMaxRetentionDays()`

UnsetMaxRetentionDays ensures that no value is present for MaxRetentionDays, not even an explicit nil
### GetVideoSettings

`func (o *InlineObject83) GetVideoSettings() NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings`

GetVideoSettings returns the VideoSettings field if non-nil, zero value otherwise.

### GetVideoSettingsOk

`func (o *InlineObject83) GetVideoSettingsOk() (*NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings, bool)`

GetVideoSettingsOk returns a tuple with the VideoSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoSettings

`func (o *InlineObject83) SetVideoSettings(v NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings)`

SetVideoSettings sets VideoSettings field to given value.

### HasVideoSettings

`func (o *InlineObject83) HasVideoSettings() bool`

HasVideoSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


