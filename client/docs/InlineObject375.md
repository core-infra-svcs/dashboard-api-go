# InlineObject375

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**OrganizationsOrganizationIdWirelessMqttSettingsNetwork**](OrganizationsOrganizationIdWirelessMqttSettingsNetwork.md) |  | 
**Mqtt** | [**OrganizationsOrganizationIdWirelessMqttSettingsMqtt**](OrganizationsOrganizationIdWirelessMqttSettingsMqtt.md) |  | 
**Ble** | Pointer to [**InlineResponse200509Ble**](InlineResponse200509Ble.md) |  | [optional] 
**Wifi** | Pointer to [**InlineResponse200509Wifi**](InlineResponse200509Wifi.md) |  | [optional] 

## Methods

### NewInlineObject375

`func NewInlineObject375(network OrganizationsOrganizationIdWirelessMqttSettingsNetwork, mqtt OrganizationsOrganizationIdWirelessMqttSettingsMqtt, ) *InlineObject375`

NewInlineObject375 instantiates a new InlineObject375 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject375WithDefaults

`func NewInlineObject375WithDefaults() *InlineObject375`

NewInlineObject375WithDefaults instantiates a new InlineObject375 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineObject375) GetNetwork() OrganizationsOrganizationIdWirelessMqttSettingsNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineObject375) GetNetworkOk() (*OrganizationsOrganizationIdWirelessMqttSettingsNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineObject375) SetNetwork(v OrganizationsOrganizationIdWirelessMqttSettingsNetwork)`

SetNetwork sets Network field to given value.


### GetMqtt

`func (o *InlineObject375) GetMqtt() OrganizationsOrganizationIdWirelessMqttSettingsMqtt`

GetMqtt returns the Mqtt field if non-nil, zero value otherwise.

### GetMqttOk

`func (o *InlineObject375) GetMqttOk() (*OrganizationsOrganizationIdWirelessMqttSettingsMqtt, bool)`

GetMqttOk returns a tuple with the Mqtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqtt

`func (o *InlineObject375) SetMqtt(v OrganizationsOrganizationIdWirelessMqttSettingsMqtt)`

SetMqtt sets Mqtt field to given value.


### GetBle

`func (o *InlineObject375) GetBle() InlineResponse200509Ble`

GetBle returns the Ble field if non-nil, zero value otherwise.

### GetBleOk

`func (o *InlineObject375) GetBleOk() (*InlineResponse200509Ble, bool)`

GetBleOk returns a tuple with the Ble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBle

`func (o *InlineObject375) SetBle(v InlineResponse200509Ble)`

SetBle sets Ble field to given value.

### HasBle

`func (o *InlineObject375) HasBle() bool`

HasBle returns a boolean if a field has been set.

### GetWifi

`func (o *InlineObject375) GetWifi() InlineResponse200509Wifi`

GetWifi returns the Wifi field if non-nil, zero value otherwise.

### GetWifiOk

`func (o *InlineObject375) GetWifiOk() (*InlineResponse200509Wifi, bool)`

GetWifiOk returns a tuple with the Wifi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifi

`func (o *InlineObject375) SetWifi(v InlineResponse200509Wifi)`

SetWifi sets Wifi field to given value.

### HasWifi

`func (o *InlineObject375) HasWifi() bool`

HasWifi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


