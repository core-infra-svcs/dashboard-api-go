# InlineObject338

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**OrganizationsOrganizationIdWirelessMqttSettingsNetwork**](OrganizationsOrganizationIdWirelessMqttSettingsNetwork.md) |  | 
**Mqtt** | [**OrganizationsOrganizationIdWirelessMqttSettingsMqtt**](OrganizationsOrganizationIdWirelessMqttSettingsMqtt.md) |  | 
**Ble** | Pointer to [**InlineResponse200448Ble**](InlineResponse200448Ble.md) |  | [optional] 
**Wifi** | Pointer to [**InlineResponse200448Wifi**](InlineResponse200448Wifi.md) |  | [optional] 

## Methods

### NewInlineObject338

`func NewInlineObject338(network OrganizationsOrganizationIdWirelessMqttSettingsNetwork, mqtt OrganizationsOrganizationIdWirelessMqttSettingsMqtt, ) *InlineObject338`

NewInlineObject338 instantiates a new InlineObject338 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject338WithDefaults

`func NewInlineObject338WithDefaults() *InlineObject338`

NewInlineObject338WithDefaults instantiates a new InlineObject338 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineObject338) GetNetwork() OrganizationsOrganizationIdWirelessMqttSettingsNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineObject338) GetNetworkOk() (*OrganizationsOrganizationIdWirelessMqttSettingsNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineObject338) SetNetwork(v OrganizationsOrganizationIdWirelessMqttSettingsNetwork)`

SetNetwork sets Network field to given value.


### GetMqtt

`func (o *InlineObject338) GetMqtt() OrganizationsOrganizationIdWirelessMqttSettingsMqtt`

GetMqtt returns the Mqtt field if non-nil, zero value otherwise.

### GetMqttOk

`func (o *InlineObject338) GetMqttOk() (*OrganizationsOrganizationIdWirelessMqttSettingsMqtt, bool)`

GetMqttOk returns a tuple with the Mqtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqtt

`func (o *InlineObject338) SetMqtt(v OrganizationsOrganizationIdWirelessMqttSettingsMqtt)`

SetMqtt sets Mqtt field to given value.


### GetBle

`func (o *InlineObject338) GetBle() InlineResponse200448Ble`

GetBle returns the Ble field if non-nil, zero value otherwise.

### GetBleOk

`func (o *InlineObject338) GetBleOk() (*InlineResponse200448Ble, bool)`

GetBleOk returns a tuple with the Ble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBle

`func (o *InlineObject338) SetBle(v InlineResponse200448Ble)`

SetBle sets Ble field to given value.

### HasBle

`func (o *InlineObject338) HasBle() bool`

HasBle returns a boolean if a field has been set.

### GetWifi

`func (o *InlineObject338) GetWifi() InlineResponse200448Wifi`

GetWifi returns the Wifi field if non-nil, zero value otherwise.

### GetWifiOk

`func (o *InlineObject338) GetWifiOk() (*InlineResponse200448Wifi, bool)`

GetWifiOk returns a tuple with the Wifi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifi

`func (o *InlineObject338) SetWifi(v InlineResponse200448Wifi)`

SetWifi sets Wifi field to given value.

### HasWifi

`func (o *InlineObject338) HasWifi() bool`

HasWifi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


