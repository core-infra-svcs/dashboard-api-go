# InlineResponse20084

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the quality retention profile | [optional] 
**NetworkId** | Pointer to **string** | Network ID | [optional] 
**Name** | Pointer to **string** | Name of the quality retention profile | [optional] 
**RestrictedBandwidthModeEnabled** | Pointer to **bool** | Whether restricted bandwidth mode is enabled | [optional] 
**MotionBasedRetentionEnabled** | Pointer to **bool** | Whether motion-based retention is enabled | [optional] 
**AudioRecordingEnabled** | Pointer to **bool** | Whether audio recording is enabled | [optional] 
**CloudArchiveEnabled** | Pointer to **bool** | Whether cloud archive is enabled | [optional] 
**MaxRetentionDays** | Pointer to **int32** | Maximum retention days | [optional] 
**ScheduleId** | Pointer to **NullableString** | Schedule ID | [optional] 
**MotionDetectorVersion** | Pointer to **int32** | Motion detector version | [optional] 
**SmartRetention** | Pointer to [**NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention**](NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention.md) |  | [optional] 
**VideoSettings** | Pointer to [**map[string]NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings**](NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings.md) | Video quality and resolution settings for camera models | [optional] 

## Methods

### NewInlineResponse20084

`func NewInlineResponse20084() *InlineResponse20084`

NewInlineResponse20084 instantiates a new InlineResponse20084 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20084WithDefaults

`func NewInlineResponse20084WithDefaults() *InlineResponse20084`

NewInlineResponse20084WithDefaults instantiates a new InlineResponse20084 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20084) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20084) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20084) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20084) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse20084) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse20084) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse20084) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse20084) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20084) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20084) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20084) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20084) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRestrictedBandwidthModeEnabled

`func (o *InlineResponse20084) GetRestrictedBandwidthModeEnabled() bool`

GetRestrictedBandwidthModeEnabled returns the RestrictedBandwidthModeEnabled field if non-nil, zero value otherwise.

### GetRestrictedBandwidthModeEnabledOk

`func (o *InlineResponse20084) GetRestrictedBandwidthModeEnabledOk() (*bool, bool)`

GetRestrictedBandwidthModeEnabledOk returns a tuple with the RestrictedBandwidthModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedBandwidthModeEnabled

`func (o *InlineResponse20084) SetRestrictedBandwidthModeEnabled(v bool)`

SetRestrictedBandwidthModeEnabled sets RestrictedBandwidthModeEnabled field to given value.

### HasRestrictedBandwidthModeEnabled

`func (o *InlineResponse20084) HasRestrictedBandwidthModeEnabled() bool`

HasRestrictedBandwidthModeEnabled returns a boolean if a field has been set.

### GetMotionBasedRetentionEnabled

`func (o *InlineResponse20084) GetMotionBasedRetentionEnabled() bool`

GetMotionBasedRetentionEnabled returns the MotionBasedRetentionEnabled field if non-nil, zero value otherwise.

### GetMotionBasedRetentionEnabledOk

`func (o *InlineResponse20084) GetMotionBasedRetentionEnabledOk() (*bool, bool)`

GetMotionBasedRetentionEnabledOk returns a tuple with the MotionBasedRetentionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionBasedRetentionEnabled

`func (o *InlineResponse20084) SetMotionBasedRetentionEnabled(v bool)`

SetMotionBasedRetentionEnabled sets MotionBasedRetentionEnabled field to given value.

### HasMotionBasedRetentionEnabled

`func (o *InlineResponse20084) HasMotionBasedRetentionEnabled() bool`

HasMotionBasedRetentionEnabled returns a boolean if a field has been set.

### GetAudioRecordingEnabled

`func (o *InlineResponse20084) GetAudioRecordingEnabled() bool`

GetAudioRecordingEnabled returns the AudioRecordingEnabled field if non-nil, zero value otherwise.

### GetAudioRecordingEnabledOk

`func (o *InlineResponse20084) GetAudioRecordingEnabledOk() (*bool, bool)`

GetAudioRecordingEnabledOk returns a tuple with the AudioRecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioRecordingEnabled

`func (o *InlineResponse20084) SetAudioRecordingEnabled(v bool)`

