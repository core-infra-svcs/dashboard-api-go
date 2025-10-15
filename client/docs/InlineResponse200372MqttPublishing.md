# InlineResponse200372MqttPublishing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | Pointer to **int32** | MQTT Publishing Frequency in seconds. Will be between 1 and 2147483647. Default is 1 second | [optional] 
**Qos** | Pointer to **int32** | MQTT Publishing QoS. Valid types are: 0, 1, 2 | [optional] 

## Methods

### NewInlineResponse200372MqttPublishing

`func NewInlineResponse200372MqttPublishing() *InlineResponse200372MqttPublishing`

NewInlineResponse200372MqttPublishing instantiates a new InlineResponse200372MqttPublishing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200372MqttPublishingWithDefaults

`func NewInlineResponse200372MqttPublishingWithDefaults() *InlineResponse200372MqttPublishing`

NewInlineResponse200372MqttPublishingWithDefaults instantiates a new InlineResponse200372MqttPublishing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *InlineResponse200372MqttPublishing) GetFrequency() int32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *InlineResponse200372MqttPublishing) GetFrequencyOk() (*int32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *InlineResponse200372MqttPublishing) SetFrequency(v int32)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *InlineResponse200372MqttPublishing) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetQos

`func (o *InlineResponse200372MqttPublishing) GetQos() int32`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *InlineResponse200372MqttPublishing) GetQosOk() (*int32, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *InlineResponse200372MqttPublishing) SetQos(v int32)`

SetQos sets Qos field to given value.

### HasQos

`func (o *InlineResponse200372MqttPublishing) HasQos() bool`

HasQos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


