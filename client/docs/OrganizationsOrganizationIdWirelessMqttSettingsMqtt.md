# OrganizationsOrganizationIdWirelessMqttSettingsMqtt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | MQTT Enabled | 
**Topic** | Pointer to **string** | MQTT Topic | [optional] 
**MessageFields** | Pointer to **[]string** | Select fields to populate in MQTT messages. Valid types are: RSSI, AP MAC address, Client MAC address, Timestamp, Radio, Network ID, Beacon type, Raw payload, Client UUID, Client major value, Client minor value, Signal power, Band, Slot ID | [optional] 
**Publishing** | Pointer to [**InlineResponse200382MqttPublishing**](InlineResponse200382MqttPublishing.md) |  | [optional] 
**Broker** | Pointer to [**OrganizationsOrganizationIdWirelessMqttSettingsMqttBroker**](OrganizationsOrganizationIdWirelessMqttSettingsMqttBroker.md) |  | [optional] 

## Methods

### NewOrganizationsOrganizationIdWirelessMqttSettingsMqtt

`func NewOrganizationsOrganizationIdWirelessMqttSettingsMqtt(enabled bool, ) *OrganizationsOrganizationIdWirelessMqttSettingsMqtt`

NewOrganizationsOrganizationIdWirelessMqttSettingsMqtt instantiates a new OrganizationsOrganizationIdWirelessMqttSettingsMqtt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdWirelessMqttSettingsMqttWithDefaults

`func NewOrganizationsOrganizationIdWirelessMqttSettingsMqttWithDefaults() *OrganizationsOrganizationIdWirelessMqttSettingsMqtt`

NewOrganizationsOrganizationIdWirelessMqttSettingsMqttWithDefaults instantiates a new OrganizationsOrganizationIdWirelessMqttSettingsMqtt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetTopic

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetMessageFields

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) GetMessageFields() []string`

GetMessageFields returns the MessageFields field if non-nil, zero value otherwise.

### GetMessageFieldsOk

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) GetMessageFieldsOk() (*[]string, bool)`

GetMessageFieldsOk returns a tuple with the MessageFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageFields

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) SetMessageFields(v []string)`

SetMessageFields sets MessageFields field to given value.

### HasMessageFields

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) HasMessageFields() bool`

HasMessageFields returns a boolean if a field has been set.

### GetPublishing

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) GetPublishing() InlineResponse200382MqttPublishing`

GetPublishing returns the Publishing field if non-nil, zero value otherwise.

### GetPublishingOk

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) GetPublishingOk() (*InlineResponse200382MqttPublishing, bool)`

GetPublishingOk returns a tuple with the Publishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishing

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) SetPublishing(v InlineResponse200382MqttPublishing)`

SetPublishing sets Publishing field to given value.

### HasPublishing

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) HasPublishing() bool`

HasPublishing returns a boolean if a field has been set.

### GetBroker

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) GetBroker() OrganizationsOrganizationIdWirelessMqttSettingsMqttBroker`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) GetBrokerOk() (*OrganizationsOrganizationIdWirelessMqttSettingsMqttBroker, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) SetBroker(v OrganizationsOrganizationIdWirelessMqttSettingsMqttBroker)`

SetBroker sets Broker field to given value.

### HasBroker

`func (o *OrganizationsOrganizationIdWirelessMqttSettingsMqtt) HasBroker() bool`

HasBroker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


