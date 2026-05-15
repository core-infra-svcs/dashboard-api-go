# InlineResponse20019

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **NullableString** | Quality and retention profile ID; null if not using a profile | [optional] 
**MotionBasedRetentionEnabled** | Pointer to **bool** | Whether motion-based retention is enabled | [optional] 
**AudioRecordingEnabled** | Pointer to **bool** | Whether audio recording is enabled | [optional] 
**RestrictedBandwidthModeEnabled** | Pointer to **bool** | Whether restricted bandwidth mode is enabled | [optional] 
**Quality** | Pointer to **string** | Video quality | [optional] 
**Resolution** | Pointer to **string** | Video resolution | [optional] 
**MotionDetectorVersion** | Pointer to **int32** | Motion detector version | [optional] 

## Methods

### NewInlineResponse20019

`func NewInlineResponse20019() *InlineResponse20019`

NewInlineResponse20019 instantiates a new InlineResponse20019 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20019WithDefaults

`func NewInlineResponse20019WithDefaults() *InlineResponse20019`

NewInlineResponse20019WithDefaults instantiates a new InlineResponse20019 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *InlineResponse20019) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *InlineResponse20019) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *InlineResponse20019) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *InlineResponse20019) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileIdNil

`func (o *InlineResponse20019) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *InlineResponse20019) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetMotionBasedRetentionEnabled

`func (o *InlineResponse20019) GetMotionBasedRetentionEnabled() bool`

GetMotionBasedRetentionEnabled returns the MotionBasedRetentionEnabled field if non-nil, zero value otherwise.

### GetMotionBasedRetentionEnabledOk

`func (o *InlineResponse20019) GetMotionBasedRetentionEnabledOk() (*bool, bool)`

GetMotionBasedRetentionEnabledOk returns a tuple with the MotionBasedRetentionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionBasedRetentionEnabled

`func (o *InlineResponse20019) SetMotionBasedRetentionEnabled(v bool)`

SetMotionBasedRetentionEnabled sets MotionBasedRetentionEnabled field to given value.

### HasMotionBasedRetentionEnabled

`func (o *InlineResponse20019) HasMotionBasedRetentionEnabled() bool`

HasMotionBasedRetentionEnabled returns a boolean if a field has been set.

### GetAudioRecordingEnabled

`func (o *InlineResponse20019) GetAudioRecordingEnabled() bool`

GetAudioRecordingEnabled returns the AudioRecordingEnabled field if non-nil, zero value otherwise.

### GetAudioRecordingEnabledOk

`func (o *InlineResponse20019) GetAudioRecordingEnabledOk() (*bool, bool)`

GetAudioRecordingEnabledOk returns a tuple with the AudioRecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioRecordingEnabled

`func (o *InlineResponse20019) SetAudioRecordingEnabled(v bool)`

SetAudioRecordingEnabled sets AudioRecordingEnabled field to given value.

### HasAudioRecordingEnabled

`func (o *InlineResponse20019) HasAudioRecordingEnabled() bool`

HasAudioRecordingEnabled returns a boolean if a field has been set.

### GetRestrictedBandwidthModeEnabled

`func (o *InlineResponse20019) GetRestrictedBandwidthModeEnabled() bool`

GetRestrictedBandwidthModeEnabled returns the RestrictedBandwidthModeEnabled field if non-nil, zero value otherwise.

### GetRestrictedBandwidthModeEnabledOk

`func (o *InlineResponse20019) GetRestrictedBandwidthModeEnabledOk() (*bool, bool)`

GetRestrictedBandwidthModeEnabledOk returns a tuple with the RestrictedBandwidthModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedBandwidthModeEnabled

`func (o *InlineResponse20019) SetRestrictedBandwidthModeEnabled(v bool)`

SetRestrictedBandwidthModeEnabled sets RestrictedBandwidthModeEnabled field to given value.

### HasRestrictedBandwidthModeEnabled

`func (o *InlineResponse20019) HasRestrictedBandwidthModeEnabled() bool`

HasRestrictedBandwidthModeEnabled returns a boolean if a field has been set.

### GetQuality

`func (o *InlineResponse20019) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *InlineResponse20019) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *InlineResponse20019) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *InlineResponse20019) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetResolution

`func (o *InlineResponse20019) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *InlineResponse20019) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *InlineResponse20019) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *InlineResponse20019) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetMotionDetectorVersion

`func (o *InlineResponse20019) GetMotionDetectorVersion() int32`

GetMotionDetectorVersion returns the MotionDetectorVersion field if non-nil, zero value otherwise.

### GetMotionDetectorVersionOk

`func (o *InlineResponse20019) GetMotionDetectorVersionOk() (*int32, bool)`

GetMotionDetectorVersionOk returns a tuple with the MotionDetectorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionDetectorVersion

`func (o *InlineResponse20019) SetMotionDetectorVersion(v int32)`

SetMotionDetectorVersion sets MotionDetectorVersion field to given value.

### HasMotionDetectorVersion

`func (o *InlineResponse20019) HasMotionDetectorVersion() bool`

HasMotionDetectorVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


