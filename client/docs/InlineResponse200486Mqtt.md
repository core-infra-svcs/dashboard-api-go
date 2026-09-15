# InlineResponse200486Mqtt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingsId** | Pointer to **string** | MQTT ID | [optional] 
**Enabled** | Pointer to **bool** | MQTT Enabled | [optional] 
**Topic** | Pointer to **string** | MQTT Topic | [optional] 
**MessageFields** | Pointer to **[]string** | Select fields to populate in MQTT messages. Valid types are: RSSI, AP MAC address, Client MAC address, Timestamp, Radio, Network ID, Beacon type, Raw payload, Client UUID, Client major value, Client minor value, Signal power, Band, Slot ID | [optional] 
**Publishing** | Pointer to [**InlineResponse200486MqttPublishing**](InlineResponse200486MqttPublishing.md) |  | [optional] 
**Broker** | Pointer to [**InlineResponse200486MqttBroker**](InlineResponse200486MqttBroker.md) |  | [optional] 

## Methods

### NewInlineResponse200486Mqtt

`func NewInlineResponse200486Mqtt() *InlineResponse200486Mqtt`

NewInlineResponse200486Mqtt instantiates a new InlineResponse200486Mqtt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200486MqttWithDefaults

`func NewInlineResponse200486MqttWithDefaults() *InlineResponse200486Mqtt`

NewInlineResponse200486MqttWithDefaults instantiates a new InlineResponse200486Mqtt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingsId

`func (o *InlineResponse200486Mqtt) GetSettingsId() string`

GetSettingsId returns the SettingsId field if non-nil, zero value otherwise.

### GetSettingsIdOk

`func (o *InlineResponse200486Mqtt) GetSettingsIdOk() (*string, bool)`

GetSettingsIdOk returns a tuple with the SettingsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsId

`func (o *InlineResponse200486Mqtt) SetSettingsId(v string)`

SetSettingsId sets SettingsId field to given value.

### HasSettingsId

`func (o *InlineResponse200486Mqtt) HasSettingsId() bool`

HasSettingsId returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200486Mqtt) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200486Mqtt) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200486Mqtt) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200486Mqtt) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTopic

`func (o *InlineResponse200486Mqtt) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *InlineResponse200486Mqtt) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *InlineResponse200486Mqtt) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *InlineResponse200486Mqtt) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetMessageFields

`func (o *InlineResponse200486Mqtt) GetMessageFields() []string`

GetMessageFields returns the MessageFields field if non-nil, zero value otherwise.

### GetMessageFieldsOk

`func (o *InlineResponse200486Mqtt) GetMessageFieldsOk() (*[]string, bool)`

GetMessageFieldsOk returns a tuple with the MessageFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageFields

`func (o *InlineResponse200486Mqtt) SetMessageFields(v []string)`

SetMessageFields sets MessageFields field to given value.

### HasMessageFields

`func (o *InlineResponse200486Mqtt) HasMessageFields() bool`

HasMessageFields returns a boolean if a field has been set.

### GetPublishing

`func (o *InlineResponse200486Mqtt) GetPublishing() InlineResponse200486MqttPublishing`

GetPublishing returns the Publishing field if non-nil, zero value otherwise.

### GetPublishingOk

`func (o *InlineResponse200486Mqtt) GetPublishingOk() (*InlineResponse200486MqttPublishing, bool)`

GetPublishingOk returns a tuple with the Publishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishing

`func (o *InlineResponse200486Mqtt) SetPublishing(v InlineResponse200486MqttPublishing)`

SetPublishing sets Publishing field to given value.

### HasPublishing

`func (o *InlineResponse200486Mqtt) HasPublishing() bool`

HasPublishing returns a boolean if a field has been set.

### GetBroker

`func (o *InlineResponse200486Mqtt) GetBroker() InlineResponse200486MqttBroker`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *InlineResponse200486Mqtt) GetBrokerOk() (*InlineResponse200486MqttBroker, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *InlineResponse200486Mqtt) SetBroker(v InlineResponse200486MqttBroker)`

SetBroker sets Broker field to given value.

### HasBroker

`func (o *InlineResponse200486Mqtt) HasBroker() bool`

HasBroker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


