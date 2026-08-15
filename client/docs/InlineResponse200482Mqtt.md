# InlineResponse200482Mqtt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingsId** | Pointer to **string** | MQTT ID | [optional] 
**Enabled** | Pointer to **bool** | MQTT Enabled | [optional] 
**Topic** | Pointer to **string** | MQTT Topic | [optional] 
**MessageFields** | Pointer to **[]string** | Select fields to populate in MQTT messages. Valid types are: RSSI, AP MAC address, Client MAC address, Timestamp, Radio, Network ID, Beacon type, Raw payload, Client UUID, Client major value, Client minor value, Signal power, Band, Slot ID | [optional] 
**Publishing** | Pointer to [**InlineResponse200482MqttPublishing**](InlineResponse200482MqttPublishing.md) |  | [optional] 
**Broker** | Pointer to [**InlineResponse200482MqttBroker**](InlineResponse200482MqttBroker.md) |  | [optional] 

## Methods

### NewInlineResponse200482Mqtt

`func NewInlineResponse200482Mqtt() *InlineResponse200482Mqtt`

NewInlineResponse200482Mqtt instantiates a new InlineResponse200482Mqtt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200482MqttWithDefaults

`func NewInlineResponse200482MqttWithDefaults() *InlineResponse200482Mqtt`

NewInlineResponse200482MqttWithDefaults instantiates a new InlineResponse200482Mqtt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingsId

`func (o *InlineResponse200482Mqtt) GetSettingsId() string`

GetSettingsId returns the SettingsId field if non-nil, zero value otherwise.

### GetSettingsIdOk

`func (o *InlineResponse200482Mqtt) GetSettingsIdOk() (*string, bool)`

GetSettingsIdOk returns a tuple with the SettingsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsId

`func (o *InlineResponse200482Mqtt) SetSettingsId(v string)`

SetSettingsId sets SettingsId field to given value.

### HasSettingsId

`func (o *InlineResponse200482Mqtt) HasSettingsId() bool`

HasSettingsId returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200482Mqtt) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200482Mqtt) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200482Mqtt) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200482Mqtt) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTopic

`func (o *InlineResponse200482Mqtt) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *InlineResponse200482Mqtt) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *InlineResponse200482Mqtt) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *InlineResponse200482Mqtt) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetMessageFields

`func (o *InlineResponse200482Mqtt) GetMessageFields() []string`

GetMessageFields returns the MessageFields field if non-nil, zero value otherwise.

### GetMessageFieldsOk

`func (o *InlineResponse200482Mqtt) GetMessageFieldsOk() (*[]string, bool)`

GetMessageFieldsOk returns a tuple with the MessageFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageFields

`func (o *InlineResponse200482Mqtt) SetMessageFields(v []string)`

SetMessageFields sets MessageFields field to given value.

### HasMessageFields

`func (o *InlineResponse200482Mqtt) HasMessageFields() bool`

HasMessageFields returns a boolean if a field has been set.

### GetPublishing

`func (o *InlineResponse200482Mqtt) GetPublishing() InlineResponse200482MqttPublishing`

GetPublishing returns the Publishing field if non-nil, zero value otherwise.

### GetPublishingOk

`func (o *InlineResponse200482Mqtt) GetPublishingOk() (*InlineResponse200482MqttPublishing, bool)`

GetPublishingOk returns a tuple with the Publishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishing

`func (o *InlineResponse200482Mqtt) SetPublishing(v InlineResponse200482MqttPublishing)`

SetPublishing sets Publishing field to given value.

### HasPublishing

`func (o *InlineResponse200482Mqtt) HasPublishing() bool`

HasPublishing returns a boolean if a field has been set.

### GetBroker

`func (o *InlineResponse200482Mqtt) GetBroker() InlineResponse200482MqttBroker`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *InlineResponse200482Mqtt) GetBrokerOk() (*InlineResponse200482MqttBroker, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *InlineResponse200482Mqtt) SetBroker(v InlineResponse200482MqttBroker)`

SetBroker sets Broker field to given value.

### HasBroker

`func (o *InlineResponse200482Mqtt) HasBroker() bool`

HasBroker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


