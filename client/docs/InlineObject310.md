# InlineObject310

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**OrganizationsOrganizationIdWirelessMqttSettingsNetwork**](OrganizationsOrganizationIdWirelessMqttSettingsNetwork.md) |  | 
**Mqtt** | [**OrganizationsOrganizationIdWirelessMqttSettingsMqtt**](OrganizationsOrganizationIdWirelessMqttSettingsMqtt.md) |  | 
**Ble** | Pointer to [**InlineResponse200382Ble**](InlineResponse200382Ble.md) |  | [optional] 
**Wifi** | Pointer to [**InlineResponse200382Wifi**](InlineResponse200382Wifi.md) |  | [optional] 

## Methods

### NewInlineObject310

`func NewInlineObject310(network OrganizationsOrganizationIdWirelessMqttSettingsNetwork, mqtt OrganizationsOrganizationIdWirelessMqttSettingsMqtt, ) *InlineObject310`

NewInlineObject310 instantiates a new InlineObject310 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject310WithDefaults

`func NewInlineObject310WithDefaults() *InlineObject310`

NewInlineObject310WithDefaults instantiates a new InlineObject310 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineObject310) GetNetwork() OrganizationsOrganizationIdWirelessMqttSettingsNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineObject310) GetNetworkOk() (*OrganizationsOrganizationIdWirelessMqttSettingsNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineObject310) SetNetwork(v OrganizationsOrganizationIdWirelessMqttSettingsNetwork)`

SetNetwork sets Network field to given value.


### GetMqtt

`func (o *InlineObject310) GetMqtt() OrganizationsOrganizationIdWirelessMqttSettingsMqtt`

GetMqtt returns the Mqtt field if non-nil, zero value otherwise.

### GetMqttOk

`func (o *InlineObject310) GetMqttOk() (*OrganizationsOrganizationIdWirelessMqttSettingsMqtt, bool)`

GetMqttOk returns a tuple with the Mqtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqtt

`func (o *InlineObject310) SetMqtt(v OrganizationsOrganizationIdWirelessMqttSettingsMqtt)`

SetMqtt sets Mqtt field to given value.


### GetBle

`func (o *InlineObject310) GetBle() InlineResponse200382Ble`

GetBle returns the Ble field if non-nil, zero value otherwise.

### GetBleOk

`func (o *InlineObject310) GetBleOk() (*InlineResponse200382Ble, bool)`

GetBleOk returns a tuple with the Ble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBle

`func (o *InlineObject310) SetBle(v InlineResponse200382Ble)`

SetBle sets Ble field to given value.

### HasBle

`func (o *InlineObject310) HasBle() bool`

HasBle returns a boolean if a field has been set.

### GetWifi

`func (o *InlineObject310) GetWifi() InlineResponse200382Wifi`

GetWifi returns the Wifi field if non-nil, zero value otherwise.

### GetWifiOk

`func (o *InlineObject310) GetWifiOk() (*InlineResponse200382Wifi, bool)`

GetWifiOk returns a tuple with the Wifi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifi

`func (o *InlineObject310) SetWifi(v InlineResponse200382Wifi)`

SetWifi sets Wifi field to given value.

### HasWifi

`func (o *InlineObject310) HasWifi() bool`

HasWifi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


