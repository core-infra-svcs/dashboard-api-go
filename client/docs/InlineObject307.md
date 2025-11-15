# InlineObject307

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**OrganizationsOrganizationIdWirelessMqttSettingsNetwork**](OrganizationsOrganizationIdWirelessMqttSettingsNetwork.md) |  | 
**Mqtt** | [**OrganizationsOrganizationIdWirelessMqttSettingsMqtt**](OrganizationsOrganizationIdWirelessMqttSettingsMqtt.md) |  | 
**Ble** | Pointer to [**InlineResponse200373Ble**](InlineResponse200373Ble.md) |  | [optional] 
**Wifi** | Pointer to [**InlineResponse200373Wifi**](InlineResponse200373Wifi.md) |  | [optional] 

## Methods

### NewInlineObject307

`func NewInlineObject307(network OrganizationsOrganizationIdWirelessMqttSettingsNetwork, mqtt OrganizationsOrganizationIdWirelessMqttSettingsMqtt, ) *InlineObject307`

NewInlineObject307 instantiates a new InlineObject307 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject307WithDefaults

`func NewInlineObject307WithDefaults() *InlineObject307`

NewInlineObject307WithDefaults instantiates a new InlineObject307 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineObject307) GetNetwork() OrganizationsOrganizationIdWirelessMqttSettingsNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineObject307) GetNetworkOk() (*OrganizationsOrganizationIdWirelessMqttSettingsNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineObject307) SetNetwork(v OrganizationsOrganizationIdWirelessMqttSettingsNetwork)`

SetNetwork sets Network field to given value.


### GetMqtt

`func (o *InlineObject307) GetMqtt() OrganizationsOrganizationIdWirelessMqttSettingsMqtt`

GetMqtt returns the Mqtt field if non-nil, zero value otherwise.

### GetMqttOk

`func (o *InlineObject307) GetMqttOk() (*OrganizationsOrganizationIdWirelessMqttSettingsMqtt, bool)`

GetMqttOk returns a tuple with the Mqtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqtt

`func (o *InlineObject307) SetMqtt(v OrganizationsOrganizationIdWirelessMqttSettingsMqtt)`

SetMqtt sets Mqtt field to given value.


### GetBle

`func (o *InlineObject307) GetBle() InlineResponse200373Ble`

GetBle returns the Ble field if non-nil, zero value otherwise.

### GetBleOk

`func (o *InlineObject307) GetBleOk() (*InlineResponse200373Ble, bool)`

GetBleOk returns a tuple with the Ble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBle

`func (o *InlineObject307) SetBle(v InlineResponse200373Ble)`

SetBle sets Ble field to given value.

### HasBle

`func (o *InlineObject307) HasBle() bool`

HasBle returns a boolean if a field has been set.

### GetWifi

`func (o *InlineObject307) GetWifi() InlineResponse200373Wifi`

GetWifi returns the Wifi field if non-nil, zero value otherwise.

### GetWifiOk

`func (o *InlineObject307) GetWifiOk() (*InlineResponse200373Wifi, bool)`

GetWifiOk returns a tuple with the Wifi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifi

`func (o *InlineObject307) SetWifi(v InlineResponse200373Wifi)`

SetWifi sets Wifi field to given value.

### HasWifi

`func (o *InlineObject307) HasWifi() bool`

HasWifi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


