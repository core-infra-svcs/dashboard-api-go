# InlineResponse200373Mqtt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingsId** | Pointer to **string** | MQTT ID | [optional] 
**Enabled** | Pointer to **bool** | MQTT Enabled | [optional] 
**Topic** | Pointer to **string** | MQTT Topic | [optional] 
**MessageFields** | Pointer to **[]string** | Select fields to populate in MQTT messages. Valid types are: RSSI, AP MAC address, Client MAC address, Timestamp, Radio, Network ID, Beacon type, Raw payload, Client UUID, Client major value, Client minor value, Signal power, Band, Slot ID | [optional] 
**Publishing** | Pointer to [**InlineResponse200373MqttPublishing**](InlineResponse200373MqttPublishing.md) |  | [optional] 
**Broker** | Pointer to [**InlineResponse200373MqttBroker**](InlineResponse200373MqttBroker.md) |  | [optional] 

## Methods

### NewInlineResponse200373Mqtt

`func NewInlineResponse200373Mqtt() *InlineResponse200373Mqtt`

NewInlineResponse200373Mqtt instantiates a new InlineResponse200373Mqtt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200373MqttWithDefaults

`func NewInlineResponse200373MqttWithDefaults() *InlineResponse200373Mqtt`

NewInlineResponse200373MqttWithDefaults instantiates a new InlineResponse200373Mqtt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingsId

`func (o *InlineResponse200373Mqtt) GetSettingsId() string`

GetSettingsId returns the SettingsId field if non-nil, zero value otherwise.

### GetSettingsIdOk

`func (o *InlineResponse200373Mqtt) GetSettingsIdOk() (*string, bool)`

GetSettingsIdOk returns a tuple with the SettingsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsId

`func (o *InlineResponse200373Mqtt) SetSettingsId(v string)`

SetSettingsId sets SettingsId field to given value.

### HasSettingsId

`func (o *InlineResponse200373Mqtt) HasSettingsId() bool`

HasSettingsId returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200373Mqtt) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200373Mqtt) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200373Mqtt) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200373Mqtt) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTopic

`func (o *InlineResponse200373Mqtt) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *InlineResponse200373Mqtt) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *InlineResponse200373Mqtt) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *InlineResponse200373Mqtt) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetMessageFields

`func (o *InlineResponse200373Mqtt) GetMessageFields() []string`

GetMessageFields returns the MessageFields field if non-nil, zero value otherwise.

### GetMessageFieldsOk

`func (o *InlineResponse200373Mqtt) GetMessageFieldsOk() (*[]string, bool)`

GetMessageFieldsOk returns a tuple with the MessageFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageFields

`func (o *InlineResponse200373Mqtt) SetMessageFields(v []string)`

SetMessageFields sets MessageFields field to given value.

### HasMessageFields

`func (o *InlineResponse200373Mqtt) HasMessageFields() bool`

HasMessageFields returns a boolean if a field has been set.

### GetPublishing

`func (o *InlineResponse200373Mqtt) GetPublishing() InlineResponse200373MqttPublishing`

GetPublishing returns the Publishing field if non-nil, zero value otherwise.

### GetPublishingOk

`func (o *InlineResponse200373Mqtt) GetPublishingOk() (*InlineResponse200373MqttPublishing, bool)`

GetPublishingOk returns a tuple with the Publishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishing

`func (o *InlineResponse200373Mqtt) SetPublishing(v InlineResponse200373MqttPublishing)`

SetPublishing sets Publishing field to given value.

### HasPublishing

`func (o *InlineResponse200373Mqtt) HasPublishing() bool`

HasPublishing returns a boolean if a field has been set.

### GetBroker

`func (o *InlineResponse200373Mqtt) GetBroker() InlineResponse200373MqttBroker`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *InlineResponse200373Mqtt) GetBrokerOk() (*InlineResponse200373MqttBroker, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *InlineResponse200373Mqtt) SetBroker(v InlineResponse200373MqttBroker)`

SetBroker sets Broker field to given value.

### HasBroker

`func (o *InlineResponse200373Mqtt) HasBroker() bool`

HasBroker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


