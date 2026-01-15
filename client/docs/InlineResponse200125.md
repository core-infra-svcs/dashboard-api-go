# InlineResponse200125

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MqttBrokerId** | Pointer to **string** | ID of the MQTT Broker. | [optional] 
**Enabled** | Pointer to **bool** | Specifies whether the broker is enabled for sensor data. Currently, only a single broker may be enabled for sensor data. | [optional] 

## Methods

### NewInlineResponse200125

`func NewInlineResponse200125() *InlineResponse200125`

NewInlineResponse200125 instantiates a new InlineResponse200125 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200125WithDefaults

`func NewInlineResponse200125WithDefaults() *InlineResponse200125`

NewInlineResponse200125WithDefaults instantiates a new InlineResponse200125 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMqttBrokerId

`func (o *InlineResponse200125) GetMqttBrokerId() string`

GetMqttBrokerId returns the MqttBrokerId field if non-nil, zero value otherwise.

### GetMqttBrokerIdOk

`func (o *InlineResponse200125) GetMqttBrokerIdOk() (*string, bool)`

GetMqttBrokerIdOk returns a tuple with the MqttBrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttBrokerId

`func (o *InlineResponse200125) SetMqttBrokerId(v string)`

SetMqttBrokerId sets MqttBrokerId field to given value.

### HasMqttBrokerId

`func (o *InlineResponse200125) HasMqttBrokerId() bool`

HasMqttBrokerId returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200125) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200125) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200125) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200125) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


