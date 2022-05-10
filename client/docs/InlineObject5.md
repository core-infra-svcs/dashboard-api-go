# InlineObject5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenseEnabled** | Pointer to **bool** | Boolean indicating if sense(license) is enabled(true) or disabled(false) on the camera | [optional] 
**MqttBrokerId** | Pointer to **string** | The ID of the MQTT broker to be enabled on the camera. A value of null will disable MQTT on the camera | [optional] 
**AudioDetection** | Pointer to [**DevicesSerialCameraSenseAudioDetection**](DevicesSerialCameraSenseAudioDetection.md) |  | [optional] 
**DetectionModelId** | Pointer to **string** | The ID of the object detection model | [optional] 

## Methods

### NewInlineObject5

`func NewInlineObject5() *InlineObject5`

NewInlineObject5 instantiates a new InlineObject5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject5WithDefaults

`func NewInlineObject5WithDefaults() *InlineObject5`

NewInlineObject5WithDefaults instantiates a new InlineObject5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSenseEnabled

`func (o *InlineObject5) GetSenseEnabled() bool`

GetSenseEnabled returns the SenseEnabled field if non-nil, zero value otherwise.

### GetSenseEnabledOk

`func (o *InlineObject5) GetSenseEnabledOk() (*bool, bool)`

GetSenseEnabledOk returns a tuple with the SenseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenseEnabled

`func (o *InlineObject5) SetSenseEnabled(v bool)`

SetSenseEnabled sets SenseEnabled field to given value.

### HasSenseEnabled

`func (o *InlineObject5) HasSenseEnabled() bool`

HasSenseEnabled returns a boolean if a field has been set.

### GetMqttBrokerId

`func (o *InlineObject5) GetMqttBrokerId() string`

GetMqttBrokerId returns the MqttBrokerId field if non-nil, zero value otherwise.

### GetMqttBrokerIdOk

`func (o *InlineObject5) GetMqttBrokerIdOk() (*string, bool)`

GetMqttBrokerIdOk returns a tuple with the MqttBrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttBrokerId

`func (o *InlineObject5) SetMqttBrokerId(v string)`

SetMqttBrokerId sets MqttBrokerId field to given value.

### HasMqttBrokerId

`func (o *InlineObject5) HasMqttBrokerId() bool`

HasMqttBrokerId returns a boolean if a field has been set.

### GetAudioDetection

`func (o *InlineObject5) GetAudioDetection() DevicesSerialCameraSenseAudioDetection`

GetAudioDetection returns the AudioDetection field if non-nil, zero value otherwise.

### GetAudioDetectionOk

`func (o *InlineObject5) GetAudioDetectionOk() (*DevicesSerialCameraSenseAudioDetection, bool)`

GetAudioDetectionOk returns a tuple with the AudioDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioDetection

`func (o *InlineObject5) SetAudioDetection(v DevicesSerialCameraSenseAudioDetection)`

SetAudioDetection sets AudioDetection field to given value.

### HasAudioDetection

`func (o *InlineObject5) HasAudioDetection() bool`

HasAudioDetection returns a boolean if a field has been set.

### GetDetectionModelId

`func (o *InlineObject5) GetDetectionModelId() string`

GetDetectionModelId returns the DetectionModelId field if non-nil, zero value otherwise.

### GetDetectionModelIdOk

`func (o *InlineObject5) GetDetectionModelIdOk() (*string, bool)`

GetDetectionModelIdOk returns a tuple with the DetectionModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionModelId

`func (o *InlineObject5) SetDetectionModelId(v string)`

SetDetectionModelId sets DetectionModelId field to given value.

### HasDetectionModelId

`func (o *InlineObject5) HasDetectionModelId() bool`

HasDetectionModelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


