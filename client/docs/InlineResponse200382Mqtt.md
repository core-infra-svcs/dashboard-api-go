# InlineResponse200382Mqtt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingsId** | Pointer to **string** | MQTT ID | [optional] 
**Enabled** | Pointer to **bool** | MQTT Enabled | [optional] 
**Topic** | Pointer to **string** | MQTT Topic | [optional] 
**MessageFields** | Pointer to **[]string** | Select fields to populate in MQTT messages. Valid types are: RSSI, AP MAC address, Client MAC address, Timestamp, Radio, Network ID, Beacon type, Raw payload, Client UUID, Client major value, Client minor value, Signal power, Band, Slot ID | [optional] 
**Publishing** | Pointer to [**InlineResponse200382MqttPublishing**](InlineResponse200382MqttPublishing.md) |  | [optional] 
**Broker** | Pointer to [**InlineResponse200382MqttBroker**](InlineResponse200382MqttBroker.md) |  | [optional] 

## Methods

### NewInlineResponse200382Mqtt

`func NewInlineResponse200382Mqtt() *InlineResponse200382Mqtt`

NewInlineResponse200382Mqtt instantiates a new InlineResponse200382Mqtt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200382MqttWithDefaults

`func NewInlineResponse200382MqttWithDefaults() *InlineResponse200382Mqtt`

NewInlineResponse200382MqttWithDefaults instantiates a new InlineResponse200382Mqtt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingsId

`func (o *InlineResponse200382Mqtt) GetSettingsId() string`

GetSettingsId returns the SettingsId field if non-nil, zero value otherwise.

### GetSettingsIdOk

`func (o *InlineResponse200382Mqtt) GetSettingsIdOk() (*string, bool)`

GetSettingsIdOk returns a tuple with the SettingsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsId

`func (o *InlineResponse200382Mqtt) SetSettingsId(v string)`

SetSettingsId sets SettingsId field to given value.

### HasSettingsId

`func (o *InlineResponse200382Mqtt) HasSettingsId() bool`

HasSettingsId returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200382Mqtt) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200382Mqtt) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200382Mqtt) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200382Mqtt) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTopic

`func (o *InlineResponse200382Mqtt) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *InlineResponse200382Mqtt) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *InlineResponse200382Mqtt) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *InlineResponse200382Mqtt) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetMessageFields

`func (o *InlineResponse200382Mqtt) GetMessageFields() []string`

GetMessageFields returns the MessageFields field if non-nil, zero value otherwise.

### GetMessageFieldsOk

`func (o *InlineResponse200382Mqtt) GetMessageFieldsOk() (*[]string, bool)`

GetMessageFieldsOk returns a tuple with the MessageFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageFields

`func (o *InlineResponse200382Mqtt) SetMessageFields(v []string)`

SetMessageFields sets MessageFields field to given value.

### HasMessageFields

`func (o *InlineResponse200382Mqtt) HasMessageFields() bool`

HasMessageFields returns a boolean if a field has been set.

### GetPublishing

`func (o *InlineResponse200382Mqtt) GetPublishing() InlineResponse200382MqttPublishing`

GetPublishing returns the Publishing field if non-nil, zero value otherwise.

### GetPublishingOk

`func (o *InlineResponse200382Mqtt) GetPublishingOk() (*InlineResponse200382MqttPublishing, bool)`

GetPublishingOk returns a tuple with the Publishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishing

`func (o *InlineResponse200382Mqtt) SetPublishing(v InlineResponse200382MqttPublishing)`

SetPublishing sets Publishing field to given value.

### HasPublishing

`func (o *InlineResponse200382Mqtt) HasPublishing() bool`

HasPublishing returns a boolean if a field has been set.

### GetBroker

`func (o *InlineResponse200382Mqtt) GetBroker() InlineResponse200382MqttBroker`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *InlineResponse200382Mqtt) GetBrokerOk() (*InlineResponse200382MqttBroker, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *InlineResponse200382Mqtt) SetBroker(v InlineResponse200382MqttBroker)`

SetBroker sets Broker field to given value.

### HasBroker

`func (o *InlineResponse200382Mqtt) HasBroker() bool`

HasBroker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


