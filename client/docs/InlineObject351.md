# InlineObject351

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**OrganizationsOrganizationIdWirelessMqttSettingsNetwork**](OrganizationsOrganizationIdWirelessMqttSettingsNetwork.md) |  | 
**Mqtt** | [**OrganizationsOrganizationIdWirelessMqttSettingsMqtt**](OrganizationsOrganizationIdWirelessMqttSettingsMqtt.md) |  | 
**Ble** | Pointer to [**InlineResponse200463Ble**](InlineResponse200463Ble.md) |  | [optional] 
**Wifi** | Pointer to [**InlineResponse200463Wifi**](InlineResponse200463Wifi.md) |  | [optional] 

## Methods

### NewInlineObject351

`func NewInlineObject351(network OrganizationsOrganizationIdWirelessMqttSettingsNetwork, mqtt OrganizationsOrganizationIdWirelessMqttSettingsMqtt, ) *InlineObject351`

NewInlineObject351 instantiates a new InlineObject351 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject351WithDefaults

`func NewInlineObject351WithDefaults() *InlineObject351`

NewInlineObject351WithDefaults instantiates a new InlineObject351 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineObject351) GetNetwork() OrganizationsOrganizationIdWirelessMqttSettingsNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineObject351) GetNetworkOk() (*OrganizationsOrganizationIdWirelessMqttSettingsNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineObject351) SetNetwork(v OrganizationsOrganizationIdWirelessMqttSettingsNetwork)`

SetNetwork sets Network field to given value.


### GetMqtt

`func (o *InlineObject351) GetMqtt() OrganizationsOrganizationIdWirelessMqttSettingsMqtt`

GetMqtt returns the Mqtt field if non-nil, zero value otherwise.

### GetMqttOk

`func (o *InlineObject351) GetMqttOk() (*OrganizationsOrganizationIdWirelessMqttSettingsMqtt, bool)`

GetMqttOk returns a tuple with the Mqtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqtt

`func (o *InlineObject351) SetMqtt(v OrganizationsOrganizationIdWirelessMqttSettingsMqtt)`

SetMqtt sets Mqtt field to given value.


### GetBle

`func (o *InlineObject351) GetBle() InlineResponse200463Ble`

GetBle returns the Ble field if non-nil, zero value otherwise.

### GetBleOk

`func (o *InlineObject351) GetBleOk() (*InlineResponse200463Ble, bool)`

GetBleOk returns a tuple with the Ble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBle

`func (o *InlineObject351) SetBle(v InlineResponse200463Ble)`

SetBle sets Ble field to given value.

### HasBle

`func (o *InlineObject351) HasBle() bool`

HasBle returns a boolean if a field has been set.

### GetWifi

`func (o *InlineObject351) GetWifi() InlineResponse200463Wifi`

GetWifi returns the Wifi field if non-nil, zero value otherwise.

### GetWifiOk

`func (o *InlineObject351) GetWifiOk() (*InlineResponse200463Wifi, bool)`

GetWifiOk returns a tuple with the Wifi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifi

`func (o *InlineObject351) SetWifi(v InlineResponse200463Wifi)`

SetWifi sets Wifi field to given value.

### HasWifi

`func (o *InlineObject351) HasWifi() bool`

HasWifi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


