# InlineResponse20018

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenseEnabled** | Pointer to **bool** | Whether MV Sense is enabled on the camera | [optional] 
**MqttBrokerId** | Pointer to **NullableString** | ID of the MQTT broker; null if MQTT is disabled | [optional] 
**MqttTopics** | Pointer to **[]string** | MQTT topics the camera publishes to | [optional] 
**AudioDetection** | Pointer to [**InlineResponse20018AudioDetection**](InlineResponse20018AudioDetection.md) |  | [optional] 
**DetectionModelId** | Pointer to **NullableString** | ID of the object detection model | [optional] 

## Methods

### NewInlineResponse20018

`func NewInlineResponse20018() *InlineResponse20018`

NewInlineResponse20018 instantiates a new InlineResponse20018 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20018WithDefaults

`func NewInlineResponse20018WithDefaults() *InlineResponse20018`

NewInlineResponse20018WithDefaults instantiates a new InlineResponse20018 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSenseEnabled

`func (o *InlineResponse20018) GetSenseEnabled() bool`

GetSenseEnabled returns the SenseEnabled field if non-nil, zero value otherwise.

### GetSenseEnabledOk

`func (o *InlineResponse20018) GetSenseEnabledOk() (*bool, bool)`

GetSenseEnabledOk returns a tuple with the SenseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenseEnabled

`func (o *InlineResponse20018) SetSenseEnabled(v bool)`

SetSenseEnabled sets SenseEnabled field to given value.

### HasSenseEnabled

`func (o *InlineResponse20018) HasSenseEnabled() bool`

HasSenseEnabled returns a boolean if a field has been set.

### GetMqttBrokerId

`func (o *InlineResponse20018) GetMqttBrokerId() string`

GetMqttBrokerId returns the MqttBrokerId field if non-nil, zero value otherwise.

### GetMqttBrokerIdOk

`func (o *InlineResponse20018) GetMqttBrokerIdOk() (*string, bool)`

GetMqttBrokerIdOk returns a tuple with the MqttBrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttBrokerId

`func (o *InlineResponse20018) SetMqttBrokerId(v string)`

SetMqttBrokerId sets MqttBrokerId field to given value.

### HasMqttBrokerId

`func (o *InlineResponse20018) HasMqttBrokerId() bool`

HasMqttBrokerId returns a boolean if a field has been set.

### SetMqttBrokerIdNil

`func (o *InlineResponse20018) SetMqttBrokerIdNil(b bool)`

 SetMqttBrokerIdNil sets the value for MqttBrokerId to be an explicit nil

### UnsetMqttBrokerId
`func (o *InlineResponse20018) UnsetMqttBrokerId()`

UnsetMqttBrokerId ensures that no value is present for MqttBrokerId, not even an explicit nil
### GetMqttTopics

`func (o *InlineResponse20018) GetMqttTopics() []string`

GetMqttTopics returns the MqttTopics field if non-nil, zero value otherwise.

### GetMqttTopicsOk

`func (o *InlineResponse20018) GetMqttTopicsOk() (*[]string, bool)`

GetMqttTopicsOk returns a tuple with the MqttTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttTopics

`func (o *InlineResponse20018) SetMqttTopics(v []string)`

SetMqttTopics sets MqttTopics field to given value.

### HasMqttTopics

`func (o *InlineResponse20018) HasMqttTopics() bool`

HasMqttTopics returns a boolean if a field has been set.

### GetAudioDetection

`func (o *InlineResponse20018) GetAudioDetection() InlineResponse20018AudioDetection`

GetAudioDetection returns the AudioDetection field if non-nil, zero value otherwise.

### GetAudioDetectionOk

`func (o *InlineResponse20018) GetAudioDetectionOk() (*InlineResponse20018AudioDetection, bool)`

GetAudioDetectionOk returns a tuple with the AudioDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioDetection

`func (o *InlineResponse20018) SetAudioDetection(v InlineResponse20018AudioDetection)`

SetAudioDetection sets AudioDetection field to given value.

### HasAudioDetection

`func (o *InlineResponse20018) HasAudioDetection() bool`

HasAudioDetection returns a boolean if a field has been set.

### GetDetectionModelId

`func (o *InlineResponse20018) GetDetectionModelId() string`

GetDetectionModelId returns the DetectionModelId field if non-nil, zero value otherwise.

### GetDetectionModelIdOk

`func (o *InlineResponse20018) GetDetectionModelIdOk() (*string, bool)`

GetDetectionModelIdOk returns a tuple with the DetectionModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionModelId

`func (o *InlineResponse20018) SetDetectionModelId(v string)`

SetDetectionModelId sets DetectionModelId field to given value.

### HasDetectionModelId

`func (o *InlineResponse20018) HasDetectionModelId() bool`

HasDetectionModelId returns a boolean if a field has been set.

### SetDetectionModelIdNil

`func (o *InlineResponse20018) SetDetectionModelIdNil(b bool)`

 SetDetectionModelIdNil sets the value for DetectionModelId to be an explicit nil

### UnsetDetectionModelId
`func (o *InlineResponse20018) UnsetDetectionModelId()`

UnsetDetectionModelId ensures that no value is present for DetectionModelId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