SetAudioRecordingEnabled sets AudioRecordingEnabled field to given value.

### HasAudioRecordingEnabled

`func (o *InlineResponse20084) HasAudioRecordingEnabled() bool`

HasAudioRecordingEnabled returns a boolean if a field has been set.

### GetCloudArchiveEnabled

`func (o *InlineResponse20084) GetCloudArchiveEnabled() bool`

GetCloudArchiveEnabled returns the CloudArchiveEnabled field if non-nil, zero value otherwise.

### GetCloudArchiveEnabledOk

`func (o *InlineResponse20084) GetCloudArchiveEnabledOk() (*bool, bool)`

GetCloudArchiveEnabledOk returns a tuple with the CloudArchiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudArchiveEnabled

`func (o *InlineResponse20084) SetCloudArchiveEnabled(v bool)`

SetCloudArchiveEnabled sets CloudArchiveEnabled field to given value.

### HasCloudArchiveEnabled

`func (o *InlineResponse20084) HasCloudArchiveEnabled() bool`

HasCloudArchiveEnabled returns a boolean if a field has been set.

### GetMaxRetentionDays

`func (o *InlineResponse20084) GetMaxRetentionDays() int32`

GetMaxRetentionDays returns the MaxRetentionDays field if non-nil, zero value otherwise.

### GetMaxRetentionDaysOk

`func (o *InlineResponse20084) GetMaxRetentionDaysOk() (*int32, bool)`

GetMaxRetentionDaysOk returns a tuple with the MaxRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetentionDays

`func (o *InlineResponse20084) SetMaxRetentionDays(v int32)`

SetMaxRetentionDays sets MaxRetentionDays field to given value.

### HasMaxRetentionDays

`func (o *InlineResponse20084) HasMaxRetentionDays() bool`

HasMaxRetentionDays returns a boolean if a field has been set.

### GetScheduleId

`func (o *InlineResponse20084) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *InlineResponse20084) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *InlineResponse20084) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *InlineResponse20084) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *InlineResponse20084) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *InlineResponse20084) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetMotionDetectorVersion

`func (o *InlineResponse20084) GetMotionDetectorVersion() int32`

GetMotionDetectorVersion returns the MotionDetectorVersion field if non-nil, zero value otherwise.

### GetMotionDetectorVersionOk

`func (o *InlineResponse20084) GetMotionDetectorVersionOk() (*int32, bool)`

GetMotionDetectorVersionOk returns a tuple with the MotionDetectorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionDetectorVersion

`func (o *InlineResponse20084) SetMotionDetectorVersion(v int32)`

SetMotionDetectorVersion sets MotionDetectorVersion field to given value.

### HasMotionDetectorVersion

`func (o *InlineResponse20084) HasMotionDetectorVersion() bool`

HasMotionDetectorVersion returns a boolean if a field has been set.

### GetSmartRetention

`func (o *InlineResponse20084) GetSmartRetention() NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention`

GetSmartRetention returns the SmartRetention field if non-nil, zero value otherwise.

### GetSmartRetentionOk

`func (o *InlineResponse20084) GetSmartRetentionOk() (*NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention, bool)`

GetSmartRetentionOk returns a tuple with the SmartRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartRetention

`func (o *InlineResponse20084) SetSmartRetention(v NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention)`

SetSmartRetention sets SmartRetention field to given value.

### HasSmartRetention

`func (o *InlineResponse20084) HasSmartRetention() bool`

HasSmartRetention returns a boolean if a field has been set.

### GetVideoSettings

`func (o *InlineResponse20084) GetVideoSettings() map[string]NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings`

GetVideoSettings returns the VideoSettings field if non-nil, zero value otherwise.

### GetVideoSettingsOk

`func (o *InlineResponse20084) GetVideoSettingsOk() (*map[string]NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings, bool)`

GetVideoSettingsOk returns a tuple with the VideoSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoSettings

`func (o *InlineResponse20084) SetVideoSettings(v map[string]NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings)`

SetVideoSettings sets VideoSettings field to given value.

### HasVideoSettings

`func (o *InlineResponse20084) HasVideoSettings() bool`

HasVideoSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


